[![Go Doc](https://pkg.go.dev/badge/github.com/rizalgowandy/synapse-go?status.svg)](https://pkg.go.dev/github.com/rizalgowandy/synapse-go?tab=doc)
[![Release](https://img.shields.io/github/release/rizalgowandy/synapse-go.svg?style=flat-square)](https://github.com/rizalgowandy/synapse-go/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/rizalgowandy/synapse-go)](https://goreportcard.com/report/github.com/rizalgowandy/synapse-go)
[![Build Status](https://github.com/rizalgowandy/synapse-go/workflows/Go/badge.svg?branch=main)](https://github.com/rizalgowandy/synapse-go/actions?query=branch%3Amain)
[![Sourcegraph](https://sourcegraph.com/github.com/rizalgowandy/synapse-go/-/badge.svg)](https://sourcegraph.com/github.com/rizalgowandy/synapse-go?badge)

![logo](https://socialify.git.ci/rizalgowandy/synapse-go/image?description=1&descriptionEditable=Interact%20with%20Synapse%20API.&logo=https%3A%2F%2Fpbs.twimg.com%2Fprofile_images%2F1560077835607281665%2FGvk6wzyA_400x400.png&owner=1&pattern=Signal&theme=Light)

## Getting Started

Every request to the Synapse API must be authenticated with a client id and secret key. You can get this information after creating a Synapse account.

## Installation

```shell
go get -v github.com/rizalgowandy/synapse-go
```

## Quick Start

```go
package main

import (
	"context"
	"log"

	"github.com/rizalgowandy/synapse-go"
	"github.com/rizalgowandy/synapse-go/pkg/api"
	"github.com/rizalgowandy/synapse-go/pkg/entity"
	"github.com/rizalgowandy/synapse-go/pkg/enum"
)

func main() {
	client, err := synapse.NewClient(&api.Config{
		ID:               "REPLACE_WITH_YOUR_API_CLIENT_ID",
		Key:              "REPLACE_WITH_YOUR_API_CLIENT_SECRET",
		Timeout:          0,    // Default: 1 min.
		RetryCount:       0,    // Default: 0 = disable.
		RetryMaxWaitTime: 0,    // Default: 2 secs.
		Debug:            true, // Turn on development for easier debugging.
	})
	if err != nil {
		log.Fatal(err)
	}

	// Generate refresh token.
	resp, err := client.GenerateRefreshToken(
		context.Background(),
		&entity.GenerateRefreshTokenReq{
			UserID:        "REPLACE_WITH_USER_ID",
			UserIPAddress: "REPLACE_WITH_USER_IP_ADDRESS",
			Email:         "REPLACE_WITH_USER_EMAIL",
			Password:      "REPLACE_WITH_USER_PASSWORD",
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%+v", resp)
}
```

For more example check [here](main_integration_test.go).

[//]: # (## Test Double / Stub)

[//]: # ()

[//]: # (Sometimes it's make sense to make an API call without actually calling the API. In order to support that this library has a built-in stub that can be triggered. You can enable stub by injecting certain value to the context data. You can also enforce that certain API call will always return error with specific type and)

[//]: # (message.)

[//]: # ()

[//]: # (```go)

[//]: # (// TODO: replace me)

[//]: # (```)

[//]: # ()

[//]: # (For more example, check [here]&#40;&#41;.)

## Supported API

Version: 2023-03-08

- [OAuth](https://docs.synapsefi.com/api-references/oauth)
   - [Create OAuth Key](https://docs.synapsefi.com/api-references/oauth/oauth-via-refresh-token)
   - [Generate Refresh Token](https://docs.synapsefi.com/api-references/oauth/generate-refresh-token)
- [Users](https://docs.synapsefi.com/api-references/users)
  - [View All Users](https://docs.synapsefi.com/api-references/users/view-all-users-paginated)
  - [View User](https://docs.synapsefi.com/api-references/users/view-user)
  - [Create User](https://docs.synapsefi.com/api-references/users/create-user)
  - [Update User](https://docs.synapsefi.com/api-references/users/update-user)
  - [Generate OBO Doc](https://docs.synapsefi.com/api-references/users/generate-ubo-doc)
  - [Manage Duplicates](https://docs.synapsefi.com/api-references/users/manage-duplicates)
    - Get Duplicates
    - Swap Duplicate Users
  - [Allowed Document Types](https://docs.synapsefi.com/api-references/users/allowed-document-types)
  - [Allowed Entity Scopes](https://docs.synapsefi.com/api-references/users/allowed-entity-scopes)
  - [Allow Entity Types](https://docs.synapsefi.com/api-references/users/allowed-entity-types)
- [Nodes](https://docs.synapsefi.com/api-references/nodes)
  - [View All User Nodes](https://docs.synapsefi.com/api-references/nodes/view-all-user-nodes)
  - [View Node](https://docs.synapsefi.com/api-references/nodes/view-node)
  - [Create Node](https://docs.synapsefi.com/api-references/nodes/create-node)
  - [Update Node](https://docs.synapsefi.com/api-references/nodes/update-node)
  - [Generate ECash Barcode](https://docs.synapsefi.com/api-references/nodes/generate-ecash-barcode)
  - [Allowed Node Types](https://docs.synapsefi.com/api-references/nodes/allowed-node-types)
  - [View ATMs](https://docs.synapsefi.com/api-references/nodes/view-atms)
  - [View Crypto Quotes](https://docs.synapsefi.com/api-references/nodes/view-crypto-quotes)
- [Subnets](https://docs.synapsefi.com/api-references/subnets)
  - [](https://docs.synapsefi.com/api-references/subnets/view-all-node-subnets)
  - [](https://docs.synapsefi.com/api-references/subnets/view-subnet)
  - [](https://docs.synapsefi.com/api-references/subnets/create-subnet)
  - [](https://docs.synapsefi.com/api-references/subnets/update-subnet)
  - [](https://docs.synapsefi.com/api-references/subnets/push-to-wallet)
