package main

import (
	"bufio"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func Test_isPrime(t *testing.T) {
	primeTests := []struct{
		name string
		testNum int
		expected bool
		msg string
	}{
		{"prime", 7,true,"7 is a prime number!"},
		{"not prime", 8,false,"8 is not a prime because it is divisible by 2!"},
		{"zero", 0,false,"0 is not prime, by definition!"},
		{"one", 1,false,"1 is not prime, by definition!"},
		{"negative", -23,false,"-23, Negative number are not prime by definition!"},
	}

	for _, e := range primeTests {
		result, msg := isPrime(e.testNum)

		if e.expected && !result {
			t.Errorf("%s, expected true but got false", e.name)
		}

		if !e.expected && result {
			t.Errorf("%s, expected false but got true", e.name)
		}

		if e.msg != msg {
			t.Errorf("%s: expected %s but got %s", e.name, e.msg, msg)
		}
	}
}

func Test_prompt(t *testing.T) {
	//Save a copy of Stdout
	oldOut := os.Stdout

	//Create a read & write pipe
	r,w,_  := os.Pipe()

	//Set os.stdOut to our write pipe.
	os.Stdout = w

	prompt()

	//close the writer
	_ = w.Close()

	//reset os.StdOut
	os.Stdout = oldOut

	// read the output of our prompt from read pipe
	out, _ := ioutil.ReadAll(r)


	if string(out) != "==>" {
		t.Errorf("incorrect prompt: expected ==> but got %s", string(out))
	}
}

func Test_intro(t *testing.T) {
	//Save a copy of Stdout
	oldOut := os.Stdout

	//Create a read & write pipe
	r,w,_  := os.Pipe()

	//Set os.stdOut to our write pipe.
	os.Stdout = w

	intro()

	//close the writer
	_ = w.Close()

	//reset os.StdOut
	os.Stdout = oldOut

	// read the output of our prompt from read pipe
	out, _ := ioutil.ReadAll(r)

	if !strings.Contains(string(out), "Enter a whole number") {
		t.Errorf("Intro text not correct; got %s", string(out))
	}
}

func Test_checknumbers(t *testing.T) {
	tests := []struct {
		name  string
		input string
		expected string
	} {
		{ name:"empty", input: "", expected: "Please enter a whole number !!"},
		{ name:"zero", input: "0", expected: "0 is not prime, by definition!"},
		{ name:"one", input: "1", expected: "1 is not prime, by definition!"},
		{ name:"seven", input: "7", expected: "7 is a prime number!"},
		{ name:"Negative Number", input: "-7", expected: "-7, Negative number are not prime by definition!"},
		{ name:"Typed", input: "three", expected: "Please enter a whole number !!"},
		{ name:"Typed", input: "three", expected: "Please enter a whole number !!"},
		{ name:"Decimal", input: "2.13", expected: "Please enter a whole number !!"},
		{ name:"Quite", input: "q", expected: ""},
		{ name:"QUITE", input: "Q", expected: ""},
	}

	for _, e := range tests {

		input := strings.NewReader(e.input)
		reader := bufio.NewScanner(input)

		res, _ := checkNumbers(reader)

		if !strings.EqualFold(res, e.expected) {
			t.Errorf("%s: expected %s, but got %s", e.name, e.expected, res)
		}
	}
}
