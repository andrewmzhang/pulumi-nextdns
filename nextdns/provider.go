package nextdns

import (
	"context"

	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
)

type Config struct {
	ApiKey string `pulumi:"apiKey" provider:"secret"`
}

// ClientFactory is a function that creates a Client based on the provider config.
type ClientFactory func(ctx context.Context, config Config) (Client, error)

func Provider(clientFactory ClientFactory) (p.Provider, error) {

	return infer.NewProviderBuilder().
		WithDisplayName("pulumi-nextdns").
		WithDescription("NextDNS provider for Pulumi").
		WithResources(
			infer.Resource(&NextDNSRewrite{getClient: clientFactory}),
		).
		WithConfig(infer.Config(&Config{})).
		WithNamespace("andrewmzhang").
		Build()
}
