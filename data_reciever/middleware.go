package main

import (
	"fmt"
	"github/princedraculla/toll-calculation/types"
	"time"

	"github.com/sirupsen/logrus"
)

type LogMiddleWare struct {
	next DataProducer
}

func NewLogMiddleWare(next DataProducer) DataProducer {
	return &LogMiddleWare{
		next: next,
	}
}

func (l *LogMiddleWare) ProduceData(data *types.OBUData) error {
	defer func(start time.Time) {
		fmt.Println("produce data function in log middleware")
		logrus.WithFields(logrus.Fields{
			"ObuId": data.ObuID,
			"lat":   data.Lat,
			"long":  data.Long,
			"took":  time.Since(start),
		}).Info("producing data from kafka")
	}(time.Now())
	return l.next.ProduceData(data)
}
