/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API. 

API version: 32.0.8
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// UpdateSourceV1Input Updates an existing Source based on a set of parameters.
type UpdateSourceV1Input struct {
	// An optional human-readable name to associate with the Source.  Config API note: equal to `displayName`.
	Name *string `json:"name,omitempty"`
	// Enable to allow the Source to send data.
	Enabled *bool `json:"enabled,omitempty"`
	// The slug that identifies the Source.  Config API note: equal to `name`.
	Slug *string `json:"slug,omitempty"`
	// A key-value object that contains instance-specific settings for the Source.  Different kinds of Sources require different kinds of input. The settings input for a Source comes from the `options` object associated with this instance of a Source.  You can find the full list of required settings by accessing the Sources catalog endpoint under `/catalog/sources`.
	Settings NullableModelMap `json:"settings,omitempty"`
}

// NewUpdateSourceV1Input instantiates a new UpdateSourceV1Input object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateSourceV1Input() *UpdateSourceV1Input {
	this := UpdateSourceV1Input{}
	return &this
}

// NewUpdateSourceV1InputWithDefaults instantiates a new UpdateSourceV1Input object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateSourceV1InputWithDefaults() *UpdateSourceV1Input {
	this := UpdateSourceV1Input{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateSourceV1Input) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSourceV1Input) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateSourceV1Input) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateSourceV1Input) SetName(v string) {
	o.Name = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *UpdateSourceV1Input) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSourceV1Input) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *UpdateSourceV1Input) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *UpdateSourceV1Input) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetSlug returns the Slug field value if set, zero value otherwise.
func (o *UpdateSourceV1Input) GetSlug() string {
	if o == nil || o.Slug == nil {
		var ret string
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSourceV1Input) GetSlugOk() (*string, bool) {
	if o == nil || o.Slug == nil {
		return nil, false
	}
	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *UpdateSourceV1Input) HasSlug() bool {
	if o != nil && o.Slug != nil {
		return true
	}

	return false
}

// SetSlug gets a reference to the given string and assigns it to the Slug field.
func (o *UpdateSourceV1Input) SetSlug(v string) {
	o.Slug = &v
}

// GetSettings returns the Settings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateSourceV1Input) GetSettings() ModelMap {
	if o == nil || o.Settings.Get() == nil {
		var ret ModelMap
		return ret
	}
	return *o.Settings.Get()
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateSourceV1Input) GetSettingsOk() (*ModelMap, bool) {
	if o == nil {
		return nil, false
	}
	return o.Settings.Get(), o.Settings.IsSet()
}

// HasSettings returns a boolean if a field has been set.
func (o *UpdateSourceV1Input) HasSettings() bool {
	if o != nil && o.Settings.IsSet() {
		return true
	}

	return false
}

// SetSettings gets a reference to the given NullableModelMap and assigns it to the Settings field.
func (o *UpdateSourceV1Input) SetSettings(v ModelMap) {
	o.Settings.Set(&v)
}
// SetSettingsNil sets the value for Settings to be an explicit nil
func (o *UpdateSourceV1Input) SetSettingsNil() {
	o.Settings.Set(nil)
}

// UnsetSettings ensures that no value is present for Settings, not even an explicit nil
func (o *UpdateSourceV1Input) UnsetSettings() {
	o.Settings.Unset()
}

func (o UpdateSourceV1Input) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Slug != nil {
		toSerialize["slug"] = o.Slug
	}
	if o.Settings.IsSet() {
		toSerialize["settings"] = o.Settings.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateSourceV1Input struct {
	value *UpdateSourceV1Input
	isSet bool
}

func (v NullableUpdateSourceV1Input) Get() *UpdateSourceV1Input {
	return v.value
}

func (v *NullableUpdateSourceV1Input) Set(val *UpdateSourceV1Input) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateSourceV1Input) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateSourceV1Input) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateSourceV1Input(val *UpdateSourceV1Input) *NullableUpdateSourceV1Input {
	return &NullableUpdateSourceV1Input{value: val, isSet: true}
}

func (v NullableUpdateSourceV1Input) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateSourceV1Input) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


