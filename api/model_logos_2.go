/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 34.1.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// Logos2 Logo information for this object.
type Logos2 struct {
	// The default URL for this logo.
	Default string `json:"default"`
	// The logo mark.
	Mark NullableString `json:"mark,omitempty"`
	// The alternative text for this logo.
	Alt NullableString `json:"alt,omitempty"`
}

// NewLogos2 instantiates a new Logos2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogos2(default_ string) *Logos2 {
	this := Logos2{}
	this.Default = default_
	return &this
}

// NewLogos2WithDefaults instantiates a new Logos2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogos2WithDefaults() *Logos2 {
	this := Logos2{}
	return &this
}

// GetDefault returns the Default field value
func (o *Logos2) GetDefault() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Default
}

// GetDefaultOk returns a tuple with the Default field value
// and a boolean to check if the value has been set.
func (o *Logos2) GetDefaultOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Default, true
}

// SetDefault sets field value
func (o *Logos2) SetDefault(v string) {
	o.Default = v
}

// GetMark returns the Mark field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Logos2) GetMark() string {
	if o == nil || o.Mark.Get() == nil {
		var ret string
		return ret
	}
	return *o.Mark.Get()
}

// GetMarkOk returns a tuple with the Mark field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Logos2) GetMarkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Mark.Get(), o.Mark.IsSet()
}

// HasMark returns a boolean if a field has been set.
func (o *Logos2) HasMark() bool {
	if o != nil && o.Mark.IsSet() {
		return true
	}

	return false
}

// SetMark gets a reference to the given NullableString and assigns it to the Mark field.
func (o *Logos2) SetMark(v string) {
	o.Mark.Set(&v)
}

// SetMarkNil sets the value for Mark to be an explicit nil
func (o *Logos2) SetMarkNil() {
	o.Mark.Set(nil)
}

// UnsetMark ensures that no value is present for Mark, not even an explicit nil
func (o *Logos2) UnsetMark() {
	o.Mark.Unset()
}

// GetAlt returns the Alt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Logos2) GetAlt() string {
	if o == nil || o.Alt.Get() == nil {
		var ret string
		return ret
	}
	return *o.Alt.Get()
}

// GetAltOk returns a tuple with the Alt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Logos2) GetAltOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Alt.Get(), o.Alt.IsSet()
}

// HasAlt returns a boolean if a field has been set.
func (o *Logos2) HasAlt() bool {
	if o != nil && o.Alt.IsSet() {
		return true
	}

	return false
}

// SetAlt gets a reference to the given NullableString and assigns it to the Alt field.
func (o *Logos2) SetAlt(v string) {
	o.Alt.Set(&v)
}

// SetAltNil sets the value for Alt to be an explicit nil
func (o *Logos2) SetAltNil() {
	o.Alt.Set(nil)
}

// UnsetAlt ensures that no value is present for Alt, not even an explicit nil
func (o *Logos2) UnsetAlt() {
	o.Alt.Unset()
}

func (o Logos2) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["default"] = o.Default
	}
	if o.Mark.IsSet() {
		toSerialize["mark"] = o.Mark.Get()
	}
	if o.Alt.IsSet() {
		toSerialize["alt"] = o.Alt.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableLogos2 struct {
	value *Logos2
	isSet bool
}

func (v NullableLogos2) Get() *Logos2 {
	return v.value
}

func (v *NullableLogos2) Set(val *Logos2) {
	v.value = val
	v.isSet = true
}

func (v NullableLogos2) IsSet() bool {
	return v.isSet
}

func (v *NullableLogos2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogos2(val *Logos2) *NullableLogos2 {
	return &NullableLogos2{value: val, isSet: true}
}

func (v NullableLogos2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogos2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
