/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API. 

API version: 33.0.3
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// UpdateTransformationBetaOutput The output of an updated Transformation.
type UpdateTransformationBetaOutput struct {
	Transformation Transformation1 `json:"transformation"`
}

// NewUpdateTransformationBetaOutput instantiates a new UpdateTransformationBetaOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateTransformationBetaOutput(transformation Transformation1) *UpdateTransformationBetaOutput {
	this := UpdateTransformationBetaOutput{}
	this.Transformation = transformation
	return &this
}

// NewUpdateTransformationBetaOutputWithDefaults instantiates a new UpdateTransformationBetaOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateTransformationBetaOutputWithDefaults() *UpdateTransformationBetaOutput {
	this := UpdateTransformationBetaOutput{}
	return &this
}

// GetTransformation returns the Transformation field value
func (o *UpdateTransformationBetaOutput) GetTransformation() Transformation1 {
	if o == nil {
		var ret Transformation1
		return ret
	}

	return o.Transformation
}

// GetTransformationOk returns a tuple with the Transformation field value
// and a boolean to check if the value has been set.
func (o *UpdateTransformationBetaOutput) GetTransformationOk() (*Transformation1, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Transformation, true
}

// SetTransformation sets field value
func (o *UpdateTransformationBetaOutput) SetTransformation(v Transformation1) {
	o.Transformation = v
}

func (o UpdateTransformationBetaOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["transformation"] = o.Transformation
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateTransformationBetaOutput struct {
	value *UpdateTransformationBetaOutput
	isSet bool
}

func (v NullableUpdateTransformationBetaOutput) Get() *UpdateTransformationBetaOutput {
	return v.value
}

func (v *NullableUpdateTransformationBetaOutput) Set(val *UpdateTransformationBetaOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateTransformationBetaOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateTransformationBetaOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateTransformationBetaOutput(val *UpdateTransformationBetaOutput) *NullableUpdateTransformationBetaOutput {
	return &NullableUpdateTransformationBetaOutput{value: val, isSet: true}
}

func (v NullableUpdateTransformationBetaOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateTransformationBetaOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


