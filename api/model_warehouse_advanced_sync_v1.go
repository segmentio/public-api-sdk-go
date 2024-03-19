/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 45.0.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the WarehouseAdvancedSyncV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WarehouseAdvancedSyncV1{}

// WarehouseAdvancedSyncV1 Determines the time of day at which a Warehouse should sync.
type WarehouseAdvancedSyncV1 struct {
	// The hour of day for which to enable/disable a sync, between 0 and 23.
	HourOfDay float32 `json:"hourOfDay"`
	// Enable to the sync at the specified hour.
	Enabled bool `json:"enabled"`
}

// NewWarehouseAdvancedSyncV1 instantiates a new WarehouseAdvancedSyncV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWarehouseAdvancedSyncV1(hourOfDay float32, enabled bool) *WarehouseAdvancedSyncV1 {
	this := WarehouseAdvancedSyncV1{}
	this.HourOfDay = hourOfDay
	this.Enabled = enabled
	return &this
}

// NewWarehouseAdvancedSyncV1WithDefaults instantiates a new WarehouseAdvancedSyncV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWarehouseAdvancedSyncV1WithDefaults() *WarehouseAdvancedSyncV1 {
	this := WarehouseAdvancedSyncV1{}
	return &this
}

// GetHourOfDay returns the HourOfDay field value
func (o *WarehouseAdvancedSyncV1) GetHourOfDay() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.HourOfDay
}

// GetHourOfDayOk returns a tuple with the HourOfDay field value
// and a boolean to check if the value has been set.
func (o *WarehouseAdvancedSyncV1) GetHourOfDayOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HourOfDay, true
}

// SetHourOfDay sets field value
func (o *WarehouseAdvancedSyncV1) SetHourOfDay(v float32) {
	o.HourOfDay = v
}

// GetEnabled returns the Enabled field value
func (o *WarehouseAdvancedSyncV1) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *WarehouseAdvancedSyncV1) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *WarehouseAdvancedSyncV1) SetEnabled(v bool) {
	o.Enabled = v
}

func (o WarehouseAdvancedSyncV1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WarehouseAdvancedSyncV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["hourOfDay"] = o.HourOfDay
	toSerialize["enabled"] = o.Enabled
	return toSerialize, nil
}

type NullableWarehouseAdvancedSyncV1 struct {
	value *WarehouseAdvancedSyncV1
	isSet bool
}

func (v NullableWarehouseAdvancedSyncV1) Get() *WarehouseAdvancedSyncV1 {
	return v.value
}

func (v *NullableWarehouseAdvancedSyncV1) Set(val *WarehouseAdvancedSyncV1) {
	v.value = val
	v.isSet = true
}

func (v NullableWarehouseAdvancedSyncV1) IsSet() bool {
	return v.isSet
}

func (v *NullableWarehouseAdvancedSyncV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWarehouseAdvancedSyncV1(
	val *WarehouseAdvancedSyncV1,
) *NullableWarehouseAdvancedSyncV1 {
	return &NullableWarehouseAdvancedSyncV1{value: val, isSet: true}
}

func (v NullableWarehouseAdvancedSyncV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWarehouseAdvancedSyncV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
