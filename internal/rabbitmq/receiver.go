package rabbitmq

import (
	"consumer/internal/handlers"
	"consumer/config"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQReceiver struct {
	ConnectionString string
	ExchangeName     string
	RoutingKey       string
}

func NewRabbitMQReceiver() *RabbitMQReceiver {
	return &RabbitMQReceiver{
		ConnectionString: config.GetEnv("RABBITMQ_URL"),
		ExchangeName:     config.GetEnv("RABBITMQ_EXCHANGE"),
		RoutingKey:       config.GetEnv("ROUTING_KEY"),
	}
}

func (r *RabbitMQReceiver) Listen() {
	conn, err := amqp.Dial(r.ConnectionString)
	if err != nil {
		log.Fatalf("Error conectando a RabbitMQ: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Error abriendo el canal: %v", err)
	}
	defer ch.Close()

	err = ch.ExchangeDeclare(
		r.ExchangeName,
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Error declarando el exchange: %v", err)
	}

	q, err := ch.QueueDeclare(
		"",
		false,
		false,
		true,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Error declarando la cola: %v", err)
	}

	err = ch.QueueBind(
		q.Name,
		r.RoutingKey,
		r.ExchangeName,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Error vinculando la cola: %v", err)
	}

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Error registrando el consumidor: %v", err)
	}

	log.Printf("Esperando mensajes en el routing key %s", r.RoutingKey)

	forever := make(chan struct{})
	go func() {
		for d := range msgs {
			handlers.HandleMessageAndSend(d.Body)
		}
	}()

	<-forever
}