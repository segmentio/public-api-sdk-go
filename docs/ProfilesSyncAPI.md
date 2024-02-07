# \ProfilesSyncAPI

All URIs are relative to *https://api.segmentapis.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateProfilesWarehouse**](ProfilesSyncAPI.md#CreateProfilesWarehouse) | **Post** /spaces/{spaceId}/profiles-warehouses | Create Profiles Warehouse
[**ListProfilesWarehouseInSpace**](ProfilesSyncAPI.md#ListProfilesWarehouseInSpace) | **Get** /spaces/{spaceId}/profiles-warehouses | List Profiles Warehouse in Space
[**ListSelectiveSyncsFromWarehouseAndSpace**](ProfilesSyncAPI.md#ListSelectiveSyncsFromWarehouseAndSpace) | **Get** /spaces/{spaceId}/profiles-warehouses/{warehouseId}/selective-syncs | List Selective Syncs from Warehouse And Space
[**RemoveProfilesWarehouseFromSpace**](ProfilesSyncAPI.md#RemoveProfilesWarehouseFromSpace) | **Delete** /spaces/{spaceId}/profiles-warehouses/{warehouseId} | Remove Profiles Warehouse from Space
[**UpdateProfilesWarehouseForSpaceWarehouse**](ProfilesSyncAPI.md#UpdateProfilesWarehouseForSpaceWarehouse) | **Patch** /spaces/{spaceId}/profiles-warehouses/{warehouseId} | Update Profiles Warehouse for Space Warehouse
[**UpdateSelectiveSyncForWarehouseAndSpace**](ProfilesSyncAPI.md#UpdateSelectiveSyncForWarehouseAndSpace) | **Patch** /spaces/{spaceId}/profiles-warehouses/{warehouseId}/selective-syncs | Update Selective Sync for Warehouse And Space



## Operation: CreateProfilesWarehouse

> CreateProfilesWarehouse201Response CreateProfilesWarehouse(ctx, spaceId).CreateProfilesWarehouseAlphaInput(createProfilesWarehouseAlphaInput).Execute()

Create Profiles Warehouse



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
    createProfilesWarehouseAlphaInput := *api.NewCreateProfilesWarehouseAlphaInput("MetadataId_example", map[string]interface{}{"key": interface{}(123)}) // CreateProfilesWarehouseAlphaInput | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.ProfilesSyncAPI.CreateProfilesWarehouse(ctx, spaceId).CreateProfilesWarehouseAlphaInput(createProfilesWarehouseAlphaInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProfilesSyncAPI.CreateProfilesWarehouse``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `CreateProfilesWarehouse`: CreateProfilesWarehouse201Response
    fmt.Fprintf(os.Stdout, "Response from `ProfilesSyncAPI.CreateProfilesWarehouse`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateProfilesWarehouseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createProfilesWarehouseAlphaInput** | [**CreateProfilesWarehouseAlphaInput**](CreateProfilesWarehouseAlphaInput.md) |  | 

### Return type

[**CreateProfilesWarehouse201Response**](CreateProfilesWarehouse201Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/vnd.segment.v1alpha+json
- **Accept**: application/vnd.segment.v1alpha+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: ListProfilesWarehouseInSpace

> ListProfilesWarehouseInSpace200Response ListProfilesWarehouseInSpace(ctx, spaceId).Pagination(pagination).Execute()

List Profiles Warehouse in Space



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
    pagination := *api.NewPaginationInput(10) // PaginationInput | Defines the pagination parameters.  This parameter exists in alpha. (optional)

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.ProfilesSyncAPI.ListProfilesWarehouseInSpace(ctx, spaceId).Pagination(pagination).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProfilesSyncAPI.ListProfilesWarehouseInSpace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `ListProfilesWarehouseInSpace`: ListProfilesWarehouseInSpace200Response
    fmt.Fprintf(os.Stdout, "Response from `ProfilesSyncAPI.ListProfilesWarehouseInSpace`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListProfilesWarehouseInSpaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pagination** | [**PaginationInput**](PaginationInput.md) | Defines the pagination parameters.  This parameter exists in alpha. | 

### Return type

[**ListProfilesWarehouseInSpace200Response**](ListProfilesWarehouseInSpace200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1alpha+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: ListSelectiveSyncsFromWarehouseAndSpace

> ListSelectiveSyncsFromWarehouseAndSpace200Response ListSelectiveSyncsFromWarehouseAndSpace(ctx, spaceId, warehouseId).Pagination(pagination).Execute()

List Selective Syncs from Warehouse And Space



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
    warehouseId := "fQyLbqjfwaqg9mr3hDQ7We" // string | 
    pagination := *api.NewPaginationInput(10) // PaginationInput | Defines the pagination parameters.  This parameter exists in alpha.

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.ProfilesSyncAPI.ListSelectiveSyncsFromWarehouseAndSpace(ctx, spaceId, warehouseId).Pagination(pagination).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProfilesSyncAPI.ListSelectiveSyncsFromWarehouseAndSpace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `ListSelectiveSyncsFromWarehouseAndSpace`: ListSelectiveSyncsFromWarehouseAndSpace200Response
    fmt.Fprintf(os.Stdout, "Response from `ProfilesSyncAPI.ListSelectiveSyncsFromWarehouseAndSpace`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** |  | 
**warehouseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSelectiveSyncsFromWarehouseAndSpaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pagination** | [**PaginationInput**](PaginationInput.md) | Defines the pagination parameters.  This parameter exists in alpha. | 

### Return type

[**ListSelectiveSyncsFromWarehouseAndSpace200Response**](ListSelectiveSyncsFromWarehouseAndSpace200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1alpha+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: RemoveProfilesWarehouseFromSpace

> RemoveProfilesWarehouseFromSpace200Response RemoveProfilesWarehouseFromSpace(ctx, spaceId, warehouseId).Execute()

Remove Profiles Warehouse from Space



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
    warehouseId := "qABd3NVTPfTLQ3kXWoBhgi" // string | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.ProfilesSyncAPI.RemoveProfilesWarehouseFromSpace(ctx, spaceId, warehouseId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProfilesSyncAPI.RemoveProfilesWarehouseFromSpace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `RemoveProfilesWarehouseFromSpace`: RemoveProfilesWarehouseFromSpace200Response
    fmt.Fprintf(os.Stdout, "Response from `ProfilesSyncAPI.RemoveProfilesWarehouseFromSpace`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** |  | 
**warehouseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveProfilesWarehouseFromSpaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RemoveProfilesWarehouseFromSpace200Response**](RemoveProfilesWarehouseFromSpace200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1alpha+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: UpdateProfilesWarehouseForSpaceWarehouse

> UpdateProfilesWarehouseForSpaceWarehouse200Response UpdateProfilesWarehouseForSpaceWarehouse(ctx, spaceId, warehouseId).UpdateProfilesWarehouseForSpaceWarehouseAlphaInput(updateProfilesWarehouseForSpaceWarehouseAlphaInput).Execute()

Update Profiles Warehouse for Space Warehouse



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
    warehouseId := "3eadBBqVMQD2DEtaWXSkqA" // string | 
    updateProfilesWarehouseForSpaceWarehouseAlphaInput := *api.NewUpdateProfilesWarehouseForSpaceWarehouseAlphaInput(map[string]interface{}{"key": interface{}(123)}) // UpdateProfilesWarehouseForSpaceWarehouseAlphaInput | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.ProfilesSyncAPI.UpdateProfilesWarehouseForSpaceWarehouse(ctx, spaceId, warehouseId).UpdateProfilesWarehouseForSpaceWarehouseAlphaInput(updateProfilesWarehouseForSpaceWarehouseAlphaInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProfilesSyncAPI.UpdateProfilesWarehouseForSpaceWarehouse``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `UpdateProfilesWarehouseForSpaceWarehouse`: UpdateProfilesWarehouseForSpaceWarehouse200Response
    fmt.Fprintf(os.Stdout, "Response from `ProfilesSyncAPI.UpdateProfilesWarehouseForSpaceWarehouse`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** |  | 
**warehouseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProfilesWarehouseForSpaceWarehouseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateProfilesWarehouseForSpaceWarehouseAlphaInput** | [**UpdateProfilesWarehouseForSpaceWarehouseAlphaInput**](UpdateProfilesWarehouseForSpaceWarehouseAlphaInput.md) |  | 

### Return type

[**UpdateProfilesWarehouseForSpaceWarehouse200Response**](UpdateProfilesWarehouseForSpaceWarehouse200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/vnd.segment.v1alpha+json
- **Accept**: application/vnd.segment.v1alpha+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: UpdateSelectiveSyncForWarehouseAndSpace

> UpdateSelectiveSyncForWarehouseAndSpace200Response UpdateSelectiveSyncForWarehouseAndSpace(ctx, spaceId, warehouseId).UpdateSelectiveSyncForWarehouseAndSpaceAlphaInput(updateSelectiveSyncForWarehouseAndSpaceAlphaInput).Execute()

Update Selective Sync for Warehouse And Space



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
    warehouseId := "qABd3NVTPfTLQ3kXWoBhgi" // string | 
    updateSelectiveSyncForWarehouseAndSpaceAlphaInput := *api.NewUpdateSelectiveSyncForWarehouseAndSpaceAlphaInput() // UpdateSelectiveSyncForWarehouseAndSpaceAlphaInput | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.ProfilesSyncAPI.UpdateSelectiveSyncForWarehouseAndSpace(ctx, spaceId, warehouseId).UpdateSelectiveSyncForWarehouseAndSpaceAlphaInput(updateSelectiveSyncForWarehouseAndSpaceAlphaInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProfilesSyncAPI.UpdateSelectiveSyncForWarehouseAndSpace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `UpdateSelectiveSyncForWarehouseAndSpace`: UpdateSelectiveSyncForWarehouseAndSpace200Response
    fmt.Fprintf(os.Stdout, "Response from `ProfilesSyncAPI.UpdateSelectiveSyncForWarehouseAndSpace`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** |  | 
**warehouseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSelectiveSyncForWarehouseAndSpaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateSelectiveSyncForWarehouseAndSpaceAlphaInput** | [**UpdateSelectiveSyncForWarehouseAndSpaceAlphaInput**](UpdateSelectiveSyncForWarehouseAndSpaceAlphaInput.md) |  | 

### Return type

[**UpdateSelectiveSyncForWarehouseAndSpace200Response**](UpdateSelectiveSyncForWarehouseAndSpace200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/vnd.segment.v1alpha+json
- **Accept**: application/vnd.segment.v1alpha+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

