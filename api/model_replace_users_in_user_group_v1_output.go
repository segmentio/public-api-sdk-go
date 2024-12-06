/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 57.1.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the ReplaceUsersInUserGroupV1Output type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReplaceUsersInUserGroupV1Output{}

// ReplaceUsersInUserGroupV1Output Returns the updated user group.
type ReplaceUsersInUserGroupV1Output struct {
	UserGroup UserGroupV1 `json:"userGroup"`
}

// NewReplaceUsersInUserGroupV1Output instantiates a new ReplaceUsersInUserGroupV1Output object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReplaceUsersInUserGroupV1Output(userGroup UserGroupV1) *ReplaceUsersInUserGroupV1Output {
	this := ReplaceUsersInUserGroupV1Output{}
	this.UserGroup = userGroup
	return &this
}

// NewReplaceUsersInUserGroupV1OutputWithDefaults instantiates a new ReplaceUsersInUserGroupV1Output object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReplaceUsersInUserGroupV1OutputWithDefaults() *ReplaceUsersInUserGroupV1Output {
	this := ReplaceUsersInUserGroupV1Output{}
	return &this
}

// GetUserGroup returns the UserGroup field value
func (o *ReplaceUsersInUserGroupV1Output) GetUserGroup() UserGroupV1 {
	if o == nil {
		var ret UserGroupV1
		return ret
	}

	return o.UserGroup
}

// GetUserGroupOk returns a tuple with the UserGroup field value
// and a boolean to check if the value has been set.
func (o *ReplaceUsersInUserGroupV1Output) GetUserGroupOk() (*UserGroupV1, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserGroup, true
}

// SetUserGroup sets field value
func (o *ReplaceUsersInUserGroupV1Output) SetUserGroup(v UserGroupV1) {
	o.UserGroup = v
}

func (o ReplaceUsersInUserGroupV1Output) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReplaceUsersInUserGroupV1Output) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["userGroup"] = o.UserGroup
	return toSerialize, nil
}

type NullableReplaceUsersInUserGroupV1Output struct {
	value *ReplaceUsersInUserGroupV1Output
	isSet bool
}

func (v NullableReplaceUsersInUserGroupV1Output) Get() *ReplaceUsersInUserGroupV1Output {
	return v.value
}

func (v *NullableReplaceUsersInUserGroupV1Output) Set(val *ReplaceUsersInUserGroupV1Output) {
	v.value = val
	v.isSet = true
}

func (v NullableReplaceUsersInUserGroupV1Output) IsSet() bool {
	return v.isSet
}

func (v *NullableReplaceUsersInUserGroupV1Output) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReplaceUsersInUserGroupV1Output(
	val *ReplaceUsersInUserGroupV1Output,
) *NullableReplaceUsersInUserGroupV1Output {
	return &NullableReplaceUsersInUserGroupV1Output{value: val, isSet: true}
}

func (v NullableReplaceUsersInUserGroupV1Output) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReplaceUsersInUserGroupV1Output) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
