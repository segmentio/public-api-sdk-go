/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 36.2.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// GetComputedTraitAlphaOutput Computed Trait output for get and update.
type GetComputedTraitAlphaOutput struct {
	ComputedTrait ComputedTrait `json:"computedTrait"`
}

// NewGetComputedTraitAlphaOutput instantiates a new GetComputedTraitAlphaOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetComputedTraitAlphaOutput(computedTrait ComputedTrait) *GetComputedTraitAlphaOutput {
	this := GetComputedTraitAlphaOutput{}
	this.ComputedTrait = computedTrait
	return &this
}

// NewGetComputedTraitAlphaOutputWithDefaults instantiates a new GetComputedTraitAlphaOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetComputedTraitAlphaOutputWithDefaults() *GetComputedTraitAlphaOutput {
	this := GetComputedTraitAlphaOutput{}
	return &this
}

// GetComputedTrait returns the ComputedTrait field value
func (o *GetComputedTraitAlphaOutput) GetComputedTrait() ComputedTrait {
	if o == nil {
		var ret ComputedTrait
		return ret
	}

	return o.ComputedTrait
}

// GetComputedTraitOk returns a tuple with the ComputedTrait field value
// and a boolean to check if the value has been set.
func (o *GetComputedTraitAlphaOutput) GetComputedTraitOk() (*ComputedTrait, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ComputedTrait, true
}

// SetComputedTrait sets field value
func (o *GetComputedTraitAlphaOutput) SetComputedTrait(v ComputedTrait) {
	o.ComputedTrait = v
}

func (o GetComputedTraitAlphaOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["computedTrait"] = o.ComputedTrait
	}
	return json.Marshal(toSerialize)
}

type NullableGetComputedTraitAlphaOutput struct {
	value *GetComputedTraitAlphaOutput
	isSet bool
}

func (v NullableGetComputedTraitAlphaOutput) Get() *GetComputedTraitAlphaOutput {
	return v.value
}

func (v *NullableGetComputedTraitAlphaOutput) Set(val *GetComputedTraitAlphaOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableGetComputedTraitAlphaOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableGetComputedTraitAlphaOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetComputedTraitAlphaOutput(
	val *GetComputedTraitAlphaOutput,
) *NullableGetComputedTraitAlphaOutput {
	return &NullableGetComputedTraitAlphaOutput{value: val, isSet: true}
}

func (v NullableGetComputedTraitAlphaOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetComputedTraitAlphaOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
