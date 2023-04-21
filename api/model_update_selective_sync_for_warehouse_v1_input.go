/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 34.3.5
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// UpdateSelectiveSyncForWarehouseV1Input Updates the schema for a Warehouse/sources pair.
type UpdateSelectiveSyncForWarehouseV1Input struct {
	// A list of sync schema overrides to apply to this Warehouse.
	SyncOverrides []WarehouseSyncOverrideV1 `json:"syncOverrides"`
}

// NewUpdateSelectiveSyncForWarehouseV1Input instantiates a new UpdateSelectiveSyncForWarehouseV1Input object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateSelectiveSyncForWarehouseV1Input(
	syncOverrides []WarehouseSyncOverrideV1,
) *UpdateSelectiveSyncForWarehouseV1Input {
	this := UpdateSelectiveSyncForWarehouseV1Input{}
	this.SyncOverrides = syncOverrides
	return &this
}

// NewUpdateSelectiveSyncForWarehouseV1InputWithDefaults instantiates a new UpdateSelectiveSyncForWarehouseV1Input object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateSelectiveSyncForWarehouseV1InputWithDefaults() *UpdateSelectiveSyncForWarehouseV1Input {
	this := UpdateSelectiveSyncForWarehouseV1Input{}
	return &this
}

// GetSyncOverrides returns the SyncOverrides field value
func (o *UpdateSelectiveSyncForWarehouseV1Input) GetSyncOverrides() []WarehouseSyncOverrideV1 {
	if o == nil {
		var ret []WarehouseSyncOverrideV1
		return ret
	}

	return o.SyncOverrides
}

// GetSyncOverridesOk returns a tuple with the SyncOverrides field value
// and a boolean to check if the value has been set.
func (o *UpdateSelectiveSyncForWarehouseV1Input) GetSyncOverridesOk() ([]WarehouseSyncOverrideV1, bool) {
	if o == nil {
		return nil, false
	}
	return o.SyncOverrides, true
}

// SetSyncOverrides sets field value
func (o *UpdateSelectiveSyncForWarehouseV1Input) SetSyncOverrides(v []WarehouseSyncOverrideV1) {
	o.SyncOverrides = v
}

func (o UpdateSelectiveSyncForWarehouseV1Input) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["syncOverrides"] = o.SyncOverrides
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateSelectiveSyncForWarehouseV1Input struct {
	value *UpdateSelectiveSyncForWarehouseV1Input
	isSet bool
}

func (v NullableUpdateSelectiveSyncForWarehouseV1Input) Get() *UpdateSelectiveSyncForWarehouseV1Input {
	return v.value
}

func (v *NullableUpdateSelectiveSyncForWarehouseV1Input) Set(
	val *UpdateSelectiveSyncForWarehouseV1Input,
) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateSelectiveSyncForWarehouseV1Input) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateSelectiveSyncForWarehouseV1Input) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateSelectiveSyncForWarehouseV1Input(
	val *UpdateSelectiveSyncForWarehouseV1Input,
) *NullableUpdateSelectiveSyncForWarehouseV1Input {
	return &NullableUpdateSelectiveSyncForWarehouseV1Input{value: val, isSet: true}
}

func (v NullableUpdateSelectiveSyncForWarehouseV1Input) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateSelectiveSyncForWarehouseV1Input) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
