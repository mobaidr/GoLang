package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

//variables - philosophers,
var philosophers = []string{"Plato", "Didi", "Nana", "Aristotle", "Socrates"}
var wg sync.WaitGroup
var sleepTime = 1 * time.Second
var eatTime = 2 * time.Second
var thinkTime = 1 * time.Second

var finishingSequence []string
var finishingMutex = sync.Mutex{}

const hunger = 3

func diningProblem(philosopher string, leftFork, rightFork *sync.Mutex) {
    defer wg.Done()

    fmt.Println(philosopher, " is seated")
	time.Sleep(sleepTime)

    for i:=hunger; i> 0; i-- {
    	fmt.Println(philosopher, " is hungry")
    	time.Sleep(sleepTime)

    	// lock for both forks
    	rightFork.Lock()
		fmt.Printf("\t%s picked up the fork to his right\n", philosopher)

    	leftFork.Lock()
		fmt.Printf("\t%s picked up the fork to his left\n", philosopher)

    	fmt.Println(philosopher, " has both forks and is eating")
    	time.Sleep(eatTime)

    	fmt.Println(philosopher, "is thinking.")
    	
		time.Sleep(thinkTime)

    	leftFork.Unlock()
		fmt.Printf("\t%s put down the fork on his left.\n", philosopher)

    	rightFork.Unlock()
		fmt.Printf("\t%s put down the fork on his right.\n", philosopher)

    	time.Sleep(sleepTime)
	}

	fmt.Println(philosopher, " is satisfied.")
	time.Sleep(sleepTime)

	fmt.Println(philosopher, " has left the table.")

    finishingMutex.Lock()
	finishingSequence = append(finishingSequence,philosopher)
    finishingMutex.Unlock()
}

func main() {
	var mutexes []sync.Mutex

	for j:=0;j<len(philosophers);j++ {
		mutexes = append(mutexes,sync.Mutex{})
	}

	numMutexes := len(mutexes)

	// print intro
	fmt.Println("Dining Philosophers problem")
	fmt.Println("===========================")

	wg.Add(len(philosophers))

	var forkLeft sync.Mutex

	//spawn one go routine for each philosophers
	for i:= 0; i < len(philosophers); i++ {

		//go diningProblem(philosophers[i],forkRight, forkLeft)
		if i==0 {
			forkLeft = mutexes[numMutexes-1]
		} else {
			forkLeft = mutexes[i-1]
		}

		forkRight := mutexes[i]

		go diningProblem(philosophers[i],&forkLeft,&forkRight)
	}

	wg.Wait()

	fmt.Println("The table is empty")
	fmt.Println("Finishing Sequence")
	fmt.Printf("Ordered finished: %s\n",strings.Join(finishingSequence,", "))
}
