# UpdateFunctionV1Input

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | The Function code. | [optional] 
**Settings** | Pointer to [**[]FunctionSettingV1**](FunctionSettingV1.md) | The list of settings for this Function. | [optional] 
**DisplayName** | Pointer to **string** | A display name for this Function. | [optional] 
**LogoUrl** | Pointer to **string** | A logo for this Function. | [optional] 
**Description** | Pointer to **string** | A description for this Function. | [optional] 

## Methods

### NewUpdateFunctionV1Input

`func NewUpdateFunctionV1Input() *UpdateFunctionV1Input`

NewUpdateFunctionV1Input instantiates a new UpdateFunctionV1Input object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateFunctionV1InputWithDefaults

`func NewUpdateFunctionV1InputWithDefaults() *UpdateFunctionV1Input`

NewUpdateFunctionV1InputWithDefaults instantiates a new UpdateFunctionV1Input object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *UpdateFunctionV1Input) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *UpdateFunctionV1Input) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *UpdateFunctionV1Input) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *UpdateFunctionV1Input) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetSettings

`func (o *UpdateFunctionV1Input) GetSettings() []FunctionSettingV1`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *UpdateFunctionV1Input) GetSettingsOk() (*[]FunctionSettingV1, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *UpdateFunctionV1Input) SetSettings(v []FunctionSettingV1)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *UpdateFunctionV1Input) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetDisplayName

`func (o *UpdateFunctionV1Input) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *UpdateFunctionV1Input) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *UpdateFunctionV1Input) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *UpdateFunctionV1Input) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetLogoUrl

`func (o *UpdateFunctionV1Input) GetLogoUrl() string`

GetLogoUrl returns the LogoUrl field if non-nil, zero value otherwise.

### GetLogoUrlOk

`func (o *UpdateFunctionV1Input) GetLogoUrlOk() (*string, bool)`

GetLogoUrlOk returns a tuple with the LogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrl

`func (o *UpdateFunctionV1Input) SetLogoUrl(v string)`

SetLogoUrl sets LogoUrl field to given value.

### HasLogoUrl

`func (o *UpdateFunctionV1Input) HasLogoUrl() bool`

HasLogoUrl returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateFunctionV1Input) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateFunctionV1Input) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateFunctionV1Input) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateFunctionV1Input) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


