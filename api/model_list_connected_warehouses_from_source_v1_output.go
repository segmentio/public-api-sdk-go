/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 38.1.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the ListConnectedWarehousesFromSourceV1Output type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListConnectedWarehousesFromSourceV1Output{}

// ListConnectedWarehousesFromSourceV1Output Returns a list of Warehouses connected to a Source.
type ListConnectedWarehousesFromSourceV1Output struct {
	// A list that contains the Warehouses connected to the Source.
	Warehouses []WarehouseV1    `json:"warehouses"`
	Pagination PaginationOutput `json:"pagination"`
}

// NewListConnectedWarehousesFromSourceV1Output instantiates a new ListConnectedWarehousesFromSourceV1Output object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListConnectedWarehousesFromSourceV1Output(
	warehouses []WarehouseV1,
	pagination PaginationOutput,
) *ListConnectedWarehousesFromSourceV1Output {
	this := ListConnectedWarehousesFromSourceV1Output{}
	this.Warehouses = warehouses
	this.Pagination = pagination
	return &this
}

// NewListConnectedWarehousesFromSourceV1OutputWithDefaults instantiates a new ListConnectedWarehousesFromSourceV1Output object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListConnectedWarehousesFromSourceV1OutputWithDefaults() *ListConnectedWarehousesFromSourceV1Output {
	this := ListConnectedWarehousesFromSourceV1Output{}
	return &this
}

// GetWarehouses returns the Warehouses field value
func (o *ListConnectedWarehousesFromSourceV1Output) GetWarehouses() []WarehouseV1 {
	if o == nil {
		var ret []WarehouseV1
		return ret
	}

	return o.Warehouses
}

// GetWarehousesOk returns a tuple with the Warehouses field value
// and a boolean to check if the value has been set.
func (o *ListConnectedWarehousesFromSourceV1Output) GetWarehousesOk() ([]WarehouseV1, bool) {
	if o == nil {
		return nil, false
	}
	return o.Warehouses, true
}

// SetWarehouses sets field value
func (o *ListConnectedWarehousesFromSourceV1Output) SetWarehouses(v []WarehouseV1) {
	o.Warehouses = v
}

// GetPagination returns the Pagination field value
func (o *ListConnectedWarehousesFromSourceV1Output) GetPagination() PaginationOutput {
	if o == nil {
		var ret PaginationOutput
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *ListConnectedWarehousesFromSourceV1Output) GetPaginationOk() (*PaginationOutput, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *ListConnectedWarehousesFromSourceV1Output) SetPagination(v PaginationOutput) {
	o.Pagination = v
}

func (o ListConnectedWarehousesFromSourceV1Output) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListConnectedWarehousesFromSourceV1Output) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["warehouses"] = o.Warehouses
	toSerialize["pagination"] = o.Pagination
	return toSerialize, nil
}

type NullableListConnectedWarehousesFromSourceV1Output struct {
	value *ListConnectedWarehousesFromSourceV1Output
	isSet bool
}

func (v NullableListConnectedWarehousesFromSourceV1Output) Get() *ListConnectedWarehousesFromSourceV1Output {
	return v.value
}

func (v *NullableListConnectedWarehousesFromSourceV1Output) Set(
	val *ListConnectedWarehousesFromSourceV1Output,
) {
	v.value = val
	v.isSet = true
}

func (v NullableListConnectedWarehousesFromSourceV1Output) IsSet() bool {
	return v.isSet
}

func (v *NullableListConnectedWarehousesFromSourceV1Output) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListConnectedWarehousesFromSourceV1Output(
	val *ListConnectedWarehousesFromSourceV1Output,
) *NullableListConnectedWarehousesFromSourceV1Output {
	return &NullableListConnectedWarehousesFromSourceV1Output{value: val, isSet: true}
}

func (v NullableListConnectedWarehousesFromSourceV1Output) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListConnectedWarehousesFromSourceV1Output) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
