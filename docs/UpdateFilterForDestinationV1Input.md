# UpdateFilterForDestinationV1Input

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**If** | Pointer to **string** | The FQL if condition to update. | [optional] 
**Actions** | Pointer to [**[]DestinationFilterActionV1**](DestinationFilterActionV1.md) | Actions for this Destination filter. | [optional] 
**Title** | Pointer to **string** | The title to update. | [optional] 
**Description** | Pointer to **NullableString** | The description of this filter. | [optional] 
**Enabled** | Pointer to **bool** | When set to true, this Destination filter is active. | [optional] 

## Methods

### NewUpdateFilterForDestinationV1Input

`func NewUpdateFilterForDestinationV1Input() *UpdateFilterForDestinationV1Input`

NewUpdateFilterForDestinationV1Input instantiates a new UpdateFilterForDestinationV1Input object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateFilterForDestinationV1InputWithDefaults

`func NewUpdateFilterForDestinationV1InputWithDefaults() *UpdateFilterForDestinationV1Input`

NewUpdateFilterForDestinationV1InputWithDefaults instantiates a new UpdateFilterForDestinationV1Input object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIf

`func (o *UpdateFilterForDestinationV1Input) GetIf() string`

GetIf returns the If field if non-nil, zero value otherwise.

### GetIfOk

`func (o *UpdateFilterForDestinationV1Input) GetIfOk() (*string, bool)`

GetIfOk returns a tuple with the If field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIf

`func (o *UpdateFilterForDestinationV1Input) SetIf(v string)`

SetIf sets If field to given value.

### HasIf

`func (o *UpdateFilterForDestinationV1Input) HasIf() bool`

HasIf returns a boolean if a field has been set.

### GetActions

`func (o *UpdateFilterForDestinationV1Input) GetActions() []DestinationFilterActionV1`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *UpdateFilterForDestinationV1Input) GetActionsOk() (*[]DestinationFilterActionV1, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *UpdateFilterForDestinationV1Input) SetActions(v []DestinationFilterActionV1)`

SetActions sets Actions field to given value.

### HasActions

`func (o *UpdateFilterForDestinationV1Input) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetTitle

`func (o *UpdateFilterForDestinationV1Input) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *UpdateFilterForDestinationV1Input) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *UpdateFilterForDestinationV1Input) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *UpdateFilterForDestinationV1Input) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateFilterForDestinationV1Input) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateFilterForDestinationV1Input) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateFilterForDestinationV1Input) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateFilterForDestinationV1Input) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UpdateFilterForDestinationV1Input) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UpdateFilterForDestinationV1Input) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEnabled

`func (o *UpdateFilterForDestinationV1Input) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateFilterForDestinationV1Input) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateFilterForDestinationV1Input) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateFilterForDestinationV1Input) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


