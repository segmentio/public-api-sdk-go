/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API. 

API version: 32.0.8
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// ReplacePermissionsForUserV1Output Returns the user's permissions, including the updated permissions.
type ReplacePermissionsForUserV1Output struct {
	// The updated permissions.
	Permissions []AccessPermissionV1 `json:"permissions"`
}

// NewReplacePermissionsForUserV1Output instantiates a new ReplacePermissionsForUserV1Output object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReplacePermissionsForUserV1Output(permissions []AccessPermissionV1) *ReplacePermissionsForUserV1Output {
	this := ReplacePermissionsForUserV1Output{}
	this.Permissions = permissions
	return &this
}

// NewReplacePermissionsForUserV1OutputWithDefaults instantiates a new ReplacePermissionsForUserV1Output object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReplacePermissionsForUserV1OutputWithDefaults() *ReplacePermissionsForUserV1Output {
	this := ReplacePermissionsForUserV1Output{}
	return &this
}

// GetPermissions returns the Permissions field value
func (o *ReplacePermissionsForUserV1Output) GetPermissions() []AccessPermissionV1 {
	if o == nil {
		var ret []AccessPermissionV1
		return ret
	}

	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value
// and a boolean to check if the value has been set.
func (o *ReplacePermissionsForUserV1Output) GetPermissionsOk() ([]AccessPermissionV1, bool) {
	if o == nil {
		return nil, false
	}
	return o.Permissions, true
}

// SetPermissions sets field value
func (o *ReplacePermissionsForUserV1Output) SetPermissions(v []AccessPermissionV1) {
	o.Permissions = v
}

func (o ReplacePermissionsForUserV1Output) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["permissions"] = o.Permissions
	}
	return json.Marshal(toSerialize)
}

type NullableReplacePermissionsForUserV1Output struct {
	value *ReplacePermissionsForUserV1Output
	isSet bool
}

func (v NullableReplacePermissionsForUserV1Output) Get() *ReplacePermissionsForUserV1Output {
	return v.value
}

func (v *NullableReplacePermissionsForUserV1Output) Set(val *ReplacePermissionsForUserV1Output) {
	v.value = val
	v.isSet = true
}

func (v NullableReplacePermissionsForUserV1Output) IsSet() bool {
	return v.isSet
}

func (v *NullableReplacePermissionsForUserV1Output) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReplacePermissionsForUserV1Output(val *ReplacePermissionsForUserV1Output) *NullableReplacePermissionsForUserV1Output {
	return &NullableReplacePermissionsForUserV1Output{value: val, isSet: true}
}

func (v NullableReplacePermissionsForUserV1Output) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReplacePermissionsForUserV1Output) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


