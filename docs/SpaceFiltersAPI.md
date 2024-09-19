# \SpaceFiltersAPI

All URIs are relative to *https://api.segmentapis.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFilterForSpace**](SpaceFiltersAPI.md#CreateFilterForSpace) | **Post** /filters | Create Filter for Space
[**DeleteFilterById**](SpaceFiltersAPI.md#DeleteFilterById) | **Delete** /filters/{id} | Delete Filter By Id
[**GetFilterById**](SpaceFiltersAPI.md#GetFilterById) | **Get** /filters/{id} | Get Filter By Id
[**ListFiltersForSpace**](SpaceFiltersAPI.md#ListFiltersForSpace) | **Get** /filters | List Filters for Space
[**UpdateFilterById**](SpaceFiltersAPI.md#UpdateFilterById) | **Patch** /filters/{id} | Update Filter By Id



## Operation: CreateFilterForSpace

> CreateFilterForSpace200Response CreateFilterForSpace(ctx).CreateFilterForSpaceInput(createFilterForSpaceInput).Execute()

Create Filter for Space



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
    createFilterForSpaceInput := *api.NewCreateFilterForSpaceInput("IntegrationId_example", "Name_example", "If_example") // CreateFilterForSpaceInput | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.SpaceFiltersAPI.CreateFilterForSpace(ctx).CreateFilterForSpaceInput(createFilterForSpaceInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpaceFiltersAPI.CreateFilterForSpace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `CreateFilterForSpace`: CreateFilterForSpace200Response
    fmt.Fprintf(os.Stdout, "Response from `SpaceFiltersAPI.CreateFilterForSpace`: %v\n", resp.GetData())
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFilterForSpaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createFilterForSpaceInput** | [**CreateFilterForSpaceInput**](CreateFilterForSpaceInput.md) |  | 

### Return type

[**CreateFilterForSpace200Response**](CreateFilterForSpace200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json
- **Accept**: application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json, application/json

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
    resp, r, err := apiClient.SpaceFiltersAPI.DeleteFilterById(ctx, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpaceFiltersAPI.DeleteFilterById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `DeleteFilterById`: DeleteFilterById200Response
    fmt.Fprintf(os.Stdout, "Response from `SpaceFiltersAPI.DeleteFilterById`: %v\n", resp.GetData())
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
- **Accept**: application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json, application/json

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
    resp, r, err := apiClient.SpaceFiltersAPI.GetFilterById(ctx, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpaceFiltersAPI.GetFilterById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `GetFilterById`: GetFilterById200Response
    fmt.Fprintf(os.Stdout, "Response from `SpaceFiltersAPI.GetFilterById`: %v\n", resp.GetData())
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
- **Accept**: application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: ListFiltersForSpace

> ListFiltersForSpace200Response ListFiltersForSpace(ctx).IntegrationId(integrationId).Pagination(pagination).Execute()

List Filters for Space



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
    integrationId := "<id>" // string | The Space Id for which to fetch filters  This parameter exists in beta.
    pagination := *api.NewListFiltersPaginationInput(float32(123)) // ListFiltersPaginationInput | Pagination parameters.  This parameter exists in beta. (optional)

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.SpaceFiltersAPI.ListFiltersForSpace(ctx).IntegrationId(integrationId).Pagination(pagination).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpaceFiltersAPI.ListFiltersForSpace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `ListFiltersForSpace`: ListFiltersForSpace200Response
    fmt.Fprintf(os.Stdout, "Response from `SpaceFiltersAPI.ListFiltersForSpace`: %v\n", resp.GetData())
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListFiltersForSpaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **integrationId** | **string** | The Space Id for which to fetch filters  This parameter exists in beta. | 
 **pagination** | [**ListFiltersPaginationInput**](ListFiltersPaginationInput.md) | Pagination parameters.  This parameter exists in beta. | 

### Return type

[**ListFiltersForSpace200Response**](ListFiltersForSpace200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json, application/json

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
    resp, r, err := apiClient.SpaceFiltersAPI.UpdateFilterById(ctx, id).UpdateFilterByIdInput(updateFilterByIdInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpaceFiltersAPI.UpdateFilterById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `UpdateFilterById`: UpdateFilterById200Response
    fmt.Fprintf(os.Stdout, "Response from `SpaceFiltersAPI.UpdateFilterById`: %v\n", resp.GetData())
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

- **Content-Type**: application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json
- **Accept**: application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

