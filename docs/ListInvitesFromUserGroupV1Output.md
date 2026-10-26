# ListInvitesFromUserGroupV1Output

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Emails** | **[]string** | The emails of the invitees to the user group. | 
**Pagination** | [**PaginationOutput**](PaginationOutput.md) |  | 

## Methods

### NewListInvitesFromUserGroupV1Output

`func NewListInvitesFromUserGroupV1Output(emails []string, pagination PaginationOutput, ) *ListInvitesFromUserGroupV1Output`

NewListInvitesFromUserGroupV1Output instantiates a new ListInvitesFromUserGroupV1Output object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListInvitesFromUserGroupV1OutputWithDefaults

`func NewListInvitesFromUserGroupV1OutputWithDefaults() *ListInvitesFromUserGroupV1Output`

NewListInvitesFromUserGroupV1OutputWithDefaults instantiates a new ListInvitesFromUserGroupV1Output object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmails

`func (o *ListInvitesFromUserGroupV1Output) GetEmails() []string`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *ListInvitesFromUserGroupV1Output) GetEmailsOk() (*[]string, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *ListInvitesFromUserGroupV1Output) SetEmails(v []string)`

SetEmails sets Emails field to given value.


### GetPagination

`func (o *ListInvitesFromUserGroupV1Output) GetPagination() PaginationOutput`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListInvitesFromUserGroupV1Output) GetPaginationOk() (*PaginationOutput, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListInvitesFromUserGroupV1Output) SetPagination(v PaginationOutput)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


