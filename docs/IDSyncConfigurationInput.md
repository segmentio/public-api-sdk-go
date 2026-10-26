# IDSyncConfigurationInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalId** | **string** | The id type to sync. Examples: user_id, email, anonymous_id. | 
**Strategy** | **string** | The rule for selecting which identifiers to sync from a profile.  Possible values: first: Syncs only the oldest recorded value. last: Syncs only the most recently updated value. all: Syncs every value found on the profile (sends multiple events). | 
**MapTo** | Pointer to **string** | Optional destination-specific identifier to map to (for example, \&quot;Email_Address_c\&quot;). | [optional] 

## Methods

### NewIDSyncConfigurationInput

`func NewIDSyncConfigurationInput(externalId string, strategy string, ) *IDSyncConfigurationInput`

NewIDSyncConfigurationInput instantiates a new IDSyncConfigurationInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIDSyncConfigurationInputWithDefaults

`func NewIDSyncConfigurationInputWithDefaults() *IDSyncConfigurationInput`

NewIDSyncConfigurationInputWithDefaults instantiates a new IDSyncConfigurationInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalId

`func (o *IDSyncConfigurationInput) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *IDSyncConfigurationInput) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *IDSyncConfigurationInput) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### GetStrategy

`func (o *IDSyncConfigurationInput) GetStrategy() string`

GetStrategy returns the Strategy field if non-nil, zero value otherwise.

### GetStrategyOk

`func (o *IDSyncConfigurationInput) GetStrategyOk() (*string, bool)`

GetStrategyOk returns a tuple with the Strategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategy

`func (o *IDSyncConfigurationInput) SetStrategy(v string)`

SetStrategy sets Strategy field to given value.


### GetMapTo

`func (o *IDSyncConfigurationInput) GetMapTo() string`

GetMapTo returns the MapTo field if non-nil, zero value otherwise.

### GetMapToOk

`func (o *IDSyncConfigurationInput) GetMapToOk() (*string, bool)`

GetMapToOk returns a tuple with the MapTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapTo

`func (o *IDSyncConfigurationInput) SetMapTo(v string)`

SetMapTo sets MapTo field to given value.

### HasMapTo

`func (o *IDSyncConfigurationInput) HasMapTo() bool`

HasMapTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


