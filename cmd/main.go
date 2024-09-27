package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/JessicaEspejo10/go-fundamentals-web-users/internal/domain"
	"github.com/JessicaEspejo10/go-fundamentals-web-users/internal/user"
)

func main() {
	server := http.NewServeMux()

	db := user.DB{
		Users: []domain.User{{
			ID:        1,
			FirstName: "Ana",
			LastName:  "Montes",
			Email:     "amontes@gmail.com",
		}, {
			ID:        2,
			FirstName: "Juan",
			LastName:  "Suarez",
			Email:     "jsuarez@gmail.com",
		}, {
			ID:        3,
			FirstName: "Luis",
			LastName:  "Naranjos",
			Email:     "lnaranjo@gmail.com",
		}, {
			ID:        4,
			FirstName: "Ruth",
			LastName:  "Castro",
			Email:     "rcastro@gmail.com",
		}}, MaxUserID: 4,
	}

	logger := log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)
	repo := user.NewRepo(db, logger)
	service := user.NewService(logger, repo)

	ctx := context.Background()

	server.HandleFunc("/users", user.MakeEndpoints(ctx, service))

	fmt.Println("server started at port 8080")
	log.Fatal(http.ListenAndServe(":8080", server))
}
