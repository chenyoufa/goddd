package main

import "gdemo1/mq/RabbitMQ"

func main() {
	rabbitmq := RabbitMQ.NewRabbitMQSimple("kuteng")
	rabbitmq.ConsumeSimple()
}
