/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 34.0.1
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// UpdateWarehouseV1Output Returns the updated Warehouse.
type UpdateWarehouseV1Output struct {
	Warehouse Warehouse2 `json:"warehouse"`
}

// NewUpdateWarehouseV1Output instantiates a new UpdateWarehouseV1Output object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateWarehouseV1Output(warehouse Warehouse2) *UpdateWarehouseV1Output {
	this := UpdateWarehouseV1Output{}
	this.Warehouse = warehouse
	return &this
}

// NewUpdateWarehouseV1OutputWithDefaults instantiates a new UpdateWarehouseV1Output object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateWarehouseV1OutputWithDefaults() *UpdateWarehouseV1Output {
	this := UpdateWarehouseV1Output{}
	return &this
}

// GetWarehouse returns the Warehouse field value
func (o *UpdateWarehouseV1Output) GetWarehouse() Warehouse2 {
	if o == nil {
		var ret Warehouse2
		return ret
	}

	return o.Warehouse
}

// GetWarehouseOk returns a tuple with the Warehouse field value
// and a boolean to check if the value has been set.
func (o *UpdateWarehouseV1Output) GetWarehouseOk() (*Warehouse2, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Warehouse, true
}

// SetWarehouse sets field value
func (o *UpdateWarehouseV1Output) SetWarehouse(v Warehouse2) {
	o.Warehouse = v
}

func (o UpdateWarehouseV1Output) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["warehouse"] = o.Warehouse
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateWarehouseV1Output struct {
	value *UpdateWarehouseV1Output
	isSet bool
}

func (v NullableUpdateWarehouseV1Output) Get() *UpdateWarehouseV1Output {
	return v.value
}

func (v *NullableUpdateWarehouseV1Output) Set(val *UpdateWarehouseV1Output) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateWarehouseV1Output) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateWarehouseV1Output) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateWarehouseV1Output(
	val *UpdateWarehouseV1Output,
) *NullableUpdateWarehouseV1Output {
	return &NullableUpdateWarehouseV1Output{value: val, isSet: true}
}

func (v NullableUpdateWarehouseV1Output) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateWarehouseV1Output) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
