package helpers

import (
    "github.com/streadway/amqp"
    "log"
)

var rabbitConn *amqp.Connection
var rabbitChannel *amqp.Channel

func InitRabbitMQ(uri string) error {
    conn, err := amqp.Dial(uri)
    if err != nil {
        log.Printf("Failed to connect to RabbitMQ: %v", err)
        return err
    }
    rabbitConn = conn
    log.Println("Connected to RabbitMQ")

    channel, err := conn.Channel()
    if err != nil {
        log.Printf("Failed to create RabbitMQ channel: %v", err)
        return err
    }
    rabbitChannel = channel
    return nil
}

// GetRabbitMQChannel retorna el canal de RabbitMQ para ser usado en otras capas.
func GetRabbitMQChannel() *amqp.Channel {
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