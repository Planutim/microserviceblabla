package main

import (
	"fmt"
	"net"
	"os"

	"github.com/Planutim/microserviceblabla/currency/data"
	protos "github.com/Planutim/microserviceblabla/currency/protos/currency"
	"github.com/Planutim/microserviceblabla/currency/server"
	hclog "github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	log := hclog.Default()

	gs := grpc.NewServer()
	rates, err := data.NewRates(log)
	if err != nil {
		log.Error("Unable to generate rates", "error", err)
	}
	cs := server.NewCurrency(rates, log)

	protos.RegisterCurrencyServer(gs, cs)
	reflection.Register(gs)
	l, err := net.Listen("tcp", ":9092")
	fmt.Println("listening")
	if err != nil {
		log.Error("Unable to listen", "error", err)
		os.Exit(1)
	}
	gs.Serve(l)
}
