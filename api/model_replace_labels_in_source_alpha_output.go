/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 34.0.1
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// ReplaceLabelsInSourceAlphaOutput Response from replacing a list of labels in a Source.
type ReplaceLabelsInSourceAlphaOutput struct {
	// All labels replaced in the Source.
	Labels []LabelAlpha `json:"labels"`
}

// NewReplaceLabelsInSourceAlphaOutput instantiates a new ReplaceLabelsInSourceAlphaOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReplaceLabelsInSourceAlphaOutput(labels []LabelAlpha) *ReplaceLabelsInSourceAlphaOutput {
	this := ReplaceLabelsInSourceAlphaOutput{}
	this.Labels = labels
	return &this
}

// NewReplaceLabelsInSourceAlphaOutputWithDefaults instantiates a new ReplaceLabelsInSourceAlphaOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReplaceLabelsInSourceAlphaOutputWithDefaults() *ReplaceLabelsInSourceAlphaOutput {
	this := ReplaceLabelsInSourceAlphaOutput{}
	return &this
}

// GetLabels returns the Labels field value
func (o *ReplaceLabelsInSourceAlphaOutput) GetLabels() []LabelAlpha {
	if o == nil {
		var ret []LabelAlpha
		return ret
	}

	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value
// and a boolean to check if the value has been set.
func (o *ReplaceLabelsInSourceAlphaOutput) GetLabelsOk() ([]LabelAlpha, bool) {
	if o == nil {
		return nil, false
	}
	return o.Labels, true
}

// SetLabels sets field value
func (o *ReplaceLabelsInSourceAlphaOutput) SetLabels(v []LabelAlpha) {
	o.Labels = v
}

func (o ReplaceLabelsInSourceAlphaOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["labels"] = o.Labels
	}
	return json.Marshal(toSerialize)
}

type NullableReplaceLabelsInSourceAlphaOutput struct {
	value *ReplaceLabelsInSourceAlphaOutput
	isSet bool
}

func (v NullableReplaceLabelsInSourceAlphaOutput) Get() *ReplaceLabelsInSourceAlphaOutput {
	return v.value
}

func (v *NullableReplaceLabelsInSourceAlphaOutput) Set(val *ReplaceLabelsInSourceAlphaOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableReplaceLabelsInSourceAlphaOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableReplaceLabelsInSourceAlphaOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReplaceLabelsInSourceAlphaOutput(
	val *ReplaceLabelsInSourceAlphaOutput,
) *NullableReplaceLabelsInSourceAlphaOutput {
	return &NullableReplaceLabelsInSourceAlphaOutput{value: val, isSet: true}
}

func (v NullableReplaceLabelsInSourceAlphaOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReplaceLabelsInSourceAlphaOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
