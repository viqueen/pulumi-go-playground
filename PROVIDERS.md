# Pulumi Provider Go SDKs

Tracking which `github.com/pulumi/*` providers ship a Go SDK, plus the latest released version and a ready-to-paste `go get` line.

- Import path convention: `github.com/pulumi/<repo>/sdk[/v<major>]/go/<pkg>`
- Verified against the default branch and the latest GitHub release tag.
- 96 of 102 candidates ship a Go SDK; 6 do not (see bottom).

## Major clouds

| Repo | Latest | go get |
| --- | --- | --- |
| pulumi-aws | v7.24.0 | `go get github.com/pulumi/pulumi-aws/sdk/v7@v7.24.0` |
| pulumi-aws-native | v1.60.0 | `go get github.com/pulumi/pulumi-aws-native/sdk@v1.60.0` |
| pulumi-awsx | v3.4.0 | `go get github.com/pulumi/pulumi-awsx/sdk/v3@v3.4.0` |
| pulumi-aws-apigateway | v3.0.0 | `go get github.com/pulumi/pulumi-aws-apigateway/sdk/v3@v3.0.0` |
| pulumi-azure | v6.34.0 | `go get github.com/pulumi/pulumi-azure/sdk/v6@v6.34.0` |
| pulumi-azure-native-sdk | per-service | `go get github.com/pulumi/pulumi-azure-native-sdk/<service>/v3@latest` (one Go module per Azure service, e.g. `storage/v3`, `compute/v3`) |
| pulumi-gcp | v9.18.0 | `go get github.com/pulumi/pulumi-gcp/sdk/v9@v9.18.0` |
| pulumi-google-native | v0.32.0 | `go get github.com/pulumi/pulumi-google-native/sdk@v0.32.0` |
| pulumi-alicloud | v3.98.0 | `go get github.com/pulumi/pulumi-alicloud/sdk/v3@v3.98.0` |
| pulumi-oci | v4.5.1 | `go get github.com/pulumi/pulumi-oci/sdk/v4@v4.5.1` |
| pulumi-yandex | v0.13.0 | `go get github.com/pulumi/pulumi-yandex/sdk@v0.13.0` |
| pulumi-ucloud | (no tags) | use commit pseudo-version |

## Kubernetes / containers

| Repo | Latest | go get |
| --- | --- | --- |
| pulumi-kubernetes | v4.28.0 | `go get github.com/pulumi/pulumi-kubernetes/sdk/v4@v4.28.0` |
| pulumi-kubernetes-cert-manager | v0.2.0 | `go get github.com/pulumi/pulumi-kubernetes-cert-manager/sdk@v0.2.0` |
| pulumi-kubernetes-coredns | v0.1.0 | `go get github.com/pulumi/pulumi-kubernetes-coredns/sdk@v0.1.0` |
| pulumi-kubernetes-ingress-nginx | v0.1.3 | `go get github.com/pulumi/pulumi-kubernetes-ingress-nginx/sdk@v0.1.3` |
| pulumi-eks | v4.2.0 | `go get github.com/pulumi/pulumi-eks/sdk/v4@v4.2.0` |
| pulumi-docker | v4.11.2 | `go get github.com/pulumi/pulumi-docker/sdk/v4@v4.11.2` |
| pulumi-docker-build | v0.0.15 | `go get github.com/pulumi/pulumi-docker-build/sdk@v0.0.15` |

## Other infra / cloud

