# UpdateDestinationV1Input

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Defines the display name of the Destination.  Config API note: equal to &#x60;displayName&#x60;. | [optional] 
**Enabled** | Pointer to **bool** | Whether this Destination should receive data. | [optional] 
**Settings** | Pointer to **map[string]interface{}** | An optional object that contains settings for the Destination based on the \&quot;required\&quot; and \&quot;advanced\&quot; settings present in the Destination metadata.  Config API note: equal to &#x60;config&#x60;. | [optional] 

## Methods

### NewUpdateDestinationV1Input

`func NewUpdateDestinationV1Input() *UpdateDestinationV1Input`

NewUpdateDestinationV1Input instantiates a new UpdateDestinationV1Input object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDestinationV1InputWithDefaults

`func NewUpdateDestinationV1InputWithDefaults() *UpdateDestinationV1Input`

NewUpdateDestinationV1InputWithDefaults instantiates a new UpdateDestinationV1Input object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateDestinationV1Input) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateDestinationV1Input) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateDestinationV1Input) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateDestinationV1Input) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *UpdateDestinationV1Input) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UpdateDestinationV1Input) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetEnabled

`func (o *UpdateDestinationV1Input) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateDestinationV1Input) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateDestinationV1Input) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateDestinationV1Input) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetSettings

`func (o *UpdateDestinationV1Input) GetSettings() map[string]interface{}`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *UpdateDestinationV1Input) GetSettingsOk() (*map[string]interface{}, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *UpdateDestinationV1Input) SetSettings(v map[string]interface{})`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *UpdateDestinationV1Input) HasSettings() bool`

HasSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


