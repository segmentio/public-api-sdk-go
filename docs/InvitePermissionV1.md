# InvitePermissionV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RoleId** | **string** | The id of the role. | 
**Resources** | Pointer to [**[]ResourceV1**](ResourceV1.md) | The resources to grant the invited users access to. | [optional] 
**Labels** | Pointer to [**[]AllowedLabelBeta**](AllowedLabelBeta.md) | The labels that determine which resources to grant users access to. | [optional] 

## Methods

### NewInvitePermissionV1

`func NewInvitePermissionV1(roleId string, ) *InvitePermissionV1`

NewInvitePermissionV1 instantiates a new InvitePermissionV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvitePermissionV1WithDefaults

`func NewInvitePermissionV1WithDefaults() *InvitePermissionV1`

NewInvitePermissionV1WithDefaults instantiates a new InvitePermissionV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoleId

`func (o *InvitePermissionV1) GetRoleId() string`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *InvitePermissionV1) GetRoleIdOk() (*string, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *InvitePermissionV1) SetRoleId(v string)`

SetRoleId sets RoleId field to given value.


### GetResources

`func (o *InvitePermissionV1) GetResources() []ResourceV1`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *InvitePermissionV1) GetResourcesOk() (*[]ResourceV1, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *InvitePermissionV1) SetResources(v []ResourceV1)`

SetResources sets Resources field to given value.

### HasResources

`func (o *InvitePermissionV1) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetLabels

`func (o *InvitePermissionV1) GetLabels() []AllowedLabelBeta`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *InvitePermissionV1) GetLabelsOk() (*[]AllowedLabelBeta, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *InvitePermissionV1) SetLabels(v []AllowedLabelBeta)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *InvitePermissionV1) HasLabels() bool`

HasLabels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


