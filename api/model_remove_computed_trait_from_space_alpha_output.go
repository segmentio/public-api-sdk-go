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

// RemoveComputedTraitFromSpaceAlphaOutput Delete computed trait endpoint output.
type RemoveComputedTraitFromSpaceAlphaOutput struct {
	// The status of the operation.
	Status string `json:"status"`
}

// NewRemoveComputedTraitFromSpaceAlphaOutput instantiates a new RemoveComputedTraitFromSpaceAlphaOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemoveComputedTraitFromSpaceAlphaOutput(
	status string,
) *RemoveComputedTraitFromSpaceAlphaOutput {
	this := RemoveComputedTraitFromSpaceAlphaOutput{}
	this.Status = status
	return &this
}

// NewRemoveComputedTraitFromSpaceAlphaOutputWithDefaults instantiates a new RemoveComputedTraitFromSpaceAlphaOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemoveComputedTraitFromSpaceAlphaOutputWithDefaults() *RemoveComputedTraitFromSpaceAlphaOutput {
	this := RemoveComputedTraitFromSpaceAlphaOutput{}
	return &this
}

// GetStatus returns the Status field value
func (o *RemoveComputedTraitFromSpaceAlphaOutput) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *RemoveComputedTraitFromSpaceAlphaOutput) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *RemoveComputedTraitFromSpaceAlphaOutput) SetStatus(v string) {
	o.Status = v
}

func (o RemoveComputedTraitFromSpaceAlphaOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableRemoveComputedTraitFromSpaceAlphaOutput struct {
	value *RemoveComputedTraitFromSpaceAlphaOutput
	isSet bool
}

func (v NullableRemoveComputedTraitFromSpaceAlphaOutput) Get() *RemoveComputedTraitFromSpaceAlphaOutput {
	return v.value
}

func (v *NullableRemoveComputedTraitFromSpaceAlphaOutput) Set(
	val *RemoveComputedTraitFromSpaceAlphaOutput,
) {
	v.value = val
	v.isSet = true
}

func (v NullableRemoveComputedTraitFromSpaceAlphaOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableRemoveComputedTraitFromSpaceAlphaOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemoveComputedTraitFromSpaceAlphaOutput(
	val *RemoveComputedTraitFromSpaceAlphaOutput,
) *NullableRemoveComputedTraitFromSpaceAlphaOutput {
	return &NullableRemoveComputedTraitFromSpaceAlphaOutput{value: val, isSet: true}
}

func (v NullableRemoveComputedTraitFromSpaceAlphaOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemoveComputedTraitFromSpaceAlphaOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}