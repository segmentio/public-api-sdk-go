/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 34.3.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// Schedule2 The schedule that contains the overrides for the Warehouse, if enabled.
type Schedule2 struct {
	// A list that contains the times when syncs should occur.
	Times []WarehouseAdvancedSyncV1 `json:"times,omitempty"`
	// A TZ-database timezone for this sync schedule.
	Timezone *string `json:"timezone,omitempty"`
}

// NewSchedule2 instantiates a new Schedule2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchedule2() *Schedule2 {
	this := Schedule2{}
	return &this
}

// NewSchedule2WithDefaults instantiates a new Schedule2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchedule2WithDefaults() *Schedule2 {
	this := Schedule2{}
	return &this
}

// GetTimes returns the Times field value if set, zero value otherwise.
func (o *Schedule2) GetTimes() []WarehouseAdvancedSyncV1 {
	if o == nil || o.Times == nil {
		var ret []WarehouseAdvancedSyncV1
		return ret
	}
	return o.Times
}

// GetTimesOk returns a tuple with the Times field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Schedule2) GetTimesOk() ([]WarehouseAdvancedSyncV1, bool) {
	if o == nil || o.Times == nil {
		return nil, false
	}
	return o.Times, true
}

// HasTimes returns a boolean if a field has been set.
func (o *Schedule2) HasTimes() bool {
	if o != nil && o.Times != nil {
		return true
	}

	return false
}

// SetTimes gets a reference to the given []WarehouseAdvancedSyncV1 and assigns it to the Times field.
func (o *Schedule2) SetTimes(v []WarehouseAdvancedSyncV1) {
	o.Times = v
}

// GetTimezone returns the Timezone field value if set, zero value otherwise.
func (o *Schedule2) GetTimezone() string {
	if o == nil || o.Timezone == nil {
		var ret string
		return ret
	}
	return *o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Schedule2) GetTimezoneOk() (*string, bool) {
	if o == nil || o.Timezone == nil {
		return nil, false
	}
	return o.Timezone, true
}

// HasTimezone returns a boolean if a field has been set.
func (o *Schedule2) HasTimezone() bool {
	if o != nil && o.Timezone != nil {
		return true
	}

	return false
}

// SetTimezone gets a reference to the given string and assigns it to the Timezone field.
func (o *Schedule2) SetTimezone(v string) {
	o.Timezone = &v
}

func (o Schedule2) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Times != nil {
		toSerialize["times"] = o.Times
	}
	if o.Timezone != nil {
		toSerialize["timezone"] = o.Timezone
	}
	return json.Marshal(toSerialize)
}

type NullableSchedule2 struct {
	value *Schedule2
	isSet bool
}

func (v NullableSchedule2) Get() *Schedule2 {
	return v.value
}

func (v *NullableSchedule2) Set(val *Schedule2) {
	v.value = val
	v.isSet = true
}

func (v NullableSchedule2) IsSet() bool {
	return v.isSet
}

func (v *NullableSchedule2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchedule2(val *Schedule2) *NullableSchedule2 {
	return &NullableSchedule2{value: val, isSet: true}
}

func (v NullableSchedule2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchedule2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
