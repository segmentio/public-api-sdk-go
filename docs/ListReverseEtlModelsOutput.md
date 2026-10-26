# ListReverseEtlModelsOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Models** | [**[]ReverseEtlModel**](ReverseEtlModel.md) | A list of Models that belong to the Workspace. | 
**Pagination** | [**PaginationOutput**](PaginationOutput.md) |  | 

## Methods

### NewListReverseEtlModelsOutput

`func NewListReverseEtlModelsOutput(models []ReverseEtlModel, pagination PaginationOutput, ) *ListReverseEtlModelsOutput`

NewListReverseEtlModelsOutput instantiates a new ListReverseEtlModelsOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListReverseEtlModelsOutputWithDefaults

`func NewListReverseEtlModelsOutputWithDefaults() *ListReverseEtlModelsOutput`

NewListReverseEtlModelsOutputWithDefaults instantiates a new ListReverseEtlModelsOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModels

`func (o *ListReverseEtlModelsOutput) GetModels() []ReverseEtlModel`

GetModels returns the Models field if non-nil, zero value otherwise.

### GetModelsOk

`func (o *ListReverseEtlModelsOutput) GetModelsOk() (*[]ReverseEtlModel, bool)`

GetModelsOk returns a tuple with the Models field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModels

`func (o *ListReverseEtlModelsOutput) SetModels(v []ReverseEtlModel)`

SetModels sets Models field to given value.


### GetPagination

`func (o *ListReverseEtlModelsOutput) GetPagination() PaginationOutput`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListReverseEtlModelsOutput) GetPaginationOk() (*PaginationOutput, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListReverseEtlModelsOutput) SetPagination(v PaginationOutput)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


