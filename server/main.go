package main

import (
	"aggregator/config"
	"log"
	"net/http"
)

func main() {
	cfg := config.Load()
	handler := config.NewRouter(cfg)

	addr := ":" + cfg.ServerPort
	log.Printf("Serving on %s (upstream: %s/flights)", addr, cfg.J1BaseURL)
	log.Fatal(http.ListenAndServe(addr, handler))
}
