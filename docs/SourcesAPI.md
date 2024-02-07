# \SourcesAPI

All URIs are relative to *https://api.segmentapis.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddLabelsToSource**](SourcesAPI.md#AddLabelsToSource) | **Post** /sources/{sourceId}/labels | Add Labels to Source
[**CreateSource**](SourcesAPI.md#CreateSource) | **Post** /sources | Create Source
[**DeleteSource**](SourcesAPI.md#DeleteSource) | **Delete** /sources/{sourceId} | Delete Source
[**GetSource**](SourcesAPI.md#GetSource) | **Get** /sources/{sourceId} | Get Source
[**ListConnectedDestinationsFromSource**](SourcesAPI.md#ListConnectedDestinationsFromSource) | **Get** /sources/{sourceId}/connected-destinations | List Connected Destinations from Source
[**ListConnectedWarehousesFromSource**](SourcesAPI.md#ListConnectedWarehousesFromSource) | **Get** /sources/{sourceId}/connected-warehouses | List Connected Warehouses from Source
[**ListSchemaSettingsInSource**](SourcesAPI.md#ListSchemaSettingsInSource) | **Get** /sources/{sourceId}/settings | List Schema Settings in Source
[**ListSources**](SourcesAPI.md#ListSources) | **Get** /sources | List Sources
[**ReplaceLabelsInSource**](SourcesAPI.md#ReplaceLabelsInSource) | **Put** /sources/{sourceId}/labels | Replace Labels in Source
[**UpdateSchemaSettingsInSource**](SourcesAPI.md#UpdateSchemaSettingsInSource) | **Patch** /sources/{sourceId}/settings | Update Schema Settings in Source
[**UpdateSource**](SourcesAPI.md#UpdateSource) | **Patch** /sources/{sourceId} | Update Source



## Operation: AddLabelsToSource

> AddLabelsToSource200Response AddLabelsToSource(ctx, sourceId).AddLabelsToSourceV1Input(addLabelsToSourceV1Input).Execute()

Add Labels to Source



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
    sourceId := "rh5BDZp6QDHvXFCkibm1pR" // string | 
    addLabelsToSourceV1Input := *api.NewAddLabelsToSourceV1Input([]api.LabelV1{*api.NewLabelV1("Key_example", "Value_example")}) // AddLabelsToSourceV1Input | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.SourcesAPI.AddLabelsToSource(ctx, sourceId).AddLabelsToSourceV1Input(addLabelsToSourceV1Input).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SourcesAPI.AddLabelsToSource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `AddLabelsToSource`: AddLabelsToSource200Response
    fmt.Fprintf(os.Stdout, "Response from `SourcesAPI.AddLabelsToSource`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddLabelsToSourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addLabelsToSourceV1Input** | [**AddLabelsToSourceV1Input**](AddLabelsToSourceV1Input.md) |  | 

### Return type

[**AddLabelsToSource200Response**](AddLabelsToSource200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json, application/vnd.segment.v1+json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json
- **Accept**: application/vnd.segment.v1+json, application/json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: CreateSource

> CreateSource201Response CreateSource(ctx).CreateSourceV1Input(createSourceV1Input).Execute()

Create Source



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
    createSourceV1Input := *api.NewCreateSourceV1Input("Slug_example", false, "MetadataId_example") // CreateSourceV1Input | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.SourcesAPI.CreateSource(ctx).CreateSourceV1Input(createSourceV1Input).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SourcesAPI.CreateSource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `CreateSource`: CreateSource201Response
    fmt.Fprintf(os.Stdout, "Response from `SourcesAPI.CreateSource`: %v\n", resp.GetData())
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createSourceV1Input** | [**CreateSourceV1Input**](CreateSourceV1Input.md) |  | 

### Return type

[**CreateSource201Response**](CreateSource201Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json, application/vnd.segment.v1+json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json
- **Accept**: application/vnd.segment.v1+json, application/json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: DeleteSource

> DeleteSource200Response DeleteSource(ctx, sourceId).Execute()

Delete Source



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
    sourceId := "rYxTjyaPtAELCjnFE5EYfM" // string | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.SourcesAPI.DeleteSource(ctx, sourceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SourcesAPI.DeleteSource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `DeleteSource`: DeleteSource200Response
    fmt.Fprintf(os.Stdout, "Response from `SourcesAPI.DeleteSource`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteSource200Response**](DeleteSource200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1+json, application/json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: GetSource

> GetSource200Response GetSource(ctx, sourceId).Execute()

Get Source



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
    resp, r, err := apiClient.SourcesAPI.GetSource(ctx, sourceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SourcesAPI.GetSource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `GetSource`: GetSource200Response
    fmt.Fprintf(os.Stdout, "Response from `SourcesAPI.GetSource`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetSource200Response**](GetSource200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1+json, application/json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: ListConnectedDestinationsFromSource

> ListConnectedDestinationsFromSource200Response ListConnectedDestinationsFromSource(ctx, sourceId).Pagination(pagination).Execute()

List Connected Destinations from Source



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
    pagination := *api.NewPaginationInput(10) // PaginationInput | Required pagination params for the request.  This parameter exists in alpha.

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.SourcesAPI.ListConnectedDestinationsFromSource(ctx, sourceId).Pagination(pagination).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SourcesAPI.ListConnectedDestinationsFromSource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `ListConnectedDestinationsFromSource`: ListConnectedDestinationsFromSource200Response
    fmt.Fprintf(os.Stdout, "Response from `SourcesAPI.ListConnectedDestinationsFromSource`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListConnectedDestinationsFromSourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pagination** | [**PaginationInput**](PaginationInput.md) | Required pagination params for the request.  This parameter exists in alpha. | 

### Return type

[**ListConnectedDestinationsFromSource200Response**](ListConnectedDestinationsFromSource200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1+json, application/json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: ListConnectedWarehousesFromSource

> ListConnectedWarehousesFromSource200Response ListConnectedWarehousesFromSource(ctx, sourceId).Pagination(pagination).Execute()

List Connected Warehouses from Source



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
    pagination := *api.NewPaginationInput(10) // PaginationInput | Required pagination params for the request.  This parameter exists in alpha.

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.SourcesAPI.ListConnectedWarehousesFromSource(ctx, sourceId).Pagination(pagination).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SourcesAPI.ListConnectedWarehousesFromSource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `ListConnectedWarehousesFromSource`: ListConnectedWarehousesFromSource200Response
    fmt.Fprintf(os.Stdout, "Response from `SourcesAPI.ListConnectedWarehousesFromSource`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListConnectedWarehousesFromSourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pagination** | [**PaginationInput**](PaginationInput.md) | Required pagination params for the request.  This parameter exists in alpha. | 

### Return type

[**ListConnectedWarehousesFromSource200Response**](ListConnectedWarehousesFromSource200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1+json, application/json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: ListSchemaSettingsInSource

> ListSchemaSettingsInSource200Response ListSchemaSettingsInSource(ctx, sourceId).Execute()

List Schema Settings in Source



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
    resp, r, err := apiClient.SourcesAPI.ListSchemaSettingsInSource(ctx, sourceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SourcesAPI.ListSchemaSettingsInSource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `ListSchemaSettingsInSource`: ListSchemaSettingsInSource200Response
    fmt.Fprintf(os.Stdout, "Response from `SourcesAPI.ListSchemaSettingsInSource`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSchemaSettingsInSourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListSchemaSettingsInSource200Response**](ListSchemaSettingsInSource200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1+json, application/json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: ListSources

> ListSources200Response ListSources(ctx).Pagination(pagination).Execute()

List Sources



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
    pagination := *api.NewPaginationInput(10) // PaginationInput | Defines the pagination parameters.  This parameter exists in alpha.

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.SourcesAPI.ListSources(ctx).Pagination(pagination).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SourcesAPI.ListSources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `ListSources`: ListSources200Response
    fmt.Fprintf(os.Stdout, "Response from `SourcesAPI.ListSources`: %v\n", resp.GetData())
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pagination** | [**PaginationInput**](PaginationInput.md) | Defines the pagination parameters.  This parameter exists in alpha. | 

### Return type

[**ListSources200Response**](ListSources200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1+json, application/json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: ReplaceLabelsInSource

> ReplaceLabelsInSource200Response ReplaceLabelsInSource(ctx, sourceId).ReplaceLabelsInSourceV1Input(replaceLabelsInSourceV1Input).Execute()

Replace Labels in Source



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
    sourceId := "rh5BDZp6QDHvXFCkibm1pR" // string | 
    replaceLabelsInSourceV1Input := *api.NewReplaceLabelsInSourceV1Input([]api.LabelV1{*api.NewLabelV1("Key_example", "Value_example")}) // ReplaceLabelsInSourceV1Input | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.SourcesAPI.ReplaceLabelsInSource(ctx, sourceId).ReplaceLabelsInSourceV1Input(replaceLabelsInSourceV1Input).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SourcesAPI.ReplaceLabelsInSource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `ReplaceLabelsInSource`: ReplaceLabelsInSource200Response
    fmt.Fprintf(os.Stdout, "Response from `SourcesAPI.ReplaceLabelsInSource`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceLabelsInSourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **replaceLabelsInSourceV1Input** | [**ReplaceLabelsInSourceV1Input**](ReplaceLabelsInSourceV1Input.md) |  | 

### Return type

[**ReplaceLabelsInSource200Response**](ReplaceLabelsInSource200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json, application/vnd.segment.v1+json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json
- **Accept**: application/vnd.segment.v1+json, application/json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: UpdateSchemaSettingsInSource

> UpdateSchemaSettingsInSource200Response UpdateSchemaSettingsInSource(ctx, sourceId).UpdateSchemaSettingsInSourceV1Input(updateSchemaSettingsInSourceV1Input).Execute()

Update Schema Settings in Source



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
    updateSchemaSettingsInSourceV1Input := *api.NewUpdateSchemaSettingsInSourceV1Input() // UpdateSchemaSettingsInSourceV1Input | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.SourcesAPI.UpdateSchemaSettingsInSource(ctx, sourceId).UpdateSchemaSettingsInSourceV1Input(updateSchemaSettingsInSourceV1Input).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SourcesAPI.UpdateSchemaSettingsInSource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `UpdateSchemaSettingsInSource`: UpdateSchemaSettingsInSource200Response
    fmt.Fprintf(os.Stdout, "Response from `SourcesAPI.UpdateSchemaSettingsInSource`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSchemaSettingsInSourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateSchemaSettingsInSourceV1Input** | [**UpdateSchemaSettingsInSourceV1Input**](UpdateSchemaSettingsInSourceV1Input.md) |  | 

### Return type

[**UpdateSchemaSettingsInSource200Response**](UpdateSchemaSettingsInSource200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json, application/vnd.segment.v1+json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json
- **Accept**: application/vnd.segment.v1+json, application/json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: UpdateSource

> UpdateSource200Response UpdateSource(ctx, sourceId).UpdateSourceV1Input(updateSourceV1Input).Execute()

Update Source



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
    sourceId := "87jXXk8QSLFPfMDGa6mtS3" // string | 
    updateSourceV1Input := *api.NewUpdateSourceV1Input() // UpdateSourceV1Input | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.SourcesAPI.UpdateSource(ctx, sourceId).UpdateSourceV1Input(updateSourceV1Input).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SourcesAPI.UpdateSource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `UpdateSource`: UpdateSource200Response
    fmt.Fprintf(os.Stdout, "Response from `SourcesAPI.UpdateSource`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateSourceV1Input** | [**UpdateSourceV1Input**](UpdateSourceV1Input.md) |  | 

### Return type

[**UpdateSource200Response**](UpdateSource200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json, application/vnd.segment.v1+json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json
- **Accept**: application/vnd.segment.v1+json, application/json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

