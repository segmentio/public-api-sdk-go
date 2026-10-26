# UpdateDestinationForAudienceAlphaInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdSyncConfiguration** | Pointer to [**[]IDSyncConfigurationInput**](IDSyncConfigurationInput.md) | Identifier sync configuration - array of external IDs to sync with their strategies. Maximum 5 items allowed. The supplied array replaces the entire stored configuration, so include every external id that should remain. Omitting mapTo on an entry clears any value stored for it. An empty array is treated the same as omitting the field and leaves the stored configuration unchanged. | [optional] 
**ConnectionSettings** | Pointer to **interface{}** | The settings that a Destination requires to create audiences on a third-party platform. The supplied object replaces the stored settings outright, so include every setting that should remain. Settings that are omitted are removed. An empty object is treated the same as omitting the field and leaves the stored settings unchanged. | [optional] 

## Methods

### NewUpdateDestinationForAudienceAlphaInput

`func NewUpdateDestinationForAudienceAlphaInput() *UpdateDestinationForAudienceAlphaInput`

NewUpdateDestinationForAudienceAlphaInput instantiates a new UpdateDestinationForAudienceAlphaInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDestinationForAudienceAlphaInputWithDefaults

`func NewUpdateDestinationForAudienceAlphaInputWithDefaults() *UpdateDestinationForAudienceAlphaInput`

NewUpdateDestinationForAudienceAlphaInputWithDefaults instantiates a new UpdateDestinationForAudienceAlphaInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdSyncConfiguration

`func (o *UpdateDestinationForAudienceAlphaInput) GetIdSyncConfiguration() []IDSyncConfigurationInput`

GetIdSyncConfiguration returns the IdSyncConfiguration field if non-nil, zero value otherwise.

### GetIdSyncConfigurationOk

`func (o *UpdateDestinationForAudienceAlphaInput) GetIdSyncConfigurationOk() (*[]IDSyncConfigurationInput, bool)`

GetIdSyncConfigurationOk returns a tuple with the IdSyncConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdSyncConfiguration

`func (o *UpdateDestinationForAudienceAlphaInput) SetIdSyncConfiguration(v []IDSyncConfigurationInput)`

SetIdSyncConfiguration sets IdSyncConfiguration field to given value.

### HasIdSyncConfiguration

`func (o *UpdateDestinationForAudienceAlphaInput) HasIdSyncConfiguration() bool`

HasIdSyncConfiguration returns a boolean if a field has been set.

### GetConnectionSettings

`func (o *UpdateDestinationForAudienceAlphaInput) GetConnectionSettings() interface{}`

GetConnectionSettings returns the ConnectionSettings field if non-nil, zero value otherwise.

### GetConnectionSettingsOk

`func (o *UpdateDestinationForAudienceAlphaInput) GetConnectionSettingsOk() (*interface{}, bool)`

GetConnectionSettingsOk returns a tuple with the ConnectionSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionSettings

`func (o *UpdateDestinationForAudienceAlphaInput) SetConnectionSettings(v interface{})`

SetConnectionSettings sets ConnectionSettings field to given value.

### HasConnectionSettings

`func (o *UpdateDestinationForAudienceAlphaInput) HasConnectionSettings() bool`

HasConnectionSettings returns a boolean if a field has been set.

### SetConnectionSettingsNil

`func (o *UpdateDestinationForAudienceAlphaInput) SetConnectionSettingsNil(b bool)`

 SetConnectionSettingsNil sets the value for ConnectionSettings to be an explicit nil

### UnsetConnectionSettings
`func (o *UpdateDestinationForAudienceAlphaInput) UnsetConnectionSettings()`

UnsetConnectionSettings ensures that no value is present for ConnectionSettings, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


