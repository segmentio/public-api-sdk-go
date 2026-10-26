# ListConnectedSourcesFromWarehouseV1Output

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sources** | [**[]SourceV1**](SourceV1.md) | A list that contains the Sources connected to the requested Warehouse. | 
**Pagination** | [**PaginationOutput**](PaginationOutput.md) |  | 

## Methods

### NewListConnectedSourcesFromWarehouseV1Output

`func NewListConnectedSourcesFromWarehouseV1Output(sources []SourceV1, pagination PaginationOutput, ) *ListConnectedSourcesFromWarehouseV1Output`

NewListConnectedSourcesFromWarehouseV1Output instantiates a new ListConnectedSourcesFromWarehouseV1Output object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListConnectedSourcesFromWarehouseV1OutputWithDefaults

`func NewListConnectedSourcesFromWarehouseV1OutputWithDefaults() *ListConnectedSourcesFromWarehouseV1Output`

NewListConnectedSourcesFromWarehouseV1OutputWithDefaults instantiates a new ListConnectedSourcesFromWarehouseV1Output object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSources

`func (o *ListConnectedSourcesFromWarehouseV1Output) GetSources() []SourceV1`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *ListConnectedSourcesFromWarehouseV1Output) GetSourcesOk() (*[]SourceV1, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *ListConnectedSourcesFromWarehouseV1Output) SetSources(v []SourceV1)`

SetSources sets Sources field to given value.


### GetPagination

`func (o *ListConnectedSourcesFromWarehouseV1Output) GetPagination() PaginationOutput`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListConnectedSourcesFromWarehouseV1Output) GetPaginationOk() (*PaginationOutput, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListConnectedSourcesFromWarehouseV1Output) SetPagination(v PaginationOutput)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


