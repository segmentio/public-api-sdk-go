# CreateCredentialV1Input

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A human-readable name for this Credential. | 
**Settings** | **map[string]interface{}** | A key-value object that contains type-specific settings for a Credential, plus a required &#x60;type&#x60; discriminator (for example &#x60;Snowflake&#x60;, &#x60;Redshift&#x60;) that determines which fields are expected. | 

## Methods

### NewCreateCredentialV1Input

`func NewCreateCredentialV1Input(name string, settings map[string]interface{}, ) *CreateCredentialV1Input`

NewCreateCredentialV1Input instantiates a new CreateCredentialV1Input object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCredentialV1InputWithDefaults

`func NewCreateCredentialV1InputWithDefaults() *CreateCredentialV1Input`

NewCreateCredentialV1InputWithDefaults instantiates a new CreateCredentialV1Input object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateCredentialV1Input) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateCredentialV1Input) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateCredentialV1Input) SetName(v string)`

SetName sets Name field to given value.


### GetSettings

`func (o *CreateCredentialV1Input) GetSettings() map[string]interface{}`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *CreateCredentialV1Input) GetSettingsOk() (*map[string]interface{}, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *CreateCredentialV1Input) SetSettings(v map[string]interface{})`

SetSettings sets Settings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


