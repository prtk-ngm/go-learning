package main

import (
	"fmt"

	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

func main() {

}

// sendIngressCarEvent produces an event for a car approaching an intersection.
func sendIngressCarEvent(producer *kafka.Producer, toppar kafka.TopicPartition) {

	producer, err := kafka.NewProducer(nil)

	kafka.TopicPartition()
	err := producer.Produce(&kafka.Message{
		TopicPartition: toppar,
		Key:            []byte("hello"),
		Value:          []byte("world")},
		nil)

	if err != nil{
		fmt.Println("Error:",err)
	}

}


