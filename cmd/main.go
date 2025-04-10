package main

import (
	"consumer/config"
	"consumer/internal/rabbitmq"
)

func main() {
	// Cargar las variables de entorno
	config.LoadEnv()

	// Inicializar el receptor de RabbitMQ
	receiver := rabbitmq.NewRabbitMQReceiver()

	// Escuchar mensajes
	receiver.Listen()
}