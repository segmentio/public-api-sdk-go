/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 34.3.5
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// CreateLabelAlphaInput Creates a new label in the current Workspace.
type CreateLabelAlphaInput struct {
	Label Label `json:"label"`
}

// NewCreateLabelAlphaInput instantiates a new CreateLabelAlphaInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateLabelAlphaInput(label Label) *CreateLabelAlphaInput {
	this := CreateLabelAlphaInput{}
	this.Label = label
	return &this
}

// NewCreateLabelAlphaInputWithDefaults instantiates a new CreateLabelAlphaInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateLabelAlphaInputWithDefaults() *CreateLabelAlphaInput {
	this := CreateLabelAlphaInput{}
	return &this
}

// GetLabel returns the Label field value
func (o *CreateLabelAlphaInput) GetLabel() Label {
	if o == nil {
		var ret Label
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *CreateLabelAlphaInput) GetLabelOk() (*Label, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *CreateLabelAlphaInput) SetLabel(v Label) {
	o.Label = v
}

func (o CreateLabelAlphaInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["label"] = o.Label
	}
	return json.Marshal(toSerialize)
}

type NullableCreateLabelAlphaInput struct {
	value *CreateLabelAlphaInput
	isSet bool
}

func (v NullableCreateLabelAlphaInput) Get() *CreateLabelAlphaInput {
	return v.value
}

func (v *NullableCreateLabelAlphaInput) Set(val *CreateLabelAlphaInput) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateLabelAlphaInput) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateLabelAlphaInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateLabelAlphaInput(val *CreateLabelAlphaInput) *NullableCreateLabelAlphaInput {
	return &NullableCreateLabelAlphaInput{value: val, isSet: true}
}

func (v NullableCreateLabelAlphaInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateLabelAlphaInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
