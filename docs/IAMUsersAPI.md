# \IAMUsersAPI

All URIs are relative to *https://api.segmentapis.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddPermissionsToUser**](IAMUsersAPI.md#AddPermissionsToUser) | **Post** /users/{userId}/permissions | Add Permissions to User
[**CreateInvites**](IAMUsersAPI.md#CreateInvites) | **Post** /invites | Create Invites
[**DeleteInvites**](IAMUsersAPI.md#DeleteInvites) | **Delete** /invites | Delete Invites
[**DeleteUsers**](IAMUsersAPI.md#DeleteUsers) | **Delete** /users | Delete Users
[**GetUser**](IAMUsersAPI.md#GetUser) | **Get** /users/{userId} | Get User
[**ListInvites**](IAMUsersAPI.md#ListInvites) | **Get** /invites | List Invites
[**ListUserGroupsFromUser**](IAMUsersAPI.md#ListUserGroupsFromUser) | **Get** /users/{userId}/groups | List User Groups from User
[**ListUsers**](IAMUsersAPI.md#ListUsers) | **Get** /users | List Users
[**ReplacePermissionsForUser**](IAMUsersAPI.md#ReplacePermissionsForUser) | **Put** /users/{userId}/permissions | Replace Permissions for User



## Operation: AddPermissionsToUser

> AddPermissionsToUser200Response AddPermissionsToUser(ctx, userId).AddPermissionsToUserV1Input(addPermissionsToUserV1Input).Execute()

Add Permissions to User



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
    userId := "sgJDWk3K21k6LE3tLU9nRK" // string | 
    addPermissionsToUserV1Input := *api.NewAddPermissionsToUserV1Input([]api.PermissionInputV1{*api.NewPermissionInputV1("RoleId_example", []api.PermissionResourceV1{*api.NewPermissionResourceV1("Id_example", "Type_example")})}) // AddPermissionsToUserV1Input | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.IAMUsersAPI.AddPermissionsToUser(ctx, userId).AddPermissionsToUserV1Input(addPermissionsToUserV1Input).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IAMUsersAPI.AddPermissionsToUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `AddPermissionsToUser`: AddPermissionsToUser200Response
    fmt.Fprintf(os.Stdout, "Response from `IAMUsersAPI.AddPermissionsToUser`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddPermissionsToUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addPermissionsToUserV1Input** | [**AddPermissionsToUserV1Input**](AddPermissionsToUserV1Input.md) |  | 

### Return type

[**AddPermissionsToUser200Response**](AddPermissionsToUser200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/vnd.segment.v1+json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json
- **Accept**: application/vnd.segment.v1+json, application/json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: CreateInvites

> CreateInvites200Response CreateInvites(ctx).CreateInvitesV1Input(createInvitesV1Input).Execute()

Create Invites



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
    createInvitesV1Input := *api.NewCreateInvitesV1Input([]api.InviteV1{*api.NewInviteV1("Email_example")}) // CreateInvitesV1Input | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.IAMUsersAPI.CreateInvites(ctx).CreateInvitesV1Input(createInvitesV1Input).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IAMUsersAPI.CreateInvites``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `CreateInvites`: CreateInvites200Response
    fmt.Fprintf(os.Stdout, "Response from `IAMUsersAPI.CreateInvites`: %v\n", resp.GetData())
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateInvitesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createInvitesV1Input** | [**CreateInvitesV1Input**](CreateInvitesV1Input.md) |  | 

### Return type

[**CreateInvites200Response**](CreateInvites200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/vnd.segment.v1+json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json
- **Accept**: application/vnd.segment.v1+json, application/json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: DeleteInvites

> DeleteInvites200Response DeleteInvites(ctx).Emails(emails).Execute()

Delete Invites



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
    emails := []string{"Inner_example"} // []string | The list of emails to delete invites for.  This parameter exists in v1.

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.IAMUsersAPI.DeleteInvites(ctx).Emails(emails).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IAMUsersAPI.DeleteInvites``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `DeleteInvites`: DeleteInvites200Response
    fmt.Fprintf(os.Stdout, "Response from `IAMUsersAPI.DeleteInvites`: %v\n", resp.GetData())
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteInvitesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **emails** | **[]string** | The list of emails to delete invites for.  This parameter exists in v1. | 

### Return type

[**DeleteInvites200Response**](DeleteInvites200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1+json, application/json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: DeleteUsers

> DeleteUsers200Response DeleteUsers(ctx).UserIds(userIds).Execute()

Delete Users



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
    userIds := []string{"Inner_example"} // []string | The ids of the users to remove.  This parameter exists in v1.

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.IAMUsersAPI.DeleteUsers(ctx).UserIds(userIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IAMUsersAPI.DeleteUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `DeleteUsers`: DeleteUsers200Response
    fmt.Fprintf(os.Stdout, "Response from `IAMUsersAPI.DeleteUsers`: %v\n", resp.GetData())
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userIds** | **[]string** | The ids of the users to remove.  This parameter exists in v1. | 

### Return type

[**DeleteUsers200Response**](DeleteUsers200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1+json, application/json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: GetUser

> GetUser200Response GetUser(ctx, userId).Execute()

Get User



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
    userId := "sgJDWk3K21k6LE3tLU9nRK" // string | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.IAMUsersAPI.GetUser(ctx, userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IAMUsersAPI.GetUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `GetUser`: GetUser200Response
    fmt.Fprintf(os.Stdout, "Response from `IAMUsersAPI.GetUser`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetUser200Response**](GetUser200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1+json, application/json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: ListInvites

> ListInvites200Response ListInvites(ctx).Pagination(pagination).Execute()

List Invites



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
    pagination := *api.NewPaginationInput(10) // PaginationInput | Defines the pagination parameters.  This parameter exists in v1.

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.IAMUsersAPI.ListInvites(ctx).Pagination(pagination).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IAMUsersAPI.ListInvites``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `ListInvites`: ListInvites200Response
    fmt.Fprintf(os.Stdout, "Response from `IAMUsersAPI.ListInvites`: %v\n", resp.GetData())
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListInvitesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pagination** | [**PaginationInput**](PaginationInput.md) | Defines the pagination parameters.  This parameter exists in v1. | 

### Return type

[**ListInvites200Response**](ListInvites200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1+json, application/json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: ListUserGroupsFromUser

> ListUserGroupsFromUser200Response ListUserGroupsFromUser(ctx, userId).Pagination(pagination).Execute()

List User Groups from User



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
    userId := "sgJDWk3K21k6LE3tLU9nRK" // string | 
    pagination := *api.NewPaginationInput(10) // PaginationInput | Pagination for groups.  This parameter exists in v1.

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.IAMUsersAPI.ListUserGroupsFromUser(ctx, userId).Pagination(pagination).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IAMUsersAPI.ListUserGroupsFromUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `ListUserGroupsFromUser`: ListUserGroupsFromUser200Response
    fmt.Fprintf(os.Stdout, "Response from `IAMUsersAPI.ListUserGroupsFromUser`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListUserGroupsFromUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pagination** | [**PaginationInput**](PaginationInput.md) | Pagination for groups.  This parameter exists in v1. | 

### Return type

[**ListUserGroupsFromUser200Response**](ListUserGroupsFromUser200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1+json, application/json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: ListUsers

> ListUsers200Response ListUsers(ctx).Pagination(pagination).Execute()

List Users



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
    pagination := *api.NewPaginationInput(10) // PaginationInput | Pagination for users.  This parameter exists in v1.

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.IAMUsersAPI.ListUsers(ctx).Pagination(pagination).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IAMUsersAPI.ListUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `ListUsers`: ListUsers200Response
    fmt.Fprintf(os.Stdout, "Response from `IAMUsersAPI.ListUsers`: %v\n", resp.GetData())
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pagination** | [**PaginationInput**](PaginationInput.md) | Pagination for users.  This parameter exists in v1. | 

### Return type

[**ListUsers200Response**](ListUsers200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1+json, application/json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: ReplacePermissionsForUser

> ReplacePermissionsForUser200Response ReplacePermissionsForUser(ctx, userId).ReplacePermissionsForUserV1Input(replacePermissionsForUserV1Input).Execute()

Replace Permissions for User



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
    userId := "sgJDWk3K21k6LE3tLU9nRK" // string | 
    replacePermissionsForUserV1Input := *api.NewReplacePermissionsForUserV1Input([]api.PermissionInputV1{*api.NewPermissionInputV1("RoleId_example", []api.PermissionResourceV1{*api.NewPermissionResourceV1("Id_example", "Type_example")})}) // ReplacePermissionsForUserV1Input | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "<BEARER_TOKEN>"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.IAMUsersAPI.ReplacePermissionsForUser(ctx, userId).ReplacePermissionsForUserV1Input(replacePermissionsForUserV1Input).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IAMUsersAPI.ReplacePermissionsForUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `ReplacePermissionsForUser`: ReplacePermissionsForUser200Response
    fmt.Fprintf(os.Stdout, "Response from `IAMUsersAPI.ReplacePermissionsForUser`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplacePermissionsForUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **replacePermissionsForUserV1Input** | [**ReplacePermissionsForUserV1Input**](ReplacePermissionsForUserV1Input.md) |  | 

### Return type

[**ReplacePermissionsForUser200Response**](ReplacePermissionsForUser200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/vnd.segment.v1+json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json
- **Accept**: application/vnd.segment.v1+json, application/json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

