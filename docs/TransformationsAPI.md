# \TransformationsAPI

All URIs are relative to *https://api.segmentapis.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTransformation**](TransformationsAPI.md#CreateTransformation) | **Post** /transformations | Create Transformation
[**DeleteTransformation**](TransformationsAPI.md#DeleteTransformation) | **Delete** /transformations/{transformationId} | Delete Transformation
[**GetTransformation**](TransformationsAPI.md#GetTransformation) | **Get** /transformations/{transformationId} | Get Transformation
[**ListTransformations**](TransformationsAPI.md#ListTransformations) | **Get** /transformations | List Transformations
[**UpdateTransformation**](TransformationsAPI.md#UpdateTransformation) | **Patch** /transformations/{transformationId} | Update Transformation



## Operation: CreateTransformation

> CreateTransformation200Response CreateTransformation(ctx).CreateTransformationV1Input(createTransformationV1Input).Execute()

Create Transformation



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
    createTransformationV1Input := *api.NewCreateTransformationV1Input("Name_example", "SourceId_example", false, "If_example") // CreateTransformationV1Input | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "INSERT BEARER TOKEN HERE"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.TransformationsAPI.CreateTransformation(ctx).CreateTransformationV1Input(createTransformationV1Input).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransformationsAPI.CreateTransformation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `CreateTransformation`: CreateTransformation200Response
    fmt.Fprintf(os.Stdout, "Response from `TransformationsAPI.CreateTransformation`: %v\n", resp.GetData())
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTransformationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTransformationV1Input** | [**CreateTransformationV1Input**](CreateTransformationV1Input.md) |  | 

### Return type

[**CreateTransformation200Response**](CreateTransformation200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/vnd.segment.v1+json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json
- **Accept**: application/vnd.segment.v1+json, application/json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: DeleteTransformation

> DeleteTransformation200Response DeleteTransformation(ctx, transformationId).Execute()

Delete Transformation



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
    transformationId := "2NhipGGSZWdzI7YrXHELB7pVXqR" // string | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "INSERT BEARER TOKEN HERE"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.TransformationsAPI.DeleteTransformation(ctx).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransformationsAPI.DeleteTransformation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `DeleteTransformation`: DeleteTransformation200Response
    fmt.Fprintf(os.Stdout, "Response from `TransformationsAPI.DeleteTransformation`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transformationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTransformationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteTransformation200Response**](DeleteTransformation200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1+json, application/json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: GetTransformation

> GetTransformation200Response GetTransformation(ctx, transformationId).Execute()

Get Transformation



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
    transformationId := "pHrD51Ds35Zjfka84yXQE6" // string | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "INSERT BEARER TOKEN HERE"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.TransformationsAPI.GetTransformation(ctx).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransformationsAPI.GetTransformation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `GetTransformation`: GetTransformation200Response
    fmt.Fprintf(os.Stdout, "Response from `TransformationsAPI.GetTransformation`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transformationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransformationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetTransformation200Response**](GetTransformation200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1+json, application/json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: ListTransformations

> ListTransformations200Response ListTransformations(ctx).Pagination(pagination).Execute()

List Transformations



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
    pagination := *api.NewPaginationInput(float32(123)) // PaginationInput | Pagination options.  This parameter exists in v1.

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "INSERT BEARER TOKEN HERE"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.TransformationsAPI.ListTransformations(ctx).Pagination(pagination).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransformationsAPI.ListTransformations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `ListTransformations`: ListTransformations200Response
    fmt.Fprintf(os.Stdout, "Response from `TransformationsAPI.ListTransformations`: %v\n", resp.GetData())
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTransformationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pagination** | [**PaginationInput**](PaginationInput.md) | Pagination options.  This parameter exists in v1. | 

### Return type

[**ListTransformations200Response**](ListTransformations200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1+json, application/json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: UpdateTransformation

> UpdateTransformation200Response UpdateTransformation(ctx, transformationId).UpdateTransformationV1Input(updateTransformationV1Input).Execute()

Update Transformation



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
    transformationId := "pHrD51Ds35Zjfka84yXQE6" // string | 
    updateTransformationV1Input := *api.NewUpdateTransformationV1Input() // UpdateTransformationV1Input | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "INSERT BEARER TOKEN HERE"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.TransformationsAPI.UpdateTransformation(ctx).UpdateTransformationV1Input(updateTransformationV1Input).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransformationsAPI.UpdateTransformation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `UpdateTransformation`: UpdateTransformation200Response
    fmt.Fprintf(os.Stdout, "Response from `TransformationsAPI.UpdateTransformation`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transformationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTransformationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateTransformationV1Input** | [**UpdateTransformationV1Input**](UpdateTransformationV1Input.md) |  | 

### Return type

[**UpdateTransformation200Response**](UpdateTransformation200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/vnd.segment.v1+json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json
- **Accept**: application/vnd.segment.v1+json, application/json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

