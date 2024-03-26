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

> GetEgressFailedMetricsFromDeliveryOverview200Response GetEgressFailedMetricsFromDeliveryOverview(ctx).SourceId(sourceId).DestinationConfigId(destinationConfigId).StartTime(startTime).EndTime(endTime).Granularity(granularity).Pagination(pagination).GroupBy(groupBy).Filter(filter).SubscriptionId(subscriptionId).Execute()

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
    sourceId := "rh5BDZp6QDHvXFCkibm1pR" // string | The sourceId for the Workspace.  This parameter exists in beta.
    destinationConfigId := "fP7qoQw2HTWt9WdMr718gn" // string | The id tied to a Workspace Destination.  This parameter exists in beta.
    startTime := "2024-01-01T00:00:00Z" // string | The ISO8601 formatted timestamp corresponding to the beginning of the requested timeframe, inclusive.  This parameter exists in beta.
    endTime := "2024-01-03T00:00:00Z" // string | The ISO8601 formatted timestamp corresponding to the end of the requested timeframe, noninclusive.  This parameter exists in beta.
    granularity := "day" // string | The size of each bucket in the requested window.  Based on the granularity chosen, there are restrictions on the time range you can query:  **Minute**: - Max time range: 4 hours - Oldest possible start time: 48 hours in the past  **Hour**: - Max Time range: 14 days - Oldest possible start time: 30 days in the past  **Day**: - Max time range: 30 days - Oldest possible start time: 30 days in the past  This parameter exists in beta.
    pagination := *api.NewPaginationInput(10) // PaginationInput | Params to specify the page cursor and count.  This parameter exists in beta.
    groupBy := []string{"Inner_example"} // []string | A comma-delimited list of strings representing one or more dimensions to group the result by.  Valid options are: `eventName`, `eventType`, `discardReason`, and `appVersion`.  This parameter exists in beta. (optional)
    filter := *api.NewDeliveryOverviewFilterBy() // DeliveryOverviewFilterBy | An optional filter for `eventName`, `eventType`, `discardReason`, and/or `appVersion` that can be applied in addition to a `groupBy`.  This parameter exists in beta. (optional)
    subscriptionId := "subscriptionId_example" // string | An optional filter for actions destinations, to filter by a specific action.  This parameter exists in beta. (optional)

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.DeliveryOverviewAPI.GetEgressFailedMetricsFromDeliveryOverview(ctx).SourceId(sourceId).DestinationConfigId(destinationConfigId).StartTime(startTime).EndTime(endTime).Granularity(granularity).Pagination(pagination).GroupBy(groupBy).Filter(filter).SubscriptionId(subscriptionId).Execute()
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
 **sourceId** | **string** | The sourceId for the Workspace.  This parameter exists in beta. | 
 **destinationConfigId** | **string** | The id tied to a Workspace Destination.  This parameter exists in beta. | 
 **startTime** | **string** | The ISO8601 formatted timestamp corresponding to the beginning of the requested timeframe, inclusive.  This parameter exists in beta. | 
 **endTime** | **string** | The ISO8601 formatted timestamp corresponding to the end of the requested timeframe, noninclusive.  This parameter exists in beta. | 
 **granularity** | **string** | The size of each bucket in the requested window.  Based on the granularity chosen, there are restrictions on the time range you can query:  **Minute**: - Max time range: 4 hours - Oldest possible start time: 48 hours in the past  **Hour**: - Max Time range: 14 days - Oldest possible start time: 30 days in the past  **Day**: - Max time range: 30 days - Oldest possible start time: 30 days in the past  This parameter exists in beta. | 
 **pagination** | [**PaginationInput**](PaginationInput.md) | Params to specify the page cursor and count.  This parameter exists in beta. | 
 **groupBy** | **[]string** | A comma-delimited list of strings representing one or more dimensions to group the result by.  Valid options are: &#x60;eventName&#x60;, &#x60;eventType&#x60;, &#x60;discardReason&#x60;, and &#x60;appVersion&#x60;.  This parameter exists in beta. | 
 **filter** | [**DeliveryOverviewFilterBy**](DeliveryOverviewFilterBy.md) | An optional filter for &#x60;eventName&#x60;, &#x60;eventType&#x60;, &#x60;discardReason&#x60;, and/or &#x60;appVersion&#x60; that can be applied in addition to a &#x60;groupBy&#x60;.  This parameter exists in beta. | 
 **subscriptionId** | **string** | An optional filter for actions destinations, to filter by a specific action.  This parameter exists in beta. | 

