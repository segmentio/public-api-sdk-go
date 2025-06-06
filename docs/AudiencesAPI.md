# \AudiencesAPI

All URIs are relative to *https://api.segmentapis.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAudience**](AudiencesAPI.md#CreateAudience) | **Post** /spaces/{spaceId}/audiences | Create Audience
[**GetAudience**](AudiencesAPI.md#GetAudience) | **Get** /spaces/{spaceId}/audiences/{id} | Get Audience
[**ListAudienceConsumersFromSpaceAndAudience**](AudiencesAPI.md#ListAudienceConsumersFromSpaceAndAudience) | **Get** /spaces/{spaceId}/audiences/{id}/audience-references | List Audience Consumers from Space And Audience
[**ListAudiences**](AudiencesAPI.md#ListAudiences) | **Get** /spaces/{spaceId}/audiences | List Audiences
[**PreviewAudience**](AudiencesAPI.md#PreviewAudience) | **Post** /spaces/{spaceId}/audiences/previews | Preview Audience
[**RemoveAudienceFromSpace**](AudiencesAPI.md#RemoveAudienceFromSpace) | **Delete** /spaces/{spaceId}/audiences/{id} | Remove Audience from Space
[**UpdateAudienceForSpace**](AudiencesAPI.md#UpdateAudienceForSpace) | **Patch** /spaces/{spaceId}/audiences/{id} | Update Audience for Space



## Operation: CreateAudience

> CreateAudience200Response CreateAudience(ctx, spaceId).CreateAudienceAlphaInput(createAudienceAlphaInput).Execute()

Create Audience



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
    spaceId := "9aQ1Lj62S4bomZKLF4DPqW" // string | 
    createAudienceAlphaInput := *api.NewCreateAudienceAlphaInput("Name_example", *api.NewAudienceDefinition("Type_example", "Query_example")) // CreateAudienceAlphaInput | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.AudiencesAPI.CreateAudience(ctx, spaceId).CreateAudienceAlphaInput(createAudienceAlphaInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AudiencesAPI.CreateAudience``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `CreateAudience`: CreateAudience200Response
    fmt.Fprintf(os.Stdout, "Response from `AudiencesAPI.CreateAudience`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAudienceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createAudienceAlphaInput** | [**CreateAudienceAlphaInput**](CreateAudienceAlphaInput.md) |  | 

### Return type

[**CreateAudience200Response**](CreateAudience200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/vnd.segment.v1alpha+json
- **Accept**: application/vnd.segment.v1alpha+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: GetAudience

> GetAudience200Response GetAudience(ctx, spaceId, id).Execute()

Get Audience



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
    spaceId := "9aQ1Lj62S4bomZKLF4DPqW" // string | 
    id := "id" // string | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.AudiencesAPI.GetAudience(ctx, spaceId, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AudiencesAPI.GetAudience``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `GetAudience`: GetAudience200Response
    fmt.Fprintf(os.Stdout, "Response from `AudiencesAPI.GetAudience`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAudienceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetAudience200Response**](GetAudience200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: ListAudienceConsumersFromSpaceAndAudience

> ListAudienceConsumersFromSpaceAndAudience200Response ListAudienceConsumersFromSpaceAndAudience(ctx, spaceId, id).Pagination(pagination).Search(search).Sort(sort).Execute()

List Audience Consumers from Space And Audience



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
    spaceId := "9aQ1Lj62S4bomZKLF4DPqW" // string | 
    id := "aud_0ujsswThIGTUYm2K8FjOOfXtY1K" // string | 
    pagination := *api.NewPaginationInput(10) // PaginationInput | Information about the pagination of this response.  [See pagination](https://docs.segmentapis.com/tag/Pagination/#section/Pagination-parameters) for more info.  This parameter exists in alpha. (optional)
    search := *api.NewListAudienceConsumersSearchInput("Type_example", "Query_example") // ListAudienceConsumersSearchInput | Optional search criteria  This parameter exists in alpha. (optional)
    sort := *api.NewListAudienceConsumersSortInput("Field_example", "Direction_example") // ListAudienceConsumersSortInput | Optional sort criteria  This parameter exists in alpha. (optional)

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.AudiencesAPI.ListAudienceConsumersFromSpaceAndAudience(ctx, spaceId, id).Pagination(pagination).Search(search).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AudiencesAPI.ListAudienceConsumersFromSpaceAndAudience``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `ListAudienceConsumersFromSpaceAndAudience`: ListAudienceConsumersFromSpaceAndAudience200Response
    fmt.Fprintf(os.Stdout, "Response from `AudiencesAPI.ListAudienceConsumersFromSpaceAndAudience`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAudienceConsumersFromSpaceAndAudienceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pagination** | [**PaginationInput**](PaginationInput.md) | Information about the pagination of this response.  [See pagination](https://docs.segmentapis.com/tag/Pagination/#section/Pagination-parameters) for more info.  This parameter exists in alpha. | 
 **search** | [**ListAudienceConsumersSearchInput**](ListAudienceConsumersSearchInput.md) | Optional search criteria  This parameter exists in alpha. | 
 **sort** | [**ListAudienceConsumersSortInput**](ListAudienceConsumersSortInput.md) | Optional sort criteria  This parameter exists in alpha. | 

### Return type

[**ListAudienceConsumersFromSpaceAndAudience200Response**](ListAudienceConsumersFromSpaceAndAudience200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1alpha+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: ListAudiences

> ListAudiences200Response ListAudiences(ctx, spaceId).Pagination(pagination).Execute()

List Audiences



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
    spaceId := "9aQ1Lj62S4bomZKLF4DPqW" // string | 
    pagination := *api.NewListAudiencesPaginationInput(float32(123)) // ListAudiencesPaginationInput | Information about the pagination of this response.  [See pagination](https://docs.segmentapis.com/tag/Pagination/#section/Pagination-parameters) for more info.  This parameter exists in alpha. (optional)

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.AudiencesAPI.ListAudiences(ctx, spaceId).Pagination(pagination).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AudiencesAPI.ListAudiences``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `ListAudiences`: ListAudiences200Response
    fmt.Fprintf(os.Stdout, "Response from `AudiencesAPI.ListAudiences`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAudiencesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pagination** | [**ListAudiencesPaginationInput**](ListAudiencesPaginationInput.md) | Information about the pagination of this response.  [See pagination](https://docs.segmentapis.com/tag/Pagination/#section/Pagination-parameters) for more info.  This parameter exists in alpha. | 

### Return type

[**ListAudiences200Response**](ListAudiences200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: PreviewAudience

> PreviewAudience200Response PreviewAudience(ctx, spaceId).PreviewAudienceInput(previewAudienceInput).Execute()

Preview Audience



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
    spaceId := "9aQ1Lj62S4bomZKLF4DPqW" // string | 
    previewAudienceInput := *api.NewPreviewAudienceInput(*api.NewAudienceDefinitionWithoutType("Query_example"), "AudienceType_example") // PreviewAudienceInput | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.AudiencesAPI.PreviewAudience(ctx, spaceId).PreviewAudienceInput(previewAudienceInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AudiencesAPI.PreviewAudience``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `PreviewAudience`: PreviewAudience200Response
    fmt.Fprintf(os.Stdout, "Response from `AudiencesAPI.PreviewAudience`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPreviewAudienceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **previewAudienceInput** | [**PreviewAudienceInput**](PreviewAudienceInput.md) |  | 

### Return type

[**PreviewAudience200Response**](PreviewAudience200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/vnd.segment.v1alpha+json
- **Accept**: application/vnd.segment.v1alpha+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: RemoveAudienceFromSpace

> RemoveAudienceFromSpace200Response RemoveAudienceFromSpace(ctx, spaceId, id).Execute()

Remove Audience from Space



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
    spaceId := "9aQ1Lj62S4bomZKLF4DPqW" // string | 
    id := "aud_0ujsswThIGTUYm2K8FjOOfXtY1K" // string | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.AudiencesAPI.RemoveAudienceFromSpace(ctx, spaceId, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AudiencesAPI.RemoveAudienceFromSpace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `RemoveAudienceFromSpace`: RemoveAudienceFromSpace200Response
    fmt.Fprintf(os.Stdout, "Response from `AudiencesAPI.RemoveAudienceFromSpace`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveAudienceFromSpaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RemoveAudienceFromSpace200Response**](RemoveAudienceFromSpace200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1alpha+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: UpdateAudienceForSpace

> UpdateAudienceForSpace200Response UpdateAudienceForSpace(ctx, spaceId, id).UpdateAudienceForSpaceAlphaInput(updateAudienceForSpaceAlphaInput).Execute()

Update Audience for Space



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
    spaceId := "9aQ1Lj62S4bomZKLF4DPqW" // string | 
    id := "aud_0ujsswThIGTUYm2K8FjOOfXtY1K" // string | 
    updateAudienceForSpaceAlphaInput := *api.NewUpdateAudienceForSpaceAlphaInput() // UpdateAudienceForSpaceAlphaInput | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.AudiencesAPI.UpdateAudienceForSpace(ctx, spaceId, id).UpdateAudienceForSpaceAlphaInput(updateAudienceForSpaceAlphaInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AudiencesAPI.UpdateAudienceForSpace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `UpdateAudienceForSpace`: UpdateAudienceForSpace200Response
    fmt.Fprintf(os.Stdout, "Response from `AudiencesAPI.UpdateAudienceForSpace`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAudienceForSpaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateAudienceForSpaceAlphaInput** | [**UpdateAudienceForSpaceAlphaInput**](UpdateAudienceForSpaceAlphaInput.md) |  | 

### Return type

[**UpdateAudienceForSpace200Response**](UpdateAudienceForSpace200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/vnd.segment.v1alpha+json
- **Accept**: application/vnd.segment.v1alpha+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

