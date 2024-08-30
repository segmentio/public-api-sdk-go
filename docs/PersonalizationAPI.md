# \PersonalizationAPI

All URIs are relative to *https://api.segmentapis.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPersonalizationData**](PersonalizationAPI.md#GetPersonalizationData) | **Get** /personalization/v1/spaces/{spaceId}/{entityType}/{entityIdentifier} | Get Personalization Data



## Operation: GetPersonalizationData

> GetPersonalizationData200Response GetPersonalizationData(ctx, spaceId, entityType, entityIdentifier).ChildrenEntityType(childrenEntityType).Execute()

Get Personalization Data



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
    entityType := "profile" // string | 
    entityIdentifier := "identifier" // string | 
    childrenEntityType := "childrenEntityType_example" // string | Child entity type.  This parameter exists in alpha. (optional)

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.PersonalizationAPI.GetPersonalizationData(ctx, spaceId, entityType, entityIdentifier).ChildrenEntityType(childrenEntityType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersonalizationAPI.GetPersonalizationData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `GetPersonalizationData`: GetPersonalizationData200Response
    fmt.Fprintf(os.Stdout, "Response from `PersonalizationAPI.GetPersonalizationData`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** |  | 
**entityType** | **string** |  | 
**entityIdentifier** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPersonalizationDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **childrenEntityType** | **string** | Child entity type.  This parameter exists in alpha. | 

### Return type

[**GetPersonalizationData200Response**](GetPersonalizationData200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1alpha+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

