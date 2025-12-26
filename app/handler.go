package app

import (
	"encoding/json"
	"fmt"
	"net/http"
)

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

	err := json.NewEncoder(w).Encode(users)
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
