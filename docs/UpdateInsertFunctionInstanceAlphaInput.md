# UpdateInsertFunctionInstanceAlphaInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Whether this insert Function instance should be enabled for the Destination. | [optional] 
**Name** | Pointer to **string** | Defines the display name of the insert Function instance. | [optional] 
**Settings** | **map[string]interface{}** | An object that contains settings for this insert Function instance based on the settings present in the insert Function class. | 

## Methods

### NewUpdateInsertFunctionInstanceAlphaInput

`func NewUpdateInsertFunctionInstanceAlphaInput(settings map[string]interface{}, ) *UpdateInsertFunctionInstanceAlphaInput`

NewUpdateInsertFunctionInstanceAlphaInput instantiates a new UpdateInsertFunctionInstanceAlphaInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateInsertFunctionInstanceAlphaInputWithDefaults

`func NewUpdateInsertFunctionInstanceAlphaInputWithDefaults() *UpdateInsertFunctionInstanceAlphaInput`

NewUpdateInsertFunctionInstanceAlphaInputWithDefaults instantiates a new UpdateInsertFunctionInstanceAlphaInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *UpdateInsertFunctionInstanceAlphaInput) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateInsertFunctionInstanceAlphaInput) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateInsertFunctionInstanceAlphaInput) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateInsertFunctionInstanceAlphaInput) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetName

`func (o *UpdateInsertFunctionInstanceAlphaInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateInsertFunctionInstanceAlphaInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateInsertFunctionInstanceAlphaInput) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateInsertFunctionInstanceAlphaInput) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSettings

`func (o *UpdateInsertFunctionInstanceAlphaInput) GetSettings() map[string]interface{}`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *UpdateInsertFunctionInstanceAlphaInput) GetSettingsOk() (*map[string]interface{}, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *UpdateInsertFunctionInstanceAlphaInput) SetSettings(v map[string]interface{})`

SetSettings sets Settings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


