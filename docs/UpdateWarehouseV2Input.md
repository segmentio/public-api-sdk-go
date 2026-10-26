# UpdateWarehouseV2Input

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | An optional human-readable name to associate with this Warehouse. | [optional] 
**Enabled** | Pointer to **bool** | Enable to allow this Warehouse to receive data. | [optional] 
**CredentialId** | Pointer to **string** | The id of a Credential to connect this Warehouse to. Swapping the Credential re-derives this Warehouse&#39;s connection settings and re-validates the connection. The Credential must be the same type as the Warehouse&#39;s existing type — a Warehouse&#39;s type cannot be changed after creation, so a mismatched Credential is rejected rather than silently ignored. | [optional] 

## Methods

### NewUpdateWarehouseV2Input

`func NewUpdateWarehouseV2Input() *UpdateWarehouseV2Input`

NewUpdateWarehouseV2Input instantiates a new UpdateWarehouseV2Input object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateWarehouseV2InputWithDefaults

`func NewUpdateWarehouseV2InputWithDefaults() *UpdateWarehouseV2Input`

NewUpdateWarehouseV2InputWithDefaults instantiates a new UpdateWarehouseV2Input object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateWarehouseV2Input) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateWarehouseV2Input) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateWarehouseV2Input) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateWarehouseV2Input) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEnabled

`func (o *UpdateWarehouseV2Input) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateWarehouseV2Input) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateWarehouseV2Input) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateWarehouseV2Input) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetCredentialId

`func (o *UpdateWarehouseV2Input) GetCredentialId() string`

GetCredentialId returns the CredentialId field if non-nil, zero value otherwise.

### GetCredentialIdOk

`func (o *UpdateWarehouseV2Input) GetCredentialIdOk() (*string, bool)`

GetCredentialIdOk returns a tuple with the CredentialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialId

`func (o *UpdateWarehouseV2Input) SetCredentialId(v string)`

SetCredentialId sets CredentialId field to given value.

### HasCredentialId

`func (o *UpdateWarehouseV2Input) HasCredentialId() bool`

HasCredentialId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


