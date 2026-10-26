# ListFunctionVersionsAlphaOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Versions** | [**[]Version**](Version.md) | An array of Functions. | 
**Pagination** | [**PaginationOutput**](PaginationOutput.md) |  | 

## Methods

### NewListFunctionVersionsAlphaOutput

`func NewListFunctionVersionsAlphaOutput(versions []Version, pagination PaginationOutput, ) *ListFunctionVersionsAlphaOutput`

NewListFunctionVersionsAlphaOutput instantiates a new ListFunctionVersionsAlphaOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListFunctionVersionsAlphaOutputWithDefaults

`func NewListFunctionVersionsAlphaOutputWithDefaults() *ListFunctionVersionsAlphaOutput`

NewListFunctionVersionsAlphaOutputWithDefaults instantiates a new ListFunctionVersionsAlphaOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersions

`func (o *ListFunctionVersionsAlphaOutput) GetVersions() []Version`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *ListFunctionVersionsAlphaOutput) GetVersionsOk() (*[]Version, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *ListFunctionVersionsAlphaOutput) SetVersions(v []Version)`

SetVersions sets Versions field to given value.


### GetPagination

`func (o *ListFunctionVersionsAlphaOutput) GetPagination() PaginationOutput`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListFunctionVersionsAlphaOutput) GetPaginationOk() (*PaginationOutput, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListFunctionVersionsAlphaOutput) SetPagination(v PaginationOutput)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


