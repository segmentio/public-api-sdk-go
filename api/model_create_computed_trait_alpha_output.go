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

// checks if the CreateComputedTraitAlphaOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateComputedTraitAlphaOutput{}

// CreateComputedTraitAlphaOutput Computed Trait output for create.
type CreateComputedTraitAlphaOutput struct {
	ComputedTrait ComputedTraitSummary `json:"computedTrait"`
}

// NewCreateComputedTraitAlphaOutput instantiates a new CreateComputedTraitAlphaOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateComputedTraitAlphaOutput(
	computedTrait ComputedTraitSummary,
) *CreateComputedTraitAlphaOutput {
	this := CreateComputedTraitAlphaOutput{}
	this.ComputedTrait = computedTrait
	return &this
}

// NewCreateComputedTraitAlphaOutputWithDefaults instantiates a new CreateComputedTraitAlphaOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateComputedTraitAlphaOutputWithDefaults() *CreateComputedTraitAlphaOutput {
	this := CreateComputedTraitAlphaOutput{}
	return &this
}

// GetComputedTrait returns the ComputedTrait field value
func (o *CreateComputedTraitAlphaOutput) GetComputedTrait() ComputedTraitSummary {
	if o == nil {
		var ret ComputedTraitSummary
		return ret
	}

	return o.ComputedTrait
}

// GetComputedTraitOk returns a tuple with the ComputedTrait field value
// and a boolean to check if the value has been set.
func (o *CreateComputedTraitAlphaOutput) GetComputedTraitOk() (*ComputedTraitSummary, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ComputedTrait, true
}

// SetComputedTrait sets field value
func (o *CreateComputedTraitAlphaOutput) SetComputedTrait(v ComputedTraitSummary) {
	o.ComputedTrait = v
}

func (o CreateComputedTraitAlphaOutput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateComputedTraitAlphaOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["computedTrait"] = o.ComputedTrait
	return toSerialize, nil
}

type NullableCreateComputedTraitAlphaOutput struct {
	value *CreateComputedTraitAlphaOutput
	isSet bool
}

func (v NullableCreateComputedTraitAlphaOutput) Get() *CreateComputedTraitAlphaOutput {
	return v.value
}

func (v *NullableCreateComputedTraitAlphaOutput) Set(val *CreateComputedTraitAlphaOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateComputedTraitAlphaOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateComputedTraitAlphaOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateComputedTraitAlphaOutput(
	val *CreateComputedTraitAlphaOutput,
) *NullableCreateComputedTraitAlphaOutput {
	return &NullableCreateComputedTraitAlphaOutput{value: val, isSet: true}
}

func (v NullableCreateComputedTraitAlphaOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateComputedTraitAlphaOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
