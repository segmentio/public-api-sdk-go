# \MonthlyTrackedUsersAPI

All URIs are relative to *https://api.segmentapis.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDailyPerSourceMTUUsage**](MonthlyTrackedUsersAPI.md#GetDailyPerSourceMTUUsage) | **Get** /usage/mtu/sources/daily | Get Daily Per Source MTU Usage
[**GetDailyWorkspaceMTUUsage**](MonthlyTrackedUsersAPI.md#GetDailyWorkspaceMTUUsage) | **Get** /usage/mtu/daily | Get Daily Workspace MTU Usage



## Operation: GetDailyPerSourceMTUUsage

> GetDailyPerSourceMTUUsage200Response GetDailyPerSourceMTUUsage(ctx).Period(period).Pagination(pagination).Execute()

Get Daily Per Source MTU Usage



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
    period := "2021-02-01" // string | The start of the usage month, in the ISO-8601 format.  This parameter exists in v1.
    pagination := *openapiclient.NewPaginationInput(float32(123)) // PaginationInput | Pagination input for per Source MTU counts.  This parameter exists in v1.

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    resp, r, err := apiClient.MonthlyTrackedUsersAPI.GetDailyPerSourceMTUUsage(context.Background()).Period(period).Pagination(pagination).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonthlyTrackedUsersAPI.GetDailyPerSourceMTUUsage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDailyPerSourceMTUUsage`: GetDailyPerSourceMTUUsage200Response
    fmt.Fprintf(os.Stdout, "Response from `MonthlyTrackedUsersAPI.GetDailyPerSourceMTUUsage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDailyPerSourceMTUUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **period** | **string** | The start of the usage month, in the ISO-8601 format.  This parameter exists in v1. | 
 **pagination** | [**PaginationInput**](PaginationInput.md) | Pagination input for per Source MTU counts.  This parameter exists in v1. | 

### Return type

[**GetDailyPerSourceMTUUsage200Response**](GetDailyPerSourceMTUUsage200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1+json, application/json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: GetDailyWorkspaceMTUUsage

> GetDailyWorkspaceMTUUsage200Response GetDailyWorkspaceMTUUsage(ctx).Period(period).Pagination(pagination).Execute()

Get Daily Workspace MTU Usage



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
    period := "2021-02-01" // string | The start of the usage month, in the ISO-8601 format.  This parameter exists in v1.
    pagination := *openapiclient.NewPaginationInput(float32(123)) // PaginationInput | Pagination input for Workspace MTU counts.  This parameter exists in v1.

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    resp, r, err := apiClient.MonthlyTrackedUsersAPI.GetDailyWorkspaceMTUUsage(context.Background()).Period(period).Pagination(pagination).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonthlyTrackedUsersAPI.GetDailyWorkspaceMTUUsage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDailyWorkspaceMTUUsage`: GetDailyWorkspaceMTUUsage200Response
    fmt.Fprintf(os.Stdout, "Response from `MonthlyTrackedUsersAPI.GetDailyWorkspaceMTUUsage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDailyWorkspaceMTUUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **period** | **string** | The start of the usage month, in the ISO-8601 format.  This parameter exists in v1. | 
 **pagination** | [**PaginationInput**](PaginationInput.md) | Pagination input for Workspace MTU counts.  This parameter exists in v1. | 

### Return type

[**GetDailyWorkspaceMTUUsage200Response**](GetDailyWorkspaceMTUUsage200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1+json, application/json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

