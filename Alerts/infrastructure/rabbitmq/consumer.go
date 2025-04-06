package rabbitmq

import (
	"log"
	"time"
	"fmt"
	_"strconv"
	_"regexp"
	"encoding/json"
	"github.com/JosephAntony37900/Multi-API-Consumer-1/Alerts/application"
	"github.com/JosephAntony37900/Multi-API-Consumer-1/helpers"
	"github.com/streadway/amqp"
	"github.com/JosephAntony37900/Multi-API-Consumer-1/Alerts/infrastructure/websocket"
)

func ConfigureAndConsume(queueName, routingKey, exchangeName string, handleMessage func(msg amqp.Delivery)) error {
	channel := helpers.GetRabbitMQChannel()
	if channel == nil {
		return logError("RabbitMQ channel is not initialized")
	}

	// Declarar el exchange
	if err := channel.ExchangeDeclare(
		exchangeName, // nombre del exchange
		"topic",      // tipo
		true,         // durable
		false,        // auto-delete
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	); err != nil {
		return logError("Failed to declare exchange: %v", err)
	}

	// Declarar la cola
	queue, err := channel.QueueDeclare(
		queueName, // nombre de la cola
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		return logError("Failed to declare queue: %v", err)
	}

	// Vincular la cola al exchange con la routing key
	if err := channel.QueueBind(
		queue.Name,   // nombre de la cola
		routingKey,   // routing key
		exchangeName, // exchange
		false,        // no-wait
		nil,          // arguments
	); err != nil {
		return logError("Failed to bind queue: %v", err)
	}

	messages, err := channel.Consume(
		queue.Name, // nombre de la cola
		"",         // consumer
		true,       // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)
	if err != nil {
		return logError("Failed to register a consumer: %v", err)
	}

	// Procesar mensajes
	go func() {
		for msg := range messages {
			log.Printf("Received a message: %s", msg.Body)
			handleMessage(msg)
		}
	}()

	log.Println("Waiting for messages...")
	return nil
}

func StartAlertConsumer(service *application.CreateAlert, queueName, routingKey, exchangeName string) error {
    handleMessage := func(msg amqp.Delivery) {
        log.Printf("Processing message: %s", msg.Body)

        // Deserializar el mensaje recibido
        var payload struct {
            Estado              string `json:"estado"`
            IdLectura           int    `json:"id_lectura"`
            CodigoIdentificador string `json:"codigo_identificador"`
            Tipo                bool   `json:"tipo"` 
        }

        err := json.Unmarshal(msg.Body, &payload)
        if err != nil {
            log.Printf("Error deserializando el mensaje: %v", err)
            return
        }

        log.Printf("Estado: %s, IdLectura: %d, CodigoIdentificador: %s, Tipo: %t",
            payload.Estado, payload.IdLectura, payload.CodigoIdentificador, payload.Tipo)

        err = service.Run(payload.IdLectura, payload.Estado, time.Now(), payload.CodigoIdentificador, payload.Tipo)
        if err != nil {
            log.Printf("Error al procesar la alerta: %v", err)
        }
		websocket.BroadcastMessage(msg.Body)
    }

    return ConfigureAndConsume(queueName, routingKey, exchangeName, handleMessage)
}

func logError(format string, args ...interface{}) error {
	log.Printf(format, args...)
	return fmt.Errorf(format, args...)
}