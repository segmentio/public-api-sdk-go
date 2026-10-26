# CreateFunctionV1Input

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | The Function code. | 
**Settings** | Pointer to [**[]FunctionSettingV1**](FunctionSettingV1.md) | The list of settings for this Function. | [optional] 
**DisplayName** | **string** | A display name for this Function.  Note that Destination Functions append the Workspace to the display name. | 
**LogoUrl** | Pointer to **string** | The URL of the logo for this Function. | [optional] 
**ResourceType** | **string** | The Function type.  Config API note: equal to &#x60;type&#x60;. | 
**Description** | Pointer to **string** | A description for this Function. | [optional] 

## Methods

### NewCreateFunctionV1Input

`func NewCreateFunctionV1Input(code string, displayName string, resourceType string, ) *CreateFunctionV1Input`

NewCreateFunctionV1Input instantiates a new CreateFunctionV1Input object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateFunctionV1InputWithDefaults

`func NewCreateFunctionV1InputWithDefaults() *CreateFunctionV1Input`

NewCreateFunctionV1InputWithDefaults instantiates a new CreateFunctionV1Input object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *CreateFunctionV1Input) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CreateFunctionV1Input) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CreateFunctionV1Input) SetCode(v string)`

SetCode sets Code field to given value.


### GetSettings

`func (o *CreateFunctionV1Input) GetSettings() []FunctionSettingV1`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *CreateFunctionV1Input) GetSettingsOk() (*[]FunctionSettingV1, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *CreateFunctionV1Input) SetSettings(v []FunctionSettingV1)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *CreateFunctionV1Input) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetDisplayName

`func (o *CreateFunctionV1Input) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CreateFunctionV1Input) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CreateFunctionV1Input) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetLogoUrl

`func (o *CreateFunctionV1Input) GetLogoUrl() string`

GetLogoUrl returns the LogoUrl field if non-nil, zero value otherwise.

### GetLogoUrlOk

`func (o *CreateFunctionV1Input) GetLogoUrlOk() (*string, bool)`

GetLogoUrlOk returns a tuple with the LogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrl

`func (o *CreateFunctionV1Input) SetLogoUrl(v string)`

SetLogoUrl sets LogoUrl field to given value.

### HasLogoUrl

`func (o *CreateFunctionV1Input) HasLogoUrl() bool`

HasLogoUrl returns a boolean if a field has been set.

### GetResourceType

`func (o *CreateFunctionV1Input) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *CreateFunctionV1Input) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *CreateFunctionV1Input) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.


### GetDescription

`func (o *CreateFunctionV1Input) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateFunctionV1Input) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateFunctionV1Input) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateFunctionV1Input) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


