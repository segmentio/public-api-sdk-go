# ListFiltersPaginationInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cursor** | Pointer to **float32** | The offset. | [optional] 
**Count** | **float32** | The number of items to retrieve in a page, between 1 and 10. | 

## Methods

### NewListFiltersPaginationInput

`func NewListFiltersPaginationInput(count float32, ) *ListFiltersPaginationInput`

NewListFiltersPaginationInput instantiates a new ListFiltersPaginationInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListFiltersPaginationInputWithDefaults

`func NewListFiltersPaginationInputWithDefaults() *ListFiltersPaginationInput`

NewListFiltersPaginationInputWithDefaults instantiates a new ListFiltersPaginationInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCursor

`func (o *ListFiltersPaginationInput) GetCursor() float32`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *ListFiltersPaginationInput) GetCursorOk() (*float32, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *ListFiltersPaginationInput) SetCursor(v float32)`

SetCursor sets Cursor field to given value.

### HasCursor

`func (o *ListFiltersPaginationInput) HasCursor() bool`

HasCursor returns a boolean if a field has been set.

### GetCount

`func (o *ListFiltersPaginationInput) GetCount() float32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ListFiltersPaginationInput) GetCountOk() (*float32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ListFiltersPaginationInput) SetCount(v float32)`

SetCount sets Count field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


