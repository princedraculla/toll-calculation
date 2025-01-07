package main

import "github/princedraculla/toll-calculation/types"

type Aggregator interface {
	AggregateDistance(types.Distance) error
}

type Store interface {
	Insert(types.Distance) error
}

type InvoiceAggregator struct {
	store Store
}

func NewInvoiceAggregator(storage Store) *InvoiceAggregator {
	return &InvoiceAggregator{
		store: storage,
	}
}

func (i *InvoiceAggregator) AggregateDistance(distance types.Distance) error {
	return i.store.Insert(distance)
}
