# ListInvitesV1Output

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Invites** | **[]string** | The list of invites. | 
**Pagination** | [**PaginationOutput**](PaginationOutput.md) |  | 

## Methods

### NewListInvitesV1Output

`func NewListInvitesV1Output(invites []string, pagination PaginationOutput, ) *ListInvitesV1Output`

NewListInvitesV1Output instantiates a new ListInvitesV1Output object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListInvitesV1OutputWithDefaults

`func NewListInvitesV1OutputWithDefaults() *ListInvitesV1Output`

NewListInvitesV1OutputWithDefaults instantiates a new ListInvitesV1Output object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvites

`func (o *ListInvitesV1Output) GetInvites() []string`

GetInvites returns the Invites field if non-nil, zero value otherwise.

### GetInvitesOk

`func (o *ListInvitesV1Output) GetInvitesOk() (*[]string, bool)`

GetInvitesOk returns a tuple with the Invites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvites

`func (o *ListInvitesV1Output) SetInvites(v []string)`

SetInvites sets Invites field to given value.


### GetPagination

`func (o *ListInvitesV1Output) GetPagination() PaginationOutput`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListInvitesV1Output) GetPaginationOk() (*PaginationOutput, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListInvitesV1Output) SetPagination(v PaginationOutput)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


