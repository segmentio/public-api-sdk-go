# \CustomerInsightsAPI

All URIs are relative to *https://api.segmentapis.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDownload**](CustomerInsightsAPI.md#CreateDownload) | **Post** /customer-insights/download | Create Download



## Operation: CreateDownload

> CreateDownload200Response CreateDownload(ctx).CreateDownloadAlphaInput(createDownloadAlphaInput).Execute()

Create Download



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    api "github.com/segmentio/public-api-sdk-go"
)

func main() {
    createDownloadAlphaInput := *api.NewCreateDownloadAlphaInput("CollectionId_example", "WorkspaceId_example", "Hour_example") // CreateDownloadAlphaInput | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.CustomerInsightsAPI.CreateDownload(ctx).CreateDownloadAlphaInput(createDownloadAlphaInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerInsightsAPI.CreateDownload``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `CreateDownload`: CreateDownload200Response
    fmt.Fprintf(os.Stdout, "Response from `CustomerInsightsAPI.CreateDownload`: %v\n", resp.GetData())
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDownloadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createDownloadAlphaInput** | [**CreateDownloadAlphaInput**](CreateDownloadAlphaInput.md) |  | 

### Return type

[**CreateDownload200Response**](CreateDownload200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/vnd.segment.v1alpha+json
- **Accept**: application/vnd.segment.v1alpha+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

