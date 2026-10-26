# PaginationInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cursor** | Pointer to **string** | The page to request.  Acceptable values to use here are in PaginationOutput objects, in the &#x60;current&#x60;, &#x60;next&#x60;, and &#x60;previous&#x60; keys.  Consumers of the API must treat this value as opaque. | [optional] 
**Count** | **int32** | The number of items to retrieve in a page, between 1 and 1000. | 

## Methods

### NewPaginationInput

`func NewPaginationInput(count int32, ) *PaginationInput`

NewPaginationInput instantiates a new PaginationInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginationInputWithDefaults

`func NewPaginationInputWithDefaults() *PaginationInput`

NewPaginationInputWithDefaults instantiates a new PaginationInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCursor

`func (o *PaginationInput) GetCursor() string`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *PaginationInput) GetCursorOk() (*string, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *PaginationInput) SetCursor(v string)`

SetCursor sets Cursor field to given value.

### HasCursor

`func (o *PaginationInput) HasCursor() bool`

HasCursor returns a boolean if a field has been set.

### GetCount

`func (o *PaginationInput) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PaginationInput) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PaginationInput) SetCount(v int32)`

SetCount sets Count field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


