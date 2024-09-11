/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 54.3.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the ReverseEtlCronScheduleConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReverseEtlCronScheduleConfig{}

// ReverseEtlCronScheduleConfig struct for ReverseEtlCronScheduleConfig
type ReverseEtlCronScheduleConfig struct {
	// 5 field cron string expression. The cron expression must be larger than 15 minutes.
	Spec string `json:"spec"`
	// Timezone for the specified times.
	Timezone string `json:"timezone"`
}

// NewReverseEtlCronScheduleConfig instantiates a new ReverseEtlCronScheduleConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReverseEtlCronScheduleConfig(spec string, timezone string) *ReverseEtlCronScheduleConfig {
	this := ReverseEtlCronScheduleConfig{}
	this.Spec = spec
	this.Timezone = timezone
	return &this
}

// NewReverseEtlCronScheduleConfigWithDefaults instantiates a new ReverseEtlCronScheduleConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReverseEtlCronScheduleConfigWithDefaults() *ReverseEtlCronScheduleConfig {
	this := ReverseEtlCronScheduleConfig{}
	return &this
}

// GetSpec returns the Spec field value
func (o *ReverseEtlCronScheduleConfig) GetSpec() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Spec
}

// GetSpecOk returns a tuple with the Spec field value
// and a boolean to check if the value has been set.
func (o *ReverseEtlCronScheduleConfig) GetSpecOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Spec, true
}

// SetSpec sets field value
func (o *ReverseEtlCronScheduleConfig) SetSpec(v string) {
	o.Spec = v
}

// GetTimezone returns the Timezone field value
func (o *ReverseEtlCronScheduleConfig) GetTimezone() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value
// and a boolean to check if the value has been set.
func (o *ReverseEtlCronScheduleConfig) GetTimezoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timezone, true
}

// SetTimezone sets field value
func (o *ReverseEtlCronScheduleConfig) SetTimezone(v string) {
	o.Timezone = v
}

func (o ReverseEtlCronScheduleConfig) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReverseEtlCronScheduleConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["spec"] = o.Spec
	toSerialize["timezone"] = o.Timezone
	return toSerialize, nil
}

type NullableReverseEtlCronScheduleConfig struct {
	value *ReverseEtlCronScheduleConfig
	isSet bool
}

func (v NullableReverseEtlCronScheduleConfig) Get() *ReverseEtlCronScheduleConfig {
	return v.value
}

func (v *NullableReverseEtlCronScheduleConfig) Set(val *ReverseEtlCronScheduleConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableReverseEtlCronScheduleConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableReverseEtlCronScheduleConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReverseEtlCronScheduleConfig(
	val *ReverseEtlCronScheduleConfig,
) *NullableReverseEtlCronScheduleConfig {
	return &NullableReverseEtlCronScheduleConfig{value: val, isSet: true}
}

func (v NullableReverseEtlCronScheduleConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReverseEtlCronScheduleConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
