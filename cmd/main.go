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
	maxID = 3
}

func main() {

	http.HandleFunc("/users", UserServer)
	fmt.Println("server started at port 8080s")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

type User struct {
	ID        uint64 `json: "id"`
	FirstName string `json: "first_name"`
	LastName  string `json: "last_name"`
	Email     string `json: "email"`
}

var users []User

var maxID uint64

func UserServer(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		GetAllUsers(w)
	case http.MethodPost:
		decode := json.NewDecoder(r.Body)
		var u User
		if err := decode.Decode(&u); err != nil {
			MsgResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		PostUser(w, u)

	default:
		InvalidMethod(w)
	}
}

func GetAllUsers(w http.ResponseWriter) {
	DataResponse(w, http.StatusOK, users)
}

func PostUser(w http.ResponseWriter, data User) {
	user := data
	if user.FirstName == "" {
		MsgResponse(w, http.StatusBadRequest, "First name is required")
		return
	}
	if user.LastName == "" {
		MsgResponse(w, http.StatusBadRequest, "Last name is required")
		return
	}
	if user.Email == "" {
		MsgResponse(w, http.StatusBadRequest, "Email is required")
		return
	}
	maxID++
	user.ID = maxID
	users = append(users, user)
	DataResponse(w, http.StatusCreated, user)

}

func InvalidMethod(w http.ResponseWriter) {
	status := http.StatusNotFound
	w.WriteHeader(status)
	fmt.Fprintf(w, `{"status": %d, "message": "%s"}`, status, "not found")

}

func MsgResponse(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)
	fmt.Fprintf(w, `{"status": %d, "message": "%s"}`, status, message)
}

func DataResponse(w http.ResponseWriter, status int, users interface{}) {
	value, err := json.Marshal(users)
	if err != nil {
		MsgResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	w.WriteHeader(status)
	fmt.Fprintf(w, `{"status": %d, "data": %s}`, status, value)
}
