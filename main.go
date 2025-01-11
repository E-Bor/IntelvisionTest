package main

import (
	"IntelvisionTest/server"
	"fmt"
	"os"
	"strconv"
)

func main() {
	numIn, numOut, err := extractArgs()
	if err {
		return
	}

	srv := server.NewServer(
		numIn,
		numOut,
	)
	srv.Start()

	// I don't know if we should get into an infinite loop here or just wait for something to happen
	defer srv.Stop()
}

func extractArgs() (
	int,
	int,
	bool,
) {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <numIn> <numOut>")
		return 0, 0, true
	}

	numIn, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf(
			"Invalid numIn: %v\n",
			err,
		)
		return 0, 0, true
	}
	numOut, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Printf(
			"Invalid numOut: %v\n",
			err,
		)
		return 0, 0, true
	}
	return numIn, numOut, false
}
