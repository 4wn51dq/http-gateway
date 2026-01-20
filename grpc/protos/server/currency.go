package server

import (
	"context"
	"grpc/protos/protos/currency"

	"github.com/hashicorp/go-hclog"
)

type Currency struct {
	currency.UnimplementedCurrencyServer
	log hclog.Logger
}

func NewCurrency(l hclog.Logger) *Currency {
	return &Currency{log: l}
}

func (c *Currency) GetRate(ctx context.Context, rr *currency.RateRequest) (*currency.RateResponse, error) {
	c.log.Info("Handle GetRate", "base", rr.GetBase(), "destination", rr.GetDestination())
	return &currency.RateResponse{Rate: 0.5}, nil
}
