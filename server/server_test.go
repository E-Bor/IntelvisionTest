package server

import (
	"testing"
)

func TestServerReadWrite(t *testing.T) {
	numIn, numOut := 3, 2

	srv := NewServer(
		numIn,
		numOut,
	)
	srv.Start()
	defer srv.Stop()

	for i := 1; i <= numIn; i++ {
		val, err := srv.Read(i)
		if err != nil {
			t.Fatalf(
				"Error while reading IN-port #%d: %v",
				i,
				err,
			)
		}
		t.Logf(
			"Read from IN-port #%d: %d",
			i,
			val,
		)
	}

	for i := 1; i <= numOut; i++ {
		if err := srv.Write(
			i,
			i,
			i%2,
		); err != nil {
			t.Fatalf(
				"error OUT-порт #%d: %v",
				i,
				err,
			)
		}
		t.Logf(
			"write to OUT-port #%d: transactionID=%d, value=%d",
			i,
			i,
			i%2,
		)
	}
}
