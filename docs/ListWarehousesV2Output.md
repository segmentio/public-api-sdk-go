# ListWarehousesV2Output

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Warehouses** | [**[]WarehouseV2**](WarehouseV2.md) | A list of Warehouses that belong to the Workspace. | 
**Pagination** | [**PaginationOutput**](PaginationOutput.md) |  | 

## Methods

### NewListWarehousesV2Output

`func NewListWarehousesV2Output(warehouses []WarehouseV2, pagination PaginationOutput, ) *ListWarehousesV2Output`

NewListWarehousesV2Output instantiates a new ListWarehousesV2Output object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListWarehousesV2OutputWithDefaults

`func NewListWarehousesV2OutputWithDefaults() *ListWarehousesV2Output`

NewListWarehousesV2OutputWithDefaults instantiates a new ListWarehousesV2Output object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWarehouses

`func (o *ListWarehousesV2Output) GetWarehouses() []WarehouseV2`

GetWarehouses returns the Warehouses field if non-nil, zero value otherwise.

### GetWarehousesOk

`func (o *ListWarehousesV2Output) GetWarehousesOk() (*[]WarehouseV2, bool)`

GetWarehousesOk returns a tuple with the Warehouses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouses

`func (o *ListWarehousesV2Output) SetWarehouses(v []WarehouseV2)`

SetWarehouses sets Warehouses field to given value.


### GetPagination

`func (o *ListWarehousesV2Output) GetPagination() PaginationOutput`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListWarehousesV2Output) GetPaginationOk() (*PaginationOutput, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListWarehousesV2Output) SetPagination(v PaginationOutput)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


