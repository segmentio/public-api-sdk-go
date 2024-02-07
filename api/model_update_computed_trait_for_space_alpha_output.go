/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 40.0.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the UpdateComputedTraitForSpaceAlphaOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateComputedTraitForSpaceAlphaOutput{}

// UpdateComputedTraitForSpaceAlphaOutput Computed Trait output for get and update.
type UpdateComputedTraitForSpaceAlphaOutput struct {
	ComputedTrait ComputedTraitSummary `json:"computedTrait"`
}

// NewUpdateComputedTraitForSpaceAlphaOutput instantiates a new UpdateComputedTraitForSpaceAlphaOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateComputedTraitForSpaceAlphaOutput(
	computedTrait ComputedTraitSummary,
) *UpdateComputedTraitForSpaceAlphaOutput {
	this := UpdateComputedTraitForSpaceAlphaOutput{}
	this.ComputedTrait = computedTrait
	return &this
}

// NewUpdateComputedTraitForSpaceAlphaOutputWithDefaults instantiates a new UpdateComputedTraitForSpaceAlphaOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateComputedTraitForSpaceAlphaOutputWithDefaults() *UpdateComputedTraitForSpaceAlphaOutput {
	this := UpdateComputedTraitForSpaceAlphaOutput{}
	return &this
}

// GetComputedTrait returns the ComputedTrait field value
func (o *UpdateComputedTraitForSpaceAlphaOutput) GetComputedTrait() ComputedTraitSummary {
	if o == nil {
		var ret ComputedTraitSummary
		return ret
	}

	return o.ComputedTrait
}

// GetComputedTraitOk returns a tuple with the ComputedTrait field value
// and a boolean to check if the value has been set.
func (o *UpdateComputedTraitForSpaceAlphaOutput) GetComputedTraitOk() (*ComputedTraitSummary, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ComputedTrait, true
}

// SetComputedTrait sets field value
func (o *UpdateComputedTraitForSpaceAlphaOutput) SetComputedTrait(v ComputedTraitSummary) {
	o.ComputedTrait = v
}

func (o UpdateComputedTraitForSpaceAlphaOutput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateComputedTraitForSpaceAlphaOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["computedTrait"] = o.ComputedTrait
	return toSerialize, nil
}

type NullableUpdateComputedTraitForSpaceAlphaOutput struct {
	value *UpdateComputedTraitForSpaceAlphaOutput
	isSet bool
}

func (v NullableUpdateComputedTraitForSpaceAlphaOutput) Get() *UpdateComputedTraitForSpaceAlphaOutput {
	return v.value
}

func (v *NullableUpdateComputedTraitForSpaceAlphaOutput) Set(
	val *UpdateComputedTraitForSpaceAlphaOutput,
) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateComputedTraitForSpaceAlphaOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateComputedTraitForSpaceAlphaOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateComputedTraitForSpaceAlphaOutput(
	val *UpdateComputedTraitForSpaceAlphaOutput,
) *NullableUpdateComputedTraitForSpaceAlphaOutput {
	return &NullableUpdateComputedTraitForSpaceAlphaOutput{value: val, isSet: true}
}

func (v NullableUpdateComputedTraitForSpaceAlphaOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateComputedTraitForSpaceAlphaOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
