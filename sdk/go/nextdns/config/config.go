// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package config

import (
	"github.com/andrewmzhang/pulumi-nextdns/sdk/go/nextdns/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

var _ = internal.GetEnvOrDefault

func GetApiKey(ctx *pulumi.Context) string {
	return config.Get(ctx, "nextdns:apiKey")
}
