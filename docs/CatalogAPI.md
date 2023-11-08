# \CatalogAPI

All URIs are relative to *https://api.segmentapis.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDestinationMetadata**](CatalogAPI.md#GetDestinationMetadata) | **Get** /catalog/destinations/{destinationMetadataId} | Get Destination Metadata
[**GetDestinationsCatalog**](CatalogAPI.md#GetDestinationsCatalog) | **Get** /catalog/destinations | Get Destinations Catalog
[**GetSourceMetadata**](CatalogAPI.md#GetSourceMetadata) | **Get** /catalog/sources/{sourceMetadataId} | Get Source Metadata
[**GetSourcesCatalog**](CatalogAPI.md#GetSourcesCatalog) | **Get** /catalog/sources | Get Sources Catalog
[**GetWarehouseMetadata**](CatalogAPI.md#GetWarehouseMetadata) | **Get** /catalog/warehouses/{warehouseMetadataId} | Get Warehouse Metadata
[**GetWarehousesCatalog**](CatalogAPI.md#GetWarehousesCatalog) | **Get** /catalog/warehouses | Get Warehouses Catalog



## Operation: GetDestinationMetadata

> GetDestinationMetadata200Response GetDestinationMetadata(ctx, destinationMetadataId).Execute()

Get Destination Metadata



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
    destinationMetadataId := "54521fd525e721e32a72ee91" // string | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "INSERT BEARER TOKEN HERE"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.CatalogAPI.GetDestinationMetadata(ctx, destinationMetadataId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CatalogAPI.GetDestinationMetadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `GetDestinationMetadata`: GetDestinationMetadata200Response
    fmt.Fprintf(os.Stdout, "Response from `CatalogAPI.GetDestinationMetadata`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**destinationMetadataId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDestinationMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetDestinationMetadata200Response**](GetDestinationMetadata200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1+json, application/json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: GetDestinationsCatalog

> GetDestinationsCatalog200Response GetDestinationsCatalog(ctx).Pagination(pagination).Execute()

Get Destinations Catalog



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
    pagination := *api.NewPaginationInput(float32(123)) // PaginationInput | Required pagination parameters used to filter the Destinations catalog.  This parameter exists in v1.

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "INSERT BEARER TOKEN HERE"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.CatalogAPI.GetDestinationsCatalog(ctx).Pagination(pagination).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CatalogAPI.GetDestinationsCatalog``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `GetDestinationsCatalog`: GetDestinationsCatalog200Response
    fmt.Fprintf(os.Stdout, "Response from `CatalogAPI.GetDestinationsCatalog`: %v\n", resp.GetData())
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDestinationsCatalogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pagination** | [**PaginationInput**](PaginationInput.md) | Required pagination parameters used to filter the Destinations catalog.  This parameter exists in v1. | 

### Return type

[**GetDestinationsCatalog200Response**](GetDestinationsCatalog200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1+json, application/json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: GetSourceMetadata

> GetSourceMetadata200Response GetSourceMetadata(ctx, sourceMetadataId).Execute()

Get Source Metadata



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
    sourceMetadataId := "1bow82lmk" // string | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "INSERT BEARER TOKEN HERE"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.CatalogAPI.GetSourceMetadata(ctx, sourceMetadataId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CatalogAPI.GetSourceMetadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `GetSourceMetadata`: GetSourceMetadata200Response
    fmt.Fprintf(os.Stdout, "Response from `CatalogAPI.GetSourceMetadata`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceMetadataId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSourceMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetSourceMetadata200Response**](GetSourceMetadata200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1+json, application/json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: GetSourcesCatalog

> GetSourcesCatalog200Response GetSourcesCatalog(ctx).Pagination(pagination).Execute()

Get Sources Catalog



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
    pagination := *api.NewPaginationInput(float32(123)) // PaginationInput | Defines the pagination parameters.  This parameter exists in v1.

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "INSERT BEARER TOKEN HERE"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.CatalogAPI.GetSourcesCatalog(ctx).Pagination(pagination).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CatalogAPI.GetSourcesCatalog``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `GetSourcesCatalog`: GetSourcesCatalog200Response
    fmt.Fprintf(os.Stdout, "Response from `CatalogAPI.GetSourcesCatalog`: %v\n", resp.GetData())
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSourcesCatalogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pagination** | [**PaginationInput**](PaginationInput.md) | Defines the pagination parameters.  This parameter exists in v1. | 

### Return type

[**GetSourcesCatalog200Response**](GetSourcesCatalog200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1+json, application/json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: GetWarehouseMetadata

> GetWarehouseMetadata200Response GetWarehouseMetadata(ctx, warehouseMetadataId).Execute()

Get Warehouse Metadata



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
    warehouseMetadataId := "55d3d3aea3c" // string | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "INSERT BEARER TOKEN HERE"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.CatalogAPI.GetWarehouseMetadata(ctx, warehouseMetadataId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CatalogAPI.GetWarehouseMetadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `GetWarehouseMetadata`: GetWarehouseMetadata200Response
    fmt.Fprintf(os.Stdout, "Response from `CatalogAPI.GetWarehouseMetadata`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**warehouseMetadataId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWarehouseMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetWarehouseMetadata200Response**](GetWarehouseMetadata200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1+json, application/json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: GetWarehousesCatalog

> GetWarehousesCatalog200Response GetWarehousesCatalog(ctx).Pagination(pagination).Execute()

Get Warehouses Catalog



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
    pagination := *api.NewPaginationInput(float32(123)) // PaginationInput | Required pagination params used to filter the Warehouses catalog.  This parameter exists in v1.

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "INSERT BEARER TOKEN HERE"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.CatalogAPI.GetWarehousesCatalog(ctx).Pagination(pagination).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CatalogAPI.GetWarehousesCatalog``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `GetWarehousesCatalog`: GetWarehousesCatalog200Response
    fmt.Fprintf(os.Stdout, "Response from `CatalogAPI.GetWarehousesCatalog`: %v\n", resp.GetData())
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWarehousesCatalogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pagination** | [**PaginationInput**](PaginationInput.md) | Required pagination params used to filter the Warehouses catalog.  This parameter exists in v1. | 

### Return type

[**GetWarehousesCatalog200Response**](GetWarehousesCatalog200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1+json, application/json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

