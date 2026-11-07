package main

import (
	"context"
	"log"
	"os"

	"github.com/whitehackerh/veylo/api_server/src/bootstrap"
)

func main() {
	app, err := bootstrap.NewApp(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer app.Pool.Close()

	port := os.Getenv("API_PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("API server listening on :%s", port)
	log.Fatal(app.Server.Start(":" + port))
}
