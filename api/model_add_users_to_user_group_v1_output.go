/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 54.0.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the AddUsersToUserGroupV1Output type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddUsersToUserGroupV1Output{}

// AddUsersToUserGroupV1Output Returns the updated user group.
type AddUsersToUserGroupV1Output struct {
	UserGroup UserGroupV1 `json:"userGroup"`
}

// NewAddUsersToUserGroupV1Output instantiates a new AddUsersToUserGroupV1Output object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddUsersToUserGroupV1Output(userGroup UserGroupV1) *AddUsersToUserGroupV1Output {
	this := AddUsersToUserGroupV1Output{}
	this.UserGroup = userGroup
	return &this
}

// NewAddUsersToUserGroupV1OutputWithDefaults instantiates a new AddUsersToUserGroupV1Output object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddUsersToUserGroupV1OutputWithDefaults() *AddUsersToUserGroupV1Output {
	this := AddUsersToUserGroupV1Output{}
	return &this
}

// GetUserGroup returns the UserGroup field value
func (o *AddUsersToUserGroupV1Output) GetUserGroup() UserGroupV1 {
	if o == nil {
		var ret UserGroupV1
		return ret
	}

	return o.UserGroup
}

// GetUserGroupOk returns a tuple with the UserGroup field value
// and a boolean to check if the value has been set.
func (o *AddUsersToUserGroupV1Output) GetUserGroupOk() (*UserGroupV1, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserGroup, true
}

// SetUserGroup sets field value
func (o *AddUsersToUserGroupV1Output) SetUserGroup(v UserGroupV1) {
	o.UserGroup = v
}

func (o AddUsersToUserGroupV1Output) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddUsersToUserGroupV1Output) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["userGroup"] = o.UserGroup
	return toSerialize, nil
}

type NullableAddUsersToUserGroupV1Output struct {
	value *AddUsersToUserGroupV1Output
	isSet bool
}

func (v NullableAddUsersToUserGroupV1Output) Get() *AddUsersToUserGroupV1Output {
	return v.value
}

func (v *NullableAddUsersToUserGroupV1Output) Set(val *AddUsersToUserGroupV1Output) {
	v.value = val
	v.isSet = true
}

func (v NullableAddUsersToUserGroupV1Output) IsSet() bool {
	return v.isSet
}

func (v *NullableAddUsersToUserGroupV1Output) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddUsersToUserGroupV1Output(
	val *AddUsersToUserGroupV1Output,
) *NullableAddUsersToUserGroupV1Output {
	return &NullableAddUsersToUserGroupV1Output{value: val, isSet: true}
}

func (v NullableAddUsersToUserGroupV1Output) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddUsersToUserGroupV1Output) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
