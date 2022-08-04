package main

import (
	"fmt"
	"sync"
)

func printSomething(str string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(str)
}

func main() {

	var wg sync.WaitGroup
	words := []string{
		"Alpha",
		"Beta",
		"delta",
		"gamma",
		"pi",
		"zeta",
		"eta",

		"theta",
	}

	wg.Add(len(words))

	 for i , x := range words {
	 	go printSomething(fmt.Sprintf("%d: %s", i, x), &wg)
	 }

	 wg.Wait()

	 wg.Add(1)

	printSomething("This is Amena", &wg)
}

