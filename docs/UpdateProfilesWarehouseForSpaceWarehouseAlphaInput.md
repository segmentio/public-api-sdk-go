# UpdateProfilesWarehouseForSpaceWarehouseAlphaInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | An optional human-readable name for this Warehouse. | [optional] 
**Enabled** | Pointer to **bool** | Enable to allow this Warehouse to receive data. Defaults to true. | [optional] 
**Settings** | **map[string]interface{}** | A key-value object that contains instance-specific Warehouse settings. | 
**SchemaName** | Pointer to **string** | The custom schema name that Segment uses on the Warehouse side. The space slug value is default otherwise. | [optional] 

## Methods

### NewUpdateProfilesWarehouseForSpaceWarehouseAlphaInput

`func NewUpdateProfilesWarehouseForSpaceWarehouseAlphaInput(settings map[string]interface{}, ) *UpdateProfilesWarehouseForSpaceWarehouseAlphaInput`

NewUpdateProfilesWarehouseForSpaceWarehouseAlphaInput instantiates a new UpdateProfilesWarehouseForSpaceWarehouseAlphaInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateProfilesWarehouseForSpaceWarehouseAlphaInputWithDefaults

`func NewUpdateProfilesWarehouseForSpaceWarehouseAlphaInputWithDefaults() *UpdateProfilesWarehouseForSpaceWarehouseAlphaInput`

NewUpdateProfilesWarehouseForSpaceWarehouseAlphaInputWithDefaults instantiates a new UpdateProfilesWarehouseForSpaceWarehouseAlphaInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateProfilesWarehouseForSpaceWarehouseAlphaInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateProfilesWarehouseForSpaceWarehouseAlphaInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateProfilesWarehouseForSpaceWarehouseAlphaInput) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateProfilesWarehouseForSpaceWarehouseAlphaInput) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEnabled

`func (o *UpdateProfilesWarehouseForSpaceWarehouseAlphaInput) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateProfilesWarehouseForSpaceWarehouseAlphaInput) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateProfilesWarehouseForSpaceWarehouseAlphaInput) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateProfilesWarehouseForSpaceWarehouseAlphaInput) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetSettings

`func (o *UpdateProfilesWarehouseForSpaceWarehouseAlphaInput) GetSettings() map[string]interface{}`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *UpdateProfilesWarehouseForSpaceWarehouseAlphaInput) GetSettingsOk() (*map[string]interface{}, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *UpdateProfilesWarehouseForSpaceWarehouseAlphaInput) SetSettings(v map[string]interface{})`

SetSettings sets Settings field to given value.


### GetSchemaName

`func (o *UpdateProfilesWarehouseForSpaceWarehouseAlphaInput) GetSchemaName() string`

GetSchemaName returns the SchemaName field if non-nil, zero value otherwise.

### GetSchemaNameOk

`func (o *UpdateProfilesWarehouseForSpaceWarehouseAlphaInput) GetSchemaNameOk() (*string, bool)`

GetSchemaNameOk returns a tuple with the SchemaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaName

`func (o *UpdateProfilesWarehouseForSpaceWarehouseAlphaInput) SetSchemaName(v string)`

SetSchemaName sets SchemaName field to given value.

### HasSchemaName

`func (o *UpdateProfilesWarehouseForSpaceWarehouseAlphaInput) HasSchemaName() bool`

HasSchemaName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


