package main

import (
	"common"
	"context"
	"encoding/json"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

var (
	amqpUser = "guest"
	amqpPass = "guest"
	amqpHost = "localhost"
	amqpPort = "5672"
)

func publishOrderMessage (ctx context.Context, order common.OrderType) {
	ch, close := common.ConnectAmqp(amqpUser, amqpPass, amqpHost, amqpPort)

	defer func ()  {
		close()

		ch.Close()
	}()

	q, err := ch.QueueDeclare(common.OrderCreatedEvent, true, false, false, false, nil)

	if err != nil {
		log.Fatal(err)
	}

	// Convert order to JSON
	orderBytes, err := json.Marshal(order)

	if err != nil {
		log.Fatal(err)
	}

	err = ch.PublishWithContext(ctx, "", q.Name, false, false, amqp.Publishing{
		ContentType: "application/json",
		Body: orderBytes,
	})

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Order published")
}