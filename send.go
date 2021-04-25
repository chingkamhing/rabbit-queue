package main

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/streadway/amqp"
)

func runSend(cmd *cobra.Command, args []string) {
	body := args[0]
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

	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	if err != nil {
		log.Fatalf("Failed to publish a message: %s", err)
	}
	log.Printf("Sent %s", body)
}
