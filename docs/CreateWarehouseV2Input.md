# CreateWarehouseV2Input

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CredentialId** | **string** | The id of the Credential to connect this Warehouse to. The Warehouse&#39;s type and connection settings are derived from this Credential. | 
**Name** | Pointer to **string** | An optional human-readable name for this Warehouse. | [optional] 
**Enabled** | Pointer to **bool** | Enable to allow this Warehouse to receive data. Defaults to true. | [optional] 

## Methods

### NewCreateWarehouseV2Input

`func NewCreateWarehouseV2Input(credentialId string, ) *CreateWarehouseV2Input`

NewCreateWarehouseV2Input instantiates a new CreateWarehouseV2Input object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateWarehouseV2InputWithDefaults

`func NewCreateWarehouseV2InputWithDefaults() *CreateWarehouseV2Input`

NewCreateWarehouseV2InputWithDefaults instantiates a new CreateWarehouseV2Input object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentialId

`func (o *CreateWarehouseV2Input) GetCredentialId() string`

GetCredentialId returns the CredentialId field if non-nil, zero value otherwise.

### GetCredentialIdOk

`func (o *CreateWarehouseV2Input) GetCredentialIdOk() (*string, bool)`

GetCredentialIdOk returns a tuple with the CredentialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialId

`func (o *CreateWarehouseV2Input) SetCredentialId(v string)`

SetCredentialId sets CredentialId field to given value.


### GetName

`func (o *CreateWarehouseV2Input) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateWarehouseV2Input) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateWarehouseV2Input) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateWarehouseV2Input) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEnabled

`func (o *CreateWarehouseV2Input) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CreateWarehouseV2Input) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CreateWarehouseV2Input) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CreateWarehouseV2Input) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


