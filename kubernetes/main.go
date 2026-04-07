package main

import (
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := kubernetes.NewProvider(ctx, "kubernetes", &kubernetes.ProviderArgs{})
		if err != nil {
			return err
		}
		return nil
	})
}
