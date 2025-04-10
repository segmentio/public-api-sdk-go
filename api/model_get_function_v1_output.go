/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 58.1.1
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the GetFunctionV1Output type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetFunctionV1Output{}

// GetFunctionV1Output Gets a single Function.
type GetFunctionV1Output struct {
	Function NullableFunctionV1 `json:"function"`
}

// NewGetFunctionV1Output instantiates a new GetFunctionV1Output object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFunctionV1Output(function NullableFunctionV1) *GetFunctionV1Output {
	this := GetFunctionV1Output{}
	this.Function = function
	return &this
}

// NewGetFunctionV1OutputWithDefaults instantiates a new GetFunctionV1Output object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFunctionV1OutputWithDefaults() *GetFunctionV1Output {
	this := GetFunctionV1Output{}
	return &this
}

// GetFunction returns the Function field value
// If the value is explicit nil, the zero value for FunctionV1 will be returned
func (o *GetFunctionV1Output) GetFunction() FunctionV1 {
	if o == nil || o.Function.Get() == nil {
		var ret FunctionV1
		return ret
	}

	return *o.Function.Get()
}

// GetFunctionOk returns a tuple with the Function field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetFunctionV1Output) GetFunctionOk() (*FunctionV1, bool) {
	if o == nil {
		return nil, false
	}
	return o.Function.Get(), o.Function.IsSet()
}

// SetFunction sets field value
func (o *GetFunctionV1Output) SetFunction(v FunctionV1) {
	o.Function.Set(&v)
}

func (o GetFunctionV1Output) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetFunctionV1Output) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["function"] = o.Function.Get()
	return toSerialize, nil
}

type NullableGetFunctionV1Output struct {
	value *GetFunctionV1Output
	isSet bool
}

func (v NullableGetFunctionV1Output) Get() *GetFunctionV1Output {
	return v.value
}

func (v *NullableGetFunctionV1Output) Set(val *GetFunctionV1Output) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFunctionV1Output) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFunctionV1Output) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFunctionV1Output(val *GetFunctionV1Output) *NullableGetFunctionV1Output {
	return &NullableGetFunctionV1Output{value: val, isSet: true}
}

func (v NullableGetFunctionV1Output) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFunctionV1Output) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
