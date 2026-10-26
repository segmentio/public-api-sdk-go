# UpdateWarehouseV1Input

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | An optional human-readable name to associate with this Warehouse. | [optional] 
**Enabled** | Pointer to **bool** | Enable to allow this Warehouse to receive data. | [optional] 
**Settings** | **map[string]interface{}** | A key-value object that contains instance-specific Warehouse settings. | 

## Methods

### NewUpdateWarehouseV1Input

`func NewUpdateWarehouseV1Input(settings map[string]interface{}, ) *UpdateWarehouseV1Input`

NewUpdateWarehouseV1Input instantiates a new UpdateWarehouseV1Input object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateWarehouseV1InputWithDefaults

`func NewUpdateWarehouseV1InputWithDefaults() *UpdateWarehouseV1Input`

NewUpdateWarehouseV1InputWithDefaults instantiates a new UpdateWarehouseV1Input object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateWarehouseV1Input) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateWarehouseV1Input) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateWarehouseV1Input) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateWarehouseV1Input) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *UpdateWarehouseV1Input) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UpdateWarehouseV1Input) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetEnabled

`func (o *UpdateWarehouseV1Input) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateWarehouseV1Input) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateWarehouseV1Input) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateWarehouseV1Input) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetSettings

`func (o *UpdateWarehouseV1Input) GetSettings() map[string]interface{}`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *UpdateWarehouseV1Input) GetSettingsOk() (*map[string]interface{}, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *UpdateWarehouseV1Input) SetSettings(v map[string]interface{})`

SetSettings sets Settings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


