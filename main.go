package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/AshokPabra/observability_assignment/app"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel/attribute"
)

func main() {

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

	http.Handle("/users", otelhttp.NewHandler(
		http.HandlerFunc(app.GetUserHandler),
		"/users",
		otelhttp.WithMetricAttributesFn(func(r *http.Request) []attribute.KeyValue {
			return []attribute.KeyValue{attribute.String("http.route", "/users")}
		}),
	))
	http.Handle("/user", otelhttp.NewHandler(
		http.HandlerFunc(app.CreateUserHandler),
		"/user",
		otelhttp.WithMetricAttributesFn(func(r *http.Request) []attribute.KeyValue {
			return []attribute.KeyValue{attribute.String("http.route", "/user")}
		}),
	))
	http.Handle("/delete", otelhttp.NewHandler(
		http.HandlerFunc(app.DeleteUserHandler),
		"/delete",
		otelhttp.WithMetricAttributesFn(func(r *http.Request) []attribute.KeyValue {
			return []attribute.KeyValue{attribute.String("http.route", "/delete")}
		}),
	))

	http.ListenAndServe(":8080", nil)
}
