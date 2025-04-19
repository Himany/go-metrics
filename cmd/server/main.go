package main

import (
	"log"

	"github.com/Himany/go-metrics/cmd/server/handlers"
	"github.com/Himany/go-metrics/cmd/server/server"
	"github.com/Himany/go-metrics/cmd/server/storage"
)

func main() {
	repo := storage.NewMemStorage()
	handler := &handlers.UpdateHandler{Repo: repo}

	if err := server.Run(handler); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
