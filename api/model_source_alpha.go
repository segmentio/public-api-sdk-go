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

// checks if the SourceAlpha type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SourceAlpha{}

// SourceAlpha Defines a data Source for Segment data.
type SourceAlpha struct {
	// The id of the Source.  Config API note: analogous to `name`.
	Id string `json:"id"`
	// The slug used to identify the Source in the Segment app.  Config API note: equal to `name`.
	Slug string `json:"slug"`
	// The name of the Source.  Config API note: equal to `displayName`.
	Name     *string          `json:"name,omitempty"`
	Metadata SourceMetadataV1 `json:"metadata"`
	// The id of the Workspace that owns the Source.  Config API note: equal to `parent`.
	WorkspaceId string `json:"workspaceId"`
	// Enable to receive data from the Source.
	Enabled bool `json:"enabled"`
	// The write keys used to send data from the Source. This field is left empty when the current token does not have the 'source admin' permission.
	WriteKeys []string `json:"writeKeys"`
	// A key-value object that contains instance-specific settings for a Source.  The `options` field in the Source metadata defines the schema of this object.
	Settings map[string]interface{} `json:"settings,omitempty"`
	// A list of labels applied to the Source.
	Labels []LabelV1 `json:"labels"`
}

// NewSourceAlpha instantiates a new SourceAlpha object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSourceAlpha(
	id string,
	slug string,
	metadata SourceMetadataV1,
	workspaceId string,
	enabled bool,
	writeKeys []string,
	labels []LabelV1,
) *SourceAlpha {
	this := SourceAlpha{}
	this.Id = id
	this.Slug = slug
	this.Metadata = metadata
	this.WorkspaceId = workspaceId
	this.Enabled = enabled
	this.WriteKeys = writeKeys
	this.Labels = labels
	return &this
}

// NewSourceAlphaWithDefaults instantiates a new SourceAlpha object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSourceAlphaWithDefaults() *SourceAlpha {
	this := SourceAlpha{}
	return &this
}

// GetId returns the Id field value
func (o *SourceAlpha) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SourceAlpha) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SourceAlpha) SetId(v string) {
	o.Id = v
}

// GetSlug returns the Slug field value
func (o *SourceAlpha) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *SourceAlpha) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *SourceAlpha) SetSlug(v string) {
	o.Slug = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SourceAlpha) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceAlpha) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SourceAlpha) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SourceAlpha) SetName(v string) {
	o.Name = &v
}

// GetMetadata returns the Metadata field value
func (o *SourceAlpha) GetMetadata() SourceMetadataV1 {
	if o == nil {
		var ret SourceMetadataV1
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *SourceAlpha) GetMetadataOk() (*SourceMetadataV1, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *SourceAlpha) SetMetadata(v SourceMetadataV1) {
	o.Metadata = v
}

// GetWorkspaceId returns the WorkspaceId field value
func (o *SourceAlpha) GetWorkspaceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WorkspaceId
}

// GetWorkspaceIdOk returns a tuple with the WorkspaceId field value
// and a boolean to check if the value has been set.
func (o *SourceAlpha) GetWorkspaceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WorkspaceId, true
}

// SetWorkspaceId sets field value
func (o *SourceAlpha) SetWorkspaceId(v string) {
	o.WorkspaceId = v
}

// GetEnabled returns the Enabled field value
func (o *SourceAlpha) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *SourceAlpha) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *SourceAlpha) SetEnabled(v bool) {
	o.Enabled = v
}

// GetWriteKeys returns the WriteKeys field value
func (o *SourceAlpha) GetWriteKeys() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.WriteKeys
}

// GetWriteKeysOk returns a tuple with the WriteKeys field value
// and a boolean to check if the value has been set.
func (o *SourceAlpha) GetWriteKeysOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.WriteKeys, true
}

// SetWriteKeys sets field value
func (o *SourceAlpha) SetWriteKeys(v []string) {
	o.WriteKeys = v
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *SourceAlpha) GetSettings() map[string]interface{} {
	if o == nil || IsNil(o.Settings) {
		var ret map[string]interface{}
		return ret
	}
	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceAlpha) GetSettingsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Settings) {
		return map[string]interface{}{}, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *SourceAlpha) HasSettings() bool {
	if o != nil && !IsNil(o.Settings) {
		return true
	}

	return false
}

// SetSettings gets a reference to the given map[string]interface{} and assigns it to the Settings field.
func (o *SourceAlpha) SetSettings(v map[string]interface{}) {
	o.Settings = v
}

// GetLabels returns the Labels field value
func (o *SourceAlpha) GetLabels() []LabelV1 {
	if o == nil {
		var ret []LabelV1
		return ret
	}

	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value
// and a boolean to check if the value has been set.
func (o *SourceAlpha) GetLabelsOk() ([]LabelV1, bool) {
	if o == nil {
		return nil, false
	}
	return o.Labels, true
}

// SetLabels sets field value
func (o *SourceAlpha) SetLabels(v []LabelV1) {
	o.Labels = v
}

func (o SourceAlpha) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SourceAlpha) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["slug"] = o.Slug
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	toSerialize["metadata"] = o.Metadata
	toSerialize["workspaceId"] = o.WorkspaceId
	toSerialize["enabled"] = o.Enabled
	toSerialize["writeKeys"] = o.WriteKeys
	if !IsNil(o.Settings) {
		toSerialize["settings"] = o.Settings
	}
	toSerialize["labels"] = o.Labels
	return toSerialize, nil
}

type NullableSourceAlpha struct {
	value *SourceAlpha
	isSet bool
}

func (v NullableSourceAlpha) Get() *SourceAlpha {
	return v.value
}

func (v *NullableSourceAlpha) Set(val *SourceAlpha) {
	v.value = val
	v.isSet = true
}

func (v NullableSourceAlpha) IsSet() bool {
	return v.isSet
}

func (v *NullableSourceAlpha) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourceAlpha(val *SourceAlpha) *NullableSourceAlpha {
	return &NullableSourceAlpha{value: val, isSet: true}
}

func (v NullableSourceAlpha) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSourceAlpha) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
