# AddPermissionsToUserV1Output

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Permissions** | [**[]AccessPermissionV1**](AccessPermissionV1.md) | The new permissions. | 

## Methods

### NewAddPermissionsToUserV1Output

`func NewAddPermissionsToUserV1Output(permissions []AccessPermissionV1, ) *AddPermissionsToUserV1Output`

NewAddPermissionsToUserV1Output instantiates a new AddPermissionsToUserV1Output object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddPermissionsToUserV1OutputWithDefaults

`func NewAddPermissionsToUserV1OutputWithDefaults() *AddPermissionsToUserV1Output`

NewAddPermissionsToUserV1OutputWithDefaults instantiates a new AddPermissionsToUserV1Output object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPermissions

`func (o *AddPermissionsToUserV1Output) GetPermissions() []AccessPermissionV1`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *AddPermissionsToUserV1Output) GetPermissionsOk() (*[]AccessPermissionV1, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *AddPermissionsToUserV1Output) SetPermissions(v []AccessPermissionV1)`

SetPermissions sets Permissions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


