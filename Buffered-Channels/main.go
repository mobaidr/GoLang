package main

import (
	"fmt"
	"time"
)

func listenTOChan(ch chan int) {
	for {
		i := <- ch
		fmt.Println("Got ",i," from channel")

		//simulate doing lots of work.
		time.Sleep(1*time.Second)
	}
}

func main() {
	ch := make(chan int,10)

	go listenTOChan(ch)

	for i:=0;i<=100;i++ {
		fmt.Println("sending ",i," to channel")
		ch <- i
		fmt.Println("sent ",i," to channel")
	}

	fmt.Println("Done")

	close(ch)
}
