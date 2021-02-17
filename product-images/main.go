package main

import (
	"context"
	"example/files"
	"example/handlers"
	"net/http"
	"os"
	"os/signal"
	"time"

	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	hclog "github.com/hashicorp/go-hclog"
)

func main() {

	l := hclog.New(
		&hclog.LoggerOptions{
			Name:  "product-images",
			Level: hclog.LevelFromString("debug"),
		},
	)
	//create a logger for the server from the default logger
	sl := l.StandardLogger(&hclog.StandardLoggerOptions{InferLevels: true})

	// create the storage class, use local storage
	// max filesize 5MB
	stor, err := files.NewLocal("./imagestore", 1024*1000*5)
	if err != nil {
		l.Error("Unable to create storage", "error", err)
		os.Exit(1)
	}

	fh := handlers.NewFiles(stor, l)
	mw := handlers.GzipHandler{}
	// create a new serve mux
	sm := mux.NewRouter()

	ch := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"*"}))
	// filename regex : {filename: [a-zA-Z]+\\.[a-z]{3}}
	// problem with FileServer is that it is dumb
	ph := sm.Methods(http.MethodPost).Subrouter()
	ph.HandleFunc("/images/{id:[0-9]+}/{filename:[a-zA-Z]+\\.[a-z]{3}}", fh.UploadREST)
	ph.HandleFunc("/", fh.UploadMultipart)
	gh := sm.Methods(http.MethodGet).Subrouter()
	gh.Handle(
		"/images/{id:[0-9]+}/{filename:[a-zA-Z]+\\.[a-z]{3}}",
		http.StripPrefix("/images/", http.FileServer(http.Dir("./imagestore"))))
	gh.Use(mw.GzipMiddleware)
	s := http.Server{
		Addr:         ":9091",
		Handler:      ch(sm),
		ErrorLog:     sl,
		ReadTimeout:  5 * time.Second,   //max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}

	//start the server
	go func() {
		l.Info("Starting server", "bind_address", ":9091")
		err := s.ListenAndServe()
		if err != nil {
			l.Error("Unable to start server", "error", err)
			os.Exit(1)
		}
	}()

	// trap sigterm or interrupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	//block until signal is received
	sig := <-c
	l.Info("Shutting down server with", "signal", sig)

	// gracefully shutdown the server, wating max 30 seconds for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)
}
