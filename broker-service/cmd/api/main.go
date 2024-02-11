package main

import (
	"fmt"
	"log"
	"net/http"
)

type Config struct {
}

const WebPort = "8080"

func main() {
	app := Config{}
	log.Printf("Starting broker service on port %s\n", WebPort)

	// define http server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", WebPort),
		Handler: app.routes(),
	}

	// start the server
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
