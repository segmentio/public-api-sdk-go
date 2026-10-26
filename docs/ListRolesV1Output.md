# ListRolesV1Output

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Roles** | [**[]RoleV1**](RoleV1.md) | The list of roles. | 
**Pagination** | [**PaginationOutput**](PaginationOutput.md) |  | 

## Methods

### NewListRolesV1Output

`func NewListRolesV1Output(roles []RoleV1, pagination PaginationOutput, ) *ListRolesV1Output`

NewListRolesV1Output instantiates a new ListRolesV1Output object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListRolesV1OutputWithDefaults

`func NewListRolesV1OutputWithDefaults() *ListRolesV1Output`

NewListRolesV1OutputWithDefaults instantiates a new ListRolesV1Output object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoles

`func (o *ListRolesV1Output) GetRoles() []RoleV1`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *ListRolesV1Output) GetRolesOk() (*[]RoleV1, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *ListRolesV1Output) SetRoles(v []RoleV1)`

SetRoles sets Roles field to given value.


### GetPagination

`func (o *ListRolesV1Output) GetPagination() PaginationOutput`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListRolesV1Output) GetPaginationOk() (*PaginationOutput, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListRolesV1Output) SetPagination(v PaginationOutput)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


