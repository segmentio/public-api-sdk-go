/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 34.3.6
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// Schedule The schedule that contains the schedule overrides for the specified Warehouse, if enabled.
type Schedule struct {
	// A list that contains the times when syncs should occur.
	Times []WarehouseAdvancedSyncV1 `json:"times,omitempty"`
	// A TZ-database timezone for this sync schedule.
	Timezone *string `json:"timezone,omitempty"`
}

// NewSchedule instantiates a new Schedule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchedule() *Schedule {
	this := Schedule{}
	return &this
}

// NewScheduleWithDefaults instantiates a new Schedule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScheduleWithDefaults() *Schedule {
	this := Schedule{}
	return &this
}

// GetTimes returns the Times field value if set, zero value otherwise.
func (o *Schedule) GetTimes() []WarehouseAdvancedSyncV1 {
	if o == nil || o.Times == nil {
		var ret []WarehouseAdvancedSyncV1
		return ret
	}
	return o.Times
}

// GetTimesOk returns a tuple with the Times field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Schedule) GetTimesOk() ([]WarehouseAdvancedSyncV1, bool) {
	if o == nil || o.Times == nil {
		return nil, false
	}
	return o.Times, true
}

// HasTimes returns a boolean if a field has been set.
func (o *Schedule) HasTimes() bool {
	if o != nil && o.Times != nil {
		return true
	}

	return false
}

// SetTimes gets a reference to the given []WarehouseAdvancedSyncV1 and assigns it to the Times field.
func (o *Schedule) SetTimes(v []WarehouseAdvancedSyncV1) {
	o.Times = v
}

// GetTimezone returns the Timezone field value if set, zero value otherwise.
func (o *Schedule) GetTimezone() string {
	if o == nil || o.Timezone == nil {
		var ret string
		return ret
	}
	return *o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Schedule) GetTimezoneOk() (*string, bool) {
	if o == nil || o.Timezone == nil {
		return nil, false
	}
	return o.Timezone, true
}

// HasTimezone returns a boolean if a field has been set.
func (o *Schedule) HasTimezone() bool {
	if o != nil && o.Timezone != nil {
		return true
	}

	return false
}

// SetTimezone gets a reference to the given string and assigns it to the Timezone field.
func (o *Schedule) SetTimezone(v string) {
	o.Timezone = &v
}

func (o Schedule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Times != nil {
		toSerialize["times"] = o.Times
	}
	if o.Timezone != nil {
		toSerialize["timezone"] = o.Timezone
	}
	return json.Marshal(toSerialize)
}

type NullableSchedule struct {
	value *Schedule
	isSet bool
}

func (v NullableSchedule) Get() *Schedule {
	return v.value
}

func (v *NullableSchedule) Set(val *Schedule) {
	v.value = val
	v.isSet = true
}

func (v NullableSchedule) IsSet() bool {
	return v.isSet
}

func (v *NullableSchedule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchedule(val *Schedule) *NullableSchedule {
	return &NullableSchedule{value: val, isSet: true}
}

func (v NullableSchedule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchedule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
