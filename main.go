package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type application struct {
	Infolog *log.Logger
}

func main() {
	logger := log.New(os.Stdout, "API-Gateway ", log.Ldate|log.Lshortfile)
	app := application{
		Infolog: logger,
	}
	port := flag.String("port", ":8010", "Http Connection Port Addres")
	serve := &http.Server{
		Addr:    *port,
		Handler: app.routes(),
	}

	app.Infolog.Print("Web is Alive!!")
	err := serve.ListenAndServe()
	if err != nil {
		panic(err)
	}

}
