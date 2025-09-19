# \AudiencesAPI

All URIs are relative to *https://api.segmentapis.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddAudienceScheduleToAudience**](AudiencesAPI.md#AddAudienceScheduleToAudience) | **Post** /spaces/{spaceId}/audiences/{id}/schedules | Add Audience Schedule to Audience
[**CreateAudience**](AudiencesAPI.md#CreateAudience) | **Post** /spaces/{spaceId}/audiences | Create Audience
[**CreateAudiencePreview**](AudiencesAPI.md#CreateAudiencePreview) | **Post** /spaces/{spaceId}/audiences/previews | Create Audience Preview
[**GetAudience**](AudiencesAPI.md#GetAudience) | **Get** /spaces/{spaceId}/audiences/{id} | Get Audience
[**GetAudiencePreview**](AudiencesAPI.md#GetAudiencePreview) | **Get** /spaces/{spaceId}/audiences/previews/{id} | Get Audience Preview
[**GetAudienceScheduleFromSpaceAndAudience**](AudiencesAPI.md#GetAudienceScheduleFromSpaceAndAudience) | **Get** /spaces/{spaceId}/audiences/{id}/schedules/{scheduleId} | Get Audience Schedule from Space And Audience
[**ListAudienceConsumersFromSpaceAndAudience**](AudiencesAPI.md#ListAudienceConsumersFromSpaceAndAudience) | **Get** /spaces/{spaceId}/audiences/{id}/audience-references | List Audience Consumers from Space And Audience
[**ListAudienceSchedulesFromSpaceAndAudience**](AudiencesAPI.md#ListAudienceSchedulesFromSpaceAndAudience) | **Get** /spaces/{spaceId}/audiences/{id}/schedules | List Audience Schedules from Space And Audience
[**ListAudiences**](AudiencesAPI.md#ListAudiences) | **Get** /spaces/{spaceId}/audiences | List Audiences
[**RemoveAudienceFromSpace**](AudiencesAPI.md#RemoveAudienceFromSpace) | **Delete** /spaces/{spaceId}/audiences/{id} | Remove Audience from Space
[**UpdateAudienceForSpace**](AudiencesAPI.md#UpdateAudienceForSpace) | **Patch** /spaces/{spaceId}/audiences/{id} | Update Audience for Space



## Operation: AddAudienceScheduleToAudience

> AddAudienceScheduleToAudience200Response AddAudienceScheduleToAudience(ctx, spaceId, id).AddAudienceScheduleToAudienceAlphaInput(addAudienceScheduleToAudienceAlphaInput).Execute()

Add Audience Schedule to Audience



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
    id := "aud_0ujsszwN8NRY24YaXiTIE2VWDTS" // string | 
    addAudienceScheduleToAudienceAlphaInput := *api.NewAddAudienceScheduleToAudienceAlphaInput(false, "Strategy_example") // AddAudienceScheduleToAudienceAlphaInput | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.AudiencesAPI.AddAudienceScheduleToAudience(ctx, spaceId, id).AddAudienceScheduleToAudienceAlphaInput(addAudienceScheduleToAudienceAlphaInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AudiencesAPI.AddAudienceScheduleToAudience``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `AddAudienceScheduleToAudience`: AddAudienceScheduleToAudience200Response
    fmt.Fprintf(os.Stdout, "Response from `AudiencesAPI.AddAudienceScheduleToAudience`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddAudienceScheduleToAudienceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **addAudienceScheduleToAudienceAlphaInput** | [**AddAudienceScheduleToAudienceAlphaInput**](AddAudienceScheduleToAudienceAlphaInput.md) |  | 

### Return type

[**AddAudienceScheduleToAudience200Response**](AddAudienceScheduleToAudience200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/vnd.segment.v1alpha+json
- **Accept**: application/vnd.segment.v1alpha+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: CreateAudience

> CreateAudience200Response CreateAudience(ctx, spaceId).CreateAudienceBetaInput(createAudienceBetaInput).Execute()

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
    createAudienceBetaInput := *api.NewCreateAudienceBetaInput("Name_example", *api.NewAudienceDefinition("Query_example"), "AudienceType_example") // CreateAudienceBetaInput | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.AudiencesAPI.CreateAudience(ctx, spaceId).CreateAudienceBetaInput(createAudienceBetaInput).Execute()
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

 **createAudienceBetaInput** | [**CreateAudienceBetaInput**](CreateAudienceBetaInput.md) |  | 

### Return type

[**CreateAudience200Response**](CreateAudience200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json
- **Accept**: application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: CreateAudiencePreview

> CreateAudiencePreview200Response CreateAudiencePreview(ctx, spaceId).CreateAudiencePreviewBetaInput(createAudiencePreviewBetaInput).Execute()

Create Audience Preview



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
    createAudiencePreviewBetaInput := *api.NewCreateAudiencePreviewBetaInput(*api.NewAudienceDefinition("Query_example"), "AudienceType_example") // CreateAudiencePreviewBetaInput | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.AudiencesAPI.CreateAudiencePreview(ctx, spaceId).CreateAudiencePreviewBetaInput(createAudiencePreviewBetaInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AudiencesAPI.CreateAudiencePreview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `CreateAudiencePreview`: CreateAudiencePreview200Response
    fmt.Fprintf(os.Stdout, "Response from `AudiencesAPI.CreateAudiencePreview`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAudiencePreviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createAudiencePreviewBetaInput** | [**CreateAudiencePreviewBetaInput**](CreateAudiencePreviewBetaInput.md) |  | 

### Return type

[**CreateAudiencePreview200Response**](CreateAudiencePreview200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json
- **Accept**: application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: GetAudience

> GetAudience200Response GetAudience(ctx, spaceId, id).Include(include).Execute()

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
    id := "aud_0ujsszwN8NRY24YaXiTIE2VWDTS" // string | 
    include := "include_example" // string | Additional resource to include, support schedules only.  This parameter exists in alpha. (optional)

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.AudiencesAPI.GetAudience(ctx, spaceId, id).Include(include).Execute()
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


 **include** | **string** | Additional resource to include, support schedules only.  This parameter exists in alpha. | 

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


## Operation: GetAudiencePreview

> GetAudiencePreview200Response GetAudiencePreview(ctx, spaceId, id).Execute()

Get Audience Preview



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
    id := "31XR7mRsdw3vz3ZqvRa1623re0q-compute_preview_execution-9aQ1Lj62S4bomZKLF4DPqW-compute_preview_execution" // string | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.AudiencesAPI.GetAudiencePreview(ctx, spaceId, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AudiencesAPI.GetAudiencePreview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `GetAudiencePreview`: GetAudiencePreview200Response
    fmt.Fprintf(os.Stdout, "Response from `AudiencesAPI.GetAudiencePreview`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAudiencePreviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetAudiencePreview200Response**](GetAudiencePreview200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: GetAudienceScheduleFromSpaceAndAudience

> GetAudienceScheduleFromSpaceAndAudience200Response GetAudienceScheduleFromSpaceAndAudience(ctx, spaceId, id, scheduleId).Execute()

Get Audience Schedule from Space And Audience



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
    id := "aud_0ujsszwN8NRY24YaXiTIE2VWDTS" // string | 
    scheduleId := "sch_0ujsszwN8NRY24YaXiTIE2VWDTS" // string | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.AudiencesAPI.GetAudienceScheduleFromSpaceAndAudience(ctx, spaceId, id, scheduleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AudiencesAPI.GetAudienceScheduleFromSpaceAndAudience``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `GetAudienceScheduleFromSpaceAndAudience`: GetAudienceScheduleFromSpaceAndAudience200Response
    fmt.Fprintf(os.Stdout, "Response from `AudiencesAPI.GetAudienceScheduleFromSpaceAndAudience`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** |  | 
**id** | **string** |  | 
**scheduleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAudienceScheduleFromSpaceAndAudienceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**GetAudienceScheduleFromSpaceAndAudience200Response**](GetAudienceScheduleFromSpaceAndAudience200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1alpha+json, application/json

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
    search := *api.NewListAudienceSearchInput("Type_example", "Query_example") // ListAudienceSearchInput | Optional search criteria  This parameter exists in alpha. (optional)
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
 **search** | [**ListAudienceSearchInput**](ListAudienceSearchInput.md) | Optional search criteria  This parameter exists in alpha. | 
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


## Operation: ListAudienceSchedulesFromSpaceAndAudience

> ListAudienceSchedulesFromSpaceAndAudience200Response ListAudienceSchedulesFromSpaceAndAudience(ctx, spaceId, id).Execute()

List Audience Schedules from Space And Audience



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
    id := "aud_0ujsszwN8NRY24YaXiTIE2VWDTS" // string | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.AudiencesAPI.ListAudienceSchedulesFromSpaceAndAudience(ctx, spaceId, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AudiencesAPI.ListAudienceSchedulesFromSpaceAndAudience``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `ListAudienceSchedulesFromSpaceAndAudience`: ListAudienceSchedulesFromSpaceAndAudience200Response
    fmt.Fprintf(os.Stdout, "Response from `AudiencesAPI.ListAudienceSchedulesFromSpaceAndAudience`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAudienceSchedulesFromSpaceAndAudienceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ListAudienceSchedulesFromSpaceAndAudience200Response**](ListAudienceSchedulesFromSpaceAndAudience200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1alpha+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: ListAudiences

> ListAudiences200Response ListAudiences(ctx, spaceId).Search(search).Pagination(pagination).Include(include).Execute()

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
    search := *api.NewListAudienceSearchInput("Type_example", "Query_example") // ListAudienceSearchInput | Optional search criteria  This parameter exists in alpha. (optional)
    pagination := *api.NewListAudiencesPaginationInput(float32(123)) // ListAudiencesPaginationInput | Information about the pagination of this response.  [See pagination](https://docs.segmentapis.com/tag/Pagination/#section/Pagination-parameters) for more info.  This parameter exists in alpha. (optional)
    include := "include_example" // string | Additional resource to include, support schedules only.  This parameter exists in alpha. (optional)

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.AudiencesAPI.ListAudiences(ctx, spaceId).Search(search).Pagination(pagination).Include(include).Execute()
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

 **search** | [**ListAudienceSearchInput**](ListAudienceSearchInput.md) | Optional search criteria  This parameter exists in alpha. | 
 **pagination** | [**ListAudiencesPaginationInput**](ListAudiencesPaginationInput.md) | Information about the pagination of this response.  [See pagination](https://docs.segmentapis.com/tag/Pagination/#section/Pagination-parameters) for more info.  This parameter exists in alpha. | 
 **include** | **string** | Additional resource to include, support schedules only.  This parameter exists in alpha. | 

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
    id := "aud_0ujsszwN8NRY24YaXiTIE2VWDTS" // string | 

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
- **Accept**: application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json, application/json

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
    id := "aud_0ujsszwN8NRY24YaXiTIE2VWDTS" // string | 
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

