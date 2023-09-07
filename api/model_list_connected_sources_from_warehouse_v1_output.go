/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 36.2.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// ListConnectedSourcesFromWarehouseV1Output Returns a list of Sources connected to a Warehouse.
type ListConnectedSourcesFromWarehouseV1Output struct {
	// A list that contains the Sources connected to the requested Warehouse.
	Sources    []SourceV1 `json:"sources"`
	Pagination Pagination `json:"pagination"`
}

// NewListConnectedSourcesFromWarehouseV1Output instantiates a new ListConnectedSourcesFromWarehouseV1Output object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListConnectedSourcesFromWarehouseV1Output(
	sources []SourceV1,
	pagination Pagination,
) *ListConnectedSourcesFromWarehouseV1Output {
	this := ListConnectedSourcesFromWarehouseV1Output{}
	this.Sources = sources
	this.Pagination = pagination
	return &this
}

// NewListConnectedSourcesFromWarehouseV1OutputWithDefaults instantiates a new ListConnectedSourcesFromWarehouseV1Output object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListConnectedSourcesFromWarehouseV1OutputWithDefaults() *ListConnectedSourcesFromWarehouseV1Output {
	this := ListConnectedSourcesFromWarehouseV1Output{}
	return &this
}

// GetSources returns the Sources field value
func (o *ListConnectedSourcesFromWarehouseV1Output) GetSources() []SourceV1 {
	if o == nil {
		var ret []SourceV1
		return ret
	}

	return o.Sources
}

// GetSourcesOk returns a tuple with the Sources field value
// and a boolean to check if the value has been set.
func (o *ListConnectedSourcesFromWarehouseV1Output) GetSourcesOk() ([]SourceV1, bool) {
	if o == nil {
		return nil, false
	}
	return o.Sources, true
}

// SetSources sets field value
func (o *ListConnectedSourcesFromWarehouseV1Output) SetSources(v []SourceV1) {
	o.Sources = v
}

// GetPagination returns the Pagination field value
func (o *ListConnectedSourcesFromWarehouseV1Output) GetPagination() Pagination {
	if o == nil {
		var ret Pagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *ListConnectedSourcesFromWarehouseV1Output) GetPaginationOk() (*Pagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *ListConnectedSourcesFromWarehouseV1Output) SetPagination(v Pagination) {
	o.Pagination = v
}

func (o ListConnectedSourcesFromWarehouseV1Output) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["sources"] = o.Sources
	}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	return json.Marshal(toSerialize)
}

type NullableListConnectedSourcesFromWarehouseV1Output struct {
	value *ListConnectedSourcesFromWarehouseV1Output
	isSet bool
}

func (v NullableListConnectedSourcesFromWarehouseV1Output) Get() *ListConnectedSourcesFromWarehouseV1Output {
	return v.value
}

func (v *NullableListConnectedSourcesFromWarehouseV1Output) Set(
	val *ListConnectedSourcesFromWarehouseV1Output,
) {
	v.value = val
	v.isSet = true
}

func (v NullableListConnectedSourcesFromWarehouseV1Output) IsSet() bool {
	return v.isSet
}

func (v *NullableListConnectedSourcesFromWarehouseV1Output) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListConnectedSourcesFromWarehouseV1Output(
	val *ListConnectedSourcesFromWarehouseV1Output,
) *NullableListConnectedSourcesFromWarehouseV1Output {
	return &NullableListConnectedSourcesFromWarehouseV1Output{value: val, isSet: true}
}

func (v NullableListConnectedSourcesFromWarehouseV1Output) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListConnectedSourcesFromWarehouseV1Output) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
