/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 44.0.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the SuppressedInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SuppressedInner{}

// SuppressedInner struct for SuppressedInner
type SuppressedInner struct {
	SubjectType string   `json:"subjectType"`
	SubjectIds  []string `json:"subjectIds"`
}

// NewSuppressedInner instantiates a new SuppressedInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSuppressedInner(subjectType string, subjectIds []string) *SuppressedInner {
	this := SuppressedInner{}
	this.SubjectType = subjectType
	this.SubjectIds = subjectIds
	return &this
}

// NewSuppressedInnerWithDefaults instantiates a new SuppressedInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSuppressedInnerWithDefaults() *SuppressedInner {
	this := SuppressedInner{}
	return &this
}

// GetSubjectType returns the SubjectType field value
func (o *SuppressedInner) GetSubjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubjectType
}

// GetSubjectTypeOk returns a tuple with the SubjectType field value
// and a boolean to check if the value has been set.
func (o *SuppressedInner) GetSubjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubjectType, true
}

// SetSubjectType sets field value
func (o *SuppressedInner) SetSubjectType(v string) {
	o.SubjectType = v
}

// GetSubjectIds returns the SubjectIds field value
func (o *SuppressedInner) GetSubjectIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.SubjectIds
}

// GetSubjectIdsOk returns a tuple with the SubjectIds field value
// and a boolean to check if the value has been set.
func (o *SuppressedInner) GetSubjectIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SubjectIds, true
}

// SetSubjectIds sets field value
func (o *SuppressedInner) SetSubjectIds(v []string) {
	o.SubjectIds = v
}

func (o SuppressedInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SuppressedInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["subjectType"] = o.SubjectType
	toSerialize["subjectIds"] = o.SubjectIds
	return toSerialize, nil
}

type NullableSuppressedInner struct {
	value *SuppressedInner
	isSet bool
}

func (v NullableSuppressedInner) Get() *SuppressedInner {
	return v.value
}

func (v *NullableSuppressedInner) Set(val *SuppressedInner) {
	v.value = val
	v.isSet = true
}

func (v NullableSuppressedInner) IsSet() bool {
	return v.isSet
}

func (v *NullableSuppressedInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSuppressedInner(val *SuppressedInner) *NullableSuppressedInner {
	return &NullableSuppressedInner{value: val, isSet: true}
}

func (v NullableSuppressedInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSuppressedInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
