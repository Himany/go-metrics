package main

import (
	"log"

	"github.com/Himany/go-metrics/handlers"
	"github.com/Himany/go-metrics/server"
	"github.com/Himany/go-metrics/storage"
)

func main() {
	repo := storage.NewMemStorage()
	handler := &handlers.UpdateHandler{Repo: repo}

	if err := server.Run(handler); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
