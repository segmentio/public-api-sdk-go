/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 36.3.1
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// CreateTransformationBetaOutput The output of a created Transformation.
type CreateTransformationBetaOutput struct {
	Transformation Transformation2 `json:"transformation"`
}

// NewCreateTransformationBetaOutput instantiates a new CreateTransformationBetaOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateTransformationBetaOutput(
	transformation Transformation2,
) *CreateTransformationBetaOutput {
	this := CreateTransformationBetaOutput{}
	this.Transformation = transformation
	return &this
}

// NewCreateTransformationBetaOutputWithDefaults instantiates a new CreateTransformationBetaOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateTransformationBetaOutputWithDefaults() *CreateTransformationBetaOutput {
	this := CreateTransformationBetaOutput{}
	return &this
}

// GetTransformation returns the Transformation field value
func (o *CreateTransformationBetaOutput) GetTransformation() Transformation2 {
	if o == nil {
		var ret Transformation2
		return ret
	}

	return o.Transformation
}

// GetTransformationOk returns a tuple with the Transformation field value
// and a boolean to check if the value has been set.
func (o *CreateTransformationBetaOutput) GetTransformationOk() (*Transformation2, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Transformation, true
}

// SetTransformation sets field value
func (o *CreateTransformationBetaOutput) SetTransformation(v Transformation2) {
	o.Transformation = v
}

func (o CreateTransformationBetaOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["transformation"] = o.Transformation
	}
	return json.Marshal(toSerialize)
}

type NullableCreateTransformationBetaOutput struct {
	value *CreateTransformationBetaOutput
	isSet bool
}

func (v NullableCreateTransformationBetaOutput) Get() *CreateTransformationBetaOutput {
	return v.value
}

func (v *NullableCreateTransformationBetaOutput) Set(val *CreateTransformationBetaOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateTransformationBetaOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateTransformationBetaOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateTransformationBetaOutput(
	val *CreateTransformationBetaOutput,
) *NullableCreateTransformationBetaOutput {
	return &NullableCreateTransformationBetaOutput{value: val, isSet: true}
}

func (v NullableCreateTransformationBetaOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateTransformationBetaOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
