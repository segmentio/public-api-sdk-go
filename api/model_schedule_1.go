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

// Schedule1 The full sync schedule for the Warehouse.
type Schedule1 struct {
	// A list that contains the times when syncs should occur.
	Times []WarehouseAdvancedSyncV1 `json:"times"`
	// A TZ-database timezone for this sync schedule.
	Timezone string `json:"timezone"`
}

// NewSchedule1 instantiates a new Schedule1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchedule1(times []WarehouseAdvancedSyncV1, timezone string) *Schedule1 {
	this := Schedule1{}
	this.Times = times
	this.Timezone = timezone
	return &this
}

// NewSchedule1WithDefaults instantiates a new Schedule1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchedule1WithDefaults() *Schedule1 {
	this := Schedule1{}
	return &this
}

// GetTimes returns the Times field value
func (o *Schedule1) GetTimes() []WarehouseAdvancedSyncV1 {
	if o == nil {
		var ret []WarehouseAdvancedSyncV1
		return ret
	}

	return o.Times
}

// GetTimesOk returns a tuple with the Times field value
// and a boolean to check if the value has been set.
func (o *Schedule1) GetTimesOk() ([]WarehouseAdvancedSyncV1, bool) {
	if o == nil {
		return nil, false
	}
	return o.Times, true
}

// SetTimes sets field value
func (o *Schedule1) SetTimes(v []WarehouseAdvancedSyncV1) {
	o.Times = v
}

// GetTimezone returns the Timezone field value
func (o *Schedule1) GetTimezone() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value
// and a boolean to check if the value has been set.
func (o *Schedule1) GetTimezoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timezone, true
}

// SetTimezone sets field value
func (o *Schedule1) SetTimezone(v string) {
	o.Timezone = v
}

func (o Schedule1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["times"] = o.Times
	}
	if true {
		toSerialize["timezone"] = o.Timezone
	}
	return json.Marshal(toSerialize)
}

type NullableSchedule1 struct {
	value *Schedule1
	isSet bool
}

func (v NullableSchedule1) Get() *Schedule1 {
	return v.value
}

func (v *NullableSchedule1) Set(val *Schedule1) {
	v.value = val
	v.isSet = true
}

func (v NullableSchedule1) IsSet() bool {
	return v.isSet
}

func (v *NullableSchedule1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchedule1(val *Schedule1) *NullableSchedule1 {
	return &NullableSchedule1{value: val, isSet: true}
}

func (v NullableSchedule1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchedule1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
