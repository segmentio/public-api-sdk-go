/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 38.5.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the MinimalUserGroupV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MinimalUserGroupV1{}

// MinimalUserGroupV1 The least amount of information needed to identify a user group.
type MinimalUserGroupV1 struct {
	// The id of the user group.
	Id string `json:"id"`
	// The name of the user group.
	Name string `json:"name"`
}

// NewMinimalUserGroupV1 instantiates a new MinimalUserGroupV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMinimalUserGroupV1(id string, name string) *MinimalUserGroupV1 {
	this := MinimalUserGroupV1{}
	this.Id = id
	this.Name = name
	return &this
}

// NewMinimalUserGroupV1WithDefaults instantiates a new MinimalUserGroupV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMinimalUserGroupV1WithDefaults() *MinimalUserGroupV1 {
	this := MinimalUserGroupV1{}
	return &this
}

// GetId returns the Id field value
func (o *MinimalUserGroupV1) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *MinimalUserGroupV1) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *MinimalUserGroupV1) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *MinimalUserGroupV1) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *MinimalUserGroupV1) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *MinimalUserGroupV1) SetName(v string) {
	o.Name = v
}

func (o MinimalUserGroupV1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MinimalUserGroupV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullableMinimalUserGroupV1 struct {
	value *MinimalUserGroupV1
	isSet bool
}

func (v NullableMinimalUserGroupV1) Get() *MinimalUserGroupV1 {
	return v.value
}

func (v *NullableMinimalUserGroupV1) Set(val *MinimalUserGroupV1) {
	v.value = val
	v.isSet = true
}

func (v NullableMinimalUserGroupV1) IsSet() bool {
	return v.isSet
}

func (v *NullableMinimalUserGroupV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMinimalUserGroupV1(val *MinimalUserGroupV1) *NullableMinimalUserGroupV1 {
	return &NullableMinimalUserGroupV1{value: val, isSet: true}
}

func (v NullableMinimalUserGroupV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMinimalUserGroupV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
