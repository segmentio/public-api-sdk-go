/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 50.3.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the ReverseEtlSpecificTimeScheduleConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReverseEtlSpecificTimeScheduleConfig{}

// ReverseEtlSpecificTimeScheduleConfig Definition for specific day and time schedule. Days is list of numbered day of the week and hours is a list of hour of the day. The corresponding Timezone is also input as string.
type ReverseEtlSpecificTimeScheduleConfig struct {
	// Days of the week.
	Days []float32 `json:"days"`
	// Hours of the day.
	Hours []float32 `json:"hours"`
	// Timezone for the specified times.
	Timezone string `json:"timezone"`
}

// NewReverseEtlSpecificTimeScheduleConfig instantiates a new ReverseEtlSpecificTimeScheduleConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReverseEtlSpecificTimeScheduleConfig(
	days []float32,
	hours []float32,
	timezone string,
) *ReverseEtlSpecificTimeScheduleConfig {
	this := ReverseEtlSpecificTimeScheduleConfig{}
	this.Days = days
	this.Hours = hours
	this.Timezone = timezone
	return &this
}

// NewReverseEtlSpecificTimeScheduleConfigWithDefaults instantiates a new ReverseEtlSpecificTimeScheduleConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReverseEtlSpecificTimeScheduleConfigWithDefaults() *ReverseEtlSpecificTimeScheduleConfig {
	this := ReverseEtlSpecificTimeScheduleConfig{}
	return &this
}

// GetDays returns the Days field value
func (o *ReverseEtlSpecificTimeScheduleConfig) GetDays() []float32 {
	if o == nil {
		var ret []float32
		return ret
	}

	return o.Days
}

// GetDaysOk returns a tuple with the Days field value
// and a boolean to check if the value has been set.
func (o *ReverseEtlSpecificTimeScheduleConfig) GetDaysOk() ([]float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Days, true
}

// SetDays sets field value
func (o *ReverseEtlSpecificTimeScheduleConfig) SetDays(v []float32) {
	o.Days = v
}

// GetHours returns the Hours field value
func (o *ReverseEtlSpecificTimeScheduleConfig) GetHours() []float32 {
	if o == nil {
		var ret []float32
		return ret
	}

	return o.Hours
}

// GetHoursOk returns a tuple with the Hours field value
// and a boolean to check if the value has been set.
func (o *ReverseEtlSpecificTimeScheduleConfig) GetHoursOk() ([]float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Hours, true
}

// SetHours sets field value
func (o *ReverseEtlSpecificTimeScheduleConfig) SetHours(v []float32) {
	o.Hours = v
}

// GetTimezone returns the Timezone field value
func (o *ReverseEtlSpecificTimeScheduleConfig) GetTimezone() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value
// and a boolean to check if the value has been set.
func (o *ReverseEtlSpecificTimeScheduleConfig) GetTimezoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timezone, true
}

// SetTimezone sets field value
func (o *ReverseEtlSpecificTimeScheduleConfig) SetTimezone(v string) {
	o.Timezone = v
}

func (o ReverseEtlSpecificTimeScheduleConfig) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReverseEtlSpecificTimeScheduleConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["days"] = o.Days
	toSerialize["hours"] = o.Hours
	toSerialize["timezone"] = o.Timezone
	return toSerialize, nil
}

type NullableReverseEtlSpecificTimeScheduleConfig struct {
	value *ReverseEtlSpecificTimeScheduleConfig
	isSet bool
}

func (v NullableReverseEtlSpecificTimeScheduleConfig) Get() *ReverseEtlSpecificTimeScheduleConfig {
	return v.value
}

func (v *NullableReverseEtlSpecificTimeScheduleConfig) Set(
	val *ReverseEtlSpecificTimeScheduleConfig,
) {
	v.value = val
	v.isSet = true
}

func (v NullableReverseEtlSpecificTimeScheduleConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableReverseEtlSpecificTimeScheduleConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReverseEtlSpecificTimeScheduleConfig(
	val *ReverseEtlSpecificTimeScheduleConfig,
) *NullableReverseEtlSpecificTimeScheduleConfig {
	return &NullableReverseEtlSpecificTimeScheduleConfig{value: val, isSet: true}
}

func (v NullableReverseEtlSpecificTimeScheduleConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReverseEtlSpecificTimeScheduleConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}