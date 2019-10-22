package main

import (
	"fmt"
	"os"

	"justanother.org/cubedengine/internal/network"
)

func main() {
	if err := network.StartHTTPServer(); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to start http server: %v\n", err)
		os.Exit(1)
	}
}
