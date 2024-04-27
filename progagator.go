package ote

import (
	"github.com/ZenLiuCN/ote/common"
	"go.opentelemetry.io/otel/propagation"
)

func NewPropagator(conf common.Config) propagation.TextMapPropagator {
	return propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	)
}
