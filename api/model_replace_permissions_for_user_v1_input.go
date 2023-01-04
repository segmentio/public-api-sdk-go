/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 33.0.4
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// ReplacePermissionsForUserV1Input Updates the list of permissions for a user.
type ReplacePermissionsForUserV1Input struct {
	// The permissions to add.
	Permissions []PermissionInputV1 `json:"permissions"`
}

// NewReplacePermissionsForUserV1Input instantiates a new ReplacePermissionsForUserV1Input object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReplacePermissionsForUserV1Input(
	permissions []PermissionInputV1,
) *ReplacePermissionsForUserV1Input {
	this := ReplacePermissionsForUserV1Input{}
	this.Permissions = permissions
	return &this
}

// NewReplacePermissionsForUserV1InputWithDefaults instantiates a new ReplacePermissionsForUserV1Input object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReplacePermissionsForUserV1InputWithDefaults() *ReplacePermissionsForUserV1Input {
	this := ReplacePermissionsForUserV1Input{}
	return &this
}

// GetPermissions returns the Permissions field value
func (o *ReplacePermissionsForUserV1Input) GetPermissions() []PermissionInputV1 {
	if o == nil {
		var ret []PermissionInputV1
		return ret
	}

	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value
// and a boolean to check if the value has been set.
func (o *ReplacePermissionsForUserV1Input) GetPermissionsOk() ([]PermissionInputV1, bool) {
	if o == nil {
		return nil, false
	}
	return o.Permissions, true
}

// SetPermissions sets field value
func (o *ReplacePermissionsForUserV1Input) SetPermissions(v []PermissionInputV1) {
	o.Permissions = v
}

func (o ReplacePermissionsForUserV1Input) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["permissions"] = o.Permissions
	}
	return json.Marshal(toSerialize)
}

type NullableReplacePermissionsForUserV1Input struct {
	value *ReplacePermissionsForUserV1Input
	isSet bool
}

func (v NullableReplacePermissionsForUserV1Input) Get() *ReplacePermissionsForUserV1Input {
	return v.value
}

func (v *NullableReplacePermissionsForUserV1Input) Set(val *ReplacePermissionsForUserV1Input) {
	v.value = val
	v.isSet = true
}

func (v NullableReplacePermissionsForUserV1Input) IsSet() bool {
	return v.isSet
}

func (v *NullableReplacePermissionsForUserV1Input) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReplacePermissionsForUserV1Input(
	val *ReplacePermissionsForUserV1Input,
) *NullableReplacePermissionsForUserV1Input {
	return &NullableReplacePermissionsForUserV1Input{value: val, isSet: true}
}

func (v NullableReplacePermissionsForUserV1Input) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReplacePermissionsForUserV1Input) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
