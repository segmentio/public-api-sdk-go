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

// Identify Identify settings.
type Identify struct {
	// Enable to allow unplanned identify traits.  Config API note: equal to `allowUnplannedIdentifyTraits`.
	AllowUnplannedTraits *bool `json:"allowUnplannedTraits,omitempty"`
	// Enable to allow identify traits on violations.  Config API note: equal to `allowIdentifyTraitsOnViolations`.
	AllowTraitsOnViolations *bool `json:"allowTraitsOnViolations,omitempty"`
	// The common identify event on violations.  Config API note: equal to `commonIdentifyEventOnViolations`.
	CommonEventOnViolations *string `json:"commonEventOnViolations,omitempty"`
}

// NewIdentify instantiates a new Identify object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentify() *Identify {
	this := Identify{}
	return &this
}

// NewIdentifyWithDefaults instantiates a new Identify object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentifyWithDefaults() *Identify {
	this := Identify{}
	return &this
}

// GetAllowUnplannedTraits returns the AllowUnplannedTraits field value if set, zero value otherwise.
func (o *Identify) GetAllowUnplannedTraits() bool {
	if o == nil || o.AllowUnplannedTraits == nil {
		var ret bool
		return ret
	}
	return *o.AllowUnplannedTraits
}

// GetAllowUnplannedTraitsOk returns a tuple with the AllowUnplannedTraits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Identify) GetAllowUnplannedTraitsOk() (*bool, bool) {
	if o == nil || o.AllowUnplannedTraits == nil {
		return nil, false
	}
	return o.AllowUnplannedTraits, true
}

// HasAllowUnplannedTraits returns a boolean if a field has been set.
func (o *Identify) HasAllowUnplannedTraits() bool {
	if o != nil && o.AllowUnplannedTraits != nil {
		return true
	}

	return false
}

// SetAllowUnplannedTraits gets a reference to the given bool and assigns it to the AllowUnplannedTraits field.
func (o *Identify) SetAllowUnplannedTraits(v bool) {
	o.AllowUnplannedTraits = &v
}

// GetAllowTraitsOnViolations returns the AllowTraitsOnViolations field value if set, zero value otherwise.
func (o *Identify) GetAllowTraitsOnViolations() bool {
	if o == nil || o.AllowTraitsOnViolations == nil {
		var ret bool
		return ret
	}
	return *o.AllowTraitsOnViolations
}

// GetAllowTraitsOnViolationsOk returns a tuple with the AllowTraitsOnViolations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Identify) GetAllowTraitsOnViolationsOk() (*bool, bool) {
	if o == nil || o.AllowTraitsOnViolations == nil {
		return nil, false
	}
	return o.AllowTraitsOnViolations, true
}

// HasAllowTraitsOnViolations returns a boolean if a field has been set.
func (o *Identify) HasAllowTraitsOnViolations() bool {
	if o != nil && o.AllowTraitsOnViolations != nil {
		return true
	}

	return false
}

// SetAllowTraitsOnViolations gets a reference to the given bool and assigns it to the AllowTraitsOnViolations field.
func (o *Identify) SetAllowTraitsOnViolations(v bool) {
	o.AllowTraitsOnViolations = &v
}

// GetCommonEventOnViolations returns the CommonEventOnViolations field value if set, zero value otherwise.
func (o *Identify) GetCommonEventOnViolations() string {
	if o == nil || o.CommonEventOnViolations == nil {
		var ret string
		return ret
	}
	return *o.CommonEventOnViolations
}

// GetCommonEventOnViolationsOk returns a tuple with the CommonEventOnViolations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Identify) GetCommonEventOnViolationsOk() (*string, bool) {
	if o == nil || o.CommonEventOnViolations == nil {
		return nil, false
	}
	return o.CommonEventOnViolations, true
}

// HasCommonEventOnViolations returns a boolean if a field has been set.
func (o *Identify) HasCommonEventOnViolations() bool {
	if o != nil && o.CommonEventOnViolations != nil {
		return true
	}

	return false
}

// SetCommonEventOnViolations gets a reference to the given string and assigns it to the CommonEventOnViolations field.
func (o *Identify) SetCommonEventOnViolations(v string) {
	o.CommonEventOnViolations = &v
}

func (o Identify) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AllowUnplannedTraits != nil {
		toSerialize["allowUnplannedTraits"] = o.AllowUnplannedTraits
	}
	if o.AllowTraitsOnViolations != nil {
		toSerialize["allowTraitsOnViolations"] = o.AllowTraitsOnViolations
	}
	if o.CommonEventOnViolations != nil {
		toSerialize["commonEventOnViolations"] = o.CommonEventOnViolations
	}
	return json.Marshal(toSerialize)
}

type NullableIdentify struct {
	value *Identify
	isSet bool
}

func (v NullableIdentify) Get() *Identify {
	return v.value
}

func (v *NullableIdentify) Set(val *Identify) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentify) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentify) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentify(val *Identify) *NullableIdentify {
	return &NullableIdentify{value: val, isSet: true}
}

func (v NullableIdentify) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentify) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


