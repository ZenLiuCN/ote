package ote

import (
	"context"
	"fmt"
	"github.com/ZenLiuCN/ote/common"
	rt "go.opentelemetry.io/contrib/instrumentation/runtime"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
	"runtime"
	"time"
)

const (
	Version = "0.0.1"
)
const ContextKey = "$telemetry"

type (
	TelemetryProviderFn func() (scope string, opt []trace.SpanStartOption)
	SpanProviderFn      func() (name string, attr []attribute.KeyValue)
)

func caller() string {
	pc, _, _, ok := runtime.Caller(1) // 0 self, 1 ByContext
	if ok {
		return fmt.Sprintf("%s", runtime.FuncForPC(pc).Name())
	}
	return "unknown"
}

// FromContext fetch context Telemetry
func FromContext(ctx context.Context) (r Telemetry) {
	if ctx == nil {
		return nil
	}
	r, ok := ctx.Value(ContextKey).(Telemetry)
	if !ok {
		return nil
	}
	return r
}

// ByContext fetch or create Telemetry from context
func ByContext(ctx context.Context, p TelemetryProviderFn) (r Telemetry, cx context.Context) {
	if ctx == nil {
		return nil, nil
	}
	r, ok := ctx.Value(ContextKey).(Telemetry)
	if !ok {
		if p != nil {
			s, o := p()
			r = NewTelemetry(s, o...)
		} else {
			r = NewTelemetry(caller())
		}
		ctx = context.WithValue(ctx, ContextKey, r)
	}
	return r, ctx
}

// SpanByContext create span with context Telemetry or create new one
func SpanByContext(ctx context.Context, p TelemetryProviderFn, sn SpanProviderFn) (te Telemetry, sp trace.Span, cx context.Context) {
	if ctx == nil {
		return nil, nil, nil
	}
	var ok bool
	te, ok = ctx.Value(ContextKey).(Telemetry)
	if !ok {
		if p != nil {
			s, o := p()
			te = NewTelemetry(s, o...)
		} else {
			te = NewTelemetry(caller())
		}
		ctx = context.WithValue(ctx, ContextKey, te)
	}
	if sn != nil {
		n, a := sn()
		cx, sp = te.StartSpan(n, ctx, a...)
	} else {
		cx, sp = te.StartSpan(caller(), ctx)
	}
	return
}

// SpanFromContext create span only context have a Telemetry
func SpanFromContext(ctx context.Context, sn SpanProviderFn) (te Telemetry, sp trace.Span, cx context.Context) {
	if ctx == nil {
		return nil, nil, nil
	}
	var ok bool
	te, ok = ctx.Value(ContextKey).(Telemetry)
	if !ok {
		return nil, nil, ctx
	}
	if sn != nil {
		n, a := sn()
		cx, sp = te.StartSpan(n, ctx, a...)
	} else {
		cx, sp = te.StartSpan(caller(), ctx)
	}
	return
}

type Telemetry interface {
	HandleError(err error)
	/*
		HandleRecover returns the recover value and if there is failure
		A sample usage like:
			defer func() {
				if r, ok := tel.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
	*/
	HandleRecover(rec any) (failure any, failed bool)

	//StartSpan if ctx is nil then the returns are nil
	StartSpan(name string, ctx context.Context, attrs ...attribute.KeyValue) (context.Context, trace.Span)

	SetContext(ctx context.Context) context.Context
}
type telemetry struct {
	spanStartOption []trace.SpanStartOption
	propagator      propagation.TextMapPropagator
	meter           metric.Meter
	tracer          trace.Tracer
}

func (t *telemetry) StartSpan(name string, ctx context.Context, attrs ...attribute.KeyValue) (cx context.Context, sp trace.Span) {
	if ctx == nil {
		return nil, nil
	}
	cx, sp = t.tracer.Start(ctx, name, t.spanStartOption...)
	if len(attrs) > 0 {
		sp.SetAttributes(attrs...)
	}
	return
}
func (t *telemetry) HandleError(err error) {
	if err != nil {
		otel.Handle(err)
	}
}
func (t *telemetry) HandleRecover(rec any) (any, bool) {
	switch r := rec.(type) {
	case nil:
		return r, false
	case error:
		t.HandleError(r)
		return r, true
	default:
		t.HandleError(fmt.Errorf("%#+v", r))
		return r, true
	}
}

func (t *telemetry) SetContext(ctx context.Context) context.Context {
	if ctx == nil {
		return nil
	}
	return context.WithValue(ctx, ContextKey, t)
}
func NewTelemetry(scope string, opts ...trace.SpanStartOption) Telemetry {
	return &telemetry{
		spanStartOption: opts,
		propagator:      otel.GetTextMapPropagator(),
		meter:           otel.GetMeterProvider().Meter(scope, metric.WithInstrumentationVersion(Version)),
		tracer:          otel.GetTracerProvider().Tracer(scope, trace.WithInstrumentationVersion(Version)),
	}
}

type ServiceFunc func(ctx context.Context)

// Instrument always start span except context is nil
func Instrument(pn TelemetryProviderFn, sn SpanProviderFn, service ServiceFunc) ServiceFunc {
	return func(ctx context.Context) {
		if te, sp, cx := SpanByContext(ctx, pn, sn); sp != nil {
			defer func() {
				if r, ok := te.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			defer sp.End()
			service(cx)
		} else {
			service(ctx)
		}

	}
}

// InstrumentOption only start span when context have Telemetry
func InstrumentOption(sn SpanProviderFn, service ServiceFunc) ServiceFunc {
	return func(ctx context.Context) {
		if te, sp, cx := SpanFromContext(ctx, sn); sp != nil {
			defer func() {
				if r, ok := te.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			defer sp.End()
			service(cx)
		} else {
			service(ctx)
		}

	}
}

// RuntimeInstrument inject runtime Telemetry
func RuntimeInstrument(c common.Config) {
	if c == nil {
		Handle(rt.Start(rt.WithMinimumReadMemStatsInterval(time.Second)))
	} else {
		Handle(rt.Start(rt.WithMinimumReadMemStatsInterval(c.GetTimeDurationInfiniteNotAllowed("telemetry.runtime.interval", time.Second))))
	}

}
func Handle(err error) {
	if err != nil {
		otel.Handle(err)
	}
}
