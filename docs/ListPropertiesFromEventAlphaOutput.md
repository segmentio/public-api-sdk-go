# ListPropertiesFromEventAlphaOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventProperties** | [**[]EventPropertyAlpha**](EventPropertyAlpha.md) | List of properties for the event. | 
**Pagination** | [**PaginationOutput**](PaginationOutput.md) |  | 

## Methods

### NewListPropertiesFromEventAlphaOutput

`func NewListPropertiesFromEventAlphaOutput(eventProperties []EventPropertyAlpha, pagination PaginationOutput, ) *ListPropertiesFromEventAlphaOutput`

NewListPropertiesFromEventAlphaOutput instantiates a new ListPropertiesFromEventAlphaOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListPropertiesFromEventAlphaOutputWithDefaults

`func NewListPropertiesFromEventAlphaOutputWithDefaults() *ListPropertiesFromEventAlphaOutput`

NewListPropertiesFromEventAlphaOutputWithDefaults instantiates a new ListPropertiesFromEventAlphaOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventProperties

`func (o *ListPropertiesFromEventAlphaOutput) GetEventProperties() []EventPropertyAlpha`

GetEventProperties returns the EventProperties field if non-nil, zero value otherwise.

### GetEventPropertiesOk

`func (o *ListPropertiesFromEventAlphaOutput) GetEventPropertiesOk() (*[]EventPropertyAlpha, bool)`

GetEventPropertiesOk returns a tuple with the EventProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventProperties

`func (o *ListPropertiesFromEventAlphaOutput) SetEventProperties(v []EventPropertyAlpha)`

SetEventProperties sets EventProperties field to given value.


### GetPagination

`func (o *ListPropertiesFromEventAlphaOutput) GetPagination() PaginationOutput`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListPropertiesFromEventAlphaOutput) GetPaginationOk() (*PaginationOutput, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListPropertiesFromEventAlphaOutput) SetPagination(v PaginationOutput)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


