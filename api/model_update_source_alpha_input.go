/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 54.0.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the UpdateSourceAlphaInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateSourceAlphaInput{}

// UpdateSourceAlphaInput Updates an existing Source based on a set of parameters.
type UpdateSourceAlphaInput struct {
	// An optional human-readable name to associate with the Source.  Config API note: equal to `displayName`.
	Name *string `json:"name,omitempty"`
	// Enable to allow the Source to send data.
	Enabled *bool `json:"enabled,omitempty"`
	// The slug that identifies the Source.  Config API note: equal to `name`.
	Slug *string `json:"slug,omitempty"`
	// A key-value object that contains instance-specific settings for a Source.  The `options` field in the Source metadata defines the schema of this object.
	Settings map[string]interface{} `json:"settings,omitempty"`
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
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSourceAlphaInput) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateSourceAlphaInput) HasName() bool {
	if o != nil && !IsNil(o.Name) {
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
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSourceAlphaInput) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *UpdateSourceAlphaInput) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
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
	if o == nil || IsNil(o.Slug) {
		var ret string
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSourceAlphaInput) GetSlugOk() (*string, bool) {
	if o == nil || IsNil(o.Slug) {
		return nil, false
	}
	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *UpdateSourceAlphaInput) HasSlug() bool {
	if o != nil && !IsNil(o.Slug) {
		return true
	}

	return false
}

// SetSlug gets a reference to the given string and assigns it to the Slug field.
func (o *UpdateSourceAlphaInput) SetSlug(v string) {
	o.Slug = &v
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *UpdateSourceAlphaInput) GetSettings() map[string]interface{} {
	if o == nil || IsNil(o.Settings) {
		var ret map[string]interface{}
		return ret
	}
	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSourceAlphaInput) GetSettingsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Settings) {
		return map[string]interface{}{}, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *UpdateSourceAlphaInput) HasSettings() bool {
	if o != nil && !IsNil(o.Settings) {
		return true
	}

	return false
}

// SetSettings gets a reference to the given map[string]interface{} and assigns it to the Settings field.
func (o *UpdateSourceAlphaInput) SetSettings(v map[string]interface{}) {
	o.Settings = v
}

func (o UpdateSourceAlphaInput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateSourceAlphaInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Slug) {
		toSerialize["slug"] = o.Slug
	}
	if !IsNil(o.Settings) {
		toSerialize["settings"] = o.Settings
	}
	return toSerialize, nil
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
