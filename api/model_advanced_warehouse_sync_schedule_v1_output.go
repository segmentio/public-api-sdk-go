/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 57.0.1
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the AdvancedWarehouseSyncScheduleV1Output type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdvancedWarehouseSyncScheduleV1Output{}

// AdvancedWarehouseSyncScheduleV1Output Defines the advanced sync schedule for a Warehouse.
type AdvancedWarehouseSyncScheduleV1Output struct {
	// A list that contains the times when syncs should occur.
	Times []WarehouseAdvancedSyncV1 `json:"times,omitempty"`
	// A TZ-database timezone for this sync schedule.
	Timezone *string `json:"timezone,omitempty"`
}

// NewAdvancedWarehouseSyncScheduleV1Output instantiates a new AdvancedWarehouseSyncScheduleV1Output object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdvancedWarehouseSyncScheduleV1Output() *AdvancedWarehouseSyncScheduleV1Output {
	this := AdvancedWarehouseSyncScheduleV1Output{}
	return &this
}

// NewAdvancedWarehouseSyncScheduleV1OutputWithDefaults instantiates a new AdvancedWarehouseSyncScheduleV1Output object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdvancedWarehouseSyncScheduleV1OutputWithDefaults() *AdvancedWarehouseSyncScheduleV1Output {
	this := AdvancedWarehouseSyncScheduleV1Output{}
	return &this
}

// GetTimes returns the Times field value if set, zero value otherwise.
func (o *AdvancedWarehouseSyncScheduleV1Output) GetTimes() []WarehouseAdvancedSyncV1 {
	if o == nil || IsNil(o.Times) {
		var ret []WarehouseAdvancedSyncV1
		return ret
	}
	return o.Times
}

// GetTimesOk returns a tuple with the Times field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvancedWarehouseSyncScheduleV1Output) GetTimesOk() ([]WarehouseAdvancedSyncV1, bool) {
	if o == nil || IsNil(o.Times) {
		return nil, false
	}
	return o.Times, true
}

// HasTimes returns a boolean if a field has been set.
func (o *AdvancedWarehouseSyncScheduleV1Output) HasTimes() bool {
	if o != nil && !IsNil(o.Times) {
		return true
	}

	return false
}

// SetTimes gets a reference to the given []WarehouseAdvancedSyncV1 and assigns it to the Times field.
func (o *AdvancedWarehouseSyncScheduleV1Output) SetTimes(v []WarehouseAdvancedSyncV1) {
	o.Times = v
}

// GetTimezone returns the Timezone field value if set, zero value otherwise.
func (o *AdvancedWarehouseSyncScheduleV1Output) GetTimezone() string {
	if o == nil || IsNil(o.Timezone) {
		var ret string
		return ret
	}
	return *o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvancedWarehouseSyncScheduleV1Output) GetTimezoneOk() (*string, bool) {
	if o == nil || IsNil(o.Timezone) {
		return nil, false
	}
	return o.Timezone, true
}

// HasTimezone returns a boolean if a field has been set.
func (o *AdvancedWarehouseSyncScheduleV1Output) HasTimezone() bool {
	if o != nil && !IsNil(o.Timezone) {
		return true
	}

	return false
}

// SetTimezone gets a reference to the given string and assigns it to the Timezone field.
func (o *AdvancedWarehouseSyncScheduleV1Output) SetTimezone(v string) {
	o.Timezone = &v
}

func (o AdvancedWarehouseSyncScheduleV1Output) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdvancedWarehouseSyncScheduleV1Output) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Times) {
		toSerialize["times"] = o.Times
	}
	if !IsNil(o.Timezone) {
		toSerialize["timezone"] = o.Timezone
	}
	return toSerialize, nil
}

type NullableAdvancedWarehouseSyncScheduleV1Output struct {
	value *AdvancedWarehouseSyncScheduleV1Output
	isSet bool
}

func (v NullableAdvancedWarehouseSyncScheduleV1Output) Get() *AdvancedWarehouseSyncScheduleV1Output {
	return v.value
}

func (v *NullableAdvancedWarehouseSyncScheduleV1Output) Set(
	val *AdvancedWarehouseSyncScheduleV1Output,
) {
	v.value = val
	v.isSet = true
}

func (v NullableAdvancedWarehouseSyncScheduleV1Output) IsSet() bool {
	return v.isSet
}

func (v *NullableAdvancedWarehouseSyncScheduleV1Output) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdvancedWarehouseSyncScheduleV1Output(
	val *AdvancedWarehouseSyncScheduleV1Output,
) *NullableAdvancedWarehouseSyncScheduleV1Output {
	return &NullableAdvancedWarehouseSyncScheduleV1Output{value: val, isSet: true}
}

func (v NullableAdvancedWarehouseSyncScheduleV1Output) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdvancedWarehouseSyncScheduleV1Output) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
