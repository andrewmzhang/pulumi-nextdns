package nextdns

import (
	"fmt"

	"github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
)

type Config struct {
	ApiKey string `pulumi:"apiKey" provider:"secret"`
}

func Provider() provider.Provider {
	p, err := infer.NewProviderBuilder().
		WithDisplayName("pulumi-nextdns").
		WithDescription("NextDNS provider for Pulumi").
		WithResources(infer.Resource(&NextDNSRewrite{})).
		WithConfig(infer.Config(&Config{})).
		WithNamespace("andrewmzhang").
		Build()
	if err != nil {
		panic(fmt.Errorf("unabled to build a provider: %w", err))
	}
	return p
}
