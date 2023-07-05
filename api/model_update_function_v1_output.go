/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 36.0.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// UpdateFunctionV1Output Create a Function.
type UpdateFunctionV1Output struct {
	Function Function2 `json:"function"`
}

// NewUpdateFunctionV1Output instantiates a new UpdateFunctionV1Output object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateFunctionV1Output(function Function2) *UpdateFunctionV1Output {
	this := UpdateFunctionV1Output{}
	this.Function = function
	return &this
}

// NewUpdateFunctionV1OutputWithDefaults instantiates a new UpdateFunctionV1Output object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateFunctionV1OutputWithDefaults() *UpdateFunctionV1Output {
	this := UpdateFunctionV1Output{}
	return &this
}

// GetFunction returns the Function field value
func (o *UpdateFunctionV1Output) GetFunction() Function2 {
	if o == nil {
		var ret Function2
		return ret
	}

	return o.Function
}

// GetFunctionOk returns a tuple with the Function field value
// and a boolean to check if the value has been set.
func (o *UpdateFunctionV1Output) GetFunctionOk() (*Function2, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Function, true
}

// SetFunction sets field value
func (o *UpdateFunctionV1Output) SetFunction(v Function2) {
	o.Function = v
}

func (o UpdateFunctionV1Output) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["function"] = o.Function
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateFunctionV1Output struct {
	value *UpdateFunctionV1Output
	isSet bool
}

func (v NullableUpdateFunctionV1Output) Get() *UpdateFunctionV1Output {
	return v.value
}

func (v *NullableUpdateFunctionV1Output) Set(val *UpdateFunctionV1Output) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateFunctionV1Output) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateFunctionV1Output) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateFunctionV1Output(
	val *UpdateFunctionV1Output,
) *NullableUpdateFunctionV1Output {
	return &NullableUpdateFunctionV1Output{value: val, isSet: true}
}

func (v NullableUpdateFunctionV1Output) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateFunctionV1Output) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
