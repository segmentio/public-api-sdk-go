/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 54.2.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the AddLabelsToSourceV1Input type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddLabelsToSourceV1Input{}

// AddLabelsToSourceV1Input Applies an existing label to an existing Source.
type AddLabelsToSourceV1Input struct {
	// The labels to associate with a Source.
	Labels []LabelV1 `json:"labels"`
}

// NewAddLabelsToSourceV1Input instantiates a new AddLabelsToSourceV1Input object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddLabelsToSourceV1Input(labels []LabelV1) *AddLabelsToSourceV1Input {
	this := AddLabelsToSourceV1Input{}
	this.Labels = labels
	return &this
}

// NewAddLabelsToSourceV1InputWithDefaults instantiates a new AddLabelsToSourceV1Input object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddLabelsToSourceV1InputWithDefaults() *AddLabelsToSourceV1Input {
	this := AddLabelsToSourceV1Input{}
	return &this
}

// GetLabels returns the Labels field value
func (o *AddLabelsToSourceV1Input) GetLabels() []LabelV1 {
	if o == nil {
		var ret []LabelV1
		return ret
	}

	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value
// and a boolean to check if the value has been set.
func (o *AddLabelsToSourceV1Input) GetLabelsOk() ([]LabelV1, bool) {
	if o == nil {
		return nil, false
	}
	return o.Labels, true
}

// SetLabels sets field value
func (o *AddLabelsToSourceV1Input) SetLabels(v []LabelV1) {
	o.Labels = v
}

func (o AddLabelsToSourceV1Input) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddLabelsToSourceV1Input) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["labels"] = o.Labels
	return toSerialize, nil
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

func NewNullableAddLabelsToSourceV1Input(
	val *AddLabelsToSourceV1Input,
) *NullableAddLabelsToSourceV1Input {
	return &NullableAddLabelsToSourceV1Input{value: val, isSet: true}
}

func (v NullableAddLabelsToSourceV1Input) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddLabelsToSourceV1Input) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
