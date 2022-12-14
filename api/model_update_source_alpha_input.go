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

// UpdateSourceAlphaInput Updates an existing Source based on a set of parameters.
type UpdateSourceAlphaInput struct {
	// An optional human-readable name to associate with the Source.  Config API note: equal to `displayName`.
	Name *string `json:"name,omitempty"`
	// Enable to allow the Source to send data.
	Enabled *bool `json:"enabled,omitempty"`
	// The slug that identifies the Source.  Config API note: equal to `name`.
	Slug *string `json:"slug,omitempty"`
	// A key-value object that contains instance-specific settings for the Source.  Different kinds of Sources require different kinds of input. The settings input for a Source comes from the `options` object associated with this instance of a Source.  You can find the full list of required settings by accessing the Sources catalog endpoint under `/catalog/sources`.
	Settings NullableModelMap `json:"settings,omitempty"`
}

// NewUpdateSourceAlphaInput instantiates a new UpdateSourceAlphaInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateSourceAlphaInput() *UpdateSourceAlphaInput {
	this := UpdateSourceAlphaInput{}
	return &this
}

// NewUpdateSourceAlphaInputWithDefaults instantiates a new UpdateSourceAlphaInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateSourceAlphaInputWithDefaults() *UpdateSourceAlphaInput {
	this := UpdateSourceAlphaInput{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateSourceAlphaInput) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSourceAlphaInput) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateSourceAlphaInput) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateSourceAlphaInput) SetName(v string) {
	o.Name = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *UpdateSourceAlphaInput) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSourceAlphaInput) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *UpdateSourceAlphaInput) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *UpdateSourceAlphaInput) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetSlug returns the Slug field value if set, zero value otherwise.
func (o *UpdateSourceAlphaInput) GetSlug() string {
	if o == nil || o.Slug == nil {
		var ret string
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSourceAlphaInput) GetSlugOk() (*string, bool) {
	if o == nil || o.Slug == nil {
		return nil, false
	}
	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *UpdateSourceAlphaInput) HasSlug() bool {
	if o != nil && o.Slug != nil {
		return true
	}

	return false
}

// SetSlug gets a reference to the given string and assigns it to the Slug field.
func (o *UpdateSourceAlphaInput) SetSlug(v string) {
	o.Slug = &v
}

// GetSettings returns the Settings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateSourceAlphaInput) GetSettings() ModelMap {
	if o == nil || o.Settings.Get() == nil {
		var ret ModelMap
		return ret
	}
	return *o.Settings.Get()
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateSourceAlphaInput) GetSettingsOk() (*ModelMap, bool) {
	if o == nil {
		return nil, false
	}
	return o.Settings.Get(), o.Settings.IsSet()
}

// HasSettings returns a boolean if a field has been set.
func (o *UpdateSourceAlphaInput) HasSettings() bool {
	if o != nil && o.Settings.IsSet() {
		return true
	}

	return false
}

// SetSettings gets a reference to the given NullableModelMap and assigns it to the Settings field.
func (o *UpdateSourceAlphaInput) SetSettings(v ModelMap) {
	o.Settings.Set(&v)
}

// SetSettingsNil sets the value for Settings to be an explicit nil
func (o *UpdateSourceAlphaInput) SetSettingsNil() {
	o.Settings.Set(nil)
}

// UnsetSettings ensures that no value is present for Settings, not even an explicit nil
func (o *UpdateSourceAlphaInput) UnsetSettings() {
	o.Settings.Unset()
}

func (o UpdateSourceAlphaInput) MarshalJSON() ([]byte, error) {
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

type NullableUpdateSourceAlphaInput struct {
	value *UpdateSourceAlphaInput
	isSet bool
}

func (v NullableUpdateSourceAlphaInput) Get() *UpdateSourceAlphaInput {
	return v.value
}

func (v *NullableUpdateSourceAlphaInput) Set(val *UpdateSourceAlphaInput) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateSourceAlphaInput) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateSourceAlphaInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateSourceAlphaInput(
	val *UpdateSourceAlphaInput,
) *NullableUpdateSourceAlphaInput {
	return &NullableUpdateSourceAlphaInput{value: val, isSet: true}
}

func (v NullableUpdateSourceAlphaInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateSourceAlphaInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
