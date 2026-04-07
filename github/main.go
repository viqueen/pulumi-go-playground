package main

import (
	"github.com/pulumi/pulumi-github/sdk/v6/go/github"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := github.NewProvider(ctx, "github", &github.ProviderArgs{})
		if err != nil {
			return err
		}
		return nil
	})
}
