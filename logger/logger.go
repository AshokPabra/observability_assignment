package logger

import (
	"context"

	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

var Log *zap.Logger

func Init() {
	var err error
	Log, err = zap.NewProduction()
	if err != nil {
		panic(err)
	}
}

// WithTraceContext returns a logger with trace_id and span_id
func WithTraceContext(ctx context.Context) *zap.Logger {
	span := trace.SpanFromContext(ctx)
	spanCtx := span.SpanContext()

	return Log.With(
		zap.String("trace_id", spanCtx.TraceID().String()),
		zap.String("span_id", spanCtx.SpanID().String()),
	)
}

// Sync flushes the logger
func Sync() {
	Log.Sync()
}
