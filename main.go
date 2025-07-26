package main

import (
	"context"
	"github.com/andrewmzhang/pulumi-nextdns/nextdns"
)

// Version is initialized by the Go linker to contain the semver of this build.
var Version string

// Name controls how this nextdns is referenced in package names and elsewhere.
const Name string = "nextdns"


func main() {
	nextdns.Provider().Run(context.Background(), "nextdns", "0.1.0")
}
