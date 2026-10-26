# ListEventsAlphaOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | [**[]EventAlpha**](EventAlpha.md) | List of events in the space. | 
**Pagination** | [**PaginationOutput**](PaginationOutput.md) |  | 

## Methods

### NewListEventsAlphaOutput

`func NewListEventsAlphaOutput(events []EventAlpha, pagination PaginationOutput, ) *ListEventsAlphaOutput`

NewListEventsAlphaOutput instantiates a new ListEventsAlphaOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListEventsAlphaOutputWithDefaults

`func NewListEventsAlphaOutputWithDefaults() *ListEventsAlphaOutput`

NewListEventsAlphaOutputWithDefaults instantiates a new ListEventsAlphaOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *ListEventsAlphaOutput) GetEvents() []EventAlpha`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *ListEventsAlphaOutput) GetEventsOk() (*[]EventAlpha, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *ListEventsAlphaOutput) SetEvents(v []EventAlpha)`

SetEvents sets Events field to given value.


### GetPagination

`func (o *ListEventsAlphaOutput) GetPagination() PaginationOutput`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListEventsAlphaOutput) GetPaginationOk() (*PaginationOutput, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListEventsAlphaOutput) SetPagination(v PaginationOutput)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


