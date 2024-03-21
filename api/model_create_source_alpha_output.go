/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 46.0.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the CreateSourceAlphaOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateSourceAlphaOutput{}

// CreateSourceAlphaOutput Returns the newly created Source.
type CreateSourceAlphaOutput struct {
	Source SourceAlpha `json:"source"`
}

// NewCreateSourceAlphaOutput instantiates a new CreateSourceAlphaOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSourceAlphaOutput(source SourceAlpha) *CreateSourceAlphaOutput {
	this := CreateSourceAlphaOutput{}
	this.Source = source
	return &this
}

// NewCreateSourceAlphaOutputWithDefaults instantiates a new CreateSourceAlphaOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSourceAlphaOutputWithDefaults() *CreateSourceAlphaOutput {
	this := CreateSourceAlphaOutput{}
	return &this
}

// GetSource returns the Source field value
func (o *CreateSourceAlphaOutput) GetSource() SourceAlpha {
	if o == nil {
		var ret SourceAlpha
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *CreateSourceAlphaOutput) GetSourceOk() (*SourceAlpha, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *CreateSourceAlphaOutput) SetSource(v SourceAlpha) {
	o.Source = v
}

func (o CreateSourceAlphaOutput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateSourceAlphaOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["source"] = o.Source
	return toSerialize, nil
}

type NullableCreateSourceAlphaOutput struct {
	value *CreateSourceAlphaOutput
	isSet bool
}

func (v NullableCreateSourceAlphaOutput) Get() *CreateSourceAlphaOutput {
	return v.value
}

func (v *NullableCreateSourceAlphaOutput) Set(val *CreateSourceAlphaOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSourceAlphaOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSourceAlphaOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSourceAlphaOutput(
	val *CreateSourceAlphaOutput,
) *NullableCreateSourceAlphaOutput {
	return &NullableCreateSourceAlphaOutput{value: val, isSet: true}
}

func (v NullableCreateSourceAlphaOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateSourceAlphaOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
