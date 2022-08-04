package main

import (
	"fmt"
	"github.com/fatih/color"
	"math/rand"
	"time"
)

//Variables
var seatingCapacity = 5
var arrivalRate = 100
var cuttingDuration = 1000 * time.Millisecond
var timeOpen = 10 * time.Second


func main() {
	//Seed our random generator
	rand.Seed(time.Now().UnixNano())

	//Print welcome message
	color.Yellow("The Sleeping Barber Problem")
	color.Yellow("===========================")

	//Create Channels
	clientChan := make(chan string, seatingCapacity)
	doneChan := make(chan bool)

	//Create data structure that represents Barber shop
	shop := BarberShop{
		ShopCapacity:    seatingCapacity,
		HairCutDuration: cuttingDuration,
		NumberOfBarbers: 0,
		BarbersDoneChan: doneChan,
		ClientsChan:     clientChan,
		Open:            true,
	}

	color.Green("The shop is OPEN for the day.")
	//Add Barbers
	shop.addBarber("Frank")
	shop.addBarber("Joseph")
	shop.addBarber("Latham")
	shop.addBarber("Mike")
	shop.addBarber("Chuck")

	//Start the barbershop as do routine
	shopClosing := make(chan bool)
	closed := make(chan bool)

	go func() {
		<-time.After(timeOpen)
		shopClosing <- true

		shop.closeShopForTheDay()

		closed <- true
	}()

	//Add clients
	i := 1

	go func() {
		for {
			// get random number  with average arrival rate
			randomMilliseconds := rand.Int() % (2 * arrivalRate)

			select {
			case <-shopClosing:
				{
					return
				}
			case <-time.After(time.Millisecond * time.Duration(randomMilliseconds)):
				{
					shop.addClient(fmt.Sprintf("Client #%d",i))
					i++
				}
			}
		}
	}()

	//Block until the barber shop is closed.
	<- closed

	close(closed)
	close(shopClosing)
}
