package ote

import (
	"context"
	"errors"
	"github.com/ZenLiuCN/ote/otlp"
	"github.com/ZenLiuCN/ote/prometheus"

	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/trace"

	"go.opentelemetry.io/otel"
)

var (
	shutdown func(context.Context) error
)

func HaveTelemetry() bool {
	return shutdown != nil
}
func SetupTelemetry(ctx context.Context, conf *otlp.TraceConfig) (s func(context.Context) error, err error) {
	if HaveTelemetry() {
		return shutdown, nil
	}
	var shutdownFunc []func(context.Context) error

	shutdown = func(ctx context.Context) error {
		var err error
		for _, fn := range shutdownFunc {
			err = errors.Join(err, fn(ctx))
		}
		shutdownFunc = nil
		return err
	}
	s = shutdown
	handleErr := func(inErr error) {
		err = errors.Join(inErr, shutdown(ctx))
	}
	prop := NewPropagator()
	otel.SetTextMapPropagator(prop)
	var tracerProvider *trace.TracerProvider
	if tracerProvider, err = otlp.NewTraceProvider(ctx, conf); err != nil {
		handleErr(err)
		return
	} else {
		shutdownFunc = append(shutdownFunc, tracerProvider.Shutdown)
		otel.SetTracerProvider(tracerProvider)
	}
	var meterProvider *metric.MeterProvider
	if meterProvider, err = prometheus.NewMeterProvider(ctx, conf.ResourceConfig); err != nil {
		handleErr(err)
		return
	} else {
		shutdownFunc = append(shutdownFunc, meterProvider.Shutdown)
		otel.SetMeterProvider(meterProvider)
	}

	return
}
