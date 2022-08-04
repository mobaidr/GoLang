package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type Income struct {
	Source string
	Amount int
}

func main() {

	// variable for bank balance
	var bankBalance int
	var balanceMutex sync.Mutex

	// print out starting values
	fmt.Printf("Initial Account balance: $%d.00", bankBalance)
	fmt.Println()

	// Define weekly revenue
	incomes := []Income{
		{Source: "Main Job", Amount: 500},
		{Source: "Gifts", Amount: 10},
		{Source: "Part time Job", Amount: 50},
		{Source: "Investments", Amount: 100},
	}

	wg.Add(len(incomes))
	//loop thru 52 weeks and print out how is made..keep running total
	for i, income := range incomes {
		go func(i int,income Income){

			defer wg.Done()

			for week :=1; week <=52; week++ {
				balanceMutex.Lock()

				bankBalance += income.Amount

				balanceMutex.Unlock()
				fmt.Printf("On week %d, you earned $%d.00 from %s \n", week, income.Amount, income.Source)
			}
		}(i,income)
	}
	wg.Wait()
	//print out final balance
	fmt.Printf("Final Bank Balance: $%d.00",bankBalance)
	fmt.Println()
}