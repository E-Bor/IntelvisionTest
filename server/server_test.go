package server

import (
	"fmt"
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

	t.Run(
		"Positive cases",
		func(t *testing.T) {
			t.Run(
				"Read from valid IN-ports",
				func(t *testing.T) {
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
				},
			)

			t.Run(
				"Write to valid OUT-ports",
				func(t *testing.T) {
					for i := 1; i <= numOut; i++ {
						transactionID := i
						value := i % 2

						err := srv.Write(
							i,
							transactionID,
							value,
						)
						if err != nil {
							t.Fatalf(
								"Error while writing to OUT-port #%d: %v",
								i,
								err,
							)
						}
						t.Logf(
							"Wrote to OUT-port #%d: transactionID=%d, value=%d",
							i,
							transactionID,
							value,
						)
					}
				},
			)
		},
	)
	invalidPorts := []int{
		0,
		-1,
		numIn + 1,
		numIn + 10,
	}

	t.Run(
		"Negative cases",
		func(t *testing.T) {
			t.Run(
				"Read with invalid port numbers",
				func(t *testing.T) {
					for _, portNum := range invalidPorts {
						t.Run(
							fmt.Sprintf(
								"IN-port #%d",
								portNum,
							),
							func(t *testing.T) {
								_, err := srv.Read(portNum)
								if err == nil {
									t.Fatalf(
										"Expected error for IN-port #%d, but got nil",
										portNum,
									)
								}
								t.Logf(
									"Got expected error for IN-port #%d: %v",
									portNum,
									err,
								)
							},
						)
					}
				},
			)

			t.Run(
				"Write with invalid port numbers",
				func(t *testing.T) {
					for _, portNum := range invalidPorts {
						t.Run(
							fmt.Sprintf(
								"OUT-port #%d",
								portNum,
							),
							func(t *testing.T) {
								transactionID := 999
								value := 1
								err := srv.Write(
									portNum,
									transactionID,
									value,
								)
								if err == nil {
									t.Fatalf(
										"Expected error for OUT-port #%d, but got nil",
										portNum,
									)
								}
								t.Logf(
									"Got expected error for OUT-port #%d: %v",
									portNum,
									err,
								)
							},
						)
					}
				},
			)
		},
	)
}
