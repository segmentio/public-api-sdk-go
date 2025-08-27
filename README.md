# Segment Public API Go SDK

:warning: This SDK is currently released as [Public Beta](https://segment.com/legal/first-access-beta-preview/). Its use in critical systems is discouraged, but [feedback is welcome](#contributing).

## Overview

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs. The full documentation is available at [https://docs.segmentapis.com](https://docs.segmentapis.com).

All endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.

See the next sections for more information on how to use the Segment Public API Go SDK.

Latest API and SDK version: 59.1.0

## Installation

In your project, you can install the package directly using `go get`:

```shell
go get github.com/segmentio/public-api-sdk-go
```

In your code, add the following in import:

```golang
import api "github.com/segmentio/public-api-sdk-go/api"
```

You are now ready to start making calls to Public API!

## Authorization

Requests are authorized by setting your API token in the context object:

```golang
token := // ...
ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
r, err := client.<Service>.<Operation>(ctx, args)
```

## Example

```golang
package main

import (
    "context"
    "fmt"
    "os"
    "github.com/segmentio/public-api-sdk-go/api"
)

func main() {
    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)

    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)

    resp, r, err := apiClient.SourcesAPI.ListSources(ctx).Pagination(api.PaginationInput{ Count: 10 }).Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SourcesApi.ListSources``: %v\n", err)
				responseErrors := api.UnwrapFullErrors(err)
				if responseErrors != nil {
					for _, responseError := range responseErrors.Errors {
						fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
					}
				}
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }

    fmt.Fprintf(os.Stdout, "Response from `SourcesApi.ListSources`: %v\n", resp.GetData())
}
```

## Contributing

The contents of this repository are automatically generated, so we can't take contributions from external developers. If you have any issues with this SDK, please raise an issue or reach out to friends@segment.com instead of opening a pull request. Pull requests will not be reviewed.
