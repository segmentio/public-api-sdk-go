# ListInsertFunctionInstancesAlphaOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Instances** | [**[]InsertFunctionInstanceAlpha**](InsertFunctionInstanceAlpha.md) | All insert Function instances found. | 
**Pagination** | [**PaginationOutput**](PaginationOutput.md) |  | 

## Methods

### NewListInsertFunctionInstancesAlphaOutput

`func NewListInsertFunctionInstancesAlphaOutput(instances []InsertFunctionInstanceAlpha, pagination PaginationOutput, ) *ListInsertFunctionInstancesAlphaOutput`

NewListInsertFunctionInstancesAlphaOutput instantiates a new ListInsertFunctionInstancesAlphaOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListInsertFunctionInstancesAlphaOutputWithDefaults

`func NewListInsertFunctionInstancesAlphaOutputWithDefaults() *ListInsertFunctionInstancesAlphaOutput`

NewListInsertFunctionInstancesAlphaOutputWithDefaults instantiates a new ListInsertFunctionInstancesAlphaOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstances

`func (o *ListInsertFunctionInstancesAlphaOutput) GetInstances() []InsertFunctionInstanceAlpha`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *ListInsertFunctionInstancesAlphaOutput) GetInstancesOk() (*[]InsertFunctionInstanceAlpha, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *ListInsertFunctionInstancesAlphaOutput) SetInstances(v []InsertFunctionInstanceAlpha)`

SetInstances sets Instances field to given value.


### GetPagination

`func (o *ListInsertFunctionInstancesAlphaOutput) GetPagination() PaginationOutput`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListInsertFunctionInstancesAlphaOutput) GetPaginationOk() (*PaginationOutput, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListInsertFunctionInstancesAlphaOutput) SetPagination(v PaginationOutput)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


