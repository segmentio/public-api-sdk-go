# DestinationSubscriptionConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionId** | **string** | The id for the action that Segment should perform on the Destination. The action cannot be changed once the Activation exists. When updating, this field is still required, but supplying a different value has no effect. | 
**Settings** | Pointer to **map[string]interface{}** | Represents settings used to configure an action subscription. | [optional] 

## Methods

### NewDestinationSubscriptionConfiguration

`func NewDestinationSubscriptionConfiguration(actionId string, ) *DestinationSubscriptionConfiguration`

NewDestinationSubscriptionConfiguration instantiates a new DestinationSubscriptionConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDestinationSubscriptionConfigurationWithDefaults

`func NewDestinationSubscriptionConfigurationWithDefaults() *DestinationSubscriptionConfiguration`

NewDestinationSubscriptionConfigurationWithDefaults instantiates a new DestinationSubscriptionConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionId

`func (o *DestinationSubscriptionConfiguration) GetActionId() string`

GetActionId returns the ActionId field if non-nil, zero value otherwise.

### GetActionIdOk

`func (o *DestinationSubscriptionConfiguration) GetActionIdOk() (*string, bool)`

GetActionIdOk returns a tuple with the ActionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionId

`func (o *DestinationSubscriptionConfiguration) SetActionId(v string)`

SetActionId sets ActionId field to given value.


### GetSettings

`func (o *DestinationSubscriptionConfiguration) GetSettings() map[string]interface{}`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *DestinationSubscriptionConfiguration) GetSettingsOk() (*map[string]interface{}, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *DestinationSubscriptionConfiguration) SetSettings(v map[string]interface{})`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *DestinationSubscriptionConfiguration) HasSettings() bool`

HasSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


