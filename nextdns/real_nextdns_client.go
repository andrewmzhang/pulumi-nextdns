package nextdns

import (
	"context"
	"errors"

	"github.com/andrewmzhang/nextdns-go"
	"github.com/andrewmzhang/nextdns-go/models"
	"github.com/andrewmzhang/nextdns-go/services"
)

// RewriteClient is a (helper) interface that interacts with nextdns.client to create Rewrites on NextDNS
type RewriteClient interface {
	CreateRewrite(profileId string, name string, content string) (*models.Rewrite, error)
	GetRewrite(profileId string, rewriteId string) (*models.Rewrite, error)
	DeleteRewrite(profileId string, rewriteId string) error
	CheckRewrite(profileId string) error
}

// Client implements all <Resource>Client interfaces
type Client interface {
	RewriteClient
}

// realClient is concrete class that implements various NextDNS<Resource>Client interfaces. This class is not meant to
// persist for long, so context is cached
type realClient struct {
	ctx    context.Context
	client *services.NextDNSClient
}

// Compile time check
var _ Client = realClient{}

// NewRealClient implements clientFactory and creates a real client based on the provider config.
func NewRealClient(ctx context.Context, config Config) (Client, error) {
	nextDNSClient, err := nextdns.NewClient(services.WithAPIKey(config.ApiKey))
	if err != nil {
		return nil, err
	}
	return &realClient{
		ctx:    ctx,
		client: nextDNSClient,
	}, nil
}

func (r realClient) GetRewrite(profileId string, rewriteId string) (*models.Rewrite, error) {
	rewrite, err := r.client.Rewrites(profileId).Bind(rewriteId).Get(r.ctx)
	if err != nil {
		return nil, err
	}
	return rewrite, nil
}

func (r realClient) CheckRewrite(profileId string) error {
	// Arguments check
	if profileId == "" {
		return errors.New("no profile id provided")
	}
	// Check that profile is not nil
	_, err := r.client.Profiles().Bind(profileId).Get(r.ctx)
	if err != nil {
		return err
	}
	return nil
}

func (r realClient) CreateRewrite(profileId string, name string, content string) (*models.Rewrite, error) {
	// Create the NextDNSRewrite
	_ = r.client.Profiles().Bind(profileId)
	rewrite, err := r.client.Rewrites(profileId).Create(r.ctx, models.Rewrite{
		Name:    name,
		Content: content,
	})
	if err != nil {
		return nil, err
	}
	return rewrite, nil
}

func (r realClient) DeleteRewrite(profileId string, rewriteId string) error {
	err := r.client.Rewrites(profileId).Bind(rewriteId).Delete(r.ctx)
	if err != nil {
		return err
	}
	return nil
}
