# \DBTAPI

All URIs are relative to *https://api.segmentapis.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDbtModelSyncTrigger**](DBTAPI.md#CreateDbtModelSyncTrigger) | **Post** /dbt-model-syncs/trigger | Create Dbt Model Sync Trigger



## Operation: CreateDbtModelSyncTrigger

> CreateDbtModelSyncTrigger200Response CreateDbtModelSyncTrigger(ctx).CreateDbtModelSyncTriggerInput(createDbtModelSyncTriggerInput).Execute()

Create Dbt Model Sync Trigger



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
    createDbtModelSyncTriggerInput := *api.NewCreateDbtModelSyncTriggerInput("SourceId_example") // CreateDbtModelSyncTriggerInput | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.DBTAPI.CreateDbtModelSyncTrigger(ctx).CreateDbtModelSyncTriggerInput(createDbtModelSyncTriggerInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DBTAPI.CreateDbtModelSyncTrigger``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `CreateDbtModelSyncTrigger`: CreateDbtModelSyncTrigger200Response
    fmt.Fprintf(os.Stdout, "Response from `DBTAPI.CreateDbtModelSyncTrigger`: %v\n", resp.GetData())
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDbtModelSyncTriggerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createDbtModelSyncTriggerInput** | [**CreateDbtModelSyncTriggerInput**](CreateDbtModelSyncTriggerInput.md) |  | 

### Return type

[**CreateDbtModelSyncTrigger200Response**](CreateDbtModelSyncTrigger200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/vnd.segment.v1beta+json
- **Accept**: application/vnd.segment.v1beta+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

