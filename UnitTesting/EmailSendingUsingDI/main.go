package main

import "fmt"

type User struct {
	Name string
	Phone string
}

func main() {
	InformOrderShipped(User{Name: "Mohammed O Rahman", Phone: "9581055556"}, "12334")
}

func InformOrderShipped(receiver User, orderID string) bool {
	message := fmt.Sprintf("Your order #%s is shipped", orderID)
	err := send(receiver,message)

	if err != nil {
		return false
	}

	return true
}
