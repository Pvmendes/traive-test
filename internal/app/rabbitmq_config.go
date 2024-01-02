package app

import (
	"log"

	"github.com/streadway/amqp"
)

// RabbitMQConfig represents the configuration for RabbitMQ.
type RabbitMQConfig struct {
	URL      string // RabbitMQ server URL (e.g., "amqp://guest:guest@localhost:5672/")
	Exchange string // Exchange name
	Queue    string // Queue name
}

func (config RabbitMQConfig) OnError(err error, msg string) {
	if err != nil {
		log.Printf("Error occurred while publishing message on '%s' queue. Error message: %s", config.Queue, msg)
	}

}

// PublishNotification publishes a notification message to RabbitMQ.
func (config RabbitMQConfig) PublishNotification(message string) {
	if config.URL == "test" {
		return
	}

	// Establish a connection to RabbitMQ
	log.Println("Establish a connection to RabbitMQ")
	log.Println(config.URL)

	conn, err := amqp.Dial(config.URL)
	config.OnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	log.Println("Create a channel")
	// Create a channel
	ch, err := conn.Channel()
	config.OnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		config.Queue, // name
		false,        // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	config.OnError(err, "Failed to declare a queue")

	log.Println("Queue Name: " + q.Name)
	err = ch.Publish(
		"",     // Exchange name
		q.Name, // Routing key
		false,  // Mandatory
		false,  // Immediate
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte(message),
		},
	)
	config.OnError(err, "Failed to publish a message")
}

// ConsumeNotifications starts a consumer to receive notifications from RabbitMQ.
func (config RabbitMQConfig) ConsumeNotifications(queue string) (<-chan amqp.Delivery, error) {
	// Establish a connection to RabbitMQ
	conn, _ := amqp.Dial(config.URL)

	// Create a channel
	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
	}

	defer conn.Close()
	defer ch.Close()

	msgs, err := ch.Consume(
		queue, // Queue name
		"",    // Consumer name (empty for auto-generated name)
		false, // Auto-Ack (false to manually acknowledge)
		false, // Exclusive
		false, // NoLocal
		false, // NoWait
		nil,   // Arguments
	)
	return msgs, err
}

// Example of how to use ConsumeNotifications:
// msgs, err := ConsumeNotifications(ch, "notifications-queue")
// for msg := range msgs {
//     // Handle the received message (e.g., send notifications to users)
//     log.Printf("Received a message: %s", msg.Body)
//     msg.Ack(false) // Acknowledge the message
// }
