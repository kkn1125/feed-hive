package library

import (
	"context"
	"encoding/json"
	"feedhive/feeds/model"
	"feedhive/feeds/variable"
	"fmt"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func Send(feed *model.Feed) {
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
		false,             // auto delete (delete when unused)
		false,             // exclusive
		false,             // no-wait
		nil,               // arguments
	)
	failOnError(err, "Failed to declare a queue")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	log.Printf(" [x] Sending feed: %v\n", feed)

	jsonValue, err := json.Marshal(feed.ID)
	if err != nil {
		log.Panicf("Failed to marshal user_id: %s", err)
	}
	err = ch.PublishWithContext(ctx,
		exchange, // exchange
		q.Name,   // routing key
		false,    // mandatory
		false,    // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(jsonValue),
		})
	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s\n", jsonValue)
}

// func SendMarkAsRead(notificationId uint) {
// 	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@localhost:5672/", variable.AMQP_USER,
// 		variable.AMQP_PASS))
// 	failOnError(err, "Failed to connect to RabbitMQ")
// 	defer conn.Close()

// 	ch, err := conn.Channel()
// 	failOnError(err, "Failed to open a channel")
// 	defer ch.Close()

// 	exchange := ""
// 	q, err := ch.QueueDeclare(
// 		variable.MARK_AS_READ, // name
// 		true,                  // durable
// 		false,                 // delete when unused
// 		false,                 // exclusive
// 		false,                 // no-wait
// 		nil,                   // arguments
// 	)
// 	failOnError(err, "Failed to declare a queue")

// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()

// 	log.Printf(" [x] Sending notificationId: %v\n", notificationId)

// 	jsonValue, err := json.Marshal(notificationId)
// 	if err != nil {
// 		log.Panicf("Failed to marshal notification_id: %s", err)
// 	}
// 	err = ch.PublishWithContext(ctx,
// 		exchange, // exchange
// 		q.Name,   // routing key
// 		false,    // mandatory
// 		false,    // immediate
// 		amqp.Publishing{
// 			ContentType: "text/plain",
// 			Body:        []byte(jsonValue),
// 		})
// 	failOnError(err, "Failed to publish a message")
// 	log.Printf(" [x] Sent %s\n", jsonValue)
// }
