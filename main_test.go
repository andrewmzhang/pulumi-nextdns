package main

import (
	"context"
	"testing"

	"github.com/andrewmzhang/nextdns-go/models"
	"github.com/blang/semver"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/integration"
	"github.com/pulumi/pulumi/sdk/v3/go/property"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	gomock "go.uber.org/mock/gomock"
)

//go:generate mockgen -typed -package main -destination mocks.gen.go -imports goprovider=github.com/pulumi/pulumi-go-provider ./nextdns Client

func TestNextDNSRewrite(t *testing.T) {
	// Configure a mock client to return a fake widget ID when CreateRewrite is called.
	ctrl := gomock.NewController(t)
	mockClient := NewMockClient(ctrl)
	mockClient.EXPECT().CreateRewrite(gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
		func(profileId string, name string, content string) (*models.Rewrite, error) {
			return &models.Rewrite{
				ID:      "fake-rewrite-id",
				Name:    name,
				Type:    "A",
				Content: content,
			}, nil
		}).AnyTimes()
	mockClient.EXPECT().CheckRewrite(gomock.Any()).DoAndReturn(
		func(profileId string) error {
			return nil
		}).AnyTimes()
	mockClient.EXPECT().DeleteRewrite(gomock.Any(), gomock.Any()).DoAndReturn(
		func(profileId string, rewriteId string) error {
			return nil
		}).AnyTimes()

	// Create the provider such that it uses the mock client.
	newMockClient := func(ctx context.Context, config Config) (Client, error) {
		assert.Equal(t, "nextdns-api-key-mock", config.ApiKey)
		return mockClient, nil
	}
	provider, err := Provider(newMockClient)
	require.NoError(t, err)

	server, err := integration.NewServer(t.Context(),
		"nextdns",
		semver.MustParse("0.1.0"),
		integration.WithProvider(provider),
	)
	require.NoError(t, err)

	// Configure the provider with a fake client key and secret.
	err = server.Configure(p.ConfigureRequest{
		Args: property.NewMap(map[string]property.Value{
			"apiKey": property.New("nextdns-api-key-mock").WithSecret(true),
		}),
	})
	require.NoError(t, err)

	// Test the lifecycle methods of the Rewrite resource, expecting it to use the mock client.
	integration.LifeCycleTest{
		Resource: "test:nextdns:NextDNSRewrite",
		Create: integration.Operation{
			Inputs: property.NewMap(map[string]property.Value{
				"profileId": property.New("profile-id"),
				"name":      property.New("example.com"),
				"content":   property.New("content.example.com"),
			}),
			Hook: func(inputs, output property.Map) {
				t.Logf("Outputs: %#v", output)
				name := output.Get("name").AsString()
				assert.Equal(t, "example.com", name)
			},
		},
	}.Run(t, server)
}
