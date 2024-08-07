/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 52.0.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the ScheduleConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScheduleConfig{}

// ScheduleConfig Depending on the chosen strategy, configures the schedule for this model.
type ScheduleConfig struct {
	// Duration is specified as a string, eg: 15m, 3h25m30s.
	Interval string `json:"interval"`
}

// NewScheduleConfig instantiates a new ScheduleConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScheduleConfig(interval string) *ScheduleConfig {
	this := ScheduleConfig{}
	this.Interval = interval
	return &this
}

// NewScheduleConfigWithDefaults instantiates a new ScheduleConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScheduleConfigWithDefaults() *ScheduleConfig {
	this := ScheduleConfig{}
	return &this
}

// GetInterval returns the Interval field value
func (o *ScheduleConfig) GetInterval() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value
// and a boolean to check if the value has been set.
func (o *ScheduleConfig) GetIntervalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Interval, true
}

// SetInterval sets field value
func (o *ScheduleConfig) SetInterval(v string) {
	o.Interval = v
}

func (o ScheduleConfig) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScheduleConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["interval"] = o.Interval
	return toSerialize, nil
}

type NullableScheduleConfig struct {
	value *ScheduleConfig
	isSet bool
}

func (v NullableScheduleConfig) Get() *ScheduleConfig {
	return v.value
}

func (v *NullableScheduleConfig) Set(val *ScheduleConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableScheduleConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableScheduleConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScheduleConfig(val *ScheduleConfig) *NullableScheduleConfig {
	return &NullableScheduleConfig{value: val, isSet: true}
}

func (v NullableScheduleConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScheduleConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
