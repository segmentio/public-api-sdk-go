/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 53.0.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the LogosBeta type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LogosBeta{}

// LogosBeta Represents a logo.
type LogosBeta struct {
	// The default URL for this logo.
	Default string `json:"default"`
	// The logo mark.
	Mark NullableString `json:"mark,omitempty"`
	// The alternative text for this logo.
	Alt NullableString `json:"alt,omitempty"`
}

// NewLogosBeta instantiates a new LogosBeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogosBeta(default_ string) *LogosBeta {
	this := LogosBeta{}
	this.Default = default_
	return &this
}

// NewLogosBetaWithDefaults instantiates a new LogosBeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogosBetaWithDefaults() *LogosBeta {
	this := LogosBeta{}
	return &this
}

// GetDefault returns the Default field value
func (o *LogosBeta) GetDefault() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Default
}

// GetDefaultOk returns a tuple with the Default field value
// and a boolean to check if the value has been set.
func (o *LogosBeta) GetDefaultOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Default, true
}

// SetDefault sets field value
func (o *LogosBeta) SetDefault(v string) {
	o.Default = v
}

// GetMark returns the Mark field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LogosBeta) GetMark() string {
	if o == nil || IsNil(o.Mark.Get()) {
		var ret string
		return ret
	}
	return *o.Mark.Get()
}

// GetMarkOk returns a tuple with the Mark field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LogosBeta) GetMarkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Mark.Get(), o.Mark.IsSet()
}

// HasMark returns a boolean if a field has been set.
func (o *LogosBeta) HasMark() bool {
	if o != nil && o.Mark.IsSet() {
		return true
	}

	return false
}

// SetMark gets a reference to the given NullableString and assigns it to the Mark field.
func (o *LogosBeta) SetMark(v string) {
	o.Mark.Set(&v)
}

// SetMarkNil sets the value for Mark to be an explicit nil
func (o *LogosBeta) SetMarkNil() {
	o.Mark.Set(nil)
}

// UnsetMark ensures that no value is present for Mark, not even an explicit nil
func (o *LogosBeta) UnsetMark() {
	o.Mark.Unset()
}

// GetAlt returns the Alt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LogosBeta) GetAlt() string {
	if o == nil || IsNil(o.Alt.Get()) {
		var ret string
		return ret
	}
	return *o.Alt.Get()
}

// GetAltOk returns a tuple with the Alt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LogosBeta) GetAltOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Alt.Get(), o.Alt.IsSet()
}

// HasAlt returns a boolean if a field has been set.
func (o *LogosBeta) HasAlt() bool {
	if o != nil && o.Alt.IsSet() {
		return true
	}

	return false
}

// SetAlt gets a reference to the given NullableString and assigns it to the Alt field.
func (o *LogosBeta) SetAlt(v string) {
	o.Alt.Set(&v)
}

// SetAltNil sets the value for Alt to be an explicit nil
func (o *LogosBeta) SetAltNil() {
	o.Alt.Set(nil)
}

// UnsetAlt ensures that no value is present for Alt, not even an explicit nil
func (o *LogosBeta) UnsetAlt() {
	o.Alt.Unset()
}

func (o LogosBeta) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LogosBeta) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["default"] = o.Default
	if o.Mark.IsSet() {
		toSerialize["mark"] = o.Mark.Get()
	}
	if o.Alt.IsSet() {
		toSerialize["alt"] = o.Alt.Get()
	}
	return toSerialize, nil
}

type NullableLogosBeta struct {
	value *LogosBeta
	isSet bool
}

func (v NullableLogosBeta) Get() *LogosBeta {
	return v.value
}

func (v *NullableLogosBeta) Set(val *LogosBeta) {
	v.value = val
	v.isSet = true
}

func (v NullableLogosBeta) IsSet() bool {
	return v.isSet
}

func (v *NullableLogosBeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogosBeta(val *LogosBeta) *NullableLogosBeta {
	return &NullableLogosBeta{value: val, isSet: true}
}

func (v NullableLogosBeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogosBeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
