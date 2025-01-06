package main

import "github/princedraculla/toll-calculation/types"

type CalculateServicer interface {
	CalculateDistance(types.OBUData) (float64, error)
}

type CalculateService struct {
}

func NewCalculatorService() CalculateServicer {
	return &CalculateService{}
}

func (cs *CalculateService) CalculateDistance(data types.OBUData) (float64, error) {
	return 0.0, nil
}
