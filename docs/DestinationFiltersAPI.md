# \DestinationFiltersAPI

All URIs are relative to *https://api.segmentapis.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFilterForDestination**](DestinationFiltersAPI.md#CreateFilterForDestination) | **Post** /destination/{destinationId}/filters | Create Filter for Destination
[**GetFilterInDestination**](DestinationFiltersAPI.md#GetFilterInDestination) | **Get** /destination/{destinationId}/filters/{filterId} | Get Filter in Destination
[**ListFiltersFromDestination**](DestinationFiltersAPI.md#ListFiltersFromDestination) | **Get** /destination/{destinationId}/filters | List Filters from Destination
[**PreviewDestinationFilter**](DestinationFiltersAPI.md#PreviewDestinationFilter) | **Post** /destination/filters/preview | Preview Destination Filter
[**RemoveFilterFromDestination**](DestinationFiltersAPI.md#RemoveFilterFromDestination) | **Delete** /destination/{destinationId}/filters/{filterId} | Remove Filter from Destination
[**UpdateFilterForDestination**](DestinationFiltersAPI.md#UpdateFilterForDestination) | **Patch** /destination/{destinationId}/filters/{filterId} | Update Filter for Destination



## Operation: CreateFilterForDestination

> CreateFilterForDestination200Response CreateFilterForDestination(ctx, destinationId).CreateFilterForDestinationV1Input(createFilterForDestinationV1Input).Execute()

Create Filter for Destination



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
    destinationId := "fP7qoQw2HTWt9WdMr718gn" // string | 
    createFilterForDestinationV1Input := *api.NewCreateFilterForDestinationV1Input("SourceId_example", "If_example", []api.DestinationFilterActionV1{*api.NewDestinationFilterActionV1("Type_example")}, "Title_example", false) // CreateFilterForDestinationV1Input | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.DestinationFiltersAPI.CreateFilterForDestination(ctx, destinationId).CreateFilterForDestinationV1Input(createFilterForDestinationV1Input).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DestinationFiltersAPI.CreateFilterForDestination``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `CreateFilterForDestination`: CreateFilterForDestination200Response
    fmt.Fprintf(os.Stdout, "Response from `DestinationFiltersAPI.CreateFilterForDestination`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**destinationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateFilterForDestinationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createFilterForDestinationV1Input** | [**CreateFilterForDestinationV1Input**](CreateFilterForDestinationV1Input.md) |  | 

### Return type

[**CreateFilterForDestination200Response**](CreateFilterForDestination200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/vnd.segment.v1+json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json
- **Accept**: application/vnd.segment.v1+json, application/json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: GetFilterInDestination

> GetFilterInDestination200Response GetFilterInDestination(ctx, destinationId, filterId).Execute()

Get Filter in Destination



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
    destinationId := "fP7qoQw2HTWt9WdMr718gn" // string | 
    filterId := "xx6AySGeFExzdv2Gw2EuhV" // string | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.DestinationFiltersAPI.GetFilterInDestination(ctx, destinationId, filterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DestinationFiltersAPI.GetFilterInDestination``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `GetFilterInDestination`: GetFilterInDestination200Response
    fmt.Fprintf(os.Stdout, "Response from `DestinationFiltersAPI.GetFilterInDestination`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**destinationId** | **string** |  | 
**filterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFilterInDestinationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetFilterInDestination200Response**](GetFilterInDestination200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1+json, application/json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: ListFiltersFromDestination

> ListFiltersFromDestination200Response ListFiltersFromDestination(ctx, destinationId).Pagination(pagination).Execute()

List Filters from Destination



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
    destinationId := "fP7qoQw2HTWt9WdMr718gn" // string | 
    pagination := *api.NewPaginationInput(10) // PaginationInput | Pagination options.  This parameter exists in v1.

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.DestinationFiltersAPI.ListFiltersFromDestination(ctx, destinationId).Pagination(pagination).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DestinationFiltersAPI.ListFiltersFromDestination``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `ListFiltersFromDestination`: ListFiltersFromDestination200Response
    fmt.Fprintf(os.Stdout, "Response from `DestinationFiltersAPI.ListFiltersFromDestination`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**destinationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListFiltersFromDestinationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pagination** | [**PaginationInput**](PaginationInput.md) | Pagination options.  This parameter exists in v1. | 

### Return type

[**ListFiltersFromDestination200Response**](ListFiltersFromDestination200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1+json, application/json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: PreviewDestinationFilter

> PreviewDestinationFilter200Response PreviewDestinationFilter(ctx).PreviewDestinationFilterV1Input(previewDestinationFilterV1Input).Execute()

Preview Destination Filter



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
    previewDestinationFilterV1Input := *api.NewPreviewDestinationFilterV1Input(*api.NewPreviewDestinationFilterV1("If_example", []api.DestinationFilterActionV1{*api.NewDestinationFilterActionV1("Type_example")}), map[string]interface{}{"key": interface{}(123)}) // PreviewDestinationFilterV1Input | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.DestinationFiltersAPI.PreviewDestinationFilter(ctx).PreviewDestinationFilterV1Input(previewDestinationFilterV1Input).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DestinationFiltersAPI.PreviewDestinationFilter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `PreviewDestinationFilter`: PreviewDestinationFilter200Response
    fmt.Fprintf(os.Stdout, "Response from `DestinationFiltersAPI.PreviewDestinationFilter`: %v\n", resp.GetData())
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPreviewDestinationFilterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **previewDestinationFilterV1Input** | [**PreviewDestinationFilterV1Input**](PreviewDestinationFilterV1Input.md) |  | 

### Return type

[**PreviewDestinationFilter200Response**](PreviewDestinationFilter200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/vnd.segment.v1+json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json
- **Accept**: application/vnd.segment.v1+json, application/json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: RemoveFilterFromDestination

> RemoveFilterFromDestination200Response RemoveFilterFromDestination(ctx, destinationId, filterId).Execute()

Remove Filter from Destination



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
    destinationId := "fP7qoQw2HTWt9WdMr718gn" // string | 
    filterId := "2DrXi3N7S85LobhzPphZz0uFzJ4" // string | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.DestinationFiltersAPI.RemoveFilterFromDestination(ctx, destinationId, filterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DestinationFiltersAPI.RemoveFilterFromDestination``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `RemoveFilterFromDestination`: RemoveFilterFromDestination200Response
    fmt.Fprintf(os.Stdout, "Response from `DestinationFiltersAPI.RemoveFilterFromDestination`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**destinationId** | **string** |  | 
**filterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveFilterFromDestinationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RemoveFilterFromDestination200Response**](RemoveFilterFromDestination200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1+json, application/json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: UpdateFilterForDestination

> UpdateFilterForDestination200Response UpdateFilterForDestination(ctx, destinationId, filterId).UpdateFilterForDestinationV1Input(updateFilterForDestinationV1Input).Execute()

Update Filter for Destination



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
    destinationId := "fP7qoQw2HTWt9WdMr718gn" // string | 
    filterId := "xx6AySGeFExzdv2Gw2EuhV" // string | 
    updateFilterForDestinationV1Input := *api.NewUpdateFilterForDestinationV1Input() // UpdateFilterForDestinationV1Input | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.DestinationFiltersAPI.UpdateFilterForDestination(ctx, destinationId, filterId).UpdateFilterForDestinationV1Input(updateFilterForDestinationV1Input).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DestinationFiltersAPI.UpdateFilterForDestination``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `UpdateFilterForDestination`: UpdateFilterForDestination200Response
    fmt.Fprintf(os.Stdout, "Response from `DestinationFiltersAPI.UpdateFilterForDestination`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**destinationId** | **string** |  | 
**filterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFilterForDestinationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateFilterForDestinationV1Input** | [**UpdateFilterForDestinationV1Input**](UpdateFilterForDestinationV1Input.md) |  | 

### Return type

[**UpdateFilterForDestination200Response**](UpdateFilterForDestination200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/vnd.segment.v1+json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json
- **Accept**: application/vnd.segment.v1+json, application/json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

