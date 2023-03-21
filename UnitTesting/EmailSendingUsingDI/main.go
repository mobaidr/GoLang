package main

import "fmt"

type User struct {
	Name string
	Phone string
}

func main() {
	d := Dispatcher{}
	InformOrderShipped(User{Name: "Amena Rahman", Phone: "1232343344"},"223344", d)
}

func InformOrderShipped(receiver User, orderID string, sender Sender) bool {
	message := fmt.Sprintf("Your order #%s is shipped", orderID)
	err := sender.Send(receiver,message)

	if err != nil {
		return false
	}

	return true
}
