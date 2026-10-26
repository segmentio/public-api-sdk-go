# CreateInsertFunctionInstanceAlphaInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FunctionId** | **string** | Insert Function id to which this instance is associated. Note: Remove the ifnd_/ifns_ prefix from the id. | 
**IntegrationId** | **string** | The Source or Destination id to be connected. | 
**Enabled** | Pointer to **bool** | Whether this insert Function instance should be enabled for the Destination. | [optional] 
**Name** | **string** | Defines the display name of the insert Function instance. | 
**Settings** | **map[string]interface{}** | An object that contains settings for this insert Function instance based on the settings present in the insert Function class. | 

## Methods

### NewCreateInsertFunctionInstanceAlphaInput

`func NewCreateInsertFunctionInstanceAlphaInput(functionId string, integrationId string, name string, settings map[string]interface{}, ) *CreateInsertFunctionInstanceAlphaInput`

NewCreateInsertFunctionInstanceAlphaInput instantiates a new CreateInsertFunctionInstanceAlphaInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateInsertFunctionInstanceAlphaInputWithDefaults

`func NewCreateInsertFunctionInstanceAlphaInputWithDefaults() *CreateInsertFunctionInstanceAlphaInput`

NewCreateInsertFunctionInstanceAlphaInputWithDefaults instantiates a new CreateInsertFunctionInstanceAlphaInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFunctionId

`func (o *CreateInsertFunctionInstanceAlphaInput) GetFunctionId() string`

GetFunctionId returns the FunctionId field if non-nil, zero value otherwise.

### GetFunctionIdOk

`func (o *CreateInsertFunctionInstanceAlphaInput) GetFunctionIdOk() (*string, bool)`

GetFunctionIdOk returns a tuple with the FunctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionId

`func (o *CreateInsertFunctionInstanceAlphaInput) SetFunctionId(v string)`

SetFunctionId sets FunctionId field to given value.


### GetIntegrationId

`func (o *CreateInsertFunctionInstanceAlphaInput) GetIntegrationId() string`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *CreateInsertFunctionInstanceAlphaInput) GetIntegrationIdOk() (*string, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationId

`func (o *CreateInsertFunctionInstanceAlphaInput) SetIntegrationId(v string)`

SetIntegrationId sets IntegrationId field to given value.


### GetEnabled

`func (o *CreateInsertFunctionInstanceAlphaInput) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CreateInsertFunctionInstanceAlphaInput) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CreateInsertFunctionInstanceAlphaInput) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CreateInsertFunctionInstanceAlphaInput) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetName

`func (o *CreateInsertFunctionInstanceAlphaInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateInsertFunctionInstanceAlphaInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateInsertFunctionInstanceAlphaInput) SetName(v string)`

SetName sets Name field to given value.


### GetSettings

`func (o *CreateInsertFunctionInstanceAlphaInput) GetSettings() map[string]interface{}`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *CreateInsertFunctionInstanceAlphaInput) GetSettingsOk() (*map[string]interface{}, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *CreateInsertFunctionInstanceAlphaInput) SetSettings(v map[string]interface{})`

SetSettings sets Settings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


