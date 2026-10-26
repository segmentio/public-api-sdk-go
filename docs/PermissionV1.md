# PermissionV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RoleName** | **string** | The name of the role associated with this permission. | 
**RoleId** | **string** | The id of the role associated with this permission. | 
**Resources** | [**[]PermissionResourceV1**](PermissionResourceV1.md) | The resources associated with this permission. | 
**Labels** | Pointer to [**[]AllowedLabelBeta**](AllowedLabelBeta.md) | The labels to attach to this permission. | [optional] 

## Methods

### NewPermissionV1

`func NewPermissionV1(roleName string, roleId string, resources []PermissionResourceV1, ) *PermissionV1`

NewPermissionV1 instantiates a new PermissionV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionV1WithDefaults

`func NewPermissionV1WithDefaults() *PermissionV1`

NewPermissionV1WithDefaults instantiates a new PermissionV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoleName

`func (o *PermissionV1) GetRoleName() string`

GetRoleName returns the RoleName field if non-nil, zero value otherwise.

### GetRoleNameOk

`func (o *PermissionV1) GetRoleNameOk() (*string, bool)`

GetRoleNameOk returns a tuple with the RoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleName

`func (o *PermissionV1) SetRoleName(v string)`

SetRoleName sets RoleName field to given value.


### GetRoleId

`func (o *PermissionV1) GetRoleId() string`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *PermissionV1) GetRoleIdOk() (*string, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *PermissionV1) SetRoleId(v string)`

SetRoleId sets RoleId field to given value.


### GetResources

`func (o *PermissionV1) GetResources() []PermissionResourceV1`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *PermissionV1) GetResourcesOk() (*[]PermissionResourceV1, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *PermissionV1) SetResources(v []PermissionResourceV1)`

SetResources sets Resources field to given value.


### GetLabels

`func (o *PermissionV1) GetLabels() []AllowedLabelBeta`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *PermissionV1) GetLabelsOk() (*[]AllowedLabelBeta, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *PermissionV1) SetLabels(v []AllowedLabelBeta)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *PermissionV1) HasLabels() bool`

HasLabels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


