# CreateWarehouseV1Input

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetadataId** | **string** | The Warehouse metadata to use. | 
**Name** | Pointer to **string** | An optional human-readable name for this Warehouse. | [optional] 
**Enabled** | Pointer to **bool** | Enable to allow this Warehouse to receive data. Defaults to true. | [optional] 
**Settings** | **map[string]interface{}** | A key-value object that contains instance-specific Warehouse settings. | 
**DisconnectAllSources** | Pointer to **bool** | Whether to disconnect all Sources from this Warehouse. | [optional] 

## Methods

### NewCreateWarehouseV1Input

`func NewCreateWarehouseV1Input(metadataId string, settings map[string]interface{}, ) *CreateWarehouseV1Input`

NewCreateWarehouseV1Input instantiates a new CreateWarehouseV1Input object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateWarehouseV1InputWithDefaults

`func NewCreateWarehouseV1InputWithDefaults() *CreateWarehouseV1Input`

NewCreateWarehouseV1InputWithDefaults instantiates a new CreateWarehouseV1Input object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadataId

`func (o *CreateWarehouseV1Input) GetMetadataId() string`

GetMetadataId returns the MetadataId field if non-nil, zero value otherwise.

### GetMetadataIdOk

`func (o *CreateWarehouseV1Input) GetMetadataIdOk() (*string, bool)`

GetMetadataIdOk returns a tuple with the MetadataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataId

`func (o *CreateWarehouseV1Input) SetMetadataId(v string)`

SetMetadataId sets MetadataId field to given value.


### GetName

`func (o *CreateWarehouseV1Input) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateWarehouseV1Input) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateWarehouseV1Input) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateWarehouseV1Input) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEnabled

`func (o *CreateWarehouseV1Input) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CreateWarehouseV1Input) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CreateWarehouseV1Input) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CreateWarehouseV1Input) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetSettings

`func (o *CreateWarehouseV1Input) GetSettings() map[string]interface{}`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *CreateWarehouseV1Input) GetSettingsOk() (*map[string]interface{}, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *CreateWarehouseV1Input) SetSettings(v map[string]interface{})`

SetSettings sets Settings field to given value.


### GetDisconnectAllSources

`func (o *CreateWarehouseV1Input) GetDisconnectAllSources() bool`

GetDisconnectAllSources returns the DisconnectAllSources field if non-nil, zero value otherwise.

### GetDisconnectAllSourcesOk

`func (o *CreateWarehouseV1Input) GetDisconnectAllSourcesOk() (*bool, bool)`

GetDisconnectAllSourcesOk returns a tuple with the DisconnectAllSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisconnectAllSources

`func (o *CreateWarehouseV1Input) SetDisconnectAllSources(v bool)`

SetDisconnectAllSources sets DisconnectAllSources field to given value.

### HasDisconnectAllSources

`func (o *CreateWarehouseV1Input) HasDisconnectAllSources() bool`

HasDisconnectAllSources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


