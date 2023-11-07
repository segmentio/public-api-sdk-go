# \AudiencesAPI

All URIs are relative to *https://api.segmentapis.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAudience**](AudiencesAPI.md#GetAudience) | **Get** /spaces/{spaceId}/audiences/{id} | Get Audience
[**ListAudiences**](AudiencesAPI.md#ListAudiences) | **Get** /spaces/{spaceId}/audiences | List Audiences
[**RemoveAudienceFromSpace**](AudiencesAPI.md#RemoveAudienceFromSpace) | **Delete** /spaces/{spaceId}/audiences/{id} | Remove Audience from Space
[**UpdateAudienceForSpace**](AudiencesAPI.md#UpdateAudienceForSpace) | **Patch** /spaces/{spaceId}/audiences/{id} | Update Audience for Space



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
    spaceId := "spaceId" // string | 
    id := "id" // string | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    resp, r, err := apiClient.AudiencesAPI.GetAudience(context.Background(), spaceId, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AudiencesAPI.GetAudience``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAudience`: GetAudience200Response
    fmt.Fprintf(os.Stdout, "Response from `AudiencesAPI.GetAudience`: %v\n", resp)
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
    spaceId := "spaceId" // string | 
    pagination := *openapiclient.NewPaginationInput(float32(123)) // PaginationInput | Information about the pagination of this response.  This parameter exists in alpha.

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    resp, r, err := apiClient.AudiencesAPI.ListAudiences(context.Background(), spaceId).Pagination(pagination).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AudiencesAPI.ListAudiences``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAudiences`: ListAudiences200Response
    fmt.Fprintf(os.Stdout, "Response from `AudiencesAPI.ListAudiences`: %v\n", resp)
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

 **pagination** | [**PaginationInput**](PaginationInput.md) | Information about the pagination of this response.  This parameter exists in alpha. | 

### Return type

[**ListAudiences200Response**](ListAudiences200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
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
    spaceId := "spaceId" // string | 
    id := "id" // string | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    resp, r, err := apiClient.AudiencesAPI.RemoveAudienceFromSpace(context.Background(), spaceId, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AudiencesAPI.RemoveAudienceFromSpace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveAudienceFromSpace`: RemoveAudienceFromSpace200Response
    fmt.Fprintf(os.Stdout, "Response from `AudiencesAPI.RemoveAudienceFromSpace`: %v\n", resp)
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

> UpdateAudienceForSpace200Response UpdateAudienceForSpace(ctx, spaceId, id).UpdateAudienceForSpaceInput(updateAudienceForSpaceInput).Execute()

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
    spaceId := "spaceId" // string | 
    id := "id" // string | 
    updateAudienceForSpaceInput := *openapiclient.NewUpdateAudienceForSpaceInput(false) // UpdateAudienceForSpaceInput | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    resp, r, err := apiClient.AudiencesAPI.UpdateAudienceForSpace(context.Background(), spaceId, id).UpdateAudienceForSpaceInput(updateAudienceForSpaceInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AudiencesAPI.UpdateAudienceForSpace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAudienceForSpace`: UpdateAudienceForSpace200Response
    fmt.Fprintf(os.Stdout, "Response from `AudiencesAPI.UpdateAudienceForSpace`: %v\n", resp)
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


 **updateAudienceForSpaceInput** | [**UpdateAudienceForSpaceInput**](UpdateAudienceForSpaceInput.md) |  | 

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
