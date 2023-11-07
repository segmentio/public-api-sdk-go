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
    createDestinationV1Input := *openapiclient.NewCreateDestinationV1Input("SourceId_example", "MetadataId_example", map[string]interface{}{"key": interface{}(123)}) // CreateDestinationV1Input | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    resp, r, err := apiClient.DestinationsAPI.CreateDestination(context.Background()).CreateDestinationV1Input(createDestinationV1Input).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DestinationsAPI.CreateDestination``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDestination`: CreateDestination200Response
    fmt.Fprintf(os.Stdout, "Response from `DestinationsAPI.CreateDestination`: %v\n", resp)
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

- **Content-Type**: application/vnd.segment.v1+json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json
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
    createDestinationSubscriptionAlphaInput := *openapiclient.NewCreateDestinationSubscriptionAlphaInput("Name_example", "ActionId_example", "Trigger_example", false) // CreateDestinationSubscriptionAlphaInput | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    resp, r, err := apiClient.DestinationsAPI.CreateDestinationSubscription(context.Background(), destinationId).CreateDestinationSubscriptionAlphaInput(createDestinationSubscriptionAlphaInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DestinationsAPI.CreateDestinationSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDestinationSubscription`: CreateDestinationSubscription200Response
    fmt.Fprintf(os.Stdout, "Response from `DestinationsAPI.CreateDestinationSubscription`: %v\n", resp)
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
    destinationId := "6424a4a5b03b3d8fe288f000" // string | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    resp, r, err := apiClient.DestinationsAPI.DeleteDestination(context.Background(), destinationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DestinationsAPI.DeleteDestination``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteDestination`: DeleteDestination200Response
    fmt.Fprintf(os.Stdout, "Response from `DestinationsAPI.DeleteDestination`: %v\n", resp)
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
    resp, r, err := apiClient.DestinationsAPI.GetDestination(context.Background(), destinationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DestinationsAPI.GetDestination``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDestination`: GetDestination200Response
    fmt.Fprintf(os.Stdout, "Response from `DestinationsAPI.GetDestination`: %v\n", resp)
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
    id := "iUyx2UdPSvp4uJtYAhjTup" // string | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    resp, r, err := apiClient.DestinationsAPI.GetSubscriptionFromDestination(context.Background(), destinationId, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DestinationsAPI.GetSubscriptionFromDestination``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSubscriptionFromDestination`: GetSubscriptionFromDestination200Response
    fmt.Fprintf(os.Stdout, "Response from `DestinationsAPI.GetSubscriptionFromDestination`: %v\n", resp)
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
    resp, r, err := apiClient.DestinationsAPI.ListDeliveryMetricsSummaryFromDestination(context.Background(), destinationId).SourceId(sourceId).StartTime(startTime).EndTime(endTime).Granularity(granularity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DestinationsAPI.ListDeliveryMetricsSummaryFromDestination``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDeliveryMetricsSummaryFromDestination`: ListDeliveryMetricsSummaryFromDestination200Response
    fmt.Fprintf(os.Stdout, "Response from `DestinationsAPI.ListDeliveryMetricsSummaryFromDestination`: %v\n", resp)
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
    pagination := *openapiclient.NewPaginationInput(float32(123)) // PaginationInput | Required pagination params for the request.  This parameter exists in v1.

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    resp, r, err := apiClient.DestinationsAPI.ListDestinations(context.Background()).Pagination(pagination).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DestinationsAPI.ListDestinations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDestinations`: ListDestinations200Response
    fmt.Fprintf(os.Stdout, "Response from `DestinationsAPI.ListDestinations`: %v\n", resp)
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
    pagination := *openapiclient.NewPaginationInput(float32(123)) // PaginationInput | Pagination options.  This parameter exists in alpha.

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    resp, r, err := apiClient.DestinationsAPI.ListSubscriptionsFromDestination(context.Background(), destinationId).Pagination(pagination).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DestinationsAPI.ListSubscriptionsFromDestination``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSubscriptionsFromDestination`: ListSubscriptionsFromDestination200Response
    fmt.Fprintf(os.Stdout, "Response from `DestinationsAPI.ListSubscriptionsFromDestination`: %v\n", resp)
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
    id := "bXbWrgh8vAeWioqx6Kg5vb" // string | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    resp, r, err := apiClient.DestinationsAPI.RemoveSubscriptionFromDestination(context.Background(), destinationId, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DestinationsAPI.RemoveSubscriptionFromDestination``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveSubscriptionFromDestination`: RemoveSubscriptionFromDestination200Response
    fmt.Fprintf(os.Stdout, "Response from `DestinationsAPI.RemoveSubscriptionFromDestination`: %v\n", resp)
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
    updateDestinationV1Input := *openapiclient.NewUpdateDestinationV1Input() // UpdateDestinationV1Input | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    resp, r, err := apiClient.DestinationsAPI.UpdateDestination(context.Background(), destinationId).UpdateDestinationV1Input(updateDestinationV1Input).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DestinationsAPI.UpdateDestination``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDestination`: UpdateDestination200Response
    fmt.Fprintf(os.Stdout, "Response from `DestinationsAPI.UpdateDestination`: %v\n", resp)
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

- **Content-Type**: application/vnd.segment.v1+json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json
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
    id := "mAdhWCVCmFadguoGcqP3vN" // string | 
    updateSubscriptionForDestinationAlphaInput := *openapiclient.NewUpdateSubscriptionForDestinationAlphaInput(*openapiclient.NewDestinationSubscriptionUpdateInput()) // UpdateSubscriptionForDestinationAlphaInput | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    resp, r, err := apiClient.DestinationsAPI.UpdateSubscriptionForDestination(context.Background(), destinationId, id).UpdateSubscriptionForDestinationAlphaInput(updateSubscriptionForDestinationAlphaInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DestinationsAPI.UpdateSubscriptionForDestination``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSubscriptionForDestination`: UpdateSubscriptionForDestination200Response
    fmt.Fprintf(os.Stdout, "Response from `DestinationsAPI.UpdateSubscriptionForDestination`: %v\n", resp)
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