### Return type

[**GetEgressFailedMetricsFromDeliveryOverview200Response**](GetEgressFailedMetricsFromDeliveryOverview200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1beta+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: GetEgressSuccessMetricsFromDeliveryOverview

> GetEgressFailedMetricsFromDeliveryOverview200Response GetEgressSuccessMetricsFromDeliveryOverview(ctx).SourceId(sourceId).DestinationConfigId(destinationConfigId).StartTime(startTime).EndTime(endTime).Granularity(granularity).Pagination(pagination).GroupBy(groupBy).Filter(filter).SubscriptionId(subscriptionId).Execute()

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
    sourceId := "rh5BDZp6QDHvXFCkibm1pR" // string | The sourceId for the Workspace.  This parameter exists in beta.
    destinationConfigId := "fP7qoQw2HTWt9WdMr718gn" // string | The id tied to a Workspace Destination.  This parameter exists in beta.
    startTime := "2024-01-01T00:00:00Z" // string | The ISO8601 formatted timestamp corresponding to the beginning of the requested timeframe, inclusive.  This parameter exists in beta.
    endTime := "2024-01-03T00:00:00Z" // string | The ISO8601 formatted timestamp corresponding to the end of the requested timeframe, noninclusive.  This parameter exists in beta.
    granularity := "day" // string | The size of each bucket in the requested window.  Based on the granularity chosen, there are restrictions on the time range you can query:  **Minute**: - Max time range: 4 hours - Oldest possible start time: 48 hours in the past  **Hour**: - Max Time range: 14 days - Oldest possible start time: 30 days in the past  **Day**: - Max time range: 30 days - Oldest possible start time: 30 days in the past  This parameter exists in beta.
    pagination := *api.NewPaginationInput(10) // PaginationInput | Params to specify the page cursor and count.  This parameter exists in beta.
    groupBy := []string{"Inner_example"} // []string | A comma-delimited list of strings representing one or more dimensions to group the result by.  Valid options are: `eventName`, `eventType`, `discardReason`, and `appVersion`.  This parameter exists in beta. (optional)
    filter := *api.NewDeliveryOverviewFilterBy() // DeliveryOverviewFilterBy | An optional filter for `eventName`, `eventType`, `discardReason`, and/or `appVersion` that can be applied in addition to a `groupBy`. If you would like to view retry attempts for a successful delivery, you can filter `discardReason` from `successes.attempt.1` through `successes.attempt.10`.  This parameter exists in beta. (optional)
    subscriptionId := "subscriptionId_example" // string | An optional filter for actions destinations, to filter by a specific action.  This parameter exists in beta. (optional)

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.DeliveryOverviewAPI.GetEgressSuccessMetricsFromDeliveryOverview(ctx).SourceId(sourceId).DestinationConfigId(destinationConfigId).StartTime(startTime).EndTime(endTime).Granularity(granularity).Pagination(pagination).GroupBy(groupBy).Filter(filter).SubscriptionId(subscriptionId).Execute()
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



### Other Parameters

Other parameters are passed through a pointer to a apiGetEgressSuccessMetricsFromDeliveryOverviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sourceId** | **string** | The sourceId for the Workspace.  This parameter exists in beta. | 
 **destinationConfigId** | **string** | The id tied to a Workspace Destination.  This parameter exists in beta. | 
 **startTime** | **string** | The ISO8601 formatted timestamp corresponding to the beginning of the requested timeframe, inclusive.  This parameter exists in beta. | 
 **endTime** | **string** | The ISO8601 formatted timestamp corresponding to the end of the requested timeframe, noninclusive.  This parameter exists in beta. | 
 **granularity** | **string** | The size of each bucket in the requested window.  Based on the granularity chosen, there are restrictions on the time range you can query:  **Minute**: - Max time range: 4 hours - Oldest possible start time: 48 hours in the past  **Hour**: - Max Time range: 14 days - Oldest possible start time: 30 days in the past  **Day**: - Max time range: 30 days - Oldest possible start time: 30 days in the past  This parameter exists in beta. | 
 **pagination** | [**PaginationInput**](PaginationInput.md) | Params to specify the page cursor and count.  This parameter exists in beta. | 
 **groupBy** | **[]string** | A comma-delimited list of strings representing one or more dimensions to group the result by.  Valid options are: &#x60;eventName&#x60;, &#x60;eventType&#x60;, &#x60;discardReason&#x60;, and &#x60;appVersion&#x60;.  This parameter exists in beta. | 
 **filter** | [**DeliveryOverviewFilterBy**](DeliveryOverviewFilterBy.md) | An optional filter for &#x60;eventName&#x60;, &#x60;eventType&#x60;, &#x60;discardReason&#x60;, and/or &#x60;appVersion&#x60; that can be applied in addition to a &#x60;groupBy&#x60;. If you would like to view retry attempts for a successful delivery, you can filter &#x60;discardReason&#x60; from &#x60;successes.attempt.1&#x60; through &#x60;successes.attempt.10&#x60;.  This parameter exists in beta. | 
 **subscriptionId** | **string** | An optional filter for actions destinations, to filter by a specific action.  This parameter exists in beta. | 

### Return type

[**GetEgressFailedMetricsFromDeliveryOverview200Response**](GetEgressFailedMetricsFromDeliveryOverview200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1beta+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: GetFilteredAtDestinationMetricsFromDeliveryOverview

> GetEgressFailedMetricsFromDeliveryOverview200Response GetFilteredAtDestinationMetricsFromDeliveryOverview(ctx).SourceId(sourceId).DestinationConfigId(destinationConfigId).StartTime(startTime).EndTime(endTime).Granularity(granularity).Pagination(pagination).GroupBy(groupBy).Filter(filter).SubscriptionId(subscriptionId).Execute()

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
    sourceId := "rh5BDZp6QDHvXFCkibm1pR" // string | The sourceId for the Workspace.  This parameter exists in beta.
    destinationConfigId := "fP7qoQw2HTWt9WdMr718gn" // string | The id tied to a Workspace Destination.  This parameter exists in beta.
    startTime := "2024-01-01T00:00:00Z" // string | The ISO8601 formatted timestamp corresponding to the beginning of the requested timeframe, inclusive.  This parameter exists in beta.
    endTime := "2024-01-03T00:00:00Z" // string | The ISO8601 formatted timestamp corresponding to the end of the requested timeframe, noninclusive.  This parameter exists in beta.
    granularity := "day" // string | The size of each bucket in the requested window.  Based on the granularity chosen, there are restrictions on the time range you can query:  **Minute**: - Max time range: 4 hours - Oldest possible start time: 48 hours in the past  **Hour**: - Max Time range: 14 days - Oldest possible start time: 30 days in the past  **Day**: - Max time range: 30 days - Oldest possible start time: 30 days in the past  This parameter exists in beta.
    pagination := *api.NewPaginationInput(10) // PaginationInput | Params to specify the page cursor and count.  This parameter exists in beta.
    groupBy := []string{"Inner_example"} // []string | A comma-delimited list of strings representing one or more dimensions to group the result by.  Valid options are: `eventName`, `eventType`, `discardReason`, and `appVersion`.  This parameter exists in beta. (optional)
    filter := *api.NewDeliveryOverviewFilterBy() // DeliveryOverviewFilterBy | An optional filter for `eventName`, `eventType`, `discardReason`, and/or `appVersion` that can be applied in addition to a `groupBy`.  This parameter exists in beta. (optional)
    subscriptionId := "subscriptionId_example" // string | An optional filter for actions destinations, to filter by a specific action.  This parameter exists in beta. (optional)

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.DeliveryOverviewAPI.GetFilteredAtDestinationMetricsFromDeliveryOverview(ctx).SourceId(sourceId).DestinationConfigId(destinationConfigId).StartTime(startTime).EndTime(endTime).Granularity(granularity).Pagination(pagination).GroupBy(groupBy).Filter(filter).SubscriptionId(subscriptionId).Execute()
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



### Other Parameters

Other parameters are passed through a pointer to a apiGetFilteredAtDestinationMetricsFromDeliveryOverviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sourceId** | **string** | The sourceId for the Workspace.  This parameter exists in beta. | 
 **destinationConfigId** | **string** | The id tied to a Workspace Destination.  This parameter exists in beta. | 
 **startTime** | **string** | The ISO8601 formatted timestamp corresponding to the beginning of the requested timeframe, inclusive.  This parameter exists in beta. | 
 **endTime** | **string** | The ISO8601 formatted timestamp corresponding to the end of the requested timeframe, noninclusive.  This parameter exists in beta. | 
 **granularity** | **string** | The size of each bucket in the requested window.  Based on the granularity chosen, there are restrictions on the time range you can query:  **Minute**: - Max time range: 4 hours - Oldest possible start time: 48 hours in the past  **Hour**: - Max Time range: 14 days - Oldest possible start time: 30 days in the past  **Day**: - Max time range: 30 days - Oldest possible start time: 30 days in the past  This parameter exists in beta. | 
 **pagination** | [**PaginationInput**](PaginationInput.md) | Params to specify the page cursor and count.  This parameter exists in beta. | 
 **groupBy** | **[]string** | A comma-delimited list of strings representing one or more dimensions to group the result by.  Valid options are: &#x60;eventName&#x60;, &#x60;eventType&#x60;, &#x60;discardReason&#x60;, and &#x60;appVersion&#x60;.  This parameter exists in beta. | 
 **filter** | [**DeliveryOverviewFilterBy**](DeliveryOverviewFilterBy.md) | An optional filter for &#x60;eventName&#x60;, &#x60;eventType&#x60;, &#x60;discardReason&#x60;, and/or &#x60;appVersion&#x60; that can be applied in addition to a &#x60;groupBy&#x60;.  This parameter exists in beta. | 
 **subscriptionId** | **string** | An optional filter for actions destinations, to filter by a specific action.  This parameter exists in beta. | 

### Return type

[**GetEgressFailedMetricsFromDeliveryOverview200Response**](GetEgressFailedMetricsFromDeliveryOverview200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1beta+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: GetFilteredAtSourceMetricsFromDeliveryOverview

> GetEgressFailedMetricsFromDeliveryOverview200Response GetFilteredAtSourceMetricsFromDeliveryOverview(ctx).SourceId(sourceId).StartTime(startTime).EndTime(endTime).Granularity(granularity).Pagination(pagination).DestinationConfigId(destinationConfigId).GroupBy(groupBy).Filter(filter).Execute()

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
    sourceId := "rh5BDZp6QDHvXFCkibm1pR" // string | The sourceId for the Workspace.  This parameter exists in beta.
    startTime := "2024-01-01T00:00:00Z" // string | The ISO8601 formatted timestamp corresponding to the beginning of the requested timeframe, inclusive.  This parameter exists in beta.
    endTime := "2024-01-03T00:00:00Z" // string | The ISO8601 formatted timestamp corresponding to the end of the requested timeframe, noninclusive.  This parameter exists in beta.
    granularity := "day" // string | The size of each bucket in the requested window.  Based on the granularity chosen, there are restrictions on the time range you can query:  **Minute**: - Max time range: 4 hours - Oldest possible start time: 48 hours in the past  **Hour**: - Max Time range: 14 days - Oldest possible start time: 30 days in the past  **Day**: - Max time range: 30 days - Oldest possible start time: 30 days in the past  This parameter exists in beta.
    pagination := *api.NewPaginationInput(10) // PaginationInput | Optional params to specify the page cursor and count.  This parameter exists in beta.
    destinationConfigId := "destinationConfigId_example" // string | The id tied to a Workspace Destination.  This parameter exists in beta. (optional)
    groupBy := []string{"Inner_example"} // []string | A comma-delimited list of strings representing one or more dimensions to group the result by.  Valid options are: `eventName`, `eventType`, `discardReason`, and `appVersion`.  This parameter exists in beta. (optional)
    filter := *api.NewDeliveryOverviewFilterBy() // DeliveryOverviewFilterBy | An optional filter for `eventName`, `eventType`, `discardReason`, and/or `appVersion` that can be applied in addition to a `groupBy`.  This parameter exists in beta. (optional)

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.DeliveryOverviewAPI.GetFilteredAtSourceMetricsFromDeliveryOverview(ctx).SourceId(sourceId).StartTime(startTime).EndTime(endTime).Granularity(granularity).Pagination(pagination).DestinationConfigId(destinationConfigId).GroupBy(groupBy).Filter(filter).Execute()
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



### Other Parameters

Other parameters are passed through a pointer to a apiGetFilteredAtSourceMetricsFromDeliveryOverviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sourceId** | **string** | The sourceId for the Workspace.  This parameter exists in beta. | 
 **startTime** | **string** | The ISO8601 formatted timestamp corresponding to the beginning of the requested timeframe, inclusive.  This parameter exists in beta. | 
 **endTime** | **string** | The ISO8601 formatted timestamp corresponding to the end of the requested timeframe, noninclusive.  This parameter exists in beta. | 
 **granularity** | **string** | The size of each bucket in the requested window.  Based on the granularity chosen, there are restrictions on the time range you can query:  **Minute**: - Max time range: 4 hours - Oldest possible start time: 48 hours in the past  **Hour**: - Max Time range: 14 days - Oldest possible start time: 30 days in the past  **Day**: - Max time range: 30 days - Oldest possible start time: 30 days in the past  This parameter exists in beta. | 
 **pagination** | [**PaginationInput**](PaginationInput.md) | Optional params to specify the page cursor and count.  This parameter exists in beta. | 
 **destinationConfigId** | **string** | The id tied to a Workspace Destination.  This parameter exists in beta. | 
 **groupBy** | **[]string** | A comma-delimited list of strings representing one or more dimensions to group the result by.  Valid options are: &#x60;eventName&#x60;, &#x60;eventType&#x60;, &#x60;discardReason&#x60;, and &#x60;appVersion&#x60;.  This parameter exists in beta. | 
 **filter** | [**DeliveryOverviewFilterBy**](DeliveryOverviewFilterBy.md) | An optional filter for &#x60;eventName&#x60;, &#x60;eventType&#x60;, &#x60;discardReason&#x60;, and/or &#x60;appVersion&#x60; that can be applied in addition to a &#x60;groupBy&#x60;.  This parameter exists in beta. | 

### Return type

[**GetEgressFailedMetricsFromDeliveryOverview200Response**](GetEgressFailedMetricsFromDeliveryOverview200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1beta+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: GetIngressFailedMetricsFromDeliveryOverview

> GetEgressFailedMetricsFromDeliveryOverview200Response GetIngressFailedMetricsFromDeliveryOverview(ctx).SourceId(sourceId).StartTime(startTime).EndTime(endTime).Granularity(granularity).Pagination(pagination).DestinationConfigId(destinationConfigId).GroupBy(groupBy).Filter(filter).Execute()

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
    sourceId := "rh5BDZp6QDHvXFCkibm1pR" // string | The sourceId for the Workspace.  This parameter exists in beta.
    startTime := "2024-01-01T00:00:00Z" // string | The ISO8601 formatted timestamp corresponding to the beginning of the requested timeframe, inclusive.  This parameter exists in beta.
    endTime := "2024-01-03T00:00:00Z" // string | The ISO8601 formatted timestamp corresponding to the end of the requested timeframe, noninclusive.  This parameter exists in beta.
    granularity := "day" // string | The size of each bucket in the requested window.  Based on the granularity chosen, there are restrictions on the time range you can query:  **Minute**: - Max time range: 4 hours - Oldest possible start time: 48 hours in the past  **Hour**: - Max Time range: 14 days - Oldest possible start time: 30 days in the past  **Day**: - Max time range: 30 days - Oldest possible start time: 30 days in the past  This parameter exists in beta.
    pagination := *api.NewPaginationInput(10) // PaginationInput | Optional params to specify the page cursor and count.  This parameter exists in beta.
    destinationConfigId := "destinationConfigId_example" // string | The id tied to a Workspace Destination.  This parameter exists in beta. (optional)
    groupBy := []string{"Inner_example"} // []string | A comma-delimited list of strings representing one or more dimensions to group the result by.  Valid options are: `eventName`, `eventType`, `discardReason`, and `appVersion`.  This parameter exists in beta. (optional)
    filter := *api.NewDeliveryOverviewFilterBy() // DeliveryOverviewFilterBy | An optional filter for `eventName`, `eventType`, `discardReason`, and/or `appVersion` that can be applied in addition to a `groupBy`.  This parameter exists in beta. (optional)

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.DeliveryOverviewAPI.GetIngressFailedMetricsFromDeliveryOverview(ctx).SourceId(sourceId).StartTime(startTime).EndTime(endTime).Granularity(granularity).Pagination(pagination).DestinationConfigId(destinationConfigId).GroupBy(groupBy).Filter(filter).Execute()
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



### Other Parameters

Other parameters are passed through a pointer to a apiGetIngressFailedMetricsFromDeliveryOverviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sourceId** | **string** | The sourceId for the Workspace.  This parameter exists in beta. | 
 **startTime** | **string** | The ISO8601 formatted timestamp corresponding to the beginning of the requested timeframe, inclusive.  This parameter exists in beta. | 
 **endTime** | **string** | The ISO8601 formatted timestamp corresponding to the end of the requested timeframe, noninclusive.  This parameter exists in beta. | 
 **granularity** | **string** | The size of each bucket in the requested window.  Based on the granularity chosen, there are restrictions on the time range you can query:  **Minute**: - Max time range: 4 hours - Oldest possible start time: 48 hours in the past  **Hour**: - Max Time range: 14 days - Oldest possible start time: 30 days in the past  **Day**: - Max time range: 30 days - Oldest possible start time: 30 days in the past  This parameter exists in beta. | 
 **pagination** | [**PaginationInput**](PaginationInput.md) | Optional params to specify the page cursor and count.  This parameter exists in beta. | 
 **destinationConfigId** | **string** | The id tied to a Workspace Destination.  This parameter exists in beta. | 
 **groupBy** | **[]string** | A comma-delimited list of strings representing one or more dimensions to group the result by.  Valid options are: &#x60;eventName&#x60;, &#x60;eventType&#x60;, &#x60;discardReason&#x60;, and &#x60;appVersion&#x60;.  This parameter exists in beta. | 
 **filter** | [**DeliveryOverviewFilterBy**](DeliveryOverviewFilterBy.md) | An optional filter for &#x60;eventName&#x60;, &#x60;eventType&#x60;, &#x60;discardReason&#x60;, and/or &#x60;appVersion&#x60; that can be applied in addition to a &#x60;groupBy&#x60;.  This parameter exists in beta. | 

### Return type

[**GetEgressFailedMetricsFromDeliveryOverview200Response**](GetEgressFailedMetricsFromDeliveryOverview200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1beta+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: GetIngressSuccessMetricsFromDeliveryOverview

> GetEgressFailedMetricsFromDeliveryOverview200Response GetIngressSuccessMetricsFromDeliveryOverview(ctx).SourceId(sourceId).StartTime(startTime).EndTime(endTime).Granularity(granularity).Pagination(pagination).DestinationConfigId(destinationConfigId).GroupBy(groupBy).Filter(filter).Execute()

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
    sourceId := "rh5BDZp6QDHvXFCkibm1pR" // string | The sourceId for the Workspace.  This parameter exists in beta.
    startTime := "2024-01-01T00:00:00Z" // string | The ISO8601 formatted timestamp corresponding to the beginning of the requested timeframe, inclusive.  This parameter exists in beta.
    endTime := "2024-01-03T00:00:00Z" // string | The ISO8601 formatted timestamp corresponding to the end of the requested timeframe, noninclusive.  This parameter exists in beta.
    granularity := "day" // string | The size of each bucket in the requested window.  Based on the granularity chosen, there are restrictions on the time range you can query:  **Minute**: - Max time range: 4 hours - Oldest possible start time: 48 hours in the past  **Hour**: - Max Time range: 14 days - Oldest possible start time: 30 days in the past  **Day**: - Max time range: 30 days - Oldest possible start time: 30 days in the past  This parameter exists in beta.
    pagination := *api.NewPaginationInput(10) // PaginationInput | Optional params to specify the page cursor and count.  This parameter exists in beta.
    destinationConfigId := "destinationConfigId_example" // string | The id tied to a Workspace Destination.  This parameter exists in beta. (optional)
    groupBy := []string{"Inner_example"} // []string | A comma-delimited list of strings representing one or more dimensions to group the result by.  Valid options are: `eventName`, `eventType`, and `appVersion`.  This parameter exists in beta. (optional)
    filter := *api.NewDeliveryOverviewFilterBy() // DeliveryOverviewFilterBy | An optional filter for `eventName`, `eventType`, and/or `appVersion` that can be applied in addition to a `groupBy`.  This parameter exists in beta. (optional)

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.DeliveryOverviewAPI.GetIngressSuccessMetricsFromDeliveryOverview(ctx).SourceId(sourceId).StartTime(startTime).EndTime(endTime).Granularity(granularity).Pagination(pagination).DestinationConfigId(destinationConfigId).GroupBy(groupBy).Filter(filter).Execute()
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



### Other Parameters

Other parameters are passed through a pointer to a apiGetIngressSuccessMetricsFromDeliveryOverviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sourceId** | **string** | The sourceId for the Workspace.  This parameter exists in beta. | 
 **startTime** | **string** | The ISO8601 formatted timestamp corresponding to the beginning of the requested timeframe, inclusive.  This parameter exists in beta. | 
 **endTime** | **string** | The ISO8601 formatted timestamp corresponding to the end of the requested timeframe, noninclusive.  This parameter exists in beta. | 
 **granularity** | **string** | The size of each bucket in the requested window.  Based on the granularity chosen, there are restrictions on the time range you can query:  **Minute**: - Max time range: 4 hours - Oldest possible start time: 48 hours in the past  **Hour**: - Max Time range: 14 days - Oldest possible start time: 30 days in the past  **Day**: - Max time range: 30 days - Oldest possible start time: 30 days in the past  This parameter exists in beta. | 
 **pagination** | [**PaginationInput**](PaginationInput.md) | Optional params to specify the page cursor and count.  This parameter exists in beta. | 
 **destinationConfigId** | **string** | The id tied to a Workspace Destination.  This parameter exists in beta. | 
 **groupBy** | **[]string** | A comma-delimited list of strings representing one or more dimensions to group the result by.  Valid options are: &#x60;eventName&#x60;, &#x60;eventType&#x60;, and &#x60;appVersion&#x60;.  This parameter exists in beta. | 
 **filter** | [**DeliveryOverviewFilterBy**](DeliveryOverviewFilterBy.md) | An optional filter for &#x60;eventName&#x60;, &#x60;eventType&#x60;, and/or &#x60;appVersion&#x60; that can be applied in addition to a &#x60;groupBy&#x60;.  This parameter exists in beta. | 

### Return type

[**GetEgressFailedMetricsFromDeliveryOverview200Response**](GetEgressFailedMetricsFromDeliveryOverview200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1beta+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

