package main

type DataConsumer interface {
	ConsumerData() error
}

type KafkaConsumer struct {
	Consumer DataConsumer
}

func NewKafkaConsumer() (DataConsumer, error) {
	return nil, nil
}

func (c *KafkaConsumer) ConsumerData(data any) error {
	return nil
}
