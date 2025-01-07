package main

import (
	"github/princedraculla/toll-calculation/types"
	"math"
)

type CalculateServicer interface {
	CalculateDistance(types.OBUData) (float64, error)
}

type CalculateService struct {
	prevPoints []float64
}

func NewCalculatorService() CalculateServicer {
	return &CalculateService{}
}

func (cs *CalculateService) CalculateDistance(data types.OBUData) (float64, error) {
	distance := 0.0

	if len(cs.prevPoints) > 0 {

		distance = calculateDistance(cs.prevPoints[0], cs.prevPoints[1], data.Lat, data.Long)
	}
	cs.prevPoints = []float64{data.Lat, data.Lat}
	return distance, nil
}

func calculateDistance(x1, x2, y1, y2 float64) float64 {
	return math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
}
