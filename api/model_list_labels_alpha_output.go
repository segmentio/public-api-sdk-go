/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 37.0.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// ListLabelsAlphaOutput Returns all available labels for the current Workspace.
type ListLabelsAlphaOutput struct {
	// All labels associated with the current Workspace.
	Labels []LabelAlpha `json:"labels"`
}

// NewListLabelsAlphaOutput instantiates a new ListLabelsAlphaOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListLabelsAlphaOutput(labels []LabelAlpha) *ListLabelsAlphaOutput {
	this := ListLabelsAlphaOutput{}
	this.Labels = labels
	return &this
}

// NewListLabelsAlphaOutputWithDefaults instantiates a new ListLabelsAlphaOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListLabelsAlphaOutputWithDefaults() *ListLabelsAlphaOutput {
	this := ListLabelsAlphaOutput{}
	return &this
}

// GetLabels returns the Labels field value
func (o *ListLabelsAlphaOutput) GetLabels() []LabelAlpha {
	if o == nil {
		var ret []LabelAlpha
		return ret
	}

	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value
// and a boolean to check if the value has been set.
func (o *ListLabelsAlphaOutput) GetLabelsOk() ([]LabelAlpha, bool) {
	if o == nil {
		return nil, false
	}
	return o.Labels, true
}

// SetLabels sets field value
func (o *ListLabelsAlphaOutput) SetLabels(v []LabelAlpha) {
	o.Labels = v
}

func (o ListLabelsAlphaOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["labels"] = o.Labels
	}
	return json.Marshal(toSerialize)
}

type NullableListLabelsAlphaOutput struct {
	value *ListLabelsAlphaOutput
	isSet bool
}

func (v NullableListLabelsAlphaOutput) Get() *ListLabelsAlphaOutput {
	return v.value
}

func (v *NullableListLabelsAlphaOutput) Set(val *ListLabelsAlphaOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableListLabelsAlphaOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableListLabelsAlphaOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListLabelsAlphaOutput(val *ListLabelsAlphaOutput) *NullableListLabelsAlphaOutput {
	return &NullableListLabelsAlphaOutput{value: val, isSet: true}
}

func (v NullableListLabelsAlphaOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListLabelsAlphaOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
