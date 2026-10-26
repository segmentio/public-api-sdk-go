# AddDestinationToAudienceAlphaOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Connection** | [**Connection**](Connection.md) |  | 
**IdSyncConfiguration** | [**[]IDSyncConfigurationInput**](IDSyncConfigurationInput.md) | Identifier sync configuration. Defines which external ids to sync and their selection strategies. Maximum 5 items allowed. If omitted, the default will be last email and last user_id, or all email and all user_id depending on the Destination. | 
**ConnectionSettings** | Pointer to **interface{}** | The settings that a Destination requires to create audiences on a third-party platform. These settings are Destination-specific. Use the List Supported Destinations from Audience endpoint to find the required connection settings. | [optional] 

## Methods

### NewAddDestinationToAudienceAlphaOutput

`func NewAddDestinationToAudienceAlphaOutput(connection Connection, idSyncConfiguration []IDSyncConfigurationInput, ) *AddDestinationToAudienceAlphaOutput`

NewAddDestinationToAudienceAlphaOutput instantiates a new AddDestinationToAudienceAlphaOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddDestinationToAudienceAlphaOutputWithDefaults

`func NewAddDestinationToAudienceAlphaOutputWithDefaults() *AddDestinationToAudienceAlphaOutput`

NewAddDestinationToAudienceAlphaOutputWithDefaults instantiates a new AddDestinationToAudienceAlphaOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnection

`func (o *AddDestinationToAudienceAlphaOutput) GetConnection() Connection`

GetConnection returns the Connection field if non-nil, zero value otherwise.

### GetConnectionOk

`func (o *AddDestinationToAudienceAlphaOutput) GetConnectionOk() (*Connection, bool)`

GetConnectionOk returns a tuple with the Connection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnection

`func (o *AddDestinationToAudienceAlphaOutput) SetConnection(v Connection)`

SetConnection sets Connection field to given value.


### GetIdSyncConfiguration

`func (o *AddDestinationToAudienceAlphaOutput) GetIdSyncConfiguration() []IDSyncConfigurationInput`

GetIdSyncConfiguration returns the IdSyncConfiguration field if non-nil, zero value otherwise.

### GetIdSyncConfigurationOk

`func (o *AddDestinationToAudienceAlphaOutput) GetIdSyncConfigurationOk() (*[]IDSyncConfigurationInput, bool)`

GetIdSyncConfigurationOk returns a tuple with the IdSyncConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdSyncConfiguration

`func (o *AddDestinationToAudienceAlphaOutput) SetIdSyncConfiguration(v []IDSyncConfigurationInput)`

SetIdSyncConfiguration sets IdSyncConfiguration field to given value.


### GetConnectionSettings

`func (o *AddDestinationToAudienceAlphaOutput) GetConnectionSettings() interface{}`

GetConnectionSettings returns the ConnectionSettings field if non-nil, zero value otherwise.

### GetConnectionSettingsOk

`func (o *AddDestinationToAudienceAlphaOutput) GetConnectionSettingsOk() (*interface{}, bool)`

GetConnectionSettingsOk returns a tuple with the ConnectionSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionSettings

`func (o *AddDestinationToAudienceAlphaOutput) SetConnectionSettings(v interface{})`

SetConnectionSettings sets ConnectionSettings field to given value.

### HasConnectionSettings

`func (o *AddDestinationToAudienceAlphaOutput) HasConnectionSettings() bool`

HasConnectionSettings returns a boolean if a field has been set.

### SetConnectionSettingsNil

`func (o *AddDestinationToAudienceAlphaOutput) SetConnectionSettingsNil(b bool)`

 SetConnectionSettingsNil sets the value for ConnectionSettings to be an explicit nil

### UnsetConnectionSettings
`func (o *AddDestinationToAudienceAlphaOutput) UnsetConnectionSettings()`

UnsetConnectionSettings ensures that no value is present for ConnectionSettings, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


