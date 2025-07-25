/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 58.13.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the AddPermissionsToUserGroupV1Input type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddPermissionsToUserGroupV1Input{}

// AddPermissionsToUserGroupV1Input Adds a list of permissions to a user group.
type AddPermissionsToUserGroupV1Input struct {
	// The permissions to add.
	Permissions []PermissionInputV1 `json:"permissions"`
}

// NewAddPermissionsToUserGroupV1Input instantiates a new AddPermissionsToUserGroupV1Input object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddPermissionsToUserGroupV1Input(
	permissions []PermissionInputV1,
) *AddPermissionsToUserGroupV1Input {
	this := AddPermissionsToUserGroupV1Input{}
	this.Permissions = permissions
	return &this
}

// NewAddPermissionsToUserGroupV1InputWithDefaults instantiates a new AddPermissionsToUserGroupV1Input object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddPermissionsToUserGroupV1InputWithDefaults() *AddPermissionsToUserGroupV1Input {
	this := AddPermissionsToUserGroupV1Input{}
	return &this
}

// GetPermissions returns the Permissions field value
func (o *AddPermissionsToUserGroupV1Input) GetPermissions() []PermissionInputV1 {
	if o == nil {
		var ret []PermissionInputV1
		return ret
	}

	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value
// and a boolean to check if the value has been set.
func (o *AddPermissionsToUserGroupV1Input) GetPermissionsOk() ([]PermissionInputV1, bool) {
	if o == nil {
		return nil, false
	}
	return o.Permissions, true
}

// SetPermissions sets field value
func (o *AddPermissionsToUserGroupV1Input) SetPermissions(v []PermissionInputV1) {
	o.Permissions = v
}

func (o AddPermissionsToUserGroupV1Input) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddPermissionsToUserGroupV1Input) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["permissions"] = o.Permissions
	return toSerialize, nil
}

type NullableAddPermissionsToUserGroupV1Input struct {
	value *AddPermissionsToUserGroupV1Input
	isSet bool
}

func (v NullableAddPermissionsToUserGroupV1Input) Get() *AddPermissionsToUserGroupV1Input {
	return v.value
}

func (v *NullableAddPermissionsToUserGroupV1Input) Set(val *AddPermissionsToUserGroupV1Input) {
	v.value = val
	v.isSet = true
}

func (v NullableAddPermissionsToUserGroupV1Input) IsSet() bool {
	return v.isSet
}

func (v *NullableAddPermissionsToUserGroupV1Input) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddPermissionsToUserGroupV1Input(
	val *AddPermissionsToUserGroupV1Input,
) *NullableAddPermissionsToUserGroupV1Input {
	return &NullableAddPermissionsToUserGroupV1Input{value: val, isSet: true}
}

func (v NullableAddPermissionsToUserGroupV1Input) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddPermissionsToUserGroupV1Input) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
