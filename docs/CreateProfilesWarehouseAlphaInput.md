# CreateProfilesWarehouseAlphaInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetadataId** | **string** | The Warehouse metadata to use. | 
**Name** | Pointer to **string** | An optional human-readable name for this Warehouse. | [optional] 
**Enabled** | Pointer to **bool** | Enable to allow this Warehouse to receive data. Defaults to true. | [optional] 
**Settings** | **map[string]interface{}** | A key-value object that contains instance-specific Warehouse settings. | 
**SchemaName** | Pointer to **string** | The custom schema name that Segment uses on the Warehouse side. The space slug value is default otherwise. The schema name cannot be an existing schema name in the Warehouse. To use an existing schema name, please create a profiles Warehouse through the Segment app UI. | [optional] 

## Methods

### NewCreateProfilesWarehouseAlphaInput

`func NewCreateProfilesWarehouseAlphaInput(metadataId string, settings map[string]interface{}, ) *CreateProfilesWarehouseAlphaInput`

NewCreateProfilesWarehouseAlphaInput instantiates a new CreateProfilesWarehouseAlphaInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateProfilesWarehouseAlphaInputWithDefaults

`func NewCreateProfilesWarehouseAlphaInputWithDefaults() *CreateProfilesWarehouseAlphaInput`

NewCreateProfilesWarehouseAlphaInputWithDefaults instantiates a new CreateProfilesWarehouseAlphaInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadataId

`func (o *CreateProfilesWarehouseAlphaInput) GetMetadataId() string`

GetMetadataId returns the MetadataId field if non-nil, zero value otherwise.

### GetMetadataIdOk

`func (o *CreateProfilesWarehouseAlphaInput) GetMetadataIdOk() (*string, bool)`

GetMetadataIdOk returns a tuple with the MetadataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataId

`func (o *CreateProfilesWarehouseAlphaInput) SetMetadataId(v string)`

SetMetadataId sets MetadataId field to given value.


### GetName

`func (o *CreateProfilesWarehouseAlphaInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateProfilesWarehouseAlphaInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateProfilesWarehouseAlphaInput) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateProfilesWarehouseAlphaInput) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEnabled

`func (o *CreateProfilesWarehouseAlphaInput) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CreateProfilesWarehouseAlphaInput) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CreateProfilesWarehouseAlphaInput) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CreateProfilesWarehouseAlphaInput) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetSettings

`func (o *CreateProfilesWarehouseAlphaInput) GetSettings() map[string]interface{}`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *CreateProfilesWarehouseAlphaInput) GetSettingsOk() (*map[string]interface{}, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *CreateProfilesWarehouseAlphaInput) SetSettings(v map[string]interface{})`

SetSettings sets Settings field to given value.


### GetSchemaName

`func (o *CreateProfilesWarehouseAlphaInput) GetSchemaName() string`

GetSchemaName returns the SchemaName field if non-nil, zero value otherwise.

### GetSchemaNameOk

`func (o *CreateProfilesWarehouseAlphaInput) GetSchemaNameOk() (*string, bool)`

GetSchemaNameOk returns a tuple with the SchemaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaName

`func (o *CreateProfilesWarehouseAlphaInput) SetSchemaName(v string)`

SetSchemaName sets SchemaName field to given value.

### HasSchemaName

`func (o *CreateProfilesWarehouseAlphaInput) HasSchemaName() bool`

HasSchemaName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


