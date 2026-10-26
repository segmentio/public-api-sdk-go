# AccessPermissionV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RoleId** | **string** | The id of the role that applies to this permission. | 
**RoleName** | **string** | The name of the role that applies to this permission. | 
**Resources** | [**[]PermissionResourceV1**](PermissionResourceV1.md) | The resources included with this permission. | 

## Methods

### NewAccessPermissionV1

`func NewAccessPermissionV1(roleId string, roleName string, resources []PermissionResourceV1, ) *AccessPermissionV1`

NewAccessPermissionV1 instantiates a new AccessPermissionV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessPermissionV1WithDefaults

`func NewAccessPermissionV1WithDefaults() *AccessPermissionV1`

NewAccessPermissionV1WithDefaults instantiates a new AccessPermissionV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoleId

`func (o *AccessPermissionV1) GetRoleId() string`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *AccessPermissionV1) GetRoleIdOk() (*string, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *AccessPermissionV1) SetRoleId(v string)`

SetRoleId sets RoleId field to given value.


### GetRoleName

`func (o *AccessPermissionV1) GetRoleName() string`

GetRoleName returns the RoleName field if non-nil, zero value otherwise.

### GetRoleNameOk

`func (o *AccessPermissionV1) GetRoleNameOk() (*string, bool)`

GetRoleNameOk returns a tuple with the RoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleName

`func (o *AccessPermissionV1) SetRoleName(v string)`

SetRoleName sets RoleName field to given value.


### GetResources

`func (o *AccessPermissionV1) GetResources() []PermissionResourceV1`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *AccessPermissionV1) GetResourcesOk() (*[]PermissionResourceV1, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *AccessPermissionV1) SetResources(v []PermissionResourceV1)`

SetResources sets Resources field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


