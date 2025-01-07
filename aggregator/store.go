package main

import (
	"fmt"
	"github/princedraculla/toll-calculation/types"
)

type MemoryStore struct{}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{}
}

func (m *MemoryStore) Insert(data types.Distance) error {
	fmt.Printf("data: %+v stored in memory", data)
	return nil
}
