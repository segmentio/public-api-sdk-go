# AddPermissionsToUserGroupV1Output

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Permissions** | [**[]AccessPermissionV1**](AccessPermissionV1.md) | The updated set of permissions assigned to the user group. | 

## Methods

### NewAddPermissionsToUserGroupV1Output

`func NewAddPermissionsToUserGroupV1Output(permissions []AccessPermissionV1, ) *AddPermissionsToUserGroupV1Output`

NewAddPermissionsToUserGroupV1Output instantiates a new AddPermissionsToUserGroupV1Output object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddPermissionsToUserGroupV1OutputWithDefaults

`func NewAddPermissionsToUserGroupV1OutputWithDefaults() *AddPermissionsToUserGroupV1Output`

NewAddPermissionsToUserGroupV1OutputWithDefaults instantiates a new AddPermissionsToUserGroupV1Output object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPermissions

`func (o *AddPermissionsToUserGroupV1Output) GetPermissions() []AccessPermissionV1`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *AddPermissionsToUserGroupV1Output) GetPermissionsOk() (*[]AccessPermissionV1, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *AddPermissionsToUserGroupV1Output) SetPermissions(v []AccessPermissionV1)`

SetPermissions sets Permissions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


