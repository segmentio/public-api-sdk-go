/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 57.2.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the UpdateSourceAlphaOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateSourceAlphaOutput{}

// UpdateSourceAlphaOutput Returns the updated Source.
type UpdateSourceAlphaOutput struct {
	Source SourceAlpha `json:"source"`
}

// NewUpdateSourceAlphaOutput instantiates a new UpdateSourceAlphaOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateSourceAlphaOutput(source SourceAlpha) *UpdateSourceAlphaOutput {
	this := UpdateSourceAlphaOutput{}
	this.Source = source
	return &this
}

// NewUpdateSourceAlphaOutputWithDefaults instantiates a new UpdateSourceAlphaOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateSourceAlphaOutputWithDefaults() *UpdateSourceAlphaOutput {
	this := UpdateSourceAlphaOutput{}
	return &this
}

// GetSource returns the Source field value
func (o *UpdateSourceAlphaOutput) GetSource() SourceAlpha {
	if o == nil {
		var ret SourceAlpha
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *UpdateSourceAlphaOutput) GetSourceOk() (*SourceAlpha, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *UpdateSourceAlphaOutput) SetSource(v SourceAlpha) {
	o.Source = v
}

func (o UpdateSourceAlphaOutput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateSourceAlphaOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["source"] = o.Source
	return toSerialize, nil
}

type NullableUpdateSourceAlphaOutput struct {
	value *UpdateSourceAlphaOutput
	isSet bool
}

func (v NullableUpdateSourceAlphaOutput) Get() *UpdateSourceAlphaOutput {
	return v.value
}

func (v *NullableUpdateSourceAlphaOutput) Set(val *UpdateSourceAlphaOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateSourceAlphaOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateSourceAlphaOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateSourceAlphaOutput(
	val *UpdateSourceAlphaOutput,
) *NullableUpdateSourceAlphaOutput {
	return &NullableUpdateSourceAlphaOutput{value: val, isSet: true}
}

func (v NullableUpdateSourceAlphaOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateSourceAlphaOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