| Repo | Latest | go get |
| --- | --- | --- |
| pulumi-digitalocean | v4.63.0 | `go get github.com/pulumi/pulumi-digitalocean/sdk/v4@v4.63.0` |
| pulumi-linode | v5.10.0 | `go get github.com/pulumi/pulumi-linode/sdk/v5@v5.10.0` |
| pulumi-hcloud | v1.32.1 | `go get github.com/pulumi/pulumi-hcloud/sdk@v1.32.1` |
| pulumi-civo | v2.4.8 | `go get github.com/pulumi/pulumi-civo/sdk/v2@v2.4.8` |
| pulumi-equinix-metal | v3.2.1 | `go get github.com/pulumi/pulumi-equinix-metal/sdk/v3@v3.2.1` |
| pulumi-packet | v3.2.2 | `go get github.com/pulumi/pulumi-packet/sdk/v3@v3.2.2` |
| pulumi-libvirt | v0.5.4 | `go get github.com/pulumi/pulumi-libvirt/sdk@v0.5.4` |
| pulumi-openstack | v5.4.1 | `go get github.com/pulumi/pulumi-openstack/sdk/v5@v5.4.1` |
| pulumi-vsphere | v4.16.5 | `go get github.com/pulumi/pulumi-vsphere/sdk/v4@v4.16.5` |
| pulumi-hyperv | (no tags) | use commit pseudo-version |
| pulumi-ec | v0.10.10 | `go get github.com/pulumi/pulumi-ec/sdk@v0.10.10` |

## DNS / networking / CDN / SD-WAN

| Repo | Latest | go get |
| --- | --- | --- |
| pulumi-cloudflare | v6.14.0 | `go get github.com/pulumi/pulumi-cloudflare/sdk/v6@v6.14.0` |
| pulumi-ns1 | v3.8.2 | `go get github.com/pulumi/pulumi-ns1/sdk/v3@v3.8.2` |
| pulumi-dnsimple | v5.0.2 | `go get github.com/pulumi/pulumi-dnsimple/sdk/v5@v5.0.2` |
| pulumi-fastly | v11.5.0 | `go get github.com/pulumi/pulumi-fastly/sdk/v11@v11.5.0` |
| pulumi-f5bigip | v3.20.1 | `go get github.com/pulumi/pulumi-f5bigip/sdk/v3@v3.20.1` |
| pulumi-tailscale | v0.27.0 | `go get github.com/pulumi/pulumi-tailscale/sdk@v0.27.0` |
| pulumi-cloudngfwaws | v1.0.2 | `go get github.com/pulumi/pulumi-cloudngfwaws/sdk@v1.0.2` |
| pulumi-junipermist | v0.8.1 | `go get github.com/pulumi/pulumi-junipermist/sdk@v0.8.1` |
| pulumi-meraki | v0.4.6 | `go get github.com/pulumi/pulumi-meraki/sdk@v0.4.6` |
| pulumi-sdwan | v0.8.1 | `go get github.com/pulumi/pulumi-sdwan/sdk@v0.8.1` |
| pulumi-ise | v0.3.0 | `go get github.com/pulumi/pulumi-ise/sdk@v0.3.0` |

## SaaS / DevOps / data platforms

| Repo | Latest | go get |
| --- | --- | --- |
| pulumi-github | v6.12.1 | `go get github.com/pulumi/pulumi-github/sdk/v6@v6.12.1` |
| pulumi-gitlab | v9.10.0 | `go get github.com/pulumi/pulumi-gitlab/sdk/v9@v9.10.0` |
| pulumi-azuredevops | v3.14.0 | `go get github.com/pulumi/pulumi-azuredevops/sdk/v3@v3.14.0` |
| pulumi-harness | v0.11.8 | `go get github.com/pulumi/pulumi-harness/sdk@v0.11.8` |
| pulumi-dbtcloud | v1.7.0 | `go get github.com/pulumi/pulumi-dbtcloud/sdk@v1.7.0` |
| pulumi-databricks | v1.90.0 | `go get github.com/pulumi/pulumi-databricks/sdk@v1.90.0` |
| pulumi-snowflake | v2.13.1 | `go get github.com/pulumi/pulumi-snowflake/sdk/v2@v2.13.1` |
| pulumi-confluentcloud | v2.63.0 | `go get github.com/pulumi/pulumi-confluentcloud/sdk/v2@v2.63.0` |
| pulumi-confluent | v0.2.2 | `go get github.com/pulumi/pulumi-confluent/sdk@v0.2.2` |
| pulumi-mongodbatlas | v4.6.0 | `go get github.com/pulumi/pulumi-mongodbatlas/sdk/v4@v4.6.0` |
| pulumi-aiven | v6.51.0 | `go get github.com/pulumi/pulumi-aiven/sdk/v6@v6.51.0` |
| pulumi-akamai | v11.1.0 | `go get github.com/pulumi/pulumi-akamai/sdk/v11@v11.1.0` |
| pulumi-artifactory | v8.10.3 | `go get github.com/pulumi/pulumi-artifactory/sdk/v8@v8.10.3` |
| pulumi-rancher2 | v11.0.1 | `go get github.com/pulumi/pulumi-rancher2/sdk/v11@v11.0.1` |
| pulumi-rke | v3.6.0 | `go get github.com/pulumi/pulumi-rke/sdk/v3@v3.6.0` |
| pulumi-pulumiservice | v0.36.0 | `go get github.com/pulumi/pulumi-pulumiservice/sdk@v0.36.0` |
| pulumi-spotinst | v3.129.0 | `go get github.com/pulumi/pulumi-spotinst/sdk/v3@v3.129.0` |

