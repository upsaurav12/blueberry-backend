package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	UserID        string `json:"id"`
	UserFirstname string `json:"firstname"`
	UserLastname  string `json:"lastname"`
	Summary       string `json:"summary"`
	Email         string `json:"email"`
	Location      string `json:"location"`
	Role          string `json:"role"`
	PhoneNumber   string `json:"phone"`
	Status        string `json:"status"`
	LinkedIn      string `json:"linkedin"`
}

var users []User // This should be a slice of User

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users) // Fix to encode 'users' slice
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			// Handle preflight request
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r) // Continue with request handling
	})
}

func main() {
	r := mux.NewRouter() // Adding a route for getUsers
	r.Use(corsMiddleware)

	// Pre-fill some users for testing
	users = append(users, User{UserID: "1", UserFirstname: "John", UserLastname: "Doe", Summary: "A summary about John.", Email: "john.doe@example.com", Location: "USA", Role: "Student", PhoneNumber: "123-456-7890", Status: "Active", LinkedIn: "linkedin.com/in/johndoe"})

	r.HandleFunc("/api/users", getUsers).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r)) // Start the server on port 8000
}
