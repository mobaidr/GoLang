package main

import (
	"errors"
	"testing"
)

type mockSender struct {
	sendingError error
}

func (ms mockSender) Send(u User, orderId string) error {
	return ms.sendingError
}

func Test_InformOrderShipped(t *testing.T) {
	cases := []struct{
		user User
		orderId string
		sendingError error
		name string
		want bool
	}{
		{
			user:         User{Name: "Amena Rahman", Phone: "2145667834"},
			orderId:      "12355",
			sendingError: nil,
			name:         "Successful Send",
			want:         true,
		},
		{
			user:         User{Name: "Hiba Rahman", Phone: "9729887456"},
			orderId:      "66655",
			sendingError: errors.New("Sending failed"),
			name:         "UnSuccessful Send",
			want:         false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T){
			ms := mockSender{tc.sendingError}
			got := InformOrderShipped(tc.user,tc.orderId,ms)

			if tc.want != got {
				t.Errorf("Want %t, Got %t", tc.want,got)
			}
		})
	}
}
