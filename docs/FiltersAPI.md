# \FiltersAPI

All URIs are relative to *https://api.segmentapis.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFilter**](FiltersAPI.md#CreateFilter) | **Post** /filters | Create Filter
[**DeleteFilterById**](FiltersAPI.md#DeleteFilterById) | **Delete** /filters/{id} | Delete Filter By Id
[**GetFilterById**](FiltersAPI.md#GetFilterById) | **Get** /filters/{id} | Get Filter By Id
[**ListFiltersByIntegrationId**](FiltersAPI.md#ListFiltersByIntegrationId) | **Get** /filters | List Filters By Integration Id
[**UpdateFilterById**](FiltersAPI.md#UpdateFilterById) | **Patch** /filters/{id} | Update Filter By Id



## Operation: CreateFilter

> CreateFilter200Response CreateFilter(ctx).CreateFilterInput(createFilterInput).Execute()

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
    createFilterInput := *api.NewCreateFilterInput("IntegrationId_example", "Name_example", "If_example") // CreateFilterInput | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.FiltersAPI.CreateFilter(ctx).CreateFilterInput(createFilterInput).Execute()
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
    // response from `CreateFilter`: CreateFilter200Response
    fmt.Fprintf(os.Stdout, "Response from `FiltersAPI.CreateFilter`: %v\n", resp.GetData())
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFilterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createFilterInput** | [**CreateFilterInput**](CreateFilterInput.md) |  | 

### Return type

[**CreateFilter200Response**](CreateFilter200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/vnd.segment.v1alpha+json
- **Accept**: application/vnd.segment.v1alpha+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: DeleteFilterById

> DeleteFilterById200Response DeleteFilterById(ctx, id).Execute()

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

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.FiltersAPI.DeleteFilterById(ctx, id).Execute()
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
    // response from `DeleteFilterById`: DeleteFilterById200Response
    fmt.Fprintf(os.Stdout, "Response from `FiltersAPI.DeleteFilterById`: %v\n", resp.GetData())
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


### Return type

[**DeleteFilterById200Response**](DeleteFilterById200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1alpha+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: GetFilterById

> GetFilterById200Response GetFilterById(ctx, id).Execute()

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

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.FiltersAPI.GetFilterById(ctx, id).Execute()
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
    // response from `GetFilterById`: GetFilterById200Response
    fmt.Fprintf(os.Stdout, "Response from `FiltersAPI.GetFilterById`: %v\n", resp.GetData())
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


### Return type

[**GetFilterById200Response**](GetFilterById200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1alpha+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: ListFiltersByIntegrationId

> ListFiltersByIntegrationId200Response ListFiltersByIntegrationId(ctx).IntegrationId(integrationId).Pagination(pagination).Execute()

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
    integrationId := "<id>" // string | The integration id used to fetch filters.  This parameter exists in alpha.
    pagination := *api.NewListFiltersPaginationInput(float32(123)) // ListFiltersPaginationInput | Pagination parameters.  This parameter exists in alpha. (optional)

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.FiltersAPI.ListFiltersByIntegrationId(ctx).IntegrationId(integrationId).Pagination(pagination).Execute()
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
    // response from `ListFiltersByIntegrationId`: ListFiltersByIntegrationId200Response
    fmt.Fprintf(os.Stdout, "Response from `FiltersAPI.ListFiltersByIntegrationId`: %v\n", resp.GetData())
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListFiltersByIntegrationIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **integrationId** | **string** | The integration id used to fetch filters.  This parameter exists in alpha. | 
 **pagination** | [**ListFiltersPaginationInput**](ListFiltersPaginationInput.md) | Pagination parameters.  This parameter exists in alpha. | 

### Return type

[**ListFiltersByIntegrationId200Response**](ListFiltersByIntegrationId200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1alpha+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: UpdateFilterById

> UpdateFilterById200Response UpdateFilterById(ctx, id).UpdateFilterByIdInput(updateFilterByIdInput).Execute()

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
    updateFilterByIdInput := *api.NewUpdateFilterByIdInput("IntegrationId_example") // UpdateFilterByIdInput | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.FiltersAPI.UpdateFilterById(ctx, id).UpdateFilterByIdInput(updateFilterByIdInput).Execute()
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
    // response from `UpdateFilterById`: UpdateFilterById200Response
    fmt.Fprintf(os.Stdout, "Response from `FiltersAPI.UpdateFilterById`: %v\n", resp.GetData())
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

[**UpdateFilterById200Response**](UpdateFilterById200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/vnd.segment.v1alpha+json
- **Accept**: application/vnd.segment.v1alpha+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

