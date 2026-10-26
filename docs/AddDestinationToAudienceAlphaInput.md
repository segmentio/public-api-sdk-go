# AddDestinationToAudienceAlphaInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Destination** | [**DestinationInput**](DestinationInput.md) |  | 
**IdSyncConfiguration** | Pointer to [**[]IDSyncConfigurationInput**](IDSyncConfigurationInput.md) | Identifier sync configuration. Defines which external ids to sync and their selection strategies. Maximum 5 items allowed. If omitted, the default will be last email and last user_id, or all email and all user_id depending on the Destination. | [optional] 
**ConnectionSettings** | Pointer to **interface{}** | The settings that a Destination requires to create audiences on a third-party platform. These settings are Destination-specific. Use the List Supported Destinations from Audience endpoint to find the required connection settings. | [optional] 

## Methods

### NewAddDestinationToAudienceAlphaInput

`func NewAddDestinationToAudienceAlphaInput(destination DestinationInput, ) *AddDestinationToAudienceAlphaInput`

NewAddDestinationToAudienceAlphaInput instantiates a new AddDestinationToAudienceAlphaInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddDestinationToAudienceAlphaInputWithDefaults

`func NewAddDestinationToAudienceAlphaInputWithDefaults() *AddDestinationToAudienceAlphaInput`

NewAddDestinationToAudienceAlphaInputWithDefaults instantiates a new AddDestinationToAudienceAlphaInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestination

`func (o *AddDestinationToAudienceAlphaInput) GetDestination() DestinationInput`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *AddDestinationToAudienceAlphaInput) GetDestinationOk() (*DestinationInput, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *AddDestinationToAudienceAlphaInput) SetDestination(v DestinationInput)`

SetDestination sets Destination field to given value.


### GetIdSyncConfiguration

`func (o *AddDestinationToAudienceAlphaInput) GetIdSyncConfiguration() []IDSyncConfigurationInput`

GetIdSyncConfiguration returns the IdSyncConfiguration field if non-nil, zero value otherwise.

### GetIdSyncConfigurationOk

`func (o *AddDestinationToAudienceAlphaInput) GetIdSyncConfigurationOk() (*[]IDSyncConfigurationInput, bool)`

GetIdSyncConfigurationOk returns a tuple with the IdSyncConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdSyncConfiguration

`func (o *AddDestinationToAudienceAlphaInput) SetIdSyncConfiguration(v []IDSyncConfigurationInput)`

SetIdSyncConfiguration sets IdSyncConfiguration field to given value.

### HasIdSyncConfiguration

`func (o *AddDestinationToAudienceAlphaInput) HasIdSyncConfiguration() bool`

HasIdSyncConfiguration returns a boolean if a field has been set.

### GetConnectionSettings

`func (o *AddDestinationToAudienceAlphaInput) GetConnectionSettings() interface{}`

GetConnectionSettings returns the ConnectionSettings field if non-nil, zero value otherwise.

### GetConnectionSettingsOk

`func (o *AddDestinationToAudienceAlphaInput) GetConnectionSettingsOk() (*interface{}, bool)`

GetConnectionSettingsOk returns a tuple with the ConnectionSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionSettings

`func (o *AddDestinationToAudienceAlphaInput) SetConnectionSettings(v interface{})`

SetConnectionSettings sets ConnectionSettings field to given value.

### HasConnectionSettings

`func (o *AddDestinationToAudienceAlphaInput) HasConnectionSettings() bool`

HasConnectionSettings returns a boolean if a field has been set.

### SetConnectionSettingsNil

`func (o *AddDestinationToAudienceAlphaInput) SetConnectionSettingsNil(b bool)`

 SetConnectionSettingsNil sets the value for ConnectionSettings to be an explicit nil

### UnsetConnectionSettings
`func (o *AddDestinationToAudienceAlphaInput) UnsetConnectionSettings()`

UnsetConnectionSettings ensures that no value is present for ConnectionSettings, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


