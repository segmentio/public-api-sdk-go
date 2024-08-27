/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 53.2.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the ReplacePermissionsForUserGroupV1Output type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReplacePermissionsForUserGroupV1Output{}

// ReplacePermissionsForUserGroupV1Output Returns the user group's permissions, including the updated permissions.
type ReplacePermissionsForUserGroupV1Output struct {
	// The updated permissions.
	Permissions []AccessPermissionV1 `json:"permissions"`
}

// NewReplacePermissionsForUserGroupV1Output instantiates a new ReplacePermissionsForUserGroupV1Output object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReplacePermissionsForUserGroupV1Output(
	permissions []AccessPermissionV1,
) *ReplacePermissionsForUserGroupV1Output {
	this := ReplacePermissionsForUserGroupV1Output{}
	this.Permissions = permissions
	return &this
}

// NewReplacePermissionsForUserGroupV1OutputWithDefaults instantiates a new ReplacePermissionsForUserGroupV1Output object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReplacePermissionsForUserGroupV1OutputWithDefaults() *ReplacePermissionsForUserGroupV1Output {
	this := ReplacePermissionsForUserGroupV1Output{}
	return &this
}

// GetPermissions returns the Permissions field value
func (o *ReplacePermissionsForUserGroupV1Output) GetPermissions() []AccessPermissionV1 {
	if o == nil {
		var ret []AccessPermissionV1
		return ret
	}

	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value
// and a boolean to check if the value has been set.
func (o *ReplacePermissionsForUserGroupV1Output) GetPermissionsOk() ([]AccessPermissionV1, bool) {
	if o == nil {
		return nil, false
	}
	return o.Permissions, true
}

// SetPermissions sets field value
func (o *ReplacePermissionsForUserGroupV1Output) SetPermissions(v []AccessPermissionV1) {
	o.Permissions = v
}

func (o ReplacePermissionsForUserGroupV1Output) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReplacePermissionsForUserGroupV1Output) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["permissions"] = o.Permissions
	return toSerialize, nil
}

type NullableReplacePermissionsForUserGroupV1Output struct {
	value *ReplacePermissionsForUserGroupV1Output
	isSet bool
}

func (v NullableReplacePermissionsForUserGroupV1Output) Get() *ReplacePermissionsForUserGroupV1Output {
	return v.value
}

func (v *NullableReplacePermissionsForUserGroupV1Output) Set(
	val *ReplacePermissionsForUserGroupV1Output,
) {
	v.value = val
	v.isSet = true
}

func (v NullableReplacePermissionsForUserGroupV1Output) IsSet() bool {
	return v.isSet
}

func (v *NullableReplacePermissionsForUserGroupV1Output) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReplacePermissionsForUserGroupV1Output(
	val *ReplacePermissionsForUserGroupV1Output,
) *NullableReplacePermissionsForUserGroupV1Output {
	return &NullableReplacePermissionsForUserGroupV1Output{value: val, isSet: true}
}

func (v NullableReplacePermissionsForUserGroupV1Output) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReplacePermissionsForUserGroupV1Output) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
