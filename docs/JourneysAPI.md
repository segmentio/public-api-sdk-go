# \JourneysAPI

All URIs are relative to *https://api.segmentapis.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateJourney**](JourneysAPI.md#CreateJourney) | **Post** /spaces/{spaceId}/journeys | Create Journey
[**GetJourney**](JourneysAPI.md#GetJourney) | **Get** /spaces/{spaceId}/journeys/{containerId} | Get Journey
[**GetJourneyAnalytics**](JourneysAPI.md#GetJourneyAnalytics) | **Get** /spaces/{spaceId}/journeys/{containerId}/analytics | Get Journey Analytics
[**RemoveJourneyFromSpace**](JourneysAPI.md#RemoveJourneyFromSpace) | **Delete** /spaces/{spaceId}/journeys/{containerId} | Remove Journey from Space
[**ReplaceStepsForJourney**](JourneysAPI.md#ReplaceStepsForJourney) | **Put** /spaces/{spaceId}/journeys/{containerId}/steps | Replace Steps for Journey
[**UpdateDestinationsForJourney**](JourneysAPI.md#UpdateDestinationsForJourney) | **Patch** /spaces/{spaceId}/journeys/{containerId}/destinations | Update Destinations for Journey
[**UpdateStatusForJourney**](JourneysAPI.md#UpdateStatusForJourney) | **Patch** /spaces/{spaceId}/journeys/{containerId}/status | Update Status for Journey



## Operation: CreateJourney

> CreateJourney201Response CreateJourney(ctx, spaceId).CreateJourneyAlphaInput(createJourneyAlphaInput).Execute()

Create Journey



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
    spaceId := "space_id" // string | 
    createJourneyAlphaInput := *api.NewCreateJourneyAlphaInput("Name_example", "Description_example", *api.NewJourneyDefinition("InitialState_example", *api.NewEntryRules("Type_example"), *api.NewExitRulesConfig(false, []api.RulesInner{*api.NewRulesInner("ExitType_example", "Condition_example", "Type_example", false, false, *api.NewKey("Id_example"), "AudienceId_example")}), []api.JourneyState{*api.NewJourneyState("Type_example", "Condition_example", []api.RandomSplitBranch{*api.NewRandomSplitBranch(float32(123), "Next_example")}, *api.NewKey("Id_example"), "AudienceId_example", []api.Destination{*api.NewDestination("DestinationId_example", []api.Activation{*api.NewActivation("EventName_example", *api.NewActionDefinition(map[string]interface{}{"key": interface{}(123)}, "ActionId_example"))})}, "Duration_example")})) // CreateJourneyAlphaInput | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.JourneysAPI.CreateJourney(ctx, spaceId).CreateJourneyAlphaInput(createJourneyAlphaInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JourneysAPI.CreateJourney``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `CreateJourney`: CreateJourney201Response
    fmt.Fprintf(os.Stdout, "Response from `JourneysAPI.CreateJourney`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateJourneyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createJourneyAlphaInput** | [**CreateJourneyAlphaInput**](CreateJourneyAlphaInput.md) |  | 

### Return type

[**CreateJourney201Response**](CreateJourney201Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/vnd.segment.v1alpha+json
- **Accept**: application/vnd.segment.v1alpha+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: GetJourney

> GetJourney200Response GetJourney(ctx, spaceId, containerId).Execute()

Get Journey



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
    spaceId := "space_id" // string | 
    containerId := "jcon_2tG95HZh4oPsgHfcOlznyfcDDAg" // string | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.JourneysAPI.GetJourney(ctx, spaceId, containerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JourneysAPI.GetJourney``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `GetJourney`: GetJourney200Response
    fmt.Fprintf(os.Stdout, "Response from `JourneysAPI.GetJourney`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** |  | 
**containerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetJourneyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetJourney200Response**](GetJourney200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1alpha+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: GetJourneyAnalytics

> GetJourneyAnalytics200Response GetJourneyAnalytics(ctx, spaceId, containerId).Version(version).FromDate(fromDate).ToDate(toDate).Execute()

Get Journey Analytics



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
    containerId := "jcon_2tG95HZh4oPsgHfcOlznyfcDDAg" // string | 
    version := float32(1) // float32 | The journey version.  This parameter exists in alpha.
    fromDate := "2006-01-02T15:04:05.000Z" // string | This parameter exists in alpha.
    toDate := "2006-01-02T15:04:05.000Z" // string | This parameter exists in alpha. (optional)

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.JourneysAPI.GetJourneyAnalytics(ctx, spaceId, containerId).Version(version).FromDate(fromDate).ToDate(toDate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JourneysAPI.GetJourneyAnalytics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `GetJourneyAnalytics`: GetJourneyAnalytics200Response
    fmt.Fprintf(os.Stdout, "Response from `JourneysAPI.GetJourneyAnalytics`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** |  | 
**containerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetJourneyAnalyticsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **version** | **float32** | The journey version.  This parameter exists in alpha. | 
 **fromDate** | **string** | This parameter exists in alpha. | 
 **toDate** | **string** | This parameter exists in alpha. | 

### Return type

[**GetJourneyAnalytics200Response**](GetJourneyAnalytics200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1alpha+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: RemoveJourneyFromSpace

> RemoveJourneyFromSpace200Response RemoveJourneyFromSpace(ctx, spaceId, containerId).Version(version).Execute()

Remove Journey from Space



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
    containerId := "jcon_2tG95HZh4oPsgHfcOlznyfcDDAg" // string | 
    version := float32(1) // float32 | The journey version.  This parameter exists in alpha.

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.JourneysAPI.RemoveJourneyFromSpace(ctx, spaceId, containerId).Version(version).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JourneysAPI.RemoveJourneyFromSpace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `RemoveJourneyFromSpace`: RemoveJourneyFromSpace200Response
    fmt.Fprintf(os.Stdout, "Response from `JourneysAPI.RemoveJourneyFromSpace`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** |  | 
**containerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveJourneyFromSpaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **version** | **float32** | The journey version.  This parameter exists in alpha. | 

### Return type

[**RemoveJourneyFromSpace200Response**](RemoveJourneyFromSpace200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1alpha+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: ReplaceStepsForJourney

> ReplaceStepsForJourney200Response ReplaceStepsForJourney(ctx, spaceId, containerId).ReplaceStepsForJourneyAlphaInput(replaceStepsForJourneyAlphaInput).Execute()

Replace Steps for Journey



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
    containerId := "jcon_2tG95HZh4oPsgHfcOlznyfcDDAg" // string | 
    replaceStepsForJourneyAlphaInput := *api.NewReplaceStepsForJourneyAlphaInput("InitialState_example", []api.JourneyState{*api.NewJourneyState("Type_example", "Condition_example", []api.RandomSplitBranch{*api.NewRandomSplitBranch(float32(123), "Next_example")}, *api.NewKey("Id_example"), "AudienceId_example", []api.Destination{*api.NewDestination("DestinationId_example", []api.Activation{*api.NewActivation("EventName_example", *api.NewActionDefinition(map[string]interface{}{"key": interface{}(123)}, "ActionId_example"))})}, "Duration_example")}) // ReplaceStepsForJourneyAlphaInput | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.JourneysAPI.ReplaceStepsForJourney(ctx, spaceId, containerId).ReplaceStepsForJourneyAlphaInput(replaceStepsForJourneyAlphaInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JourneysAPI.ReplaceStepsForJourney``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `ReplaceStepsForJourney`: ReplaceStepsForJourney200Response
    fmt.Fprintf(os.Stdout, "Response from `JourneysAPI.ReplaceStepsForJourney`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** |  | 
**containerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceStepsForJourneyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **replaceStepsForJourneyAlphaInput** | [**ReplaceStepsForJourneyAlphaInput**](ReplaceStepsForJourneyAlphaInput.md) |  | 

### Return type

[**ReplaceStepsForJourney200Response**](ReplaceStepsForJourney200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/vnd.segment.v1alpha+json
- **Accept**: application/vnd.segment.v1alpha+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: UpdateDestinationsForJourney

> UpdateDestinationsForJourney200Response UpdateDestinationsForJourney(ctx, spaceId, containerId).UpdateDestinationsForJourneyAlphaInput(updateDestinationsForJourneyAlphaInput).Execute()

Update Destinations for Journey



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
    containerId := "jcon_2tG95HZh4oPsgHfcOlznyfcDDAg" // string | 
    updateDestinationsForJourneyAlphaInput := *api.NewUpdateDestinationsForJourneyAlphaInput([]api.UpdateDestinationsInput{*api.NewUpdateDestinationsInput("StateId_example", []api.Destination{*api.NewDestination("DestinationId_example", []api.Activation{*api.NewActivation("EventName_example", *api.NewActionDefinition(map[string]interface{}{"key": interface{}(123)}, "ActionId_example"))})})}) // UpdateDestinationsForJourneyAlphaInput | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.JourneysAPI.UpdateDestinationsForJourney(ctx, spaceId, containerId).UpdateDestinationsForJourneyAlphaInput(updateDestinationsForJourneyAlphaInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JourneysAPI.UpdateDestinationsForJourney``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `UpdateDestinationsForJourney`: UpdateDestinationsForJourney200Response
    fmt.Fprintf(os.Stdout, "Response from `JourneysAPI.UpdateDestinationsForJourney`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** |  | 
**containerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDestinationsForJourneyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateDestinationsForJourneyAlphaInput** | [**UpdateDestinationsForJourneyAlphaInput**](UpdateDestinationsForJourneyAlphaInput.md) |  | 

### Return type

[**UpdateDestinationsForJourney200Response**](UpdateDestinationsForJourney200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/vnd.segment.v1alpha+json
- **Accept**: application/vnd.segment.v1alpha+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: UpdateStatusForJourney

> UpdateStatusForJourney200Response UpdateStatusForJourney(ctx, spaceId, containerId).UpdateStatusForJourneyAlphaInput(updateStatusForJourneyAlphaInput).Execute()

Update Status for Journey



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
    containerId := "jcon_2tG95HZh4oPsgHfcOlznyfcDDAg" // string | 
    updateStatusForJourneyAlphaInput := *api.NewUpdateStatusForJourneyAlphaInput(float32(123), "Action_example") // UpdateStatusForJourneyAlphaInput | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.JourneysAPI.UpdateStatusForJourney(ctx, spaceId, containerId).UpdateStatusForJourneyAlphaInput(updateStatusForJourneyAlphaInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JourneysAPI.UpdateStatusForJourney``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `UpdateStatusForJourney`: UpdateStatusForJourney200Response
    fmt.Fprintf(os.Stdout, "Response from `JourneysAPI.UpdateStatusForJourney`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** |  | 
**containerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStatusForJourneyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateStatusForJourneyAlphaInput** | [**UpdateStatusForJourneyAlphaInput**](UpdateStatusForJourneyAlphaInput.md) |  | 

### Return type

[**UpdateStatusForJourney200Response**](UpdateStatusForJourney200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/vnd.segment.v1alpha+json
- **Accept**: application/vnd.segment.v1alpha+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

