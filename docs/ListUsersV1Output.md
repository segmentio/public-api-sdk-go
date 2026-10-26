# ListUsersV1Output

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Users** | [**[]UserV1**](UserV1.md) | The list of users. | 
**Pagination** | [**PaginationOutput**](PaginationOutput.md) |  | 

## Methods

### NewListUsersV1Output

`func NewListUsersV1Output(users []UserV1, pagination PaginationOutput, ) *ListUsersV1Output`

NewListUsersV1Output instantiates a new ListUsersV1Output object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListUsersV1OutputWithDefaults

`func NewListUsersV1OutputWithDefaults() *ListUsersV1Output`

NewListUsersV1OutputWithDefaults instantiates a new ListUsersV1Output object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsers

`func (o *ListUsersV1Output) GetUsers() []UserV1`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *ListUsersV1Output) GetUsersOk() (*[]UserV1, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *ListUsersV1Output) SetUsers(v []UserV1)`

SetUsers sets Users field to given value.


### GetPagination

`func (o *ListUsersV1Output) GetPagination() PaginationOutput`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListUsersV1Output) GetPaginationOk() (*PaginationOutput, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListUsersV1Output) SetPagination(v PaginationOutput)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


