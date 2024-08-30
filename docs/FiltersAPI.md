# \FiltersAPI

All URIs are relative to *https://api.segmentapis.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFilter**](FiltersAPI.md#CreateFilter) | **Post** /filters/create/{integrationId} | Create Filter
[**DeleteFilterById**](FiltersAPI.md#DeleteFilterById) | **Delete** /filters/delete/{id} | Delete Filter By Id
[**GetFilterById**](FiltersAPI.md#GetFilterById) | **Get** /filters/filter/{id} | Get Filter By Id
[**ListFiltersByIntegrationId**](FiltersAPI.md#ListFiltersByIntegrationId) | **Get** /filters/{integrationId} | List Filters By Integration Id
[**UpdateFilterById**](FiltersAPI.md#UpdateFilterById) | **Patch** /filters/update/{id} | Update Filter By Id



## Operation: CreateFilter

> CreateFilter(ctx, integrationId).CreateFilterInput(createFilterInput).Execute()

Create Filter



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
    integrationId := "<id>" // string | 
    createFilterInput := *api.NewCreateFilterInput("Name_example", "If_example") // CreateFilterInput | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    r, err := apiClient.FiltersAPI.CreateFilter(ctx, integrationId).CreateFilterInput(createFilterInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FiltersAPI.CreateFilter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateFilterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createFilterInput** | [**CreateFilterInput**](CreateFilterInput.md) |  | 

### Return type

 (empty response body)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/vnd.segment.v1alpha+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: DeleteFilterById

> DeleteFilterById(ctx, id).ProductArea(productArea).Execute()

Delete Filter By Id



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
    id := "<id>" // string | 
    productArea := "spaces" // string | The product area of the filter  This parameter exists in alpha.

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    r, err := apiClient.FiltersAPI.DeleteFilterById(ctx, id).ProductArea(productArea).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FiltersAPI.DeleteFilterById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFilterByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **productArea** | **string** | The product area of the filter  This parameter exists in alpha. | 

### Return type

 (empty response body)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: GetFilterById

> GetFilterById(ctx, id).ProductArea(productArea).Execute()

Get Filter By Id



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
    id := "<id>" // string | 
    productArea := "spaces" // string | The product area of the filter, which should be spaces (endpoint table should be able to determine the resource)  This parameter exists in alpha.

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    r, err := apiClient.FiltersAPI.GetFilterById(ctx, id).ProductArea(productArea).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FiltersAPI.GetFilterById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFilterByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **productArea** | **string** | The product area of the filter, which should be spaces (endpoint table should be able to determine the resource)  This parameter exists in alpha. | 

### Return type

 (empty response body)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: ListFiltersByIntegrationId

> ListFiltersByIntegrationId(ctx, integrationId).ProductArea(productArea).Pagination(pagination).Execute()

List Filters By Integration Id



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
    integrationId := "<id>" // string | 
    productArea := "spaces" // string | The product area of the filter, which should be spaces (endpoint table should be able to determine the resource)  This parameter exists in alpha.
    pagination := *api.NewListFiltersPaginationInput(float32(123)) // ListFiltersPaginationInput | Pagination parameters.  This parameter exists in alpha. (optional)

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    r, err := apiClient.FiltersAPI.ListFiltersByIntegrationId(ctx, integrationId).ProductArea(productArea).Pagination(pagination).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FiltersAPI.ListFiltersByIntegrationId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListFiltersByIntegrationIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **productArea** | **string** | The product area of the filter, which should be spaces (endpoint table should be able to determine the resource)  This parameter exists in alpha. | 
 **pagination** | [**ListFiltersPaginationInput**](ListFiltersPaginationInput.md) | Pagination parameters.  This parameter exists in alpha. | 

### Return type

 (empty response body)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: UpdateFilterById

> UpdateFilterById(ctx, id).UpdateFilterByIdInput(updateFilterByIdInput).Execute()

Update Filter By Id



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
    id := "<id>" // string | 
    updateFilterByIdInput := *api.NewUpdateFilterByIdInput("IntegrationId_example", "Name_example", "If_example") // UpdateFilterByIdInput | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    r, err := apiClient.FiltersAPI.UpdateFilterById(ctx, id).UpdateFilterByIdInput(updateFilterByIdInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FiltersAPI.UpdateFilterById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFilterByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateFilterByIdInput** | [**UpdateFilterByIdInput**](UpdateFilterByIdInput.md) |  | 

### Return type

 (empty response body)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/vnd.segment.v1alpha+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

