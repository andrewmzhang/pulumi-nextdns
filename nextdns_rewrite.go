package main

import (
	"context"
	"github.com/pulumi/pulumi-go-provider/infer"
)

// TODO: Call the master object NextDNSProfileRewrite
// NextDNSRewrite Resource Controlling struct
type NextDNSRewrite struct {
}

// NextDNSRewriteArgs Resource input struct
type NextDNSRewriteArgs struct {
	ProfileId string `pulumi:"profileId,optional"`
	Name      string `pulumi:"name"`
	Content   string `pulumi:"content"`
}

type NextDNSRewriteState struct {
	RewriteId string `json:"rewriteId"`
	NextDNSRewriteArgs
}

func (f *NextDNSRewrite) Annotate(a infer.Annotator) {
	a.Describe(&f, "A NextDNS Rewrite into a pulumi resource")
}

func (f *NextDNSRewriteArgs) Annotate(a infer.Annotator) {
	a.Describe(&f.ProfileId, "Profile Id to apply rewrite to. This overrides the default profile id.")
	a.Describe(&f.Name, "Domain name to apply rewrite to.")
	a.Describe(&f.Content, "IP Address or Domain to rewrite domain name to.")
}

func (f *NextDNSRewriteState) Annotate(a infer.Annotator) {
}

func (f *NextDNSRewrite) Create(ctx context.Context, req infer.CreateRequest[NextDNSRewriteArgs]) (resp infer.CreateResponse[NextDNSRewriteState], err error) {
	// Get provider config
	config := infer.GetConfig[Config](ctx)
	apiKey := config.ApiKey

	var id string
	if !req.DryRun {
		id, err = createRewrite(apiKey, req.Inputs.ProfileId, req.Inputs.Name, req.Inputs.Content)
		if err != nil {
			panic(err)
		}
	}

	return infer.CreateResponse[NextDNSRewriteState]{
		ID: id,
		Output: NextDNSRewriteState{
			RewriteId:          id,
			NextDNSRewriteArgs: req.Inputs,
		},
	}, nil
}

func (*NextDNSRewrite) Delete(ctx context.Context, req infer.DeleteRequest[NextDNSRewriteState]) (infer.DeleteResponse, error) {
	// Get provider config
	config := infer.GetConfig[Config](ctx)
	apiKey := config.ApiKey
	profileID := req.State.ProfileId

	deleteRewrite(apiKey, profileID, req.State.RewriteId)
	return infer.DeleteResponse{}, nil
}
