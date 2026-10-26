# CreateDestinationSubscriptionAlphaInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A user-defined name for the subscription. | 
**ActionId** | **string** | The associated action id the subscription should trigger. | 
**Trigger** | **string** | The FQL statement. | 
**Enabled** | **bool** | Is the subscription enabled. | 
**Settings** | Pointer to **map[string]interface{}** | Represents settings used to configure an action subscription. | [optional] 
**ModelId** | Pointer to **string** | When creating a Reverse ETL connection, indicates the Model being used to extract data. | [optional] 

## Methods

### NewCreateDestinationSubscriptionAlphaInput

`func NewCreateDestinationSubscriptionAlphaInput(name string, actionId string, trigger string, enabled bool, ) *CreateDestinationSubscriptionAlphaInput`

NewCreateDestinationSubscriptionAlphaInput instantiates a new CreateDestinationSubscriptionAlphaInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDestinationSubscriptionAlphaInputWithDefaults

`func NewCreateDestinationSubscriptionAlphaInputWithDefaults() *CreateDestinationSubscriptionAlphaInput`

NewCreateDestinationSubscriptionAlphaInputWithDefaults instantiates a new CreateDestinationSubscriptionAlphaInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateDestinationSubscriptionAlphaInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateDestinationSubscriptionAlphaInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateDestinationSubscriptionAlphaInput) SetName(v string)`

SetName sets Name field to given value.


### GetActionId

`func (o *CreateDestinationSubscriptionAlphaInput) GetActionId() string`

GetActionId returns the ActionId field if non-nil, zero value otherwise.

### GetActionIdOk

`func (o *CreateDestinationSubscriptionAlphaInput) GetActionIdOk() (*string, bool)`

GetActionIdOk returns a tuple with the ActionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionId

`func (o *CreateDestinationSubscriptionAlphaInput) SetActionId(v string)`

SetActionId sets ActionId field to given value.


### GetTrigger

`func (o *CreateDestinationSubscriptionAlphaInput) GetTrigger() string`

GetTrigger returns the Trigger field if non-nil, zero value otherwise.

### GetTriggerOk

`func (o *CreateDestinationSubscriptionAlphaInput) GetTriggerOk() (*string, bool)`

GetTriggerOk returns a tuple with the Trigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrigger

`func (o *CreateDestinationSubscriptionAlphaInput) SetTrigger(v string)`

SetTrigger sets Trigger field to given value.


### GetEnabled

`func (o *CreateDestinationSubscriptionAlphaInput) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CreateDestinationSubscriptionAlphaInput) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CreateDestinationSubscriptionAlphaInput) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetSettings

`func (o *CreateDestinationSubscriptionAlphaInput) GetSettings() map[string]interface{}`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *CreateDestinationSubscriptionAlphaInput) GetSettingsOk() (*map[string]interface{}, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *CreateDestinationSubscriptionAlphaInput) SetSettings(v map[string]interface{})`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *CreateDestinationSubscriptionAlphaInput) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetModelId

`func (o *CreateDestinationSubscriptionAlphaInput) GetModelId() string`

GetModelId returns the ModelId field if non-nil, zero value otherwise.

### GetModelIdOk

`func (o *CreateDestinationSubscriptionAlphaInput) GetModelIdOk() (*string, bool)`

GetModelIdOk returns a tuple with the ModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelId

`func (o *CreateDestinationSubscriptionAlphaInput) SetModelId(v string)`

SetModelId sets ModelId field to given value.

### HasModelId

`func (o *CreateDestinationSubscriptionAlphaInput) HasModelId() bool`

HasModelId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


