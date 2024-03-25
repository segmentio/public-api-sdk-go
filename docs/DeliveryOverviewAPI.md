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

> GetEgressFailedMetricsFromDeliveryOverview200Response GetEgressFailedMetricsFromDeliveryOverview(ctx).Metrics(metrics).Execute()

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
    metrics := *api.NewGetDeliveryOverviewDestMetricsBetaInput("SourceId_example", "DestinationConfigId_example", "StartTime_example", "EndTime_example", "Granularity_example", *api.NewPaginationInput(10)) // GetDeliveryOverviewDestMetricsBetaInput | Metrics for this Destination pipeline step.  This parameter exists in beta.

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.DeliveryOverviewAPI.GetEgressFailedMetricsFromDeliveryOverview(ctx).Metrics(metrics).Execute()
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
 **metrics** | [**GetDeliveryOverviewDestMetricsBetaInput**](GetDeliveryOverviewDestMetricsBetaInput.md) | Metrics for this Destination pipeline step.  This parameter exists in beta. | 

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

> GetEgressFailedMetricsFromDeliveryOverview200Response GetEgressSuccessMetricsFromDeliveryOverview(ctx).Metrics(metrics).Execute()

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
    metrics := *api.NewGetDeliveryOverviewDestMetricsBetaInput("SourceId_example", "DestinationConfigId_example", "StartTime_example", "EndTime_example", "Granularity_example", *api.NewPaginationInput(10)) // GetDeliveryOverviewDestMetricsBetaInput | Metrics for this Destination pipeline step.  This parameter exists in beta.

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.DeliveryOverviewAPI.GetEgressSuccessMetricsFromDeliveryOverview(ctx).Metrics(metrics).Execute()
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
 **metrics** | [**GetDeliveryOverviewDestMetricsBetaInput**](GetDeliveryOverviewDestMetricsBetaInput.md) | Metrics for this Destination pipeline step.  This parameter exists in beta. | 

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

> GetEgressFailedMetricsFromDeliveryOverview200Response GetFilteredAtDestinationMetricsFromDeliveryOverview(ctx).Metrics(metrics).Execute()

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
    metrics := *api.NewGetDeliveryOverviewDestMetricsBetaInput("SourceId_example", "DestinationConfigId_example", "StartTime_example", "EndTime_example", "Granularity_example", *api.NewPaginationInput(10)) // GetDeliveryOverviewDestMetricsBetaInput | Metrics for this Destination pipeline step.  This parameter exists in beta.

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.DeliveryOverviewAPI.GetFilteredAtDestinationMetricsFromDeliveryOverview(ctx).Metrics(metrics).Execute()
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
 **metrics** | [**GetDeliveryOverviewDestMetricsBetaInput**](GetDeliveryOverviewDestMetricsBetaInput.md) | Metrics for this Destination pipeline step.  This parameter exists in beta. | 

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

> GetEgressFailedMetricsFromDeliveryOverview200Response GetFilteredAtSourceMetricsFromDeliveryOverview(ctx).Metrics(metrics).Execute()

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
    metrics := *api.NewGetDeliveryOverviewSourceMetricsBetaInput("SourceId_example", "StartTime_example", "EndTime_example", "Granularity_example", *api.NewPaginationInput(10)) // GetDeliveryOverviewSourceMetricsBetaInput | Metrics for this Source pipeline step.  This parameter exists in beta.

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.DeliveryOverviewAPI.GetFilteredAtSourceMetricsFromDeliveryOverview(ctx).Metrics(metrics).Execute()
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
 **metrics** | [**GetDeliveryOverviewSourceMetricsBetaInput**](GetDeliveryOverviewSourceMetricsBetaInput.md) | Metrics for this Source pipeline step.  This parameter exists in beta. | 

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

> GetEgressFailedMetricsFromDeliveryOverview200Response GetIngressFailedMetricsFromDeliveryOverview(ctx).Metrics(metrics).Execute()

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
    metrics := *api.NewGetDeliveryOverviewSourceMetricsBetaInput("SourceId_example", "StartTime_example", "EndTime_example", "Granularity_example", *api.NewPaginationInput(10)) // GetDeliveryOverviewSourceMetricsBetaInput | Metrics for this Source pipeline step.  This parameter exists in beta.

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.DeliveryOverviewAPI.GetIngressFailedMetricsFromDeliveryOverview(ctx).Metrics(metrics).Execute()
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
 **metrics** | [**GetDeliveryOverviewSourceMetricsBetaInput**](GetDeliveryOverviewSourceMetricsBetaInput.md) | Metrics for this Source pipeline step.  This parameter exists in beta. | 

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

> GetEgressFailedMetricsFromDeliveryOverview200Response GetIngressSuccessMetricsFromDeliveryOverview(ctx).Metrics(metrics).Execute()

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
    metrics := *api.NewGetDeliveryOverviewSourceMetricsBetaInput("SourceId_example", "StartTime_example", "EndTime_example", "Granularity_example", *api.NewPaginationInput(10)) // GetDeliveryOverviewSourceMetricsBetaInput | Metrics for this Source pipeline step.  This parameter exists in beta.

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.DeliveryOverviewAPI.GetIngressSuccessMetricsFromDeliveryOverview(ctx).Metrics(metrics).Execute()
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
 **metrics** | [**GetDeliveryOverviewSourceMetricsBetaInput**](GetDeliveryOverviewSourceMetricsBetaInput.md) | Metrics for this Source pipeline step.  This parameter exists in beta. | 

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

