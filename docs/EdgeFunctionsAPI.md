# \EdgeFunctionsAPI

All URIs are relative to *https://api.segmentapis.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEdgeFunction**](EdgeFunctionsAPI.md#CreateEdgeFunction) | **Post** /sources/{sourceId}/edge-functions/create | Create Edge Function
[**DeleteEdgeFunctionCode**](EdgeFunctionsAPI.md#DeleteEdgeFunctionCode) | **Delete** /sources/{sourceId}/edge-functions/delete-code | Delete Edge Function Code
[**DisableEdgeFunctions**](EdgeFunctionsAPI.md#DisableEdgeFunctions) | **Patch** /sources/{sourceId}/edge-functions/disable | Disable Edge Functions
[**GenerateUploadURLForEdgeFunctions**](EdgeFunctionsAPI.md#GenerateUploadURLForEdgeFunctions) | **Post** /sources/{sourceId}/edge-functions/upload-url | Generate Upload URL for Edge Functions
[**GetLatestFromEdgeFunctions**](EdgeFunctionsAPI.md#GetLatestFromEdgeFunctions) | **Get** /sources/{sourceId}/edge-functions/latest | Get Latest from Edge Functions



## Operation: CreateEdgeFunction

> CreateEdgeFunction200Response CreateEdgeFunction(ctx, sourceId).CreateEdgeFunctionAlphaInput(createEdgeFunctionAlphaInput).Execute()

Create Edge Function



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
    sourceId := "qQEHquLrjRDN9j1ByrChyn" // string | 
    createEdgeFunctionAlphaInput := *api.NewCreateEdgeFunctionAlphaInput("Code_example") // CreateEdgeFunctionAlphaInput | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.EdgeFunctionsAPI.CreateEdgeFunction(ctx, sourceId).CreateEdgeFunctionAlphaInput(createEdgeFunctionAlphaInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EdgeFunctionsAPI.CreateEdgeFunction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `CreateEdgeFunction`: CreateEdgeFunction200Response
    fmt.Fprintf(os.Stdout, "Response from `EdgeFunctionsAPI.CreateEdgeFunction`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateEdgeFunctionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createEdgeFunctionAlphaInput** | [**CreateEdgeFunctionAlphaInput**](CreateEdgeFunctionAlphaInput.md) |  | 

### Return type

[**CreateEdgeFunction200Response**](CreateEdgeFunction200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/vnd.segment.v1alpha+json
- **Accept**: application/vnd.segment.v1alpha+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: DeleteEdgeFunctionCode

> DeleteEdgeFunctionCode200Response DeleteEdgeFunctionCode(ctx, sourceId).Execute()

Delete Edge Function Code



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
    sourceId := "qQEHquLrjRDN9j1ByrChyn" // string | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.EdgeFunctionsAPI.DeleteEdgeFunctionCode(ctx, sourceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EdgeFunctionsAPI.DeleteEdgeFunctionCode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `DeleteEdgeFunctionCode`: DeleteEdgeFunctionCode200Response
    fmt.Fprintf(os.Stdout, "Response from `EdgeFunctionsAPI.DeleteEdgeFunctionCode`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEdgeFunctionCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteEdgeFunctionCode200Response**](DeleteEdgeFunctionCode200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1alpha+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: DisableEdgeFunctions

> DisableEdgeFunctions200Response DisableEdgeFunctions(ctx, sourceId).Execute()

Disable Edge Functions



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
    sourceId := "qQEHquLrjRDN9j1ByrChyn" // string | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.EdgeFunctionsAPI.DisableEdgeFunctions(ctx, sourceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EdgeFunctionsAPI.DisableEdgeFunctions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `DisableEdgeFunctions`: DisableEdgeFunctions200Response
    fmt.Fprintf(os.Stdout, "Response from `EdgeFunctionsAPI.DisableEdgeFunctions`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisableEdgeFunctionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DisableEdgeFunctions200Response**](DisableEdgeFunctions200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1alpha+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: GenerateUploadURLForEdgeFunctions

> GenerateUploadURLForEdgeFunctions200Response GenerateUploadURLForEdgeFunctions(ctx, sourceId).Execute()

Generate Upload URL for Edge Functions



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
    sourceId := "qQEHquLrjRDN9j1ByrChyn" // string | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.EdgeFunctionsAPI.GenerateUploadURLForEdgeFunctions(ctx, sourceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EdgeFunctionsAPI.GenerateUploadURLForEdgeFunctions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `GenerateUploadURLForEdgeFunctions`: GenerateUploadURLForEdgeFunctions200Response
    fmt.Fprintf(os.Stdout, "Response from `EdgeFunctionsAPI.GenerateUploadURLForEdgeFunctions`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGenerateUploadURLForEdgeFunctionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GenerateUploadURLForEdgeFunctions200Response**](GenerateUploadURLForEdgeFunctions200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1alpha+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: GetLatestFromEdgeFunctions

> GetLatestFromEdgeFunctions200Response GetLatestFromEdgeFunctions(ctx, sourceId).Execute()

Get Latest from Edge Functions



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
    sourceId := "qQEHquLrjRDN9j1ByrChyn" // string | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.EdgeFunctionsAPI.GetLatestFromEdgeFunctions(ctx, sourceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EdgeFunctionsAPI.GetLatestFromEdgeFunctions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `GetLatestFromEdgeFunctions`: GetLatestFromEdgeFunctions200Response
    fmt.Fprintf(os.Stdout, "Response from `EdgeFunctionsAPI.GetLatestFromEdgeFunctions`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLatestFromEdgeFunctionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetLatestFromEdgeFunctions200Response**](GetLatestFromEdgeFunctions200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1alpha+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

