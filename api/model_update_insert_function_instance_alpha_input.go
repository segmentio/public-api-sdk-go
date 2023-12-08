/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 38.3.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the UpdateInsertFunctionInstanceAlphaInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateInsertFunctionInstanceAlphaInput{}

// UpdateInsertFunctionInstanceAlphaInput Updates an insert Function instance.
type UpdateInsertFunctionInstanceAlphaInput struct {
	// Whether this insert Function instance should be enabled for the Destination.
	Enabled *bool `json:"enabled,omitempty"`
	// Defines the display name of the insert Function instance.
	Name *string `json:"name,omitempty"`
	// An object that contains settings for this insert Function instance based on the settings present in the insert Function class.
	Settings map[string]interface{} `json:"settings"`
}

// NewUpdateInsertFunctionInstanceAlphaInput instantiates a new UpdateInsertFunctionInstanceAlphaInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateInsertFunctionInstanceAlphaInput(
	settings map[string]interface{},
) *UpdateInsertFunctionInstanceAlphaInput {
	this := UpdateInsertFunctionInstanceAlphaInput{}
	this.Settings = settings
	return &this
}

// NewUpdateInsertFunctionInstanceAlphaInputWithDefaults instantiates a new UpdateInsertFunctionInstanceAlphaInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateInsertFunctionInstanceAlphaInputWithDefaults() *UpdateInsertFunctionInstanceAlphaInput {
	this := UpdateInsertFunctionInstanceAlphaInput{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *UpdateInsertFunctionInstanceAlphaInput) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateInsertFunctionInstanceAlphaInput) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *UpdateInsertFunctionInstanceAlphaInput) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *UpdateInsertFunctionInstanceAlphaInput) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateInsertFunctionInstanceAlphaInput) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateInsertFunctionInstanceAlphaInput) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateInsertFunctionInstanceAlphaInput) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateInsertFunctionInstanceAlphaInput) SetName(v string) {
	o.Name = &v
}

// GetSettings returns the Settings field value
func (o *UpdateInsertFunctionInstanceAlphaInput) GetSettings() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value
// and a boolean to check if the value has been set.
func (o *UpdateInsertFunctionInstanceAlphaInput) GetSettingsOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Settings, true
}

// SetSettings sets field value
func (o *UpdateInsertFunctionInstanceAlphaInput) SetSettings(v map[string]interface{}) {
	o.Settings = v
}

func (o UpdateInsertFunctionInstanceAlphaInput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateInsertFunctionInstanceAlphaInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	toSerialize["settings"] = o.Settings
	return toSerialize, nil
}

type NullableUpdateInsertFunctionInstanceAlphaInput struct {
	value *UpdateInsertFunctionInstanceAlphaInput
	isSet bool
}

func (v NullableUpdateInsertFunctionInstanceAlphaInput) Get() *UpdateInsertFunctionInstanceAlphaInput {
	return v.value
}

func (v *NullableUpdateInsertFunctionInstanceAlphaInput) Set(
	val *UpdateInsertFunctionInstanceAlphaInput,
) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateInsertFunctionInstanceAlphaInput) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateInsertFunctionInstanceAlphaInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateInsertFunctionInstanceAlphaInput(
	val *UpdateInsertFunctionInstanceAlphaInput,
) *NullableUpdateInsertFunctionInstanceAlphaInput {
	return &NullableUpdateInsertFunctionInstanceAlphaInput{value: val, isSet: true}
}

func (v NullableUpdateInsertFunctionInstanceAlphaInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateInsertFunctionInstanceAlphaInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
