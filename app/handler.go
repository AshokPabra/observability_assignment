package app

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

var tracer = otel.Tracer("user-service/app")

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var users []User

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "invalid user", http.StatusBadRequest)
		return
	}
	users = append(users, user)

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "user created successfully \n")
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	users, err := getUserList(r.Context())

	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		http.Error(w, "error in getting users list", http.StatusInternalServerError)
		return
	}
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	var user User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, "invalid user", http.StatusBadRequest)
		return
	}

	for i, v := range users {
		if v.Id == user.Id {
			users[i] = User{}
		}
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "user deleted successfully \n")
}

func getUserList(ctx context.Context) ([]User, error) {

	ctx, span := tracer.Start(ctx, "getUserList")
	defer span.End()

	list_of_users := users
	span.SetAttributes(attribute.Int("user.count", len(list_of_users)))
	return list_of_users, nil
}
