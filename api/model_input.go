/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 33.0.4
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// Input A set of valid Destination input params required for updating.
type Input struct {
	// The user-defined name for the subscription.
	Name *string `json:"name,omitempty"`
	// The fql statement.
	Trigger *string `json:"trigger,omitempty"`
	// Is the subscription enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// The fields used for configuring this action.
	Settings NullableModelMap `json:"settings,omitempty"`
}

// NewInput instantiates a new Input object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInput() *Input {
	this := Input{}
	return &this
}

// NewInputWithDefaults instantiates a new Input object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInputWithDefaults() *Input {
	this := Input{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Input) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Input) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Input) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Input) SetName(v string) {
	o.Name = &v
}

// GetTrigger returns the Trigger field value if set, zero value otherwise.
func (o *Input) GetTrigger() string {
	if o == nil || o.Trigger == nil {
		var ret string
		return ret
	}
	return *o.Trigger
}

// GetTriggerOk returns a tuple with the Trigger field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Input) GetTriggerOk() (*string, bool) {
	if o == nil || o.Trigger == nil {
		return nil, false
	}
	return o.Trigger, true
}

// HasTrigger returns a boolean if a field has been set.
func (o *Input) HasTrigger() bool {
	if o != nil && o.Trigger != nil {
		return true
	}

	return false
}

// SetTrigger gets a reference to the given string and assigns it to the Trigger field.
func (o *Input) SetTrigger(v string) {
	o.Trigger = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *Input) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Input) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *Input) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *Input) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetSettings returns the Settings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Input) GetSettings() ModelMap {
	if o == nil || o.Settings.Get() == nil {
		var ret ModelMap
		return ret
	}
	return *o.Settings.Get()
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Input) GetSettingsOk() (*ModelMap, bool) {
	if o == nil {
		return nil, false
	}
	return o.Settings.Get(), o.Settings.IsSet()
}

// HasSettings returns a boolean if a field has been set.
func (o *Input) HasSettings() bool {
	if o != nil && o.Settings.IsSet() {
		return true
	}

	return false
}

// SetSettings gets a reference to the given NullableModelMap and assigns it to the Settings field.
func (o *Input) SetSettings(v ModelMap) {
	o.Settings.Set(&v)
}

// SetSettingsNil sets the value for Settings to be an explicit nil
func (o *Input) SetSettingsNil() {
	o.Settings.Set(nil)
}

// UnsetSettings ensures that no value is present for Settings, not even an explicit nil
func (o *Input) UnsetSettings() {
	o.Settings.Unset()
}

func (o Input) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Trigger != nil {
		toSerialize["trigger"] = o.Trigger
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Settings.IsSet() {
		toSerialize["settings"] = o.Settings.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableInput struct {
	value *Input
	isSet bool
}

func (v NullableInput) Get() *Input {
	return v.value
}

func (v *NullableInput) Set(val *Input) {
	v.value = val
	v.isSet = true
}

func (v NullableInput) IsSet() bool {
	return v.isSet
}

func (v *NullableInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInput(val *Input) *NullableInput {
	return &NullableInput{value: val, isSet: true}
}

func (v NullableInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
