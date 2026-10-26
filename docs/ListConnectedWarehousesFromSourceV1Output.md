# ListConnectedWarehousesFromSourceV1Output

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Warehouses** | [**[]WarehouseV1**](WarehouseV1.md) | A list that contains the Warehouses connected to the Source. | 
**Pagination** | [**PaginationOutput**](PaginationOutput.md) |  | 

## Methods

### NewListConnectedWarehousesFromSourceV1Output

`func NewListConnectedWarehousesFromSourceV1Output(warehouses []WarehouseV1, pagination PaginationOutput, ) *ListConnectedWarehousesFromSourceV1Output`

NewListConnectedWarehousesFromSourceV1Output instantiates a new ListConnectedWarehousesFromSourceV1Output object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListConnectedWarehousesFromSourceV1OutputWithDefaults

`func NewListConnectedWarehousesFromSourceV1OutputWithDefaults() *ListConnectedWarehousesFromSourceV1Output`

NewListConnectedWarehousesFromSourceV1OutputWithDefaults instantiates a new ListConnectedWarehousesFromSourceV1Output object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWarehouses

`func (o *ListConnectedWarehousesFromSourceV1Output) GetWarehouses() []WarehouseV1`

GetWarehouses returns the Warehouses field if non-nil, zero value otherwise.

### GetWarehousesOk

`func (o *ListConnectedWarehousesFromSourceV1Output) GetWarehousesOk() (*[]WarehouseV1, bool)`

GetWarehousesOk returns a tuple with the Warehouses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouses

`func (o *ListConnectedWarehousesFromSourceV1Output) SetWarehouses(v []WarehouseV1)`

SetWarehouses sets Warehouses field to given value.


### GetPagination

`func (o *ListConnectedWarehousesFromSourceV1Output) GetPagination() PaginationOutput`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListConnectedWarehousesFromSourceV1Output) GetPaginationOk() (*PaginationOutput, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListConnectedWarehousesFromSourceV1Output) SetPagination(v PaginationOutput)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


