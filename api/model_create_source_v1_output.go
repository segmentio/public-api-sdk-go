/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 36.1.1
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// CreateSourceV1Output Returns a newly created Source.
type CreateSourceV1Output struct {
	Source Source5 `json:"source"`
}

// NewCreateSourceV1Output instantiates a new CreateSourceV1Output object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSourceV1Output(source Source5) *CreateSourceV1Output {
	this := CreateSourceV1Output{}
	this.Source = source
	return &this
}

// NewCreateSourceV1OutputWithDefaults instantiates a new CreateSourceV1Output object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSourceV1OutputWithDefaults() *CreateSourceV1Output {
	this := CreateSourceV1Output{}
	return &this
}

// GetSource returns the Source field value
func (o *CreateSourceV1Output) GetSource() Source5 {
	if o == nil {
		var ret Source5
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *CreateSourceV1Output) GetSourceOk() (*Source5, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *CreateSourceV1Output) SetSource(v Source5) {
	o.Source = v
}

func (o CreateSourceV1Output) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["source"] = o.Source
	}
	return json.Marshal(toSerialize)
}

type NullableCreateSourceV1Output struct {
	value *CreateSourceV1Output
	isSet bool
}

func (v NullableCreateSourceV1Output) Get() *CreateSourceV1Output {
	return v.value
}

func (v *NullableCreateSourceV1Output) Set(val *CreateSourceV1Output) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSourceV1Output) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSourceV1Output) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSourceV1Output(val *CreateSourceV1Output) *NullableCreateSourceV1Output {
	return &NullableCreateSourceV1Output{value: val, isSet: true}
}

func (v NullableCreateSourceV1Output) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateSourceV1Output) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
