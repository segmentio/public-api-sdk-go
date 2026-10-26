# ListSelectiveSyncsFromWarehouseAndSpaceAlphaOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]SpaceWarehouseSelectiveSyncItemAlpha**](SpaceWarehouseSelectiveSyncItemAlpha.md) | Represents a list of collections and properties synced to the Warehouse. | 
**EnableEventTables** | Pointer to **bool** | Optional. Represents the enabled state of all event tables. | [optional] 
**Pagination** | [**PaginationOutput**](PaginationOutput.md) |  | 

## Methods

### NewListSelectiveSyncsFromWarehouseAndSpaceAlphaOutput

`func NewListSelectiveSyncsFromWarehouseAndSpaceAlphaOutput(items []SpaceWarehouseSelectiveSyncItemAlpha, pagination PaginationOutput, ) *ListSelectiveSyncsFromWarehouseAndSpaceAlphaOutput`

NewListSelectiveSyncsFromWarehouseAndSpaceAlphaOutput instantiates a new ListSelectiveSyncsFromWarehouseAndSpaceAlphaOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSelectiveSyncsFromWarehouseAndSpaceAlphaOutputWithDefaults

`func NewListSelectiveSyncsFromWarehouseAndSpaceAlphaOutputWithDefaults() *ListSelectiveSyncsFromWarehouseAndSpaceAlphaOutput`

NewListSelectiveSyncsFromWarehouseAndSpaceAlphaOutputWithDefaults instantiates a new ListSelectiveSyncsFromWarehouseAndSpaceAlphaOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *ListSelectiveSyncsFromWarehouseAndSpaceAlphaOutput) GetItems() []SpaceWarehouseSelectiveSyncItemAlpha`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ListSelectiveSyncsFromWarehouseAndSpaceAlphaOutput) GetItemsOk() (*[]SpaceWarehouseSelectiveSyncItemAlpha, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ListSelectiveSyncsFromWarehouseAndSpaceAlphaOutput) SetItems(v []SpaceWarehouseSelectiveSyncItemAlpha)`

SetItems sets Items field to given value.


### GetEnableEventTables

`func (o *ListSelectiveSyncsFromWarehouseAndSpaceAlphaOutput) GetEnableEventTables() bool`

GetEnableEventTables returns the EnableEventTables field if non-nil, zero value otherwise.

### GetEnableEventTablesOk

`func (o *ListSelectiveSyncsFromWarehouseAndSpaceAlphaOutput) GetEnableEventTablesOk() (*bool, bool)`

GetEnableEventTablesOk returns a tuple with the EnableEventTables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableEventTables

`func (o *ListSelectiveSyncsFromWarehouseAndSpaceAlphaOutput) SetEnableEventTables(v bool)`

SetEnableEventTables sets EnableEventTables field to given value.

### HasEnableEventTables

`func (o *ListSelectiveSyncsFromWarehouseAndSpaceAlphaOutput) HasEnableEventTables() bool`

HasEnableEventTables returns a boolean if a field has been set.

### GetPagination

`func (o *ListSelectiveSyncsFromWarehouseAndSpaceAlphaOutput) GetPagination() PaginationOutput`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListSelectiveSyncsFromWarehouseAndSpaceAlphaOutput) GetPaginationOk() (*PaginationOutput, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListSelectiveSyncsFromWarehouseAndSpaceAlphaOutput) SetPagination(v PaginationOutput)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


