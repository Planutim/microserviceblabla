package main

import (
	"net"
	"os"

	"example.com/protos"

	hclog "github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	log := hclog.Default()

	gs := grpc.NewServer()
	cs := currency.NewCurrency(log)

	protos.RegisterCurrencyServer(gs, cs)
	reflection.Register(gs)
	l, err := net.Listen("tcp", ":9091")
	if err != nil {
		log.Error("Unable to listen", "error", err)
		os.Exit(1)
	}
	gs.Serve(l)
}
