package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	//Print a welcome message
	intro()

	//Create a channel to indicate when user wants to quit
	doneChan := make(chan bool)

	//start to goroutine to read user input and run program
	go readUserInput(os.Stdin, doneChan)

	//block until the done chan gets a value.
	<-doneChan

	//close the channel
	close(doneChan)

	//say goodbye.
	fmt.Println("Goodbye.")
}

func readUserInput(in io.Reader, doneChan chan bool) {
	scanner := bufio.NewScanner(in)

	for {
		res, done := checkNumbers(scanner)

		if done {
			doneChan <- true
			return
		}

		fmt.Println(res)
		prompt()
	}
}

func checkNumbers(scanner *bufio.Scanner) (string, bool) {
	// read user input
	scanner.Scan()

	// check to see if user wants to quit
	if strings.EqualFold(scanner.Text(), "q") {
		return "", true
	}

	//Try to convert what user typed into int
	numToCheck, err := strconv.Atoi(scanner.Text())

	if err != nil {
		return "Please enter a whole number !!", false
	}

	_, msg := isPrime(numToCheck)

	return msg, false
}

func intro() {
	fmt.Println("Is it Prime ?")
	fmt.Println("=============")

	fmt.Println("Enter a whole number and we will tell you if it is a prime number or not. Enter q to quit")

	prompt()
}

func prompt() {
	fmt.Print("==>")
}

func isPrime(n int) (bool, string) {

	// 0 and 1 are not prime
	if n == 0 || n == 1 {
		return false, fmt.Sprintf("%d is not prime, by definition!", n)
	}

	if n < 0 {
		return false, fmt.Sprintf("%d, Negative number are not prime by definition!", n)
	}

	// use the modulus operator repeatedly to see if we have a prime number
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return false, fmt.Sprintf("%d is not a prime because it is divisible by %d!", n, i)
		}
	}

	return true, fmt.Sprintf("%d is a prime number!", n)
}
