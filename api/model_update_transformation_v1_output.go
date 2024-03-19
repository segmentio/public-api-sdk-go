/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 45.0.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the UpdateTransformationV1Output type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateTransformationV1Output{}

// UpdateTransformationV1Output The output of an updated Transformation.
type UpdateTransformationV1Output struct {
	Transformation TransformationV1 `json:"transformation"`
}

// NewUpdateTransformationV1Output instantiates a new UpdateTransformationV1Output object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateTransformationV1Output(
	transformation TransformationV1,
) *UpdateTransformationV1Output {
	this := UpdateTransformationV1Output{}
	this.Transformation = transformation
	return &this
}

// NewUpdateTransformationV1OutputWithDefaults instantiates a new UpdateTransformationV1Output object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateTransformationV1OutputWithDefaults() *UpdateTransformationV1Output {
	this := UpdateTransformationV1Output{}
	return &this
}

// GetTransformation returns the Transformation field value
func (o *UpdateTransformationV1Output) GetTransformation() TransformationV1 {
	if o == nil {
		var ret TransformationV1
		return ret
	}

	return o.Transformation
}

// GetTransformationOk returns a tuple with the Transformation field value
// and a boolean to check if the value has been set.
func (o *UpdateTransformationV1Output) GetTransformationOk() (*TransformationV1, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Transformation, true
}

// SetTransformation sets field value
func (o *UpdateTransformationV1Output) SetTransformation(v TransformationV1) {
	o.Transformation = v
}

func (o UpdateTransformationV1Output) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateTransformationV1Output) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["transformation"] = o.Transformation
	return toSerialize, nil
}

type NullableUpdateTransformationV1Output struct {
	value *UpdateTransformationV1Output
	isSet bool
}

func (v NullableUpdateTransformationV1Output) Get() *UpdateTransformationV1Output {
	return v.value
}

func (v *NullableUpdateTransformationV1Output) Set(val *UpdateTransformationV1Output) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateTransformationV1Output) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateTransformationV1Output) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateTransformationV1Output(
	val *UpdateTransformationV1Output,
) *NullableUpdateTransformationV1Output {
	return &NullableUpdateTransformationV1Output{value: val, isSet: true}
}

func (v NullableUpdateTransformationV1Output) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateTransformationV1Output) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
