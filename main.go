package main

import (
	"context"
	"fmt"
	"github.com/pulumi/pulumi-go-provider/infer"
	"os"
)

func main() {
	provider, err := infer.NewProviderBuilder().
		WithResources(infer.Resource(&NextDNSRewrite{})).
		WithConfig(infer.Config(&Config{})).
		WithNamespace("andrewmzhang").
		Build()

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

type Config struct {
	ApiKey string `pulumi:"apiKey"`
}
