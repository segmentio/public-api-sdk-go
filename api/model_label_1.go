/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 34.3.6
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// Label1 The new label to create in the Workspace.
type Label1 struct {
	// The key that represents the name of this label.
	Key string `json:"key"`
	// The value associated with the key of this label.
	Value string `json:"value"`
	// An optional description of the purpose of this label.
	Description *string `json:"description,omitempty"`
}

// NewLabel1 instantiates a new Label1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLabel1(key string, value string) *Label1 {
	this := Label1{}
	this.Key = key
	this.Value = value
	return &this
}

// NewLabel1WithDefaults instantiates a new Label1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLabel1WithDefaults() *Label1 {
	this := Label1{}
	return &this
}

// GetKey returns the Key field value
func (o *Label1) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *Label1) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *Label1) SetKey(v string) {
	o.Key = v
}

// GetValue returns the Value field value
func (o *Label1) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *Label1) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *Label1) SetValue(v string) {
	o.Value = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Label1) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Label1) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Label1) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Label1) SetDescription(v string) {
	o.Description = &v
}

func (o Label1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["key"] = o.Key
	}
	if true {
		toSerialize["value"] = o.Value
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableLabel1 struct {
	value *Label1
	isSet bool
}

func (v NullableLabel1) Get() *Label1 {
	return v.value
}

func (v *NullableLabel1) Set(val *Label1) {
	v.value = val
	v.isSet = true
}

func (v NullableLabel1) IsSet() bool {
	return v.isSet
}

func (v *NullableLabel1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLabel1(val *Label1) *NullableLabel1 {
	return &NullableLabel1{value: val, isSet: true}
}

func (v NullableLabel1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLabel1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
