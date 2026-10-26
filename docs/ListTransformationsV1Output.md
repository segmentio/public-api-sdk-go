# ListTransformationsV1Output

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Transformations** | [**[]TransformationV1**](TransformationV1.md) | A paginated list of Transformations. | 
**Pagination** | [**PaginationOutput**](PaginationOutput.md) |  | 

## Methods

### NewListTransformationsV1Output

`func NewListTransformationsV1Output(transformations []TransformationV1, pagination PaginationOutput, ) *ListTransformationsV1Output`

NewListTransformationsV1Output instantiates a new ListTransformationsV1Output object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTransformationsV1OutputWithDefaults

`func NewListTransformationsV1OutputWithDefaults() *ListTransformationsV1Output`

NewListTransformationsV1OutputWithDefaults instantiates a new ListTransformationsV1Output object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransformations

`func (o *ListTransformationsV1Output) GetTransformations() []TransformationV1`

GetTransformations returns the Transformations field if non-nil, zero value otherwise.

### GetTransformationsOk

`func (o *ListTransformationsV1Output) GetTransformationsOk() (*[]TransformationV1, bool)`

GetTransformationsOk returns a tuple with the Transformations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransformations

`func (o *ListTransformationsV1Output) SetTransformations(v []TransformationV1)`

SetTransformations sets Transformations field to given value.


### GetPagination

`func (o *ListTransformationsV1Output) GetPagination() PaginationOutput`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListTransformationsV1Output) GetPaginationOk() (*PaginationOutput, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListTransformationsV1Output) SetPagination(v PaginationOutput)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


