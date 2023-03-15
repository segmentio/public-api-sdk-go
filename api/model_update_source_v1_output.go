/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 34.3.2
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// UpdateSourceV1Output Returns the updated Source.
type UpdateSourceV1Output struct {
	Source Source6 `json:"source"`
}

// NewUpdateSourceV1Output instantiates a new UpdateSourceV1Output object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateSourceV1Output(source Source6) *UpdateSourceV1Output {
	this := UpdateSourceV1Output{}
	this.Source = source
	return &this
}

// NewUpdateSourceV1OutputWithDefaults instantiates a new UpdateSourceV1Output object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateSourceV1OutputWithDefaults() *UpdateSourceV1Output {
	this := UpdateSourceV1Output{}
	return &this
}

// GetSource returns the Source field value
func (o *UpdateSourceV1Output) GetSource() Source6 {
	if o == nil {
		var ret Source6
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *UpdateSourceV1Output) GetSourceOk() (*Source6, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *UpdateSourceV1Output) SetSource(v Source6) {
	o.Source = v
}

func (o UpdateSourceV1Output) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["source"] = o.Source
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateSourceV1Output struct {
	value *UpdateSourceV1Output
	isSet bool
}

func (v NullableUpdateSourceV1Output) Get() *UpdateSourceV1Output {
	return v.value
}

func (v *NullableUpdateSourceV1Output) Set(val *UpdateSourceV1Output) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateSourceV1Output) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateSourceV1Output) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateSourceV1Output(val *UpdateSourceV1Output) *NullableUpdateSourceV1Output {
	return &NullableUpdateSourceV1Output{value: val, isSet: true}
}

func (v NullableUpdateSourceV1Output) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateSourceV1Output) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
