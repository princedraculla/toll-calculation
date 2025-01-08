package main

import (
	"github/princedraculla/toll-calculation/types"
)

const BasePrice = 3.15

type Aggregator interface {
	AggregateDistance(types.Distance) error
	CalculateInvoice(int) (*types.Invoice, error)
}

type Store interface {
	Insert(types.Distance) error
	Get(int) (float64, error)
}

type InvoiceAggregator struct {
	store Store
}

func NewInvoiceAggregator(storage Store) Aggregator {
	return &InvoiceAggregator{
		store: storage,
	}
}

func (i *InvoiceAggregator) AggregateDistance(distance types.Distance) error {
	return i.store.Insert(distance)
}

func (i *InvoiceAggregator) CalculateInvoice(id int) (*types.Invoice, error) {
	dist, err := i.store.Get(id)
	if err != nil {
		return nil, err
	}
	inv := &types.Invoice{
		OBUID:         id,
		TotalDistance: dist,
		TotalAmount:   BasePrice * dist,
	}
	return inv, nil
}
