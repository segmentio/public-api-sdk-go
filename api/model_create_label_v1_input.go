/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 57.3.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the CreateLabelV1Input type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateLabelV1Input{}

// CreateLabelV1Input Creates a new label in the current Workspace.
type CreateLabelV1Input struct {
	Label LabelV1 `json:"label"`
}

// NewCreateLabelV1Input instantiates a new CreateLabelV1Input object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateLabelV1Input(label LabelV1) *CreateLabelV1Input {
	this := CreateLabelV1Input{}
	this.Label = label
	return &this
}

// NewCreateLabelV1InputWithDefaults instantiates a new CreateLabelV1Input object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateLabelV1InputWithDefaults() *CreateLabelV1Input {
	this := CreateLabelV1Input{}
	return &this
}

// GetLabel returns the Label field value
func (o *CreateLabelV1Input) GetLabel() LabelV1 {
	if o == nil {
		var ret LabelV1
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *CreateLabelV1Input) GetLabelOk() (*LabelV1, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *CreateLabelV1Input) SetLabel(v LabelV1) {
	o.Label = v
}

func (o CreateLabelV1Input) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateLabelV1Input) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["label"] = o.Label
	return toSerialize, nil
}

type NullableCreateLabelV1Input struct {
	value *CreateLabelV1Input
	isSet bool
}

func (v NullableCreateLabelV1Input) Get() *CreateLabelV1Input {
	return v.value
}

func (v *NullableCreateLabelV1Input) Set(val *CreateLabelV1Input) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateLabelV1Input) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateLabelV1Input) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateLabelV1Input(val *CreateLabelV1Input) *NullableCreateLabelV1Input {
	return &NullableCreateLabelV1Input{value: val, isSet: true}
}

func (v NullableCreateLabelV1Input) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateLabelV1Input) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
