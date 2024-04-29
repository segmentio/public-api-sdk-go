/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 50.0.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the UpdateUserGroupV1Output type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateUserGroupV1Output{}

// UpdateUserGroupV1Output Returns the updated user group.
type UpdateUserGroupV1Output struct {
	UserGroup UserGroupV1 `json:"userGroup"`
}

// NewUpdateUserGroupV1Output instantiates a new UpdateUserGroupV1Output object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateUserGroupV1Output(userGroup UserGroupV1) *UpdateUserGroupV1Output {
	this := UpdateUserGroupV1Output{}
	this.UserGroup = userGroup
	return &this
}

// NewUpdateUserGroupV1OutputWithDefaults instantiates a new UpdateUserGroupV1Output object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateUserGroupV1OutputWithDefaults() *UpdateUserGroupV1Output {
	this := UpdateUserGroupV1Output{}
	return &this
}

// GetUserGroup returns the UserGroup field value
func (o *UpdateUserGroupV1Output) GetUserGroup() UserGroupV1 {
	if o == nil {
		var ret UserGroupV1
		return ret
	}

	return o.UserGroup
}

// GetUserGroupOk returns a tuple with the UserGroup field value
// and a boolean to check if the value has been set.
func (o *UpdateUserGroupV1Output) GetUserGroupOk() (*UserGroupV1, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserGroup, true
}

// SetUserGroup sets field value
func (o *UpdateUserGroupV1Output) SetUserGroup(v UserGroupV1) {
	o.UserGroup = v
}

func (o UpdateUserGroupV1Output) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateUserGroupV1Output) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["userGroup"] = o.UserGroup
	return toSerialize, nil
}

type NullableUpdateUserGroupV1Output struct {
	value *UpdateUserGroupV1Output
	isSet bool
}

func (v NullableUpdateUserGroupV1Output) Get() *UpdateUserGroupV1Output {
	return v.value
}

func (v *NullableUpdateUserGroupV1Output) Set(val *UpdateUserGroupV1Output) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateUserGroupV1Output) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateUserGroupV1Output) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateUserGroupV1Output(
	val *UpdateUserGroupV1Output,
) *NullableUpdateUserGroupV1Output {
	return &NullableUpdateUserGroupV1Output{value: val, isSet: true}
}

func (v NullableUpdateUserGroupV1Output) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateUserGroupV1Output) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
