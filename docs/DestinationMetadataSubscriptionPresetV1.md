# DestinationMetadataSubscriptionPresetV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionId** | **string** | The unique identifier for the Destination Action to trigger. | 
**Name** | **string** | The name of the subscription. | 
**Fields** | **map[string]interface{}** | The default settings for action fields. | 
**Trigger** | **string** | FQL string that describes what events should trigger an action. See https://segment.com/docs/config-api/fql/ for more information regarding Segment&#39;s Filter Query Language (FQL). | 

## Methods

### NewDestinationMetadataSubscriptionPresetV1

`func NewDestinationMetadataSubscriptionPresetV1(actionId string, name string, fields map[string]interface{}, trigger string, ) *DestinationMetadataSubscriptionPresetV1`

NewDestinationMetadataSubscriptionPresetV1 instantiates a new DestinationMetadataSubscriptionPresetV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDestinationMetadataSubscriptionPresetV1WithDefaults

`func NewDestinationMetadataSubscriptionPresetV1WithDefaults() *DestinationMetadataSubscriptionPresetV1`

NewDestinationMetadataSubscriptionPresetV1WithDefaults instantiates a new DestinationMetadataSubscriptionPresetV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionId

`func (o *DestinationMetadataSubscriptionPresetV1) GetActionId() string`

GetActionId returns the ActionId field if non-nil, zero value otherwise.

### GetActionIdOk

`func (o *DestinationMetadataSubscriptionPresetV1) GetActionIdOk() (*string, bool)`

GetActionIdOk returns a tuple with the ActionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionId

`func (o *DestinationMetadataSubscriptionPresetV1) SetActionId(v string)`

SetActionId sets ActionId field to given value.


### GetName

`func (o *DestinationMetadataSubscriptionPresetV1) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DestinationMetadataSubscriptionPresetV1) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DestinationMetadataSubscriptionPresetV1) SetName(v string)`

SetName sets Name field to given value.


### GetFields

`func (o *DestinationMetadataSubscriptionPresetV1) GetFields() map[string]interface{}`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *DestinationMetadataSubscriptionPresetV1) GetFieldsOk() (*map[string]interface{}, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *DestinationMetadataSubscriptionPresetV1) SetFields(v map[string]interface{})`

SetFields sets Fields field to given value.


### GetTrigger

`func (o *DestinationMetadataSubscriptionPresetV1) GetTrigger() string`

GetTrigger returns the Trigger field if non-nil, zero value otherwise.

### GetTriggerOk

`func (o *DestinationMetadataSubscriptionPresetV1) GetTriggerOk() (*string, bool)`

GetTriggerOk returns a tuple with the Trigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrigger

`func (o *DestinationMetadataSubscriptionPresetV1) SetTrigger(v string)`

SetTrigger sets Trigger field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


