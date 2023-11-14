# \TestingAPI

All URIs are relative to *https://api.segmentapis.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Echo**](TestingAPI.md#Echo) | **Get** /echo | Echo



## Operation: Echo

> Echo200Response Echo(ctx).Message(message).Delay(delay).TriggerError(triggerError).TriggerMultipleErrors(triggerMultipleErrors).TriggerUnexpectedError(triggerUnexpectedError).StatusCode(statusCode).Execute()

Echo



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
    message := "Hello, World!" // string | Sets the response `message` field. The response contains this field's entry.  This parameter exists in alpha.
    delay := float32(8.14) // float32 | The desired response delay, in milliseconds.  This parameter exists in alpha. (optional)
    triggerError := true // bool | If `true`, returns an HTTP `4xx` error that contains the string in `message`.  This parameter exists in alpha. (optional)
    triggerMultipleErrors := true // bool | If `true`, returns an HTTP `4xx` error that contains the value of the `message` field in the error message array.  This has no effect if the request sets `triggerError`.  This parameter exists in alpha. (optional)
    triggerUnexpectedError := true // bool | If `true`, triggers a `500` error.  This has no effect if the request sets either `triggerError` or `triggerMultipleErrors`.  This parameter exists in alpha. (optional)
    statusCode := float32(8.14) // float32 | Sets the HTTP status code to return.  This parameter exists in alpha. (optional)

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.TestingAPI.Echo(ctx).Message(message).Delay(delay).TriggerError(triggerError).TriggerMultipleErrors(triggerMultipleErrors).TriggerUnexpectedError(triggerUnexpectedError).StatusCode(statusCode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TestingAPI.Echo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `Echo`: Echo200Response
    fmt.Fprintf(os.Stdout, "Response from `TestingAPI.Echo`: %v\n", resp.GetData())
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEchoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **message** | **string** | Sets the response &#x60;message&#x60; field. The response contains this field&#39;s entry.  This parameter exists in alpha. | 
 **delay** | **float32** | The desired response delay, in milliseconds.  This parameter exists in alpha. | 
 **triggerError** | **bool** | If &#x60;true&#x60;, returns an HTTP &#x60;4xx&#x60; error that contains the string in &#x60;message&#x60;.  This parameter exists in alpha. | 
 **triggerMultipleErrors** | **bool** | If &#x60;true&#x60;, returns an HTTP &#x60;4xx&#x60; error that contains the value of the &#x60;message&#x60; field in the error message array.  This has no effect if the request sets &#x60;triggerError&#x60;.  This parameter exists in alpha. | 
 **triggerUnexpectedError** | **bool** | If &#x60;true&#x60;, triggers a &#x60;500&#x60; error.  This has no effect if the request sets either &#x60;triggerError&#x60; or &#x60;triggerMultipleErrors&#x60;.  This parameter exists in alpha. | 
 **statusCode** | **float32** | Sets the HTTP status code to return.  This parameter exists in alpha. | 

### Return type

[**Echo200Response**](Echo200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1+json, application/json, application/vnd.segment.v1alpha+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

