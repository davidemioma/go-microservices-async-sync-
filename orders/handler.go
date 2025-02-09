package main

import (
	"common"
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

	var formmatedItems []common.ItemType

	for _, item := range req.Items {
		formmatedItems = append(formmatedItems, common.ItemType{
			ID: item.ID,
			Name: item.Name,
			Quantity: item.Quantity,
			PriceID: item.PriceID,
		})
	}

	// Publish order message for rabbitMq
	publishOrderMessage(ctx, common.OrderType{
		ID: "1234",
		CustomerID: req.CustomerID,
		Items: formmatedItems,
		Status: "snsmxmx",
	})

	return order, nil
}