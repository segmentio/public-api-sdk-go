# \ReverseETLAPI

All URIs are relative to *https://api.segmentapis.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateReverseEtlModel**](ReverseETLAPI.md#CreateReverseEtlModel) | **Post** /reverse-etl-models | Create Reverse Etl Model
[**DeleteReverseEtlModel**](ReverseETLAPI.md#DeleteReverseEtlModel) | **Delete** /reverse-etl-models/{modelId} | Delete Reverse Etl Model
[**GetReverseEtlModel**](ReverseETLAPI.md#GetReverseEtlModel) | **Get** /reverse-etl-models/{modelId} | Get Reverse Etl Model
[**ListReverseEtlModels**](ReverseETLAPI.md#ListReverseEtlModels) | **Get** /reverse-etl-models | List Reverse Etl Models
[**UpdateReverseEtlModel**](ReverseETLAPI.md#UpdateReverseEtlModel) | **Patch** /reverse-etl-models/{modelId} | Update Reverse Etl Model



## Operation: CreateReverseEtlModel

> CreateReverseEtlModel200Response CreateReverseEtlModel(ctx).CreateReverseEtlModelInput(createReverseEtlModelInput).Execute()

Create Reverse Etl Model



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
    createReverseEtlModelInput := *api.NewCreateReverseEtlModelInput("SourceId_example", "Name_example", "Description_example", false, "ScheduleStrategy_example", map[string]interface{}{"key": interface{}(123)}, "Query_example", "QueryIdentifierColumn_example") // CreateReverseEtlModelInput | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.ReverseETLAPI.CreateReverseEtlModel(ctx).CreateReverseEtlModelInput(createReverseEtlModelInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReverseETLAPI.CreateReverseEtlModel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `CreateReverseEtlModel`: CreateReverseEtlModel200Response
    fmt.Fprintf(os.Stdout, "Response from `ReverseETLAPI.CreateReverseEtlModel`: %v\n", resp.GetData())
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateReverseEtlModelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createReverseEtlModelInput** | [**CreateReverseEtlModelInput**](CreateReverseEtlModelInput.md) |  | 

### Return type

[**CreateReverseEtlModel200Response**](CreateReverseEtlModel200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/vnd.segment.v1alpha+json
- **Accept**: application/vnd.segment.v1alpha+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: DeleteReverseEtlModel

> DeleteReverseEtlModel200Response DeleteReverseEtlModel(ctx, modelId).Execute()

Delete Reverse Etl Model



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
    modelId := "aow61ZsjXFRsUqB5wWmZES" // string | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.ReverseETLAPI.DeleteReverseEtlModel(ctx).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReverseETLAPI.DeleteReverseEtlModel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `DeleteReverseEtlModel`: DeleteReverseEtlModel200Response
    fmt.Fprintf(os.Stdout, "Response from `ReverseETLAPI.DeleteReverseEtlModel`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**modelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteReverseEtlModelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteReverseEtlModel200Response**](DeleteReverseEtlModel200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1alpha+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: GetReverseEtlModel

> GetReverseEtlModel200Response GetReverseEtlModel(ctx, modelId).Execute()

Get Reverse Etl Model



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
    modelId := "MaAeg9yDd1UZTBeEYDiVw" // string | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.ReverseETLAPI.GetReverseEtlModel(ctx).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReverseETLAPI.GetReverseEtlModel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `GetReverseEtlModel`: GetReverseEtlModel200Response
    fmt.Fprintf(os.Stdout, "Response from `ReverseETLAPI.GetReverseEtlModel`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**modelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReverseEtlModelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetReverseEtlModel200Response**](GetReverseEtlModel200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1alpha+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: ListReverseEtlModels

> ListReverseEtlModels200Response ListReverseEtlModels(ctx).Pagination(pagination).Execute()

List Reverse Etl Models



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
    pagination := *api.NewPaginationInput(float32(123)) // PaginationInput | Defines the pagination parameters.  This parameter exists in alpha.

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.ReverseETLAPI.ListReverseEtlModels(ctx).Pagination(pagination).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReverseETLAPI.ListReverseEtlModels``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `ListReverseEtlModels`: ListReverseEtlModels200Response
    fmt.Fprintf(os.Stdout, "Response from `ReverseETLAPI.ListReverseEtlModels`: %v\n", resp.GetData())
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListReverseEtlModelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pagination** | [**PaginationInput**](PaginationInput.md) | Defines the pagination parameters.  This parameter exists in alpha. | 

### Return type

[**ListReverseEtlModels200Response**](ListReverseEtlModels200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1alpha+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: UpdateReverseEtlModel

> UpdateReverseEtlModel200Response UpdateReverseEtlModel(ctx, modelId).UpdateReverseEtlModelInput(updateReverseEtlModelInput).Execute()

Update Reverse Etl Model



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
    modelId := "37YCmBXayzqG4sit63P5pH" // string | 
    updateReverseEtlModelInput := *api.NewUpdateReverseEtlModelInput() // UpdateReverseEtlModelInput | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.ReverseETLAPI.UpdateReverseEtlModel(ctx).UpdateReverseEtlModelInput(updateReverseEtlModelInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReverseETLAPI.UpdateReverseEtlModel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `UpdateReverseEtlModel`: UpdateReverseEtlModel200Response
    fmt.Fprintf(os.Stdout, "Response from `ReverseETLAPI.UpdateReverseEtlModel`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**modelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateReverseEtlModelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateReverseEtlModelInput** | [**UpdateReverseEtlModelInput**](UpdateReverseEtlModelInput.md) |  | 

### Return type

[**UpdateReverseEtlModel200Response**](UpdateReverseEtlModel200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/vnd.segment.v1alpha+json
- **Accept**: application/vnd.segment.v1alpha+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

