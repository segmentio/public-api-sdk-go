/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 37.0.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// GetWarehouseV1Output Returns a Warehouse.
type GetWarehouseV1Output struct {
	Warehouse Warehouse `json:"warehouse"`
}

// NewGetWarehouseV1Output instantiates a new GetWarehouseV1Output object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetWarehouseV1Output(warehouse Warehouse) *GetWarehouseV1Output {
	this := GetWarehouseV1Output{}
	this.Warehouse = warehouse
	return &this
}

// NewGetWarehouseV1OutputWithDefaults instantiates a new GetWarehouseV1Output object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetWarehouseV1OutputWithDefaults() *GetWarehouseV1Output {
	this := GetWarehouseV1Output{}
	return &this
}

// GetWarehouse returns the Warehouse field value
func (o *GetWarehouseV1Output) GetWarehouse() Warehouse {
	if o == nil {
		var ret Warehouse
		return ret
	}

	return o.Warehouse
}

// GetWarehouseOk returns a tuple with the Warehouse field value
// and a boolean to check if the value has been set.
func (o *GetWarehouseV1Output) GetWarehouseOk() (*Warehouse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Warehouse, true
}

// SetWarehouse sets field value
func (o *GetWarehouseV1Output) SetWarehouse(v Warehouse) {
	o.Warehouse = v
}

func (o GetWarehouseV1Output) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["warehouse"] = o.Warehouse
	}
	return json.Marshal(toSerialize)
}

type NullableGetWarehouseV1Output struct {
	value *GetWarehouseV1Output
	isSet bool
}

func (v NullableGetWarehouseV1Output) Get() *GetWarehouseV1Output {
	return v.value
}

func (v *NullableGetWarehouseV1Output) Set(val *GetWarehouseV1Output) {
	v.value = val
	v.isSet = true
}

func (v NullableGetWarehouseV1Output) IsSet() bool {
	return v.isSet
}

func (v *NullableGetWarehouseV1Output) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetWarehouseV1Output(val *GetWarehouseV1Output) *NullableGetWarehouseV1Output {
	return &NullableGetWarehouseV1Output{value: val, isSet: true}
}

func (v NullableGetWarehouseV1Output) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetWarehouseV1Output) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
