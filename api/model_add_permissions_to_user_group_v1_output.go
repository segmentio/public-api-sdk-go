/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 36.1.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// AddPermissionsToUserGroupV1Output Returns the group's permissions, including the added permissions.
type AddPermissionsToUserGroupV1Output struct {
	// The updated set of permissions assigned to the user group.
	Permissions []AccessPermissionV1 `json:"permissions"`
}

// NewAddPermissionsToUserGroupV1Output instantiates a new AddPermissionsToUserGroupV1Output object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddPermissionsToUserGroupV1Output(
	permissions []AccessPermissionV1,
) *AddPermissionsToUserGroupV1Output {
	this := AddPermissionsToUserGroupV1Output{}
	this.Permissions = permissions
	return &this
}

// NewAddPermissionsToUserGroupV1OutputWithDefaults instantiates a new AddPermissionsToUserGroupV1Output object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddPermissionsToUserGroupV1OutputWithDefaults() *AddPermissionsToUserGroupV1Output {
	this := AddPermissionsToUserGroupV1Output{}
	return &this
}

// GetPermissions returns the Permissions field value
func (o *AddPermissionsToUserGroupV1Output) GetPermissions() []AccessPermissionV1 {
	if o == nil {
		var ret []AccessPermissionV1
		return ret
	}

	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value
// and a boolean to check if the value has been set.
func (o *AddPermissionsToUserGroupV1Output) GetPermissionsOk() ([]AccessPermissionV1, bool) {
	if o == nil {
		return nil, false
	}
	return o.Permissions, true
}

// SetPermissions sets field value
func (o *AddPermissionsToUserGroupV1Output) SetPermissions(v []AccessPermissionV1) {
	o.Permissions = v
}

func (o AddPermissionsToUserGroupV1Output) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["permissions"] = o.Permissions
	}
	return json.Marshal(toSerialize)
}

type NullableAddPermissionsToUserGroupV1Output struct {
	value *AddPermissionsToUserGroupV1Output
	isSet bool
}

func (v NullableAddPermissionsToUserGroupV1Output) Get() *AddPermissionsToUserGroupV1Output {
	return v.value
}

func (v *NullableAddPermissionsToUserGroupV1Output) Set(val *AddPermissionsToUserGroupV1Output) {
	v.value = val
	v.isSet = true
}

func (v NullableAddPermissionsToUserGroupV1Output) IsSet() bool {
	return v.isSet
}

func (v *NullableAddPermissionsToUserGroupV1Output) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddPermissionsToUserGroupV1Output(
	val *AddPermissionsToUserGroupV1Output,
) *NullableAddPermissionsToUserGroupV1Output {
	return &NullableAddPermissionsToUserGroupV1Output{value: val, isSet: true}
}

func (v NullableAddPermissionsToUserGroupV1Output) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddPermissionsToUserGroupV1Output) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
