/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 56.0.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the ReplaceLabelsInSourceV1Output type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReplaceLabelsInSourceV1Output{}

// ReplaceLabelsInSourceV1Output Response from replacing a list of labels in a Source.
type ReplaceLabelsInSourceV1Output struct {
	// All labels replaced in the Source.
	Labels []LabelV1 `json:"labels"`
}

// NewReplaceLabelsInSourceV1Output instantiates a new ReplaceLabelsInSourceV1Output object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReplaceLabelsInSourceV1Output(labels []LabelV1) *ReplaceLabelsInSourceV1Output {
	this := ReplaceLabelsInSourceV1Output{}
	this.Labels = labels
	return &this
}

// NewReplaceLabelsInSourceV1OutputWithDefaults instantiates a new ReplaceLabelsInSourceV1Output object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReplaceLabelsInSourceV1OutputWithDefaults() *ReplaceLabelsInSourceV1Output {
	this := ReplaceLabelsInSourceV1Output{}
	return &this
}

// GetLabels returns the Labels field value
func (o *ReplaceLabelsInSourceV1Output) GetLabels() []LabelV1 {
	if o == nil {
		var ret []LabelV1
		return ret
	}

	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value
// and a boolean to check if the value has been set.
func (o *ReplaceLabelsInSourceV1Output) GetLabelsOk() ([]LabelV1, bool) {
	if o == nil {
		return nil, false
	}
	return o.Labels, true
}

// SetLabels sets field value
func (o *ReplaceLabelsInSourceV1Output) SetLabels(v []LabelV1) {
	o.Labels = v
}

func (o ReplaceLabelsInSourceV1Output) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReplaceLabelsInSourceV1Output) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["labels"] = o.Labels
	return toSerialize, nil
}

type NullableReplaceLabelsInSourceV1Output struct {
	value *ReplaceLabelsInSourceV1Output
	isSet bool
}

func (v NullableReplaceLabelsInSourceV1Output) Get() *ReplaceLabelsInSourceV1Output {
	return v.value
}

func (v *NullableReplaceLabelsInSourceV1Output) Set(val *ReplaceLabelsInSourceV1Output) {
	v.value = val
	v.isSet = true
}

func (v NullableReplaceLabelsInSourceV1Output) IsSet() bool {
	return v.isSet
}

func (v *NullableReplaceLabelsInSourceV1Output) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReplaceLabelsInSourceV1Output(
	val *ReplaceLabelsInSourceV1Output,
) *NullableReplaceLabelsInSourceV1Output {
	return &NullableReplaceLabelsInSourceV1Output{value: val, isSet: true}
}

func (v NullableReplaceLabelsInSourceV1Output) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReplaceLabelsInSourceV1Output) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
