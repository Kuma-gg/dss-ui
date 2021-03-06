package main

import (
	"github.com/streadway/amqp"
)

//Mail structure
type Mail struct {
	Name string
	Mail string
}

func sendEmailChannel(dataFile []byte) {
	conn, err := amqp.Dial(rabbitServer)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		emailRequestQueue, // name
		true,             // durable
		false,            // delete when unused
		false,            // exclusive
		false,            // no-wait
		nil,              // arguments

	)
	failOnError(err, "Failed to declare a queue")
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        dataFile,
		})
	failOnError(err, "Failed to publish a message")
}
