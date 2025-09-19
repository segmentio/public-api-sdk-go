# \ActivationsAPI

All URIs are relative to *https://api.segmentapis.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddActivationToAudience**](ActivationsAPI.md#AddActivationToAudience) | **Post** /spaces/{spaceId}/audiences/{audienceId}/destination-connections/{connectionId}/activations | Add Activation to Audience
[**AddDestinationToAudience**](ActivationsAPI.md#AddDestinationToAudience) | **Post** /spaces/{spaceId}/audiences/{audienceId}/destination-connections | Add Destination to Audience
[**GetActivationFromAudience**](ActivationsAPI.md#GetActivationFromAudience) | **Get** /spaces/{spaceId}/audiences/{audienceId}/activations/{id} | Get Activation from Audience
[**ListActivationsFromAudience**](ActivationsAPI.md#ListActivationsFromAudience) | **Get** /spaces/{spaceId}/audiences/{audienceId}/activations | List Activations from Audience
[**ListDestinationsFromAudience**](ActivationsAPI.md#ListDestinationsFromAudience) | **Get** /spaces/{spaceId}/audiences/{audienceId}/destination-connections | List Destinations from Audience
[**RemoveActivationFromAudience**](ActivationsAPI.md#RemoveActivationFromAudience) | **Delete** /spaces/{spaceId}/audiences/{audienceId}/activations/{id} | Remove Activation from Audience
[**UpdateActivationForAudience**](ActivationsAPI.md#UpdateActivationForAudience) | **Patch** /spaces/{spaceId}/audiences/{audienceId}/activations/{id} | Update Activation for Audience



## Operation: AddActivationToAudience

> AddActivationToAudience200Response AddActivationToAudience(ctx, spaceId, audienceId, connectionId).AddActivationToAudienceAlphaInput(addActivationToAudienceAlphaInput).Execute()

Add Activation to Audience



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
    spaceId := "spa_9aQ1Lj62S4bomZKLF4DPqW" // string | 
    audienceId := "aud_0ujsszwN8NRY24YaXiTIE2VWDTS" // string | 
    connectionId := "ii_123456789" // string | 
    addActivationToAudienceAlphaInput := *api.NewAddActivationToAudienceAlphaInput(false, "ActivationType_example", "ActivationName_example", *api.NewPersonalizationInput(*api.NewProfile([]string{"Properties_example"})), *api.NewDestinationSubscriptionConfiguration("ActionId_example")) // AddActivationToAudienceAlphaInput | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.ActivationsAPI.AddActivationToAudience(ctx, spaceId, audienceId, connectionId).AddActivationToAudienceAlphaInput(addActivationToAudienceAlphaInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActivationsAPI.AddActivationToAudience``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `AddActivationToAudience`: AddActivationToAudience200Response
    fmt.Fprintf(os.Stdout, "Response from `ActivationsAPI.AddActivationToAudience`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** |  | 
**audienceId** | **string** |  | 
**connectionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddActivationToAudienceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **addActivationToAudienceAlphaInput** | [**AddActivationToAudienceAlphaInput**](AddActivationToAudienceAlphaInput.md) |  | 

### Return type

[**AddActivationToAudience200Response**](AddActivationToAudience200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/vnd.segment.v1alpha+json
- **Accept**: application/vnd.segment.v1alpha+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: AddDestinationToAudience

> AddDestinationToAudience200Response AddDestinationToAudience(ctx, spaceId, audienceId).AddDestinationToAudienceAlphaInput(addDestinationToAudienceAlphaInput).Execute()

Add Destination to Audience



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
    spaceId := "spa_9aQ1Lj62S4bomZKLF4DPqW" // string | 
    audienceId := "aud_0ujsszwN8NRY24YaXiTIE2VWDTS" // string | 
    addDestinationToAudienceAlphaInput := *api.NewAddDestinationToAudienceAlphaInput(*api.NewDestinationInput("Id_example", "Type_example")) // AddDestinationToAudienceAlphaInput | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.ActivationsAPI.AddDestinationToAudience(ctx, spaceId, audienceId).AddDestinationToAudienceAlphaInput(addDestinationToAudienceAlphaInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActivationsAPI.AddDestinationToAudience``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `AddDestinationToAudience`: AddDestinationToAudience200Response
    fmt.Fprintf(os.Stdout, "Response from `ActivationsAPI.AddDestinationToAudience`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** |  | 
**audienceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddDestinationToAudienceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **addDestinationToAudienceAlphaInput** | [**AddDestinationToAudienceAlphaInput**](AddDestinationToAudienceAlphaInput.md) |  | 

### Return type

[**AddDestinationToAudience200Response**](AddDestinationToAudience200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/vnd.segment.v1alpha+json
- **Accept**: application/vnd.segment.v1alpha+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: GetActivationFromAudience

> GetActivationFromAudience200Response GetActivationFromAudience(ctx, spaceId, audienceId, id).Execute()

Get Activation from Audience



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
    spaceId := "spa_9aQ1Lj62S4bomZKLF4DPqW" // string | 
    audienceId := "aud_0ujsszwN8NRY24YaXiTIE2VWDTS" // string | 
    id := "act_987654321" // string | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.ActivationsAPI.GetActivationFromAudience(ctx, spaceId, audienceId, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActivationsAPI.GetActivationFromAudience``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `GetActivationFromAudience`: GetActivationFromAudience200Response
    fmt.Fprintf(os.Stdout, "Response from `ActivationsAPI.GetActivationFromAudience`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** |  | 
**audienceId** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetActivationFromAudienceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**GetActivationFromAudience200Response**](GetActivationFromAudience200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1alpha+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: ListActivationsFromAudience

> ListActivationsFromAudience200Response ListActivationsFromAudience(ctx, spaceId, audienceId).Pagination(pagination).Execute()

List Activations from Audience



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
    spaceId := "spa_9aQ1Lj62S4bomZKLF4DPqW" // string | 
    audienceId := "aud_0ujsszwN8NRY24YaXiTIE2VWDTS" // string | 
    pagination := *api.NewPaginationInput(10) // PaginationInput | Optional pagination.  This parameter exists in alpha. (optional)

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.ActivationsAPI.ListActivationsFromAudience(ctx, spaceId, audienceId).Pagination(pagination).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActivationsAPI.ListActivationsFromAudience``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `ListActivationsFromAudience`: ListActivationsFromAudience200Response
    fmt.Fprintf(os.Stdout, "Response from `ActivationsAPI.ListActivationsFromAudience`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** |  | 
**audienceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListActivationsFromAudienceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pagination** | [**PaginationInput**](PaginationInput.md) | Optional pagination.  This parameter exists in alpha. | 

### Return type

[**ListActivationsFromAudience200Response**](ListActivationsFromAudience200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1alpha+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: ListDestinationsFromAudience

> ListDestinationsFromAudience200Response ListDestinationsFromAudience(ctx, spaceId, audienceId).Pagination(pagination).Execute()

List Destinations from Audience



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
    spaceId := "spa_9aQ1Lj62S4bomZKLF4DPqW" // string | 
    audienceId := "aud_0ujsszwN8NRY24YaXiTIE2VWDTS" // string | 
    pagination := *api.NewPaginationInput(10) // PaginationInput | Optional pagination.  This parameter exists in alpha. (optional)

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.ActivationsAPI.ListDestinationsFromAudience(ctx, spaceId, audienceId).Pagination(pagination).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActivationsAPI.ListDestinationsFromAudience``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `ListDestinationsFromAudience`: ListDestinationsFromAudience200Response
    fmt.Fprintf(os.Stdout, "Response from `ActivationsAPI.ListDestinationsFromAudience`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** |  | 
**audienceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListDestinationsFromAudienceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pagination** | [**PaginationInput**](PaginationInput.md) | Optional pagination.  This parameter exists in alpha. | 

### Return type

[**ListDestinationsFromAudience200Response**](ListDestinationsFromAudience200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1alpha+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: RemoveActivationFromAudience

> RemoveActivationFromAudience200Response RemoveActivationFromAudience(ctx, spaceId, audienceId, id).Execute()

Remove Activation from Audience



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
    spaceId := "spa_9aQ1Lj62S4bomZKLF4DPqW" // string | 
    audienceId := "aud_0ujsszwN8NRY24YaXiTIE2VWDTS" // string | 
    id := "act_987654321" // string | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.ActivationsAPI.RemoveActivationFromAudience(ctx, spaceId, audienceId, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActivationsAPI.RemoveActivationFromAudience``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `RemoveActivationFromAudience`: RemoveActivationFromAudience200Response
    fmt.Fprintf(os.Stdout, "Response from `ActivationsAPI.RemoveActivationFromAudience`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** |  | 
**audienceId** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveActivationFromAudienceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**RemoveActivationFromAudience200Response**](RemoveActivationFromAudience200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1alpha+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: UpdateActivationForAudience

> UpdateActivationForAudience200Response UpdateActivationForAudience(ctx, spaceId, audienceId, id).UpdateActivationForAudienceAlphaInput(updateActivationForAudienceAlphaInput).Execute()

Update Activation for Audience



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
    spaceId := "spa_9aQ1Lj62S4bomZKLF4DPqW" // string | 
    audienceId := "aud_0ujsszwN8NRY24YaXiTIE2VWDTS" // string | 
    id := "act_987654321" // string | 
    updateActivationForAudienceAlphaInput := *api.NewUpdateActivationForAudienceAlphaInput() // UpdateActivationForAudienceAlphaInput | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.ActivationsAPI.UpdateActivationForAudience(ctx, spaceId, audienceId, id).UpdateActivationForAudienceAlphaInput(updateActivationForAudienceAlphaInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActivationsAPI.UpdateActivationForAudience``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `UpdateActivationForAudience`: UpdateActivationForAudience200Response
    fmt.Fprintf(os.Stdout, "Response from `ActivationsAPI.UpdateActivationForAudience`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** |  | 
**audienceId** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateActivationForAudienceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **updateActivationForAudienceAlphaInput** | [**UpdateActivationForAudienceAlphaInput**](UpdateActivationForAudienceAlphaInput.md) |  | 

### Return type

[**UpdateActivationForAudience200Response**](UpdateActivationForAudience200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/vnd.segment.v1alpha+json
- **Accept**: application/vnd.segment.v1alpha+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

