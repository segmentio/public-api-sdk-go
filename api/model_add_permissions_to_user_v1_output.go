/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 49.1.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the AddPermissionsToUserV1Output type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddPermissionsToUserV1Output{}

// AddPermissionsToUserV1Output Returns the user's permissions, including the added permissions.
type AddPermissionsToUserV1Output struct {
	// The new permissions.
	Permissions []AccessPermissionV1 `json:"permissions"`
}

// NewAddPermissionsToUserV1Output instantiates a new AddPermissionsToUserV1Output object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddPermissionsToUserV1Output(
	permissions []AccessPermissionV1,
) *AddPermissionsToUserV1Output {
	this := AddPermissionsToUserV1Output{}
	this.Permissions = permissions
	return &this
}

// NewAddPermissionsToUserV1OutputWithDefaults instantiates a new AddPermissionsToUserV1Output object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddPermissionsToUserV1OutputWithDefaults() *AddPermissionsToUserV1Output {
	this := AddPermissionsToUserV1Output{}
	return &this
}

// GetPermissions returns the Permissions field value
func (o *AddPermissionsToUserV1Output) GetPermissions() []AccessPermissionV1 {
	if o == nil {
		var ret []AccessPermissionV1
		return ret
	}

	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value
// and a boolean to check if the value has been set.
func (o *AddPermissionsToUserV1Output) GetPermissionsOk() ([]AccessPermissionV1, bool) {
	if o == nil {
		return nil, false
	}
	return o.Permissions, true
}

// SetPermissions sets field value
func (o *AddPermissionsToUserV1Output) SetPermissions(v []AccessPermissionV1) {
	o.Permissions = v
}

func (o AddPermissionsToUserV1Output) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddPermissionsToUserV1Output) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["permissions"] = o.Permissions
	return toSerialize, nil
}

type NullableAddPermissionsToUserV1Output struct {
	value *AddPermissionsToUserV1Output
	isSet bool
}

func (v NullableAddPermissionsToUserV1Output) Get() *AddPermissionsToUserV1Output {
	return v.value
}

func (v *NullableAddPermissionsToUserV1Output) Set(val *AddPermissionsToUserV1Output) {
	v.value = val
	v.isSet = true
}

func (v NullableAddPermissionsToUserV1Output) IsSet() bool {
	return v.isSet
}

func (v *NullableAddPermissionsToUserV1Output) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddPermissionsToUserV1Output(
	val *AddPermissionsToUserV1Output,
) *NullableAddPermissionsToUserV1Output {
	return &NullableAddPermissionsToUserV1Output{value: val, isSet: true}
}

func (v NullableAddPermissionsToUserV1Output) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddPermissionsToUserV1Output) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
