/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API. 

API version: 32.0.2
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// CreateUserGroupV1Input Creates a user group, used to bundle permissions for its members, within a Workspace.
type CreateUserGroupV1Input struct {
	// The name of the user group to create.
	Name string `json:"name"`
}

// NewCreateUserGroupV1Input instantiates a new CreateUserGroupV1Input object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateUserGroupV1Input(name string) *CreateUserGroupV1Input {
	this := CreateUserGroupV1Input{}
	this.Name = name
	return &this
}

// NewCreateUserGroupV1InputWithDefaults instantiates a new CreateUserGroupV1Input object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateUserGroupV1InputWithDefaults() *CreateUserGroupV1Input {
	this := CreateUserGroupV1Input{}
	return &this
}

// GetName returns the Name field value
func (o *CreateUserGroupV1Input) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateUserGroupV1Input) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateUserGroupV1Input) SetName(v string) {
	o.Name = v
}

func (o CreateUserGroupV1Input) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableCreateUserGroupV1Input struct {
	value *CreateUserGroupV1Input
	isSet bool
}

func (v NullableCreateUserGroupV1Input) Get() *CreateUserGroupV1Input {
	return v.value
}

func (v *NullableCreateUserGroupV1Input) Set(val *CreateUserGroupV1Input) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateUserGroupV1Input) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateUserGroupV1Input) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateUserGroupV1Input(val *CreateUserGroupV1Input) *NullableCreateUserGroupV1Input {
	return &NullableCreateUserGroupV1Input{value: val, isSet: true}
}

func (v NullableCreateUserGroupV1Input) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateUserGroupV1Input) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


