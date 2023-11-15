# \SelectiveSyncAPI

All URIs are relative to *https://api.segmentapis.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAdvancedSyncScheduleFromWarehouse**](SelectiveSyncAPI.md#GetAdvancedSyncScheduleFromWarehouse) | **Get** /warehouses/{warehouseId}/advanced-sync-schedule | Get Advanced Sync Schedule from Warehouse
[**ListSelectiveSyncsFromWarehouseAndSource**](SelectiveSyncAPI.md#ListSelectiveSyncsFromWarehouseAndSource) | **Get** /warehouses/{warehouseId}/connected-sources/{sourceId}/selective-syncs | List Selective Syncs from Warehouse And Source
[**ListSyncsFromWarehouse**](SelectiveSyncAPI.md#ListSyncsFromWarehouse) | **Get** /warehouses/{warehouseId}/syncs | List Syncs from Warehouse
[**ListSyncsFromWarehouseAndSource**](SelectiveSyncAPI.md#ListSyncsFromWarehouseAndSource) | **Get** /warehouses/{warehouseId}/connected-sources/{sourceId}/syncs | List Syncs from Warehouse And Source
[**ReplaceAdvancedSyncScheduleForWarehouse**](SelectiveSyncAPI.md#ReplaceAdvancedSyncScheduleForWarehouse) | **Put** /warehouses/{warehouseId}/advanced-sync-schedule | Replace Advanced Sync Schedule for Warehouse
[**UpdateSelectiveSyncForWarehouse**](SelectiveSyncAPI.md#UpdateSelectiveSyncForWarehouse) | **Patch** /warehouses/{warehouseId}/selective-sync | Update Selective Sync for Warehouse



## Operation: GetAdvancedSyncScheduleFromWarehouse

> GetAdvancedSyncScheduleFromWarehouse200Response GetAdvancedSyncScheduleFromWarehouse(ctx, warehouseId).Execute()

Get Advanced Sync Schedule from Warehouse



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
    warehouseId := "kjU72LCJexvrqL7G4TMHHN" // string | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.SelectiveSyncAPI.GetAdvancedSyncScheduleFromWarehouse(ctx, warehouseId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelectiveSyncAPI.GetAdvancedSyncScheduleFromWarehouse``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `GetAdvancedSyncScheduleFromWarehouse`: GetAdvancedSyncScheduleFromWarehouse200Response
    fmt.Fprintf(os.Stdout, "Response from `SelectiveSyncAPI.GetAdvancedSyncScheduleFromWarehouse`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**warehouseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAdvancedSyncScheduleFromWarehouseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetAdvancedSyncScheduleFromWarehouse200Response**](GetAdvancedSyncScheduleFromWarehouse200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1+json, application/json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: ListSelectiveSyncsFromWarehouseAndSource

> ListSelectiveSyncsFromWarehouseAndSource200Response ListSelectiveSyncsFromWarehouseAndSource(ctx, warehouseId, sourceId).Pagination(pagination).Execute()

List Selective Syncs from Warehouse And Source



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
    warehouseId := "kjU72LCJexvrqL7G4TMHHN" // string | 
    sourceId := "rh5BDZp6QDHvXFCkibm1pR" // string | 
    pagination := *api.NewPaginationInput(int32(123)) // PaginationInput | Defines the pagination parameters.  This parameter exists in v1.

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.SelectiveSyncAPI.ListSelectiveSyncsFromWarehouseAndSource(ctx, warehouseId, sourceId).Pagination(pagination).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelectiveSyncAPI.ListSelectiveSyncsFromWarehouseAndSource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `ListSelectiveSyncsFromWarehouseAndSource`: ListSelectiveSyncsFromWarehouseAndSource200Response
    fmt.Fprintf(os.Stdout, "Response from `SelectiveSyncAPI.ListSelectiveSyncsFromWarehouseAndSource`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**warehouseId** | **string** |  | 
**sourceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSelectiveSyncsFromWarehouseAndSourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pagination** | [**PaginationInput**](PaginationInput.md) | Defines the pagination parameters.  This parameter exists in v1. | 

### Return type

[**ListSelectiveSyncsFromWarehouseAndSource200Response**](ListSelectiveSyncsFromWarehouseAndSource200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1+json, application/json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: ListSyncsFromWarehouse

> ListSyncsFromWarehouse200Response ListSyncsFromWarehouse(ctx, warehouseId).Pagination(pagination).Execute()

List Syncs from Warehouse



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
    warehouseId := "kjU72LCJexvrqL7G4TMHHN" // string | 
    pagination := *api.NewPaginationInput(int32(123)) // PaginationInput | Defines the pagination parameters.  This parameter exists in v1.

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.SelectiveSyncAPI.ListSyncsFromWarehouse(ctx, warehouseId).Pagination(pagination).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelectiveSyncAPI.ListSyncsFromWarehouse``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `ListSyncsFromWarehouse`: ListSyncsFromWarehouse200Response
    fmt.Fprintf(os.Stdout, "Response from `SelectiveSyncAPI.ListSyncsFromWarehouse`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**warehouseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSyncsFromWarehouseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pagination** | [**PaginationInput**](PaginationInput.md) | Defines the pagination parameters.  This parameter exists in v1. | 

### Return type

[**ListSyncsFromWarehouse200Response**](ListSyncsFromWarehouse200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1+json, application/json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: ListSyncsFromWarehouseAndSource

> ListSyncsFromWarehouseAndSource200Response ListSyncsFromWarehouseAndSource(ctx, warehouseId, sourceId).Pagination(pagination).Execute()

List Syncs from Warehouse And Source



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
    warehouseId := "kjU72LCJexvrqL7G4TMHHN" // string | 
    sourceId := "rh5BDZp6QDHvXFCkibm1pR" // string | 
    pagination := *api.NewPaginationInput(int32(123)) // PaginationInput | Defines the pagination parameters.  This parameter exists in v1.

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.SelectiveSyncAPI.ListSyncsFromWarehouseAndSource(ctx, warehouseId, sourceId).Pagination(pagination).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelectiveSyncAPI.ListSyncsFromWarehouseAndSource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `ListSyncsFromWarehouseAndSource`: ListSyncsFromWarehouseAndSource200Response
    fmt.Fprintf(os.Stdout, "Response from `SelectiveSyncAPI.ListSyncsFromWarehouseAndSource`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**warehouseId** | **string** |  | 
**sourceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSyncsFromWarehouseAndSourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pagination** | [**PaginationInput**](PaginationInput.md) | Defines the pagination parameters.  This parameter exists in v1. | 

### Return type

[**ListSyncsFromWarehouseAndSource200Response**](ListSyncsFromWarehouseAndSource200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1+json, application/json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: ReplaceAdvancedSyncScheduleForWarehouse

> ReplaceAdvancedSyncScheduleForWarehouse200Response ReplaceAdvancedSyncScheduleForWarehouse(ctx, warehouseId).ReplaceAdvancedSyncScheduleForWarehouseV1Input(replaceAdvancedSyncScheduleForWarehouseV1Input).Execute()

Replace Advanced Sync Schedule for Warehouse



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
    warehouseId := "kjU72LCJexvrqL7G4TMHHN" // string | 
    replaceAdvancedSyncScheduleForWarehouseV1Input := *api.NewReplaceAdvancedSyncScheduleForWarehouseV1Input(false) // ReplaceAdvancedSyncScheduleForWarehouseV1Input | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.SelectiveSyncAPI.ReplaceAdvancedSyncScheduleForWarehouse(ctx, warehouseId).ReplaceAdvancedSyncScheduleForWarehouseV1Input(replaceAdvancedSyncScheduleForWarehouseV1Input).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelectiveSyncAPI.ReplaceAdvancedSyncScheduleForWarehouse``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `ReplaceAdvancedSyncScheduleForWarehouse`: ReplaceAdvancedSyncScheduleForWarehouse200Response
    fmt.Fprintf(os.Stdout, "Response from `SelectiveSyncAPI.ReplaceAdvancedSyncScheduleForWarehouse`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**warehouseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceAdvancedSyncScheduleForWarehouseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **replaceAdvancedSyncScheduleForWarehouseV1Input** | [**ReplaceAdvancedSyncScheduleForWarehouseV1Input**](ReplaceAdvancedSyncScheduleForWarehouseV1Input.md) |  | 

### Return type

[**ReplaceAdvancedSyncScheduleForWarehouse200Response**](ReplaceAdvancedSyncScheduleForWarehouse200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/vnd.segment.v1+json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json
- **Accept**: application/vnd.segment.v1+json, application/json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: UpdateSelectiveSyncForWarehouse

> UpdateSelectiveSyncForWarehouse200Response UpdateSelectiveSyncForWarehouse(ctx, warehouseId).UpdateSelectiveSyncForWarehouseV1Input(updateSelectiveSyncForWarehouseV1Input).Execute()

Update Selective Sync for Warehouse



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
    warehouseId := "kjU72LCJexvrqL7G4TMHHN" // string | 
    updateSelectiveSyncForWarehouseV1Input := *api.NewUpdateSelectiveSyncForWarehouseV1Input([]api.WarehouseSyncOverrideV1{*api.NewWarehouseSyncOverrideV1("SourceId_example", false)}) // UpdateSelectiveSyncForWarehouseV1Input | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.SelectiveSyncAPI.UpdateSelectiveSyncForWarehouse(ctx, warehouseId).UpdateSelectiveSyncForWarehouseV1Input(updateSelectiveSyncForWarehouseV1Input).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelectiveSyncAPI.UpdateSelectiveSyncForWarehouse``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `UpdateSelectiveSyncForWarehouse`: UpdateSelectiveSyncForWarehouse200Response
    fmt.Fprintf(os.Stdout, "Response from `SelectiveSyncAPI.UpdateSelectiveSyncForWarehouse`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**warehouseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSelectiveSyncForWarehouseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateSelectiveSyncForWarehouseV1Input** | [**UpdateSelectiveSyncForWarehouseV1Input**](UpdateSelectiveSyncForWarehouseV1Input.md) |  | 

### Return type

[**UpdateSelectiveSyncForWarehouse200Response**](UpdateSelectiveSyncForWarehouse200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/vnd.segment.v1+json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json
- **Accept**: application/vnd.segment.v1+json, application/json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

