# ListUserGroupsFromUserV1Output

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Groups** | [**[]MinimalUserGroupV1**](MinimalUserGroupV1.md) | The user groups that the user belongs to. | 
**Pagination** | [**PaginationOutput**](PaginationOutput.md) |  | 

## Methods

### NewListUserGroupsFromUserV1Output

`func NewListUserGroupsFromUserV1Output(groups []MinimalUserGroupV1, pagination PaginationOutput, ) *ListUserGroupsFromUserV1Output`

NewListUserGroupsFromUserV1Output instantiates a new ListUserGroupsFromUserV1Output object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListUserGroupsFromUserV1OutputWithDefaults

`func NewListUserGroupsFromUserV1OutputWithDefaults() *ListUserGroupsFromUserV1Output`

NewListUserGroupsFromUserV1OutputWithDefaults instantiates a new ListUserGroupsFromUserV1Output object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroups

`func (o *ListUserGroupsFromUserV1Output) GetGroups() []MinimalUserGroupV1`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *ListUserGroupsFromUserV1Output) GetGroupsOk() (*[]MinimalUserGroupV1, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *ListUserGroupsFromUserV1Output) SetGroups(v []MinimalUserGroupV1)`

SetGroups sets Groups field to given value.


### GetPagination

`func (o *ListUserGroupsFromUserV1Output) GetPagination() PaginationOutput`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListUserGroupsFromUserV1Output) GetPaginationOk() (*PaginationOutput, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListUserGroupsFromUserV1Output) SetPagination(v PaginationOutput)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


