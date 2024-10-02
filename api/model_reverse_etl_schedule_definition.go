/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 55.0.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the ReverseEtlScheduleDefinition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReverseEtlScheduleDefinition{}

// ReverseEtlScheduleDefinition Defines a configuration object used for scheduling, which can vary depending on the configured strategy.
type ReverseEtlScheduleDefinition struct {
	// Strategy supports three modes: Periodic, Specific Days, or Manual.
	Strategy string         `json:"strategy"`
	Config   NullableConfig `json:"config,omitempty"`
}

// NewReverseEtlScheduleDefinition instantiates a new ReverseEtlScheduleDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReverseEtlScheduleDefinition(strategy string) *ReverseEtlScheduleDefinition {
	this := ReverseEtlScheduleDefinition{}
	this.Strategy = strategy
	return &this
}

// NewReverseEtlScheduleDefinitionWithDefaults instantiates a new ReverseEtlScheduleDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReverseEtlScheduleDefinitionWithDefaults() *ReverseEtlScheduleDefinition {
	this := ReverseEtlScheduleDefinition{}
	return &this
}

// GetStrategy returns the Strategy field value
func (o *ReverseEtlScheduleDefinition) GetStrategy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Strategy
}

// GetStrategyOk returns a tuple with the Strategy field value
// and a boolean to check if the value has been set.
func (o *ReverseEtlScheduleDefinition) GetStrategyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Strategy, true
}

// SetStrategy sets field value
func (o *ReverseEtlScheduleDefinition) SetStrategy(v string) {
	o.Strategy = v
}

// GetConfig returns the Config field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReverseEtlScheduleDefinition) GetConfig() Config {
	if o == nil || IsNil(o.Config.Get()) {
		var ret Config
		return ret
	}
	return *o.Config.Get()
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReverseEtlScheduleDefinition) GetConfigOk() (*Config, bool) {
	if o == nil {
		return nil, false
	}
	return o.Config.Get(), o.Config.IsSet()
}

// HasConfig returns a boolean if a field has been set.
func (o *ReverseEtlScheduleDefinition) HasConfig() bool {
	if o != nil && o.Config.IsSet() {
		return true
	}

	return false
}

// SetConfig gets a reference to the given NullableConfig and assigns it to the Config field.
func (o *ReverseEtlScheduleDefinition) SetConfig(v Config) {
	o.Config.Set(&v)
}

// SetConfigNil sets the value for Config to be an explicit nil
func (o *ReverseEtlScheduleDefinition) SetConfigNil() {
	o.Config.Set(nil)
}

// UnsetConfig ensures that no value is present for Config, not even an explicit nil
func (o *ReverseEtlScheduleDefinition) UnsetConfig() {
	o.Config.Unset()
}

func (o ReverseEtlScheduleDefinition) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReverseEtlScheduleDefinition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["strategy"] = o.Strategy
	if o.Config.IsSet() {
		toSerialize["config"] = o.Config.Get()
	}
	return toSerialize, nil
}

type NullableReverseEtlScheduleDefinition struct {
	value *ReverseEtlScheduleDefinition
	isSet bool
}

func (v NullableReverseEtlScheduleDefinition) Get() *ReverseEtlScheduleDefinition {
	return v.value
}

func (v *NullableReverseEtlScheduleDefinition) Set(val *ReverseEtlScheduleDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableReverseEtlScheduleDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableReverseEtlScheduleDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReverseEtlScheduleDefinition(
	val *ReverseEtlScheduleDefinition,
) *NullableReverseEtlScheduleDefinition {
	return &NullableReverseEtlScheduleDefinition{value: val, isSet: true}
}

func (v NullableReverseEtlScheduleDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReverseEtlScheduleDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
