package main

import (
	"fmt"
	"time"
)

type Sender interface {
	Send(receiver User, message string) error
}


type Dispatcher struct {

}

func (d Dispatcher) Send(receiver User, message string) error {
	fmt.Println("Sending : ",message)
	time.Sleep(3*time.Second)
	return nil
}
