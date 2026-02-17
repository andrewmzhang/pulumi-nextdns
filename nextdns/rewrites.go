package nextdns

import (
	"context"

	"github.com/pulumi/pulumi-go-provider/infer"
)

// NextDNSRewrite is Rewrite Resource Controlling struct
// Golang lacks function covariance, but ClientFactory produces Client which is implements RewriteClient
type NextDNSRewrite struct {
	getClient ClientFactory
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
	// Get nextdns config
	config := infer.GetConfig[Config](ctx)

	// Use the client factory to create a client based on the current config.
	client, err := f.getClient(ctx, config)
	if err != nil {
		return infer.CreateResponse[NextDNSRewriteState]{}, err
	}

	// Do nothing on dry-run
	if req.DryRun {
		return infer.CreateResponse[NextDNSRewriteState]{ID: ""}, nil
	}

	rewrite, err := client.CreateRewrite(req.Inputs.ProfileID, req.Inputs.Name, req.Inputs.Content)
	if err != nil {
		return infer.CreateResponse[NextDNSRewriteState]{}, err
	}

	return infer.CreateResponse[NextDNSRewriteState]{
		ID: rewrite.ID,
		Output: NextDNSRewriteState{
			RewriteID:          rewrite.ID,
			NextDNSRewriteArgs: req.Inputs,
		},
	}, nil
}

func (f *NextDNSRewrite) Delete(ctx context.Context, req infer.DeleteRequest[NextDNSRewriteState]) (infer.DeleteResponse, error) {
	// Get nextdns config
	config := infer.GetConfig[Config](ctx)

	// Use the client factory to create a client based on the current config.
	client, err := f.getClient(ctx, config)
	if err != nil {
		return infer.DeleteResponse{}, err
	}

	err = client.DeleteRewrite(req.State.ProfileID, req.State.RewriteID)
	if err != nil {
		return infer.DeleteResponse{}, err
	}

	return infer.DeleteResponse{}, nil
}

func (n *NextDNSRewrite) Check(ctx context.Context, req infer.CheckRequest) (infer.CheckResponse[NextDNSRewriteArgs], error) {
	args, f, err := infer.DefaultCheck[NextDNSRewriteArgs](ctx, req.NewInputs)
	if err != nil {
		return infer.CheckResponse[NextDNSRewriteArgs]{
			Inputs:   args,
			Failures: f,
		}, err
	}

	// Get nextdns config
	config := infer.GetConfig[Config](ctx)

	// Use the client factory to create a client based on the current config.
	client, err := n.getClient(ctx, config)
	if err != nil {
		return infer.CheckResponse[NextDNSRewriteArgs]{}, err
	}
	// Check that profile exists
	err = client.CheckRewrite(args.ProfileID)
	if err != nil {
		return infer.CheckResponse[NextDNSRewriteArgs]{}, err
	}
	return infer.CheckResponse[NextDNSRewriteArgs]{
		Inputs:   args,
		Failures: f,
	}, err
}

func (n *NextDNSRewrite) Read(ctx context.Context, req infer.ReadRequest[NextDNSRewriteArgs, NextDNSRewriteState]) (infer.ReadResponse[NextDNSRewriteArgs, NextDNSRewriteState], error) {
	// Get nextdns config
	config := infer.GetConfig[Config](ctx)
	// Use the client factory to create a client based on the current config.
	client, err := n.getClient(ctx, config)
	if err != nil {
		return infer.ReadResponse[NextDNSRewriteArgs, NextDNSRewriteState]{}, err
	}

	rewrite, err := client.GetRewrite(req.State.ProfileID, req.ID)
	if err != nil {
		return infer.ReadResponse[NextDNSRewriteArgs, NextDNSRewriteState]{}, err
	}

	return infer.ReadResponse[NextDNSRewriteArgs, NextDNSRewriteState]{
		ID: req.ID,
		Inputs: NextDNSRewriteArgs{
			ProfileID: req.Inputs.ProfileID,
			Name:      rewrite.Name,
			Content:   rewrite.Content,
		},
		State: NextDNSRewriteState{
			RewriteID: rewrite.ID,
			NextDNSRewriteArgs: NextDNSRewriteArgs{
				ProfileID: req.Inputs.ProfileID,
				Name:      rewrite.Name,
				Content:   rewrite.Content,
			},
		},
	}, nil
}

func (*NextDNSRewrite) WireDependencies(f infer.FieldSelector, args *NextDNSRewriteArgs, state *NextDNSRewriteState) {
	f.OutputField(&state.Content).DependsOn(f.InputField(&args.Content))
	f.OutputField(&state.Name).DependsOn(f.InputField(&args.Name))
	f.OutputField(&state.ProfileID).DependsOn(f.InputField(&args.ProfileID))
}
