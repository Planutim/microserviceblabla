package main

import (
	"fmt"
	"net"
	"os"

	protos "github.com/Planutim/microserviceblabla/currency/protos/currency"
	"github.com/Planutim/microserviceblabla/currency/server"
	hclog "github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	log := hclog.Default()

	gs := grpc.NewServer()
	cs := server.NewCurrency(log)

	protos.RegisterCurrencyServer(gs, cs)
	reflection.Register(gs)
	l, err := net.Listen("tcp", ":9091")
	fmt.Println("listening")
	if err != nil {
		log.Error("Unable to listen", "error", err)
		os.Exit(1)
	}
	gs.Serve(l)
}
