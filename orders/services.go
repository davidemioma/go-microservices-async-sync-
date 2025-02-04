package main

import "context"

type Services struct {
	store OrderStore
}

func NewService(store OrderStore) *Services{
	return &Services{
		store,
	}
}

func (s *Services) CreateOrder (context.Context) error {
	return nil
}