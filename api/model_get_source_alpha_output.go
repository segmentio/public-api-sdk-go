/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 36.1.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// GetSourceAlphaOutput Returns a Source.
type GetSourceAlphaOutput struct {
	Source Source1 `json:"source"`
}

// NewGetSourceAlphaOutput instantiates a new GetSourceAlphaOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSourceAlphaOutput(source Source1) *GetSourceAlphaOutput {
	this := GetSourceAlphaOutput{}
	this.Source = source
	return &this
}

// NewGetSourceAlphaOutputWithDefaults instantiates a new GetSourceAlphaOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSourceAlphaOutputWithDefaults() *GetSourceAlphaOutput {
	this := GetSourceAlphaOutput{}
	return &this
}

// GetSource returns the Source field value
func (o *GetSourceAlphaOutput) GetSource() Source1 {
	if o == nil {
		var ret Source1
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *GetSourceAlphaOutput) GetSourceOk() (*Source1, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *GetSourceAlphaOutput) SetSource(v Source1) {
	o.Source = v
}

func (o GetSourceAlphaOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["source"] = o.Source
	}
	return json.Marshal(toSerialize)
}

type NullableGetSourceAlphaOutput struct {
	value *GetSourceAlphaOutput
	isSet bool
}

func (v NullableGetSourceAlphaOutput) Get() *GetSourceAlphaOutput {
	return v.value
}

func (v *NullableGetSourceAlphaOutput) Set(val *GetSourceAlphaOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSourceAlphaOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSourceAlphaOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSourceAlphaOutput(val *GetSourceAlphaOutput) *NullableGetSourceAlphaOutput {
	return &NullableGetSourceAlphaOutput{value: val, isSet: true}
}

func (v NullableGetSourceAlphaOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSourceAlphaOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
