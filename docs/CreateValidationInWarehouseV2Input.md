# CreateValidationInWarehouseV2Input

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CredentialId** | **string** | The id of the Credential to validate the connection with. The Warehouse type is derived from this Credential. | 
**Settings** | Pointer to **map[string]interface{}** | A key-value object that contains the settings for a Warehouse, derived from its associated Credential. | [optional] 

## Methods

### NewCreateValidationInWarehouseV2Input

`func NewCreateValidationInWarehouseV2Input(credentialId string, ) *CreateValidationInWarehouseV2Input`

NewCreateValidationInWarehouseV2Input instantiates a new CreateValidationInWarehouseV2Input object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateValidationInWarehouseV2InputWithDefaults

`func NewCreateValidationInWarehouseV2InputWithDefaults() *CreateValidationInWarehouseV2Input`

NewCreateValidationInWarehouseV2InputWithDefaults instantiates a new CreateValidationInWarehouseV2Input object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentialId

`func (o *CreateValidationInWarehouseV2Input) GetCredentialId() string`

GetCredentialId returns the CredentialId field if non-nil, zero value otherwise.

### GetCredentialIdOk

`func (o *CreateValidationInWarehouseV2Input) GetCredentialIdOk() (*string, bool)`

GetCredentialIdOk returns a tuple with the CredentialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialId

`func (o *CreateValidationInWarehouseV2Input) SetCredentialId(v string)`

SetCredentialId sets CredentialId field to given value.


### GetSettings

`func (o *CreateValidationInWarehouseV2Input) GetSettings() map[string]interface{}`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *CreateValidationInWarehouseV2Input) GetSettingsOk() (*map[string]interface{}, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *CreateValidationInWarehouseV2Input) SetSettings(v map[string]interface{})`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *CreateValidationInWarehouseV2Input) HasSettings() bool`

HasSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


