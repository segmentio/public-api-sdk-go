# FunctionSettingV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of this Function Setting. | 
**Label** | **string** | The label for this Function Setting. | 
**Description** | **string** | A description of this Function Setting. | 
**Type** | **string** | The type of this Function Setting. | 
**Required** | **bool** | Whether this Function Setting is required. | 
**Sensitive** | **bool** | Whether this Function Setting contains sensitive information. | 

## Methods

### NewFunctionSettingV1

`func NewFunctionSettingV1(name string, label string, description string, type_ string, required bool, sensitive bool, ) *FunctionSettingV1`

NewFunctionSettingV1 instantiates a new FunctionSettingV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunctionSettingV1WithDefaults

`func NewFunctionSettingV1WithDefaults() *FunctionSettingV1`

NewFunctionSettingV1WithDefaults instantiates a new FunctionSettingV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *FunctionSettingV1) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FunctionSettingV1) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FunctionSettingV1) SetName(v string)`

SetName sets Name field to given value.


### GetLabel

`func (o *FunctionSettingV1) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *FunctionSettingV1) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *FunctionSettingV1) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetDescription

`func (o *FunctionSettingV1) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FunctionSettingV1) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FunctionSettingV1) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetType

`func (o *FunctionSettingV1) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FunctionSettingV1) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FunctionSettingV1) SetType(v string)`

SetType sets Type field to given value.


### GetRequired

`func (o *FunctionSettingV1) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *FunctionSettingV1) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *FunctionSettingV1) SetRequired(v bool)`

SetRequired sets Required field to given value.


### GetSensitive

`func (o *FunctionSettingV1) GetSensitive() bool`

GetSensitive returns the Sensitive field if non-nil, zero value otherwise.

### GetSensitiveOk

`func (o *FunctionSettingV1) GetSensitiveOk() (*bool, bool)`

GetSensitiveOk returns a tuple with the Sensitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensitive

`func (o *FunctionSettingV1) SetSensitive(v bool)`

SetSensitive sets Sensitive field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


