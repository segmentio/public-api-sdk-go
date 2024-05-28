/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 50.2.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the GroupSourceSettingsV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupSourceSettingsV1{}

// GroupSourceSettingsV1 struct for GroupSourceSettingsV1
type GroupSourceSettingsV1 struct {
	// Enable to allow unplanned group traits.  Config API note: equal to `allowUnplannedGroupTraits`.
	AllowUnplannedTraits *bool `json:"allowUnplannedTraits,omitempty"`
	// Enable to allow group traits on violations.  Config API note: equal to `allowGroupTraitsOnViolations`.
	AllowTraitsOnViolations *bool `json:"allowTraitsOnViolations,omitempty"`
	// The common group event on violations.  Config API note: equal to `commonGroupEventOnViolations`.
	CommonEventOnViolations *string `json:"commonEventOnViolations,omitempty"`
}

// NewGroupSourceSettingsV1 instantiates a new GroupSourceSettingsV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupSourceSettingsV1() *GroupSourceSettingsV1 {
	this := GroupSourceSettingsV1{}
	return &this
}

// NewGroupSourceSettingsV1WithDefaults instantiates a new GroupSourceSettingsV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupSourceSettingsV1WithDefaults() *GroupSourceSettingsV1 {
	this := GroupSourceSettingsV1{}
	return &this
}

// GetAllowUnplannedTraits returns the AllowUnplannedTraits field value if set, zero value otherwise.
func (o *GroupSourceSettingsV1) GetAllowUnplannedTraits() bool {
	if o == nil || IsNil(o.AllowUnplannedTraits) {
		var ret bool
		return ret
	}
	return *o.AllowUnplannedTraits
}

// GetAllowUnplannedTraitsOk returns a tuple with the AllowUnplannedTraits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupSourceSettingsV1) GetAllowUnplannedTraitsOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowUnplannedTraits) {
		return nil, false
	}
	return o.AllowUnplannedTraits, true
}

// HasAllowUnplannedTraits returns a boolean if a field has been set.
func (o *GroupSourceSettingsV1) HasAllowUnplannedTraits() bool {
	if o != nil && !IsNil(o.AllowUnplannedTraits) {
		return true
	}

	return false
}

// SetAllowUnplannedTraits gets a reference to the given bool and assigns it to the AllowUnplannedTraits field.
func (o *GroupSourceSettingsV1) SetAllowUnplannedTraits(v bool) {
	o.AllowUnplannedTraits = &v
}

// GetAllowTraitsOnViolations returns the AllowTraitsOnViolations field value if set, zero value otherwise.
func (o *GroupSourceSettingsV1) GetAllowTraitsOnViolations() bool {
	if o == nil || IsNil(o.AllowTraitsOnViolations) {
		var ret bool
		return ret
	}
	return *o.AllowTraitsOnViolations
}

// GetAllowTraitsOnViolationsOk returns a tuple with the AllowTraitsOnViolations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupSourceSettingsV1) GetAllowTraitsOnViolationsOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowTraitsOnViolations) {
		return nil, false
	}
	return o.AllowTraitsOnViolations, true
}

// HasAllowTraitsOnViolations returns a boolean if a field has been set.
func (o *GroupSourceSettingsV1) HasAllowTraitsOnViolations() bool {
	if o != nil && !IsNil(o.AllowTraitsOnViolations) {
		return true
	}

	return false
}

// SetAllowTraitsOnViolations gets a reference to the given bool and assigns it to the AllowTraitsOnViolations field.
func (o *GroupSourceSettingsV1) SetAllowTraitsOnViolations(v bool) {
	o.AllowTraitsOnViolations = &v
}

// GetCommonEventOnViolations returns the CommonEventOnViolations field value if set, zero value otherwise.
func (o *GroupSourceSettingsV1) GetCommonEventOnViolations() string {
	if o == nil || IsNil(o.CommonEventOnViolations) {
		var ret string
		return ret
	}
	return *o.CommonEventOnViolations
}

// GetCommonEventOnViolationsOk returns a tuple with the CommonEventOnViolations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupSourceSettingsV1) GetCommonEventOnViolationsOk() (*string, bool) {
	if o == nil || IsNil(o.CommonEventOnViolations) {
		return nil, false
	}
	return o.CommonEventOnViolations, true
}

// HasCommonEventOnViolations returns a boolean if a field has been set.
func (o *GroupSourceSettingsV1) HasCommonEventOnViolations() bool {
	if o != nil && !IsNil(o.CommonEventOnViolations) {
		return true
	}

	return false
}

// SetCommonEventOnViolations gets a reference to the given string and assigns it to the CommonEventOnViolations field.
func (o *GroupSourceSettingsV1) SetCommonEventOnViolations(v string) {
	o.CommonEventOnViolations = &v
}

func (o GroupSourceSettingsV1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupSourceSettingsV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AllowUnplannedTraits) {
		toSerialize["allowUnplannedTraits"] = o.AllowUnplannedTraits
	}
	if !IsNil(o.AllowTraitsOnViolations) {
		toSerialize["allowTraitsOnViolations"] = o.AllowTraitsOnViolations
	}
	if !IsNil(o.CommonEventOnViolations) {
		toSerialize["commonEventOnViolations"] = o.CommonEventOnViolations
	}
	return toSerialize, nil
}

type NullableGroupSourceSettingsV1 struct {
	value *GroupSourceSettingsV1
	isSet bool
}

func (v NullableGroupSourceSettingsV1) Get() *GroupSourceSettingsV1 {
	return v.value
}

func (v *NullableGroupSourceSettingsV1) Set(val *GroupSourceSettingsV1) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupSourceSettingsV1) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupSourceSettingsV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupSourceSettingsV1(val *GroupSourceSettingsV1) *NullableGroupSourceSettingsV1 {
	return &NullableGroupSourceSettingsV1{value: val, isSet: true}
}

func (v NullableGroupSourceSettingsV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupSourceSettingsV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
