# SimpleDestination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The id of the Integration. | 
**Name** | Pointer to **string** | The name of the Destination. | [optional] 
**SourceId** | **string** | The Source of the Destination. | 
**Enabled** | **bool** | Whether the Integration is enabled or not. | 
**CreatedAt** | **string** | When the Integration connection was created. | 
**UpdatedAt** | **string** | When the Integration connection was last updated. | 
**Settings** | **map[string]interface{}** | The Destination settings. | 
**DestinationId** | **string** | The Destination id. | 
**Metadata** | Pointer to [**Metadata**](Metadata.md) |  | [optional] 
**IdSyncConfiguration** | Pointer to [**[]IDSyncConfigurationInput**](IDSyncConfigurationInput.md) | Identifier sync configuration. Defines which external ids to sync and their selection strategies. Maximum 5 items allowed. If omitted, the default will be last email and last user_id, or all email and all user_id depending on the Destination. | [optional] 
**ConnectionSettings** | Pointer to **interface{}** | The settings that a Destination requires to create audiences on a third-party platform. These settings are Destination-specific. Use the List Supported Destinations from Audience endpoint to find the required connection settings. | [optional] 

## Methods

### NewSimpleDestination

`func NewSimpleDestination(id string, sourceId string, enabled bool, createdAt string, updatedAt string, settings map[string]interface{}, destinationId string, ) *SimpleDestination`

NewSimpleDestination instantiates a new SimpleDestination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimpleDestinationWithDefaults

`func NewSimpleDestinationWithDefaults() *SimpleDestination`

NewSimpleDestinationWithDefaults instantiates a new SimpleDestination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SimpleDestination) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SimpleDestination) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SimpleDestination) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *SimpleDestination) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SimpleDestination) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SimpleDestination) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SimpleDestination) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSourceId

`func (o *SimpleDestination) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *SimpleDestination) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *SimpleDestination) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.


### GetEnabled

`func (o *SimpleDestination) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SimpleDestination) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SimpleDestination) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetCreatedAt

`func (o *SimpleDestination) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SimpleDestination) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SimpleDestination) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *SimpleDestination) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SimpleDestination) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SimpleDestination) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetSettings

`func (o *SimpleDestination) GetSettings() map[string]interface{}`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *SimpleDestination) GetSettingsOk() (*map[string]interface{}, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *SimpleDestination) SetSettings(v map[string]interface{})`

SetSettings sets Settings field to given value.


### GetDestinationId

`func (o *SimpleDestination) GetDestinationId() string`

GetDestinationId returns the DestinationId field if non-nil, zero value otherwise.

### GetDestinationIdOk

`func (o *SimpleDestination) GetDestinationIdOk() (*string, bool)`

GetDestinationIdOk returns a tuple with the DestinationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationId

`func (o *SimpleDestination) SetDestinationId(v string)`

SetDestinationId sets DestinationId field to given value.


### GetMetadata

`func (o *SimpleDestination) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SimpleDestination) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SimpleDestination) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *SimpleDestination) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetIdSyncConfiguration

`func (o *SimpleDestination) GetIdSyncConfiguration() []IDSyncConfigurationInput`

GetIdSyncConfiguration returns the IdSyncConfiguration field if non-nil, zero value otherwise.

### GetIdSyncConfigurationOk

`func (o *SimpleDestination) GetIdSyncConfigurationOk() (*[]IDSyncConfigurationInput, bool)`

GetIdSyncConfigurationOk returns a tuple with the IdSyncConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdSyncConfiguration

`func (o *SimpleDestination) SetIdSyncConfiguration(v []IDSyncConfigurationInput)`

SetIdSyncConfiguration sets IdSyncConfiguration field to given value.

### HasIdSyncConfiguration

`func (o *SimpleDestination) HasIdSyncConfiguration() bool`

HasIdSyncConfiguration returns a boolean if a field has been set.

### GetConnectionSettings

`func (o *SimpleDestination) GetConnectionSettings() interface{}`

GetConnectionSettings returns the ConnectionSettings field if non-nil, zero value otherwise.

### GetConnectionSettingsOk

`func (o *SimpleDestination) GetConnectionSettingsOk() (*interface{}, bool)`

GetConnectionSettingsOk returns a tuple with the ConnectionSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionSettings

`func (o *SimpleDestination) SetConnectionSettings(v interface{})`

SetConnectionSettings sets ConnectionSettings field to given value.

### HasConnectionSettings

`func (o *SimpleDestination) HasConnectionSettings() bool`

HasConnectionSettings returns a boolean if a field has been set.

### SetConnectionSettingsNil

`func (o *SimpleDestination) SetConnectionSettingsNil(b bool)`

 SetConnectionSettingsNil sets the value for ConnectionSettings to be an explicit nil

### UnsetConnectionSettings
`func (o *SimpleDestination) UnsetConnectionSettings()`

UnsetConnectionSettings ensures that no value is present for ConnectionSettings, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


