/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API. 

API version: 32.0.4
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// CreateWarehouseV1Output Returns the newly created Warehouse.
type CreateWarehouseV1Output struct {
	Warehouse Warehouse1 `json:"warehouse"`
}

// NewCreateWarehouseV1Output instantiates a new CreateWarehouseV1Output object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateWarehouseV1Output(warehouse Warehouse1) *CreateWarehouseV1Output {
	this := CreateWarehouseV1Output{}
	this.Warehouse = warehouse
	return &this
}

// NewCreateWarehouseV1OutputWithDefaults instantiates a new CreateWarehouseV1Output object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateWarehouseV1OutputWithDefaults() *CreateWarehouseV1Output {
	this := CreateWarehouseV1Output{}
	return &this
}

// GetWarehouse returns the Warehouse field value
func (o *CreateWarehouseV1Output) GetWarehouse() Warehouse1 {
	if o == nil {
		var ret Warehouse1
		return ret
	}

	return o.Warehouse
}

// GetWarehouseOk returns a tuple with the Warehouse field value
// and a boolean to check if the value has been set.
func (o *CreateWarehouseV1Output) GetWarehouseOk() (*Warehouse1, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Warehouse, true
}

// SetWarehouse sets field value
func (o *CreateWarehouseV1Output) SetWarehouse(v Warehouse1) {
	o.Warehouse = v
}

func (o CreateWarehouseV1Output) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["warehouse"] = o.Warehouse
	}
	return json.Marshal(toSerialize)
}

type NullableCreateWarehouseV1Output struct {
	value *CreateWarehouseV1Output
	isSet bool
}

func (v NullableCreateWarehouseV1Output) Get() *CreateWarehouseV1Output {
	return v.value
}

func (v *NullableCreateWarehouseV1Output) Set(val *CreateWarehouseV1Output) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateWarehouseV1Output) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateWarehouseV1Output) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateWarehouseV1Output(val *CreateWarehouseV1Output) *NullableCreateWarehouseV1Output {
	return &NullableCreateWarehouseV1Output{value: val, isSet: true}
}

func (v NullableCreateWarehouseV1Output) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateWarehouseV1Output) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


