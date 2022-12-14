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

// GetAdvancedSyncScheduleFromWarehouseV1Output Returns the advanced sync schedule for a Warehouse.
type GetAdvancedSyncScheduleFromWarehouseV1Output struct {
	// Indicates if an advanced sync schedule is enabled for this Warehouse.
	Enabled  bool      `json:"enabled"`
	Schedule *Schedule `json:"schedule,omitempty"`
}

// NewGetAdvancedSyncScheduleFromWarehouseV1Output instantiates a new GetAdvancedSyncScheduleFromWarehouseV1Output object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAdvancedSyncScheduleFromWarehouseV1Output(
	enabled bool,
) *GetAdvancedSyncScheduleFromWarehouseV1Output {
	this := GetAdvancedSyncScheduleFromWarehouseV1Output{}
	this.Enabled = enabled
	return &this
}

// NewGetAdvancedSyncScheduleFromWarehouseV1OutputWithDefaults instantiates a new GetAdvancedSyncScheduleFromWarehouseV1Output object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAdvancedSyncScheduleFromWarehouseV1OutputWithDefaults() *GetAdvancedSyncScheduleFromWarehouseV1Output {
	this := GetAdvancedSyncScheduleFromWarehouseV1Output{}
	return &this
}

// GetEnabled returns the Enabled field value
func (o *GetAdvancedSyncScheduleFromWarehouseV1Output) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *GetAdvancedSyncScheduleFromWarehouseV1Output) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *GetAdvancedSyncScheduleFromWarehouseV1Output) SetEnabled(v bool) {
	o.Enabled = v
}

// GetSchedule returns the Schedule field value if set, zero value otherwise.
func (o *GetAdvancedSyncScheduleFromWarehouseV1Output) GetSchedule() Schedule {
	if o == nil || o.Schedule == nil {
		var ret Schedule
		return ret
	}
	return *o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAdvancedSyncScheduleFromWarehouseV1Output) GetScheduleOk() (*Schedule, bool) {
	if o == nil || o.Schedule == nil {
		return nil, false
	}
	return o.Schedule, true
}

// HasSchedule returns a boolean if a field has been set.
func (o *GetAdvancedSyncScheduleFromWarehouseV1Output) HasSchedule() bool {
	if o != nil && o.Schedule != nil {
		return true
	}

	return false
}

// SetSchedule gets a reference to the given Schedule and assigns it to the Schedule field.
func (o *GetAdvancedSyncScheduleFromWarehouseV1Output) SetSchedule(v Schedule) {
	o.Schedule = &v
}

func (o GetAdvancedSyncScheduleFromWarehouseV1Output) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Schedule != nil {
		toSerialize["schedule"] = o.Schedule
	}
	return json.Marshal(toSerialize)
}

type NullableGetAdvancedSyncScheduleFromWarehouseV1Output struct {
	value *GetAdvancedSyncScheduleFromWarehouseV1Output
	isSet bool
}

func (v NullableGetAdvancedSyncScheduleFromWarehouseV1Output) Get() *GetAdvancedSyncScheduleFromWarehouseV1Output {
	return v.value
}

func (v *NullableGetAdvancedSyncScheduleFromWarehouseV1Output) Set(
	val *GetAdvancedSyncScheduleFromWarehouseV1Output,
) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAdvancedSyncScheduleFromWarehouseV1Output) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAdvancedSyncScheduleFromWarehouseV1Output) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAdvancedSyncScheduleFromWarehouseV1Output(
	val *GetAdvancedSyncScheduleFromWarehouseV1Output,
) *NullableGetAdvancedSyncScheduleFromWarehouseV1Output {
	return &NullableGetAdvancedSyncScheduleFromWarehouseV1Output{value: val, isSet: true}
}

func (v NullableGetAdvancedSyncScheduleFromWarehouseV1Output) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAdvancedSyncScheduleFromWarehouseV1Output) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
