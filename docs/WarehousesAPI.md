# \WarehousesAPI

All URIs are relative to *https://api.segmentapis.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddConnectionFromSourceToWarehouse**](WarehousesAPI.md#AddConnectionFromSourceToWarehouse) | **Post** /warehouses/{warehouseId}/connected-sources/{sourceId} | Add Connection from Source to Warehouse
[**CreateValidationInWarehouse**](WarehousesAPI.md#CreateValidationInWarehouse) | **Post** /warehouses/validation | Create Validation in Warehouse
[**CreateWarehouse**](WarehousesAPI.md#CreateWarehouse) | **Post** /warehouses | Create Warehouse
[**DeleteWarehouse**](WarehousesAPI.md#DeleteWarehouse) | **Delete** /warehouses/{warehouseId} | Delete Warehouse
[**GetConnectionStateFromWarehouse**](WarehousesAPI.md#GetConnectionStateFromWarehouse) | **Get** /warehouses/{warehouseId}/connection-state | Get Connection State from Warehouse
[**GetWarehouse**](WarehousesAPI.md#GetWarehouse) | **Get** /warehouses/{warehouseId} | Get Warehouse
[**ListConnectedSourcesFromWarehouse**](WarehousesAPI.md#ListConnectedSourcesFromWarehouse) | **Get** /warehouses/{warehouseId}/connected-sources | List Connected Sources from Warehouse
[**ListWarehouses**](WarehousesAPI.md#ListWarehouses) | **Get** /warehouses | List Warehouses
[**RemoveSourceConnectionFromWarehouse**](WarehousesAPI.md#RemoveSourceConnectionFromWarehouse) | **Delete** /warehouses/{warehouseId}/connected-sources/{sourceId} | Remove Source Connection from Warehouse
[**UpdateWarehouse**](WarehousesAPI.md#UpdateWarehouse) | **Patch** /warehouses/{warehouseId} | Update Warehouse



## Operation: AddConnectionFromSourceToWarehouse

> AddConnectionFromSourceToWarehouse201Response AddConnectionFromSourceToWarehouse(ctx, warehouseId, sourceId).Execute()

Add Connection from Source to Warehouse



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

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.WarehousesAPI.AddConnectionFromSourceToWarehouse(ctx, warehouseId, sourceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WarehousesAPI.AddConnectionFromSourceToWarehouse``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `AddConnectionFromSourceToWarehouse`: AddConnectionFromSourceToWarehouse201Response
    fmt.Fprintf(os.Stdout, "Response from `WarehousesAPI.AddConnectionFromSourceToWarehouse`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**warehouseId** | **string** |  | 
**sourceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddConnectionFromSourceToWarehouseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AddConnectionFromSourceToWarehouse201Response**](AddConnectionFromSourceToWarehouse201Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1+json, application/json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: CreateValidationInWarehouse

> CreateValidationInWarehouse200Response CreateValidationInWarehouse(ctx).CreateValidationInWarehouseV1Input(createValidationInWarehouseV1Input).Execute()

Create Validation in Warehouse



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
    createValidationInWarehouseV1Input := *api.NewCreateValidationInWarehouseV1Input("MetadataId_example", map[string]interface{}{"key": interface{}(123)}) // CreateValidationInWarehouseV1Input | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.WarehousesAPI.CreateValidationInWarehouse(ctx).CreateValidationInWarehouseV1Input(createValidationInWarehouseV1Input).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WarehousesAPI.CreateValidationInWarehouse``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `CreateValidationInWarehouse`: CreateValidationInWarehouse200Response
    fmt.Fprintf(os.Stdout, "Response from `WarehousesAPI.CreateValidationInWarehouse`: %v\n", resp.GetData())
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateValidationInWarehouseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createValidationInWarehouseV1Input** | [**CreateValidationInWarehouseV1Input**](CreateValidationInWarehouseV1Input.md) |  | 

### Return type

[**CreateValidationInWarehouse200Response**](CreateValidationInWarehouse200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json, application/vnd.segment.v1+json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json
- **Accept**: application/vnd.segment.v1+json, application/json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: CreateWarehouse

> CreateWarehouse201Response CreateWarehouse(ctx).CreateWarehouseV1Input(createWarehouseV1Input).Execute()

Create Warehouse



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
    createWarehouseV1Input := *api.NewCreateWarehouseV1Input("MetadataId_example", map[string]interface{}{"key": interface{}(123)}) // CreateWarehouseV1Input | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.WarehousesAPI.CreateWarehouse(ctx).CreateWarehouseV1Input(createWarehouseV1Input).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WarehousesAPI.CreateWarehouse``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `CreateWarehouse`: CreateWarehouse201Response
    fmt.Fprintf(os.Stdout, "Response from `WarehousesAPI.CreateWarehouse`: %v\n", resp.GetData())
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWarehouseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createWarehouseV1Input** | [**CreateWarehouseV1Input**](CreateWarehouseV1Input.md) |  | 

### Return type

[**CreateWarehouse201Response**](CreateWarehouse201Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json, application/vnd.segment.v1+json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json
- **Accept**: application/vnd.segment.v1+json, application/json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: DeleteWarehouse

> DeleteWarehouse200Response DeleteWarehouse(ctx, warehouseId).Execute()

Delete Warehouse



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
    warehouseId := "tmiTtiPi58udvDAjcxKUJY" // string | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.WarehousesAPI.DeleteWarehouse(ctx, warehouseId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WarehousesAPI.DeleteWarehouse``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `DeleteWarehouse`: DeleteWarehouse200Response
    fmt.Fprintf(os.Stdout, "Response from `WarehousesAPI.DeleteWarehouse`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**warehouseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWarehouseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteWarehouse200Response**](DeleteWarehouse200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1+json, application/json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: GetConnectionStateFromWarehouse

> GetConnectionStateFromWarehouse200Response GetConnectionStateFromWarehouse(ctx, warehouseId).Execute()

Get Connection State from Warehouse



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
    resp, r, err := apiClient.WarehousesAPI.GetConnectionStateFromWarehouse(ctx, warehouseId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WarehousesAPI.GetConnectionStateFromWarehouse``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `GetConnectionStateFromWarehouse`: GetConnectionStateFromWarehouse200Response
    fmt.Fprintf(os.Stdout, "Response from `WarehousesAPI.GetConnectionStateFromWarehouse`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**warehouseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectionStateFromWarehouseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetConnectionStateFromWarehouse200Response**](GetConnectionStateFromWarehouse200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1+json, application/json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: GetWarehouse

> GetWarehouse200Response GetWarehouse(ctx, warehouseId).Execute()

Get Warehouse



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
    resp, r, err := apiClient.WarehousesAPI.GetWarehouse(ctx, warehouseId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WarehousesAPI.GetWarehouse``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `GetWarehouse`: GetWarehouse200Response
    fmt.Fprintf(os.Stdout, "Response from `WarehousesAPI.GetWarehouse`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**warehouseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWarehouseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetWarehouse200Response**](GetWarehouse200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1+json, application/json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: ListConnectedSourcesFromWarehouse

> ListConnectedSourcesFromWarehouse200Response ListConnectedSourcesFromWarehouse(ctx, warehouseId).Pagination(pagination).Execute()

List Connected Sources from Warehouse



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
    pagination := *api.NewPaginationInput(10) // PaginationInput | Defines the pagination parameters.  This parameter exists in v1. (optional)

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.WarehousesAPI.ListConnectedSourcesFromWarehouse(ctx, warehouseId).Pagination(pagination).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WarehousesAPI.ListConnectedSourcesFromWarehouse``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `ListConnectedSourcesFromWarehouse`: ListConnectedSourcesFromWarehouse200Response
    fmt.Fprintf(os.Stdout, "Response from `WarehousesAPI.ListConnectedSourcesFromWarehouse`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**warehouseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListConnectedSourcesFromWarehouseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pagination** | [**PaginationInput**](PaginationInput.md) | Defines the pagination parameters.  This parameter exists in v1. | 

### Return type

[**ListConnectedSourcesFromWarehouse200Response**](ListConnectedSourcesFromWarehouse200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1+json, application/json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: ListWarehouses

> ListWarehouses200Response ListWarehouses(ctx).Pagination(pagination).Execute()

List Warehouses



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
    pagination := *api.NewPaginationInput(10) // PaginationInput | Defines the pagination parameters.  This parameter exists in v1. (optional)

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.WarehousesAPI.ListWarehouses(ctx).Pagination(pagination).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WarehousesAPI.ListWarehouses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `ListWarehouses`: ListWarehouses200Response
    fmt.Fprintf(os.Stdout, "Response from `WarehousesAPI.ListWarehouses`: %v\n", resp.GetData())
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListWarehousesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pagination** | [**PaginationInput**](PaginationInput.md) | Defines the pagination parameters.  This parameter exists in v1. | 

### Return type

[**ListWarehouses200Response**](ListWarehouses200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1+json, application/json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: RemoveSourceConnectionFromWarehouse

> RemoveSourceConnectionFromWarehouse200Response RemoveSourceConnectionFromWarehouse(ctx, warehouseId, sourceId).Execute()

Remove Source Connection from Warehouse



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

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.WarehousesAPI.RemoveSourceConnectionFromWarehouse(ctx, warehouseId, sourceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WarehousesAPI.RemoveSourceConnectionFromWarehouse``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `RemoveSourceConnectionFromWarehouse`: RemoveSourceConnectionFromWarehouse200Response
    fmt.Fprintf(os.Stdout, "Response from `WarehousesAPI.RemoveSourceConnectionFromWarehouse`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**warehouseId** | **string** |  | 
**sourceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveSourceConnectionFromWarehouseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RemoveSourceConnectionFromWarehouse200Response**](RemoveSourceConnectionFromWarehouse200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1+json, application/json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: UpdateWarehouse

> UpdateWarehouse200Response UpdateWarehouse(ctx, warehouseId).UpdateWarehouseV1Input(updateWarehouseV1Input).Execute()

Update Warehouse



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
    updateWarehouseV1Input := *api.NewUpdateWarehouseV1Input(map[string]interface{}{"key": interface{}(123)}) // UpdateWarehouseV1Input | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.WarehousesAPI.UpdateWarehouse(ctx, warehouseId).UpdateWarehouseV1Input(updateWarehouseV1Input).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WarehousesAPI.UpdateWarehouse``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `UpdateWarehouse`: UpdateWarehouse200Response
    fmt.Fprintf(os.Stdout, "Response from `WarehousesAPI.UpdateWarehouse`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**warehouseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWarehouseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateWarehouseV1Input** | [**UpdateWarehouseV1Input**](UpdateWarehouseV1Input.md) |  | 

### Return type

[**UpdateWarehouse200Response**](UpdateWarehouse200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json, application/vnd.segment.v1+json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json
- **Accept**: application/vnd.segment.v1+json, application/json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

