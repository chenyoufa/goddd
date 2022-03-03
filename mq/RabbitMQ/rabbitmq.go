package RabbitMQ

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

const MQURL = "amqp://admin:123456@127.0.0.1:5672/kuteng"

type RabbitMQ struct {
	conn      *amqp.Connection
	channel   *amqp.Channel
	QueueName string
	Exchange  string
	Key       string
	Mqurll    string
}

func NewRabbitMQ(queueName string, exchange string, key string) *RabbitMQ {
	return &RabbitMQ{
		QueueName: queueName,
		Exchange:  exchange,
		Key:       key,
		Mqurll:    MQURL,
	}
}

func (r *RabbitMQ) Destory() {
	r.channel.Close()
	r.conn.Close()
}

func (r *RabbitMQ) failOnErr(err error, message string) {
	if err != nil {
		log.Fatalf("%s:%s", message, err)
		panic(fmt.Sprintf("%s:%s", message, err))
	}
}

func NewRabbitMQSimple(queueName string) *RabbitMQ {
	rabbitmq := NewRabbitMQ(queueName, "", "")
	var err error
	rabbitmq.conn, err = amqp.Dial(rabbitmq.Mqurll)
	rabbitmq.failOnErr(err, "failed to connect rabbitmq!")

	rabbitmq.channel, err = rabbitmq.conn.Channel()
	rabbitmq.failOnErr(err, "failed to open a channel !")
	return rabbitmq
}

func (r *RabbitMQ) PublishSimple(message string) {

	_, err := r.channel.QueueDeclare(
		r.QueueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println(err)
	}
	r.channel.Publish(
		"",
		r.QueueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
}

func (r *RabbitMQ) ConsumeSimple() {
	q, err := r.channel.QueueDeclare(
		r.QueueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println(err)
	}
	msgs, err := r.channel.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println(err)
	}
	forever := make(chan bool)
	go func() {
		for d := range msgs {
			log.Printf("Received a message:%s ", d.Body)
		}
	}()
	log.Printf("[*] Waiting for message . to exit press ctrl +c")
	<-forever
}
