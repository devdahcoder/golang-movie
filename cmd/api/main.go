package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
}

type application struct {
	config config
	log    *log.Logger
}

func main() {

	var cfg config

	flag.IntVar(&cfg.port, "port", 8080, "API server port")

	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")

	flag.Parse()

	logger := log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile|log.Ldate|log.Ltime)

	app := &application{config: cfg, log: logger}

	mux := http.NewServeMux()

	mux.HandleFunc("/health-check", app.healthcheck)

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      mux,
		ErrorLog:     logger,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Printf("starting %s server on %s", cfg.env, srv.Addr)

	err := srv.ListenAndServe()

	if err != nil {
		logger.Fatal(err)
	}

}
