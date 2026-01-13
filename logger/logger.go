package logger

import (
	"context"
	"os"

	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *zap.Logger

func Init() {
	// Create encoder config for JSON output
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	// Create core that writes only to stdout
	// Docker will capture stdout and write to its JSON log file
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.AddSync(os.Stdout),
		zap.InfoLevel,
	)

	Log = zap.New(core)
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
