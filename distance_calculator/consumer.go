package main

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/sirupsen/logrus"
)

type KafkaConsumer struct {
	Consumer  *kafka.Consumer
	IsRunning bool
}

func NewKafkaConsumer(topic string) (*KafkaConsumer, error) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost",
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		return nil, err
	}

	err = c.SubscribeTopics([]string{topic}, nil)
	if err != nil {
		return nil, err
	}

	return &KafkaConsumer{
		Consumer: c,
	}, nil
}

func (c *KafkaConsumer) Start() {
	logrus.Info("kafka consumer is started")
	c.IsRunning = true
	c.readMessageLoop()
}

func (c *KafkaConsumer) readMessageLoop() {
	for c.IsRunning {
		msg, err := c.Consumer.ReadMessage(-1)
		if err != nil {
			logrus.Errorf("error while consuming message: %s", err)
			continue
		}
		fmt.Println(msg.Value)
	}
}
