# \LivePluginsAPI

All URIs are relative to *https://api.segmentapis.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLivePlugin**](LivePluginsAPI.md#CreateLivePlugin) | **Post** /sources/{sourceId}/live-plugins/create | Create Live Plugin
[**DeleteLivePluginCode**](LivePluginsAPI.md#DeleteLivePluginCode) | **Delete** /sources/{sourceId}/live-plugins/delete-code | Delete Live Plugin Code
[**GetLatestFromLivePlugins**](LivePluginsAPI.md#GetLatestFromLivePlugins) | **Get** /sources/{sourceId}/live-plugins/latest | Get Latest from Live Plugins



## Operation: CreateLivePlugin

> CreateLivePlugin200Response CreateLivePlugin(ctx, sourceId).CreateLivePluginAlphaInput(createLivePluginAlphaInput).Execute()

Create Live Plugin



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
    createLivePluginAlphaInput := *api.NewCreateLivePluginAlphaInput("Code_example") // CreateLivePluginAlphaInput | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.LivePluginsAPI.CreateLivePlugin(ctx, sourceId).CreateLivePluginAlphaInput(createLivePluginAlphaInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LivePluginsAPI.CreateLivePlugin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `CreateLivePlugin`: CreateLivePlugin200Response
    fmt.Fprintf(os.Stdout, "Response from `LivePluginsAPI.CreateLivePlugin`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLivePluginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createLivePluginAlphaInput** | [**CreateLivePluginAlphaInput**](CreateLivePluginAlphaInput.md) |  | 

### Return type

[**CreateLivePlugin200Response**](CreateLivePlugin200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/vnd.segment.v1alpha+json
- **Accept**: application/vnd.segment.v1alpha+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: DeleteLivePluginCode

> DeleteLivePluginCode200Response DeleteLivePluginCode(ctx, sourceId).Execute()

Delete Live Plugin Code



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
    resp, r, err := apiClient.LivePluginsAPI.DeleteLivePluginCode(ctx, sourceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LivePluginsAPI.DeleteLivePluginCode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `DeleteLivePluginCode`: DeleteLivePluginCode200Response
    fmt.Fprintf(os.Stdout, "Response from `LivePluginsAPI.DeleteLivePluginCode`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLivePluginCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteLivePluginCode200Response**](DeleteLivePluginCode200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1alpha+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: GetLatestFromLivePlugins

> GetLatestFromLivePlugins200Response GetLatestFromLivePlugins(ctx, sourceId).Execute()

Get Latest from Live Plugins



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
    resp, r, err := apiClient.LivePluginsAPI.GetLatestFromLivePlugins(ctx, sourceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LivePluginsAPI.GetLatestFromLivePlugins``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `GetLatestFromLivePlugins`: GetLatestFromLivePlugins200Response
    fmt.Fprintf(os.Stdout, "Response from `LivePluginsAPI.GetLatestFromLivePlugins`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLatestFromLivePluginsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetLatestFromLivePlugins200Response**](GetLatestFromLivePlugins200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1alpha+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

