package server

import (
	"context"
	"io"
	"time"

	"github.com/Planutim/microserviceblabla/currency/data"
	protos "github.com/Planutim/microserviceblabla/currency/protos/currency"

	hclog "github.com/hashicorp/go-hclog"
)

type Currency struct {
	log   hclog.Logger
	rates *data.ExchangeRates
}

func (c *Currency) GetRate(ctx context.Context, rr *protos.RateRequest) (*protos.RateResponse, error) {
	c.log.Info("Handle GetRate", "base", rr.GetBase(), "destination", rr.Destination)

	// return &protos.RateResponse{Rate: 0.5}, nil
	rate, err := c.rates.GetRate(rr.GetBase().String(), rr.GetDestination().String())
	if err != nil {
		return nil, err
	}
	return &protos.RateResponse{Rate: rate}, nil

}

func (c *Currency) SubscribeRates(src protos.Currency_SubscribeRatesServer) error {
	go func() {
		for {
			rr, err := src.Recv()
			if err == io.EOF {
				c.log.Info("Client has closed connection")
				break
			}
			if err != nil {
				c.log.Error("Unable to read from client", "error", err)
				break
			}
			c.log.Info("Handle client request", "request", rr)
		}
	}()
	for {
		err := src.Send(&protos.RateResponse{Rate: 12.1})
		if err != nil {
			return err
		}
		time.Sleep(5 * time.Second)
	}
}

func NewCurrency(r *data.ExchangeRates, l hclog.Logger) *Currency {
	return &Currency{log: l, rates: r}
}
