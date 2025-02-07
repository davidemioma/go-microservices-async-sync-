package main

import (
	"common/api"
	"context"

	"google.golang.org/grpc"
)

type grpcHandler struct {
	api.UnimplementedOrderServiceServer
}

func NewGrpcHandler (server *grpc.Server) {
	handler := &grpcHandler{}

	// register the OrderServiceServer to GRPC
	api.RegisterOrderServiceServer(server, handler)
}

func (h *grpcHandler) createOrder(ctx context.Context, req *api.CreateOrderRequest) (*api.Order, error) {
	return &api.Order{
		ID: "",
		CustomerID: req.CustomerID,
		Status: "Pending",
	}, nil
}