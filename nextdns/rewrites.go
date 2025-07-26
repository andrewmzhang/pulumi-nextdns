package nextdns

import (
	"context"
	"errors"
	"fmt"
	"github.com/amalucelli/nextdns-go/nextdns"
	"github.com/pulumi/pulumi-go-provider/infer"
	"os"
)


// NextDNSRewrite TODO: Call the master object NextDNSProfileRewrite
// NextDNSRewrite Resource Controlling struct
type NextDNSRewrite struct {
}

// NextDNSRewriteArgs Resource input struct
type NextDNSRewriteArgs struct {
	ProfileID string `pulumi:"profileId"`
	Name      string `pulumi:"name"`
	Content   string `pulumi:"content"`
}

type NextDNSRewriteState struct {
	RewriteID string `pulumi:"rewriteId"`
	NextDNSRewriteArgs
}

func (f *NextDNSRewrite) Annotate(a infer.Annotator) {
	a.Describe(&f, "A NextDNS Rewrite into a pulumi resource")
}

func (f *NextDNSRewriteArgs) Annotate(a infer.Annotator) {
	a.Describe(&f.ProfileID, "Profile Id to apply rewrite to.")
	a.Describe(&f.Name, "Domain name to apply rewrite to.")
	a.Describe(&f.Content, "IP Address or Domain to rewrite domain name to.")
}

func (f *NextDNSRewriteState) Annotate(a infer.Annotator) {
	a.Describe(&f.RewriteID, "RewriteID as represented in NextDNS")
	a.Describe(&f.ProfileID, "ProfileId of the profile this rewrite belongs to.")
	a.Describe(&f.Name, "Domain name to apply rewrite to.")
	a.Describe(&f.Content, "IP Address or Domain to rewrite domain name to.")
}

func (f *NextDNSRewrite) Create(ctx context.Context, req infer.CreateRequest[NextDNSRewriteArgs]) (resp infer.CreateResponse[NextDNSRewriteState], err error) {
	// Get nextdns config and retrieve NextDNS API Key
	config := infer.GetConfig[Config](ctx)
	apiKey := config.ApiKey

	if req.DryRun {
		return infer.CreateResponse[NextDNSRewriteState]{ID: ""}, nil
	}

	// Not dry run, create the NextDNSRewrite
	client, err := nextdns.New(nextdns.WithAPIKey(apiKey))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
		os.Exit(1)
	}
	_, err = client.Profiles.Get(ctx, &nextdns.GetProfileRequest{ProfileID: req.Inputs.ProfileID})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
		os.Exit(1)
	}
	createRewriteRequest := &nextdns.CreateRewritesRequest{
		ProfileID: req.Inputs.ProfileID,
		Rewrites: &nextdns.Rewrites{
			Name:    req.Inputs.Name,
			Content: req.Inputs.Content,
		},
	}

	rewriteID, err := client.Rewrites.Create(ctx, createRewriteRequest)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
		os.Exit(1)
	}

	return infer.CreateResponse[NextDNSRewriteState]{
		ID: rewriteID,
		Output: NextDNSRewriteState{
			RewriteID:          rewriteID,
			NextDNSRewriteArgs: req.Inputs,
		},
	}, nil
}

func (*NextDNSRewrite) Delete(ctx context.Context, req infer.DeleteRequest[NextDNSRewriteState]) (infer.DeleteResponse, error) {
	// Get nextdns config
	config := infer.GetConfig[Config](ctx)
	apiKey := config.ApiKey
	profileID := req.State.ProfileID

	// Not dry run, create the NextDNSRewrite
	client, err := nextdns.New(nextdns.WithAPIKey(apiKey))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
		os.Exit(1)
	}

	client.Rewrites.Delete(ctx, &nextdns.DeleteRewritesRequest{
		ProfileID: profileID,
		ID:        req.State.RewriteID,
	})

	return infer.DeleteResponse{}, nil
}

func (*NextDNSRewrite) Check(ctx context.Context, req infer.CheckRequest) (infer.CheckResponse[NextDNSRewriteArgs], error) {
	args, f, err := infer.DefaultCheck[NextDNSRewriteArgs](ctx, req.NewInputs)
	if err != nil {
		return infer.CheckResponse[NextDNSRewriteArgs]{
			Inputs:   args,
			Failures: f,
		}, err
	}

	// Get nextdns config
	config := infer.GetConfig[Config](ctx)
	apiKey := config.ApiKey
	client, err := nextdns.New(nextdns.WithAPIKey(apiKey))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
		os.Exit(1)
	}

	profileID := args.ProfileID
	if profileID != "" {
		profile, err := client.Profiles.Get(ctx, &nextdns.GetProfileRequest{ProfileID: profileID})
		if err != nil || profile == nil {
			fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
			return infer.CheckResponse[NextDNSRewriteArgs]{
				Inputs:   args,
				Failures: f,
			}, err
		}
	}

	return infer.CheckResponse[NextDNSRewriteArgs]{
		Inputs:   args,
		Failures: f,
	}, err
}

func (*NextDNSRewrite) Read(ctx context.Context, req infer.ReadRequest[NextDNSRewriteArgs, NextDNSRewriteState]) (infer.ReadResponse[NextDNSRewriteArgs, NextDNSRewriteState], error) {
	// Get nextdns config
	config := infer.GetConfig[Config](ctx)
	apiKey := config.ApiKey
	client, err := nextdns.New(nextdns.WithAPIKey(apiKey))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
		os.Exit(1)
	}

	rewrites, err := client.Rewrites.List(ctx, &nextdns.ListRewritesRequest{
		ProfileID: req.Inputs.ProfileID,
	})
	if err != nil {
		return infer.ReadResponse[NextDNSRewriteArgs, NextDNSRewriteState]{}, err
	}

	var rewriteFound *nextdns.Rewrites = nil
	for _, rewrite := range rewrites {
		if rewrite.ID == req.ID {
			rewriteFound = rewrite
			break
		}
	}
	if rewriteFound == nil {
		return infer.ReadResponse[NextDNSRewriteArgs, NextDNSRewriteState]{}, errors.New("rewrite does not exist")
	}

	return infer.ReadResponse[NextDNSRewriteArgs, NextDNSRewriteState]{
		ID: req.ID,
		Inputs: NextDNSRewriteArgs{
			ProfileID: req.Inputs.ProfileID,
			Name:      rewriteFound.Name,
			Content:   rewriteFound.Content,
		},
		State: NextDNSRewriteState{
			RewriteID: rewriteFound.ID,
			NextDNSRewriteArgs: NextDNSRewriteArgs{
				ProfileID: req.Inputs.ProfileID,
				Name:      rewriteFound.Name,
				Content:   rewriteFound.Content,
			},
		},
	}, nil
}

func (*NextDNSRewrite) WireDependencies(f infer.FieldSelector, args *NextDNSRewriteArgs, state *NextDNSRewriteState) {
	f.OutputField(&state.Content).DependsOn(f.InputField(&args.Content))
	f.OutputField(&state.Name).DependsOn(f.InputField(&args.Name))
	f.OutputField(&state.ProfileID).DependsOn(f.InputField(&args.ProfileID))
}
