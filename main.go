package main

import (
	"context"
	"flag"
	"log"

	"github.com/erenhncr/go-api-structure/api"
	"github.com/erenhncr/go-api-structure/storage"
	"github.com/joho/godotenv"
)

func main() {
	listenAddr := flag.String("listenaddr", "localhost:3000", "the server port")
	envFile := flag.String("env-file", ".env", "env file location")
	flag.Parse()

	err := godotenv.Load(*envFile)
	if err != nil {
		log.Fatal("error loading .env file")
		return
	}
	log.Printf("environment file: %v", *envFile)

	ctx := context.Background()
	store := storage.NewStorage()
	err = store.Connect(ctx)

	if err != nil {
		log.Fatal(err.Error())
		return
	}

	defer store.Disconnect(ctx)

	server := api.NewServer(*listenAddr, store)
	log.Println("server running on port:", *listenAddr)
	log.Fatal(server.Start())
}
