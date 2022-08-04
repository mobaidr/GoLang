package main

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func Test_main(t *testing.T) {
	stdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	_ = w.Close()

	result, _ := ioutil.ReadAll(r)
	output := string(result)

	os.Stdout = stdout

	if ! strings.Contains(output,"$34320.00") {
		t.Error("Wrong Balance returned.")
	}
}
