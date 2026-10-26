# GetWarehousesCatalogV1Output

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WarehousesCatalog** | [**[]WarehouseMetadataV1**](WarehouseMetadataV1.md) | All Warehouse catalog items contained within the requested page. | 
**Pagination** | [**PaginationOutput**](PaginationOutput.md) |  | 

## Methods

### NewGetWarehousesCatalogV1Output

`func NewGetWarehousesCatalogV1Output(warehousesCatalog []WarehouseMetadataV1, pagination PaginationOutput, ) *GetWarehousesCatalogV1Output`

NewGetWarehousesCatalogV1Output instantiates a new GetWarehousesCatalogV1Output object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetWarehousesCatalogV1OutputWithDefaults

`func NewGetWarehousesCatalogV1OutputWithDefaults() *GetWarehousesCatalogV1Output`

NewGetWarehousesCatalogV1OutputWithDefaults instantiates a new GetWarehousesCatalogV1Output object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWarehousesCatalog

`func (o *GetWarehousesCatalogV1Output) GetWarehousesCatalog() []WarehouseMetadataV1`

GetWarehousesCatalog returns the WarehousesCatalog field if non-nil, zero value otherwise.

### GetWarehousesCatalogOk

`func (o *GetWarehousesCatalogV1Output) GetWarehousesCatalogOk() (*[]WarehouseMetadataV1, bool)`

GetWarehousesCatalogOk returns a tuple with the WarehousesCatalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehousesCatalog

`func (o *GetWarehousesCatalogV1Output) SetWarehousesCatalog(v []WarehouseMetadataV1)`

SetWarehousesCatalog sets WarehousesCatalog field to given value.


### GetPagination

`func (o *GetWarehousesCatalogV1Output) GetPagination() PaginationOutput`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *GetWarehousesCatalogV1Output) GetPaginationOk() (*PaginationOutput, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *GetWarehousesCatalogV1Output) SetPagination(v PaginationOutput)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


