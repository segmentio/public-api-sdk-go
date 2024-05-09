# \ComputedTraitsAPI

All URIs are relative to *https://api.segmentapis.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetComputedTrait**](ComputedTraitsAPI.md#GetComputedTrait) | **Get** /spaces/{spaceId}/computed-traits/{id} | Get Computed Trait
[**ListComputedTraits**](ComputedTraitsAPI.md#ListComputedTraits) | **Get** /spaces/{spaceId}/computed-traits | List Computed Traits
[**RemoveComputedTraitFromSpace**](ComputedTraitsAPI.md#RemoveComputedTraitFromSpace) | **Delete** /spaces/{spaceId}/computed-traits/{id} | Remove Computed Trait from Space
[**UpdateComputedTraitForSpace**](ComputedTraitsAPI.md#UpdateComputedTraitForSpace) | **Patch** /spaces/{spaceId}/computed-traits/{id} | Update Computed Trait for Space



## Operation: GetComputedTrait

> GetComputedTrait200Response GetComputedTrait(ctx, spaceId, id).Execute()

Get Computed Trait



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
    spaceId := "spaceId" // string | 
    id := "id" // string | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.ComputedTraitsAPI.GetComputedTrait(ctx, spaceId, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComputedTraitsAPI.GetComputedTrait``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `GetComputedTrait`: GetComputedTrait200Response
    fmt.Fprintf(os.Stdout, "Response from `ComputedTraitsAPI.GetComputedTrait`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetComputedTraitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetComputedTrait200Response**](GetComputedTrait200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1alpha+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: ListComputedTraits

> ListComputedTraits200Response ListComputedTraits(ctx, spaceId).Pagination(pagination).Execute()

List Computed Traits



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
    spaceId := "spaceId" // string | 
    pagination := *api.NewPaginationInput(10) // PaginationInput | Information about the pagination of this response.  This parameter exists in alpha. (optional)

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.ComputedTraitsAPI.ListComputedTraits(ctx, spaceId).Pagination(pagination).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComputedTraitsAPI.ListComputedTraits``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `ListComputedTraits`: ListComputedTraits200Response
    fmt.Fprintf(os.Stdout, "Response from `ComputedTraitsAPI.ListComputedTraits`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListComputedTraitsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pagination** | [**PaginationInput**](PaginationInput.md) | Information about the pagination of this response.  This parameter exists in alpha. | 

### Return type

[**ListComputedTraits200Response**](ListComputedTraits200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1alpha+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: RemoveComputedTraitFromSpace

> RemoveComputedTraitFromSpace200Response RemoveComputedTraitFromSpace(ctx, spaceId, id).Execute()

Remove Computed Trait from Space



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
    spaceId := "spaceId" // string | 
    id := "id" // string | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.ComputedTraitsAPI.RemoveComputedTraitFromSpace(ctx, spaceId, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComputedTraitsAPI.RemoveComputedTraitFromSpace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `RemoveComputedTraitFromSpace`: RemoveComputedTraitFromSpace200Response
    fmt.Fprintf(os.Stdout, "Response from `ComputedTraitsAPI.RemoveComputedTraitFromSpace`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveComputedTraitFromSpaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RemoveComputedTraitFromSpace200Response**](RemoveComputedTraitFromSpace200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1alpha+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: UpdateComputedTraitForSpace

> UpdateComputedTraitForSpace200Response UpdateComputedTraitForSpace(ctx, spaceId, id).UpdateComputedTraitForSpaceAlphaInput(updateComputedTraitForSpaceAlphaInput).Execute()

Update Computed Trait for Space



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
    spaceId := "spaceId" // string | 
    id := "id" // string | 
    updateComputedTraitForSpaceAlphaInput := *api.NewUpdateComputedTraitForSpaceAlphaInput() // UpdateComputedTraitForSpaceAlphaInput | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.ComputedTraitsAPI.UpdateComputedTraitForSpace(ctx, spaceId, id).UpdateComputedTraitForSpaceAlphaInput(updateComputedTraitForSpaceAlphaInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComputedTraitsAPI.UpdateComputedTraitForSpace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `UpdateComputedTraitForSpace`: UpdateComputedTraitForSpace200Response
    fmt.Fprintf(os.Stdout, "Response from `ComputedTraitsAPI.UpdateComputedTraitForSpace`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateComputedTraitForSpaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateComputedTraitForSpaceAlphaInput** | [**UpdateComputedTraitForSpaceAlphaInput**](UpdateComputedTraitForSpaceAlphaInput.md) |  | 

### Return type

[**UpdateComputedTraitForSpace200Response**](UpdateComputedTraitForSpace200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/vnd.segment.v1alpha+json
- **Accept**: application/vnd.segment.v1alpha+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

