# \EventsAPI

All URIs are relative to *https://api.segmentapis.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEventsVolumeFromWorkspace**](EventsAPI.md#GetEventsVolumeFromWorkspace) | **Get** /events/volume | Get Events Volume from Workspace



## Operation: GetEventsVolumeFromWorkspace

> GetEventsVolumeFromWorkspace200Response GetEventsVolumeFromWorkspace(ctx).Granularity(granularity).StartTime(startTime).EndTime(endTime).GroupBy(groupBy).SourceId(sourceId).EventName(eventName).EventType(eventType).AppVersion(appVersion).Pagination(pagination).Execute()

Get Events Volume from Workspace



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
    granularity := "DAY" // string | The size of each bucket in the requested window.  This parameter exists in v1.
    startTime := "2021-10-28T00:00:00Z" // string | The ISO8601 formatted timestamp that corresponds to the beginning of the requested time frame, inclusive.  This parameter exists in v1.
    endTime := "2021-10-29T16:40:00Z" // string | The ISO8601 formatted timestamp that corresponds to the end of the requested time frame, noninclusive. Segment recommends that you lag queries 1 minute behind clock time to reduce the risk for latency to impact the counts.  This parameter exists in v1.
    groupBy := []string{[]string{"GroupBy_example"}} // []string | A comma-delimited list of strings that represents the dimensions to group the result by. The options are: `eventName`, `eventType` and `source`.  This parameter exists in v1. (optional)
    sourceId := []string{[]string{"SourceId_example"}} // []string | A list of strings which filters the results to the given SourceIds.  This parameter exists in v1. (optional)
    eventName := []string{[]string{"EventName_example"}} // []string | A list of strings which filters the results to the given EventNames.  This parameter exists in v1. (optional)
    eventType := []string{[]string{"EventType_example"}} // []string | A list of strings which filters the results to the given EventTypes.  This parameter exists in v1. (optional)
    appVersion := []string{[]string{"AppVersion_example"}} // []string | A list of strings which filters the results to the given AppVersions.  This parameter exists in v1. (optional)
    pagination := *api.NewPaginationInput(10) // PaginationInput | Pagination input for event volume by Workspace.  This parameter exists in v1. (optional)

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.EventsAPI.GetEventsVolumeFromWorkspace(ctx).Granularity(granularity).StartTime(startTime).EndTime(endTime).GroupBy(groupBy).SourceId(sourceId).EventName(eventName).EventType(eventType).AppVersion(appVersion).Pagination(pagination).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.GetEventsVolumeFromWorkspace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `GetEventsVolumeFromWorkspace`: GetEventsVolumeFromWorkspace200Response
    fmt.Fprintf(os.Stdout, "Response from `EventsAPI.GetEventsVolumeFromWorkspace`: %v\n", resp.GetData())
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEventsVolumeFromWorkspaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **granularity** | **string** | The size of each bucket in the requested window.  This parameter exists in v1. | 
 **startTime** | **string** | The ISO8601 formatted timestamp that corresponds to the beginning of the requested time frame, inclusive.  This parameter exists in v1. | 
 **endTime** | **string** | The ISO8601 formatted timestamp that corresponds to the end of the requested time frame, noninclusive. Segment recommends that you lag queries 1 minute behind clock time to reduce the risk for latency to impact the counts.  This parameter exists in v1. | 
 **groupBy** | **[][]string** | A comma-delimited list of strings that represents the dimensions to group the result by. The options are: &#x60;eventName&#x60;, &#x60;eventType&#x60; and &#x60;source&#x60;.  This parameter exists in v1. | 
 **sourceId** | **[][]string** | A list of strings which filters the results to the given SourceIds.  This parameter exists in v1. | 
 **eventName** | **[][]string** | A list of strings which filters the results to the given EventNames.  This parameter exists in v1. | 
 **eventType** | **[][]string** | A list of strings which filters the results to the given EventTypes.  This parameter exists in v1. | 
 **appVersion** | **[][]string** | A list of strings which filters the results to the given AppVersions.  This parameter exists in v1. | 
 **pagination** | [**PaginationInput**](PaginationInput.md) | Pagination input for event volume by Workspace.  This parameter exists in v1. | 

### Return type

[**GetEventsVolumeFromWorkspace200Response**](GetEventsVolumeFromWorkspace200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1+json, application/json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

