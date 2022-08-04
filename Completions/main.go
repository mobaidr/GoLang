package main

import (
	"encoding/json"
	"fmt"
)

type message struct {
	Field1 int
	Field2 string
}

//AddToField
func (m *message) AddToField(field1 int, field2 string) {
	m.Field1=field1
	m.Field2=field2
}

func main() {
	msg := message {
		Field1: 1,
		Field2: "Obaid",
	}

	marshal, err := json.Marshal(msg)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(marshal)
}
