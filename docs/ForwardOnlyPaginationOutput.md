# ForwardOnlyPaginationOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Next** | Pointer to **NullableString** | A pointer to the next page.  This does not return when you retrieve the last page.  Consumers of the API must treat this value as opaque. | [optional] 
**Current** | **string** | The current cursor within a collection.  Consumers of the API must treat this value as opaque. | 
**TotalEntries** | Pointer to **float32** | The total number of entries available in the collection.  If calculating it impacts performance, the response may omit this field. | [optional] 

## Methods

### NewForwardOnlyPaginationOutput

`func NewForwardOnlyPaginationOutput(current string, ) *ForwardOnlyPaginationOutput`

NewForwardOnlyPaginationOutput instantiates a new ForwardOnlyPaginationOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewForwardOnlyPaginationOutputWithDefaults

`func NewForwardOnlyPaginationOutputWithDefaults() *ForwardOnlyPaginationOutput`

NewForwardOnlyPaginationOutputWithDefaults instantiates a new ForwardOnlyPaginationOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNext

`func (o *ForwardOnlyPaginationOutput) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *ForwardOnlyPaginationOutput) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *ForwardOnlyPaginationOutput) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *ForwardOnlyPaginationOutput) HasNext() bool`

HasNext returns a boolean if a field has been set.

### SetNextNil

`func (o *ForwardOnlyPaginationOutput) SetNextNil(b bool)`

 SetNextNil sets the value for Next to be an explicit nil

### UnsetNext
`func (o *ForwardOnlyPaginationOutput) UnsetNext()`

UnsetNext ensures that no value is present for Next, not even an explicit nil
### GetCurrent

`func (o *ForwardOnlyPaginationOutput) GetCurrent() string`

GetCurrent returns the Current field if non-nil, zero value otherwise.

### GetCurrentOk

`func (o *ForwardOnlyPaginationOutput) GetCurrentOk() (*string, bool)`

GetCurrentOk returns a tuple with the Current field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrent

`func (o *ForwardOnlyPaginationOutput) SetCurrent(v string)`

SetCurrent sets Current field to given value.


### GetTotalEntries

`func (o *ForwardOnlyPaginationOutput) GetTotalEntries() float32`

GetTotalEntries returns the TotalEntries field if non-nil, zero value otherwise.

### GetTotalEntriesOk

`func (o *ForwardOnlyPaginationOutput) GetTotalEntriesOk() (*float32, bool)`

GetTotalEntriesOk returns a tuple with the TotalEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalEntries

`func (o *ForwardOnlyPaginationOutput) SetTotalEntries(v float32)`

SetTotalEntries sets TotalEntries field to given value.

### HasTotalEntries

`func (o *ForwardOnlyPaginationOutput) HasTotalEntries() bool`

HasTotalEntries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


