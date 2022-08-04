package main

import (
	"fmt"
	"sync"
)

var msg1 string
var wg1 sync.WaitGroup

func updateMessageStr(s string) {
	defer wg1.Done()

	msg1 = s
}

func main() {
	msg1 = "Hello World!!!"


	wg1.Add(3)
	go updateMessageStr("Hello, Universe!!!")
	go updateMessageStr("Hello, World!!!")
	go updateMessageStr("Hello, Obaid!!!")

	wg1.Wait()

	fmt.Println(msg1)
}
