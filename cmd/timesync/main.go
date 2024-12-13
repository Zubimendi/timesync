package main

import (
	"fmt"
	"os"

	"timesync/internal/cli"
)

func main() {
	err := cli.Execute()
	if err != nil {
		fmt.Fprint(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}