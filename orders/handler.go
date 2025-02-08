package main

import (
	"common/api"
	"context"

	"google.golang.org/grpc"
)

type grpcHandler struct {
	ordersService OrderService
	api.UnimplementedOrderServiceServer
}

func NewGrpcHandler (server *grpc.Server, ordersService OrderService) {
	handler := &grpcHandler{
		ordersService: ordersService,
	}

	// register the OrderServiceServer to GRPC
	api.RegisterOrderServiceServer(server, handler)
}

func (h *grpcHandler) CreateOrder(ctx context.Context, req *api.CreateOrderRequest) (*api.Order, error) {
	order := &api.Order{
		ID: "1234",
		CustomerID: req.CustomerID,
		Items: req.Items,
		Status: "snsmxmx",
	}

	err := h.ordersService.CreateOrder(ctx, order)

	if err != nil {
		return nil, err
	}

	return order, nil
}