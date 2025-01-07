package main

import (
	"github/princedraculla/toll-calculation/types"
	"time"

	"github.com/sirupsen/logrus"
)

type LogMiddleWare struct {
	next DataProducer
}

func NewLogMiddleWare(next DataProducer) *LogMiddleWare {
	return &LogMiddleWare{
		next: next,
	}
}

func (lp *LogMiddleWare) ProduceData(data *types.OBUData) error {
	defer func(start time.Time) {
		logrus.WithFields(logrus.Fields{
			"ObuId": data.ObuID,
			"lat":   data.Lat,
			"long":  data.Long,
			"took":  time.Since(start),
		}).Info("producing data from kafka")
	}(time.Now())
	return lp.next.ProduceData(data)
}
