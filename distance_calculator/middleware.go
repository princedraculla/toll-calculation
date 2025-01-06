package main

type LogMiddleWareConsumer struct {
	next DataConsumer
}

func NewLogMiddleWareConsumer(next DataConsumer) *LogMiddleWareConsumer {
	return &LogMiddleWareConsumer{
		next: next,
	}
}

func (lc *LogMiddleWareConsumer) ConsumerData(data any) error {
	return nil
}
