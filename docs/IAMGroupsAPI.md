# \IAMGroupsAPI

All URIs are relative to *https://api.segmentapis.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddPermissionsToUserGroup**](IAMGroupsAPI.md#AddPermissionsToUserGroup) | **Post** /groups/{userGroupId}/permissions | Add Permissions to User Group
[**AddUsersToUserGroup**](IAMGroupsAPI.md#AddUsersToUserGroup) | **Post** /groups/{userGroupId}/users | Add Users to User Group
[**CreateUserGroup**](IAMGroupsAPI.md#CreateUserGroup) | **Post** /groups | Create User Group
[**DeleteUserGroup**](IAMGroupsAPI.md#DeleteUserGroup) | **Delete** /groups/{userGroupId} | Delete User Group
[**GetUserGroup**](IAMGroupsAPI.md#GetUserGroup) | **Get** /groups/{userGroupId} | Get User Group
[**ListInvitesFromUserGroup**](IAMGroupsAPI.md#ListInvitesFromUserGroup) | **Get** /groups/{userGroupId}/invites | List Invites from User Group
[**ListUserGroups**](IAMGroupsAPI.md#ListUserGroups) | **Get** /groups | List User Groups
[**ListUsersFromUserGroup**](IAMGroupsAPI.md#ListUsersFromUserGroup) | **Get** /groups/{userGroupId}/users | List Users from User Group
[**RemoveUsersFromUserGroup**](IAMGroupsAPI.md#RemoveUsersFromUserGroup) | **Delete** /group/{userGroupId}/users | Remove Users from User Group
[**ReplacePermissionsForUserGroup**](IAMGroupsAPI.md#ReplacePermissionsForUserGroup) | **Put** /groups/{userGroupId}/permissions | Replace Permissions for User Group
[**ReplaceUsersInUserGroup**](IAMGroupsAPI.md#ReplaceUsersInUserGroup) | **Put** /group/{userGroupId}/users | Replace Users in User Group
[**UpdateUserGroup**](IAMGroupsAPI.md#UpdateUserGroup) | **Patch** /groups/{userGroupId} | Update User Group



## Operation: AddPermissionsToUserGroup

> AddPermissionsToUserGroup200Response AddPermissionsToUserGroup(ctx, userGroupId).AddPermissionsToUserGroupV1Input(addPermissionsToUserGroupV1Input).Execute()

Add Permissions to User Group



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
    userGroupId := "bBABwqbaDf2QdwTbW8bNEm" // string | 
    addPermissionsToUserGroupV1Input := *api.NewAddPermissionsToUserGroupV1Input([]api.PermissionInputV1{*api.NewPermissionInputV1("RoleId_example", []api.PermissionResourceV1{*api.NewPermissionResourceV1("Id_example", "Type_example")})}) // AddPermissionsToUserGroupV1Input | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "INSERT BEARER TOKEN HERE"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.IAMGroupsAPI.AddPermissionsToUserGroup(ctx, userGroupId).AddPermissionsToUserGroupV1Input(addPermissionsToUserGroupV1Input).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IAMGroupsAPI.AddPermissionsToUserGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `AddPermissionsToUserGroup`: AddPermissionsToUserGroup200Response
    fmt.Fprintf(os.Stdout, "Response from `IAMGroupsAPI.AddPermissionsToUserGroup`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGroupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddPermissionsToUserGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addPermissionsToUserGroupV1Input** | [**AddPermissionsToUserGroupV1Input**](AddPermissionsToUserGroupV1Input.md) |  | 

### Return type

[**AddPermissionsToUserGroup200Response**](AddPermissionsToUserGroup200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/vnd.segment.v1+json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json
- **Accept**: application/vnd.segment.v1+json, application/json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: AddUsersToUserGroup

> AddUsersToUserGroup200Response AddUsersToUserGroup(ctx, userGroupId).AddUsersToUserGroupV1Input(addUsersToUserGroupV1Input).Execute()

Add Users to User Group



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
    userGroupId := "bBABwqbaDf2QdwTbW8bNEm" // string | 
    addUsersToUserGroupV1Input := *api.NewAddUsersToUserGroupV1Input([]string{"Emails_example"}) // AddUsersToUserGroupV1Input | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "INSERT BEARER TOKEN HERE"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.IAMGroupsAPI.AddUsersToUserGroup(ctx, userGroupId).AddUsersToUserGroupV1Input(addUsersToUserGroupV1Input).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IAMGroupsAPI.AddUsersToUserGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `AddUsersToUserGroup`: AddUsersToUserGroup200Response
    fmt.Fprintf(os.Stdout, "Response from `IAMGroupsAPI.AddUsersToUserGroup`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGroupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddUsersToUserGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addUsersToUserGroupV1Input** | [**AddUsersToUserGroupV1Input**](AddUsersToUserGroupV1Input.md) |  | 

### Return type

[**AddUsersToUserGroup200Response**](AddUsersToUserGroup200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/vnd.segment.v1+json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json
- **Accept**: application/vnd.segment.v1+json, application/json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: CreateUserGroup

> CreateUserGroup200Response CreateUserGroup(ctx).CreateUserGroupV1Input(createUserGroupV1Input).Execute()

Create User Group



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
    createUserGroupV1Input := *api.NewCreateUserGroupV1Input("Name_example") // CreateUserGroupV1Input | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "INSERT BEARER TOKEN HERE"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.IAMGroupsAPI.CreateUserGroup(ctx).CreateUserGroupV1Input(createUserGroupV1Input).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IAMGroupsAPI.CreateUserGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `CreateUserGroup`: CreateUserGroup200Response
    fmt.Fprintf(os.Stdout, "Response from `IAMGroupsAPI.CreateUserGroup`: %v\n", resp.GetData())
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createUserGroupV1Input** | [**CreateUserGroupV1Input**](CreateUserGroupV1Input.md) |  | 

### Return type

[**CreateUserGroup200Response**](CreateUserGroup200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/vnd.segment.v1+json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json
- **Accept**: application/vnd.segment.v1+json, application/json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: DeleteUserGroup

> DeleteUserGroup200Response DeleteUserGroup(ctx, userGroupId).Execute()

Delete User Group



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
    userGroupId := "2Nhj3GVumKbR0cjDly4nCdGo6cT" // string | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "INSERT BEARER TOKEN HERE"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.IAMGroupsAPI.DeleteUserGroup(ctx, userGroupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IAMGroupsAPI.DeleteUserGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `DeleteUserGroup`: DeleteUserGroup200Response
    fmt.Fprintf(os.Stdout, "Response from `IAMGroupsAPI.DeleteUserGroup`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGroupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteUserGroup200Response**](DeleteUserGroup200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1+json, application/json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: GetUserGroup

> GetUserGroup200Response GetUserGroup(ctx, userGroupId).Execute()

Get User Group



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
    userGroupId := "bBABwqbaDf2QdwTbW8bNEm" // string | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "INSERT BEARER TOKEN HERE"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.IAMGroupsAPI.GetUserGroup(ctx, userGroupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IAMGroupsAPI.GetUserGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `GetUserGroup`: GetUserGroup200Response
    fmt.Fprintf(os.Stdout, "Response from `IAMGroupsAPI.GetUserGroup`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGroupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetUserGroup200Response**](GetUserGroup200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1+json, application/json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: ListInvitesFromUserGroup

> ListInvitesFromUserGroup200Response ListInvitesFromUserGroup(ctx, userGroupId).Pagination(pagination).Execute()

List Invites from User Group



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
    userGroupId := "bBABwqbaDf2QdwTbW8bNEm" // string | 
    pagination := *api.NewPaginationInput(float32(123)) // PaginationInput | Pagination for invites to the group.  This parameter exists in v1.

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "INSERT BEARER TOKEN HERE"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.IAMGroupsAPI.ListInvitesFromUserGroup(ctx, userGroupId).Pagination(pagination).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IAMGroupsAPI.ListInvitesFromUserGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `ListInvitesFromUserGroup`: ListInvitesFromUserGroup200Response
    fmt.Fprintf(os.Stdout, "Response from `IAMGroupsAPI.ListInvitesFromUserGroup`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGroupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListInvitesFromUserGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pagination** | [**PaginationInput**](PaginationInput.md) | Pagination for invites to the group.  This parameter exists in v1. | 

### Return type

[**ListInvitesFromUserGroup200Response**](ListInvitesFromUserGroup200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1+json, application/json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: ListUserGroups

> ListUserGroups200Response ListUserGroups(ctx).Pagination(pagination).Execute()

List User Groups



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
    pagination := *api.NewPaginationInput(float32(123)) // PaginationInput | Pagination for user groups.  This parameter exists in v1.

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "INSERT BEARER TOKEN HERE"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.IAMGroupsAPI.ListUserGroups(ctx).Pagination(pagination).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IAMGroupsAPI.ListUserGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `ListUserGroups`: ListUserGroups200Response
    fmt.Fprintf(os.Stdout, "Response from `IAMGroupsAPI.ListUserGroups`: %v\n", resp.GetData())
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListUserGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pagination** | [**PaginationInput**](PaginationInput.md) | Pagination for user groups.  This parameter exists in v1. | 

### Return type

[**ListUserGroups200Response**](ListUserGroups200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1+json, application/json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: ListUsersFromUserGroup

> ListUsersFromUserGroup200Response ListUsersFromUserGroup(ctx, userGroupId).Pagination(pagination).Execute()

List Users from User Group



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
    userGroupId := "bBABwqbaDf2QdwTbW8bNEm" // string | 
    pagination := *api.NewPaginationInput(float32(123)) // PaginationInput | Pagination for members of a group.  This parameter exists in v1.

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "INSERT BEARER TOKEN HERE"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.IAMGroupsAPI.ListUsersFromUserGroup(ctx, userGroupId).Pagination(pagination).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IAMGroupsAPI.ListUsersFromUserGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `ListUsersFromUserGroup`: ListUsersFromUserGroup200Response
    fmt.Fprintf(os.Stdout, "Response from `IAMGroupsAPI.ListUsersFromUserGroup`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGroupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListUsersFromUserGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pagination** | [**PaginationInput**](PaginationInput.md) | Pagination for members of a group.  This parameter exists in v1. | 

### Return type

[**ListUsersFromUserGroup200Response**](ListUsersFromUserGroup200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1+json, application/json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: RemoveUsersFromUserGroup

> RemoveUsersFromUserGroup200Response RemoveUsersFromUserGroup(ctx, userGroupId).Emails(emails).Execute()

Remove Users from User Group



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
    userGroupId := "bBABwqbaDf2QdwTbW8bNEm" // string | 
    emails := []string{[]string{"Emails_example"}} // []string | The list of emails to remove from the user group.  This parameter exists in v1.

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "INSERT BEARER TOKEN HERE"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.IAMGroupsAPI.RemoveUsersFromUserGroup(ctx, userGroupId).Emails(emails).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IAMGroupsAPI.RemoveUsersFromUserGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `RemoveUsersFromUserGroup`: RemoveUsersFromUserGroup200Response
    fmt.Fprintf(os.Stdout, "Response from `IAMGroupsAPI.RemoveUsersFromUserGroup`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGroupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveUsersFromUserGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **emails** | **[][]string** | The list of emails to remove from the user group.  This parameter exists in v1. | 

### Return type

[**RemoveUsersFromUserGroup200Response**](RemoveUsersFromUserGroup200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.segment.v1+json, application/json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: ReplacePermissionsForUserGroup

> ReplacePermissionsForUserGroup200Response ReplacePermissionsForUserGroup(ctx, userGroupId).ReplacePermissionsForUserGroupV1Input(replacePermissionsForUserGroupV1Input).Execute()

Replace Permissions for User Group



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
    userGroupId := "bBABwqbaDf2QdwTbW8bNEm" // string | 
    replacePermissionsForUserGroupV1Input := *api.NewReplacePermissionsForUserGroupV1Input([]api.PermissionInputV1{*api.NewPermissionInputV1("RoleId_example", []api.PermissionResourceV1{*api.NewPermissionResourceV1("Id_example", "Type_example")})}) // ReplacePermissionsForUserGroupV1Input | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "INSERT BEARER TOKEN HERE"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.IAMGroupsAPI.ReplacePermissionsForUserGroup(ctx, userGroupId).ReplacePermissionsForUserGroupV1Input(replacePermissionsForUserGroupV1Input).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IAMGroupsAPI.ReplacePermissionsForUserGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `ReplacePermissionsForUserGroup`: ReplacePermissionsForUserGroup200Response
    fmt.Fprintf(os.Stdout, "Response from `IAMGroupsAPI.ReplacePermissionsForUserGroup`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGroupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplacePermissionsForUserGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **replacePermissionsForUserGroupV1Input** | [**ReplacePermissionsForUserGroupV1Input**](ReplacePermissionsForUserGroupV1Input.md) |  | 

### Return type

[**ReplacePermissionsForUserGroup200Response**](ReplacePermissionsForUserGroup200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/vnd.segment.v1+json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json
- **Accept**: application/vnd.segment.v1+json, application/json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: ReplaceUsersInUserGroup

> ReplaceUsersInUserGroup200Response ReplaceUsersInUserGroup(ctx, userGroupId).ReplaceUsersInUserGroupV1Input(replaceUsersInUserGroupV1Input).Execute()

Replace Users in User Group



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
    userGroupId := "bBABwqbaDf2QdwTbW8bNEm" // string | 
    replaceUsersInUserGroupV1Input := *api.NewReplaceUsersInUserGroupV1Input([]string{"Emails_example"}) // ReplaceUsersInUserGroupV1Input | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "INSERT BEARER TOKEN HERE"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.IAMGroupsAPI.ReplaceUsersInUserGroup(ctx, userGroupId).ReplaceUsersInUserGroupV1Input(replaceUsersInUserGroupV1Input).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IAMGroupsAPI.ReplaceUsersInUserGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `ReplaceUsersInUserGroup`: ReplaceUsersInUserGroup200Response
    fmt.Fprintf(os.Stdout, "Response from `IAMGroupsAPI.ReplaceUsersInUserGroup`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGroupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceUsersInUserGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **replaceUsersInUserGroupV1Input** | [**ReplaceUsersInUserGroupV1Input**](ReplaceUsersInUserGroupV1Input.md) |  | 

### Return type

[**ReplaceUsersInUserGroup200Response**](ReplaceUsersInUserGroup200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/vnd.segment.v1+json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json
- **Accept**: application/vnd.segment.v1+json, application/json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Operation: UpdateUserGroup

> UpdateUserGroup200Response UpdateUserGroup(ctx, userGroupId).UpdateUserGroupV1Input(updateUserGroupV1Input).Execute()

Update User Group



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
    userGroupId := "bBABwqbaDf2QdwTbW8bNEm" // string | 
    updateUserGroupV1Input := *api.NewUpdateUserGroupV1Input("Name_example") // UpdateUserGroupV1Input | 

    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    token := "INSERT BEARER TOKEN HERE"
    ctx := context.WithValue(context.Background(), api.ContextAccessToken, token)
    resp, r, err := apiClient.IAMGroupsAPI.UpdateUserGroup(ctx, userGroupId).UpdateUserGroupV1Input(updateUserGroupV1Input).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IAMGroupsAPI.UpdateUserGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        responseErrors := api.UnwrapFullErrors(err)
        if responseErrors != nil {
            for _, responseError := range responseErrors.Errors {
                fmt.Fprintf(os.Stderr, "Full error message: %v\n", *responseError.Message)
            }
        }
    }
    // response from `UpdateUserGroup`: UpdateUserGroup200Response
    fmt.Fprintf(os.Stdout, "Response from `IAMGroupsAPI.UpdateUserGroup`: %v\n", resp.GetData())
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userGroupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateUserGroupV1Input** | [**UpdateUserGroupV1Input**](UpdateUserGroupV1Input.md) |  | 

### Return type

[**UpdateUserGroup200Response**](UpdateUserGroup200Response.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/vnd.segment.v1+json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json
- **Accept**: application/vnd.segment.v1+json, application/json, application/vnd.segment.v1beta+json, application/vnd.segment.v1alpha+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

