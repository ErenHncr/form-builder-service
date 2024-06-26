package main

import (
	"context"
	"flag"
	"log"

	"github.com/erenhncr/go-api-structure/api"
	"github.com/erenhncr/go-api-structure/storage"
	"github.com/joho/godotenv"
)

// @title Form Builder API
// @description This is a simple form builder API
// @BasePath /
func main() {
	listenAddr := flag.String("listenaddr", "localhost:3000", "the server port")
	envFile := flag.String("env-file", ".env", "env file location")
	flag.Parse()

	if err := godotenv.Load(*envFile); err != nil {
		log.Fatal("error loading .env file")
		return
	}
	log.Printf("environment file: %v", *envFile)

	ctx := context.Background()
	store := storage.NewStorage()

	if err := store.Connect(ctx); err != nil {
		log.Fatal(err.Error())
		return
	}
	defer store.Disconnect(ctx)

	server := api.NewServer(*listenAddr, store)
	log.Println("server running on port:", *listenAddr)
	log.Fatal(server.Start())
}
