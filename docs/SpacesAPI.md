# \SpacesAPI

All URIs are relative to *https://api.segmentapis.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BatchQueryMessagingSubscriptionsForSpace**](SpacesAPI.md#BatchQueryMessagingSubscriptionsForSpace) | **Post** /spaces/{spaceId}/messaging-subscriptions/batch | Batch Query Messaging Subscriptions for Space
[**GetSpace**](SpacesAPI.md#GetSpace) | **Get** /spaces/{spaceId} | Get Space
[**ReplaceMessagingSubscriptionsInSpaces**](SpacesAPI.md#ReplaceMessagingSubscriptionsInSpaces) | **Put** /spaces/{spaceId}/messaging-subscriptions | Replace Messaging Subscriptions in Spaces



## Operation: BatchQueryMessagingSubscriptionsForSpace

> BatchQueryMessagingSubscriptionsForSpace200Response BatchQueryMessagingSubscriptionsForSpace(ctx, spaceId).BatchQueryMessagingSubscriptionsForSpaceAlphaInput(batchQueryMessagingSubscriptionsForSpaceAlphaInput).Execute()

Batch Query Messaging Subscriptions for Space



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
    batchQueryMessagingSubscriptionsForSpaceAlphaInput := *api.NewBatchQueryMessagingSubscriptionsForSpaceAlphaInput([]api.GetSubscriptionRequest{*api.NewGetSubscriptionRequest("Key_example", "Type_example")}) // BatchQueryMessagingSubscriptionsForSpaceAlphaInput | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "INSERT BEARER TOKEN HERE"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.SpacesAPI.BatchQueryMessagingSubscriptionsForSpace(ctx, spaceId).BatchQueryMessagingSubscriptionsForSpaceAlphaInput(batchQueryMessagingSubscriptionsForSpaceAlphaInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpacesAPI.BatchQueryMessagingSubscriptionsForSpace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `BatchQueryMessagingSubscriptionsForSpace`: BatchQueryMessagingSubscriptionsForSpace200Response
    fmt.Fprintf(os.Stdout, "Response from `SpacesAPI.BatchQueryMessagingSubscriptionsForSpace`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchQueryMessagingSubscriptionsForSpaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **batchQueryMessagingSubscriptionsForSpaceAlphaInput** | [**BatchQueryMessagingSubscriptionsForSpaceAlphaInput**](BatchQueryMessagingSubscriptionsForSpaceAlphaInput.md) |  | 

### Return type

[**BatchQueryMessagingSubscriptionsForSpace200Response**](BatchQueryMessagingSubscriptionsForSpace200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/vnd.segment.v1alpha+json
- **Accept**: application/vnd.segment.v1alpha+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: GetSpace

> GetSpace200Response GetSpace(ctx, spaceId).Execute()

Get Space



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

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "INSERT BEARER TOKEN HERE"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.SpacesAPI.GetSpace(ctx, spaceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpacesAPI.GetSpace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `GetSpace`: GetSpace200Response
    fmt.Fprintf(os.Stdout, "Response from `SpacesAPI.GetSpace`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSpaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetSpace200Response**](GetSpace200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1alpha+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: ReplaceMessagingSubscriptionsInSpaces

> ReplaceMessagingSubscriptionsInSpaces200Response ReplaceMessagingSubscriptionsInSpaces(ctx, spaceId).ReplaceMessagingSubscriptionsInSpacesAlphaInput(replaceMessagingSubscriptionsInSpacesAlphaInput).Execute()

Replace Messaging Subscriptions in Spaces



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
    replaceMessagingSubscriptionsInSpacesAlphaInput := *api.NewReplaceMessagingSubscriptionsInSpacesAlphaInput([]api.MessagesSubscriptionRequest{*api.NewMessagesSubscriptionRequest("Key_example", "Type_example")}) // ReplaceMessagingSubscriptionsInSpacesAlphaInput | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "INSERT BEARER TOKEN HERE"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.SpacesAPI.ReplaceMessagingSubscriptionsInSpaces(ctx, spaceId).ReplaceMessagingSubscriptionsInSpacesAlphaInput(replaceMessagingSubscriptionsInSpacesAlphaInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpacesAPI.ReplaceMessagingSubscriptionsInSpaces``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `ReplaceMessagingSubscriptionsInSpaces`: ReplaceMessagingSubscriptionsInSpaces200Response
    fmt.Fprintf(os.Stdout, "Response from `SpacesAPI.ReplaceMessagingSubscriptionsInSpaces`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceMessagingSubscriptionsInSpacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **replaceMessagingSubscriptionsInSpacesAlphaInput** | [**ReplaceMessagingSubscriptionsInSpacesAlphaInput**](ReplaceMessagingSubscriptionsInSpacesAlphaInput.md) |  | 

### Return type

[**ReplaceMessagingSubscriptionsInSpaces200Response**](ReplaceMessagingSubscriptionsInSpaces200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/vnd.segment.v1alpha+json
- **Accept**: application/vnd.segment.v1alpha+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

