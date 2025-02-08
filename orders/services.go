package main

import (
	"common/api"
	"context"
)

var ordersDb []*api.Order

type Services struct {
}

func NewService() *Services{
	return &Services{}
}

func (s *Services) CreateOrder (ctx context.Context, order *api.Order) error {
	ordersDb = append(ordersDb, order)

	return nil
}