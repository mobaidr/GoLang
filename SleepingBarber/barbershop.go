package main

import (
	"github.com/fatih/color"
	"time"
)

type BarberShop struct {
	ShopCapacity    int
	HairCutDuration time.Duration
	NumberOfBarbers int
	BarbersDoneChan chan bool
	ClientsChan     chan string
	Open            bool
}

func (shop *BarberShop) addBarber(barber string) {
	shop.NumberOfBarbers++

	go func () {
		isSleeping := false

		color.Yellow("%s goes to the waiting to check for clients", barber)

		for {
			// if there are no clients, barber goes to sleep
			if len(shop.ClientsChan) == 0 {
				color.Yellow("There is nothing to do so %s takes a nap", barber)
				isSleeping = true
			}

			client, shopOpen := <- shop.ClientsChan

			if shopOpen {
				if isSleeping {
					color.Yellow("%s wakes %s up", client, barber)
					isSleeping = false
				}

				//Cut the client's hair
				shop.CutHair(barber, client)
			} else {
				// Shop is closed, so send the barber home and close this go routine
				shop.sendBarberHome(barber)

				return
			}
		}
	}()
}

func (shop *BarberShop) CutHair(barber, client string) {
	color.Green("%s is cutting %s's hair", barber, client)
	time.Sleep(shop.HairCutDuration)

	color.Green("%s is finished cutting %s's hair", barber, client)
}

func (shop *BarberShop) sendBarberHome(barber string) {
	color.Red("%s is going home", barber)

	shop.BarbersDoneChan <- true
}

func (shop *BarberShop) closeShopForTheDay() {
	color.Cyan("Closing shop for the day")
	close(shop.ClientsChan)
	shop.Open = false

	for a :=1; a <= shop.NumberOfBarbers; a++ {
		<- shop.BarbersDoneChan
	}

	close(shop.BarbersDoneChan)

	color.Green("---------------------------------------------------------------------")
	color.Green("The Barber shop is now closed for the day, and everyone has gone home")
	color.Green("---------------------------------------------------------------------")
}

func (shop *BarberShop) addClient(client string) {
	color.Green("***** %s arrives",client)

	if shop.Open {
		select {
			case shop.ClientsChan <- client : {
				color.Yellow("%s takes a seat in the waiting room", client)
			}
			default: {
				color.Red("Waiting room is full, so %s leaves", client)
			}
		}
	} else {
		color.Red("The shop is already closed, so %s leaves!", client)
	}
}