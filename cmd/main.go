package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/JessicaEspejo10/go-fundamentals-web-users/internal/user"
	"github.com/JessicaEspejo10/go-fundamentals-web-users/pkg/bootstrap"
)

func main() {
	server := http.NewServeMux()

	db := bootstrap.NewDB()

	logger := bootstrap.NewLogger()
	repo := user.NewRepo(db, logger)
	service := user.NewService(logger, repo)

	ctx := context.Background()

	server.HandleFunc("/users", user.MakeEndpoints(ctx, service))

	fmt.Println("server started at port 8080")
	log.Fatal(http.ListenAndServe(":8080", server))
}
