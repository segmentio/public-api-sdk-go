# \SpaceSchemaAPI

All URIs are relative to *https://api.segmentapis.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListEvents**](SpaceSchemaAPI.md#ListEvents) | **Get** /spaces/{spaceId}/events | List Events
[**ListPropertiesFromEvent**](SpaceSchemaAPI.md#ListPropertiesFromEvent) | **Get** /spaces/{spaceId}/events/{eventName}/properties | List Properties from Event
[**ListSampleValuesFromEventProperty**](SpaceSchemaAPI.md#ListSampleValuesFromEventProperty) | **Get** /spaces/{spaceId}/events/{eventName}/properties/{propertyName}/sample-values | List Sample Values from Event Property



## Operation: ListEvents

> ListEvents200Response ListEvents(ctx, spaceId).Pagination(pagination).SortBy(sortBy).SortDir(sortDir).Search(search).Execute()

List Events



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
    pagination := *api.NewPaginationInput(10) // PaginationInput | Pagination params. Defaults to count 200.  This parameter exists in alpha. (optional)
    sortBy := "lastSeenAt" // string | Field to sort by. Defaults to 'lastSeenAt'.  This parameter exists in alpha. (optional)
    sortDir := "desc" // string | Sort direction. Defaults to 'desc'.  This parameter exists in alpha. (optional)
    search := "search_example" // string | Filter events by name substring.  This parameter exists in alpha. (optional)

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.SpaceSchemaAPI.ListEvents(ctx, spaceId).Pagination(pagination).SortBy(sortBy).SortDir(sortDir).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpaceSchemaAPI.ListEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `ListEvents`: ListEvents200Response
    fmt.Fprintf(os.Stdout, "Response from `SpaceSchemaAPI.ListEvents`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pagination** | [**PaginationInput**](PaginationInput.md) | Pagination params. Defaults to count 200.  This parameter exists in alpha. | 
 **sortBy** | **string** | Field to sort by. Defaults to &#39;lastSeenAt&#39;.  This parameter exists in alpha. | 
 **sortDir** | **string** | Sort direction. Defaults to &#39;desc&#39;.  This parameter exists in alpha. | 
 **search** | **string** | Filter events by name substring.  This parameter exists in alpha. | 

### Return type

[**ListEvents200Response**](ListEvents200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1alpha+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: ListPropertiesFromEvent

> ListPropertiesFromEvent200Response ListPropertiesFromEvent(ctx, spaceId, eventName).Pagination(pagination).SortBy(sortBy).SortDir(sortDir).Search(search).IncludeSampleValues(includeSampleValues).SamplesCount(samplesCount).Execute()

List Properties from Event



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
    eventName := "Order Completed" // string | 
    pagination := *api.NewPaginationInput(10) // PaginationInput | Pagination params. Defaults to count 200.  This parameter exists in alpha. (optional)
    sortBy := "name" // string | Field to sort by. Defaults to 'lastSeenAt'.  This parameter exists in alpha. (optional)
    sortDir := "asc" // string | Sort direction. Defaults to 'desc'.  This parameter exists in alpha. (optional)
    search := "search_example" // string | Filter properties by name substring.  This parameter exists in alpha. (optional)
    includeSampleValues := true // bool | When true, include sample values for each property. Defaults to false.  This parameter exists in alpha. (optional)
    samplesCount := float32(8.14) // float32 | Max number of sample values to return per property. Defaults to 20, min 1, max 100.  This parameter exists in alpha. (optional)

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.SpaceSchemaAPI.ListPropertiesFromEvent(ctx, spaceId, eventName).Pagination(pagination).SortBy(sortBy).SortDir(sortDir).Search(search).IncludeSampleValues(includeSampleValues).SamplesCount(samplesCount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpaceSchemaAPI.ListPropertiesFromEvent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `ListPropertiesFromEvent`: ListPropertiesFromEvent200Response
    fmt.Fprintf(os.Stdout, "Response from `SpaceSchemaAPI.ListPropertiesFromEvent`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** |  | 
**eventName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListPropertiesFromEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pagination** | [**PaginationInput**](PaginationInput.md) | Pagination params. Defaults to count 200.  This parameter exists in alpha. | 
 **sortBy** | **string** | Field to sort by. Defaults to &#39;lastSeenAt&#39;.  This parameter exists in alpha. | 
 **sortDir** | **string** | Sort direction. Defaults to &#39;desc&#39;.  This parameter exists in alpha. | 
 **search** | **string** | Filter properties by name substring.  This parameter exists in alpha. | 
 **includeSampleValues** | **bool** | When true, include sample values for each property. Defaults to false.  This parameter exists in alpha. | 
 **samplesCount** | **float32** | Max number of sample values to return per property. Defaults to 20, min 1, max 100.  This parameter exists in alpha. | 

### Return type

[**ListPropertiesFromEvent200Response**](ListPropertiesFromEvent200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1alpha+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: ListSampleValuesFromEventProperty

> ListSampleValuesFromEventProperty200Response ListSampleValuesFromEventProperty(ctx, spaceId, eventName, propertyName).PropertyType(propertyType).Execute()

List Sample Values from Event Property



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
    eventName := "Order Completed" // string | 
    propertyName := "revenue" // string | 
    propertyType := api.EventPropertyType("CONTEXT") // EventPropertyType | The property type.  This parameter exists in alpha.

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.SpaceSchemaAPI.ListSampleValuesFromEventProperty(ctx, spaceId, eventName, propertyName).PropertyType(propertyType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpaceSchemaAPI.ListSampleValuesFromEventProperty``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `ListSampleValuesFromEventProperty`: ListSampleValuesFromEventProperty200Response
    fmt.Fprintf(os.Stdout, "Response from `SpaceSchemaAPI.ListSampleValuesFromEventProperty`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** |  | 
**eventName** | **string** |  | 
**propertyName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSampleValuesFromEventPropertyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **propertyType** | [**EventPropertyType**](EventPropertyType.md) | The property type.  This parameter exists in alpha. | 

### Return type

[**ListSampleValuesFromEventProperty200Response**](ListSampleValuesFromEventProperty200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1alpha+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

