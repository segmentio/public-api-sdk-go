# \LabelsAPI

All URIs are relative to *https://api.segmentapis.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLabel**](LabelsAPI.md#CreateLabel) | **Post** /labels | Create Label
[**DeleteLabel**](LabelsAPI.md#DeleteLabel) | **Delete** /labels/{key}/{value} | Delete Label
[**ListLabels**](LabelsAPI.md#ListLabels) | **Get** /labels | List Labels



## Operation: CreateLabel

> CreateLabel200Response CreateLabel(ctx).CreateLabelV1Input(createLabelV1Input).Execute()

Create Label



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
    createLabelV1Input := *openapiclient.NewCreateLabelV1Input(*openapiclient.NewLabelV1("Key_example", "Value_example")) // CreateLabelV1Input | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    resp, r, err := apiClient.LabelsAPI.CreateLabel(context.Background()).CreateLabelV1Input(createLabelV1Input).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LabelsAPI.CreateLabel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLabel`: CreateLabel200Response
    fmt.Fprintf(os.Stdout, "Response from `LabelsAPI.CreateLabel`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateLabelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createLabelV1Input** | [**CreateLabelV1Input**](CreateLabelV1Input.md) |  | 

### Return type

[**CreateLabel200Response**](CreateLabel200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/vnd.segment.v1+json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json
- **Accept**: application/vnd.segment.v1+json, application/json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: DeleteLabel

> DeleteLabel200Response DeleteLabel(ctx, key, value).Execute()

Delete Label



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
    key := "environment" // string | 
    value := "yolo" // string | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    resp, r, err := apiClient.LabelsAPI.DeleteLabel(context.Background(), key, value).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LabelsAPI.DeleteLabel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteLabel`: DeleteLabel200Response
    fmt.Fprintf(os.Stdout, "Response from `LabelsAPI.DeleteLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** |  | 
**value** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLabelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DeleteLabel200Response**](DeleteLabel200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1+json, application/json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: ListLabels

> ListLabels200Response ListLabels(ctx).Execute()

List Labels



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
    resp, r, err := apiClient.LabelsAPI.ListLabels(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LabelsAPI.ListLabels``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLabels`: ListLabels200Response
    fmt.Fprintf(os.Stdout, "Response from `LabelsAPI.ListLabels`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListLabelsRequest struct via the builder pattern


### Return type

[**ListLabels200Response**](ListLabels200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1+json, application/json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

