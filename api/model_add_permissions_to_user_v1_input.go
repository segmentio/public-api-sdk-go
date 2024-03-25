/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 47.0.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the AddPermissionsToUserV1Input type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddPermissionsToUserV1Input{}

// AddPermissionsToUserV1Input Adds a list of permissions to a user.
type AddPermissionsToUserV1Input struct {
	// The permissions to add.
	Permissions []PermissionInputV1 `json:"permissions"`
}

// NewAddPermissionsToUserV1Input instantiates a new AddPermissionsToUserV1Input object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddPermissionsToUserV1Input(permissions []PermissionInputV1) *AddPermissionsToUserV1Input {
	this := AddPermissionsToUserV1Input{}
	this.Permissions = permissions
	return &this
}

// NewAddPermissionsToUserV1InputWithDefaults instantiates a new AddPermissionsToUserV1Input object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddPermissionsToUserV1InputWithDefaults() *AddPermissionsToUserV1Input {
	this := AddPermissionsToUserV1Input{}
	return &this
}

// GetPermissions returns the Permissions field value
func (o *AddPermissionsToUserV1Input) GetPermissions() []PermissionInputV1 {
	if o == nil {
		var ret []PermissionInputV1
		return ret
	}

	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value
// and a boolean to check if the value has been set.
func (o *AddPermissionsToUserV1Input) GetPermissionsOk() ([]PermissionInputV1, bool) {
	if o == nil {
		return nil, false
	}
	return o.Permissions, true
}

// SetPermissions sets field value
func (o *AddPermissionsToUserV1Input) SetPermissions(v []PermissionInputV1) {
	o.Permissions = v
}

func (o AddPermissionsToUserV1Input) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddPermissionsToUserV1Input) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["permissions"] = o.Permissions
	return toSerialize, nil
}

type NullableAddPermissionsToUserV1Input struct {
	value *AddPermissionsToUserV1Input
	isSet bool
}

func (v NullableAddPermissionsToUserV1Input) Get() *AddPermissionsToUserV1Input {
	return v.value
}

func (v *NullableAddPermissionsToUserV1Input) Set(val *AddPermissionsToUserV1Input) {
	v.value = val
	v.isSet = true
}

func (v NullableAddPermissionsToUserV1Input) IsSet() bool {
	return v.isSet
}

func (v *NullableAddPermissionsToUserV1Input) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddPermissionsToUserV1Input(
	val *AddPermissionsToUserV1Input,
) *NullableAddPermissionsToUserV1Input {
	return &NullableAddPermissionsToUserV1Input{value: val, isSet: true}
}

func (v NullableAddPermissionsToUserV1Input) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddPermissionsToUserV1Input) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
