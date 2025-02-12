package library

import (
	"encoding/json"
	"feedhive/notifications/repository"
	"feedhive/notifications/variable"
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func Receive() {
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@localhost:5672/", variable.AMQP_USER,
		variable.AMQP_PASS))
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	exchange := ""
	q, err := ch.QueueDeclare(
		variable.ADD_FEED, // name
		true,              // durable
		false,             // delete when unused
		false,             // exclusive
		false,             // no-wait
		nil,               // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name,   // queue
		exchange, // consumer
		true,     // auto-ack
		false,    // exclusive
		false,    // no-local
		false,    // no-wait
		nil,      // args
	)
	failOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			repo := repository.NewNotificationRepository()
			var feedId uint
			json.Unmarshal(d.Body, &feedId)
			repo.CreateFeedNotification(feedId)
		}
	}()

	log.Printf(" [*] Add Feed: Waiting for messages. To exit press CTRL+C")
	<-forever
}

// func ReceiveMarkAsRead() {
// 	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@localhost:5672/", variable.AMQP_USER,
// 		variable.AMQP_PASS))
// 	failOnError(err, "Failed to connect to RabbitMQ")
// 	defer conn.Close()

// 	ch, err := conn.Channel()
// 	failOnError(err, "Failed to open a channel")
// 	defer ch.Close()

// 	q, err := ch.QueueDeclare(
// 		variable.MARK_AS_READ, // name
// 		true,                  // durable
// 		false,                 // delete when unused
// 		false,                 // exclusive
// 		false,                 // no-wait
// 		nil,                   // arguments
// 	)
// 	failOnError(err, "Failed to declare a queue")

// 	msgs, err := ch.Consume(
// 		q.Name, // queue
// 		"",     // consumer
// 		true,   // auto-ack
// 		false,  // exclusive
// 		false,  // no-local
// 		false,  // no-wait
// 		nil,    // args
// 	)
// 	failOnError(err, "Failed to register a consumer")

// 	var forever chan struct{}

// 	go func() {
// 		for d := range msgs {
// 			log.Printf("Received a message: %s", d.Body)
// 			repo := repository.NewNotificationRepository()
// 			var notificationId uint
// 			json.Unmarshal(d.Body, &notificationId)
// 			repo.MarkAsRead(notificationId)
// 		}
// 	}()

// 	log.Printf(" [*] Mark Read: Waiting for messages. To exit press CTRL+C")
// 	<-forever
// }
