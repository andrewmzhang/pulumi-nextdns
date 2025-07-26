package tests

import (
	"github.com/andrewmzhang/pulumi-nextdns/nextdns/nextdns"
	"github.com/blang/semver"
	"github.com/pulumi/pulumi-go-provider/integration"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAutoName(t *testing.T) {
	t.Parallel()
	provider := nextdns.Provider()

	_, err := integration.NewServer(t.Context(),
		"autoname",
		semver.MustParse("0.1.0"),
		integration.WithProvider(provider),
	)
	require.NoError(t, err)

}
