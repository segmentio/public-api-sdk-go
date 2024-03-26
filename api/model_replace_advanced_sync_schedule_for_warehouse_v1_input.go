/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 47.0.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the ReplaceAdvancedSyncScheduleForWarehouseV1Input type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReplaceAdvancedSyncScheduleForWarehouseV1Input{}

// ReplaceAdvancedSyncScheduleForWarehouseV1Input Replaces the advanced sync schedule for a Warehouse.
type ReplaceAdvancedSyncScheduleForWarehouseV1Input struct {
	// Enable to turn on an advanced sync schedule for the Warehouse.
	Enabled  bool                                  `json:"enabled"`
	Schedule *AdvancedWarehouseSyncScheduleV1Input `json:"schedule,omitempty"`
}

// NewReplaceAdvancedSyncScheduleForWarehouseV1Input instantiates a new ReplaceAdvancedSyncScheduleForWarehouseV1Input object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReplaceAdvancedSyncScheduleForWarehouseV1Input(
	enabled bool,
) *ReplaceAdvancedSyncScheduleForWarehouseV1Input {
	this := ReplaceAdvancedSyncScheduleForWarehouseV1Input{}
	this.Enabled = enabled
	return &this
}

// NewReplaceAdvancedSyncScheduleForWarehouseV1InputWithDefaults instantiates a new ReplaceAdvancedSyncScheduleForWarehouseV1Input object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReplaceAdvancedSyncScheduleForWarehouseV1InputWithDefaults() *ReplaceAdvancedSyncScheduleForWarehouseV1Input {
	this := ReplaceAdvancedSyncScheduleForWarehouseV1Input{}
	return &this
}

// GetEnabled returns the Enabled field value
func (o *ReplaceAdvancedSyncScheduleForWarehouseV1Input) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *ReplaceAdvancedSyncScheduleForWarehouseV1Input) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *ReplaceAdvancedSyncScheduleForWarehouseV1Input) SetEnabled(v bool) {
	o.Enabled = v
}

// GetSchedule returns the Schedule field value if set, zero value otherwise.
func (o *ReplaceAdvancedSyncScheduleForWarehouseV1Input) GetSchedule() AdvancedWarehouseSyncScheduleV1Input {
	if o == nil || IsNil(o.Schedule) {
		var ret AdvancedWarehouseSyncScheduleV1Input
		return ret
	}
	return *o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplaceAdvancedSyncScheduleForWarehouseV1Input) GetScheduleOk() (*AdvancedWarehouseSyncScheduleV1Input, bool) {
	if o == nil || IsNil(o.Schedule) {
		return nil, false
	}
	return o.Schedule, true
}

// HasSchedule returns a boolean if a field has been set.
func (o *ReplaceAdvancedSyncScheduleForWarehouseV1Input) HasSchedule() bool {
	if o != nil && !IsNil(o.Schedule) {
		return true
	}

	return false
}

// SetSchedule gets a reference to the given AdvancedWarehouseSyncScheduleV1Input and assigns it to the Schedule field.
func (o *ReplaceAdvancedSyncScheduleForWarehouseV1Input) SetSchedule(
	v AdvancedWarehouseSyncScheduleV1Input,
) {
	o.Schedule = &v
}

func (o ReplaceAdvancedSyncScheduleForWarehouseV1Input) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReplaceAdvancedSyncScheduleForWarehouseV1Input) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["enabled"] = o.Enabled
	if !IsNil(o.Schedule) {
		toSerialize["schedule"] = o.Schedule
	}
	return toSerialize, nil
}

type NullableReplaceAdvancedSyncScheduleForWarehouseV1Input struct {
	value *ReplaceAdvancedSyncScheduleForWarehouseV1Input
	isSet bool
}

func (v NullableReplaceAdvancedSyncScheduleForWarehouseV1Input) Get() *ReplaceAdvancedSyncScheduleForWarehouseV1Input {
	return v.value
}

func (v *NullableReplaceAdvancedSyncScheduleForWarehouseV1Input) Set(
	val *ReplaceAdvancedSyncScheduleForWarehouseV1Input,
) {
	v.value = val
	v.isSet = true
}

func (v NullableReplaceAdvancedSyncScheduleForWarehouseV1Input) IsSet() bool {
	return v.isSet
}

func (v *NullableReplaceAdvancedSyncScheduleForWarehouseV1Input) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReplaceAdvancedSyncScheduleForWarehouseV1Input(
	val *ReplaceAdvancedSyncScheduleForWarehouseV1Input,
) *NullableReplaceAdvancedSyncScheduleForWarehouseV1Input {
	return &NullableReplaceAdvancedSyncScheduleForWarehouseV1Input{value: val, isSet: true}
}

func (v NullableReplaceAdvancedSyncScheduleForWarehouseV1Input) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReplaceAdvancedSyncScheduleForWarehouseV1Input) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
