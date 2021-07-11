package main

import (
	"github.com/jpastorm/redis-cache/cache/connection"
	"github.com/jpastorm/redis-cache/cache/handler"
	"github.com/jpastorm/redis-cache/cache/storage"
	"github.com/jpastorm/redis-cache/cache/usecase"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	client := connection.NewRedis()
	newStorage := storage.NewStorage(client)
	cache := usecase.NewCache(newStorage)
	handler.NewRouter(mux, cache)

	server := &http.Server{
		Addr: ":8082",
		Handler: mux,
	}
	log.Fatal(server.ListenAndServe())
}
