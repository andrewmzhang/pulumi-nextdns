package main

import (
	"context"
	"fmt"
	"os"

	"github.com/andrewmzhang/pulumi-nextdns/nextdns"
)

// Version is initialized by the Go linker to contain the semver of this build.
var Version string

// Name controls how this nextdns is referenced in package names and elsewhere.
const Name string = "nextdns"

func main() {
	provider, err := nextdns.Provider(nextdns.NewRealClient)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
		os.Exit(1)
	}
	err = provider.Run(context.Background(), "nextdns", "0.1.0")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
		os.Exit(1)
	}
}
