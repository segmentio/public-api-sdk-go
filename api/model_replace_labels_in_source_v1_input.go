/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 58.13.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the ReplaceLabelsInSourceV1Input type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReplaceLabelsInSourceV1Input{}

// ReplaceLabelsInSourceV1Input Replaces all labels in a Source with a list of new labels.
type ReplaceLabelsInSourceV1Input struct {
	// The list of labels to replace in the Source.
	Labels []LabelV1 `json:"labels"`
}

// NewReplaceLabelsInSourceV1Input instantiates a new ReplaceLabelsInSourceV1Input object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReplaceLabelsInSourceV1Input(labels []LabelV1) *ReplaceLabelsInSourceV1Input {
	this := ReplaceLabelsInSourceV1Input{}
	this.Labels = labels
	return &this
}

// NewReplaceLabelsInSourceV1InputWithDefaults instantiates a new ReplaceLabelsInSourceV1Input object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReplaceLabelsInSourceV1InputWithDefaults() *ReplaceLabelsInSourceV1Input {
	this := ReplaceLabelsInSourceV1Input{}
	return &this
}

// GetLabels returns the Labels field value
func (o *ReplaceLabelsInSourceV1Input) GetLabels() []LabelV1 {
	if o == nil {
		var ret []LabelV1
		return ret
	}

	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value
// and a boolean to check if the value has been set.
func (o *ReplaceLabelsInSourceV1Input) GetLabelsOk() ([]LabelV1, bool) {
	if o == nil {
		return nil, false
	}
	return o.Labels, true
}

// SetLabels sets field value
func (o *ReplaceLabelsInSourceV1Input) SetLabels(v []LabelV1) {
	o.Labels = v
}

func (o ReplaceLabelsInSourceV1Input) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReplaceLabelsInSourceV1Input) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["labels"] = o.Labels
	return toSerialize, nil
}

type NullableReplaceLabelsInSourceV1Input struct {
	value *ReplaceLabelsInSourceV1Input
	isSet bool
}

func (v NullableReplaceLabelsInSourceV1Input) Get() *ReplaceLabelsInSourceV1Input {
	return v.value
}

func (v *NullableReplaceLabelsInSourceV1Input) Set(val *ReplaceLabelsInSourceV1Input) {
	v.value = val
	v.isSet = true
}

func (v NullableReplaceLabelsInSourceV1Input) IsSet() bool {
	return v.isSet
}

func (v *NullableReplaceLabelsInSourceV1Input) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReplaceLabelsInSourceV1Input(
	val *ReplaceLabelsInSourceV1Input,
) *NullableReplaceLabelsInSourceV1Input {
	return &NullableReplaceLabelsInSourceV1Input{value: val, isSet: true}
}

func (v NullableReplaceLabelsInSourceV1Input) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReplaceLabelsInSourceV1Input) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
