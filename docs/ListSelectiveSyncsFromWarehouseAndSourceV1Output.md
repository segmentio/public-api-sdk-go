# ListSelectiveSyncsFromWarehouseAndSourceV1Output

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]WarehouseSelectiveSyncItemV1**](WarehouseSelectiveSyncItemV1.md) | Represents a list of Source, collection, and properties synced to the Warehouse. | 
**Pagination** | [**PaginationOutput**](PaginationOutput.md) |  | 

## Methods

### NewListSelectiveSyncsFromWarehouseAndSourceV1Output

`func NewListSelectiveSyncsFromWarehouseAndSourceV1Output(items []WarehouseSelectiveSyncItemV1, pagination PaginationOutput, ) *ListSelectiveSyncsFromWarehouseAndSourceV1Output`

NewListSelectiveSyncsFromWarehouseAndSourceV1Output instantiates a new ListSelectiveSyncsFromWarehouseAndSourceV1Output object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSelectiveSyncsFromWarehouseAndSourceV1OutputWithDefaults

`func NewListSelectiveSyncsFromWarehouseAndSourceV1OutputWithDefaults() *ListSelectiveSyncsFromWarehouseAndSourceV1Output`

NewListSelectiveSyncsFromWarehouseAndSourceV1OutputWithDefaults instantiates a new ListSelectiveSyncsFromWarehouseAndSourceV1Output object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *ListSelectiveSyncsFromWarehouseAndSourceV1Output) GetItems() []WarehouseSelectiveSyncItemV1`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ListSelectiveSyncsFromWarehouseAndSourceV1Output) GetItemsOk() (*[]WarehouseSelectiveSyncItemV1, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ListSelectiveSyncsFromWarehouseAndSourceV1Output) SetItems(v []WarehouseSelectiveSyncItemV1)`

SetItems sets Items field to given value.


### GetPagination

`func (o *ListSelectiveSyncsFromWarehouseAndSourceV1Output) GetPagination() PaginationOutput`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListSelectiveSyncsFromWarehouseAndSourceV1Output) GetPaginationOk() (*PaginationOutput, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListSelectiveSyncsFromWarehouseAndSourceV1Output) SetPagination(v PaginationOutput)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


