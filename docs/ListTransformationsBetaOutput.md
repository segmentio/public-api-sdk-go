# ListTransformationsBetaOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Transformations** | [**[]TransformationBeta**](TransformationBeta.md) | A paginated list of Transformations. | 
**Pagination** | [**PaginationOutput**](PaginationOutput.md) |  | 

## Methods

### NewListTransformationsBetaOutput

`func NewListTransformationsBetaOutput(transformations []TransformationBeta, pagination PaginationOutput, ) *ListTransformationsBetaOutput`

NewListTransformationsBetaOutput instantiates a new ListTransformationsBetaOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTransformationsBetaOutputWithDefaults

`func NewListTransformationsBetaOutputWithDefaults() *ListTransformationsBetaOutput`

NewListTransformationsBetaOutputWithDefaults instantiates a new ListTransformationsBetaOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransformations

`func (o *ListTransformationsBetaOutput) GetTransformations() []TransformationBeta`

GetTransformations returns the Transformations field if non-nil, zero value otherwise.

### GetTransformationsOk

`func (o *ListTransformationsBetaOutput) GetTransformationsOk() (*[]TransformationBeta, bool)`

GetTransformationsOk returns a tuple with the Transformations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransformations

`func (o *ListTransformationsBetaOutput) SetTransformations(v []TransformationBeta)`

SetTransformations sets Transformations field to given value.


### GetPagination

`func (o *ListTransformationsBetaOutput) GetPagination() PaginationOutput`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListTransformationsBetaOutput) GetPaginationOk() (*PaginationOutput, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListTransformationsBetaOutput) SetPagination(v PaginationOutput)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


