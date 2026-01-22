package tests

import (
	"github.com/andrewmzhang/pulumi-nextdns/nextdns"
	"github.com/blang/semver"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/integration"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/property"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAutoName(t *testing.T) {
	t.Parallel()
	provider := nextdns.Provider()

	s, err := integration.NewServer(t.Context(),
		"autoname",
		semver.MustParse("0.1.0"),
		integration.WithProvider(provider),
	)
	require.NoError(t, err)

	_, err = s.Check(p.CheckRequest{
		Urn: resource.CreateURN("name", "test:nextdns:NextDNSRewrite", "", "proj", "stack"),
		Inputs: property.NewMap(map[string]property.Value{
			"field": property.New("value"),
		}),
	})
	require.Error(t, err)

}
