# \SpaceSchemaAPIAPI

All URIs are relative to *https://api.segmentapis.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListEntityPaths**](SpaceSchemaAPIAPI.md#ListEntityPaths) | **Get** /spaces/{spaceId}/entity-paths | List Entity Paths
[**ListEvents**](SpaceSchemaAPIAPI.md#ListEvents) | **Get** /spaces/{spaceId}/events | List Events
[**ListPropertiesFromEntity**](SpaceSchemaAPIAPI.md#ListPropertiesFromEntity) | **Get** /spaces/{spaceId}/entities/{entitySlug}/properties | List Properties from Entity
[**ListPropertiesFromEvent**](SpaceSchemaAPIAPI.md#ListPropertiesFromEvent) | **Get** /spaces/{spaceId}/events/{eventName}/properties | List Properties from Event
[**ListSampleValuesFromEntityProperty**](SpaceSchemaAPIAPI.md#ListSampleValuesFromEntityProperty) | **Get** /spaces/{spaceId}/entities/{entitySlug}/properties/{propertyName}/sample-values | List Sample Values from Entity Property
[**ListSampleValuesFromEventProperty**](SpaceSchemaAPIAPI.md#ListSampleValuesFromEventProperty) | **Get** /spaces/{spaceId}/events/{eventName}/properties/{propertyName}/sample-values | List Sample Values from Event Property
[**ListSampleValuesFromTrait**](SpaceSchemaAPIAPI.md#ListSampleValuesFromTrait) | **Get** /spaces/{spaceId}/traits/{traitKey}/sample-values | List Sample Values from Trait
[**ListTraits**](SpaceSchemaAPIAPI.md#ListTraits) | **Get** /spaces/{spaceId}/traits | List Traits



## Operation: ListEntityPaths

> ListEntityPaths200Response ListEntityPaths(ctx, spaceId).Pagination(pagination).Search(search).Execute()

List Entity Paths



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
    search := "search_example" // string | Filter paths by entity name or path name.  This parameter exists in alpha. (optional)

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.SpaceSchemaAPIAPI.ListEntityPaths(ctx, spaceId).Pagination(pagination).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpaceSchemaAPIAPI.ListEntityPaths``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `ListEntityPaths`: ListEntityPaths200Response
    fmt.Fprintf(os.Stdout, "Response from `SpaceSchemaAPIAPI.ListEntityPaths`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListEntityPathsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pagination** | [**PaginationInput**](PaginationInput.md) | Pagination params. Defaults to count 200.  This parameter exists in alpha. | 
 **search** | **string** | Filter paths by entity name or path name.  This parameter exists in alpha. | 

### Return type

[**ListEntityPaths200Response**](ListEntityPaths200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1alpha+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
    resp, r, err := apiClient.SpaceSchemaAPIAPI.ListEvents(ctx, spaceId).Pagination(pagination).SortBy(sortBy).SortDir(sortDir).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpaceSchemaAPIAPI.ListEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `ListEvents`: ListEvents200Response
    fmt.Fprintf(os.Stdout, "Response from `SpaceSchemaAPIAPI.ListEvents`: %v\n", resp.GetData())
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


## Operation: ListPropertiesFromEntity

> ListPropertiesFromEntity200Response ListPropertiesFromEntity(ctx, spaceId, entitySlug).Pagination(pagination).IncludeSampleValues(includeSampleValues).SamplesCount(samplesCount).Execute()

List Properties from Entity



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
    entitySlug := "my-entity" // string | 
    pagination := *api.NewPaginationInput(10) // PaginationInput | Pagination params. Defaults to count 200.  This parameter exists in alpha. (optional)
    includeSampleValues := true // bool | When true, include sample values for each property. Defaults to false.  This parameter exists in alpha. (optional)
    samplesCount := float32(8.14) // float32 | Max number of sample values to return per property. Defaults to 20, min 1, max 100.  This parameter exists in alpha. (optional)

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.SpaceSchemaAPIAPI.ListPropertiesFromEntity(ctx, spaceId, entitySlug).Pagination(pagination).IncludeSampleValues(includeSampleValues).SamplesCount(samplesCount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpaceSchemaAPIAPI.ListPropertiesFromEntity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `ListPropertiesFromEntity`: ListPropertiesFromEntity200Response
    fmt.Fprintf(os.Stdout, "Response from `SpaceSchemaAPIAPI.ListPropertiesFromEntity`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** |  | 
**entitySlug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListPropertiesFromEntityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pagination** | [**PaginationInput**](PaginationInput.md) | Pagination params. Defaults to count 200.  This parameter exists in alpha. | 
 **includeSampleValues** | **bool** | When true, include sample values for each property. Defaults to false.  This parameter exists in alpha. | 
 **samplesCount** | **float32** | Max number of sample values to return per property. Defaults to 20, min 1, max 100.  This parameter exists in alpha. | 

### Return type

[**ListPropertiesFromEntity200Response**](ListPropertiesFromEntity200Response.md)

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
    resp, r, err := apiClient.SpaceSchemaAPIAPI.ListPropertiesFromEvent(ctx, spaceId, eventName).Pagination(pagination).SortBy(sortBy).SortDir(sortDir).Search(search).IncludeSampleValues(includeSampleValues).SamplesCount(samplesCount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpaceSchemaAPIAPI.ListPropertiesFromEvent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `ListPropertiesFromEvent`: ListPropertiesFromEvent200Response
    fmt.Fprintf(os.Stdout, "Response from `SpaceSchemaAPIAPI.ListPropertiesFromEvent`: %v\n", resp.GetData())
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


## Operation: ListSampleValuesFromEntityProperty

> ListSampleValuesFromEntityProperty200Response ListSampleValuesFromEntityProperty(ctx, spaceId, entitySlug, propertyName).Execute()

List Sample Values from Entity Property



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
    entitySlug := "my-entity" // string | 
    propertyName := "email" // string | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.SpaceSchemaAPIAPI.ListSampleValuesFromEntityProperty(ctx, spaceId, entitySlug, propertyName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpaceSchemaAPIAPI.ListSampleValuesFromEntityProperty``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `ListSampleValuesFromEntityProperty`: ListSampleValuesFromEntityProperty200Response
    fmt.Fprintf(os.Stdout, "Response from `SpaceSchemaAPIAPI.ListSampleValuesFromEntityProperty`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** |  | 
**entitySlug** | **string** |  | 
**propertyName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSampleValuesFromEntityPropertyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ListSampleValuesFromEntityProperty200Response**](ListSampleValuesFromEntityProperty200Response.md)

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
    resp, r, err := apiClient.SpaceSchemaAPIAPI.ListSampleValuesFromEventProperty(ctx, spaceId, eventName, propertyName).PropertyType(propertyType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpaceSchemaAPIAPI.ListSampleValuesFromEventProperty``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `ListSampleValuesFromEventProperty`: ListSampleValuesFromEventProperty200Response
    fmt.Fprintf(os.Stdout, "Response from `SpaceSchemaAPIAPI.ListSampleValuesFromEventProperty`: %v\n", resp.GetData())
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


## Operation: ListSampleValuesFromTrait

> ListSampleValuesFromTrait200Response ListSampleValuesFromTrait(ctx, spaceId, traitKey).Collection(collection).Execute()

List Sample Values from Trait



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
    traitKey := "email" // string | 
    collection := "collection_example" // string | Collection to get trait values for. Defaults to 'users'.  This parameter exists in alpha. (optional)

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.SpaceSchemaAPIAPI.ListSampleValuesFromTrait(ctx, spaceId, traitKey).Collection(collection).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpaceSchemaAPIAPI.ListSampleValuesFromTrait``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `ListSampleValuesFromTrait`: ListSampleValuesFromTrait200Response
    fmt.Fprintf(os.Stdout, "Response from `SpaceSchemaAPIAPI.ListSampleValuesFromTrait`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** |  | 
**traitKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSampleValuesFromTraitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **collection** | **string** | Collection to get trait values for. Defaults to &#39;users&#39;.  This parameter exists in alpha. | 

### Return type

[**ListSampleValuesFromTrait200Response**](ListSampleValuesFromTrait200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1alpha+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: ListTraits

> ListTraits200Response ListTraits(ctx, spaceId).Pagination(pagination).SortBy(sortBy).SortDir(sortDir).Search(search).Collection(collection).IncludeSampleValues(includeSampleValues).SamplesCount(samplesCount).Execute()

List Traits



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
    sortBy := "trait" // string | Field to sort by. Defaults to 'trait'.  This parameter exists in alpha. (optional)
    sortDir := "asc" // string | Sort direction. Defaults to 'asc'.  This parameter exists in alpha. (optional)
    search := "search_example" // string | Filter traits by key substring.  This parameter exists in alpha. (optional)
    collection := "collection_example" // string | Collection to list traits for. Defaults to 'users'.  This parameter exists in alpha. (optional)
    includeSampleValues := true // bool | When true, include sample values for each trait. Defaults to false.  This parameter exists in alpha. (optional)
    samplesCount := float32(8.14) // float32 | Max number of sample values to return per trait. Defaults to 20, min 1, max 100.  This parameter exists in alpha. (optional)

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.SpaceSchemaAPIAPI.ListTraits(ctx, spaceId).Pagination(pagination).SortBy(sortBy).SortDir(sortDir).Search(search).Collection(collection).IncludeSampleValues(includeSampleValues).SamplesCount(samplesCount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpaceSchemaAPIAPI.ListTraits``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `ListTraits`: ListTraits200Response
    fmt.Fprintf(os.Stdout, "Response from `SpaceSchemaAPIAPI.ListTraits`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTraitsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pagination** | [**PaginationInput**](PaginationInput.md) | Pagination params. Defaults to count 200.  This parameter exists in alpha. | 
 **sortBy** | **string** | Field to sort by. Defaults to &#39;trait&#39;.  This parameter exists in alpha. | 
 **sortDir** | **string** | Sort direction. Defaults to &#39;asc&#39;.  This parameter exists in alpha. | 
 **search** | **string** | Filter traits by key substring.  This parameter exists in alpha. | 
 **collection** | **string** | Collection to list traits for. Defaults to &#39;users&#39;.  This parameter exists in alpha. | 
 **includeSampleValues** | **bool** | When true, include sample values for each trait. Defaults to false.  This parameter exists in alpha. | 
 **samplesCount** | **float32** | Max number of sample values to return per trait. Defaults to 20, min 1, max 100.  This parameter exists in alpha. | 

### Return type

[**ListTraits200Response**](ListTraits200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1alpha+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

