package main

import (
	"testing"
	"time"
)

func Test_main(t *testing.T) {
	eatTime = 0 * time.Second
	thinkTime = 0 * time.Second
	sleepTime = 0 * time.Second

	for i:=0; i<100; i++ {
		main()

		if len(finishingSequence) != 5 {
			t.Error("Wrong number of entries in slice")
		}

		finishingSequence = []string{}
	}
}
