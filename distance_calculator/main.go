package main

import "log"

const kafkaTopic = "obudata"

func main() {
	var (
		err error
		svc CalculateServicer
	)
	svc = NewCalculatorService()
	svc = NewLogMiddleWareConsumer(svc)
	kafkaConsumer, err := NewKafkaConsumer(kafkaTopic, svc)
	if err != nil {
		log.Fatal(err)
	}
	kafkaConsumer.Start()
}
