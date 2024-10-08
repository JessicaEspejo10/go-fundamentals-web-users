package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/JessicaEspejo10/go-fundamentals-web-users/internal/user"
	"github.com/JessicaEspejo10/go-fundamentals-web-users/pkg/bootstrap"
	"github.com/JessicaEspejo10/go-fundamentals-web-users/pkg/handler"
)

func main() {
	server := http.NewServeMux()

	db, err := bootstrap.NewDB()

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	logger := bootstrap.NewLogger()
	repo := user.NewRepo(db, logger)
	service := user.NewService(logger, repo)

	ctx := context.Background()

	handler.NewUserHTTPServer(ctx, server, user.MakeEndpoints(ctx, service))

	fmt.Println("server started at port 8080")
	log.Fatal(http.ListenAndServe(":8080", server))
}
