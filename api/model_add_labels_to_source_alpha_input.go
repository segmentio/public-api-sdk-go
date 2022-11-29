/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 33.0.2
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// AddLabelsToSourceAlphaInput Applies an existing label to an existing Source.
type AddLabelsToSourceAlphaInput struct {
	// The labels to associate with a Source.
	Labels []LabelAlpha `json:"labels"`
}

// NewAddLabelsToSourceAlphaInput instantiates a new AddLabelsToSourceAlphaInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddLabelsToSourceAlphaInput(labels []LabelAlpha) *AddLabelsToSourceAlphaInput {
	this := AddLabelsToSourceAlphaInput{}
	this.Labels = labels
	return &this
}

// NewAddLabelsToSourceAlphaInputWithDefaults instantiates a new AddLabelsToSourceAlphaInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddLabelsToSourceAlphaInputWithDefaults() *AddLabelsToSourceAlphaInput {
	this := AddLabelsToSourceAlphaInput{}
	return &this
}

// GetLabels returns the Labels field value
func (o *AddLabelsToSourceAlphaInput) GetLabels() []LabelAlpha {
	if o == nil {
		var ret []LabelAlpha
		return ret
	}

	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value
// and a boolean to check if the value has been set.
func (o *AddLabelsToSourceAlphaInput) GetLabelsOk() ([]LabelAlpha, bool) {
	if o == nil {
		return nil, false
	}
	return o.Labels, true
}

// SetLabels sets field value
func (o *AddLabelsToSourceAlphaInput) SetLabels(v []LabelAlpha) {
	o.Labels = v
}

func (o AddLabelsToSourceAlphaInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["labels"] = o.Labels
	}
	return json.Marshal(toSerialize)
}

type NullableAddLabelsToSourceAlphaInput struct {
	value *AddLabelsToSourceAlphaInput
	isSet bool
}

func (v NullableAddLabelsToSourceAlphaInput) Get() *AddLabelsToSourceAlphaInput {
	return v.value
}

func (v *NullableAddLabelsToSourceAlphaInput) Set(val *AddLabelsToSourceAlphaInput) {
	v.value = val
	v.isSet = true
}

func (v NullableAddLabelsToSourceAlphaInput) IsSet() bool {
	return v.isSet
}

func (v *NullableAddLabelsToSourceAlphaInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddLabelsToSourceAlphaInput(
	val *AddLabelsToSourceAlphaInput,
) *NullableAddLabelsToSourceAlphaInput {
	return &NullableAddLabelsToSourceAlphaInput{value: val, isSet: true}
}

func (v NullableAddLabelsToSourceAlphaInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddLabelsToSourceAlphaInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
