# UpdateSelectiveSyncForWarehouseAndSpaceAlphaInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SyncOverrides** | Pointer to [**[]SpaceWarehouseSchemaOverride**](SpaceWarehouseSchemaOverride.md) | A list of sync Schema overrides to apply to this Space Warehouse. Note: Selective Sync is not supported if the enableEventTables flag is false. | [optional] 
**EnableEventTables** | Pointer to **bool** | A flag to enable or disable all event Tables. This field is optional. | [optional] 

## Methods

### NewUpdateSelectiveSyncForWarehouseAndSpaceAlphaInput

`func NewUpdateSelectiveSyncForWarehouseAndSpaceAlphaInput() *UpdateSelectiveSyncForWarehouseAndSpaceAlphaInput`

NewUpdateSelectiveSyncForWarehouseAndSpaceAlphaInput instantiates a new UpdateSelectiveSyncForWarehouseAndSpaceAlphaInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSelectiveSyncForWarehouseAndSpaceAlphaInputWithDefaults

`func NewUpdateSelectiveSyncForWarehouseAndSpaceAlphaInputWithDefaults() *UpdateSelectiveSyncForWarehouseAndSpaceAlphaInput`

NewUpdateSelectiveSyncForWarehouseAndSpaceAlphaInputWithDefaults instantiates a new UpdateSelectiveSyncForWarehouseAndSpaceAlphaInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSyncOverrides

`func (o *UpdateSelectiveSyncForWarehouseAndSpaceAlphaInput) GetSyncOverrides() []SpaceWarehouseSchemaOverride`

GetSyncOverrides returns the SyncOverrides field if non-nil, zero value otherwise.

### GetSyncOverridesOk

`func (o *UpdateSelectiveSyncForWarehouseAndSpaceAlphaInput) GetSyncOverridesOk() (*[]SpaceWarehouseSchemaOverride, bool)`

GetSyncOverridesOk returns a tuple with the SyncOverrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncOverrides

`func (o *UpdateSelectiveSyncForWarehouseAndSpaceAlphaInput) SetSyncOverrides(v []SpaceWarehouseSchemaOverride)`

SetSyncOverrides sets SyncOverrides field to given value.

### HasSyncOverrides

`func (o *UpdateSelectiveSyncForWarehouseAndSpaceAlphaInput) HasSyncOverrides() bool`

HasSyncOverrides returns a boolean if a field has been set.

### GetEnableEventTables

`func (o *UpdateSelectiveSyncForWarehouseAndSpaceAlphaInput) GetEnableEventTables() bool`

GetEnableEventTables returns the EnableEventTables field if non-nil, zero value otherwise.

### GetEnableEventTablesOk

`func (o *UpdateSelectiveSyncForWarehouseAndSpaceAlphaInput) GetEnableEventTablesOk() (*bool, bool)`

GetEnableEventTablesOk returns a tuple with the EnableEventTables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableEventTables

`func (o *UpdateSelectiveSyncForWarehouseAndSpaceAlphaInput) SetEnableEventTables(v bool)`

SetEnableEventTables sets EnableEventTables field to given value.

### HasEnableEventTables

`func (o *UpdateSelectiveSyncForWarehouseAndSpaceAlphaInput) HasEnableEventTables() bool`

HasEnableEventTables returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


