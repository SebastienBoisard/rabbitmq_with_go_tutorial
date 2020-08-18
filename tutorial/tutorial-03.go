package tutorial

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"math/rand"
	"strings"
	"time"
)

func PlayTutorial03() {

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

	// Declare a queue to be able to consume messages from it (same producer queue as in tutorial-01)
	q, err := ch.QueueDeclare(
		"task_queue", // name
		true,         // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue [%s]", err)
	}
	fmt.Println("Queue declared")

	fmt.Println(string(rand.Intn(10)))

	rand.Seed(time.Now().UnixNano())
	body := fmt.Sprintf("test_%d%s", rand.Intn(10), strings.Repeat(".", rand.Intn(5)))

	fmt.Println("publish a message")
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte(body),
		})
	if err != nil {
		log.Fatalf("Failed to register a consumer [%s]", err)
	}

	fmt.Printf("Message '%s' sent to RabbitMQ\n", body)
}
