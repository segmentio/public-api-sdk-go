# \DeliveryOverviewAPI

All URIs are relative to *https://api.segmentapis.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEgressFailedMetricsFromDeliveryOverview**](DeliveryOverviewAPI.md#GetEgressFailedMetricsFromDeliveryOverview) | **Get** /delivery-overview/failed-delivery | Get Egress Failed Metrics from Delivery Overview
[**GetEgressSuccessMetricsFromDeliveryOverview**](DeliveryOverviewAPI.md#GetEgressSuccessMetricsFromDeliveryOverview) | **Get** /delivery-overview/successful-delivery | Get Egress Success Metrics from Delivery Overview
[**GetFilteredAtDestinationMetricsFromDeliveryOverview**](DeliveryOverviewAPI.md#GetFilteredAtDestinationMetricsFromDeliveryOverview) | **Get** /delivery-overview/filtered-at-destination | Get Filtered At Destination Metrics from Delivery Overview
[**GetFilteredAtSourceMetricsFromDeliveryOverview**](DeliveryOverviewAPI.md#GetFilteredAtSourceMetricsFromDeliveryOverview) | **Get** /delivery-overview/filtered-at-source | Get Filtered At Source Metrics from Delivery Overview
[**GetIngressFailedMetricsFromDeliveryOverview**](DeliveryOverviewAPI.md#GetIngressFailedMetricsFromDeliveryOverview) | **Get** /delivery-overview/failed-on-ingest | Get Ingress Failed Metrics from Delivery Overview
[**GetIngressSuccessMetricsFromDeliveryOverview**](DeliveryOverviewAPI.md#GetIngressSuccessMetricsFromDeliveryOverview) | **Get** /delivery-overview/successfully-received | Get Ingress Success Metrics from Delivery Overview



## Operation: GetEgressFailedMetricsFromDeliveryOverview

> GetEgressFailedMetricsFromDeliveryOverview200Response GetEgressFailedMetricsFromDeliveryOverview(ctx).SourceId(sourceId).StartTime(startTime).EndTime(endTime).Granularity(granularity).DestinationConfigId(destinationConfigId).GroupBy(groupBy).Filter(filter).Pagination(pagination).SubscriptionId(subscriptionId).Execute()

Get Egress Failed Metrics from Delivery Overview



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
    sourceId := "rh5BDZp6QDHvXFCkibm1pR" // string | The sourceId for the workspace.  This parameter exists in alpha.
    startTime := "2024-01-01T00:00:00Z" // string | The ISO8601 formatted timestamp corresponding to the beginning of the requested timeframe, inclusive.  This parameter exists in alpha.
    endTime := "2024-01-03T00:00:00Z" // string | The ISO8601 formatted timestamp corresponding to the end of the requested timeframe, noninclusive.  This parameter exists in alpha.
    granularity := "day" // string | The size of each bucket in the requested window.  Based on the granularity chosen, there are restrictions on the time range you can query:  **Minute**: - Max time range: 4 hours - Oldest possible start time: 48 hours in the past  **Hour**: - Max Time range: 14 days - Oldest possible start time: 30 days in the past  **Day**: - Max time range: 30 days - Oldest possible start time: 30 days in the past  This parameter exists in alpha.
    destinationConfigId := "fP7qoQw2HTWt9WdMr718gn" // string | The ID tied to a workspace destination. DestinationConfigId is a required input for queries to Filtered at Destination, Failed Delivery, and Successful Delivery.  This parameter exists in alpha. (optional)
    groupBy := []string{[]string{"GroupBy_example"}} // []string | A comma-delimited list of strings representing one or more dimensions to group the result by.  Valid options are: `eventName`, `eventType`, `discardReason`, and `appVersion`.  This parameter exists in alpha. (optional)
    filter := *api.NewDeliveryOverviewFilterBy() // DeliveryOverviewFilterBy | An optional filter for `eventName`, `eventType`, `discardReason`, and/or `appVersion` that can be applied in addition to a `groupBy`. Ex: `filter: {discardReason: ['discard1'], eventName: ['name1', 'name2'], eventType: ['type1']}`.  This parameter exists in alpha. (optional)
    pagination := *api.NewPaginationInput(10) // PaginationInput | Optional params to specify the page cursor and count.  This parameter exists in alpha. (optional)
    subscriptionId := "subscriptionId_example" // string | An optional filter for actions destinations, to filter by a specific action.  This parameter exists in alpha. (optional)

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.DeliveryOverviewAPI.GetEgressFailedMetricsFromDeliveryOverview(ctx).SourceId(sourceId).StartTime(startTime).EndTime(endTime).Granularity(granularity).DestinationConfigId(destinationConfigId).GroupBy(groupBy).Filter(filter).Pagination(pagination).SubscriptionId(subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryOverviewAPI.GetEgressFailedMetricsFromDeliveryOverview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `GetEgressFailedMetricsFromDeliveryOverview`: GetEgressFailedMetricsFromDeliveryOverview200Response
    fmt.Fprintf(os.Stdout, "Response from `DeliveryOverviewAPI.GetEgressFailedMetricsFromDeliveryOverview`: %v\n", resp.GetData())
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEgressFailedMetricsFromDeliveryOverviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sourceId** | **string** | The sourceId for the workspace.  This parameter exists in alpha. | 
 **startTime** | **string** | The ISO8601 formatted timestamp corresponding to the beginning of the requested timeframe, inclusive.  This parameter exists in alpha. | 
 **endTime** | **string** | The ISO8601 formatted timestamp corresponding to the end of the requested timeframe, noninclusive.  This parameter exists in alpha. | 
 **granularity** | **string** | The size of each bucket in the requested window.  Based on the granularity chosen, there are restrictions on the time range you can query:  **Minute**: - Max time range: 4 hours - Oldest possible start time: 48 hours in the past  **Hour**: - Max Time range: 14 days - Oldest possible start time: 30 days in the past  **Day**: - Max time range: 30 days - Oldest possible start time: 30 days in the past  This parameter exists in alpha. | 
 **destinationConfigId** | **string** | The ID tied to a workspace destination. DestinationConfigId is a required input for queries to Filtered at Destination, Failed Delivery, and Successful Delivery.  This parameter exists in alpha. | 
 **groupBy** | **[][]string** | A comma-delimited list of strings representing one or more dimensions to group the result by.  Valid options are: &#x60;eventName&#x60;, &#x60;eventType&#x60;, &#x60;discardReason&#x60;, and &#x60;appVersion&#x60;.  This parameter exists in alpha. | 
 **filter** | [**DeliveryOverviewFilterBy**](DeliveryOverviewFilterBy.md) | An optional filter for &#x60;eventName&#x60;, &#x60;eventType&#x60;, &#x60;discardReason&#x60;, and/or &#x60;appVersion&#x60; that can be applied in addition to a &#x60;groupBy&#x60;. Ex: &#x60;filter: {discardReason: [&#39;discard1&#39;], eventName: [&#39;name1&#39;, &#39;name2&#39;], eventType: [&#39;type1&#39;]}&#x60;.  This parameter exists in alpha. | 
 **pagination** | [**PaginationInput**](PaginationInput.md) | Optional params to specify the page cursor and count.  This parameter exists in alpha. | 
 **subscriptionId** | **string** | An optional filter for actions destinations, to filter by a specific action.  This parameter exists in alpha. | 

### Return type

[**GetEgressFailedMetricsFromDeliveryOverview200Response**](GetEgressFailedMetricsFromDeliveryOverview200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1alpha+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: GetEgressSuccessMetricsFromDeliveryOverview

> GetEgressFailedMetricsFromDeliveryOverview200Response GetEgressSuccessMetricsFromDeliveryOverview(ctx).Execute()

Get Egress Success Metrics from Delivery Overview



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

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.DeliveryOverviewAPI.GetEgressSuccessMetricsFromDeliveryOverview(ctx).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryOverviewAPI.GetEgressSuccessMetricsFromDeliveryOverview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `GetEgressSuccessMetricsFromDeliveryOverview`: GetEgressFailedMetricsFromDeliveryOverview200Response
    fmt.Fprintf(os.Stdout, "Response from `DeliveryOverviewAPI.GetEgressSuccessMetricsFromDeliveryOverview`: %v\n", resp.GetData())
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetEgressSuccessMetricsFromDeliveryOverviewRequest struct via the builder pattern


### Return type

[**GetEgressFailedMetricsFromDeliveryOverview200Response**](GetEgressFailedMetricsFromDeliveryOverview200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1alpha+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: GetFilteredAtDestinationMetricsFromDeliveryOverview

> GetEgressFailedMetricsFromDeliveryOverview200Response GetFilteredAtDestinationMetricsFromDeliveryOverview(ctx).Execute()

Get Filtered At Destination Metrics from Delivery Overview



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

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.DeliveryOverviewAPI.GetFilteredAtDestinationMetricsFromDeliveryOverview(ctx).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryOverviewAPI.GetFilteredAtDestinationMetricsFromDeliveryOverview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `GetFilteredAtDestinationMetricsFromDeliveryOverview`: GetEgressFailedMetricsFromDeliveryOverview200Response
    fmt.Fprintf(os.Stdout, "Response from `DeliveryOverviewAPI.GetFilteredAtDestinationMetricsFromDeliveryOverview`: %v\n", resp.GetData())
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetFilteredAtDestinationMetricsFromDeliveryOverviewRequest struct via the builder pattern


### Return type

[**GetEgressFailedMetricsFromDeliveryOverview200Response**](GetEgressFailedMetricsFromDeliveryOverview200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1alpha+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: GetFilteredAtSourceMetricsFromDeliveryOverview

> GetEgressFailedMetricsFromDeliveryOverview200Response GetFilteredAtSourceMetricsFromDeliveryOverview(ctx).Execute()

Get Filtered At Source Metrics from Delivery Overview



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

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.DeliveryOverviewAPI.GetFilteredAtSourceMetricsFromDeliveryOverview(ctx).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryOverviewAPI.GetFilteredAtSourceMetricsFromDeliveryOverview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `GetFilteredAtSourceMetricsFromDeliveryOverview`: GetEgressFailedMetricsFromDeliveryOverview200Response
    fmt.Fprintf(os.Stdout, "Response from `DeliveryOverviewAPI.GetFilteredAtSourceMetricsFromDeliveryOverview`: %v\n", resp.GetData())
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetFilteredAtSourceMetricsFromDeliveryOverviewRequest struct via the builder pattern


### Return type

[**GetEgressFailedMetricsFromDeliveryOverview200Response**](GetEgressFailedMetricsFromDeliveryOverview200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1alpha+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: GetIngressFailedMetricsFromDeliveryOverview

> GetEgressFailedMetricsFromDeliveryOverview200Response GetIngressFailedMetricsFromDeliveryOverview(ctx).Execute()

Get Ingress Failed Metrics from Delivery Overview



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

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.DeliveryOverviewAPI.GetIngressFailedMetricsFromDeliveryOverview(ctx).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryOverviewAPI.GetIngressFailedMetricsFromDeliveryOverview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `GetIngressFailedMetricsFromDeliveryOverview`: GetEgressFailedMetricsFromDeliveryOverview200Response
    fmt.Fprintf(os.Stdout, "Response from `DeliveryOverviewAPI.GetIngressFailedMetricsFromDeliveryOverview`: %v\n", resp.GetData())
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIngressFailedMetricsFromDeliveryOverviewRequest struct via the builder pattern


### Return type

[**GetEgressFailedMetricsFromDeliveryOverview200Response**](GetEgressFailedMetricsFromDeliveryOverview200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1alpha+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: GetIngressSuccessMetricsFromDeliveryOverview

> GetEgressFailedMetricsFromDeliveryOverview200Response GetIngressSuccessMetricsFromDeliveryOverview(ctx).Execute()

Get Ingress Success Metrics from Delivery Overview



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

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.DeliveryOverviewAPI.GetIngressSuccessMetricsFromDeliveryOverview(ctx).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryOverviewAPI.GetIngressSuccessMetricsFromDeliveryOverview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `GetIngressSuccessMetricsFromDeliveryOverview`: GetEgressFailedMetricsFromDeliveryOverview200Response
    fmt.Fprintf(os.Stdout, "Response from `DeliveryOverviewAPI.GetIngressSuccessMetricsFromDeliveryOverview`: %v\n", resp.GetData())
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIngressSuccessMetricsFromDeliveryOverviewRequest struct via the builder pattern


### Return type

[**GetEgressFailedMetricsFromDeliveryOverview200Response**](GetEgressFailedMetricsFromDeliveryOverview200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1alpha+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

