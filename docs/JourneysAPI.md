# \JourneysAPI

All URIs are relative to *https://api.segmentapis.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateJourney**](JourneysAPI.md#CreateJourney) | **Post** /spaces/{spaceId}/journeys | Create Journey



## Operation: CreateJourney

> CreateJourney200Response CreateJourney(ctx, spaceId).CreateJourneyAlphaInput(createJourneyAlphaInput).Execute()

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
    spaceId := "9aQ1Lj62S4bomZKLF4DPqW" // string | 
    createJourneyAlphaInput := *api.NewCreateJourneyAlphaInput("Name_example", *api.NewJourneyDefinition("Initial_example", map[string]interface{}{"key": interface{}(123)})) // CreateJourneyAlphaInput | 

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
    // response from `CreateJourney`: CreateJourney200Response
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

[**CreateJourney200Response**](CreateJourney200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/vnd.segment.v1alpha+json
- **Accept**: application/vnd.segment.v1alpha+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

