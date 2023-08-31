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

// UpdateSelectiveSyncForWarehouseAndSpaceAlphaOutput Results from a SelectiveSync patch to a Space Warehouse connection.
type UpdateSelectiveSyncForWarehouseAndSpaceAlphaOutput struct {
	// Status of the update operation.
	Status string `json:"status"`
}

// NewUpdateSelectiveSyncForWarehouseAndSpaceAlphaOutput instantiates a new UpdateSelectiveSyncForWarehouseAndSpaceAlphaOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateSelectiveSyncForWarehouseAndSpaceAlphaOutput(
	status string,
) *UpdateSelectiveSyncForWarehouseAndSpaceAlphaOutput {
	this := UpdateSelectiveSyncForWarehouseAndSpaceAlphaOutput{}
	this.Status = status
	return &this
}

// NewUpdateSelectiveSyncForWarehouseAndSpaceAlphaOutputWithDefaults instantiates a new UpdateSelectiveSyncForWarehouseAndSpaceAlphaOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateSelectiveSyncForWarehouseAndSpaceAlphaOutputWithDefaults() *UpdateSelectiveSyncForWarehouseAndSpaceAlphaOutput {
	this := UpdateSelectiveSyncForWarehouseAndSpaceAlphaOutput{}
	return &this
}

// GetStatus returns the Status field value
func (o *UpdateSelectiveSyncForWarehouseAndSpaceAlphaOutput) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *UpdateSelectiveSyncForWarehouseAndSpaceAlphaOutput) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *UpdateSelectiveSyncForWarehouseAndSpaceAlphaOutput) SetStatus(v string) {
	o.Status = v
}

func (o UpdateSelectiveSyncForWarehouseAndSpaceAlphaOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateSelectiveSyncForWarehouseAndSpaceAlphaOutput struct {
	value *UpdateSelectiveSyncForWarehouseAndSpaceAlphaOutput
	isSet bool
}

func (v NullableUpdateSelectiveSyncForWarehouseAndSpaceAlphaOutput) Get() *UpdateSelectiveSyncForWarehouseAndSpaceAlphaOutput {
	return v.value
}

func (v *NullableUpdateSelectiveSyncForWarehouseAndSpaceAlphaOutput) Set(
	val *UpdateSelectiveSyncForWarehouseAndSpaceAlphaOutput,
) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateSelectiveSyncForWarehouseAndSpaceAlphaOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateSelectiveSyncForWarehouseAndSpaceAlphaOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateSelectiveSyncForWarehouseAndSpaceAlphaOutput(
	val *UpdateSelectiveSyncForWarehouseAndSpaceAlphaOutput,
) *NullableUpdateSelectiveSyncForWarehouseAndSpaceAlphaOutput {
	return &NullableUpdateSelectiveSyncForWarehouseAndSpaceAlphaOutput{value: val, isSet: true}
}

func (v NullableUpdateSelectiveSyncForWarehouseAndSpaceAlphaOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateSelectiveSyncForWarehouseAndSpaceAlphaOutput) UnmarshalJSON(
	src []byte,
) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
