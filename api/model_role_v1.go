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

// checks if the RoleV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoleV1{}

// RoleV1 Represents a role.
type RoleV1 struct {
	// The unique identifier of the role.
	Id string `json:"id"`
	// The human-readable name of the role.
	Name string `json:"name"`
	// The human-readable description of the role.
	Description string `json:"description"`
}

// NewRoleV1 instantiates a new RoleV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleV1(id string, name string, description string) *RoleV1 {
	this := RoleV1{}
	this.Id = id
	this.Name = name
	this.Description = description
	return &this
}

// NewRoleV1WithDefaults instantiates a new RoleV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleV1WithDefaults() *RoleV1 {
	this := RoleV1{}
	return &this
}

// GetId returns the Id field value
func (o *RoleV1) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RoleV1) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *RoleV1) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *RoleV1) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RoleV1) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RoleV1) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *RoleV1) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *RoleV1) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *RoleV1) SetDescription(v string) {
	o.Description = v
}

func (o RoleV1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoleV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description
	return toSerialize, nil
}

type NullableRoleV1 struct {
	value *RoleV1
	isSet bool
}

func (v NullableRoleV1) Get() *RoleV1 {
	return v.value
}

func (v *NullableRoleV1) Set(val *RoleV1) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleV1) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleV1(val *RoleV1) *NullableRoleV1 {
	return &NullableRoleV1{value: val, isSet: true}
}

func (v NullableRoleV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
