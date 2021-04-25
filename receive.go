package main

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/streadway/amqp"
)

func runReceive(cmd *cobra.Command, args []string) {
	var amqpUrl string
	if username != "" && password != "" {
		amqpUrl = fmt.Sprintf("%s://%s:%s@%s:%d/", scheme, username, password, host, port)
	} else {
		amqpUrl = fmt.Sprintf("%s://%s:%d/", scheme, host, port)
	}
	conn, err := amqp.Dial(amqpUrl)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %s", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %s", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %s", err)
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %s", err)
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf("Waiting for messages. To exit press CTRL+C")
	<-forever
}
