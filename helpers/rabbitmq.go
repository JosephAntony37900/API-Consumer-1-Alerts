package helpers

import (
	"log"

	"github.com/streadway/amqp"
)

var rabbitConn *amqp.Connection
var rabbitChannel *amqp.Channel

func InitRabbitMQ(uri string) error {
	var err error
	rabbitConn, err = amqp.Dial(uri)
	if err != nil {
		log.Printf("Failed to connect to RabbitMQ: %v", err)
		return err
	}
	log.Println("Connected to RabbitMQ")

	rabbitChannel, err = rabbitConn.Channel()
	if err != nil {
		log.Printf("Failed to create RabbitMQ channel: %v", err)
		return err
	}
	log.Println("RabbitMQ channel created")
	return nil
}

func GetRabbitMQChannel() *amqp.Channel {
	if rabbitChannel == nil {
		log.Println("RabbitMQ channel is not initialized")
	}
	return rabbitChannel
}

func CloseRabbitMQ() {
	if rabbitChannel != nil {
		rabbitChannel.Close()
	}
	if rabbitConn != nil {
		rabbitConn.Close()
	}
	log.Println("RabbitMQ connection closed")
}