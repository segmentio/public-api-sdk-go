/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 33.0.4
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// CreateLabelAlphaOutput Result of creating a new label in the current Workspace.
type CreateLabelAlphaOutput struct {
	// All labels associated with the current Workspace.
	Labels []LabelAlpha `json:"labels"`
}

// NewCreateLabelAlphaOutput instantiates a new CreateLabelAlphaOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateLabelAlphaOutput(labels []LabelAlpha) *CreateLabelAlphaOutput {
	this := CreateLabelAlphaOutput{}
	this.Labels = labels
	return &this
}

// NewCreateLabelAlphaOutputWithDefaults instantiates a new CreateLabelAlphaOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateLabelAlphaOutputWithDefaults() *CreateLabelAlphaOutput {
	this := CreateLabelAlphaOutput{}
	return &this
}

// GetLabels returns the Labels field value
func (o *CreateLabelAlphaOutput) GetLabels() []LabelAlpha {
	if o == nil {
		var ret []LabelAlpha
		return ret
	}

	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value
// and a boolean to check if the value has been set.
func (o *CreateLabelAlphaOutput) GetLabelsOk() ([]LabelAlpha, bool) {
	if o == nil {
		return nil, false
	}
	return o.Labels, true
}

// SetLabels sets field value
func (o *CreateLabelAlphaOutput) SetLabels(v []LabelAlpha) {
	o.Labels = v
}

func (o CreateLabelAlphaOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["labels"] = o.Labels
	}
	return json.Marshal(toSerialize)
}

type NullableCreateLabelAlphaOutput struct {
	value *CreateLabelAlphaOutput
	isSet bool
}

func (v NullableCreateLabelAlphaOutput) Get() *CreateLabelAlphaOutput {
	return v.value
}

func (v *NullableCreateLabelAlphaOutput) Set(val *CreateLabelAlphaOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateLabelAlphaOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateLabelAlphaOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateLabelAlphaOutput(
	val *CreateLabelAlphaOutput,
) *NullableCreateLabelAlphaOutput {
	return &NullableCreateLabelAlphaOutput{value: val, isSet: true}
}

func (v NullableCreateLabelAlphaOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateLabelAlphaOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
