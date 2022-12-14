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

// GetSourceV1Output Returns a Source.
type GetSourceV1Output struct {
	Source Source4 `json:"source"`
}

// NewGetSourceV1Output instantiates a new GetSourceV1Output object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSourceV1Output(source Source4) *GetSourceV1Output {
	this := GetSourceV1Output{}
	this.Source = source
	return &this
}

// NewGetSourceV1OutputWithDefaults instantiates a new GetSourceV1Output object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSourceV1OutputWithDefaults() *GetSourceV1Output {
	this := GetSourceV1Output{}
	return &this
}

// GetSource returns the Source field value
func (o *GetSourceV1Output) GetSource() Source4 {
	if o == nil {
		var ret Source4
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *GetSourceV1Output) GetSourceOk() (*Source4, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *GetSourceV1Output) SetSource(v Source4) {
	o.Source = v
}

func (o GetSourceV1Output) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["source"] = o.Source
	}
	return json.Marshal(toSerialize)
}

type NullableGetSourceV1Output struct {
	value *GetSourceV1Output
	isSet bool
}

func (v NullableGetSourceV1Output) Get() *GetSourceV1Output {
	return v.value
}

func (v *NullableGetSourceV1Output) Set(val *GetSourceV1Output) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSourceV1Output) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSourceV1Output) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSourceV1Output(val *GetSourceV1Output) *NullableGetSourceV1Output {
	return &NullableGetSourceV1Output{value: val, isSet: true}
}

func (v NullableGetSourceV1Output) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSourceV1Output) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
