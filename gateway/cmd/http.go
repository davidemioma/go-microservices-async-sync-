package main

import (
	"common/api"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Http struct {
	client api.OrderServiceClient
}

func NewHttpHandler (client api.OrderServiceClient) *Http {
	return &Http{
		client,
	}
}

// Handle Routes
func (h *Http) mount() *fiber.App {
	r := fiber.New()

	// Cors
	r.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowMethods:     "GET, POST, PUT, DELETE, OPTIONS",
		AllowHeaders:     "Accept, Authorization, Content-Type, Set-Cookie",
		ExposeHeaders:    "Content-Length, Link, Set-Cookie",
		AllowCredentials: true,  // Change this to true to allow cookies
		MaxAge:           int(time.Hour * 12 / time.Second),
	}))

	// Routes
	api := r.Group("/api")

	api.Post("/customers/:customerId/orders", h.createOrderHandler)

	return r
}

// Run Server
func (h *Http) run(port string, handler *fiber.App) error {
	err := handler.Listen(port)

	if err != nil {
		return err
	}

	return nil
}