/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 34.3.4
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// CreateSourceV1Input Create a new Source based on a set of parameters.
type CreateSourceV1Input struct {
	// The slug by which to identify the Source in the Segment app.
	Slug string `json:"slug"`
	// Enable to allow this Source to send data. Defaults to true.
	Enabled bool `json:"enabled"`
	// The id of the Source metadata from which this instance of the Source derives.  All Source metadata is available under `/catalog/sources`.
	MetadataId string `json:"metadataId"`
	// A key-value object that contains instance-specific settings for the Source.
	Settings NullableModelMap `json:"settings,omitempty"`
}

// NewCreateSourceV1Input instantiates a new CreateSourceV1Input object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSourceV1Input(slug string, enabled bool, metadataId string) *CreateSourceV1Input {
	this := CreateSourceV1Input{}
	this.Slug = slug
	this.Enabled = enabled
	this.MetadataId = metadataId
	return &this
}

// NewCreateSourceV1InputWithDefaults instantiates a new CreateSourceV1Input object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSourceV1InputWithDefaults() *CreateSourceV1Input {
	this := CreateSourceV1Input{}
	return &this
}

// GetSlug returns the Slug field value
func (o *CreateSourceV1Input) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *CreateSourceV1Input) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *CreateSourceV1Input) SetSlug(v string) {
	o.Slug = v
}

// GetEnabled returns the Enabled field value
func (o *CreateSourceV1Input) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *CreateSourceV1Input) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *CreateSourceV1Input) SetEnabled(v bool) {
	o.Enabled = v
}

// GetMetadataId returns the MetadataId field value
func (o *CreateSourceV1Input) GetMetadataId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MetadataId
}

// GetMetadataIdOk returns a tuple with the MetadataId field value
// and a boolean to check if the value has been set.
func (o *CreateSourceV1Input) GetMetadataIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetadataId, true
}

// SetMetadataId sets field value
func (o *CreateSourceV1Input) SetMetadataId(v string) {
	o.MetadataId = v
}

// GetSettings returns the Settings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateSourceV1Input) GetSettings() ModelMap {
	if o == nil || o.Settings.Get() == nil {
		var ret ModelMap
		return ret
	}
	return *o.Settings.Get()
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateSourceV1Input) GetSettingsOk() (*ModelMap, bool) {
	if o == nil {
		return nil, false
	}
	return o.Settings.Get(), o.Settings.IsSet()
}

// HasSettings returns a boolean if a field has been set.
func (o *CreateSourceV1Input) HasSettings() bool {
	if o != nil && o.Settings.IsSet() {
		return true
	}

	return false
}

// SetSettings gets a reference to the given NullableModelMap and assigns it to the Settings field.
func (o *CreateSourceV1Input) SetSettings(v ModelMap) {
	o.Settings.Set(&v)
}

// SetSettingsNil sets the value for Settings to be an explicit nil
func (o *CreateSourceV1Input) SetSettingsNil() {
	o.Settings.Set(nil)
}

// UnsetSettings ensures that no value is present for Settings, not even an explicit nil
func (o *CreateSourceV1Input) UnsetSettings() {
	o.Settings.Unset()
}

func (o CreateSourceV1Input) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["slug"] = o.Slug
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	if true {
		toSerialize["metadataId"] = o.MetadataId
	}
	if o.Settings.IsSet() {
		toSerialize["settings"] = o.Settings.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableCreateSourceV1Input struct {
	value *CreateSourceV1Input
	isSet bool
}

func (v NullableCreateSourceV1Input) Get() *CreateSourceV1Input {
	return v.value
}

func (v *NullableCreateSourceV1Input) Set(val *CreateSourceV1Input) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSourceV1Input) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSourceV1Input) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSourceV1Input(val *CreateSourceV1Input) *NullableCreateSourceV1Input {
	return &NullableCreateSourceV1Input{value: val, isSet: true}
}

func (v NullableCreateSourceV1Input) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateSourceV1Input) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