## Auth / identity / secrets / certs

| Repo | Latest | go get |
| --- | --- | --- |
| pulumi-auth0 | v3.39.0 | `go get github.com/pulumi/pulumi-auth0/sdk/v3@v3.39.0` |
| pulumi-okta | v6.4.0 | `go get github.com/pulumi/pulumi-okta/sdk/v6@v6.4.0` |
| pulumi-onelogin | v0.6.9 | `go get github.com/pulumi/pulumi-onelogin/sdk@v0.6.9` |
| pulumi-keycloak | v6.10.0 | `go get github.com/pulumi/pulumi-keycloak/sdk/v6@v6.10.0` |
| pulumi-azuread | v6.9.0 | `go get github.com/pulumi/pulumi-azuread/sdk/v6@v6.9.0` |
| pulumi-vault | v7.8.0 | `go get github.com/pulumi/pulumi-vault/sdk/v7@v7.8.0` |
| pulumi-tls | v5.3.1 | `go get github.com/pulumi/pulumi-tls/sdk/v5@v5.3.1` |
| pulumi-venafi | v1.12.3 | `go get github.com/pulumi/pulumi-venafi/sdk@v1.12.3` |
| pulumi-scm | v1.0.5 | `go get github.com/pulumi/pulumi-scm/sdk@v1.0.5` |

## Observability / incident

| Repo | Latest | go get |
| --- | --- | --- |
| pulumi-datadog | v5.1.0 | `go get github.com/pulumi/pulumi-datadog/sdk/v5@v5.1.0` |
| pulumi-newrelic | v5.64.1 | `go get github.com/pulumi/pulumi-newrelic/sdk/v5@v5.64.1` |
| pulumi-signalfx | v7.24.0 | `go get github.com/pulumi/pulumi-signalfx/sdk/v7@v7.24.0` |
| pulumi-splunk | v1.3.0 | `go get github.com/pulumi/pulumi-splunk/sdk@v1.3.0` |
| pulumi-sumologic | v1.0.11 | `go get github.com/pulumi/pulumi-sumologic/sdk@v1.0.11` |
| pulumi-wavefront | v3.1.12 | `go get github.com/pulumi/pulumi-wavefront/sdk/v3@v3.1.12` |
| pulumi-pagerduty | v4.31.1 | `go get github.com/pulumi/pulumi-pagerduty/sdk/v4@v4.31.1` |
| pulumi-opsgenie | v1.3.20 | `go get github.com/pulumi/pulumi-opsgenie/sdk@v1.3.20` |

## Data stores / messaging / orchestration

