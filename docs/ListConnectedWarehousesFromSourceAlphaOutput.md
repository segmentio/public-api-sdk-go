# ListConnectedWarehousesFromSourceAlphaOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Warehouses** | [**[]WarehouseV1**](WarehouseV1.md) | A list that contains the Warehouses connected to the Source. | 
**Pagination** | [**PaginationOutput**](PaginationOutput.md) |  | 

## Methods

### NewListConnectedWarehousesFromSourceAlphaOutput

`func NewListConnectedWarehousesFromSourceAlphaOutput(warehouses []WarehouseV1, pagination PaginationOutput, ) *ListConnectedWarehousesFromSourceAlphaOutput`

NewListConnectedWarehousesFromSourceAlphaOutput instantiates a new ListConnectedWarehousesFromSourceAlphaOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListConnectedWarehousesFromSourceAlphaOutputWithDefaults

`func NewListConnectedWarehousesFromSourceAlphaOutputWithDefaults() *ListConnectedWarehousesFromSourceAlphaOutput`

NewListConnectedWarehousesFromSourceAlphaOutputWithDefaults instantiates a new ListConnectedWarehousesFromSourceAlphaOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWarehouses

`func (o *ListConnectedWarehousesFromSourceAlphaOutput) GetWarehouses() []WarehouseV1`

GetWarehouses returns the Warehouses field if non-nil, zero value otherwise.

### GetWarehousesOk

`func (o *ListConnectedWarehousesFromSourceAlphaOutput) GetWarehousesOk() (*[]WarehouseV1, bool)`

GetWarehousesOk returns a tuple with the Warehouses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouses

`func (o *ListConnectedWarehousesFromSourceAlphaOutput) SetWarehouses(v []WarehouseV1)`

SetWarehouses sets Warehouses field to given value.


### GetPagination

`func (o *ListConnectedWarehousesFromSourceAlphaOutput) GetPagination() PaginationOutput`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListConnectedWarehousesFromSourceAlphaOutput) GetPaginationOk() (*PaginationOutput, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListConnectedWarehousesFromSourceAlphaOutput) SetPagination(v PaginationOutput)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


