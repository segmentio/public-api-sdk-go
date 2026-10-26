# CreateFilterForDestinationV1Input

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceId** | **string** | The id of the Source associated with this filter. | 
**If** | **string** | The filter&#39;s condition. | 
**Actions** | [**[]DestinationFilterActionV1**](DestinationFilterActionV1.md) | Actions for the Destination filter. | 
**Title** | **string** | The title of the filter. | 
**Description** | Pointer to **string** | The description of the filter. | [optional] 
**Enabled** | **bool** | When set to true, the Destination filter is active. | 

## Methods

### NewCreateFilterForDestinationV1Input

`func NewCreateFilterForDestinationV1Input(sourceId string, if_ string, actions []DestinationFilterActionV1, title string, enabled bool, ) *CreateFilterForDestinationV1Input`

NewCreateFilterForDestinationV1Input instantiates a new CreateFilterForDestinationV1Input object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateFilterForDestinationV1InputWithDefaults

`func NewCreateFilterForDestinationV1InputWithDefaults() *CreateFilterForDestinationV1Input`

NewCreateFilterForDestinationV1InputWithDefaults instantiates a new CreateFilterForDestinationV1Input object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceId

`func (o *CreateFilterForDestinationV1Input) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *CreateFilterForDestinationV1Input) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *CreateFilterForDestinationV1Input) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.


### GetIf

`func (o *CreateFilterForDestinationV1Input) GetIf() string`

GetIf returns the If field if non-nil, zero value otherwise.

### GetIfOk

`func (o *CreateFilterForDestinationV1Input) GetIfOk() (*string, bool)`

GetIfOk returns a tuple with the If field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIf

`func (o *CreateFilterForDestinationV1Input) SetIf(v string)`

SetIf sets If field to given value.


### GetActions

`func (o *CreateFilterForDestinationV1Input) GetActions() []DestinationFilterActionV1`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *CreateFilterForDestinationV1Input) GetActionsOk() (*[]DestinationFilterActionV1, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *CreateFilterForDestinationV1Input) SetActions(v []DestinationFilterActionV1)`

SetActions sets Actions field to given value.


### GetTitle

`func (o *CreateFilterForDestinationV1Input) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CreateFilterForDestinationV1Input) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CreateFilterForDestinationV1Input) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *CreateFilterForDestinationV1Input) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateFilterForDestinationV1Input) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateFilterForDestinationV1Input) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateFilterForDestinationV1Input) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *CreateFilterForDestinationV1Input) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CreateFilterForDestinationV1Input) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CreateFilterForDestinationV1Input) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


