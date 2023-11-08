# \EdgeFunctionsAPI

All URIs are relative to *https://api.segmentapis.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEdgeFunctions**](EdgeFunctionsAPI.md#CreateEdgeFunctions) | **Post** /sources/{sourceId}/edge-functions | Create Edge Functions
[**DisableEdgeFunctions**](EdgeFunctionsAPI.md#DisableEdgeFunctions) | **Patch** /sources/{sourceId}/edge-functions/disable | Disable Edge Functions
[**GenerateUploadURLForEdgeFunctions**](EdgeFunctionsAPI.md#GenerateUploadURLForEdgeFunctions) | **Post** /sources/{sourceId}/edge-functions/upload-url | Generate Upload URL for Edge Functions
[**GetLatestFromEdgeFunctions**](EdgeFunctionsAPI.md#GetLatestFromEdgeFunctions) | **Get** /sources/{sourceId}/edge-functions/latest | Get Latest from Edge Functions



## Operation: CreateEdgeFunctions

> CreateEdgeFunctions200Response CreateEdgeFunctions(ctx, sourceId).CreateEdgeFunctionsAlphaInput(createEdgeFunctionsAlphaInput).Execute()

Create Edge Functions



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
    createEdgeFunctionsAlphaInput := *api.NewCreateEdgeFunctionsAlphaInput("UploadURL_example") // CreateEdgeFunctionsAlphaInput | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.EdgeFunctionsAPI.CreateEdgeFunctions(ctx).CreateEdgeFunctionsAlphaInput(createEdgeFunctionsAlphaInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EdgeFunctionsAPI.CreateEdgeFunctions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `CreateEdgeFunctions`: CreateEdgeFunctions200Response
    fmt.Fprintf(os.Stdout, "Response from `EdgeFunctionsAPI.CreateEdgeFunctions`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateEdgeFunctionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createEdgeFunctionsAlphaInput** | [**CreateEdgeFunctionsAlphaInput**](CreateEdgeFunctionsAlphaInput.md) |  | 

### Return type

[**CreateEdgeFunctions200Response**](CreateEdgeFunctions200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/vnd.segment.v1alpha+json
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
    resp, r, err := apiClient.EdgeFunctionsAPI.DisableEdgeFunctions(ctx).Execute()
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
    resp, r, err := apiClient.EdgeFunctionsAPI.GenerateUploadURLForEdgeFunctions(ctx).Execute()
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
    resp, r, err := apiClient.EdgeFunctionsAPI.GetLatestFromEdgeFunctions(ctx).Execute()
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

