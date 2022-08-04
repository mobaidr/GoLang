package main

import (
	"fmt"
	"github.com/fatih/color"
	"math/rand"
	"time"
)

const (
	NumberOfPizzas = 10
)

var pizzasMade, pizzasFailed, total int

type Producer struct {
	data chan PizzaOrder
	quit chan chan error
}

type PizzaOrder struct{
	pizzaNumber int
	message string
	success bool
}

func (p *Producer) Close() error {
	ch:= make(chan error)

	p.quit <- ch

	return <- ch
}

func makePizza(pizzaNumber int) *PizzaOrder {
	pizzaNumber++

	if pizzaNumber <= NumberOfPizzas {
		delay := rand.Intn(5) + 1
		fmt.Printf("Received order number %d\n", pizzaNumber)

		rnd := rand.Intn(12) + 1
		msg := ""
		success := false

		if rnd < 5 {
			pizzasFailed++
		} else {
			pizzasMade++
		}

		total++

		fmt.Printf("Making Pizza number %d it will take %d seconds \n", pizzaNumber, delay)

		//delay for a bit
		time.Sleep(time.Duration(delay) * time.Second)


		if rnd <= 2 {
			msg = fmt.Sprintf("*** We ran out of ingredients for pizza #%d", pizzaNumber)
		} else if rnd <= 4 {
			msg = fmt.Sprintf("*** The cook quite while making pizza #%d", pizzaNumber)
		} else {
			success = true
			msg = fmt.Sprintf("Pizza order #%d is ready", pizzaNumber)
		}

		return &PizzaOrder{
			pizzaNumber: pizzaNumber,
			message:     msg,
			success:     success,
		}
	}

	return &PizzaOrder{pizzaNumber: pizzaNumber}
}

func pizzeria(pizzaMaker *Producer) {
	// Keep track of which pizza are we making
	var i = 0

	//try to make pizzas

	for {
		currentPizza := makePizza(i)

		if currentPizza != nil {
			i = currentPizza.pizzaNumber
			select {
				//We tried to make a pizza (we sent something to the data channel)
				case pizzaMaker.data <- *currentPizza: {}
				case quitChan := <- pizzaMaker.quit:
					{
						close(pizzaMaker.data)
						close(quitChan)
						return
					}
			}
		}
	}
}

func main() {
	// Producers
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	//print out the message that program has started
	fmt.Println("The Pizzeria is Open for business")
	fmt.Println("---------------------------------")

	//create a producer
	pizzaJob := &Producer{
		data: make(chan PizzaOrder),
		quit: make(chan chan error),
	}

	// run the producer in the background
	go pizzeria(pizzaJob)

	// Create and run consumer
	for p := range  pizzaJob.data {
		if p.pizzaNumber <= NumberOfPizzas {
			if p.success {
				color.Green(p.message)
				color.Green("Order # %d is out for delivery", p.pizzaNumber)
			} else {
				color.Red(p.message)
				color.Red("Customer is really mad")
			}
		} else {
			color.Yellow("Done making pizzas")
			err := pizzaJob.Close()
			if err != nil {
				color.Red("*** Error closing Channel!", err)
			}
		}
	}

	// print out the ending message
	color.Green("================")
	color.Green("Done for the day")

	color.Green("We made %d pizzas, but failed to make %d, with %d attempts in total", pizzasMade,pizzasFailed,total)

	switch {
	case pizzasFailed > 9 :
		color.Red("It was an awful day")
	case pizzasFailed >= 6:
		color.Red("It was not a very good day")
	case pizzasFailed >=4:
		color.Yellow("It was an OK day")
	case pizzasFailed >=2:
		color.Green("It was great day.")

	}
}
