package main

import (
	"github.com/pulumi/pulumi-cloudflare/sdk/v6/go/cloudflare"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := cloudflare.NewProvider(ctx, "cloudflare", &cloudflare.ProviderArgs{})
		if err != nil {
			return err
		}
		return nil
	})
}
