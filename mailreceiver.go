package main

import (
	"encoding/json"
	"log"

	"github.com/streadway/amqp"
)

//MailConfirmationQueue structure
type MailConfirmationQueue struct {
	Type    string
	Message string
}

func receiverMailMessage() {
	conn, err := amqp.Dial(rabbitServer)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()
	log.Printf("iniciando:")
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		storageResponseQueue, // name
		true,                 // durable
		false,                // delete when unused
		false,                // exclusive
		false,                // no-wait
		nil,                  // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	go func() {
		for d := range msgs {
			var confirmation ConfirmationQueue
			errDecoding := json.Unmarshal(d.Body, &confirmation)
			if errDecoding != nil {
				panic(errDecoding)
			}
			log.Println(confirmation.Type + " : " + confirmation.Message)
			d.Ack(false)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	forever := make(chan bool)
	<-forever
}
