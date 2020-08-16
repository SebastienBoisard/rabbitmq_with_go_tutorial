package tutorial

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

func PlayTutorial01() {

	// Connect to RabbitMQ server
	// The connection abstracts the socket connection, and takes care of protocol version negotiation and authentication
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ [%s]", err)
	}
	defer conn.Close()
	fmt.Println("Connect to RabbitMQ server on port 5672")

	// Create a channel, which is where most of the API for getting things done resides
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel [%s]", err)
	}
	defer ch.Close()
	fmt.Println("Channel created")

	// Declare a queue to be able to publish messages to it
	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue [%s]", err)
	}
	fmt.Println("Queue declared")


	fmt.Println("Publish message")
	body := "Hello World!"
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing {
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	if err != nil {
		log.Fatalf("Failed to publish a message [%s]", err)
	}
	fmt.Printf(" Message '%s' sent to RabbitMQ", body)
}
