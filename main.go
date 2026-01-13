package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/AshokPabra/observability_assignment/app"
	"github.com/AshokPabra/observability_assignment/logger"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel/attribute"
)

func main() {

	logger.Init()
	defer logger.Sync()

	tp, err := initTracer()
	if err != nil {
		log.Fatal("failed to initialize tracer:", err)
	}
	defer tp.Shutdown(context.Background())

	mp, err := initMeter()
	if err != nil {
		log.Fatal("failed to initialize meter:", err)
	}
	defer mp.Shutdown(context.Background())

	fmt.Println("server is started on port :8080")

	// Health check endpoint - no API key required
	http.Handle("/health", otelhttp.NewHandler(
		http.HandlerFunc(app.HealthCheckHandler),
		"/health",
		otelhttp.WithMetricAttributesFn(func(r *http.Request) []attribute.KeyValue {
			return []attribute.KeyValue{attribute.String("http.route", "/health")}
		}),
	))

	// Protected endpoints - API key required
	http.Handle("/users", app.APIKeyMiddleware(otelhttp.NewHandler(
		http.HandlerFunc(app.GetUserHandler),
		"/users",
		otelhttp.WithMetricAttributesFn(func(r *http.Request) []attribute.KeyValue {
			return []attribute.KeyValue{attribute.String("http.route", "/users")}
		}),
	)))
	http.Handle("/user", app.APIKeyMiddleware(otelhttp.NewHandler(
		http.HandlerFunc(app.CreateUserHandler),
		"/user",
		otelhttp.WithMetricAttributesFn(func(r *http.Request) []attribute.KeyValue {
			return []attribute.KeyValue{attribute.String("http.route", "/user")}
		}),
	)))
	http.Handle("/delete", app.APIKeyMiddleware(otelhttp.NewHandler(
		http.HandlerFunc(app.DeleteUserHandler),
		"/delete",
		otelhttp.WithMetricAttributesFn(func(r *http.Request) []attribute.KeyValue {
			return []attribute.KeyValue{attribute.String("http.route", "/delete")}
		}),
	)))
	http.Handle("/search", app.APIKeyMiddleware(otelhttp.NewHandler(
		http.HandlerFunc(app.SearchUserByEmailHandler),
		"/search",
		otelhttp.WithMetricAttributesFn(func(r *http.Request) []attribute.KeyValue {
			return []attribute.KeyValue{attribute.String("http.route", "/search")}
		}),
	)))

	http.ListenAndServe(":8080", nil)
}
