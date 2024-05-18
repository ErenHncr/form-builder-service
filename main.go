package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/erenhncr/go-api-structure/api"
	"github.com/erenhncr/go-api-structure/storage"
)

func main() {
	listenAddr := flag.String("listenaddr", "localhost:3000", "the server port")
	flag.Parse()

	store := storage.NewMemoryStorage()

	server := api.NewServer(*listenAddr, store)
	fmt.Println("server running on port:", *listenAddr)
	log.Fatal(server.Start())
}
