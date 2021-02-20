package main

import (
	// "example/handlers"

	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Planutim/microserviceblabla/product-api/data"
	"github.com/Planutim/microserviceblabla/product-api/handlers"

	"context"

	protos "github.com/Planutim/microserviceblabla/product-api/currency/"
	"github.com/go-openapi/runtime/middleware"
	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

func main() {

	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	v := data.NewValidation()

	conn, err := grpc.Dial("localhost:9092")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	cc := protos.NewCurrencyClient(conn)
	ph := handlers.NewProducts(l, v, cc)
	// sm := http.NewServeMux()
	sm := mux.NewRouter()
	getRouter := sm.Methods(http.MethodGet).Subrouter()
	// sm.Handle("/", ph)
	// sm.Handle("/products", ph).Method("GET")
	getRouter.HandleFunc("/products", ph.ListAll)
	getRouter.HandleFunc("/products/{id:[0-9]+}", ph.ListSingle)

	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/products/{id:[0-9]+}", ph.UpdateProducts)
	putRouter.Use(ph.MiddleWareProductValidation)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/products", ph.AddProduct)
	postRouter.Use(ph.MiddleWareProductValidation)

	deleteRouter := sm.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/products/{id:[0-9]+}", ph.DeleteProduct)

	opts := middleware.RedocOpts{
		SpecURL: "/swagger.yaml",
	}
	sh := middleware.Redoc(opts, nil)
	getRouter.Handle("/docs", sh)
	getRouter.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	ch := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"*"}))
	s := &http.Server{
		Addr:         ":9890",
		Handler:      ch(sm),
		ErrorLog:     l,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		l.Println("Starting server on port 9890")
		err := s.ListenAndServe()
		if err != nil {
			l.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Println("Received terminate, graceful shutdown", sig)
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}
