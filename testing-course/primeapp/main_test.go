package main

import "testing"

func Test_isPrime(t *testing.T) {
	result, msg := isPrime(0)

	if result {
		t.Errorf("with 0, as test parameter, got true but expected false")
	}

	if msg != "0 is not prime, by definition!" {
		t.Error("wrong message returned: ", msg)
	}
}
