package main

import (
	"github/princedraculla/toll-calculation/types"
	"math"
)

type CalculateServicer interface {
	CalculateDistance(types.OBUData) (float64, error)
}

type CalculateService struct {
	points [][]float64
}

func NewCalculatorService() CalculateServicer {
	return &CalculateService{
		points: make([][]float64, 0),
	}
}

func (cs *CalculateService) CalculateDistance(data types.OBUData) (float64, error) {
	distance := 0.0

	if len(cs.points) > 0 {
		prevPoint := cs.points[len(cs.points)-1]
		distance = calculateDistance(prevPoint[0], prevPoint[1], data.Lat, data.Long)
	}
	cs.points = append(cs.points, []float64{data.Lat, data.Long})

	return distance, nil
}

func calculateDistance(x1, x2, y1, y2 float64) float64 {
	return math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
}
