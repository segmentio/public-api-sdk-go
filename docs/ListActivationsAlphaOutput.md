# ListActivationsAlphaOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Activations** | [**[]ActivationOutput**](ActivationOutput.md) | A list of activation results. | 
**Pagination** | Pointer to [**PaginationOutput**](PaginationOutput.md) |  | [optional] 

## Methods

### NewListActivationsAlphaOutput

`func NewListActivationsAlphaOutput(activations []ActivationOutput, ) *ListActivationsAlphaOutput`

NewListActivationsAlphaOutput instantiates a new ListActivationsAlphaOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListActivationsAlphaOutputWithDefaults

`func NewListActivationsAlphaOutputWithDefaults() *ListActivationsAlphaOutput`

NewListActivationsAlphaOutputWithDefaults instantiates a new ListActivationsAlphaOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivations

`func (o *ListActivationsAlphaOutput) GetActivations() []ActivationOutput`

GetActivations returns the Activations field if non-nil, zero value otherwise.

### GetActivationsOk

`func (o *ListActivationsAlphaOutput) GetActivationsOk() (*[]ActivationOutput, bool)`

GetActivationsOk returns a tuple with the Activations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivations

`func (o *ListActivationsAlphaOutput) SetActivations(v []ActivationOutput)`

SetActivations sets Activations field to given value.


### GetPagination

`func (o *ListActivationsAlphaOutput) GetPagination() PaginationOutput`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListActivationsAlphaOutput) GetPaginationOk() (*PaginationOutput, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListActivationsAlphaOutput) SetPagination(v PaginationOutput)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *ListActivationsAlphaOutput) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


