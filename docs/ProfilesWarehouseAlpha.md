# ProfilesWarehouseAlpha

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The id of the Warehouse. | 
**SpaceId** | **string** | The Space id. | 
**Metadata** | [**WarehouseMetadataV1**](WarehouseMetadataV1.md) |  | 
**WorkspaceId** | **string** | The id of the Workspace that owns this Warehouse. | 
**Enabled** | **bool** | When set to true, this Warehouse receives data. | 
**Settings** | **map[string]interface{}** | A key-value object that contains instance-specific Warehouse settings. | 
**SchemaName** | Pointer to **string** | The custom schema name that Segment uses on the Warehouse side. | [optional] 

## Methods

### NewProfilesWarehouseAlpha

`func NewProfilesWarehouseAlpha(id string, spaceId string, metadata WarehouseMetadataV1, workspaceId string, enabled bool, settings map[string]interface{}, ) *ProfilesWarehouseAlpha`

NewProfilesWarehouseAlpha instantiates a new ProfilesWarehouseAlpha object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProfilesWarehouseAlphaWithDefaults

`func NewProfilesWarehouseAlphaWithDefaults() *ProfilesWarehouseAlpha`

NewProfilesWarehouseAlphaWithDefaults instantiates a new ProfilesWarehouseAlpha object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProfilesWarehouseAlpha) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProfilesWarehouseAlpha) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProfilesWarehouseAlpha) SetId(v string)`

SetId sets Id field to given value.


### GetSpaceId

`func (o *ProfilesWarehouseAlpha) GetSpaceId() string`

GetSpaceId returns the SpaceId field if non-nil, zero value otherwise.

### GetSpaceIdOk

`func (o *ProfilesWarehouseAlpha) GetSpaceIdOk() (*string, bool)`

GetSpaceIdOk returns a tuple with the SpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceId

`func (o *ProfilesWarehouseAlpha) SetSpaceId(v string)`

SetSpaceId sets SpaceId field to given value.


### GetMetadata

`func (o *ProfilesWarehouseAlpha) GetMetadata() WarehouseMetadataV1`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ProfilesWarehouseAlpha) GetMetadataOk() (*WarehouseMetadataV1, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ProfilesWarehouseAlpha) SetMetadata(v WarehouseMetadataV1)`

SetMetadata sets Metadata field to given value.


### GetWorkspaceId

`func (o *ProfilesWarehouseAlpha) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *ProfilesWarehouseAlpha) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *ProfilesWarehouseAlpha) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.


### GetEnabled

`func (o *ProfilesWarehouseAlpha) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ProfilesWarehouseAlpha) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ProfilesWarehouseAlpha) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetSettings

`func (o *ProfilesWarehouseAlpha) GetSettings() map[string]interface{}`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *ProfilesWarehouseAlpha) GetSettingsOk() (*map[string]interface{}, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *ProfilesWarehouseAlpha) SetSettings(v map[string]interface{})`

SetSettings sets Settings field to given value.


### GetSchemaName

`func (o *ProfilesWarehouseAlpha) GetSchemaName() string`

GetSchemaName returns the SchemaName field if non-nil, zero value otherwise.

### GetSchemaNameOk

`func (o *ProfilesWarehouseAlpha) GetSchemaNameOk() (*string, bool)`

GetSchemaNameOk returns a tuple with the SchemaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaName

`func (o *ProfilesWarehouseAlpha) SetSchemaName(v string)`

SetSchemaName sets SchemaName field to given value.

### HasSchemaName

`func (o *ProfilesWarehouseAlpha) HasSchemaName() bool`

HasSchemaName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


