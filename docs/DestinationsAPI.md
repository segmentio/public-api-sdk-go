# \DestinationsAPI

All URIs are relative to *https://api.segmentapis.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDestination**](DestinationsAPI.md#CreateDestination) | **Post** /destinations | Create Destination
[**CreateDestinationSubscription**](DestinationsAPI.md#CreateDestinationSubscription) | **Post** /destinations/{destinationId}/subscriptions | Create Destination Subscription
[**DeleteDestination**](DestinationsAPI.md#DeleteDestination) | **Delete** /destinations/{destinationId} | Delete Destination
[**GetDestination**](DestinationsAPI.md#GetDestination) | **Get** /destinations/{destinationId} | Get Destination
[**GetSubscriptionFromDestination**](DestinationsAPI.md#GetSubscriptionFromDestination) | **Get** /destinations/{destinationId}/subscriptions/{id} | Get Subscription from Destination
[**ListDeliveryMetricsSummaryFromDestination**](DestinationsAPI.md#ListDeliveryMetricsSummaryFromDestination) | **Get** /destinations/{destinationId}/delivery-metrics | List Delivery Metrics Summary from Destination
[**ListDestinations**](DestinationsAPI.md#ListDestinations) | **Get** /destinations | List Destinations
[**ListSubscriptionsFromDestination**](DestinationsAPI.md#ListSubscriptionsFromDestination) | **Get** /destinations/{destinationId}/subscriptions | List Subscriptions from Destination
[**RemoveSubscriptionFromDestination**](DestinationsAPI.md#RemoveSubscriptionFromDestination) | **Delete** /destinations/{destinationId}/subscriptions/{id} | Remove Subscription from Destination
[**UpdateDestination**](DestinationsAPI.md#UpdateDestination) | **Patch** /destinations/{destinationId} | Update Destination
[**UpdateSubscriptionForDestination**](DestinationsAPI.md#UpdateSubscriptionForDestination) | **Patch** /destinations/{destinationId}/subscriptions/{id} | Update Subscription for Destination



## Operation: CreateDestination

> CreateDestination200Response CreateDestination(ctx).CreateDestinationV1Input(createDestinationV1Input).Execute()

Create Destination



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
    createDestinationV1Input := *api.NewCreateDestinationV1Input("SourceId_example", "MetadataId_example", map[string]interface{}{"key": interface{}(123)}) // CreateDestinationV1Input | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.DestinationsAPI.CreateDestination(ctx).CreateDestinationV1Input(createDestinationV1Input).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DestinationsAPI.CreateDestination``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `CreateDestination`: CreateDestination200Response
    fmt.Fprintf(os.Stdout, "Response from `DestinationsAPI.CreateDestination`: %v\n", resp.GetData())
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDestinationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createDestinationV1Input** | [**CreateDestinationV1Input**](CreateDestinationV1Input.md) |  | 

### Return type

[**CreateDestination200Response**](CreateDestination200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json, application/vnd.segment.v1+json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json
- **Accept**: application/vnd.segment.v1+json, application/json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: CreateDestinationSubscription

> CreateDestinationSubscription200Response CreateDestinationSubscription(ctx, destinationId).CreateDestinationSubscriptionAlphaInput(createDestinationSubscriptionAlphaInput).Execute()

Create Destination Subscription



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
    destinationId := "fP7qoQw2HTWt9WdMr718gn" // string | 
    createDestinationSubscriptionAlphaInput := *api.NewCreateDestinationSubscriptionAlphaInput("Name_example", "ActionId_example", "Trigger_example", false) // CreateDestinationSubscriptionAlphaInput | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.DestinationsAPI.CreateDestinationSubscription(ctx, destinationId).CreateDestinationSubscriptionAlphaInput(createDestinationSubscriptionAlphaInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DestinationsAPI.CreateDestinationSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `CreateDestinationSubscription`: CreateDestinationSubscription200Response
    fmt.Fprintf(os.Stdout, "Response from `DestinationsAPI.CreateDestinationSubscription`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**destinationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDestinationSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createDestinationSubscriptionAlphaInput** | [**CreateDestinationSubscriptionAlphaInput**](CreateDestinationSubscriptionAlphaInput.md) |  | 

### Return type

[**CreateDestinationSubscription200Response**](CreateDestinationSubscription200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/vnd.segment.v1alpha+json
- **Accept**: application/vnd.segment.v1alpha+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: DeleteDestination

> DeleteDestination200Response DeleteDestination(ctx, destinationId).Execute()

Delete Destination



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
    destinationId := "65c2bdbede6f2d8297f943db" // string | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.DestinationsAPI.DeleteDestination(ctx, destinationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DestinationsAPI.DeleteDestination``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `DeleteDestination`: DeleteDestination200Response
    fmt.Fprintf(os.Stdout, "Response from `DestinationsAPI.DeleteDestination`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**destinationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDestinationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteDestination200Response**](DeleteDestination200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1+json, application/json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: GetDestination

> GetDestination200Response GetDestination(ctx, destinationId).Execute()

Get Destination



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
    destinationId := "fP7qoQw2HTWt9WdMr718gn" // string | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.DestinationsAPI.GetDestination(ctx, destinationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DestinationsAPI.GetDestination``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `GetDestination`: GetDestination200Response
    fmt.Fprintf(os.Stdout, "Response from `DestinationsAPI.GetDestination`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**destinationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDestinationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetDestination200Response**](GetDestination200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1+json, application/json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: GetSubscriptionFromDestination

> GetSubscriptionFromDestination200Response GetSubscriptionFromDestination(ctx, destinationId, id).Execute()

Get Subscription from Destination



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
    destinationId := "fP7qoQw2HTWt9WdMr718gn" // string | 
    id := "kyMKN6LUgMvF8dwRMEz3cX" // string | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.DestinationsAPI.GetSubscriptionFromDestination(ctx, destinationId, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DestinationsAPI.GetSubscriptionFromDestination``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `GetSubscriptionFromDestination`: GetSubscriptionFromDestination200Response
    fmt.Fprintf(os.Stdout, "Response from `DestinationsAPI.GetSubscriptionFromDestination`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**destinationId** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSubscriptionFromDestinationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetSubscriptionFromDestination200Response**](GetSubscriptionFromDestination200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1alpha+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: ListDeliveryMetricsSummaryFromDestination

> ListDeliveryMetricsSummaryFromDestination200Response ListDeliveryMetricsSummaryFromDestination(ctx, destinationId).SourceId(sourceId).StartTime(startTime).EndTime(endTime).Granularity(granularity).Execute()

List Delivery Metrics Summary from Destination



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
    destinationId := "fP7qoQw2HTWt9WdMr718gn" // string | 
    sourceId := "rh5BDZp6QDHvXFCkibm1pR" // string | The id of the Source linked to the Destination.  Config API note: analogous to `parent`.  This parameter exists in beta.
    startTime := "2006-01-02T15:04:05.000Z" // string | Filter events that happened after this time.  Defaults to: - 1 hour ago if granularity is `MINUTE`. - 7 days ago if granularity is `HOUR`. - 30 days ago if granularity is `DAY`.  This parameter exists in beta. (optional)
    endTime := "2006-01-02T15:04:05.000Z" // string | Filter events that happened before this time. Defaults to now if not set.  This parameter exists in beta. (optional)
    granularity := "DAY" // string | The granularity to filter metrics to. Either `MINUTE`, `HOUR` or `DAY`.  Defaults to `MINUTE` if not set.  This parameter exists in beta. (optional)

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.DestinationsAPI.ListDeliveryMetricsSummaryFromDestination(ctx, destinationId).SourceId(sourceId).StartTime(startTime).EndTime(endTime).Granularity(granularity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DestinationsAPI.ListDeliveryMetricsSummaryFromDestination``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `ListDeliveryMetricsSummaryFromDestination`: ListDeliveryMetricsSummaryFromDestination200Response
    fmt.Fprintf(os.Stdout, "Response from `DestinationsAPI.ListDeliveryMetricsSummaryFromDestination`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**destinationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListDeliveryMetricsSummaryFromDestinationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sourceId** | **string** | The id of the Source linked to the Destination.  Config API note: analogous to &#x60;parent&#x60;.  This parameter exists in beta. | 
 **startTime** | **string** | Filter events that happened after this time.  Defaults to: - 1 hour ago if granularity is &#x60;MINUTE&#x60;. - 7 days ago if granularity is &#x60;HOUR&#x60;. - 30 days ago if granularity is &#x60;DAY&#x60;.  This parameter exists in beta. | 
 **endTime** | **string** | Filter events that happened before this time. Defaults to now if not set.  This parameter exists in beta. | 
 **granularity** | **string** | The granularity to filter metrics to. Either &#x60;MINUTE&#x60;, &#x60;HOUR&#x60; or &#x60;DAY&#x60;.  Defaults to &#x60;MINUTE&#x60; if not set.  This parameter exists in beta. | 

### Return type

[**ListDeliveryMetricsSummaryFromDestination200Response**](ListDeliveryMetricsSummaryFromDestination200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: ListDestinations

> ListDestinations200Response ListDestinations(ctx).Pagination(pagination).Execute()

List Destinations



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
    pagination := *api.NewPaginationInput(10) // PaginationInput | Required pagination params for the request.  This parameter exists in v1. (optional)

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.DestinationsAPI.ListDestinations(ctx).Pagination(pagination).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DestinationsAPI.ListDestinations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `ListDestinations`: ListDestinations200Response
    fmt.Fprintf(os.Stdout, "Response from `DestinationsAPI.ListDestinations`: %v\n", resp.GetData())
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDestinationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pagination** | [**PaginationInput**](PaginationInput.md) | Required pagination params for the request.  This parameter exists in v1. | 

### Return type

[**ListDestinations200Response**](ListDestinations200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1+json, application/json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: ListSubscriptionsFromDestination

> ListSubscriptionsFromDestination200Response ListSubscriptionsFromDestination(ctx, destinationId).Pagination(pagination).Execute()

List Subscriptions from Destination



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
    destinationId := "fP7qoQw2HTWt9WdMr718gn" // string | 
    pagination := *api.NewPaginationInput(10) // PaginationInput | Pagination options.  This parameter exists in alpha. (optional)

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.DestinationsAPI.ListSubscriptionsFromDestination(ctx, destinationId).Pagination(pagination).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DestinationsAPI.ListSubscriptionsFromDestination``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `ListSubscriptionsFromDestination`: ListSubscriptionsFromDestination200Response
    fmt.Fprintf(os.Stdout, "Response from `DestinationsAPI.ListSubscriptionsFromDestination`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**destinationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSubscriptionsFromDestinationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pagination** | [**PaginationInput**](PaginationInput.md) | Pagination options.  This parameter exists in alpha. | 

### Return type

[**ListSubscriptionsFromDestination200Response**](ListSubscriptionsFromDestination200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1alpha+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: RemoveSubscriptionFromDestination

> RemoveSubscriptionFromDestination200Response RemoveSubscriptionFromDestination(ctx, destinationId, id).Execute()

Remove Subscription from Destination



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
    destinationId := "fP7qoQw2HTWt9WdMr718gn" // string | 
    id := "zXCqmEMHJojkD45GcBAPt" // string | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.DestinationsAPI.RemoveSubscriptionFromDestination(ctx, destinationId, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DestinationsAPI.RemoveSubscriptionFromDestination``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `RemoveSubscriptionFromDestination`: RemoveSubscriptionFromDestination200Response
    fmt.Fprintf(os.Stdout, "Response from `DestinationsAPI.RemoveSubscriptionFromDestination`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**destinationId** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveSubscriptionFromDestinationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RemoveSubscriptionFromDestination200Response**](RemoveSubscriptionFromDestination200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1alpha+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: UpdateDestination

> UpdateDestination200Response UpdateDestination(ctx, destinationId).UpdateDestinationV1Input(updateDestinationV1Input).Execute()

Update Destination



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
    destinationId := "fP7qoQw2HTWt9WdMr718gn" // string | 
    updateDestinationV1Input := *api.NewUpdateDestinationV1Input() // UpdateDestinationV1Input | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.DestinationsAPI.UpdateDestination(ctx, destinationId).UpdateDestinationV1Input(updateDestinationV1Input).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DestinationsAPI.UpdateDestination``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `UpdateDestination`: UpdateDestination200Response
    fmt.Fprintf(os.Stdout, "Response from `DestinationsAPI.UpdateDestination`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**destinationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDestinationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDestinationV1Input** | [**UpdateDestinationV1Input**](UpdateDestinationV1Input.md) |  | 

### Return type

[**UpdateDestination200Response**](UpdateDestination200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json, application/vnd.segment.v1+json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json
- **Accept**: application/vnd.segment.v1+json, application/json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: UpdateSubscriptionForDestination

> UpdateSubscriptionForDestination200Response UpdateSubscriptionForDestination(ctx, destinationId, id).UpdateSubscriptionForDestinationAlphaInput(updateSubscriptionForDestinationAlphaInput).Execute()

Update Subscription for Destination



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
    destinationId := "fP7qoQw2HTWt9WdMr718gn" // string | 
    id := "3ELMSracBm5MMikXBYfo1c" // string | 
    updateSubscriptionForDestinationAlphaInput := *api.NewUpdateSubscriptionForDestinationAlphaInput(*api.NewDestinationSubscriptionUpdateInput()) // UpdateSubscriptionForDestinationAlphaInput | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.DestinationsAPI.UpdateSubscriptionForDestination(ctx, destinationId, id).UpdateSubscriptionForDestinationAlphaInput(updateSubscriptionForDestinationAlphaInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DestinationsAPI.UpdateSubscriptionForDestination``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `UpdateSubscriptionForDestination`: UpdateSubscriptionForDestination200Response
    fmt.Fprintf(os.Stdout, "Response from `DestinationsAPI.UpdateSubscriptionForDestination`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**destinationId** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSubscriptionForDestinationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateSubscriptionForDestinationAlphaInput** | [**UpdateSubscriptionForDestinationAlphaInput**](UpdateSubscriptionForDestinationAlphaInput.md) |  | 

### Return type

[**UpdateSubscriptionForDestination200Response**](UpdateSubscriptionForDestination200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/vnd.segment.v1alpha+json
- **Accept**: application/vnd.segment.v1alpha+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

