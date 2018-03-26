package main

import (
	"fmt"
	"log"
	"net/http"

	"hubimage/src/config"
	"hubimage/src/router"
)

func main() {
	conf := config.GetConfig()
	err := config.CheckConfig()
	if err != nil {
		log.Fatalf("Err args: %v", err)
	}
	addr := fmt.Sprintf("%s:%s", conf.Address, conf.Port)
	log.Printf("hubimage server %s starting...\n", addr)

	server := &http.Server{
		Addr:    addr,
		Handler: router.CreateHTTPRouter(),
	}

	err = server.ListenAndServe()
	log.Fatalf("create listen server error: %v", err)
}
