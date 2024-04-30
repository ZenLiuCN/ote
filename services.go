package ote

import (
	"context"
	"fmt"
	rt "go.opentelemetry.io/contrib/instrumentation/runtime"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
	"net/http"
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

// ByContext fetch or create Telemetry from context
func ByContext(ctx context.Context, p TelemetryProviderFn) (r Telemetry, cx context.Context) {
	if ctx == nil {
		return nil, nil
	}
	if !HaveTelemetry() {
		return nil, ctx
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
func ByRequestContext(req *http.Request, p TelemetryProviderFn) (r Telemetry, rx *http.Request) {
	var cx context.Context
	r, cx = ByContext(req.Context(), p)
	if cx != nil {
		return r, req.WithContext(cx)
	}
	return nil, req
}

// SpanByContext create span with context Telemetry or create new one
func SpanByContext(ctx context.Context, p TelemetryProviderFn, sn SpanProviderFn) (te Telemetry, sp trace.Span, cx context.Context) {
	if ctx == nil {
		return nil, nil, nil
	}
	if !HaveTelemetry() {
		return nil, nil, ctx
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
func SpanByRequestContext(req *http.Request, p TelemetryProviderFn, sn SpanProviderFn) (te Telemetry, sp trace.Span, rx *http.Request) {
	var cx context.Context
	te, sp, cx = SpanByContext(req.Context(), p, sn)
	if te != nil {
		rx = req.WithContext(cx)
		return
	}
	return nil, nil, req
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

// NewTelemetry create Telemetry returns nil if not setup telemetry
func NewTelemetry(scope string, opts ...trace.SpanStartOption) Telemetry {
	if shutdown == nil {
		return nil
	}
	return &telemetry{
		spanStartOption: opts,
		propagator:      otel.GetTextMapPropagator(),
		meter:           otel.GetMeterProvider().Meter(scope, metric.WithInstrumentationVersion(Version)),
		tracer:          otel.GetTracerProvider().Tracer(scope, trace.WithInstrumentationVersion(Version)),
	}
}

// RuntimeInstrument inject runtime Telemetry
func RuntimeInstrument(interval time.Duration) {
	if interval == 0 {
		Handle(rt.Start(rt.WithMinimumReadMemStatsInterval(time.Second)))
	} else {
		Handle(rt.Start(rt.WithMinimumReadMemStatsInterval(interval)))
	}

}
func Handle(err error) {
	if err != nil {
		otel.Handle(err)
	}
}
