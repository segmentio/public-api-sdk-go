/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API. 

API version: 32.0.5
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// Label2 The newly created label.
type Label2 struct {
	// The key that represents the name of this label.
	Key string `json:"key"`
	// The value associated with the key of this label.
	Value string `json:"value"`
	// An optional description of the purpose of this label.
	Description *string `json:"description,omitempty"`
}

// NewLabel2 instantiates a new Label2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLabel2(key string, value string) *Label2 {
	this := Label2{}
	this.Key = key
	this.Value = value
	return &this
}

// NewLabel2WithDefaults instantiates a new Label2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLabel2WithDefaults() *Label2 {
	this := Label2{}
	return &this
}

// GetKey returns the Key field value
func (o *Label2) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *Label2) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *Label2) SetKey(v string) {
	o.Key = v
}

// GetValue returns the Value field value
func (o *Label2) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *Label2) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *Label2) SetValue(v string) {
	o.Value = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Label2) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Label2) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Label2) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Label2) SetDescription(v string) {
	o.Description = &v
}

func (o Label2) MarshalJSON() ([]byte, error) {
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

type NullableLabel2 struct {
	value *Label2
	isSet bool
}

func (v NullableLabel2) Get() *Label2 {
	return v.value
}

func (v *NullableLabel2) Set(val *Label2) {
	v.value = val
	v.isSet = true
}

func (v NullableLabel2) IsSet() bool {
	return v.isSet
}

func (v *NullableLabel2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLabel2(val *Label2) *NullableLabel2 {
	return &NullableLabel2{value: val, isSet: true}
}

func (v NullableLabel2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLabel2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


