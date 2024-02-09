/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 42.0.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the GetTransformationBetaInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTransformationBetaInput{}

// GetTransformationBetaInput The input of Transformation retrieval.
type GetTransformationBetaInput struct {
	// The Transformation id.
	TransformationId string `json:"transformationId"`
}

// NewGetTransformationBetaInput instantiates a new GetTransformationBetaInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTransformationBetaInput(transformationId string) *GetTransformationBetaInput {
	this := GetTransformationBetaInput{}
	this.TransformationId = transformationId
	return &this
}

// NewGetTransformationBetaInputWithDefaults instantiates a new GetTransformationBetaInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTransformationBetaInputWithDefaults() *GetTransformationBetaInput {
	this := GetTransformationBetaInput{}
	return &this
}

// GetTransformationId returns the TransformationId field value
func (o *GetTransformationBetaInput) GetTransformationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransformationId
}

// GetTransformationIdOk returns a tuple with the TransformationId field value
// and a boolean to check if the value has been set.
func (o *GetTransformationBetaInput) GetTransformationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransformationId, true
}

// SetTransformationId sets field value
func (o *GetTransformationBetaInput) SetTransformationId(v string) {
	o.TransformationId = v
}

func (o GetTransformationBetaInput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTransformationBetaInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["transformationId"] = o.TransformationId
	return toSerialize, nil
}

type NullableGetTransformationBetaInput struct {
	value *GetTransformationBetaInput
	isSet bool
}

func (v NullableGetTransformationBetaInput) Get() *GetTransformationBetaInput {
	return v.value
}

func (v *NullableGetTransformationBetaInput) Set(val *GetTransformationBetaInput) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTransformationBetaInput) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTransformationBetaInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTransformationBetaInput(
	val *GetTransformationBetaInput,
) *NullableGetTransformationBetaInput {
	return &NullableGetTransformationBetaInput{value: val, isSet: true}
}

func (v NullableGetTransformationBetaInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTransformationBetaInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
