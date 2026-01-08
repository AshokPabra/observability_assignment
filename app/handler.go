package app

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/AshokPabra/observability_assignment/logger"
	"go.uber.org/zap"
)

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

	err := json.NewEncoder(w).Encode(users)
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
