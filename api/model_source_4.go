/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API. 

API version: 32.0.4
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// Source4 The returned Source object.
type Source4 struct {
	// The id of the Source.  Config API note: analogous to `name`.
	Id string `json:"id"`
	// The slug used to identify the Source in the Segment app.  Config API note: equal to `name`.
	Slug string `json:"slug"`
	// The name of the Source.  Config API note: equal to `displayName`.
	Name *string `json:"name,omitempty"`
	Metadata Metadata1 `json:"metadata"`
	// The id of the Workspace that owns the Source.  Config API note: equal to `parent`.
	WorkspaceId string `json:"workspaceId"`
	// Enable to receive data from the Source.
	Enabled bool `json:"enabled"`
	// The write keys used to send data from the Source. This field is left empty when the current token does not have the 'source admin' permission.
	WriteKeys []string `json:"writeKeys"`
	// The settings associated with the Source.
	Settings NullableModelMap `json:"settings,omitempty"`
	// A list of labels applied to the Source.
	Labels []LabelV1 `json:"labels"`
}

// NewSource4 instantiates a new Source4 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSource4(id string, slug string, metadata Metadata1, workspaceId string, enabled bool, writeKeys []string, labels []LabelV1) *Source4 {
	this := Source4{}
	this.Id = id
	this.Slug = slug
	this.Metadata = metadata
	this.WorkspaceId = workspaceId
	this.Enabled = enabled
	this.WriteKeys = writeKeys
	this.Labels = labels
	return &this
}

// NewSource4WithDefaults instantiates a new Source4 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSource4WithDefaults() *Source4 {
	this := Source4{}
	return &this
}

// GetId returns the Id field value
func (o *Source4) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Source4) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Source4) SetId(v string) {
	o.Id = v
}

// GetSlug returns the Slug field value
func (o *Source4) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *Source4) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *Source4) SetSlug(v string) {
	o.Slug = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Source4) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Source4) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Source4) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Source4) SetName(v string) {
	o.Name = &v
}

// GetMetadata returns the Metadata field value
func (o *Source4) GetMetadata() Metadata1 {
	if o == nil {
		var ret Metadata1
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *Source4) GetMetadataOk() (*Metadata1, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *Source4) SetMetadata(v Metadata1) {
	o.Metadata = v
}

// GetWorkspaceId returns the WorkspaceId field value
func (o *Source4) GetWorkspaceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WorkspaceId
}

// GetWorkspaceIdOk returns a tuple with the WorkspaceId field value
// and a boolean to check if the value has been set.
func (o *Source4) GetWorkspaceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WorkspaceId, true
}

// SetWorkspaceId sets field value
func (o *Source4) SetWorkspaceId(v string) {
	o.WorkspaceId = v
}

// GetEnabled returns the Enabled field value
func (o *Source4) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *Source4) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *Source4) SetEnabled(v bool) {
	o.Enabled = v
}

// GetWriteKeys returns the WriteKeys field value
func (o *Source4) GetWriteKeys() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.WriteKeys
}

// GetWriteKeysOk returns a tuple with the WriteKeys field value
// and a boolean to check if the value has been set.
func (o *Source4) GetWriteKeysOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.WriteKeys, true
}

// SetWriteKeys sets field value
func (o *Source4) SetWriteKeys(v []string) {
	o.WriteKeys = v
}

// GetSettings returns the Settings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Source4) GetSettings() ModelMap {
	if o == nil || o.Settings.Get() == nil {
		var ret ModelMap
		return ret
	}
	return *o.Settings.Get()
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Source4) GetSettingsOk() (*ModelMap, bool) {
	if o == nil {
		return nil, false
	}
	return o.Settings.Get(), o.Settings.IsSet()
}

// HasSettings returns a boolean if a field has been set.
func (o *Source4) HasSettings() bool {
	if o != nil && o.Settings.IsSet() {
		return true
	}

	return false
}

// SetSettings gets a reference to the given NullableModelMap and assigns it to the Settings field.
func (o *Source4) SetSettings(v ModelMap) {
	o.Settings.Set(&v)
}
// SetSettingsNil sets the value for Settings to be an explicit nil
func (o *Source4) SetSettingsNil() {
	o.Settings.Set(nil)
}

// UnsetSettings ensures that no value is present for Settings, not even an explicit nil
func (o *Source4) UnsetSettings() {
	o.Settings.Unset()
}

// GetLabels returns the Labels field value
func (o *Source4) GetLabels() []LabelV1 {
	if o == nil {
		var ret []LabelV1
		return ret
	}

	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value
// and a boolean to check if the value has been set.
func (o *Source4) GetLabelsOk() ([]LabelV1, bool) {
	if o == nil {
		return nil, false
	}
	return o.Labels, true
}

// SetLabels sets field value
func (o *Source4) SetLabels(v []LabelV1) {
	o.Labels = v
}

func (o Source4) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["slug"] = o.Slug
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["metadata"] = o.Metadata
	}
	if true {
		toSerialize["workspaceId"] = o.WorkspaceId
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	if true {
		toSerialize["writeKeys"] = o.WriteKeys
	}
	if o.Settings.IsSet() {
		toSerialize["settings"] = o.Settings.Get()
	}
	if true {
		toSerialize["labels"] = o.Labels
	}
	return json.Marshal(toSerialize)
}

type NullableSource4 struct {
	value *Source4
	isSet bool
}

func (v NullableSource4) Get() *Source4 {
	return v.value
}

func (v *NullableSource4) Set(val *Source4) {
	v.value = val
	v.isSet = true
}

func (v NullableSource4) IsSet() bool {
	return v.isSet
}

func (v *NullableSource4) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSource4(val *Source4) *NullableSource4 {
	return &NullableSource4{value: val, isSet: true}
}

func (v NullableSource4) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSource4) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


