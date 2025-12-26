package main

import (
	"fmt"
	"net/http"

	"github.com/AshokPabra/observability_assignment/app"
)

func main() {
	fmt.Println("server is started on port :8080")
	http.HandleFunc("/users", app.GetUserHandler)
	http.HandleFunc("/user", app.CreateUserHandler)
	http.HandleFunc("/delete", app.DeleteUserHandler)
	http.ListenAndServe(":8080", nil)
}
