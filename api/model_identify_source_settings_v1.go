/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API. 

API version: 32.0.8
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// IdentifySourceSettingsV1 struct for IdentifySourceSettingsV1
type IdentifySourceSettingsV1 struct {
	// Enable to allow unplanned identify traits.  Config API note: equal to `allowUnplannedIdentifyTraits`.
	AllowUnplannedTraits *bool `json:"allowUnplannedTraits,omitempty"`
	// Enable to allow identify traits on violations.  Config API note: equal to `allowIdentifyTraitsOnViolations`.
	AllowTraitsOnViolations *bool `json:"allowTraitsOnViolations,omitempty"`
	// The common identify event on violations.  Config API note: equal to `commonIdentifyEventOnViolations`.
	CommonEventOnViolations *string `json:"commonEventOnViolations,omitempty"`
}

// NewIdentifySourceSettingsV1 instantiates a new IdentifySourceSettingsV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentifySourceSettingsV1() *IdentifySourceSettingsV1 {
	this := IdentifySourceSettingsV1{}
	return &this
}

// NewIdentifySourceSettingsV1WithDefaults instantiates a new IdentifySourceSettingsV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentifySourceSettingsV1WithDefaults() *IdentifySourceSettingsV1 {
	this := IdentifySourceSettingsV1{}
	return &this
}

// GetAllowUnplannedTraits returns the AllowUnplannedTraits field value if set, zero value otherwise.
func (o *IdentifySourceSettingsV1) GetAllowUnplannedTraits() bool {
	if o == nil || o.AllowUnplannedTraits == nil {
		var ret bool
		return ret
	}
	return *o.AllowUnplannedTraits
}

// GetAllowUnplannedTraitsOk returns a tuple with the AllowUnplannedTraits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentifySourceSettingsV1) GetAllowUnplannedTraitsOk() (*bool, bool) {
	if o == nil || o.AllowUnplannedTraits == nil {
		return nil, false
	}
	return o.AllowUnplannedTraits, true
}

// HasAllowUnplannedTraits returns a boolean if a field has been set.
func (o *IdentifySourceSettingsV1) HasAllowUnplannedTraits() bool {
	if o != nil && o.AllowUnplannedTraits != nil {
		return true
	}

	return false
}

// SetAllowUnplannedTraits gets a reference to the given bool and assigns it to the AllowUnplannedTraits field.
func (o *IdentifySourceSettingsV1) SetAllowUnplannedTraits(v bool) {
	o.AllowUnplannedTraits = &v
}

// GetAllowTraitsOnViolations returns the AllowTraitsOnViolations field value if set, zero value otherwise.
func (o *IdentifySourceSettingsV1) GetAllowTraitsOnViolations() bool {
	if o == nil || o.AllowTraitsOnViolations == nil {
		var ret bool
		return ret
	}
	return *o.AllowTraitsOnViolations
}

// GetAllowTraitsOnViolationsOk returns a tuple with the AllowTraitsOnViolations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentifySourceSettingsV1) GetAllowTraitsOnViolationsOk() (*bool, bool) {
	if o == nil || o.AllowTraitsOnViolations == nil {
		return nil, false
	}
	return o.AllowTraitsOnViolations, true
}

// HasAllowTraitsOnViolations returns a boolean if a field has been set.
func (o *IdentifySourceSettingsV1) HasAllowTraitsOnViolations() bool {
	if o != nil && o.AllowTraitsOnViolations != nil {
		return true
	}

	return false
}

// SetAllowTraitsOnViolations gets a reference to the given bool and assigns it to the AllowTraitsOnViolations field.
func (o *IdentifySourceSettingsV1) SetAllowTraitsOnViolations(v bool) {
	o.AllowTraitsOnViolations = &v
}

// GetCommonEventOnViolations returns the CommonEventOnViolations field value if set, zero value otherwise.
func (o *IdentifySourceSettingsV1) GetCommonEventOnViolations() string {
	if o == nil || o.CommonEventOnViolations == nil {
		var ret string
		return ret
	}
	return *o.CommonEventOnViolations
}

// GetCommonEventOnViolationsOk returns a tuple with the CommonEventOnViolations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentifySourceSettingsV1) GetCommonEventOnViolationsOk() (*string, bool) {
	if o == nil || o.CommonEventOnViolations == nil {
		return nil, false
	}
	return o.CommonEventOnViolations, true
}

// HasCommonEventOnViolations returns a boolean if a field has been set.
func (o *IdentifySourceSettingsV1) HasCommonEventOnViolations() bool {
	if o != nil && o.CommonEventOnViolations != nil {
		return true
	}

	return false
}

// SetCommonEventOnViolations gets a reference to the given string and assigns it to the CommonEventOnViolations field.
func (o *IdentifySourceSettingsV1) SetCommonEventOnViolations(v string) {
	o.CommonEventOnViolations = &v
}

func (o IdentifySourceSettingsV1) MarshalJSON() ([]byte, error) {
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

type NullableIdentifySourceSettingsV1 struct {
	value *IdentifySourceSettingsV1
	isSet bool
}

func (v NullableIdentifySourceSettingsV1) Get() *IdentifySourceSettingsV1 {
	return v.value
}

func (v *NullableIdentifySourceSettingsV1) Set(val *IdentifySourceSettingsV1) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentifySourceSettingsV1) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentifySourceSettingsV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentifySourceSettingsV1(val *IdentifySourceSettingsV1) *NullableIdentifySourceSettingsV1 {
	return &NullableIdentifySourceSettingsV1{value: val, isSet: true}
}

func (v NullableIdentifySourceSettingsV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentifySourceSettingsV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


