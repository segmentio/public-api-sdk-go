# InviteV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | The invited user&#39;s email to attach the permissions to. | 
**Permissions** | Pointer to [**[]InvitePermissionV1**](InvitePermissionV1.md) | The permissions to attach to the invited user. | [optional] 

## Methods

### NewInviteV1

`func NewInviteV1(email string, ) *InviteV1`

NewInviteV1 instantiates a new InviteV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInviteV1WithDefaults

`func NewInviteV1WithDefaults() *InviteV1`

NewInviteV1WithDefaults instantiates a new InviteV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *InviteV1) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *InviteV1) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *InviteV1) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetPermissions

`func (o *InviteV1) GetPermissions() []InvitePermissionV1`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *InviteV1) GetPermissionsOk() (*[]InvitePermissionV1, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *InviteV1) SetPermissions(v []InvitePermissionV1)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *InviteV1) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


