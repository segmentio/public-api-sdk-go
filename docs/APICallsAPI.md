# \APICallsAPI

All URIs are relative to *https://api.segmentapis.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDailyPerSourceAPICallsUsage**](APICallsAPI.md#GetDailyPerSourceAPICallsUsage) | **Get** /usage/api-calls/sources/daily | Get Daily Per Source API Calls Usage
[**GetDailyWorkspaceAPICallsUsage**](APICallsAPI.md#GetDailyWorkspaceAPICallsUsage) | **Get** /usage/api-calls/daily | Get Daily Workspace API Calls Usage



## Operation: GetDailyPerSourceAPICallsUsage

> GetDailyPerSourceAPICallsUsage200Response GetDailyPerSourceAPICallsUsage(ctx).Period(period).Pagination(pagination).Execute()

Get Daily Per Source API Calls Usage



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
    period := "2021-02-01" // string | The start of the usage month in the ISO-8601 format.  This parameter exists in v1.
    pagination := *api.NewPaginationInput(10) // PaginationInput | Pagination input for per Source API calls counts.  This parameter exists in v1. (optional)

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.APICallsAPI.GetDailyPerSourceAPICallsUsage(ctx).Period(period).Pagination(pagination).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APICallsAPI.GetDailyPerSourceAPICallsUsage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `GetDailyPerSourceAPICallsUsage`: GetDailyPerSourceAPICallsUsage200Response
    fmt.Fprintf(os.Stdout, "Response from `APICallsAPI.GetDailyPerSourceAPICallsUsage`: %v\n", resp.GetData())
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDailyPerSourceAPICallsUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **period** | **string** | The start of the usage month in the ISO-8601 format.  This parameter exists in v1. | 
 **pagination** | [**PaginationInput**](PaginationInput.md) | Pagination input for per Source API calls counts.  This parameter exists in v1. | 

### Return type

[**GetDailyPerSourceAPICallsUsage200Response**](GetDailyPerSourceAPICallsUsage200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1+json, application/json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: GetDailyWorkspaceAPICallsUsage

> GetDailyWorkspaceAPICallsUsage200Response GetDailyWorkspaceAPICallsUsage(ctx).Period(period).Pagination(pagination).Execute()

Get Daily Workspace API Calls Usage



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
    period := "2021-02-01" // string | The start of the usage month in the ISO-8601 format.  This parameter exists in v1.
    pagination := *api.NewPaginationInput(10) // PaginationInput | Pagination input for Workspace API call counts.  This parameter exists in v1. (optional)

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.APICallsAPI.GetDailyWorkspaceAPICallsUsage(ctx).Period(period).Pagination(pagination).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APICallsAPI.GetDailyWorkspaceAPICallsUsage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `GetDailyWorkspaceAPICallsUsage`: GetDailyWorkspaceAPICallsUsage200Response
    fmt.Fprintf(os.Stdout, "Response from `APICallsAPI.GetDailyWorkspaceAPICallsUsage`: %v\n", resp.GetData())
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDailyWorkspaceAPICallsUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **period** | **string** | The start of the usage month in the ISO-8601 format.  This parameter exists in v1. | 
 **pagination** | [**PaginationInput**](PaginationInput.md) | Pagination input for Workspace API call counts.  This parameter exists in v1. | 

### Return type

[**GetDailyWorkspaceAPICallsUsage200Response**](GetDailyWorkspaceAPICallsUsage200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1+json, application/json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

