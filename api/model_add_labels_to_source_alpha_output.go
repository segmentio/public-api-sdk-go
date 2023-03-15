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

// AddLabelsToSourceAlphaOutput Applies an existing label to an existing Source.
type AddLabelsToSourceAlphaOutput struct {
	// All labels applied to the Source.
	Labels []LabelAlpha `json:"labels"`
}

// NewAddLabelsToSourceAlphaOutput instantiates a new AddLabelsToSourceAlphaOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddLabelsToSourceAlphaOutput(labels []LabelAlpha) *AddLabelsToSourceAlphaOutput {
	this := AddLabelsToSourceAlphaOutput{}
	this.Labels = labels
	return &this
}

// NewAddLabelsToSourceAlphaOutputWithDefaults instantiates a new AddLabelsToSourceAlphaOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddLabelsToSourceAlphaOutputWithDefaults() *AddLabelsToSourceAlphaOutput {
	this := AddLabelsToSourceAlphaOutput{}
	return &this
}

// GetLabels returns the Labels field value
func (o *AddLabelsToSourceAlphaOutput) GetLabels() []LabelAlpha {
	if o == nil {
		var ret []LabelAlpha
		return ret
	}

	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value
// and a boolean to check if the value has been set.
func (o *AddLabelsToSourceAlphaOutput) GetLabelsOk() ([]LabelAlpha, bool) {
	if o == nil {
		return nil, false
	}
	return o.Labels, true
}

// SetLabels sets field value
func (o *AddLabelsToSourceAlphaOutput) SetLabels(v []LabelAlpha) {
	o.Labels = v
}

func (o AddLabelsToSourceAlphaOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["labels"] = o.Labels
	}
	return json.Marshal(toSerialize)
}

type NullableAddLabelsToSourceAlphaOutput struct {
	value *AddLabelsToSourceAlphaOutput
	isSet bool
}

func (v NullableAddLabelsToSourceAlphaOutput) Get() *AddLabelsToSourceAlphaOutput {
	return v.value
}

func (v *NullableAddLabelsToSourceAlphaOutput) Set(val *AddLabelsToSourceAlphaOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableAddLabelsToSourceAlphaOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableAddLabelsToSourceAlphaOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddLabelsToSourceAlphaOutput(
	val *AddLabelsToSourceAlphaOutput,
) *NullableAddLabelsToSourceAlphaOutput {
	return &NullableAddLabelsToSourceAlphaOutput{value: val, isSet: true}
}

func (v NullableAddLabelsToSourceAlphaOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddLabelsToSourceAlphaOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
