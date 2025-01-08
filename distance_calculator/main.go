package main

import (
	"github/princedraculla/toll-calculation/aggregator/client"
	"log"
)

const (
	kafkaTopic  = "obudata"
	aggEndpoint = "http://127.0.0.1:5000/aggregate"
)

func main() {
	var (
		err error
		svc CalculateServicer
	)
	svc = NewCalculatorService()
	svc = NewLogMiddleWareConsumer(svc)
	kafkaConsumer, err := NewKafkaConsumer(kafkaTopic, svc, client.NewClient(aggEndpoint))
	if err != nil {
		log.Fatal(err)
	}
	kafkaConsumer.Start()
}