| Repo | Latest | go get |
| --- | --- | --- |
| pulumi-mysql | v3.2.14 | `go get github.com/pulumi/pulumi-mysql/sdk/v3@v3.2.14` |
| pulumi-postgresql | v3.16.2 | `go get github.com/pulumi/pulumi-postgresql/sdk/v3@v3.16.2` |
| pulumi-kafka | v3.12.3 | `go get github.com/pulumi/pulumi-kafka/sdk/v3@v3.12.3` |
| pulumi-rabbitmq | v3.4.2 | `go get github.com/pulumi/pulumi-rabbitmq/sdk/v3@v3.4.2` |
| pulumi-cloudamqp | v3.27.1 | `go get github.com/pulumi/pulumi-cloudamqp/sdk/v3@v3.27.1` |
| pulumi-minio | v0.16.8 | `go get github.com/pulumi/pulumi-minio/sdk@v0.16.8` |
| pulumi-nomad | v2.5.5 | `go get github.com/pulumi/pulumi-nomad/sdk/v2@v2.5.5` |
| pulumi-consul | v3.14.1 | `go get github.com/pulumi/pulumi-consul/sdk/v3@v3.14.1` |

## Comms / misc SaaS

| Repo | Latest | go get |
| --- | --- | --- |
| pulumi-slack | v0.4.16 | `go get github.com/pulumi/pulumi-slack/sdk@v0.4.16` |
| pulumi-mailgun | v3.7.1 | `go get github.com/pulumi/pulumi-mailgun/sdk/v3@v3.7.1` |

## Utility providers

| Repo | Latest | go get |
| --- | --- | --- |
| pulumi-random | v4.19.2 | `go get github.com/pulumi/pulumi-random/sdk/v4@v4.19.2` |
| pulumi-null | v0.0.15 | `go get github.com/pulumi/pulumi-null/sdk@v0.0.15` |
| pulumi-command | v1.2.1 | `go get github.com/pulumi/pulumi-command/sdk@v1.2.1` |
| pulumi-local | v0.1.6 | `go get github.com/pulumi/pulumi-local/sdk@v0.1.6` |
| pulumi-http | v0.1.4 | `go get github.com/pulumi/pulumi-http/sdk@v0.1.4` |
| pulumi-archive | v0.3.7 | `go get github.com/pulumi/pulumi-archive/sdk@v0.3.7` |
| pulumi-cloudinit | v1.4.16 | `go get github.com/pulumi/pulumi-cloudinit/sdk@v1.4.16` |
| pulumi-std | v2.3.2 | `go get github.com/pulumi/pulumi-std/sdk/v2@v2.3.2` |
| pulumi-synced-folder | v0.12.4 | `go get github.com/pulumi/pulumi-synced-folder/sdk@v0.12.4` |
| pulumi-tls-self-signed-cert | v0.1.3 | `go get github.com/pulumi/pulumi-tls-self-signed-cert/sdk@v0.1.3` |
| pulumi-external | v0.0.18 | `go get github.com/pulumi/pulumi-external/sdk@v0.0.18` |
| pulumi-terraform | v6.0.1 | `go get github.com/pulumi/pulumi-terraform/sdk/v6@v6.0.1` |

## No Go SDK

| Repo | Reason |
| --- | --- |
| pulumi-azure-native | Go SDK split out into [`pulumi-azure-native-sdk`](https://github.com/pulumi/pulumi-azure-native-sdk) (one Go module per Azure service) |
| pulumi-cdk | TypeScript-only library |
| pulumi-cloud | Legacy multi-language framework, no `sdk/go` |
| pulumi-terraform-module | Generates per-module SDKs at provider runtime |
| pulumi-hubspot | No SDK published |
| pulumi-chocolatey | No SDK published |

## Notes

- The `/v<major>` segment in the import path is required by Go modules whenever the module's major version is `>= 2`. For `v0.x` and `v1.x` modules, omit it.
- Each provider's Go package(s) live under `sdk[/v<major>]/go/<pkg>`. For most providers `<pkg>` matches the provider name (e.g. `aws`, `gcp`); some providers expose multiple subpackages (e.g. `kubernetes/core/v1`).
- `pulumi-azure-native-sdk` does not have a single version — each Azure service is its own Go module versioned independently (tags look like `storage/v3.16.0`). Run `go get github.com/pulumi/pulumi-azure-native-sdk/<service>/v3@latest` per service you need.
- `pulumi-hyperv` and `pulumi-ucloud` ship `sdk/go` in source but have no tagged Go releases — you'd need a commit pseudo-version (`go get github.com/pulumi/pulumi-hyperv/sdk@<commit>`).
