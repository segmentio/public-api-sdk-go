/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 33.0.4
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// DestinationFilterActionV1 Represents a Destination filter action.
type DestinationFilterActionV1 struct {
	// The kind of Transformation to apply to any matched properties.
	Type string `json:"type"`
	// A dictionary of paths to object keys that this filter applies to.   The literal string '' represents the top level of the object.
	Fields map[string]interface{} `json:"fields,omitempty"`
	// A decimal between 0 and 1 used for 'sample' type events and influences the likelihood of sampling to occur.
	Percent *float32 `json:"percent,omitempty"`
	// The JSON path to a property within a payload object from which Segment generates a deterministic sampling rate.
	Path *string `json:"path,omitempty"`
}

// NewDestinationFilterActionV1 instantiates a new DestinationFilterActionV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDestinationFilterActionV1(type_ string) *DestinationFilterActionV1 {
	this := DestinationFilterActionV1{}
	this.Type = type_
	return &this
}

// NewDestinationFilterActionV1WithDefaults instantiates a new DestinationFilterActionV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDestinationFilterActionV1WithDefaults() *DestinationFilterActionV1 {
	this := DestinationFilterActionV1{}
	return &this
}

// GetType returns the Type field value
func (o *DestinationFilterActionV1) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DestinationFilterActionV1) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *DestinationFilterActionV1) SetType(v string) {
	o.Type = v
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *DestinationFilterActionV1) GetFields() map[string]interface{} {
	if o == nil || o.Fields == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DestinationFilterActionV1) GetFieldsOk() (map[string]interface{}, bool) {
	if o == nil || o.Fields == nil {
		return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *DestinationFilterActionV1) HasFields() bool {
	if o != nil && o.Fields != nil {
		return true
	}

	return false
}

// SetFields gets a reference to the given map[string]interface{} and assigns it to the Fields field.
func (o *DestinationFilterActionV1) SetFields(v map[string]interface{}) {
	o.Fields = v
}

// GetPercent returns the Percent field value if set, zero value otherwise.
func (o *DestinationFilterActionV1) GetPercent() float32 {
	if o == nil || o.Percent == nil {
		var ret float32
		return ret
	}
	return *o.Percent
}

// GetPercentOk returns a tuple with the Percent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DestinationFilterActionV1) GetPercentOk() (*float32, bool) {
	if o == nil || o.Percent == nil {
		return nil, false
	}
	return o.Percent, true
}

// HasPercent returns a boolean if a field has been set.
func (o *DestinationFilterActionV1) HasPercent() bool {
	if o != nil && o.Percent != nil {
		return true
	}

	return false
}

// SetPercent gets a reference to the given float32 and assigns it to the Percent field.
func (o *DestinationFilterActionV1) SetPercent(v float32) {
	o.Percent = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *DestinationFilterActionV1) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DestinationFilterActionV1) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *DestinationFilterActionV1) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *DestinationFilterActionV1) SetPath(v string) {
	o.Path = &v
}

func (o DestinationFilterActionV1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.Fields != nil {
		toSerialize["fields"] = o.Fields
	}
	if o.Percent != nil {
		toSerialize["percent"] = o.Percent
	}
	if o.Path != nil {
		toSerialize["path"] = o.Path
	}
	return json.Marshal(toSerialize)
}

type NullableDestinationFilterActionV1 struct {
	value *DestinationFilterActionV1
	isSet bool
}

func (v NullableDestinationFilterActionV1) Get() *DestinationFilterActionV1 {
	return v.value
}

func (v *NullableDestinationFilterActionV1) Set(val *DestinationFilterActionV1) {
	v.value = val
	v.isSet = true
}

func (v NullableDestinationFilterActionV1) IsSet() bool {
	return v.isSet
}

func (v *NullableDestinationFilterActionV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDestinationFilterActionV1(
	val *DestinationFilterActionV1,
) *NullableDestinationFilterActionV1 {
	return &NullableDestinationFilterActionV1{value: val, isSet: true}
}

func (v NullableDestinationFilterActionV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDestinationFilterActionV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
