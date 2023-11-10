# \DeletionAndSuppressionAPI

All URIs are relative to *https://api.segmentapis.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCloudSourceRegulation**](DeletionAndSuppressionAPI.md#CreateCloudSourceRegulation) | **Post** /regulations/cloudsources/{sourceId} | Create Cloud Source Regulation
[**CreateSourceRegulation**](DeletionAndSuppressionAPI.md#CreateSourceRegulation) | **Post** /regulations/sources/{sourceId} | Create Source Regulation
[**CreateWorkspaceRegulation**](DeletionAndSuppressionAPI.md#CreateWorkspaceRegulation) | **Post** /regulations | Create Workspace Regulation
[**DeleteRegulation**](DeletionAndSuppressionAPI.md#DeleteRegulation) | **Delete** /regulations/{regulateId} | Delete Regulation
[**GetRegulation**](DeletionAndSuppressionAPI.md#GetRegulation) | **Get** /regulations/{regulateId} | Get Regulation
[**ListRegulationsFromSource**](DeletionAndSuppressionAPI.md#ListRegulationsFromSource) | **Get** /regulations/sources/{sourceId} | List Regulations from Source
[**ListSuppressions**](DeletionAndSuppressionAPI.md#ListSuppressions) | **Get** /suppressions | List Suppressions
[**ListWorkspaceRegulations**](DeletionAndSuppressionAPI.md#ListWorkspaceRegulations) | **Get** /regulations | List Workspace Regulations



## Operation: CreateCloudSourceRegulation

> CreateCloudSourceRegulation200Response CreateCloudSourceRegulation(ctx, sourceId).CreateCloudSourceRegulationV1Input(createCloudSourceRegulationV1Input).Execute()

Create Cloud Source Regulation



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
    createCloudSourceRegulationV1Input := *api.NewCreateCloudSourceRegulationV1Input("RegulationType_example", "SubjectType_example", []string{"SubjectIds_example"}, "Collection_example") // CreateCloudSourceRegulationV1Input | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "INSERT BEARER TOKEN HERE"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.DeletionAndSuppressionAPI.CreateCloudSourceRegulation(ctx, sourceId).CreateCloudSourceRegulationV1Input(createCloudSourceRegulationV1Input).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeletionAndSuppressionAPI.CreateCloudSourceRegulation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `CreateCloudSourceRegulation`: CreateCloudSourceRegulation200Response
    fmt.Fprintf(os.Stdout, "Response from `DeletionAndSuppressionAPI.CreateCloudSourceRegulation`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCloudSourceRegulationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createCloudSourceRegulationV1Input** | [**CreateCloudSourceRegulationV1Input**](CreateCloudSourceRegulationV1Input.md) |  | 

### Return type

[**CreateCloudSourceRegulation200Response**](CreateCloudSourceRegulation200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/vnd.segment.v1+json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json
- **Accept**: application/vnd.segment.v1+json, application/json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: CreateSourceRegulation

> CreateSourceRegulation200Response CreateSourceRegulation(ctx, sourceId).CreateSourceRegulationV1Input(createSourceRegulationV1Input).Execute()

Create Source Regulation



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
    createSourceRegulationV1Input := *api.NewCreateSourceRegulationV1Input("RegulationType_example", "SubjectType_example", []string{"SubjectIds_example"}) // CreateSourceRegulationV1Input | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "INSERT BEARER TOKEN HERE"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.DeletionAndSuppressionAPI.CreateSourceRegulation(ctx, sourceId).CreateSourceRegulationV1Input(createSourceRegulationV1Input).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeletionAndSuppressionAPI.CreateSourceRegulation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `CreateSourceRegulation`: CreateSourceRegulation200Response
    fmt.Fprintf(os.Stdout, "Response from `DeletionAndSuppressionAPI.CreateSourceRegulation`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSourceRegulationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createSourceRegulationV1Input** | [**CreateSourceRegulationV1Input**](CreateSourceRegulationV1Input.md) |  | 

### Return type

[**CreateSourceRegulation200Response**](CreateSourceRegulation200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/vnd.segment.v1+json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json
- **Accept**: application/vnd.segment.v1+json, application/json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: CreateWorkspaceRegulation

> CreateWorkspaceRegulation200Response CreateWorkspaceRegulation(ctx).CreateWorkspaceRegulationV1Input(createWorkspaceRegulationV1Input).Execute()

Create Workspace Regulation



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
    createWorkspaceRegulationV1Input := *api.NewCreateWorkspaceRegulationV1Input("RegulationType_example", "SubjectType_example", []string{"SubjectIds_example"}) // CreateWorkspaceRegulationV1Input | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "INSERT BEARER TOKEN HERE"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.DeletionAndSuppressionAPI.CreateWorkspaceRegulation(ctx).CreateWorkspaceRegulationV1Input(createWorkspaceRegulationV1Input).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeletionAndSuppressionAPI.CreateWorkspaceRegulation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `CreateWorkspaceRegulation`: CreateWorkspaceRegulation200Response
    fmt.Fprintf(os.Stdout, "Response from `DeletionAndSuppressionAPI.CreateWorkspaceRegulation`: %v\n", resp.GetData())
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorkspaceRegulationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createWorkspaceRegulationV1Input** | [**CreateWorkspaceRegulationV1Input**](CreateWorkspaceRegulationV1Input.md) |  | 

### Return type

[**CreateWorkspaceRegulation200Response**](CreateWorkspaceRegulation200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/vnd.segment.v1+json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json
- **Accept**: application/vnd.segment.v1+json, application/json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: DeleteRegulation

> DeleteRegulation200Response DeleteRegulation(ctx, regulateId).Execute()

Delete Regulation



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
    regulateId := "1qJkfE1tpwvQcklImGksLN629wn" // string | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "INSERT BEARER TOKEN HERE"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.DeletionAndSuppressionAPI.DeleteRegulation(ctx, regulateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeletionAndSuppressionAPI.DeleteRegulation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `DeleteRegulation`: DeleteRegulation200Response
    fmt.Fprintf(os.Stdout, "Response from `DeletionAndSuppressionAPI.DeleteRegulation`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**regulateId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRegulationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteRegulation200Response**](DeleteRegulation200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1+json, application/json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: GetRegulation

> GetRegulation200Response GetRegulation(ctx, regulateId).Execute()

Get Regulation



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
    regulateId := "1qJkfE1tpwvQcklImGksLN629wn" // string | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "INSERT BEARER TOKEN HERE"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.DeletionAndSuppressionAPI.GetRegulation(ctx, regulateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeletionAndSuppressionAPI.GetRegulation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `GetRegulation`: GetRegulation200Response
    fmt.Fprintf(os.Stdout, "Response from `DeletionAndSuppressionAPI.GetRegulation`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**regulateId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRegulationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetRegulation200Response**](GetRegulation200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1+json, application/json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: ListRegulationsFromSource

> ListRegulationsFromSource200Response ListRegulationsFromSource(ctx, sourceId).Pagination(pagination).Status(status).RegulationTypes(regulationTypes).Execute()

List Regulations from Source



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
    pagination := *api.NewPaginationInput(float32(123)) // PaginationInput | Pagination parameters.  This parameter exists in v1.
    status := "status_example" // string | The status on which to filter returned regulations.  This parameter exists in v1. (optional)
    regulationTypes := []string{"RegulationTypes_example"} // []string | The regulation types on which to filter returned regulations.  This parameter exists in v1. (optional)

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "INSERT BEARER TOKEN HERE"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.DeletionAndSuppressionAPI.ListRegulationsFromSource(ctx, sourceId).Pagination(pagination).Status(status).RegulationTypes(regulationTypes).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeletionAndSuppressionAPI.ListRegulationsFromSource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `ListRegulationsFromSource`: ListRegulationsFromSource200Response
    fmt.Fprintf(os.Stdout, "Response from `DeletionAndSuppressionAPI.ListRegulationsFromSource`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRegulationsFromSourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pagination** | [**PaginationInput**](PaginationInput.md) | Pagination parameters.  This parameter exists in v1. | 
 **status** | **string** | The status on which to filter returned regulations.  This parameter exists in v1. | 
 **regulationTypes** | **[]string** | The regulation types on which to filter returned regulations.  This parameter exists in v1. | 

### Return type

[**ListRegulationsFromSource200Response**](ListRegulationsFromSource200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1+json, application/json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: ListSuppressions

> ListSuppressions200Response ListSuppressions(ctx).Pagination(pagination).Execute()

List Suppressions



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
    pagination := *api.NewPaginationInput(float32(123)) // PaginationInput | Pagination parameters.  This parameter exists in v1.

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "INSERT BEARER TOKEN HERE"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.DeletionAndSuppressionAPI.ListSuppressions(ctx).Pagination(pagination).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeletionAndSuppressionAPI.ListSuppressions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `ListSuppressions`: ListSuppressions200Response
    fmt.Fprintf(os.Stdout, "Response from `DeletionAndSuppressionAPI.ListSuppressions`: %v\n", resp.GetData())
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSuppressionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pagination** | [**PaginationInput**](PaginationInput.md) | Pagination parameters.  This parameter exists in v1. | 

### Return type

[**ListSuppressions200Response**](ListSuppressions200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1+json, application/json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: ListWorkspaceRegulations

> ListWorkspaceRegulations200Response ListWorkspaceRegulations(ctx).Pagination(pagination).Status(status).RegulationTypes(regulationTypes).Execute()

List Workspace Regulations



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
    pagination := *api.NewPaginationInput(float32(123)) // PaginationInput | Pagination parameters.  This parameter exists in v1.
    status := "status_example" // string | The status on which to filter the returned regulations.  This parameter exists in v1. (optional)
    regulationTypes := []string{"RegulationTypes_example"} // []string | The regulation types on which to filter returned regulations.  This parameter exists in v1. (optional)

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "INSERT BEARER TOKEN HERE"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.DeletionAndSuppressionAPI.ListWorkspaceRegulations(ctx).Pagination(pagination).Status(status).RegulationTypes(regulationTypes).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeletionAndSuppressionAPI.ListWorkspaceRegulations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `ListWorkspaceRegulations`: ListWorkspaceRegulations200Response
    fmt.Fprintf(os.Stdout, "Response from `DeletionAndSuppressionAPI.ListWorkspaceRegulations`: %v\n", resp.GetData())
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListWorkspaceRegulationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pagination** | [**PaginationInput**](PaginationInput.md) | Pagination parameters.  This parameter exists in v1. | 
 **status** | **string** | The status on which to filter the returned regulations.  This parameter exists in v1. | 
 **regulationTypes** | **[]string** | The regulation types on which to filter returned regulations.  This parameter exists in v1. | 

### Return type

[**ListWorkspaceRegulations200Response**](ListWorkspaceRegulations200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1+json, application/json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

