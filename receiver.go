package main

import (
	"encoding/json"
	"github.com/streadway/amqp"
	"log"
)

type ConfirmationQueue struct {
	Type    string
	Message string
	Event string
}

type EmailMessage  struct {
	Users []User
	Event string
}

func receiverFileMessageStorage() {
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
			if confirmation.Type != "error" {
				users := getUsers()

				var emailsMessages   EmailMessage
				emailsMessages.Event = confirmation.Event
				emailsMessages.Users = users
				documentJSON, err := json.Marshal(emailsMessages)
				if err != nil {
					panic(err)
				}

				sendEmailMessage(documentJSON);
			}
			log.Println(confirmation.Type + " : " + confirmation.Message)
			d.Ack(false)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	forever := make(chan bool)
	<-forever
}


func receiverEmailMessage() {
	conn, err := amqp.Dial(rabbitServer)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()
	log.Printf("iniciando:")
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		emailResponseQueue, // name
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