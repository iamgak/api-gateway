package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/redis/go-redis/v9"
)

type application struct {
	Infolog *log.Logger
	Redis   *redis.Client
}

func main() {

	logger := log.New(os.Stdout, "API-Gateway ", log.Ldate|log.Lshortfile)
	app := application{
		Infolog: logger,
		Redis:   InitRedis("localhost", "6379", ""),
	}

	port := flag.String("port", ":8010", "Http Connection Port Addres")

	serve := &http.Server{
		Addr:    *port,
		Handler: app.routes(),
	}

	app.Infolog.Print("API-Gateway is Alive & Kicking !!")
	err := serve.ListenAndServe()
	if err != nil {
		panic(err)
	}

}
