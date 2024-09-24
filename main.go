package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func init() {
	users = []User{
		{
			ID:        1,
			FirstName: "Paquita",
			LastName:  "Perez",
			Email:     "pperez@gmail.com",
		}, {
			ID:        2,
			FirstName: "Andres",
			LastName:  "Roman",
			Email:     "aroman@gmail.com",
		}, {
			ID:        3,
			FirstName: "Juana",
			LastName:  "Lopez",
			Email:     "jlopez@gmail.com",
		},
	}
}

func main() {

	http.HandleFunc("/users", UserServer)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

type User struct {
	ID        uint64 `json: "id"`
	FirstName string `json: "first_name"`
	LastName  string `json: "last_name"`
	Email     string `json: "email"`
}

var users []User

func UserServer(w http.ResponseWriter, r *http.Request) {
	var status int
	switch r.Method {
	case http.MethodGet:
		GetAllUsers(w)
	case http.MethodPost:
		status = 200
		w.WriteHeader(status)
		fmt.Fprintf(w, `{"status": %d, "message": "%s"}`, status, "success in post")
	default:
		status = 404
		w.WriteHeader(status)
		fmt.Fprintf(w, `{"status": %d, "message": "%s"}`, status, "not found")
	}
}

func GetAllUsers(w http.ResponseWriter) {
	DataResponse(w, http.StatusOK, users)
}

func DataResponse(w http.ResponseWriter, status int, users interface{}) {
	value, _ := json.Marshal(users)
	w.WriteHeader(status)
	fmt.Fprintf(w, `{"status": %d, "data": %s}`, status, value)
}
