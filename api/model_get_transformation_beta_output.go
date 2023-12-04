/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 38.2.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the GetTransformationBetaOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTransformationBetaOutput{}

// GetTransformationBetaOutput The output of Transformation retrieval.
type GetTransformationBetaOutput struct {
	Transformation TransformationBeta `json:"transformation"`
}

// NewGetTransformationBetaOutput instantiates a new GetTransformationBetaOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTransformationBetaOutput(
	transformation TransformationBeta,
) *GetTransformationBetaOutput {
	this := GetTransformationBetaOutput{}
	this.Transformation = transformation
	return &this
}

// NewGetTransformationBetaOutputWithDefaults instantiates a new GetTransformationBetaOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTransformationBetaOutputWithDefaults() *GetTransformationBetaOutput {
	this := GetTransformationBetaOutput{}
	return &this
}

// GetTransformation returns the Transformation field value
func (o *GetTransformationBetaOutput) GetTransformation() TransformationBeta {
	if o == nil {
		var ret TransformationBeta
		return ret
	}

	return o.Transformation
}

// GetTransformationOk returns a tuple with the Transformation field value
// and a boolean to check if the value has been set.
func (o *GetTransformationBetaOutput) GetTransformationOk() (*TransformationBeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Transformation, true
}

// SetTransformation sets field value
func (o *GetTransformationBetaOutput) SetTransformation(v TransformationBeta) {
	o.Transformation = v
}

func (o GetTransformationBetaOutput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTransformationBetaOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["transformation"] = o.Transformation
	return toSerialize, nil
}

type NullableGetTransformationBetaOutput struct {
	value *GetTransformationBetaOutput
	isSet bool
}

func (v NullableGetTransformationBetaOutput) Get() *GetTransformationBetaOutput {
	return v.value
}

func (v *NullableGetTransformationBetaOutput) Set(val *GetTransformationBetaOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTransformationBetaOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTransformationBetaOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTransformationBetaOutput(
	val *GetTransformationBetaOutput,
) *NullableGetTransformationBetaOutput {
	return &NullableGetTransformationBetaOutput{value: val, isSet: true}
}

func (v NullableGetTransformationBetaOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTransformationBetaOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
