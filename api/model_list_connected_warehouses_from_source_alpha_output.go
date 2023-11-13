/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 37.2.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// ListConnectedWarehousesFromSourceAlphaOutput Returns a list of Warehouses connected to a Source.
type ListConnectedWarehousesFromSourceAlphaOutput struct {
	// A list that contains the Warehouses connected to the Source.
	Warehouses []WarehouseV1 `json:"warehouses"`
	Pagination Pagination    `json:"pagination"`
}

// NewListConnectedWarehousesFromSourceAlphaOutput instantiates a new ListConnectedWarehousesFromSourceAlphaOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListConnectedWarehousesFromSourceAlphaOutput(
	warehouses []WarehouseV1,
	pagination Pagination,
) *ListConnectedWarehousesFromSourceAlphaOutput {
	this := ListConnectedWarehousesFromSourceAlphaOutput{}
	this.Warehouses = warehouses
	this.Pagination = pagination
	return &this
}

// NewListConnectedWarehousesFromSourceAlphaOutputWithDefaults instantiates a new ListConnectedWarehousesFromSourceAlphaOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListConnectedWarehousesFromSourceAlphaOutputWithDefaults() *ListConnectedWarehousesFromSourceAlphaOutput {
	this := ListConnectedWarehousesFromSourceAlphaOutput{}
	return &this
}

// GetWarehouses returns the Warehouses field value
func (o *ListConnectedWarehousesFromSourceAlphaOutput) GetWarehouses() []WarehouseV1 {
	if o == nil {
		var ret []WarehouseV1
		return ret
	}

	return o.Warehouses
}

// GetWarehousesOk returns a tuple with the Warehouses field value
// and a boolean to check if the value has been set.
func (o *ListConnectedWarehousesFromSourceAlphaOutput) GetWarehousesOk() ([]WarehouseV1, bool) {
	if o == nil {
		return nil, false
	}
	return o.Warehouses, true
}

// SetWarehouses sets field value
func (o *ListConnectedWarehousesFromSourceAlphaOutput) SetWarehouses(v []WarehouseV1) {
	o.Warehouses = v
}

// GetPagination returns the Pagination field value
func (o *ListConnectedWarehousesFromSourceAlphaOutput) GetPagination() Pagination {
	if o == nil {
		var ret Pagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *ListConnectedWarehousesFromSourceAlphaOutput) GetPaginationOk() (*Pagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *ListConnectedWarehousesFromSourceAlphaOutput) SetPagination(v Pagination) {
	o.Pagination = v
}

func (o ListConnectedWarehousesFromSourceAlphaOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["warehouses"] = o.Warehouses
	}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	return json.Marshal(toSerialize)
}

type NullableListConnectedWarehousesFromSourceAlphaOutput struct {
	value *ListConnectedWarehousesFromSourceAlphaOutput
	isSet bool
}

func (v NullableListConnectedWarehousesFromSourceAlphaOutput) Get() *ListConnectedWarehousesFromSourceAlphaOutput {
	return v.value
}

func (v *NullableListConnectedWarehousesFromSourceAlphaOutput) Set(
	val *ListConnectedWarehousesFromSourceAlphaOutput,
) {
	v.value = val
	v.isSet = true
}

func (v NullableListConnectedWarehousesFromSourceAlphaOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableListConnectedWarehousesFromSourceAlphaOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListConnectedWarehousesFromSourceAlphaOutput(
	val *ListConnectedWarehousesFromSourceAlphaOutput,
) *NullableListConnectedWarehousesFromSourceAlphaOutput {
	return &NullableListConnectedWarehousesFromSourceAlphaOutput{value: val, isSet: true}
}

func (v NullableListConnectedWarehousesFromSourceAlphaOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListConnectedWarehousesFromSourceAlphaOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
