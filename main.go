package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"crud_api/database"
	"crud_api/handlers"
)

func main() {
	database.Init()

	database.CreateUsersTable()

	r := mux.NewRouter()

	r.HandleFunc("/users", handlers.CreateUserHandler).Methods("POST")
	r.HandleFunc("/users/{id}", handlers.GetUserByIDHandler).Methods("GET")
	r.HandleFunc("/users/{id}", handlers.UpdateUserHandler).Methods("PUT")
	r.HandleFunc("/users/{id}", handlers.DeleteUserHandler).Methods("DELETE")

	log.Println("Successfully started.")
	log.Fatal(http.ListenAndServe(":8080", r))
}
