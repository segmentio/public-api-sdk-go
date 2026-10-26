# WarehouseV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The id of the Warehouse. | 
**Metadata** | [**WarehouseMetadataV1**](WarehouseMetadataV1.md) |  | 
**WorkspaceId** | **string** | The id of the Workspace that owns this Warehouse. | 
**Enabled** | **bool** | When set to true, this Warehouse receives data. | 
**Settings** | **map[string]interface{}** | A key-value object that contains instance-specific Warehouse settings. | 

## Methods

### NewWarehouseV1

`func NewWarehouseV1(id string, metadata WarehouseMetadataV1, workspaceId string, enabled bool, settings map[string]interface{}, ) *WarehouseV1`

NewWarehouseV1 instantiates a new WarehouseV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWarehouseV1WithDefaults

`func NewWarehouseV1WithDefaults() *WarehouseV1`

NewWarehouseV1WithDefaults instantiates a new WarehouseV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WarehouseV1) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WarehouseV1) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WarehouseV1) SetId(v string)`

SetId sets Id field to given value.


### GetMetadata

`func (o *WarehouseV1) GetMetadata() WarehouseMetadataV1`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *WarehouseV1) GetMetadataOk() (*WarehouseMetadataV1, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *WarehouseV1) SetMetadata(v WarehouseMetadataV1)`

SetMetadata sets Metadata field to given value.


### GetWorkspaceId

`func (o *WarehouseV1) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *WarehouseV1) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *WarehouseV1) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.


### GetEnabled

`func (o *WarehouseV1) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *WarehouseV1) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *WarehouseV1) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetSettings

`func (o *WarehouseV1) GetSettings() map[string]interface{}`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *WarehouseV1) GetSettingsOk() (*map[string]interface{}, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *WarehouseV1) SetSettings(v map[string]interface{})`

SetSettings sets Settings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


