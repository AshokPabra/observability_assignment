package app

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/AshokPabra/observability_assignment/logger"
	"go.uber.org/zap"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

var tracer = otel.Tracer("user-service/app")

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var users []User

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	log := logger.WithTraceContext(r.Context())
	log.Info("CreateUserHandler started")

	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "invalid user", http.StatusBadRequest)
		return
	}
	users = append(users, user)

	log.Info("user created successfully", zap.Int("user_id", user.Id), zap.String("user_name", user.Name))

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "user created successfully \n")
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	log := logger.WithTraceContext(r.Context())

	log.Info("getUserHandler started ")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	span := trace.SpanFromContext(r.Context())

	traceId := span.SpanContext().TraceID().String()

	spanId := span.SpanContext().SpanID().String()

	ctx1 := context.WithValue(r.Context(), "traceId", traceId)

	ctx2 := context.WithValue(ctx1, "spanId", spanId)
	users, err := getUserList(ctx2)

	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		log.Error("failed to encode users", zap.Error(err))
		http.Error(w, "error in getting users list", http.StatusInternalServerError)
		return
	}

	log.Info("users fetched successfully !! ", zap.Int("user_count", len(users)))
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	log := logger.WithTraceContext(r.Context())
	log.Info("DeleteUserHandler started")
	var user User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		log.Error("failed to decode user ", zap.Error(err))
		http.Error(w, "invalid user", http.StatusBadRequest)
		return
	}

	for i, v := range users {
		if v.Id == user.Id {
			users[i] = User{}
			log.Info("user deleted", zap.Int("user_id", user.Id))
		}
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "user deleted successfully \n")
}

func getUserList(ctx context.Context) ([]User, error) {

	ctx, span := tracer.Start(ctx, "getUserList")
	traceId := ctx.Value("traceId").(string)
	spanId := ctx.Value("spanId").(string)

	spanIdpresent := span.SpanContext().SpanID().String()

	defer span.End()
	fmt.Println()
	fmt.Printf("traceId: %s, parent-spanId: %s, spanId: %s", traceId, spanId, spanIdpresent)
	fmt.Println()

	sleepfunc(ctx)

	list_of_users := users
	span.SetAttributes(attribute.Int("user.count", len(list_of_users)))
	return list_of_users, nil
}

func sleepfunc(ctx context.Context) {
	ctx, span := tracer.Start(ctx, "getUserList")
	defer span.End()
	time.Sleep(5 * time.Microsecond)
}
