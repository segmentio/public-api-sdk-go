# CreateComputedTraitAlphaInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the computation. | 
**Enabled** | Pointer to **bool** | Determines whether a computation is enabled. | [optional] 
**Description** | Pointer to **string** | The description of the computation. | [optional] 
**Definition** | [**TraitDefinition**](TraitDefinition.md) |  | 
**Options** | Pointer to [**TraitOptions**](TraitOptions.md) |  | [optional] 

## Methods

### NewCreateComputedTraitAlphaInput

`func NewCreateComputedTraitAlphaInput(name string, definition TraitDefinition, ) *CreateComputedTraitAlphaInput`

NewCreateComputedTraitAlphaInput instantiates a new CreateComputedTraitAlphaInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateComputedTraitAlphaInputWithDefaults

`func NewCreateComputedTraitAlphaInputWithDefaults() *CreateComputedTraitAlphaInput`

NewCreateComputedTraitAlphaInputWithDefaults instantiates a new CreateComputedTraitAlphaInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateComputedTraitAlphaInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateComputedTraitAlphaInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateComputedTraitAlphaInput) SetName(v string)`

SetName sets Name field to given value.


### GetEnabled

`func (o *CreateComputedTraitAlphaInput) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CreateComputedTraitAlphaInput) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CreateComputedTraitAlphaInput) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CreateComputedTraitAlphaInput) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetDescription

`func (o *CreateComputedTraitAlphaInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateComputedTraitAlphaInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateComputedTraitAlphaInput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateComputedTraitAlphaInput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDefinition

`func (o *CreateComputedTraitAlphaInput) GetDefinition() TraitDefinition`

GetDefinition returns the Definition field if non-nil, zero value otherwise.

### GetDefinitionOk

`func (o *CreateComputedTraitAlphaInput) GetDefinitionOk() (*TraitDefinition, bool)`

GetDefinitionOk returns a tuple with the Definition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinition

`func (o *CreateComputedTraitAlphaInput) SetDefinition(v TraitDefinition)`

SetDefinition sets Definition field to given value.


### GetOptions

`func (o *CreateComputedTraitAlphaInput) GetOptions() TraitOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *CreateComputedTraitAlphaInput) GetOptionsOk() (*TraitOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *CreateComputedTraitAlphaInput) SetOptions(v TraitOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *CreateComputedTraitAlphaInput) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


