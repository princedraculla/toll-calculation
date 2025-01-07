package main

import (
	"github/princedraculla/toll-calculation/types"
	"time"

	"github.com/sirupsen/logrus"
)

type LogMiddleWareConsumer struct {
	next CalculateServicer
}

func NewLogMiddleWareConsumer(next CalculateServicer) *LogMiddleWareConsumer {
	return &LogMiddleWareConsumer{
		next: next,
	}
}

func (lm *LogMiddleWareConsumer) CalculateDistance(data types.OBUData) (dist float64, err error) {
	defer func(start time.Time) {
		logrus.WithFields(logrus.Fields{
			"took":     time.Since(start),
			"error":    err,
			"distance": dist,
		}).Info("calculation of distance")
	}(time.Now())

	return lm.next.CalculateDistance(data)
}
