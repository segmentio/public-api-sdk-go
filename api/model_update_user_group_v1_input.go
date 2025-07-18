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

// checks if the UpdateUserGroupV1Input type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateUserGroupV1Input{}

// UpdateUserGroupV1Input Updates a user group with a given id.
type UpdateUserGroupV1Input struct {
	// The intended value to rename the user group to.
	Name string `json:"name"`
}

// NewUpdateUserGroupV1Input instantiates a new UpdateUserGroupV1Input object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateUserGroupV1Input(name string) *UpdateUserGroupV1Input {
	this := UpdateUserGroupV1Input{}
	this.Name = name
	return &this
}

// NewUpdateUserGroupV1InputWithDefaults instantiates a new UpdateUserGroupV1Input object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateUserGroupV1InputWithDefaults() *UpdateUserGroupV1Input {
	this := UpdateUserGroupV1Input{}
	return &this
}

// GetName returns the Name field value
func (o *UpdateUserGroupV1Input) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UpdateUserGroupV1Input) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UpdateUserGroupV1Input) SetName(v string) {
	o.Name = v
}

func (o UpdateUserGroupV1Input) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateUserGroupV1Input) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullableUpdateUserGroupV1Input struct {
	value *UpdateUserGroupV1Input
	isSet bool
}

func (v NullableUpdateUserGroupV1Input) Get() *UpdateUserGroupV1Input {
	return v.value
}

func (v *NullableUpdateUserGroupV1Input) Set(val *UpdateUserGroupV1Input) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateUserGroupV1Input) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateUserGroupV1Input) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateUserGroupV1Input(
	val *UpdateUserGroupV1Input,
) *NullableUpdateUserGroupV1Input {
	return &NullableUpdateUserGroupV1Input{value: val, isSet: true}
}

func (v NullableUpdateUserGroupV1Input) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateUserGroupV1Input) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
