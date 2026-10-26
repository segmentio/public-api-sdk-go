# WarehouseV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The id of the Warehouse. | 
**WorkspaceId** | **string** | The id of the Workspace that owns this Warehouse. | 
**Name** | Pointer to **string** | A human-readable name for this Warehouse. | [optional] 
**Enabled** | **bool** | When set to true, this Warehouse receives data. | 
**CredentialId** | **string** | The id of the Credential this Warehouse is connected to. | 
**Settings** | **map[string]interface{}** | A key-value object that contains the settings for a Warehouse, derived from its associated Credential. | 
**CreatedAt** | **NullableString** | The time this Warehouse was created. | 
**UpdatedAt** | **NullableString** | The time this Warehouse was last updated. | 

## Methods

### NewWarehouseV2

`func NewWarehouseV2(id string, workspaceId string, enabled bool, credentialId string, settings map[string]interface{}, createdAt NullableString, updatedAt NullableString, ) *WarehouseV2`

NewWarehouseV2 instantiates a new WarehouseV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWarehouseV2WithDefaults

`func NewWarehouseV2WithDefaults() *WarehouseV2`

NewWarehouseV2WithDefaults instantiates a new WarehouseV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WarehouseV2) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WarehouseV2) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WarehouseV2) SetId(v string)`

SetId sets Id field to given value.


### GetWorkspaceId

`func (o *WarehouseV2) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *WarehouseV2) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *WarehouseV2) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.


### GetName

`func (o *WarehouseV2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WarehouseV2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WarehouseV2) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WarehouseV2) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEnabled

`func (o *WarehouseV2) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *WarehouseV2) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *WarehouseV2) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetCredentialId

`func (o *WarehouseV2) GetCredentialId() string`

GetCredentialId returns the CredentialId field if non-nil, zero value otherwise.

### GetCredentialIdOk

`func (o *WarehouseV2) GetCredentialIdOk() (*string, bool)`

GetCredentialIdOk returns a tuple with the CredentialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialId

`func (o *WarehouseV2) SetCredentialId(v string)`

SetCredentialId sets CredentialId field to given value.


### GetSettings

`func (o *WarehouseV2) GetSettings() map[string]interface{}`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *WarehouseV2) GetSettingsOk() (*map[string]interface{}, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *WarehouseV2) SetSettings(v map[string]interface{})`

SetSettings sets Settings field to given value.


### GetCreatedAt

`func (o *WarehouseV2) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WarehouseV2) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WarehouseV2) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *WarehouseV2) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *WarehouseV2) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *WarehouseV2) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *WarehouseV2) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *WarehouseV2) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### SetUpdatedAtNil

`func (o *WarehouseV2) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *WarehouseV2) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


