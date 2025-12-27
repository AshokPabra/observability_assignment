package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/AshokPabra/observability_assignment/app"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

func main() {

	tp, err := initTracer()
	if err != nil {
		log.Fatal("failed to initialize tracer:", err)
	}
	defer tp.Shutdown(context.Background())

	fmt.Println("server is started on port :8080")

	http.Handle("/users", otelhttp.NewHandler(http.HandlerFunc(app.GetUserHandler), "/users"))
	http.Handle("/user", otelhttp.NewHandler(http.HandlerFunc(app.CreateUserHandler), "/user"))
	http.Handle("/delete", otelhttp.NewHandler(http.HandlerFunc(app.DeleteUserHandler), "/delete"))

	http.ListenAndServe(":8080", nil)
}
