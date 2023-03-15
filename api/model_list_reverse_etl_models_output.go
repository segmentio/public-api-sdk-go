/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 34.3.3
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// ListReverseEtlModelsOutput Defines the result of listing Models.
type ListReverseEtlModelsOutput struct {
	// A list of Models that belong to the Workspace.
	Models     []ReverseEtlModel `json:"models"`
	Pagination Pagination        `json:"pagination"`
}

// NewListReverseEtlModelsOutput instantiates a new ListReverseEtlModelsOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListReverseEtlModelsOutput(
	models []ReverseEtlModel,
	pagination Pagination,
) *ListReverseEtlModelsOutput {
	this := ListReverseEtlModelsOutput{}
	this.Models = models
	this.Pagination = pagination
	return &this
}

// NewListReverseEtlModelsOutputWithDefaults instantiates a new ListReverseEtlModelsOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListReverseEtlModelsOutputWithDefaults() *ListReverseEtlModelsOutput {
	this := ListReverseEtlModelsOutput{}
	return &this
}

// GetModels returns the Models field value
func (o *ListReverseEtlModelsOutput) GetModels() []ReverseEtlModel {
	if o == nil {
		var ret []ReverseEtlModel
		return ret
	}

	return o.Models
}

// GetModelsOk returns a tuple with the Models field value
// and a boolean to check if the value has been set.
func (o *ListReverseEtlModelsOutput) GetModelsOk() ([]ReverseEtlModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.Models, true
}

// SetModels sets field value
func (o *ListReverseEtlModelsOutput) SetModels(v []ReverseEtlModel) {
	o.Models = v
}

// GetPagination returns the Pagination field value
func (o *ListReverseEtlModelsOutput) GetPagination() Pagination {
	if o == nil {
		var ret Pagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *ListReverseEtlModelsOutput) GetPaginationOk() (*Pagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *ListReverseEtlModelsOutput) SetPagination(v Pagination) {
	o.Pagination = v
}

func (o ListReverseEtlModelsOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["models"] = o.Models
	}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	return json.Marshal(toSerialize)
}

type NullableListReverseEtlModelsOutput struct {
	value *ListReverseEtlModelsOutput
	isSet bool
}

func (v NullableListReverseEtlModelsOutput) Get() *ListReverseEtlModelsOutput {
	return v.value
}

func (v *NullableListReverseEtlModelsOutput) Set(val *ListReverseEtlModelsOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableListReverseEtlModelsOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableListReverseEtlModelsOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListReverseEtlModelsOutput(
	val *ListReverseEtlModelsOutput,
) *NullableListReverseEtlModelsOutput {
	return &NullableListReverseEtlModelsOutput{value: val, isSet: true}
}

func (v NullableListReverseEtlModelsOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListReverseEtlModelsOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
