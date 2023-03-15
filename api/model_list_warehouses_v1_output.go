/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 34.3.2
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// ListWarehousesV1Output Returns a list of Warehouses that belong to a Workspace.
type ListWarehousesV1Output struct {
	// A list of Warehouses that belong to the Workspace.
	Warehouses []WarehouseV1 `json:"warehouses"`
	Pagination Pagination    `json:"pagination"`
}

// NewListWarehousesV1Output instantiates a new ListWarehousesV1Output object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListWarehousesV1Output(
	warehouses []WarehouseV1,
	pagination Pagination,
) *ListWarehousesV1Output {
	this := ListWarehousesV1Output{}
	this.Warehouses = warehouses
	this.Pagination = pagination
	return &this
}

// NewListWarehousesV1OutputWithDefaults instantiates a new ListWarehousesV1Output object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListWarehousesV1OutputWithDefaults() *ListWarehousesV1Output {
	this := ListWarehousesV1Output{}
	return &this
}

// GetWarehouses returns the Warehouses field value
func (o *ListWarehousesV1Output) GetWarehouses() []WarehouseV1 {
	if o == nil {
		var ret []WarehouseV1
		return ret
	}

	return o.Warehouses
}

// GetWarehousesOk returns a tuple with the Warehouses field value
// and a boolean to check if the value has been set.
func (o *ListWarehousesV1Output) GetWarehousesOk() ([]WarehouseV1, bool) {
	if o == nil {
		return nil, false
	}
	return o.Warehouses, true
}

// SetWarehouses sets field value
func (o *ListWarehousesV1Output) SetWarehouses(v []WarehouseV1) {
	o.Warehouses = v
}

// GetPagination returns the Pagination field value
func (o *ListWarehousesV1Output) GetPagination() Pagination {
	if o == nil {
		var ret Pagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *ListWarehousesV1Output) GetPaginationOk() (*Pagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *ListWarehousesV1Output) SetPagination(v Pagination) {
	o.Pagination = v
}

func (o ListWarehousesV1Output) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["warehouses"] = o.Warehouses
	}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	return json.Marshal(toSerialize)
}

type NullableListWarehousesV1Output struct {
	value *ListWarehousesV1Output
	isSet bool
}

func (v NullableListWarehousesV1Output) Get() *ListWarehousesV1Output {
	return v.value
}

func (v *NullableListWarehousesV1Output) Set(val *ListWarehousesV1Output) {
	v.value = val
	v.isSet = true
}

func (v NullableListWarehousesV1Output) IsSet() bool {
	return v.isSet
}

func (v *NullableListWarehousesV1Output) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListWarehousesV1Output(
	val *ListWarehousesV1Output,
) *NullableListWarehousesV1Output {
	return &NullableListWarehousesV1Output{value: val, isSet: true}
}

func (v NullableListWarehousesV1Output) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListWarehousesV1Output) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
