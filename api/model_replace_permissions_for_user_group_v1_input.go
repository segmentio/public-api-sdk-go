/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 35.0.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// ReplacePermissionsForUserGroupV1Input Updates the list of permissions for a user group.
type ReplacePermissionsForUserGroupV1Input struct {
	// The permissions to replace with.
	Permissions []PermissionInputV1 `json:"permissions"`
}

// NewReplacePermissionsForUserGroupV1Input instantiates a new ReplacePermissionsForUserGroupV1Input object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReplacePermissionsForUserGroupV1Input(
	permissions []PermissionInputV1,
) *ReplacePermissionsForUserGroupV1Input {
	this := ReplacePermissionsForUserGroupV1Input{}
	this.Permissions = permissions
	return &this
}

// NewReplacePermissionsForUserGroupV1InputWithDefaults instantiates a new ReplacePermissionsForUserGroupV1Input object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReplacePermissionsForUserGroupV1InputWithDefaults() *ReplacePermissionsForUserGroupV1Input {
	this := ReplacePermissionsForUserGroupV1Input{}
	return &this
}

// GetPermissions returns the Permissions field value
func (o *ReplacePermissionsForUserGroupV1Input) GetPermissions() []PermissionInputV1 {
	if o == nil {
		var ret []PermissionInputV1
		return ret
	}

	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value
// and a boolean to check if the value has been set.
func (o *ReplacePermissionsForUserGroupV1Input) GetPermissionsOk() ([]PermissionInputV1, bool) {
	if o == nil {
		return nil, false
	}
	return o.Permissions, true
}

// SetPermissions sets field value
func (o *ReplacePermissionsForUserGroupV1Input) SetPermissions(v []PermissionInputV1) {
	o.Permissions = v
}

func (o ReplacePermissionsForUserGroupV1Input) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["permissions"] = o.Permissions
	}
	return json.Marshal(toSerialize)
}

type NullableReplacePermissionsForUserGroupV1Input struct {
	value *ReplacePermissionsForUserGroupV1Input
	isSet bool
}

func (v NullableReplacePermissionsForUserGroupV1Input) Get() *ReplacePermissionsForUserGroupV1Input {
	return v.value
}

func (v *NullableReplacePermissionsForUserGroupV1Input) Set(
	val *ReplacePermissionsForUserGroupV1Input,
) {
	v.value = val
	v.isSet = true
}

func (v NullableReplacePermissionsForUserGroupV1Input) IsSet() bool {
	return v.isSet
}

func (v *NullableReplacePermissionsForUserGroupV1Input) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReplacePermissionsForUserGroupV1Input(
	val *ReplacePermissionsForUserGroupV1Input,
) *NullableReplacePermissionsForUserGroupV1Input {
	return &NullableReplacePermissionsForUserGroupV1Input{value: val, isSet: true}
}

func (v NullableReplacePermissionsForUserGroupV1Input) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReplacePermissionsForUserGroupV1Input) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
