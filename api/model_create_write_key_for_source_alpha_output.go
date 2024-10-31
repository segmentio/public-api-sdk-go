/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 55.2.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the CreateWriteKeyForSourceAlphaOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateWriteKeyForSourceAlphaOutput{}

// CreateWriteKeyForSourceAlphaOutput Returns the updated Source.
type CreateWriteKeyForSourceAlphaOutput struct {
	Source SourceAlpha `json:"source"`
}

// NewCreateWriteKeyForSourceAlphaOutput instantiates a new CreateWriteKeyForSourceAlphaOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateWriteKeyForSourceAlphaOutput(source SourceAlpha) *CreateWriteKeyForSourceAlphaOutput {
	this := CreateWriteKeyForSourceAlphaOutput{}
	this.Source = source
	return &this
}

// NewCreateWriteKeyForSourceAlphaOutputWithDefaults instantiates a new CreateWriteKeyForSourceAlphaOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateWriteKeyForSourceAlphaOutputWithDefaults() *CreateWriteKeyForSourceAlphaOutput {
	this := CreateWriteKeyForSourceAlphaOutput{}
	return &this
}

// GetSource returns the Source field value
func (o *CreateWriteKeyForSourceAlphaOutput) GetSource() SourceAlpha {
	if o == nil {
		var ret SourceAlpha
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *CreateWriteKeyForSourceAlphaOutput) GetSourceOk() (*SourceAlpha, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *CreateWriteKeyForSourceAlphaOutput) SetSource(v SourceAlpha) {
	o.Source = v
}

func (o CreateWriteKeyForSourceAlphaOutput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateWriteKeyForSourceAlphaOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["source"] = o.Source
	return toSerialize, nil
}

type NullableCreateWriteKeyForSourceAlphaOutput struct {
	value *CreateWriteKeyForSourceAlphaOutput
	isSet bool
}

func (v NullableCreateWriteKeyForSourceAlphaOutput) Get() *CreateWriteKeyForSourceAlphaOutput {
	return v.value
}

func (v *NullableCreateWriteKeyForSourceAlphaOutput) Set(val *CreateWriteKeyForSourceAlphaOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateWriteKeyForSourceAlphaOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateWriteKeyForSourceAlphaOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateWriteKeyForSourceAlphaOutput(
	val *CreateWriteKeyForSourceAlphaOutput,
) *NullableCreateWriteKeyForSourceAlphaOutput {
	return &NullableCreateWriteKeyForSourceAlphaOutput{value: val, isSet: true}
}

func (v NullableCreateWriteKeyForSourceAlphaOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateWriteKeyForSourceAlphaOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
