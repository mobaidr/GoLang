package main

import (
	"io/ioutil"
	"os"
	"strings"
	"sync"
	"testing"
)

func Test_printSomething(t *testing.T) {
	stdOut := os.Stdout

	r,w,_ := os.Pipe()

	os.Stdout = w

	var wg sync.WaitGroup

	wg.Add(1)

	go printSomething("Obaid", &wg)

	wg.Wait()

	_ = w.Close()

	result, _ := ioutil.ReadAll(r)

	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output,"Obaid") {
		t.Error("Expected to find Obaid, but it is not there")
	}
}
