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

// GetTransformationV1Output The output of Transformation retrieval.
type GetTransformationV1Output struct {
	Transformation Transformation3 `json:"transformation"`
}

// NewGetTransformationV1Output instantiates a new GetTransformationV1Output object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTransformationV1Output(transformation Transformation3) *GetTransformationV1Output {
	this := GetTransformationV1Output{}
	this.Transformation = transformation
	return &this
}

// NewGetTransformationV1OutputWithDefaults instantiates a new GetTransformationV1Output object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTransformationV1OutputWithDefaults() *GetTransformationV1Output {
	this := GetTransformationV1Output{}
	return &this
}

// GetTransformation returns the Transformation field value
func (o *GetTransformationV1Output) GetTransformation() Transformation3 {
	if o == nil {
		var ret Transformation3
		return ret
	}

	return o.Transformation
}

// GetTransformationOk returns a tuple with the Transformation field value
// and a boolean to check if the value has been set.
func (o *GetTransformationV1Output) GetTransformationOk() (*Transformation3, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Transformation, true
}

// SetTransformation sets field value
func (o *GetTransformationV1Output) SetTransformation(v Transformation3) {
	o.Transformation = v
}

func (o GetTransformationV1Output) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["transformation"] = o.Transformation
	}
	return json.Marshal(toSerialize)
}

type NullableGetTransformationV1Output struct {
	value *GetTransformationV1Output
	isSet bool
}

func (v NullableGetTransformationV1Output) Get() *GetTransformationV1Output {
	return v.value
}

func (v *NullableGetTransformationV1Output) Set(val *GetTransformationV1Output) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTransformationV1Output) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTransformationV1Output) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTransformationV1Output(
	val *GetTransformationV1Output,
) *NullableGetTransformationV1Output {
	return &NullableGetTransformationV1Output{value: val, isSet: true}
}

func (v NullableGetTransformationV1Output) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTransformationV1Output) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
