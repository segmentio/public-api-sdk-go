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

// AddConnectionFromSourceToWarehouseV1Output Response indicating whether the connection was successful.
type AddConnectionFromSourceToWarehouseV1Output struct {
	// The status of the connection between the Source and Warehouse.
	Status string `json:"status"`
}

// NewAddConnectionFromSourceToWarehouseV1Output instantiates a new AddConnectionFromSourceToWarehouseV1Output object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddConnectionFromSourceToWarehouseV1Output(status string) *AddConnectionFromSourceToWarehouseV1Output {
	this := AddConnectionFromSourceToWarehouseV1Output{}
	this.Status = status
	return &this
}

// NewAddConnectionFromSourceToWarehouseV1OutputWithDefaults instantiates a new AddConnectionFromSourceToWarehouseV1Output object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddConnectionFromSourceToWarehouseV1OutputWithDefaults() *AddConnectionFromSourceToWarehouseV1Output {
	this := AddConnectionFromSourceToWarehouseV1Output{}
	return &this
}

// GetStatus returns the Status field value
func (o *AddConnectionFromSourceToWarehouseV1Output) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *AddConnectionFromSourceToWarehouseV1Output) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *AddConnectionFromSourceToWarehouseV1Output) SetStatus(v string) {
	o.Status = v
}

func (o AddConnectionFromSourceToWarehouseV1Output) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableAddConnectionFromSourceToWarehouseV1Output struct {
	value *AddConnectionFromSourceToWarehouseV1Output
	isSet bool
}

func (v NullableAddConnectionFromSourceToWarehouseV1Output) Get() *AddConnectionFromSourceToWarehouseV1Output {
	return v.value
}

func (v *NullableAddConnectionFromSourceToWarehouseV1Output) Set(val *AddConnectionFromSourceToWarehouseV1Output) {
	v.value = val
	v.isSet = true
}

func (v NullableAddConnectionFromSourceToWarehouseV1Output) IsSet() bool {
	return v.isSet
}

func (v *NullableAddConnectionFromSourceToWarehouseV1Output) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddConnectionFromSourceToWarehouseV1Output(val *AddConnectionFromSourceToWarehouseV1Output) *NullableAddConnectionFromSourceToWarehouseV1Output {
	return &NullableAddConnectionFromSourceToWarehouseV1Output{value: val, isSet: true}
}

func (v NullableAddConnectionFromSourceToWarehouseV1Output) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddConnectionFromSourceToWarehouseV1Output) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


