# UserGroupV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MemberCount** | **float32** | The number of members in the user group. | 
**Permissions** | Pointer to [**[]PermissionV1**](PermissionV1.md) | The permissions associated with the user group. | [optional] 
**Id** | **string** | The id of the user group. | 
**Name** | **string** | The name of the user group. | 

## Methods

### NewUserGroupV1

`func NewUserGroupV1(memberCount float32, id string, name string, ) *UserGroupV1`

NewUserGroupV1 instantiates a new UserGroupV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserGroupV1WithDefaults

`func NewUserGroupV1WithDefaults() *UserGroupV1`

NewUserGroupV1WithDefaults instantiates a new UserGroupV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMemberCount

`func (o *UserGroupV1) GetMemberCount() float32`

GetMemberCount returns the MemberCount field if non-nil, zero value otherwise.

### GetMemberCountOk

`func (o *UserGroupV1) GetMemberCountOk() (*float32, bool)`

GetMemberCountOk returns a tuple with the MemberCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberCount

`func (o *UserGroupV1) SetMemberCount(v float32)`

SetMemberCount sets MemberCount field to given value.


### GetPermissions

`func (o *UserGroupV1) GetPermissions() []PermissionV1`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *UserGroupV1) GetPermissionsOk() (*[]PermissionV1, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *UserGroupV1) SetPermissions(v []PermissionV1)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *UserGroupV1) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetId

`func (o *UserGroupV1) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserGroupV1) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserGroupV1) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *UserGroupV1) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserGroupV1) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserGroupV1) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


