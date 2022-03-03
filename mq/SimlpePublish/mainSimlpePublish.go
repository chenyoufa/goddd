package main

import (
	"fmt"
	"gdemo1/mq/RabbitMQ"
	"strconv"
	"time"
)

func main() {
	rabbitmq := RabbitMQ.NewRabbitMQSimple("kuteng")
	for i := 0; i <= 100; i++ {
		rabbitmq.PublishSimple("Hello kuteng!" + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}
}
