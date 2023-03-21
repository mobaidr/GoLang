package main

import "testing"

func Test_InformOrderShipped(t *testing.T) {
	user := User {
		Name: "Peggy",
		Phone: "+1 203 456 2233",
	}
	orderId := "1234"

	got := InformOrderShipped(user, orderId)
	want := true

	if got != want {
		t.Errorf("Want %t, Got %t", want, got)
	}
}
