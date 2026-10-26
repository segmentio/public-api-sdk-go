# ListSourcesAlphaOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sources** | [**[]SourceAlpha**](SourceAlpha.md) | A list of Sources that belong to the Workspace. | 
**Pagination** | [**PaginationOutput**](PaginationOutput.md) |  | 

## Methods

### NewListSourcesAlphaOutput

`func NewListSourcesAlphaOutput(sources []SourceAlpha, pagination PaginationOutput, ) *ListSourcesAlphaOutput`

NewListSourcesAlphaOutput instantiates a new ListSourcesAlphaOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSourcesAlphaOutputWithDefaults

`func NewListSourcesAlphaOutputWithDefaults() *ListSourcesAlphaOutput`

NewListSourcesAlphaOutputWithDefaults instantiates a new ListSourcesAlphaOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSources

`func (o *ListSourcesAlphaOutput) GetSources() []SourceAlpha`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *ListSourcesAlphaOutput) GetSourcesOk() (*[]SourceAlpha, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *ListSourcesAlphaOutput) SetSources(v []SourceAlpha)`

SetSources sets Sources field to given value.


### GetPagination

`func (o *ListSourcesAlphaOutput) GetPagination() PaginationOutput`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListSourcesAlphaOutput) GetPaginationOk() (*PaginationOutput, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListSourcesAlphaOutput) SetPagination(v PaginationOutput)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


