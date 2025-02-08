package main

import (
	"common/api"
	"context"
)

type OrderService interface {
	CreateOrder(context.Context, *api.Order) error
}