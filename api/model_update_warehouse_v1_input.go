/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 34.3.5
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// UpdateWarehouseV1Input Updates an existing Warehouse based on a set of parameters.
type UpdateWarehouseV1Input struct {
	// An optional human-readable name to associate with this Warehouse.
	Name NullableString `json:"name,omitempty"`
	// Enable to allow this Warehouse to receive data.
	Enabled *bool `json:"enabled,omitempty"`
	// A key-value object that contains instance-specific settings for a Warehouse.  Different kinds of Warehouses require different settings. The required and optional settings for a Warehouse are described in the `options` object of the associated Warehouse metadata.  You can find the full list of Warehouse metadata and related settings information in the `/catalog/warehouses` endpoint.
	Settings NullableModelMap `json:"settings"`
}

// NewUpdateWarehouseV1Input instantiates a new UpdateWarehouseV1Input object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateWarehouseV1Input(settings NullableModelMap) *UpdateWarehouseV1Input {
	this := UpdateWarehouseV1Input{}
	this.Settings = settings
	return &this
}

// NewUpdateWarehouseV1InputWithDefaults instantiates a new UpdateWarehouseV1Input object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateWarehouseV1InputWithDefaults() *UpdateWarehouseV1Input {
	this := UpdateWarehouseV1Input{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateWarehouseV1Input) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateWarehouseV1Input) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *UpdateWarehouseV1Input) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *UpdateWarehouseV1Input) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *UpdateWarehouseV1Input) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *UpdateWarehouseV1Input) UnsetName() {
	o.Name.Unset()
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *UpdateWarehouseV1Input) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWarehouseV1Input) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *UpdateWarehouseV1Input) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *UpdateWarehouseV1Input) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetSettings returns the Settings field value
// If the value is explicit nil, the zero value for ModelMap will be returned
func (o *UpdateWarehouseV1Input) GetSettings() ModelMap {
	if o == nil || o.Settings.Get() == nil {
		var ret ModelMap
		return ret
	}

	return *o.Settings.Get()
}

// GetSettingsOk returns a tuple with the Settings field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateWarehouseV1Input) GetSettingsOk() (*ModelMap, bool) {
	if o == nil {
		return nil, false
	}
	return o.Settings.Get(), o.Settings.IsSet()
}

// SetSettings sets field value
func (o *UpdateWarehouseV1Input) SetSettings(v ModelMap) {
	o.Settings.Set(&v)
}

func (o UpdateWarehouseV1Input) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if true {
		toSerialize["settings"] = o.Settings.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateWarehouseV1Input struct {
	value *UpdateWarehouseV1Input
	isSet bool
}

func (v NullableUpdateWarehouseV1Input) Get() *UpdateWarehouseV1Input {
	return v.value
}

func (v *NullableUpdateWarehouseV1Input) Set(val *UpdateWarehouseV1Input) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateWarehouseV1Input) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateWarehouseV1Input) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateWarehouseV1Input(
	val *UpdateWarehouseV1Input,
) *NullableUpdateWarehouseV1Input {
	return &NullableUpdateWarehouseV1Input{value: val, isSet: true}
}

func (v NullableUpdateWarehouseV1Input) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateWarehouseV1Input) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
