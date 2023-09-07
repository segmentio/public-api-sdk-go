/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 36.1.1
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// ReplaceAdvancedSyncScheduleForWarehouseV1Output Returns the advanced sync schedule for a Warehouse.
type ReplaceAdvancedSyncScheduleForWarehouseV1Output struct {
	// Indicates if an advanced sync schedule is enabled for the Warehouse.
	Enabled  bool       `json:"enabled"`
	Schedule *Schedule2 `json:"schedule,omitempty"`
}

// NewReplaceAdvancedSyncScheduleForWarehouseV1Output instantiates a new ReplaceAdvancedSyncScheduleForWarehouseV1Output object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReplaceAdvancedSyncScheduleForWarehouseV1Output(
	enabled bool,
) *ReplaceAdvancedSyncScheduleForWarehouseV1Output {
	this := ReplaceAdvancedSyncScheduleForWarehouseV1Output{}
	this.Enabled = enabled
	return &this
}

// NewReplaceAdvancedSyncScheduleForWarehouseV1OutputWithDefaults instantiates a new ReplaceAdvancedSyncScheduleForWarehouseV1Output object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReplaceAdvancedSyncScheduleForWarehouseV1OutputWithDefaults() *ReplaceAdvancedSyncScheduleForWarehouseV1Output {
	this := ReplaceAdvancedSyncScheduleForWarehouseV1Output{}
	return &this
}

// GetEnabled returns the Enabled field value
func (o *ReplaceAdvancedSyncScheduleForWarehouseV1Output) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *ReplaceAdvancedSyncScheduleForWarehouseV1Output) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *ReplaceAdvancedSyncScheduleForWarehouseV1Output) SetEnabled(v bool) {
	o.Enabled = v
}

// GetSchedule returns the Schedule field value if set, zero value otherwise.
func (o *ReplaceAdvancedSyncScheduleForWarehouseV1Output) GetSchedule() Schedule2 {
	if o == nil || o.Schedule == nil {
		var ret Schedule2
		return ret
	}
	return *o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplaceAdvancedSyncScheduleForWarehouseV1Output) GetScheduleOk() (*Schedule2, bool) {
	if o == nil || o.Schedule == nil {
		return nil, false
	}
	return o.Schedule, true
}

// HasSchedule returns a boolean if a field has been set.
func (o *ReplaceAdvancedSyncScheduleForWarehouseV1Output) HasSchedule() bool {
	if o != nil && o.Schedule != nil {
		return true
	}

	return false
}

// SetSchedule gets a reference to the given Schedule2 and assigns it to the Schedule field.
func (o *ReplaceAdvancedSyncScheduleForWarehouseV1Output) SetSchedule(v Schedule2) {
	o.Schedule = &v
}

func (o ReplaceAdvancedSyncScheduleForWarehouseV1Output) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Schedule != nil {
		toSerialize["schedule"] = o.Schedule
	}
	return json.Marshal(toSerialize)
}

type NullableReplaceAdvancedSyncScheduleForWarehouseV1Output struct {
	value *ReplaceAdvancedSyncScheduleForWarehouseV1Output
	isSet bool
}

func (v NullableReplaceAdvancedSyncScheduleForWarehouseV1Output) Get() *ReplaceAdvancedSyncScheduleForWarehouseV1Output {
	return v.value
}

func (v *NullableReplaceAdvancedSyncScheduleForWarehouseV1Output) Set(
	val *ReplaceAdvancedSyncScheduleForWarehouseV1Output,
) {
	v.value = val
	v.isSet = true
}

func (v NullableReplaceAdvancedSyncScheduleForWarehouseV1Output) IsSet() bool {
	return v.isSet
}

func (v *NullableReplaceAdvancedSyncScheduleForWarehouseV1Output) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReplaceAdvancedSyncScheduleForWarehouseV1Output(
	val *ReplaceAdvancedSyncScheduleForWarehouseV1Output,
) *NullableReplaceAdvancedSyncScheduleForWarehouseV1Output {
	return &NullableReplaceAdvancedSyncScheduleForWarehouseV1Output{value: val, isSet: true}
}

func (v NullableReplaceAdvancedSyncScheduleForWarehouseV1Output) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReplaceAdvancedSyncScheduleForWarehouseV1Output) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
