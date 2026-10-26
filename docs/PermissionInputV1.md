# PermissionInputV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RoleId** | **string** | The role id of this permission. | 
**Resources** | [**[]PermissionResourceV1**](PermissionResourceV1.md) | The resources to link the selected role to. | 

## Methods

### NewPermissionInputV1

`func NewPermissionInputV1(roleId string, resources []PermissionResourceV1, ) *PermissionInputV1`

NewPermissionInputV1 instantiates a new PermissionInputV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionInputV1WithDefaults

`func NewPermissionInputV1WithDefaults() *PermissionInputV1`

NewPermissionInputV1WithDefaults instantiates a new PermissionInputV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoleId

`func (o *PermissionInputV1) GetRoleId() string`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *PermissionInputV1) GetRoleIdOk() (*string, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *PermissionInputV1) SetRoleId(v string)`

SetRoleId sets RoleId field to given value.


### GetResources

`func (o *PermissionInputV1) GetResources() []PermissionResourceV1`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *PermissionInputV1) GetResourcesOk() (*[]PermissionResourceV1, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *PermissionInputV1) SetResources(v []PermissionResourceV1)`

SetResources sets Resources field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


