# DestinationSubscriptionUpdateInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The user-defined name for the subscription. | [optional] 
**Trigger** | Pointer to **string** | The fql statement. | [optional] 
**Enabled** | Pointer to **bool** | Is the subscription enabled. | [optional] 
**Settings** | Pointer to **map[string]interface{}** | Represents settings used to configure an action subscription. | [optional] 
**ReverseETLModelId** | Pointer to **string** | (Reverse ETL only) The reverse ETL model to attach this subscription to. | [optional] 
**ReverseETLSchedule** | Pointer to [**ReverseEtlScheduleDefinition**](ReverseEtlScheduleDefinition.md) |  | [optional] 

## Methods

### NewDestinationSubscriptionUpdateInput

`func NewDestinationSubscriptionUpdateInput() *DestinationSubscriptionUpdateInput`

NewDestinationSubscriptionUpdateInput instantiates a new DestinationSubscriptionUpdateInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDestinationSubscriptionUpdateInputWithDefaults

`func NewDestinationSubscriptionUpdateInputWithDefaults() *DestinationSubscriptionUpdateInput`

NewDestinationSubscriptionUpdateInputWithDefaults instantiates a new DestinationSubscriptionUpdateInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DestinationSubscriptionUpdateInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DestinationSubscriptionUpdateInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DestinationSubscriptionUpdateInput) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DestinationSubscriptionUpdateInput) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTrigger

`func (o *DestinationSubscriptionUpdateInput) GetTrigger() string`

GetTrigger returns the Trigger field if non-nil, zero value otherwise.

### GetTriggerOk

`func (o *DestinationSubscriptionUpdateInput) GetTriggerOk() (*string, bool)`

GetTriggerOk returns a tuple with the Trigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrigger

`func (o *DestinationSubscriptionUpdateInput) SetTrigger(v string)`

SetTrigger sets Trigger field to given value.

### HasTrigger

`func (o *DestinationSubscriptionUpdateInput) HasTrigger() bool`

HasTrigger returns a boolean if a field has been set.

### GetEnabled

`func (o *DestinationSubscriptionUpdateInput) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DestinationSubscriptionUpdateInput) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DestinationSubscriptionUpdateInput) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *DestinationSubscriptionUpdateInput) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetSettings

`func (o *DestinationSubscriptionUpdateInput) GetSettings() map[string]interface{}`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *DestinationSubscriptionUpdateInput) GetSettingsOk() (*map[string]interface{}, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *DestinationSubscriptionUpdateInput) SetSettings(v map[string]interface{})`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *DestinationSubscriptionUpdateInput) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetReverseETLModelId

`func (o *DestinationSubscriptionUpdateInput) GetReverseETLModelId() string`

GetReverseETLModelId returns the ReverseETLModelId field if non-nil, zero value otherwise.

### GetReverseETLModelIdOk

`func (o *DestinationSubscriptionUpdateInput) GetReverseETLModelIdOk() (*string, bool)`

GetReverseETLModelIdOk returns a tuple with the ReverseETLModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReverseETLModelId

`func (o *DestinationSubscriptionUpdateInput) SetReverseETLModelId(v string)`

SetReverseETLModelId sets ReverseETLModelId field to given value.

### HasReverseETLModelId

`func (o *DestinationSubscriptionUpdateInput) HasReverseETLModelId() bool`

HasReverseETLModelId returns a boolean if a field has been set.

### GetReverseETLSchedule

`func (o *DestinationSubscriptionUpdateInput) GetReverseETLSchedule() ReverseEtlScheduleDefinition`

GetReverseETLSchedule returns the ReverseETLSchedule field if non-nil, zero value otherwise.

### GetReverseETLScheduleOk

`func (o *DestinationSubscriptionUpdateInput) GetReverseETLScheduleOk() (*ReverseEtlScheduleDefinition, bool)`

GetReverseETLScheduleOk returns a tuple with the ReverseETLSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReverseETLSchedule

`func (o *DestinationSubscriptionUpdateInput) SetReverseETLSchedule(v ReverseEtlScheduleDefinition)`

SetReverseETLSchedule sets ReverseETLSchedule field to given value.

### HasReverseETLSchedule

`func (o *DestinationSubscriptionUpdateInput) HasReverseETLSchedule() bool`

HasReverseETLSchedule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


