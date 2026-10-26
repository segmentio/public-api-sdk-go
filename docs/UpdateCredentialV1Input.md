# UpdateCredentialV1Input

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | An optional human-readable name to associate with this Credential. | [optional] 
**Settings** | Pointer to **map[string]interface{}** | A key-value object that contains type-specific settings for a Credential, plus a required &#x60;type&#x60; discriminator (for example &#x60;Snowflake&#x60;, &#x60;Redshift&#x60;) that determines which fields are expected. | [optional] 

## Methods

### NewUpdateCredentialV1Input

`func NewUpdateCredentialV1Input() *UpdateCredentialV1Input`

NewUpdateCredentialV1Input instantiates a new UpdateCredentialV1Input object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCredentialV1InputWithDefaults

`func NewUpdateCredentialV1InputWithDefaults() *UpdateCredentialV1Input`

NewUpdateCredentialV1InputWithDefaults instantiates a new UpdateCredentialV1Input object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateCredentialV1Input) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateCredentialV1Input) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateCredentialV1Input) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateCredentialV1Input) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSettings

`func (o *UpdateCredentialV1Input) GetSettings() map[string]interface{}`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *UpdateCredentialV1Input) GetSettingsOk() (*map[string]interface{}, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *UpdateCredentialV1Input) SetSettings(v map[string]interface{})`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *UpdateCredentialV1Input) HasSettings() bool`

HasSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


