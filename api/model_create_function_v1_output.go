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

// checks if the CreateFunctionV1Output type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateFunctionV1Output{}

// CreateFunctionV1Output Create a Function.
type CreateFunctionV1Output struct {
	Function FunctionV1 `json:"function"`
}

// NewCreateFunctionV1Output instantiates a new CreateFunctionV1Output object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateFunctionV1Output(function FunctionV1) *CreateFunctionV1Output {
	this := CreateFunctionV1Output{}
	this.Function = function
	return &this
}

// NewCreateFunctionV1OutputWithDefaults instantiates a new CreateFunctionV1Output object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateFunctionV1OutputWithDefaults() *CreateFunctionV1Output {
	this := CreateFunctionV1Output{}
	return &this
}

// GetFunction returns the Function field value
func (o *CreateFunctionV1Output) GetFunction() FunctionV1 {
	if o == nil {
		var ret FunctionV1
		return ret
	}

	return o.Function
}

// GetFunctionOk returns a tuple with the Function field value
// and a boolean to check if the value has been set.
func (o *CreateFunctionV1Output) GetFunctionOk() (*FunctionV1, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Function, true
}

// SetFunction sets field value
func (o *CreateFunctionV1Output) SetFunction(v FunctionV1) {
	o.Function = v
}

func (o CreateFunctionV1Output) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateFunctionV1Output) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["function"] = o.Function
	return toSerialize, nil
}

type NullableCreateFunctionV1Output struct {
	value *CreateFunctionV1Output
	isSet bool
}

func (v NullableCreateFunctionV1Output) Get() *CreateFunctionV1Output {
	return v.value
}

func (v *NullableCreateFunctionV1Output) Set(val *CreateFunctionV1Output) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateFunctionV1Output) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateFunctionV1Output) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateFunctionV1Output(
	val *CreateFunctionV1Output,
) *NullableCreateFunctionV1Output {
	return &NullableCreateFunctionV1Output{value: val, isSet: true}
}

func (v NullableCreateFunctionV1Output) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateFunctionV1Output) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
