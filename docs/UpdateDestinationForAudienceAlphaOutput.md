# UpdateDestinationForAudienceAlphaOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Destination** | [**Destination**](Destination.md) |  | 
**IdSyncConfiguration** | Pointer to [**[]IDSyncConfigurationInput**](IDSyncConfigurationInput.md) | The id sync configuration for the Destination - array of external ids with their strategies. | [optional] 
**ConnectionSettings** | Pointer to **interface{}** | The settings that a Destination requires to create audiences on a third-party platform. | [optional] 

## Methods

### NewUpdateDestinationForAudienceAlphaOutput

`func NewUpdateDestinationForAudienceAlphaOutput(destination Destination, ) *UpdateDestinationForAudienceAlphaOutput`

NewUpdateDestinationForAudienceAlphaOutput instantiates a new UpdateDestinationForAudienceAlphaOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDestinationForAudienceAlphaOutputWithDefaults

`func NewUpdateDestinationForAudienceAlphaOutputWithDefaults() *UpdateDestinationForAudienceAlphaOutput`

NewUpdateDestinationForAudienceAlphaOutputWithDefaults instantiates a new UpdateDestinationForAudienceAlphaOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestination

`func (o *UpdateDestinationForAudienceAlphaOutput) GetDestination() Destination`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *UpdateDestinationForAudienceAlphaOutput) GetDestinationOk() (*Destination, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *UpdateDestinationForAudienceAlphaOutput) SetDestination(v Destination)`

SetDestination sets Destination field to given value.


### GetIdSyncConfiguration

`func (o *UpdateDestinationForAudienceAlphaOutput) GetIdSyncConfiguration() []IDSyncConfigurationInput`

GetIdSyncConfiguration returns the IdSyncConfiguration field if non-nil, zero value otherwise.

### GetIdSyncConfigurationOk

`func (o *UpdateDestinationForAudienceAlphaOutput) GetIdSyncConfigurationOk() (*[]IDSyncConfigurationInput, bool)`

GetIdSyncConfigurationOk returns a tuple with the IdSyncConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdSyncConfiguration

`func (o *UpdateDestinationForAudienceAlphaOutput) SetIdSyncConfiguration(v []IDSyncConfigurationInput)`

SetIdSyncConfiguration sets IdSyncConfiguration field to given value.

### HasIdSyncConfiguration

`func (o *UpdateDestinationForAudienceAlphaOutput) HasIdSyncConfiguration() bool`

HasIdSyncConfiguration returns a boolean if a field has been set.

### GetConnectionSettings

`func (o *UpdateDestinationForAudienceAlphaOutput) GetConnectionSettings() interface{}`

GetConnectionSettings returns the ConnectionSettings field if non-nil, zero value otherwise.

### GetConnectionSettingsOk

`func (o *UpdateDestinationForAudienceAlphaOutput) GetConnectionSettingsOk() (*interface{}, bool)`

GetConnectionSettingsOk returns a tuple with the ConnectionSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionSettings

`func (o *UpdateDestinationForAudienceAlphaOutput) SetConnectionSettings(v interface{})`

SetConnectionSettings sets ConnectionSettings field to given value.

### HasConnectionSettings

`func (o *UpdateDestinationForAudienceAlphaOutput) HasConnectionSettings() bool`

HasConnectionSettings returns a boolean if a field has been set.

### SetConnectionSettingsNil

`func (o *UpdateDestinationForAudienceAlphaOutput) SetConnectionSettingsNil(b bool)`

 SetConnectionSettingsNil sets the value for ConnectionSettings to be an explicit nil

### UnsetConnectionSettings
`func (o *UpdateDestinationForAudienceAlphaOutput) UnsetConnectionSettings()`

UnsetConnectionSettings ensures that no value is present for ConnectionSettings, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


