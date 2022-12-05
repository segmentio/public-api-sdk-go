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

// ListFunctionsV1Output Lists Functions in a Workspace.
type ListFunctionsV1Output struct {
	// An array of Functions.
	Functions  []ListFunctionItemV1 `json:"functions"`
	Pagination Pagination           `json:"pagination"`
}

// NewListFunctionsV1Output instantiates a new ListFunctionsV1Output object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListFunctionsV1Output(
	functions []ListFunctionItemV1,
	pagination Pagination,
) *ListFunctionsV1Output {
	this := ListFunctionsV1Output{}
	this.Functions = functions
	this.Pagination = pagination
	return &this
}

// NewListFunctionsV1OutputWithDefaults instantiates a new ListFunctionsV1Output object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListFunctionsV1OutputWithDefaults() *ListFunctionsV1Output {
	this := ListFunctionsV1Output{}
	return &this
}

// GetFunctions returns the Functions field value
func (o *ListFunctionsV1Output) GetFunctions() []ListFunctionItemV1 {
	if o == nil {
		var ret []ListFunctionItemV1
		return ret
	}

	return o.Functions
}

// GetFunctionsOk returns a tuple with the Functions field value
// and a boolean to check if the value has been set.
func (o *ListFunctionsV1Output) GetFunctionsOk() ([]ListFunctionItemV1, bool) {
	if o == nil {
		return nil, false
	}
	return o.Functions, true
}

// SetFunctions sets field value
func (o *ListFunctionsV1Output) SetFunctions(v []ListFunctionItemV1) {
	o.Functions = v
}

// GetPagination returns the Pagination field value
func (o *ListFunctionsV1Output) GetPagination() Pagination {
	if o == nil {
		var ret Pagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *ListFunctionsV1Output) GetPaginationOk() (*Pagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *ListFunctionsV1Output) SetPagination(v Pagination) {
	o.Pagination = v
}

func (o ListFunctionsV1Output) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["functions"] = o.Functions
	}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	return json.Marshal(toSerialize)
}

type NullableListFunctionsV1Output struct {
	value *ListFunctionsV1Output
	isSet bool
}

func (v NullableListFunctionsV1Output) Get() *ListFunctionsV1Output {
	return v.value
}

func (v *NullableListFunctionsV1Output) Set(val *ListFunctionsV1Output) {
	v.value = val
	v.isSet = true
}

func (v NullableListFunctionsV1Output) IsSet() bool {
	return v.isSet
}

func (v *NullableListFunctionsV1Output) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListFunctionsV1Output(val *ListFunctionsV1Output) *NullableListFunctionsV1Output {
	return &NullableListFunctionsV1Output{value: val, isSet: true}
}

func (v NullableListFunctionsV1Output) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListFunctionsV1Output) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
