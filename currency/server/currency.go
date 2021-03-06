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
	log           hclog.Logger
	rates         *data.ExchangeRates
	subscriptions map[protos.Currency_SubscribeRatesServer][]*protos.RateRequest
}

func (c *Currency) GetRate(ctx context.Context, rr *protos.RateRequest) (*protos.RateResponse, error) {
	c.log.Info("Handle GetRate", "base", rr.GetBase(), "destination", rr.GetDestination())

	// return &protos.RateResponse{Rate: 0.5}, nil
	rate, err := c.rates.GetRate(rr.GetBase().String(), rr.GetDestination().String())
	if err != nil {
		return nil, err
	}
	return &protos.RateResponse{Base: rr.Base, Destination: rr.Destination, Rate: rate}, nil

}

func (c *Currency) SubscribeRates(src protos.Currency_SubscribeRatesServer) error {
	for {
		rr, err := src.Recv()
		if err == io.EOF {
			c.log.Info("Client has closed connection")
			return err
		}
		if err != nil {
			c.log.Error("Unable to read from client", "error", err)
			return err
		}
		c.log.Info("Handle client request", "request", rr)
		rrs, ok := c.subscriptions[src]
		if !ok {
			rrs = []*protos.RateRequest{}
		}
		rrs = append(rrs, rr)
		c.subscriptions[src] = rrs
	}
	return nil
	// for {
	// 	err := src.Send(&protos.RateResponse{Rate: 12.1})
	// 	if err != nil {
	// 		return err
	// 	}
	// 	time.Sleep(5 * time.Second)
	// }
}

func NewCurrency(r *data.ExchangeRates, l hclog.Logger) *Currency {

	c := &Currency{log: l, rates: r, subscriptions: make(map[protos.Currency_SubscribeRatesServer][]*protos.RateRequest)}

	go c.handleUpdates()
	return c
}

func (c *Currency) handleUpdates() {
	ru := c.rates.MonitorRates(5 * time.Second)
	for range ru {
		c.log.Info("Got updates rates")
		// loop over subscribed clients
		for k, v := range c.subscriptions {
			for _, rr := range v {
				r, err := c.rates.GetRate(rr.GetBase().String(), rr.GetDestination().String())
				if err != nil {
					c.log.Error("Unable to get update rate", "base", rr.GetBase().String(), "destination", rr.GetDestination().String())
					return
				}
				err = k.Send(&protos.RateResponse{Rate: r})
				if err != nil {
					c.log.Error("Unable to send updated rate", "base", rr.GetBase().String(), "destination", rr.GetDestination().String())
				}
			}
		}
	}
}
