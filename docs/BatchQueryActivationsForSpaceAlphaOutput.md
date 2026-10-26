# BatchQueryActivationsForSpaceAlphaOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Activations** | [**[]ActivationOutput**](ActivationOutput.md) | Activation results across all valid requested audiences. | 
**Errors** | [**[]ErrorsInner**](ErrorsInner.md) | Audience IDs from the request that could not be resolved, with why - for example not found, or wrong space. When &#x60;audienceIds&#x60; was omitted from the request (space-scan mode), this never contains a caller-supplied bad id - there isn&#39;t one to be wrong - but it can still contain per-activation hydration failures, same as the &#x60;audienceIds&#x60;-provided path. | 
**Pagination** | Pointer to [**PaginationOutput**](PaginationOutput.md) |  | [optional] 

## Methods

### NewBatchQueryActivationsForSpaceAlphaOutput

`func NewBatchQueryActivationsForSpaceAlphaOutput(activations []ActivationOutput, errors []ErrorsInner, ) *BatchQueryActivationsForSpaceAlphaOutput`

NewBatchQueryActivationsForSpaceAlphaOutput instantiates a new BatchQueryActivationsForSpaceAlphaOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchQueryActivationsForSpaceAlphaOutputWithDefaults

`func NewBatchQueryActivationsForSpaceAlphaOutputWithDefaults() *BatchQueryActivationsForSpaceAlphaOutput`

NewBatchQueryActivationsForSpaceAlphaOutputWithDefaults instantiates a new BatchQueryActivationsForSpaceAlphaOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivations

`func (o *BatchQueryActivationsForSpaceAlphaOutput) GetActivations() []ActivationOutput`

GetActivations returns the Activations field if non-nil, zero value otherwise.

### GetActivationsOk

`func (o *BatchQueryActivationsForSpaceAlphaOutput) GetActivationsOk() (*[]ActivationOutput, bool)`

GetActivationsOk returns a tuple with the Activations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivations

`func (o *BatchQueryActivationsForSpaceAlphaOutput) SetActivations(v []ActivationOutput)`

SetActivations sets Activations field to given value.


### GetErrors

`func (o *BatchQueryActivationsForSpaceAlphaOutput) GetErrors() []ErrorsInner`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *BatchQueryActivationsForSpaceAlphaOutput) GetErrorsOk() (*[]ErrorsInner, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *BatchQueryActivationsForSpaceAlphaOutput) SetErrors(v []ErrorsInner)`

SetErrors sets Errors field to given value.


### GetPagination

`func (o *BatchQueryActivationsForSpaceAlphaOutput) GetPagination() PaginationOutput`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *BatchQueryActivationsForSpaceAlphaOutput) GetPaginationOk() (*PaginationOutput, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *BatchQueryActivationsForSpaceAlphaOutput) SetPagination(v PaginationOutput)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *BatchQueryActivationsForSpaceAlphaOutput) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


