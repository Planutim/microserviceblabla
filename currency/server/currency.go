package server

import (
	"context"

	protos "github.com/Planutim/microserviceblabla/currency/protos/currency"

	hclog "github.com/hashicorp/go-hclog"
)

type Currency struct {
	log hclog.Logger
}

func (c *Currency) GetRate(ctx context.Context, rr *protos.RateRequest) (*protos.RateResponse, error) {
	c.log.Info("Handle GetRate", "base", rr.GetBase(), "destination", rr.Destination)

	return &protos.RateResponse{Rate: 0.5}, nil
}

func NewCurrency(l hclog.Logger) *Currency {
	return &Currency{l}
}
