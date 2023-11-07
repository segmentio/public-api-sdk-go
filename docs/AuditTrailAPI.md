# \AuditTrailAPI

All URIs are relative to *https://api.segmentapis.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListAuditEvents**](AuditTrailAPI.md#ListAuditEvents) | **Get** /audit-events | List Audit Events



## Operation: ListAuditEvents

> ListAuditEvents200Response ListAuditEvents(ctx).Pagination(pagination).StartTime(startTime).EndTime(endTime).ResourceId(resourceId).ResourceType(resourceType).Execute()

List Audit Events



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
    pagination := *openapiclient.NewPaginationInput(float32(123)) // PaginationInput | Defines the pagination parameters.  This parameter exists in v1.
    startTime := "startTime_example" // string | Filter response to events that happened after this time.  This parameter exists in v1. (optional)
    endTime := "endTime_example" // string | Filter response to events that happened before this time. Defaults to the current time, or the end time from the pagination cursor.  This parameter exists in v1. (optional)
    resourceId := "9aQ1Lj62S4bomZKLF4DPqW" // string | Filter response to events that affect a specific resource, for example, a single Source.  This parameter exists in v1. (optional)
    resourceType := "resourceType_example" // string | Filter response to events that affect a specific type, for example, Sources, Warehouses, and Tracking Plans.  This parameter exists in v1. (optional)

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    resp, r, err := apiClient.AuditTrailAPI.ListAuditEvents(context.Background()).Pagination(pagination).StartTime(startTime).EndTime(endTime).ResourceId(resourceId).ResourceType(resourceType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuditTrailAPI.ListAuditEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAuditEvents`: ListAuditEvents200Response
    fmt.Fprintf(os.Stdout, "Response from `AuditTrailAPI.ListAuditEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAuditEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pagination** | [**PaginationInput**](PaginationInput.md) | Defines the pagination parameters.  This parameter exists in v1. | 
 **startTime** | **string** | Filter response to events that happened after this time.  This parameter exists in v1. | 
 **endTime** | **string** | Filter response to events that happened before this time. Defaults to the current time, or the end time from the pagination cursor.  This parameter exists in v1. | 
 **resourceId** | **string** | Filter response to events that affect a specific resource, for example, a single Source.  This parameter exists in v1. | 
 **resourceType** | **string** | Filter response to events that affect a specific type, for example, Sources, Warehouses, and Tracking Plans.  This parameter exists in v1. | 

### Return type

[**ListAuditEvents200Response**](ListAuditEvents200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1+json, application/json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

