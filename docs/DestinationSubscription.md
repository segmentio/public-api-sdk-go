# DestinationSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier for the subscription. | 
**Name** | **string** | The name of the subscription. | 
**ActionId** | **string** | The unique identifier for the Destination action to trigger. | 
**ActionSlug** | **string** | The URL-friendly key for the associated Destination action. | 
**DestinationId** | **string** | The associated Destination instance id. | 
**Enabled** | **bool** | Is the subscription enabled. | 
**Settings** | **map[string]interface{}** | Represents settings used to configure an action subscription. | 
**Trigger** | **string** | FQL string that describes what events should trigger a Destination action. | 
**ModelId** | Pointer to **string** | The unique identifier for the linked ReverseETLModel, if this part of a Reverse ETL connection. | [optional] 
**ReverseETLSchedule** | Pointer to [**ReverseEtlScheduleDefinition**](ReverseEtlScheduleDefinition.md) |  | [optional] 

## Methods

### NewDestinationSubscription

`func NewDestinationSubscription(id string, name string, actionId string, actionSlug string, destinationId string, enabled bool, settings map[string]interface{}, trigger string, ) *DestinationSubscription`

NewDestinationSubscription instantiates a new DestinationSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDestinationSubscriptionWithDefaults

`func NewDestinationSubscriptionWithDefaults() *DestinationSubscription`

NewDestinationSubscriptionWithDefaults instantiates a new DestinationSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DestinationSubscription) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DestinationSubscription) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DestinationSubscription) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *DestinationSubscription) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DestinationSubscription) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DestinationSubscription) SetName(v string)`

SetName sets Name field to given value.


### GetActionId

`func (o *DestinationSubscription) GetActionId() string`

GetActionId returns the ActionId field if non-nil, zero value otherwise.

### GetActionIdOk

`func (o *DestinationSubscription) GetActionIdOk() (*string, bool)`

GetActionIdOk returns a tuple with the ActionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionId

`func (o *DestinationSubscription) SetActionId(v string)`

SetActionId sets ActionId field to given value.


### GetActionSlug

`func (o *DestinationSubscription) GetActionSlug() string`

GetActionSlug returns the ActionSlug field if non-nil, zero value otherwise.

### GetActionSlugOk

`func (o *DestinationSubscription) GetActionSlugOk() (*string, bool)`

GetActionSlugOk returns a tuple with the ActionSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionSlug

`func (o *DestinationSubscription) SetActionSlug(v string)`

SetActionSlug sets ActionSlug field to given value.


### GetDestinationId

`func (o *DestinationSubscription) GetDestinationId() string`

GetDestinationId returns the DestinationId field if non-nil, zero value otherwise.

### GetDestinationIdOk

`func (o *DestinationSubscription) GetDestinationIdOk() (*string, bool)`

GetDestinationIdOk returns a tuple with the DestinationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationId

`func (o *DestinationSubscription) SetDestinationId(v string)`

SetDestinationId sets DestinationId field to given value.


### GetEnabled

`func (o *DestinationSubscription) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DestinationSubscription) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DestinationSubscription) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetSettings

`func (o *DestinationSubscription) GetSettings() map[string]interface{}`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *DestinationSubscription) GetSettingsOk() (*map[string]interface{}, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *DestinationSubscription) SetSettings(v map[string]interface{})`

SetSettings sets Settings field to given value.


### GetTrigger

`func (o *DestinationSubscription) GetTrigger() string`

GetTrigger returns the Trigger field if non-nil, zero value otherwise.

### GetTriggerOk

`func (o *DestinationSubscription) GetTriggerOk() (*string, bool)`

GetTriggerOk returns a tuple with the Trigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrigger

`func (o *DestinationSubscription) SetTrigger(v string)`

SetTrigger sets Trigger field to given value.


### GetModelId

`func (o *DestinationSubscription) GetModelId() string`

GetModelId returns the ModelId field if non-nil, zero value otherwise.

### GetModelIdOk

`func (o *DestinationSubscription) GetModelIdOk() (*string, bool)`

GetModelIdOk returns a tuple with the ModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelId

`func (o *DestinationSubscription) SetModelId(v string)`

SetModelId sets ModelId field to given value.

### HasModelId

`func (o *DestinationSubscription) HasModelId() bool`

HasModelId returns a boolean if a field has been set.

### GetReverseETLSchedule

`func (o *DestinationSubscription) GetReverseETLSchedule() ReverseEtlScheduleDefinition`

GetReverseETLSchedule returns the ReverseETLSchedule field if non-nil, zero value otherwise.

### GetReverseETLScheduleOk

`func (o *DestinationSubscription) GetReverseETLScheduleOk() (*ReverseEtlScheduleDefinition, bool)`

GetReverseETLScheduleOk returns a tuple with the ReverseETLSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReverseETLSchedule

`func (o *DestinationSubscription) SetReverseETLSchedule(v ReverseEtlScheduleDefinition)`

SetReverseETLSchedule sets ReverseETLSchedule field to given value.

### HasReverseETLSchedule

`func (o *DestinationSubscription) HasReverseETLSchedule() bool`

HasReverseETLSchedule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


