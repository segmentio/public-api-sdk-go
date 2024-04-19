/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 49.0.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the AdvancedWarehouseSyncScheduleV1Input type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdvancedWarehouseSyncScheduleV1Input{}

// AdvancedWarehouseSyncScheduleV1Input Defines the advanced sync schedule for a Warehouse.
type AdvancedWarehouseSyncScheduleV1Input struct {
	// A list that contains the times when syncs should occur.
	Times []WarehouseAdvancedSyncV1 `json:"times"`
	// A TZ-database timezone for this sync schedule.
	Timezone string `json:"timezone"`
}

// NewAdvancedWarehouseSyncScheduleV1Input instantiates a new AdvancedWarehouseSyncScheduleV1Input object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdvancedWarehouseSyncScheduleV1Input(
	times []WarehouseAdvancedSyncV1,
	timezone string,
) *AdvancedWarehouseSyncScheduleV1Input {
	this := AdvancedWarehouseSyncScheduleV1Input{}
	this.Times = times
	this.Timezone = timezone
	return &this
}

// NewAdvancedWarehouseSyncScheduleV1InputWithDefaults instantiates a new AdvancedWarehouseSyncScheduleV1Input object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdvancedWarehouseSyncScheduleV1InputWithDefaults() *AdvancedWarehouseSyncScheduleV1Input {
	this := AdvancedWarehouseSyncScheduleV1Input{}
	return &this
}

// GetTimes returns the Times field value
func (o *AdvancedWarehouseSyncScheduleV1Input) GetTimes() []WarehouseAdvancedSyncV1 {
	if o == nil {
		var ret []WarehouseAdvancedSyncV1
		return ret
	}

	return o.Times
}

// GetTimesOk returns a tuple with the Times field value
// and a boolean to check if the value has been set.
func (o *AdvancedWarehouseSyncScheduleV1Input) GetTimesOk() ([]WarehouseAdvancedSyncV1, bool) {
	if o == nil {
		return nil, false
	}
	return o.Times, true
}

// SetTimes sets field value
func (o *AdvancedWarehouseSyncScheduleV1Input) SetTimes(v []WarehouseAdvancedSyncV1) {
	o.Times = v
}

// GetTimezone returns the Timezone field value
func (o *AdvancedWarehouseSyncScheduleV1Input) GetTimezone() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value
// and a boolean to check if the value has been set.
func (o *AdvancedWarehouseSyncScheduleV1Input) GetTimezoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timezone, true
}

// SetTimezone sets field value
func (o *AdvancedWarehouseSyncScheduleV1Input) SetTimezone(v string) {
	o.Timezone = v
}

func (o AdvancedWarehouseSyncScheduleV1Input) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdvancedWarehouseSyncScheduleV1Input) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["times"] = o.Times
	toSerialize["timezone"] = o.Timezone
	return toSerialize, nil
}

type NullableAdvancedWarehouseSyncScheduleV1Input struct {
	value *AdvancedWarehouseSyncScheduleV1Input
	isSet bool
}

func (v NullableAdvancedWarehouseSyncScheduleV1Input) Get() *AdvancedWarehouseSyncScheduleV1Input {
	return v.value
}

func (v *NullableAdvancedWarehouseSyncScheduleV1Input) Set(
	val *AdvancedWarehouseSyncScheduleV1Input,
) {
	v.value = val
	v.isSet = true
}

func (v NullableAdvancedWarehouseSyncScheduleV1Input) IsSet() bool {
	return v.isSet
}

func (v *NullableAdvancedWarehouseSyncScheduleV1Input) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdvancedWarehouseSyncScheduleV1Input(
	val *AdvancedWarehouseSyncScheduleV1Input,
) *NullableAdvancedWarehouseSyncScheduleV1Input {
	return &NullableAdvancedWarehouseSyncScheduleV1Input{value: val, isSet: true}
}

func (v NullableAdvancedWarehouseSyncScheduleV1Input) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdvancedWarehouseSyncScheduleV1Input) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
