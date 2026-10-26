# ListUsersFromUserGroupV1Output

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Users** | [**[]MinimalUserV1**](MinimalUserV1.md) | The users of the user group. | 
**Pagination** | [**PaginationOutput**](PaginationOutput.md) |  | 

## Methods

### NewListUsersFromUserGroupV1Output

`func NewListUsersFromUserGroupV1Output(users []MinimalUserV1, pagination PaginationOutput, ) *ListUsersFromUserGroupV1Output`

NewListUsersFromUserGroupV1Output instantiates a new ListUsersFromUserGroupV1Output object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListUsersFromUserGroupV1OutputWithDefaults

`func NewListUsersFromUserGroupV1OutputWithDefaults() *ListUsersFromUserGroupV1Output`

NewListUsersFromUserGroupV1OutputWithDefaults instantiates a new ListUsersFromUserGroupV1Output object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsers

`func (o *ListUsersFromUserGroupV1Output) GetUsers() []MinimalUserV1`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *ListUsersFromUserGroupV1Output) GetUsersOk() (*[]MinimalUserV1, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *ListUsersFromUserGroupV1Output) SetUsers(v []MinimalUserV1)`

SetUsers sets Users field to given value.


### GetPagination

`func (o *ListUsersFromUserGroupV1Output) GetPagination() PaginationOutput`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListUsersFromUserGroupV1Output) GetPaginationOk() (*PaginationOutput, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListUsersFromUserGroupV1Output) SetPagination(v PaginationOutput)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


