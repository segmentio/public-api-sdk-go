/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API. 

API version: 33.0.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// AllowedLabelBeta Defines a label that you may apply to resources within a Workspace.
type AllowedLabelBeta struct {
	// The key identifier for this label.
	Key string `json:"key"`
	// The value of this label.
	Value string `json:"value"`
	// A description of what this label represents.
	Description *string `json:"description,omitempty"`
}

// NewAllowedLabelBeta instantiates a new AllowedLabelBeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAllowedLabelBeta(key string, value string) *AllowedLabelBeta {
	this := AllowedLabelBeta{}
	this.Key = key
	this.Value = value
	return &this
}

// NewAllowedLabelBetaWithDefaults instantiates a new AllowedLabelBeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAllowedLabelBetaWithDefaults() *AllowedLabelBeta {
	this := AllowedLabelBeta{}
	return &this
}

// GetKey returns the Key field value
func (o *AllowedLabelBeta) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *AllowedLabelBeta) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *AllowedLabelBeta) SetKey(v string) {
	o.Key = v
}

// GetValue returns the Value field value
func (o *AllowedLabelBeta) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *AllowedLabelBeta) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *AllowedLabelBeta) SetValue(v string) {
	o.Value = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AllowedLabelBeta) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllowedLabelBeta) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AllowedLabelBeta) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AllowedLabelBeta) SetDescription(v string) {
	o.Description = &v
}

func (o AllowedLabelBeta) MarshalJSON() ([]byte, error) {
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

type NullableAllowedLabelBeta struct {
	value *AllowedLabelBeta
	isSet bool
}

func (v NullableAllowedLabelBeta) Get() *AllowedLabelBeta {
	return v.value
}

func (v *NullableAllowedLabelBeta) Set(val *AllowedLabelBeta) {
	v.value = val
	v.isSet = true
}

func (v NullableAllowedLabelBeta) IsSet() bool {
	return v.isSet
}

func (v *NullableAllowedLabelBeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAllowedLabelBeta(val *AllowedLabelBeta) *NullableAllowedLabelBeta {
	return &NullableAllowedLabelBeta{value: val, isSet: true}
}

func (v NullableAllowedLabelBeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAllowedLabelBeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


