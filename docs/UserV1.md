# UserV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier of this user.  Config API note: analogous to &#x60;name&#x60;. | 
**Name** | **string** | The human-readable name of this user. | 
**Email** | **string** | The email address associated with this user. | 
**Permissions** | Pointer to [**[]PermissionV1**](PermissionV1.md) | The permissions associated with this user. | [optional] 

## Methods

### NewUserV1

`func NewUserV1(id string, name string, email string, ) *UserV1`

NewUserV1 instantiates a new UserV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserV1WithDefaults

`func NewUserV1WithDefaults() *UserV1`

NewUserV1WithDefaults instantiates a new UserV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserV1) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserV1) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserV1) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *UserV1) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserV1) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserV1) SetName(v string)`

SetName sets Name field to given value.


### GetEmail

`func (o *UserV1) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserV1) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserV1) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetPermissions

`func (o *UserV1) GetPermissions() []PermissionV1`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *UserV1) GetPermissionsOk() (*[]PermissionV1, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *UserV1) SetPermissions(v []PermissionV1)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *UserV1) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


