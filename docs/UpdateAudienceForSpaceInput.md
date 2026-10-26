# UpdateAudienceForSpaceInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Enabled/disabled status for the audience. | [optional] 
**Name** | Pointer to **string** | The name of the computation. | [optional] 
**Description** | Pointer to **string** | The description of the computation. | [optional] 
**Definition** | Pointer to [**AudienceDefinitionWithoutTargetEntity**](AudienceDefinitionWithoutTargetEntity.md) |  | [optional] 
**Options** | Pointer to [**AudienceOptions**](AudienceOptions.md) |  | [optional] 

## Methods

### NewUpdateAudienceForSpaceInput

`func NewUpdateAudienceForSpaceInput() *UpdateAudienceForSpaceInput`

NewUpdateAudienceForSpaceInput instantiates a new UpdateAudienceForSpaceInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAudienceForSpaceInputWithDefaults

`func NewUpdateAudienceForSpaceInputWithDefaults() *UpdateAudienceForSpaceInput`

NewUpdateAudienceForSpaceInputWithDefaults instantiates a new UpdateAudienceForSpaceInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *UpdateAudienceForSpaceInput) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateAudienceForSpaceInput) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateAudienceForSpaceInput) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateAudienceForSpaceInput) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetName

`func (o *UpdateAudienceForSpaceInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateAudienceForSpaceInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateAudienceForSpaceInput) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateAudienceForSpaceInput) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateAudienceForSpaceInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateAudienceForSpaceInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateAudienceForSpaceInput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateAudienceForSpaceInput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDefinition

`func (o *UpdateAudienceForSpaceInput) GetDefinition() AudienceDefinitionWithoutTargetEntity`

GetDefinition returns the Definition field if non-nil, zero value otherwise.

### GetDefinitionOk

`func (o *UpdateAudienceForSpaceInput) GetDefinitionOk() (*AudienceDefinitionWithoutTargetEntity, bool)`

GetDefinitionOk returns a tuple with the Definition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinition

`func (o *UpdateAudienceForSpaceInput) SetDefinition(v AudienceDefinitionWithoutTargetEntity)`

SetDefinition sets Definition field to given value.

### HasDefinition

`func (o *UpdateAudienceForSpaceInput) HasDefinition() bool`

HasDefinition returns a boolean if a field has been set.

### GetOptions

`func (o *UpdateAudienceForSpaceInput) GetOptions() AudienceOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *UpdateAudienceForSpaceInput) GetOptionsOk() (*AudienceOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *UpdateAudienceForSpaceInput) SetOptions(v AudienceOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *UpdateAudienceForSpaceInput) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


