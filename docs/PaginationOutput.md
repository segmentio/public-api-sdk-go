# PaginationOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Current** | **string** | The current cursor within a collection.  Consumers of the API must treat this value as opaque. | 
**Next** | Pointer to **NullableString** | A pointer to the next page.  This does not return when you retrieve the last page.  Consumers of the API must treat this value as opaque. | [optional] 
**Previous** | Pointer to **NullableString** | A pointer to the previous page.  This does not return when you retrieve the first page.  Consumers of the API must treat this value as opaque. | [optional] 
**TotalEntries** | Pointer to **float32** | The total number of entries available in the collection.  If calculating it impacts performance, the response may omit this field. | [optional] 

## Methods

### NewPaginationOutput

`func NewPaginationOutput(current string, ) *PaginationOutput`

NewPaginationOutput instantiates a new PaginationOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginationOutputWithDefaults

`func NewPaginationOutputWithDefaults() *PaginationOutput`

NewPaginationOutputWithDefaults instantiates a new PaginationOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrent

`func (o *PaginationOutput) GetCurrent() string`

GetCurrent returns the Current field if non-nil, zero value otherwise.

### GetCurrentOk

`func (o *PaginationOutput) GetCurrentOk() (*string, bool)`

GetCurrentOk returns a tuple with the Current field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrent

`func (o *PaginationOutput) SetCurrent(v string)`

SetCurrent sets Current field to given value.


### GetNext

`func (o *PaginationOutput) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *PaginationOutput) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *PaginationOutput) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *PaginationOutput) HasNext() bool`

HasNext returns a boolean if a field has been set.

### SetNextNil

`func (o *PaginationOutput) SetNextNil(b bool)`

 SetNextNil sets the value for Next to be an explicit nil

### UnsetNext
`func (o *PaginationOutput) UnsetNext()`

UnsetNext ensures that no value is present for Next, not even an explicit nil
### GetPrevious

`func (o *PaginationOutput) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *PaginationOutput) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *PaginationOutput) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *PaginationOutput) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### SetPreviousNil

`func (o *PaginationOutput) SetPreviousNil(b bool)`

 SetPreviousNil sets the value for Previous to be an explicit nil

### UnsetPrevious
`func (o *PaginationOutput) UnsetPrevious()`

UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
### GetTotalEntries

`func (o *PaginationOutput) GetTotalEntries() float32`

GetTotalEntries returns the TotalEntries field if non-nil, zero value otherwise.

### GetTotalEntriesOk

`func (o *PaginationOutput) GetTotalEntriesOk() (*float32, bool)`

GetTotalEntriesOk returns a tuple with the TotalEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalEntries

`func (o *PaginationOutput) SetTotalEntries(v float32)`

SetTotalEntries sets TotalEntries field to given value.

### HasTotalEntries

`func (o *PaginationOutput) HasTotalEntries() bool`

HasTotalEntries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


