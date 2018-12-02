package main

import (
	"io/ioutil"
	"log"
	"github.com/streadway/amqp"
)
func receiverFileMessage(){
	conn, err := amqp.Dial(rabbitServer)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()
	log.Printf("iniciando:")
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		storageRequestQueue, // name
		true,    // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
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
			//log.Printf("Received a message %s",d.Body)
			err := ioutil.WriteFile(d.MessageId, d.Body, 777)

			if err != nil {
				// handle error
			}
			d.Ack(false)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	forever := make(chan bool)
	<-forever
}

