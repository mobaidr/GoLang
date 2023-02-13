package main

import "fmt"

func main() {
	n := -98

	_, msg := isPrime(n)

	fmt.Println(msg)
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
