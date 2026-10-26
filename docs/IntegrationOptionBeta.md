# IntegrationOptionBeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name identifying this option in the context of a Segment Integration. | 
**Type** | **string** | Defines the type for this option in the schema. Types are most commonly strings, but may also represent other primitive types, such as booleans, and numbers, as well as complex types, such as objects and arrays. | 
**Required** | **bool** | Whether this is a required option when setting up the Integration. | 
**Description** | Pointer to **string** | An optional short text description of the field. | [optional] 
**DefaultValue** | Pointer to **interface{}** | An optional default value for the field. | [optional] 
**Label** | Pointer to **string** | An optional label for this field. | [optional] 

## Methods

### NewIntegrationOptionBeta

`func NewIntegrationOptionBeta(name string, type_ string, required bool, ) *IntegrationOptionBeta`

NewIntegrationOptionBeta instantiates a new IntegrationOptionBeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationOptionBetaWithDefaults

`func NewIntegrationOptionBetaWithDefaults() *IntegrationOptionBeta`

NewIntegrationOptionBetaWithDefaults instantiates a new IntegrationOptionBeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *IntegrationOptionBeta) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IntegrationOptionBeta) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IntegrationOptionBeta) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *IntegrationOptionBeta) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IntegrationOptionBeta) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IntegrationOptionBeta) SetType(v string)`

SetType sets Type field to given value.


### GetRequired

`func (o *IntegrationOptionBeta) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *IntegrationOptionBeta) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *IntegrationOptionBeta) SetRequired(v bool)`

SetRequired sets Required field to given value.


### GetDescription

`func (o *IntegrationOptionBeta) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IntegrationOptionBeta) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IntegrationOptionBeta) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IntegrationOptionBeta) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDefaultValue

`func (o *IntegrationOptionBeta) GetDefaultValue() interface{}`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *IntegrationOptionBeta) GetDefaultValueOk() (*interface{}, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *IntegrationOptionBeta) SetDefaultValue(v interface{})`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *IntegrationOptionBeta) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### SetDefaultValueNil

`func (o *IntegrationOptionBeta) SetDefaultValueNil(b bool)`

 SetDefaultValueNil sets the value for DefaultValue to be an explicit nil

### UnsetDefaultValue
`func (o *IntegrationOptionBeta) UnsetDefaultValue()`

UnsetDefaultValue ensures that no value is present for DefaultValue, not even an explicit nil
### GetLabel

`func (o *IntegrationOptionBeta) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *IntegrationOptionBeta) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *IntegrationOptionBeta) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *IntegrationOptionBeta) HasLabel() bool`

HasLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


