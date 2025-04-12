//go:build wireinject

package main

import "github.com/google/wire"

var providerSet wire.ProviderSet = wire.NewSet(wire.Struct(new(Message), "Content", "FCode"), NewGreeter)

func InitializeEvent(content string, code int, fCode float64) (Event, error) {

	wire.Build(providerSet, NewEvent)
	return Event{}, nil
}
