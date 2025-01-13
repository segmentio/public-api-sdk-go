/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 57.3.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the MinimalUserV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MinimalUserV1{}

// MinimalUserV1 A user belonging to a Segment Workspace.
type MinimalUserV1 struct {
	// The unique identifier of this user.  Config API note: analogous to `name`.
	Id string `json:"id"`
	// The human-readable name of this user.
	Name string `json:"name"`
	// The email address associated with this user.
	Email string `json:"email"`
}

// NewMinimalUserV1 instantiates a new MinimalUserV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMinimalUserV1(id string, name string, email string) *MinimalUserV1 {
	this := MinimalUserV1{}
	this.Id = id
	this.Name = name
	this.Email = email
	return &this
}

// NewMinimalUserV1WithDefaults instantiates a new MinimalUserV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMinimalUserV1WithDefaults() *MinimalUserV1 {
	this := MinimalUserV1{}
	return &this
}

// GetId returns the Id field value
func (o *MinimalUserV1) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *MinimalUserV1) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *MinimalUserV1) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *MinimalUserV1) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *MinimalUserV1) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *MinimalUserV1) SetName(v string) {
	o.Name = v
}

// GetEmail returns the Email field value
func (o *MinimalUserV1) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *MinimalUserV1) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *MinimalUserV1) SetEmail(v string) {
	o.Email = v
}

func (o MinimalUserV1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MinimalUserV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["email"] = o.Email
	return toSerialize, nil
}

type NullableMinimalUserV1 struct {
	value *MinimalUserV1
	isSet bool
}

func (v NullableMinimalUserV1) Get() *MinimalUserV1 {
	return v.value
}

func (v *NullableMinimalUserV1) Set(val *MinimalUserV1) {
	v.value = val
	v.isSet = true
}

func (v NullableMinimalUserV1) IsSet() bool {
	return v.isSet
}

func (v *NullableMinimalUserV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMinimalUserV1(val *MinimalUserV1) *NullableMinimalUserV1 {
	return &NullableMinimalUserV1{value: val, isSet: true}
}

func (v NullableMinimalUserV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMinimalUserV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
