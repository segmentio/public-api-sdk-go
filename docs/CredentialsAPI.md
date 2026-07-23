# \CredentialsAPI

All URIs are relative to *https://api.segmentapis.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCredential**](CredentialsAPI.md#CreateCredential) | **Post** /credentials | Create Credential
[**UpdateCredential**](CredentialsAPI.md#UpdateCredential) | **Patch** /credentials/{credentialId} | Update Credential



## Operation: CreateCredential

> CreateCredential201Response CreateCredential(ctx).CreateCredentialV1Input(createCredentialV1Input).Execute()

Create Credential



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
    createCredentialV1Input := *api.NewCreateCredentialV1Input("Name_example", map[string]interface{}{"key": interface{}(123)}) // CreateCredentialV1Input | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.CredentialsAPI.CreateCredential(ctx).CreateCredentialV1Input(createCredentialV1Input).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CredentialsAPI.CreateCredential``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `CreateCredential`: CreateCredential201Response
    fmt.Fprintf(os.Stdout, "Response from `CredentialsAPI.CreateCredential`: %v\n", resp.GetData())
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createCredentialV1Input** | [**CreateCredentialV1Input**](CreateCredentialV1Input.md) |  | 

### Return type

[**CreateCredential201Response**](CreateCredential201Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json, application/vnd.segment.v1+json
- **Accept**: application/vnd.segment.v1+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: UpdateCredential

> UpdateCredential200Response UpdateCredential(ctx, credentialId).UpdateCredentialV1Input(updateCredentialV1Input).Execute()

Update Credential



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
    credentialId := "cred_2JzKWb8FGhGVYZ3xVqQGc7NkYPl" // string | 
    updateCredentialV1Input := *api.NewUpdateCredentialV1Input() // UpdateCredentialV1Input | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.CredentialsAPI.UpdateCredential(ctx, credentialId).UpdateCredentialV1Input(updateCredentialV1Input).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CredentialsAPI.UpdateCredential``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `UpdateCredential`: UpdateCredential200Response
    fmt.Fprintf(os.Stdout, "Response from `CredentialsAPI.UpdateCredential`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**credentialId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateCredentialV1Input** | [**UpdateCredentialV1Input**](UpdateCredentialV1Input.md) |  | 

### Return type

[**UpdateCredential200Response**](UpdateCredential200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json, application/vnd.segment.v1+json
- **Accept**: application/vnd.segment.v1+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

