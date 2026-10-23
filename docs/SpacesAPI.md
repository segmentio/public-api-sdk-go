# \SpacesAPI

All URIs are relative to *https://api.segmentapis.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSpace**](SpacesAPI.md#GetSpace) | **Get** /spaces/{spaceId} | Get Space
[**ListSpaces**](SpacesAPI.md#ListSpaces) | **Get** /spaces | List Spaces



## Operation: GetSpace

> GetSpace200Response GetSpace(ctx, spaceId).Execute()

Get Space



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
    spaceId := "9aQ1Lj62S4bomZKLF4DPqW" // string | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.SpacesAPI.GetSpace(ctx, spaceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpacesAPI.GetSpace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `GetSpace`: GetSpace200Response
    fmt.Fprintf(os.Stdout, "Response from `SpacesAPI.GetSpace`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSpaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetSpace200Response**](GetSpace200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1alpha+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: ListSpaces

> ListSpaces200Response ListSpaces(ctx).Pagination(pagination).Execute()

List Spaces



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
    pagination := *api.NewPaginationInput(10) // PaginationInput | Pagination params  This parameter exists in alpha. (optional)

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.SpacesAPI.ListSpaces(ctx).Pagination(pagination).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpacesAPI.ListSpaces``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `ListSpaces`: ListSpaces200Response
    fmt.Fprintf(os.Stdout, "Response from `SpacesAPI.ListSpaces`: %v\n", resp.GetData())
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSpacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pagination** | [**PaginationInput**](PaginationInput.md) | Pagination params  This parameter exists in alpha. | 

### Return type

[**ListSpaces200Response**](ListSpaces200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1alpha+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

