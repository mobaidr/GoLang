//+build wireinject

package main

import "github.com/google/wire"

func InitializeEvent() (Event, error) {
	wire.Build(GetMessage, GetGreeter, GetEvent)
	return Event{}, nil
}
