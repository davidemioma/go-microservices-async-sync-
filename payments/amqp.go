package main

import (
	"common"
	"encoding/json"
	"log"
)

var (
	amqpUser = "guest"
	amqpPass = "guest"
	amqpHost = "localhost"
	amqpPort = "5672"
)

func CreatepaymentLink () (string, error) {
	return "dummy-payment-link.com", nil
}

func listenForOrderMessage () {
	ch, close := common.ConnectAmqp(amqpUser, amqpPass, amqpHost, amqpPort)

	defer func ()  {
		close()

		ch.Close()
	}()

	q, err := ch.QueueDeclare(common.OrderCreatedEvent, true, false, false, false, nil)

	if err != nil {
		log.Fatal(err)
	}

	// Get the message
	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)

	if err != nil {
		log.Fatal(err)
	}

	forever := make(chan struct{})

	go func() {
		for m := range msgs {
			log.Printf("Recieved a message: %s", m.Body)

			// Convert order from JSON to regular OrderType
			o := common.OrderType{}

			err := json.Unmarshal(m.Body, &o)

			if err != nil {
				// Nack negatively acknowledge the delivery of message(s)
				m.Nack(false, false)

				log.Printf("Unable to unmarshal order: %v", err)

				continue
			}

			// Create payment link
			paymentLink, err := CreatepaymentLink()

			if err != nil {
				log.Printf("Unable to create payment link: %v", err)

				// Handle retry here...

				continue
			}

			log.Printf("Payment link generated: %v", paymentLink)
		}
	}()

	log.Printf("AMQP listening. To exit press CMD+C (Mac) or CTRL+C (windows)")

	<-forever
}