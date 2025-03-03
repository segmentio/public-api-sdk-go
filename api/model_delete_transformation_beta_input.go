/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 57.4.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the DeleteTransformationBetaInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteTransformationBetaInput{}

// DeleteTransformationBetaInput The input of delete Transformation.
type DeleteTransformationBetaInput struct {
	// The Transformation id.
	TransformationId string `json:"transformationId"`
}

// NewDeleteTransformationBetaInput instantiates a new DeleteTransformationBetaInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteTransformationBetaInput(transformationId string) *DeleteTransformationBetaInput {
	this := DeleteTransformationBetaInput{}
	this.TransformationId = transformationId
	return &this
}

// NewDeleteTransformationBetaInputWithDefaults instantiates a new DeleteTransformationBetaInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteTransformationBetaInputWithDefaults() *DeleteTransformationBetaInput {
	this := DeleteTransformationBetaInput{}
	return &this
}

// GetTransformationId returns the TransformationId field value
func (o *DeleteTransformationBetaInput) GetTransformationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransformationId
}

// GetTransformationIdOk returns a tuple with the TransformationId field value
// and a boolean to check if the value has been set.
func (o *DeleteTransformationBetaInput) GetTransformationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransformationId, true
}

// SetTransformationId sets field value
func (o *DeleteTransformationBetaInput) SetTransformationId(v string) {
	o.TransformationId = v
}

func (o DeleteTransformationBetaInput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteTransformationBetaInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["transformationId"] = o.TransformationId
	return toSerialize, nil
}

type NullableDeleteTransformationBetaInput struct {
	value *DeleteTransformationBetaInput
	isSet bool
}

func (v NullableDeleteTransformationBetaInput) Get() *DeleteTransformationBetaInput {
	return v.value
}

func (v *NullableDeleteTransformationBetaInput) Set(val *DeleteTransformationBetaInput) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteTransformationBetaInput) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteTransformationBetaInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteTransformationBetaInput(
	val *DeleteTransformationBetaInput,
) *NullableDeleteTransformationBetaInput {
	return &NullableDeleteTransformationBetaInput{value: val, isSet: true}
}

func (v NullableDeleteTransformationBetaInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteTransformationBetaInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
