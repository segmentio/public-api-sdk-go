/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 49.0.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the CreateLabelV1Output type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateLabelV1Output{}

// CreateLabelV1Output Result of creating a new label in the current Workspace.
type CreateLabelV1Output struct {
	Label LabelV1 `json:"label"`
}

// NewCreateLabelV1Output instantiates a new CreateLabelV1Output object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateLabelV1Output(label LabelV1) *CreateLabelV1Output {
	this := CreateLabelV1Output{}
	this.Label = label
	return &this
}

// NewCreateLabelV1OutputWithDefaults instantiates a new CreateLabelV1Output object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateLabelV1OutputWithDefaults() *CreateLabelV1Output {
	this := CreateLabelV1Output{}
	return &this
}

// GetLabel returns the Label field value
func (o *CreateLabelV1Output) GetLabel() LabelV1 {
	if o == nil {
		var ret LabelV1
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *CreateLabelV1Output) GetLabelOk() (*LabelV1, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *CreateLabelV1Output) SetLabel(v LabelV1) {
	o.Label = v
}

func (o CreateLabelV1Output) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateLabelV1Output) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["label"] = o.Label
	return toSerialize, nil
}

type NullableCreateLabelV1Output struct {
	value *CreateLabelV1Output
	isSet bool
}

func (v NullableCreateLabelV1Output) Get() *CreateLabelV1Output {
	return v.value
}

func (v *NullableCreateLabelV1Output) Set(val *CreateLabelV1Output) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateLabelV1Output) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateLabelV1Output) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateLabelV1Output(val *CreateLabelV1Output) *NullableCreateLabelV1Output {
	return &NullableCreateLabelV1Output{value: val, isSet: true}
}

func (v NullableCreateLabelV1Output) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateLabelV1Output) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
