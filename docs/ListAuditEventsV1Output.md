# ListAuditEventsV1Output

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | [**[]AuditEventV1**](AuditEventV1.md) | Audit trail events for the current Workspace. | 
**Pagination** | Pointer to [**PaginationOutput**](PaginationOutput.md) |  | [optional] 

## Methods

### NewListAuditEventsV1Output

`func NewListAuditEventsV1Output(events []AuditEventV1, ) *ListAuditEventsV1Output`

NewListAuditEventsV1Output instantiates a new ListAuditEventsV1Output object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAuditEventsV1OutputWithDefaults

`func NewListAuditEventsV1OutputWithDefaults() *ListAuditEventsV1Output`

NewListAuditEventsV1OutputWithDefaults instantiates a new ListAuditEventsV1Output object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *ListAuditEventsV1Output) GetEvents() []AuditEventV1`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *ListAuditEventsV1Output) GetEventsOk() (*[]AuditEventV1, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *ListAuditEventsV1Output) SetEvents(v []AuditEventV1)`

SetEvents sets Events field to given value.


### GetPagination

`func (o *ListAuditEventsV1Output) GetPagination() PaginationOutput`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListAuditEventsV1Output) GetPaginationOk() (*PaginationOutput, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListAuditEventsV1Output) SetPagination(v PaginationOutput)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *ListAuditEventsV1Output) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


