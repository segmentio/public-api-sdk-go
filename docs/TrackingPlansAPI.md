# \TrackingPlansAPI

All URIs are relative to *https://api.segmentapis.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddSourceToTrackingPlan**](TrackingPlansAPI.md#AddSourceToTrackingPlan) | **Post** /tracking-plans/{trackingPlanId}/sources | Add Source to Tracking Plan
[**CreateTrackingPlan**](TrackingPlansAPI.md#CreateTrackingPlan) | **Post** /tracking-plans | Create Tracking Plan
[**DeleteTrackingPlan**](TrackingPlansAPI.md#DeleteTrackingPlan) | **Delete** /tracking-plans/{trackingPlanId} | Delete Tracking Plan
[**GetTrackingPlan**](TrackingPlansAPI.md#GetTrackingPlan) | **Get** /tracking-plans/{trackingPlanId} | Get Tracking Plan
[**ListRulesFromTrackingPlan**](TrackingPlansAPI.md#ListRulesFromTrackingPlan) | **Get** /tracking-plans/{trackingPlanId}/rules | List Rules from Tracking Plan
[**ListSourcesFromTrackingPlan**](TrackingPlansAPI.md#ListSourcesFromTrackingPlan) | **Get** /tracking-plans/{trackingPlanId}/sources | List Sources from Tracking Plan
[**ListTrackingPlans**](TrackingPlansAPI.md#ListTrackingPlans) | **Get** /tracking-plans | List Tracking Plans
[**RemoveRulesFromTrackingPlan**](TrackingPlansAPI.md#RemoveRulesFromTrackingPlan) | **Delete** /tracking-plans/{trackingPlanId}/rules | Remove Rules from Tracking Plan
[**RemoveSourceFromTrackingPlan**](TrackingPlansAPI.md#RemoveSourceFromTrackingPlan) | **Delete** /tracking-plans/{trackingPlanId}/sources | Remove Source from Tracking Plan
[**ReplaceRulesInTrackingPlan**](TrackingPlansAPI.md#ReplaceRulesInTrackingPlan) | **Put** /tracking-plans/{trackingPlanId}/rules | Replace Rules in Tracking Plan
[**UpdateRulesInTrackingPlan**](TrackingPlansAPI.md#UpdateRulesInTrackingPlan) | **Patch** /tracking-plans/{trackingPlanId}/rules | Update Rules in Tracking Plan
[**UpdateTrackingPlan**](TrackingPlansAPI.md#UpdateTrackingPlan) | **Patch** /tracking-plans/{trackingPlanId} | Update Tracking Plan



## Operation: AddSourceToTrackingPlan

> AddSourceToTrackingPlan200Response AddSourceToTrackingPlan(ctx, trackingPlanId).AddSourceToTrackingPlanV1Input(addSourceToTrackingPlanV1Input).Execute()

Add Source to Tracking Plan



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
    trackingPlanId := "tp_sprout_rVGCC6WdrNxjCf6JpCHP" // string | 
    addSourceToTrackingPlanV1Input := *api.NewAddSourceToTrackingPlanV1Input("SourceId_example") // AddSourceToTrackingPlanV1Input | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.TrackingPlansAPI.AddSourceToTrackingPlan(ctx).AddSourceToTrackingPlanV1Input(addSourceToTrackingPlanV1Input).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrackingPlansAPI.AddSourceToTrackingPlan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `AddSourceToTrackingPlan`: AddSourceToTrackingPlan200Response
    fmt.Fprintf(os.Stdout, "Response from `TrackingPlansAPI.AddSourceToTrackingPlan`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**trackingPlanId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddSourceToTrackingPlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addSourceToTrackingPlanV1Input** | [**AddSourceToTrackingPlanV1Input**](AddSourceToTrackingPlanV1Input.md) |  | 

### Return type

[**AddSourceToTrackingPlan200Response**](AddSourceToTrackingPlan200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/vnd.segment.v1+json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json
- **Accept**: application/vnd.segment.v1+json, application/json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: CreateTrackingPlan

> CreateTrackingPlan200Response CreateTrackingPlan(ctx).CreateTrackingPlanV1Input(createTrackingPlanV1Input).Execute()

Create Tracking Plan



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
    createTrackingPlanV1Input := *api.NewCreateTrackingPlanV1Input("Name_example", "Type_example") // CreateTrackingPlanV1Input | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.TrackingPlansAPI.CreateTrackingPlan(ctx).CreateTrackingPlanV1Input(createTrackingPlanV1Input).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrackingPlansAPI.CreateTrackingPlan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `CreateTrackingPlan`: CreateTrackingPlan200Response
    fmt.Fprintf(os.Stdout, "Response from `TrackingPlansAPI.CreateTrackingPlan`: %v\n", resp.GetData())
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTrackingPlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTrackingPlanV1Input** | [**CreateTrackingPlanV1Input**](CreateTrackingPlanV1Input.md) |  | 

### Return type

[**CreateTrackingPlan200Response**](CreateTrackingPlan200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/vnd.segment.v1+json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json
- **Accept**: application/vnd.segment.v1+json, application/json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: DeleteTrackingPlan

> DeleteTrackingPlan200Response DeleteTrackingPlan(ctx, trackingPlanId).Execute()

Delete Tracking Plan



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
    trackingPlanId := "tp_sprout_rVGCC6WdrNxjCf6JpCHP" // string | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.TrackingPlansAPI.DeleteTrackingPlan(ctx).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrackingPlansAPI.DeleteTrackingPlan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `DeleteTrackingPlan`: DeleteTrackingPlan200Response
    fmt.Fprintf(os.Stdout, "Response from `TrackingPlansAPI.DeleteTrackingPlan`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**trackingPlanId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTrackingPlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteTrackingPlan200Response**](DeleteTrackingPlan200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1+json, application/json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: GetTrackingPlan

> GetTrackingPlan200Response GetTrackingPlan(ctx, trackingPlanId).Execute()

Get Tracking Plan



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
    trackingPlanId := "tp_sprout_rVGCC6WdrNxjCf6JpCHP" // string | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.TrackingPlansAPI.GetTrackingPlan(ctx).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrackingPlansAPI.GetTrackingPlan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `GetTrackingPlan`: GetTrackingPlan200Response
    fmt.Fprintf(os.Stdout, "Response from `TrackingPlansAPI.GetTrackingPlan`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**trackingPlanId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTrackingPlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetTrackingPlan200Response**](GetTrackingPlan200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1+json, application/json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: ListRulesFromTrackingPlan

> ListRulesFromTrackingPlan200Response ListRulesFromTrackingPlan(ctx, trackingPlanId).Pagination(pagination).Execute()

List Rules from Tracking Plan



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
    trackingPlanId := "tp_sprout_rVGCC6WdrNxjCf6JpCHP" // string | 
    pagination := *api.NewPaginationInput(float32(123)) // PaginationInput | Pagination options.  This parameter exists in v1.

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.TrackingPlansAPI.ListRulesFromTrackingPlan(ctx).Pagination(pagination).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrackingPlansAPI.ListRulesFromTrackingPlan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `ListRulesFromTrackingPlan`: ListRulesFromTrackingPlan200Response
    fmt.Fprintf(os.Stdout, "Response from `TrackingPlansAPI.ListRulesFromTrackingPlan`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**trackingPlanId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRulesFromTrackingPlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pagination** | [**PaginationInput**](PaginationInput.md) | Pagination options.  This parameter exists in v1. | 

### Return type

[**ListRulesFromTrackingPlan200Response**](ListRulesFromTrackingPlan200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1+json, application/json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: ListSourcesFromTrackingPlan

> ListSourcesFromTrackingPlan200Response ListSourcesFromTrackingPlan(ctx, trackingPlanId).Pagination(pagination).Execute()

List Sources from Tracking Plan



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
    trackingPlanId := "tp_sprout_rVGCC6WdrNxjCf6JpCHP" // string | 
    pagination := *api.NewPaginationInput(float32(123)) // PaginationInput | Pagination options.  This parameter exists in v1.

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.TrackingPlansAPI.ListSourcesFromTrackingPlan(ctx).Pagination(pagination).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrackingPlansAPI.ListSourcesFromTrackingPlan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `ListSourcesFromTrackingPlan`: ListSourcesFromTrackingPlan200Response
    fmt.Fprintf(os.Stdout, "Response from `TrackingPlansAPI.ListSourcesFromTrackingPlan`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**trackingPlanId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSourcesFromTrackingPlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pagination** | [**PaginationInput**](PaginationInput.md) | Pagination options.  This parameter exists in v1. | 

### Return type

[**ListSourcesFromTrackingPlan200Response**](ListSourcesFromTrackingPlan200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1+json, application/json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: ListTrackingPlans

> ListTrackingPlans200Response ListTrackingPlans(ctx).Pagination(pagination).Type_(type_).Execute()

List Tracking Plans



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
    pagination := *api.NewPaginationInput(float32(123)) // PaginationInput | Pagination options.  This parameter exists in v1.
    type_ := "LIVE" // string | Requests Tracking Plans of a certain type. If omitted, lists all types.  This parameter exists in v1. (optional)

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.TrackingPlansAPI.ListTrackingPlans(ctx).Pagination(pagination).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrackingPlansAPI.ListTrackingPlans``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `ListTrackingPlans`: ListTrackingPlans200Response
    fmt.Fprintf(os.Stdout, "Response from `TrackingPlansAPI.ListTrackingPlans`: %v\n", resp.GetData())
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTrackingPlansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pagination** | [**PaginationInput**](PaginationInput.md) | Pagination options.  This parameter exists in v1. | 
 **type_** | **string** | Requests Tracking Plans of a certain type. If omitted, lists all types.  This parameter exists in v1. | 

### Return type

[**ListTrackingPlans200Response**](ListTrackingPlans200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1+json, application/json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: RemoveRulesFromTrackingPlan

> RemoveRulesFromTrackingPlan200Response RemoveRulesFromTrackingPlan(ctx, trackingPlanId).Rules(rules).Execute()

Remove Rules from Tracking Plan



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
    trackingPlanId := "tp_sprout_rVGCC6WdrNxjCf6JpCHP" // string | 
    rules := []api.RemoveRuleV1{[]api.RemoveRuleV1{*api.NewRemoveRuleV1("Type_example", float32(123))}} // []RemoveRuleV1 | Rules to delete.  This parameter exists in v1.

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.TrackingPlansAPI.RemoveRulesFromTrackingPlan(ctx).Rules(rules).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrackingPlansAPI.RemoveRulesFromTrackingPlan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `RemoveRulesFromTrackingPlan`: RemoveRulesFromTrackingPlan200Response
    fmt.Fprintf(os.Stdout, "Response from `TrackingPlansAPI.RemoveRulesFromTrackingPlan`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**trackingPlanId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveRulesFromTrackingPlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **rules** | [**[][]RemoveRuleV1**](array.md) | Rules to delete.  This parameter exists in v1. | 

### Return type

[**RemoveRulesFromTrackingPlan200Response**](RemoveRulesFromTrackingPlan200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1+json, application/json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: RemoveSourceFromTrackingPlan

> RemoveSourceFromTrackingPlan200Response RemoveSourceFromTrackingPlan(ctx, trackingPlanId).SourceId(sourceId).Execute()

Remove Source from Tracking Plan



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
    trackingPlanId := "tp_sprout_rVGCC6WdrNxjCf6JpCHP" // string | 
    sourceId := "qQEHquLrjRDN9j1ByrChyn" // string | The id of the Source associated with the Tracking Plan.  Config API note: analogous to `sourceName`.  This parameter exists in v1.

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.TrackingPlansAPI.RemoveSourceFromTrackingPlan(ctx).SourceId(sourceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrackingPlansAPI.RemoveSourceFromTrackingPlan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `RemoveSourceFromTrackingPlan`: RemoveSourceFromTrackingPlan200Response
    fmt.Fprintf(os.Stdout, "Response from `TrackingPlansAPI.RemoveSourceFromTrackingPlan`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**trackingPlanId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveSourceFromTrackingPlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sourceId** | **string** | The id of the Source associated with the Tracking Plan.  Config API note: analogous to &#x60;sourceName&#x60;.  This parameter exists in v1. | 

### Return type

[**RemoveSourceFromTrackingPlan200Response**](RemoveSourceFromTrackingPlan200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1+json, application/json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: ReplaceRulesInTrackingPlan

> ReplaceRulesInTrackingPlan200Response ReplaceRulesInTrackingPlan(ctx, trackingPlanId).ReplaceRulesInTrackingPlanV1Input(replaceRulesInTrackingPlanV1Input).Execute()

Replace Rules in Tracking Plan



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
    trackingPlanId := "tp_sprout_rVGCC6WdrNxjCf6JpCHP" // string | 
    replaceRulesInTrackingPlanV1Input := *api.NewReplaceRulesInTrackingPlanV1Input([]api.RuleInputV1{*api.NewRuleInputV1("Type_example", interface{}(123), float32(123))}) // ReplaceRulesInTrackingPlanV1Input | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.TrackingPlansAPI.ReplaceRulesInTrackingPlan(ctx).ReplaceRulesInTrackingPlanV1Input(replaceRulesInTrackingPlanV1Input).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrackingPlansAPI.ReplaceRulesInTrackingPlan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `ReplaceRulesInTrackingPlan`: ReplaceRulesInTrackingPlan200Response
    fmt.Fprintf(os.Stdout, "Response from `TrackingPlansAPI.ReplaceRulesInTrackingPlan`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**trackingPlanId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceRulesInTrackingPlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **replaceRulesInTrackingPlanV1Input** | [**ReplaceRulesInTrackingPlanV1Input**](ReplaceRulesInTrackingPlanV1Input.md) |  | 

### Return type

[**ReplaceRulesInTrackingPlan200Response**](ReplaceRulesInTrackingPlan200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/vnd.segment.v1+json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json
- **Accept**: application/vnd.segment.v1+json, application/json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: UpdateRulesInTrackingPlan

> UpdateRulesInTrackingPlan200Response UpdateRulesInTrackingPlan(ctx, trackingPlanId).UpdateRulesInTrackingPlanV1Input(updateRulesInTrackingPlanV1Input).Execute()

Update Rules in Tracking Plan



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
    trackingPlanId := "tp_sprout_rVGCC6WdrNxjCf6JpCHP" // string | 
    updateRulesInTrackingPlanV1Input := *api.NewUpdateRulesInTrackingPlanV1Input([]api.UpsertRuleV1{*api.NewUpsertRuleV1("Type_example", interface{}(123), float32(123))}) // UpdateRulesInTrackingPlanV1Input | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.TrackingPlansAPI.UpdateRulesInTrackingPlan(ctx).UpdateRulesInTrackingPlanV1Input(updateRulesInTrackingPlanV1Input).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrackingPlansAPI.UpdateRulesInTrackingPlan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `UpdateRulesInTrackingPlan`: UpdateRulesInTrackingPlan200Response
    fmt.Fprintf(os.Stdout, "Response from `TrackingPlansAPI.UpdateRulesInTrackingPlan`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**trackingPlanId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRulesInTrackingPlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRulesInTrackingPlanV1Input** | [**UpdateRulesInTrackingPlanV1Input**](UpdateRulesInTrackingPlanV1Input.md) |  | 

### Return type

[**UpdateRulesInTrackingPlan200Response**](UpdateRulesInTrackingPlan200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/vnd.segment.v1+json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json
- **Accept**: application/vnd.segment.v1+json, application/json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: UpdateTrackingPlan

> UpdateTrackingPlan200Response UpdateTrackingPlan(ctx, trackingPlanId).UpdateTrackingPlanV1Input(updateTrackingPlanV1Input).Execute()

Update Tracking Plan



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
    trackingPlanId := "tp_sprout_rVGCC6WdrNxjCf6JpCHP" // string | 
    updateTrackingPlanV1Input := *api.NewUpdateTrackingPlanV1Input() // UpdateTrackingPlanV1Input | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.TrackingPlansAPI.UpdateTrackingPlan(ctx).UpdateTrackingPlanV1Input(updateTrackingPlanV1Input).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrackingPlansAPI.UpdateTrackingPlan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `UpdateTrackingPlan`: UpdateTrackingPlan200Response
    fmt.Fprintf(os.Stdout, "Response from `TrackingPlansAPI.UpdateTrackingPlan`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**trackingPlanId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTrackingPlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateTrackingPlanV1Input** | [**UpdateTrackingPlanV1Input**](UpdateTrackingPlanV1Input.md) |  | 

### Return type

[**UpdateTrackingPlan200Response**](UpdateTrackingPlan200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/vnd.segment.v1+json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json
- **Accept**: application/vnd.segment.v1+json, application/json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

