# \FunctionsAPI

All URIs are relative to *https://api.segmentapis.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFunction**](FunctionsAPI.md#CreateFunction) | **Post** /functions | Create Function
[**CreateFunctionDeployment**](FunctionsAPI.md#CreateFunctionDeployment) | **Post** /functions/{functionId}/deploy | Create Function Deployment
[**CreateInsertFunctionInstance**](FunctionsAPI.md#CreateInsertFunctionInstance) | **Post** /insert-function-instances | Create Insert Function Instance
[**DeleteFunction**](FunctionsAPI.md#DeleteFunction) | **Delete** /functions/{functionId} | Delete Function
[**DeleteInsertFunctionInstance**](FunctionsAPI.md#DeleteInsertFunctionInstance) | **Delete** /insert-function-instances/{instanceId} | Delete Insert Function Instance
[**GetFunction**](FunctionsAPI.md#GetFunction) | **Get** /functions/{functionId} | Get Function
[**GetFunctionVersion**](FunctionsAPI.md#GetFunctionVersion) | **Get** /functions/{functionId}/versions/{versionId} | Get Function Version
[**GetInsertFunctionInstance**](FunctionsAPI.md#GetInsertFunctionInstance) | **Get** /insert-function-instances/{instanceId} | Get Insert Function Instance
[**ListFunctionVersions**](FunctionsAPI.md#ListFunctionVersions) | **Get** /functions/{functionId}/versions | List Function Versions
[**ListFunctions**](FunctionsAPI.md#ListFunctions) | **Get** /functions | List Functions
[**ListInsertFunctionInstances**](FunctionsAPI.md#ListInsertFunctionInstances) | **Get** /insert-function-instances | List Insert Function Instances
[**RestoreFunctionVersion**](FunctionsAPI.md#RestoreFunctionVersion) | **Post** /functions/{functionId}/versions | Restore Function Version
[**UpdateFunction**](FunctionsAPI.md#UpdateFunction) | **Patch** /functions/{functionId} | Update Function
[**UpdateInsertFunctionInstance**](FunctionsAPI.md#UpdateInsertFunctionInstance) | **Patch** /insert-function-instances/{instanceId} | Update Insert Function Instance



## Operation: CreateFunction

> CreateFunction200Response CreateFunction(ctx).CreateFunctionV1Input(createFunctionV1Input).Execute()

Create Function



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
    createFunctionV1Input := *openapiclient.NewCreateFunctionV1Input("Code_example", "DisplayName_example", "ResourceType_example") // CreateFunctionV1Input | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    resp, r, err := apiClient.FunctionsAPI.CreateFunction(context.Background()).CreateFunctionV1Input(createFunctionV1Input).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAPI.CreateFunction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateFunction`: CreateFunction200Response
    fmt.Fprintf(os.Stdout, "Response from `FunctionsAPI.CreateFunction`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFunctionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createFunctionV1Input** | [**CreateFunctionV1Input**](CreateFunctionV1Input.md) |  | 

### Return type

[**CreateFunction200Response**](CreateFunction200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/vnd.segment.v1+json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json
- **Accept**: application/vnd.segment.v1+json, application/json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: CreateFunctionDeployment

> CreateFunctionDeployment200Response CreateFunctionDeployment(ctx, functionId).Execute()

Create Function Deployment



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
    functionId := "sfn_rh5BDZp6QDHvXFCkibm1pR" // string | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    resp, r, err := apiClient.FunctionsAPI.CreateFunctionDeployment(context.Background(), functionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAPI.CreateFunctionDeployment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateFunctionDeployment`: CreateFunctionDeployment200Response
    fmt.Fprintf(os.Stdout, "Response from `FunctionsAPI.CreateFunctionDeployment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**functionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateFunctionDeploymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateFunctionDeployment200Response**](CreateFunctionDeployment200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1+json, application/json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: CreateInsertFunctionInstance

> CreateInsertFunctionInstance200Response CreateInsertFunctionInstance(ctx).CreateInsertFunctionInstanceAlphaInput(createInsertFunctionInstanceAlphaInput).Execute()

Create Insert Function Instance



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
    createInsertFunctionInstanceAlphaInput := *openapiclient.NewCreateInsertFunctionInstanceAlphaInput("FunctionId_example", "IntegrationId_example", "Name_example", map[string]interface{}{"key": interface{}(123)}) // CreateInsertFunctionInstanceAlphaInput | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    resp, r, err := apiClient.FunctionsAPI.CreateInsertFunctionInstance(context.Background()).CreateInsertFunctionInstanceAlphaInput(createInsertFunctionInstanceAlphaInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAPI.CreateInsertFunctionInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateInsertFunctionInstance`: CreateInsertFunctionInstance200Response
    fmt.Fprintf(os.Stdout, "Response from `FunctionsAPI.CreateInsertFunctionInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateInsertFunctionInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createInsertFunctionInstanceAlphaInput** | [**CreateInsertFunctionInstanceAlphaInput**](CreateInsertFunctionInstanceAlphaInput.md) |  | 

### Return type

[**CreateInsertFunctionInstance200Response**](CreateInsertFunctionInstance200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/vnd.segment.v1alpha+json
- **Accept**: application/vnd.segment.v1alpha+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: DeleteFunction

> DeleteFunction200Response DeleteFunction(ctx, functionId).Execute()

Delete Function



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
    functionId := "sfnc_wXzcDGFR3KmjLDrtSawNHf" // string | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    resp, r, err := apiClient.FunctionsAPI.DeleteFunction(context.Background(), functionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAPI.DeleteFunction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteFunction`: DeleteFunction200Response
    fmt.Fprintf(os.Stdout, "Response from `FunctionsAPI.DeleteFunction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**functionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFunctionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteFunction200Response**](DeleteFunction200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1+json, application/json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: DeleteInsertFunctionInstance

> DeleteInsertFunctionInstance200Response DeleteInsertFunctionInstance(ctx, instanceId).Execute()

Delete Insert Function Instance



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
    instanceId := "64edfa5ac43ed978ee4f2b7f" // string | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    resp, r, err := apiClient.FunctionsAPI.DeleteInsertFunctionInstance(context.Background(), instanceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAPI.DeleteInsertFunctionInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteInsertFunctionInstance`: DeleteInsertFunctionInstance200Response
    fmt.Fprintf(os.Stdout, "Response from `FunctionsAPI.DeleteInsertFunctionInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteInsertFunctionInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteInsertFunctionInstance200Response**](DeleteInsertFunctionInstance200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1alpha+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: GetFunction

> GetFunction200Response GetFunction(ctx, functionId).Execute()

Get Function



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
    functionId := "sfnc_wXzcDGFR3KmjLDrtSawNHf" // string | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    resp, r, err := apiClient.FunctionsAPI.GetFunction(context.Background(), functionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAPI.GetFunction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFunction`: GetFunction200Response
    fmt.Fprintf(os.Stdout, "Response from `FunctionsAPI.GetFunction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**functionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFunctionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetFunction200Response**](GetFunction200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1+json, application/json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: GetFunctionVersion

> GetFunctionVersion200Response GetFunctionVersion(ctx, functionId, versionId).Execute()

Get Function Version



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
    functionId := "sfnc_wXzcDGFR3KmjLDrtSawNHf" // string | 
    versionId := "2Ma03fahSqLoEzQfV5o2aSfVEHs" // string | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    resp, r, err := apiClient.FunctionsAPI.GetFunctionVersion(context.Background(), functionId, versionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAPI.GetFunctionVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFunctionVersion`: GetFunctionVersion200Response
    fmt.Fprintf(os.Stdout, "Response from `FunctionsAPI.GetFunctionVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**functionId** | **string** |  | 
**versionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFunctionVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetFunctionVersion200Response**](GetFunctionVersion200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1alpha+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: GetInsertFunctionInstance

> GetInsertFunctionInstance200Response GetInsertFunctionInstance(ctx, instanceId).Execute()

Get Insert Function Instance



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
    instanceId := "65414b1bf05ef7ecfd8b6736" // string | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    resp, r, err := apiClient.FunctionsAPI.GetInsertFunctionInstance(context.Background(), instanceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAPI.GetInsertFunctionInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInsertFunctionInstance`: GetInsertFunctionInstance200Response
    fmt.Fprintf(os.Stdout, "Response from `FunctionsAPI.GetInsertFunctionInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInsertFunctionInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetInsertFunctionInstance200Response**](GetInsertFunctionInstance200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1alpha+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: ListFunctionVersions

> ListFunctionVersions200Response ListFunctionVersions(ctx, functionId).Pagination(pagination).Execute()

List Function Versions



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
    functionId := "sfnc_wXzcDGFR3KmjLDrtSawNHf" // string | 
    pagination := *openapiclient.NewPaginationInput(float32(123)) // PaginationInput | Pagination parameters.  This parameter exists in alpha.

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    resp, r, err := apiClient.FunctionsAPI.ListFunctionVersions(context.Background(), functionId).Pagination(pagination).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAPI.ListFunctionVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListFunctionVersions`: ListFunctionVersions200Response
    fmt.Fprintf(os.Stdout, "Response from `FunctionsAPI.ListFunctionVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**functionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListFunctionVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pagination** | [**PaginationInput**](PaginationInput.md) | Pagination parameters.  This parameter exists in alpha. | 

### Return type

[**ListFunctionVersions200Response**](ListFunctionVersions200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1alpha+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: ListFunctions

> ListFunctions200Response ListFunctions(ctx).Pagination(pagination).ResourceType(resourceType).Execute()

List Functions



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
    pagination := *openapiclient.NewPaginationInput(float32(123)) // PaginationInput | Pagination parameters.  This parameter exists in v1.
    resourceType := "SOURCE" // string | The Function type.  Config API note: equal to `type`.  This parameter exists in v1.

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    resp, r, err := apiClient.FunctionsAPI.ListFunctions(context.Background()).Pagination(pagination).ResourceType(resourceType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAPI.ListFunctions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListFunctions`: ListFunctions200Response
    fmt.Fprintf(os.Stdout, "Response from `FunctionsAPI.ListFunctions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListFunctionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pagination** | [**PaginationInput**](PaginationInput.md) | Pagination parameters.  This parameter exists in v1. | 
 **resourceType** | **string** | The Function type.  Config API note: equal to &#x60;type&#x60;.  This parameter exists in v1. | 

### Return type

[**ListFunctions200Response**](ListFunctions200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1+json, application/json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: ListInsertFunctionInstances

> ListInsertFunctionInstances200Response ListInsertFunctionInstances(ctx).Pagination(pagination).FunctionId(functionId).Execute()

List Insert Function Instances



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
    pagination := *openapiclient.NewPaginationInput(float32(123)) // PaginationInput | Pagination parameters.  This parameter exists in alpha.
    functionId := "76365637324e715a67535831" // string | The insert Function class id to lookup.  This parameter exists in alpha.

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    resp, r, err := apiClient.FunctionsAPI.ListInsertFunctionInstances(context.Background()).Pagination(pagination).FunctionId(functionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAPI.ListInsertFunctionInstances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListInsertFunctionInstances`: ListInsertFunctionInstances200Response
    fmt.Fprintf(os.Stdout, "Response from `FunctionsAPI.ListInsertFunctionInstances`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListInsertFunctionInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pagination** | [**PaginationInput**](PaginationInput.md) | Pagination parameters.  This parameter exists in alpha. | 
 **functionId** | **string** | The insert Function class id to lookup.  This parameter exists in alpha. | 

### Return type

[**ListInsertFunctionInstances200Response**](ListInsertFunctionInstances200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1alpha+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: RestoreFunctionVersion

> RestoreFunctionVersion200Response RestoreFunctionVersion(ctx, functionId).RestoreFunctionVersionAlphaInput(restoreFunctionVersionAlphaInput).Execute()

Restore Function Version



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
    functionId := "sfnc_wXzcDGFR3KmjLDrtSawNHf" // string | 
    restoreFunctionVersionAlphaInput := *openapiclient.NewRestoreFunctionVersionAlphaInput("VersionId_example") // RestoreFunctionVersionAlphaInput | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    resp, r, err := apiClient.FunctionsAPI.RestoreFunctionVersion(context.Background(), functionId).RestoreFunctionVersionAlphaInput(restoreFunctionVersionAlphaInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAPI.RestoreFunctionVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RestoreFunctionVersion`: RestoreFunctionVersion200Response
    fmt.Fprintf(os.Stdout, "Response from `FunctionsAPI.RestoreFunctionVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**functionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestoreFunctionVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **restoreFunctionVersionAlphaInput** | [**RestoreFunctionVersionAlphaInput**](RestoreFunctionVersionAlphaInput.md) |  | 

### Return type

[**RestoreFunctionVersion200Response**](RestoreFunctionVersion200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/vnd.segment.v1alpha+json
- **Accept**: application/vnd.segment.v1alpha+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: UpdateFunction

> UpdateFunction200Response UpdateFunction(ctx, functionId).UpdateFunctionV1Input(updateFunctionV1Input).Execute()

Update Function



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
    functionId := "sfnc_wXzcDGFR3KmjLDrtSawNHf" // string | 
    updateFunctionV1Input := *openapiclient.NewUpdateFunctionV1Input() // UpdateFunctionV1Input | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    resp, r, err := apiClient.FunctionsAPI.UpdateFunction(context.Background(), functionId).UpdateFunctionV1Input(updateFunctionV1Input).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAPI.UpdateFunction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateFunction`: UpdateFunction200Response
    fmt.Fprintf(os.Stdout, "Response from `FunctionsAPI.UpdateFunction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**functionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFunctionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateFunctionV1Input** | [**UpdateFunctionV1Input**](UpdateFunctionV1Input.md) |  | 

### Return type

[**UpdateFunction200Response**](UpdateFunction200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/vnd.segment.v1+json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json
- **Accept**: application/vnd.segment.v1+json, application/json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: UpdateInsertFunctionInstance

> UpdateInsertFunctionInstance200Response UpdateInsertFunctionInstance(ctx, instanceId).UpdateInsertFunctionInstanceAlphaInput(updateInsertFunctionInstanceAlphaInput).Execute()

Update Insert Function Instance



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
    instanceId := "64edfa58c43ed978ee4f2b7e" // string | 
    updateInsertFunctionInstanceAlphaInput := *openapiclient.NewUpdateInsertFunctionInstanceAlphaInput(map[string]interface{}{"key": interface{}(123)}) // UpdateInsertFunctionInstanceAlphaInput | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    resp, r, err := apiClient.FunctionsAPI.UpdateInsertFunctionInstance(context.Background(), instanceId).UpdateInsertFunctionInstanceAlphaInput(updateInsertFunctionInstanceAlphaInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAPI.UpdateInsertFunctionInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateInsertFunctionInstance`: UpdateInsertFunctionInstance200Response
    fmt.Fprintf(os.Stdout, "Response from `FunctionsAPI.UpdateInsertFunctionInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateInsertFunctionInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateInsertFunctionInstanceAlphaInput** | [**UpdateInsertFunctionInstanceAlphaInput**](UpdateInsertFunctionInstanceAlphaInput.md) |  | 

### Return type

[**UpdateInsertFunctionInstance200Response**](UpdateInsertFunctionInstance200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/vnd.segment.v1alpha+json
- **Accept**: application/vnd.segment.v1alpha+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

