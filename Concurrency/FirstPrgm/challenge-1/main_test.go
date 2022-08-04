package main

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func Test_updateMessage(t *testing.T) {
	wg.Add(1)

	updateMessage("epsilon")
	wg.Wait()

	if msg != "epsilon" {
		t.Error("Expected to find epsilon, but it is not there")
	}
}

func Test_printMessage(t *testing.T) {
	stdout := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	msg = "epsilon"
	printMessage()

	_ = w.Close()

	result, _ := ioutil.ReadAll(r)
	output := string(result)

	os.Stdout= stdout

	if !strings.Contains(output, "epsilon") {
		t.Error("Expected to find epsilon, but it is not there")
	}
}

func Test_main(t *testing.T) {
	stdout := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	_ = w.Close()

	result, _ := ioutil.ReadAll(r)
	output := string(result)

	os.Stdout= stdout

	if !strings.Contains(output, "Hello, universe!") {
		t.Error("Expected to find epsilon, but it is not there")
	}

	if !strings.Contains(output, "Hello, cosmos!") {
		t.Error("Expected to find epsilon, but it is not there")
	}

	if !strings.Contains(output, "Hello, world!") {
		t.Error("Expected to find epsilon, but it is not there")
	}

}
