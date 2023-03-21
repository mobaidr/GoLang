package main

import (
	"fmt"
	"time"
)

func send(receiver User, message string) error {
	fmt.Println("Sending : ",message)
	time.Sleep(3*time.Second)
	return nil
}
