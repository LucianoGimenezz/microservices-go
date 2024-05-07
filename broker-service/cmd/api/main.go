package main

import (
	"fmt"
	"log"
	"net/http"
)

const webPort = "80"

type Config struct{}

func main() {

	app := Config{}
	svr := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	err := svr.ListenAndServe()

	if err != nil {
		log.Panic(err)
	}

	log.Printf("Broker server listening on port %s", webPort)
}
