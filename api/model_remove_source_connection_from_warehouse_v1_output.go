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

// RemoveSourceConnectionFromWarehouseV1Output Response indicating whether the disconnection was successful.
type RemoveSourceConnectionFromWarehouseV1Output struct {
	// The status of the request to disconnect the Source and Warehouse.
	Status string `json:"status"`
}

// NewRemoveSourceConnectionFromWarehouseV1Output instantiates a new RemoveSourceConnectionFromWarehouseV1Output object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemoveSourceConnectionFromWarehouseV1Output(
	status string,
) *RemoveSourceConnectionFromWarehouseV1Output {
	this := RemoveSourceConnectionFromWarehouseV1Output{}
	this.Status = status
	return &this
}

// NewRemoveSourceConnectionFromWarehouseV1OutputWithDefaults instantiates a new RemoveSourceConnectionFromWarehouseV1Output object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemoveSourceConnectionFromWarehouseV1OutputWithDefaults() *RemoveSourceConnectionFromWarehouseV1Output {
	this := RemoveSourceConnectionFromWarehouseV1Output{}
	return &this
}

// GetStatus returns the Status field value
func (o *RemoveSourceConnectionFromWarehouseV1Output) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *RemoveSourceConnectionFromWarehouseV1Output) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *RemoveSourceConnectionFromWarehouseV1Output) SetStatus(v string) {
	o.Status = v
}

func (o RemoveSourceConnectionFromWarehouseV1Output) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableRemoveSourceConnectionFromWarehouseV1Output struct {
	value *RemoveSourceConnectionFromWarehouseV1Output
	isSet bool
}

func (v NullableRemoveSourceConnectionFromWarehouseV1Output) Get() *RemoveSourceConnectionFromWarehouseV1Output {
	return v.value
}

func (v *NullableRemoveSourceConnectionFromWarehouseV1Output) Set(
	val *RemoveSourceConnectionFromWarehouseV1Output,
) {
	v.value = val
	v.isSet = true
}

func (v NullableRemoveSourceConnectionFromWarehouseV1Output) IsSet() bool {
	return v.isSet
}

func (v *NullableRemoveSourceConnectionFromWarehouseV1Output) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemoveSourceConnectionFromWarehouseV1Output(
	val *RemoveSourceConnectionFromWarehouseV1Output,
) *NullableRemoveSourceConnectionFromWarehouseV1Output {
	return &NullableRemoveSourceConnectionFromWarehouseV1Output{value: val, isSet: true}
}

func (v NullableRemoveSourceConnectionFromWarehouseV1Output) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemoveSourceConnectionFromWarehouseV1Output) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
