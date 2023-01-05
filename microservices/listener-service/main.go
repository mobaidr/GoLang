package main

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"listener/event"
	"log"
	"math"
	"os"
	"time"
)

func main() {
	// try to connect RabbitMQ
	rabbitCon, err := connect()

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	defer rabbitCon.Close()


	// start Listening for messages.
	log.Println("Listening for and consuming RabbitMQ Messages....")

	// Create a consumer
	consumer, err := event.NewConsumer(rabbitCon)

	if err != nil {
		panic(err)
	}

	// watch the queue  and consume events
	err = consumer.Listen([]string{"log.INFO", "log.WARNING", "log.ERROR"})

	if err != nil {
		log.Println(err)
	}
}

func connect() (*amqp.Connection, error) {
	var counts int64
	var backoff = 1 * time.Second
	var connection *amqp.Connection

	// Do not continue until rabbit is ready
	for {
		c, err := amqp.Dial("amqp://guest:guest@localhost")

		if err != nil {
			fmt.Println("Rabbit MQ is not ready")
			counts++
		} else {
			log.Println("Connected to RabbitMQ")
			connection = c
			break
		}

		if counts > 5 {
			fmt.Println(err)
			return nil, err
		}

		backoff = time.Duration(math.Pow(float64(counts), 2)) * time.Second
		log.Println("backing off..")
		time.Sleep(backoff)

		continue
	}

	return connection, nil
}
