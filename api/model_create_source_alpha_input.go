/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 58.13.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the CreateSourceAlphaInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateSourceAlphaInput{}

// CreateSourceAlphaInput Create a new Source based on a set of parameters.
type CreateSourceAlphaInput struct {
	// The slug by which to identify the Source in the Segment app.
	Slug string `json:"slug"`
	// Enable to allow this Source to send data. Defaults to true.
	Enabled bool `json:"enabled"`
	// An optional human-readable name for this Source.
	Name *string `json:"name,omitempty"`
	// The id of the Source metadata from which this instance of the Source derives.  All Source metadata is available under `/catalog/sources`.
	MetadataId string `json:"metadataId"`
	// A key-value object that contains instance-specific settings for a Source.  The `options` field in the Source metadata defines the schema of this object.
	Settings map[string]interface{} `json:"settings,omitempty"`
}

// NewCreateSourceAlphaInput instantiates a new CreateSourceAlphaInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSourceAlphaInput(
	slug string,
	enabled bool,
	metadataId string,
) *CreateSourceAlphaInput {
	this := CreateSourceAlphaInput{}
	this.Slug = slug
	this.Enabled = enabled
	this.MetadataId = metadataId
	return &this
}

// NewCreateSourceAlphaInputWithDefaults instantiates a new CreateSourceAlphaInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSourceAlphaInputWithDefaults() *CreateSourceAlphaInput {
	this := CreateSourceAlphaInput{}
	return &this
}

// GetSlug returns the Slug field value
func (o *CreateSourceAlphaInput) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *CreateSourceAlphaInput) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *CreateSourceAlphaInput) SetSlug(v string) {
	o.Slug = v
}

// GetEnabled returns the Enabled field value
func (o *CreateSourceAlphaInput) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *CreateSourceAlphaInput) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *CreateSourceAlphaInput) SetEnabled(v bool) {
	o.Enabled = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateSourceAlphaInput) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSourceAlphaInput) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateSourceAlphaInput) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreateSourceAlphaInput) SetName(v string) {
	o.Name = &v
}

// GetMetadataId returns the MetadataId field value
func (o *CreateSourceAlphaInput) GetMetadataId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MetadataId
}

// GetMetadataIdOk returns a tuple with the MetadataId field value
// and a boolean to check if the value has been set.
func (o *CreateSourceAlphaInput) GetMetadataIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetadataId, true
}

// SetMetadataId sets field value
func (o *CreateSourceAlphaInput) SetMetadataId(v string) {
	o.MetadataId = v
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *CreateSourceAlphaInput) GetSettings() map[string]interface{} {
	if o == nil || IsNil(o.Settings) {
		var ret map[string]interface{}
		return ret
	}
	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSourceAlphaInput) GetSettingsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Settings) {
		return map[string]interface{}{}, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *CreateSourceAlphaInput) HasSettings() bool {
	if o != nil && !IsNil(o.Settings) {
		return true
	}

	return false
}

// SetSettings gets a reference to the given map[string]interface{} and assigns it to the Settings field.
func (o *CreateSourceAlphaInput) SetSettings(v map[string]interface{}) {
	o.Settings = v
}

func (o CreateSourceAlphaInput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateSourceAlphaInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["slug"] = o.Slug
	toSerialize["enabled"] = o.Enabled
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	toSerialize["metadataId"] = o.MetadataId
	if !IsNil(o.Settings) {
		toSerialize["settings"] = o.Settings
	}
	return toSerialize, nil
}

type NullableCreateSourceAlphaInput struct {
	value *CreateSourceAlphaInput
	isSet bool
}

func (v NullableCreateSourceAlphaInput) Get() *CreateSourceAlphaInput {
	return v.value
}

func (v *NullableCreateSourceAlphaInput) Set(val *CreateSourceAlphaInput) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSourceAlphaInput) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSourceAlphaInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSourceAlphaInput(
	val *CreateSourceAlphaInput,
) *NullableCreateSourceAlphaInput {
	return &NullableCreateSourceAlphaInput{value: val, isSet: true}
}

func (v NullableCreateSourceAlphaInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateSourceAlphaInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
