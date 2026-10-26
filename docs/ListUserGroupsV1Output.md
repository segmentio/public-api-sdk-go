# ListUserGroupsV1Output

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserGroups** | [**[]UserGroupV1**](UserGroupV1.md) | The user group returned from the query. | 
**Pagination** | [**PaginationOutput**](PaginationOutput.md) |  | 

## Methods

### NewListUserGroupsV1Output

`func NewListUserGroupsV1Output(userGroups []UserGroupV1, pagination PaginationOutput, ) *ListUserGroupsV1Output`

NewListUserGroupsV1Output instantiates a new ListUserGroupsV1Output object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListUserGroupsV1OutputWithDefaults

`func NewListUserGroupsV1OutputWithDefaults() *ListUserGroupsV1Output`

NewListUserGroupsV1OutputWithDefaults instantiates a new ListUserGroupsV1Output object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserGroups

`func (o *ListUserGroupsV1Output) GetUserGroups() []UserGroupV1`

GetUserGroups returns the UserGroups field if non-nil, zero value otherwise.

### GetUserGroupsOk

`func (o *ListUserGroupsV1Output) GetUserGroupsOk() (*[]UserGroupV1, bool)`

GetUserGroupsOk returns a tuple with the UserGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroups

`func (o *ListUserGroupsV1Output) SetUserGroups(v []UserGroupV1)`

SetUserGroups sets UserGroups field to given value.


### GetPagination

`func (o *ListUserGroupsV1Output) GetPagination() PaginationOutput`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListUserGroupsV1Output) GetPaginationOk() (*PaginationOutput, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListUserGroupsV1Output) SetPagination(v PaginationOutput)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


