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

// AddLabelsToSourceV1Input Applies an existing label to an existing Source.
type AddLabelsToSourceV1Input struct {
	// The labels to associate with a Source.
	Labels []LabelV1 `json:"labels,omitempty"`
}

// NewAddLabelsToSourceV1Input instantiates a new AddLabelsToSourceV1Input object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddLabelsToSourceV1Input() *AddLabelsToSourceV1Input {
	this := AddLabelsToSourceV1Input{}
	return &this
}

// NewAddLabelsToSourceV1InputWithDefaults instantiates a new AddLabelsToSourceV1Input object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddLabelsToSourceV1InputWithDefaults() *AddLabelsToSourceV1Input {
	this := AddLabelsToSourceV1Input{}
	return &this
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *AddLabelsToSourceV1Input) GetLabels() []LabelV1 {
	if o == nil || o.Labels == nil {
		var ret []LabelV1
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddLabelsToSourceV1Input) GetLabelsOk() ([]LabelV1, bool) {
	if o == nil || o.Labels == nil {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *AddLabelsToSourceV1Input) HasLabels() bool {
	if o != nil && o.Labels != nil {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []LabelV1 and assigns it to the Labels field.
func (o *AddLabelsToSourceV1Input) SetLabels(v []LabelV1) {
	o.Labels = v
}

func (o AddLabelsToSourceV1Input) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Labels != nil {
		toSerialize["labels"] = o.Labels
	}
	return json.Marshal(toSerialize)
}

type NullableAddLabelsToSourceV1Input struct {
	value *AddLabelsToSourceV1Input
	isSet bool
}

func (v NullableAddLabelsToSourceV1Input) Get() *AddLabelsToSourceV1Input {
	return v.value
}

func (v *NullableAddLabelsToSourceV1Input) Set(val *AddLabelsToSourceV1Input) {
	v.value = val
	v.isSet = true
}

func (v NullableAddLabelsToSourceV1Input) IsSet() bool {
	return v.isSet
}

func (v *NullableAddLabelsToSourceV1Input) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddLabelsToSourceV1Input(val *AddLabelsToSourceV1Input) *NullableAddLabelsToSourceV1Input {
	return &NullableAddLabelsToSourceV1Input{value: val, isSet: true}
}

func (v NullableAddLabelsToSourceV1Input) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddLabelsToSourceV1Input) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


