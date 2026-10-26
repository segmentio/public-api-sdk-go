# ListEntityPathsAlphaOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityPaths** | [**[]EntityPathAlpha**](EntityPathAlpha.md) | List of entity paths in the space. | 
**Pagination** | [**PaginationOutput**](PaginationOutput.md) |  | 

## Methods

### NewListEntityPathsAlphaOutput

`func NewListEntityPathsAlphaOutput(entityPaths []EntityPathAlpha, pagination PaginationOutput, ) *ListEntityPathsAlphaOutput`

NewListEntityPathsAlphaOutput instantiates a new ListEntityPathsAlphaOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListEntityPathsAlphaOutputWithDefaults

`func NewListEntityPathsAlphaOutputWithDefaults() *ListEntityPathsAlphaOutput`

NewListEntityPathsAlphaOutputWithDefaults instantiates a new ListEntityPathsAlphaOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityPaths

`func (o *ListEntityPathsAlphaOutput) GetEntityPaths() []EntityPathAlpha`

GetEntityPaths returns the EntityPaths field if non-nil, zero value otherwise.

### GetEntityPathsOk

`func (o *ListEntityPathsAlphaOutput) GetEntityPathsOk() (*[]EntityPathAlpha, bool)`

GetEntityPathsOk returns a tuple with the EntityPaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityPaths

`func (o *ListEntityPathsAlphaOutput) SetEntityPaths(v []EntityPathAlpha)`

SetEntityPaths sets EntityPaths field to given value.


### GetPagination

`func (o *ListEntityPathsAlphaOutput) GetPagination() PaginationOutput`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListEntityPathsAlphaOutput) GetPaginationOk() (*PaginationOutput, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListEntityPathsAlphaOutput) SetPagination(v PaginationOutput)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


