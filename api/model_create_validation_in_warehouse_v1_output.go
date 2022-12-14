/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 33.0.4
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// CreateValidationInWarehouseV1Output Returns the status of a Warehouse connection settings after an attempt to connect to it.
type CreateValidationInWarehouseV1Output struct {
	// Represents the status for the current connection settings.
	Status string `json:"status"`
}

// NewCreateValidationInWarehouseV1Output instantiates a new CreateValidationInWarehouseV1Output object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateValidationInWarehouseV1Output(status string) *CreateValidationInWarehouseV1Output {
	this := CreateValidationInWarehouseV1Output{}
	this.Status = status
	return &this
}

// NewCreateValidationInWarehouseV1OutputWithDefaults instantiates a new CreateValidationInWarehouseV1Output object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateValidationInWarehouseV1OutputWithDefaults() *CreateValidationInWarehouseV1Output {
	this := CreateValidationInWarehouseV1Output{}
	return &this
}

// GetStatus returns the Status field value
func (o *CreateValidationInWarehouseV1Output) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *CreateValidationInWarehouseV1Output) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *CreateValidationInWarehouseV1Output) SetStatus(v string) {
	o.Status = v
}

func (o CreateValidationInWarehouseV1Output) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableCreateValidationInWarehouseV1Output struct {
	value *CreateValidationInWarehouseV1Output
	isSet bool
}

func (v NullableCreateValidationInWarehouseV1Output) Get() *CreateValidationInWarehouseV1Output {
	return v.value
}

func (v *NullableCreateValidationInWarehouseV1Output) Set(
	val *CreateValidationInWarehouseV1Output,
) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateValidationInWarehouseV1Output) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateValidationInWarehouseV1Output) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateValidationInWarehouseV1Output(
	val *CreateValidationInWarehouseV1Output,
) *NullableCreateValidationInWarehouseV1Output {
	return &NullableCreateValidationInWarehouseV1Output{value: val, isSet: true}
}

func (v NullableCreateValidationInWarehouseV1Output) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateValidationInWarehouseV1Output) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
