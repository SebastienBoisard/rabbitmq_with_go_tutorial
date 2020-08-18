package tutorial

import (
	"bytes"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"time"
)

func PlayTutorial04() {

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

	// Define how many messages or how many bytes RabbitMQ will try to keep on.
	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	if err != nil {
		log.Fatalf("Failed to set QoS [%s]", err)
	}

	fmt.Println("Create a consumer")
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer [%s]", err)
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			fmt.Printf("Received a message: %s\n", d.Body)
			dotCount := bytes.Count(d.Body, []byte("."))
			t := time.Duration(dotCount)
			time.Sleep(t * time.Second)
			log.Printf("Done")
			err := d.Ack(false)
			if err != nil {
				log.Fatalf("Failed to send an acknowledgement to RabbitMQ [%s]", err)
			}
		}
	}()

	fmt.Println("Waiting for messages (to exit press CTRL+C)")
	<-forever
}
