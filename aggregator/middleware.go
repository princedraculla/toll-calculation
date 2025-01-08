package main

import (
	"github/princedraculla/toll-calculation/types"
	"time"

	"github.com/sirupsen/logrus"
)

type LogMiddleWareAggregator struct {
	next Aggregator
}

func NewLogMiddleWareAggregator(next Aggregator) Aggregator {
	return &LogMiddleWareAggregator{
		next: next,
	}
}

func (lm *LogMiddleWareAggregator) AggregateDistance(dist types.Distance) (err error) {
	defer func(start time.Time) {
		logrus.WithFields(logrus.Fields{
			"took":  time.Since(start),
			"error": err,
		}).Info("aggregate distance")
	}(time.Now())
	err = lm.next.AggregateDistance(dist)
	return err
}
