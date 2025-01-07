package main

import (
	"encoding/json"
	"github/princedraculla/toll-calculation/types"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/sirupsen/logrus"
)

type KafkaConsumer struct {
	Consumer    *kafka.Consumer
	IsRunning   bool
	CalcService CalculateServicer
}

func NewKafkaConsumer(topic string, svc CalculateServicer) (*KafkaConsumer, error) {
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
		Consumer:    c,
		CalcService: svc,
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
		var data types.OBUData
		if err := json.Unmarshal(msg.Value, &data); err != nil {
			logrus.Errorf("can unmarshal to obu data with error: %s\n", err)
			continue
		}
		distance, err := c.CalcService.CalculateDistance(data)
		if err != nil {
			logrus.Errorf("calculation error with error :%s\n", err)
		}
		_ = distance
	}
}
