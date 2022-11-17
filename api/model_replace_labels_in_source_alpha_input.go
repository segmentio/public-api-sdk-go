/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API. 

API version: 32.0.7
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// ReplaceLabelsInSourceAlphaInput Replaces all labels in a Source with a list of new labels.
type ReplaceLabelsInSourceAlphaInput struct {
	// The list of labels to replace in the Source.
	Labels []LabelAlpha `json:"labels"`
}

// NewReplaceLabelsInSourceAlphaInput instantiates a new ReplaceLabelsInSourceAlphaInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReplaceLabelsInSourceAlphaInput(labels []LabelAlpha) *ReplaceLabelsInSourceAlphaInput {
	this := ReplaceLabelsInSourceAlphaInput{}
	this.Labels = labels
	return &this
}

// NewReplaceLabelsInSourceAlphaInputWithDefaults instantiates a new ReplaceLabelsInSourceAlphaInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReplaceLabelsInSourceAlphaInputWithDefaults() *ReplaceLabelsInSourceAlphaInput {
	this := ReplaceLabelsInSourceAlphaInput{}
	return &this
}

// GetLabels returns the Labels field value
func (o *ReplaceLabelsInSourceAlphaInput) GetLabels() []LabelAlpha {
	if o == nil {
		var ret []LabelAlpha
		return ret
	}

	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value
// and a boolean to check if the value has been set.
func (o *ReplaceLabelsInSourceAlphaInput) GetLabelsOk() ([]LabelAlpha, bool) {
	if o == nil {
		return nil, false
	}
	return o.Labels, true
}

// SetLabels sets field value
func (o *ReplaceLabelsInSourceAlphaInput) SetLabels(v []LabelAlpha) {
	o.Labels = v
}

func (o ReplaceLabelsInSourceAlphaInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["labels"] = o.Labels
	}
	return json.Marshal(toSerialize)
}

type NullableReplaceLabelsInSourceAlphaInput struct {
	value *ReplaceLabelsInSourceAlphaInput
	isSet bool
}

func (v NullableReplaceLabelsInSourceAlphaInput) Get() *ReplaceLabelsInSourceAlphaInput {
	return v.value
}

func (v *NullableReplaceLabelsInSourceAlphaInput) Set(val *ReplaceLabelsInSourceAlphaInput) {
	v.value = val
	v.isSet = true
}

func (v NullableReplaceLabelsInSourceAlphaInput) IsSet() bool {
	return v.isSet
}

func (v *NullableReplaceLabelsInSourceAlphaInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReplaceLabelsInSourceAlphaInput(val *ReplaceLabelsInSourceAlphaInput) *NullableReplaceLabelsInSourceAlphaInput {
	return &NullableReplaceLabelsInSourceAlphaInput{value: val, isSet: true}
}

func (v NullableReplaceLabelsInSourceAlphaInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReplaceLabelsInSourceAlphaInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


