/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 37.2.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// Source5 The newly created Source.
type Source5 struct {
	// The id of the Source.  Config API note: analogous to `name`.
	Id string `json:"id"`
	// The slug used to identify the Source in the Segment app.  Config API note: equal to `name`.
	Slug string `json:"slug"`
	// The name of the Source.  Config API note: equal to `displayName`.
	Name     *string   `json:"name,omitempty"`
	Metadata Metadata2 `json:"metadata"`
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

// NewSource5 instantiates a new Source5 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSource5(
	id string,
	slug string,
	metadata Metadata2,
	workspaceId string,
	enabled bool,
	writeKeys []string,
	labels []LabelV1,
) *Source5 {
	this := Source5{}
	this.Id = id
	this.Slug = slug
	this.Metadata = metadata
	this.WorkspaceId = workspaceId
	this.Enabled = enabled
	this.WriteKeys = writeKeys
	this.Labels = labels
	return &this
}

// NewSource5WithDefaults instantiates a new Source5 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSource5WithDefaults() *Source5 {
	this := Source5{}
	return &this
}

// GetId returns the Id field value
func (o *Source5) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Source5) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Source5) SetId(v string) {
	o.Id = v
}

// GetSlug returns the Slug field value
func (o *Source5) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *Source5) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *Source5) SetSlug(v string) {
	o.Slug = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Source5) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Source5) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Source5) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Source5) SetName(v string) {
	o.Name = &v
}

// GetMetadata returns the Metadata field value
func (o *Source5) GetMetadata() Metadata2 {
	if o == nil {
		var ret Metadata2
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *Source5) GetMetadataOk() (*Metadata2, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *Source5) SetMetadata(v Metadata2) {
	o.Metadata = v
}

// GetWorkspaceId returns the WorkspaceId field value
func (o *Source5) GetWorkspaceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WorkspaceId
}

// GetWorkspaceIdOk returns a tuple with the WorkspaceId field value
// and a boolean to check if the value has been set.
func (o *Source5) GetWorkspaceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WorkspaceId, true
}

// SetWorkspaceId sets field value
func (o *Source5) SetWorkspaceId(v string) {
	o.WorkspaceId = v
}

// GetEnabled returns the Enabled field value
func (o *Source5) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *Source5) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *Source5) SetEnabled(v bool) {
	o.Enabled = v
}

// GetWriteKeys returns the WriteKeys field value
func (o *Source5) GetWriteKeys() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.WriteKeys
}

// GetWriteKeysOk returns a tuple with the WriteKeys field value
// and a boolean to check if the value has been set.
func (o *Source5) GetWriteKeysOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.WriteKeys, true
}

// SetWriteKeys sets field value
func (o *Source5) SetWriteKeys(v []string) {
	o.WriteKeys = v
}

// GetSettings returns the Settings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Source5) GetSettings() ModelMap {
	if o == nil || o.Settings.Get() == nil {
		var ret ModelMap
		return ret
	}
	return *o.Settings.Get()
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Source5) GetSettingsOk() (*ModelMap, bool) {
	if o == nil {
		return nil, false
	}
	return o.Settings.Get(), o.Settings.IsSet()
}

// HasSettings returns a boolean if a field has been set.
func (o *Source5) HasSettings() bool {
	if o != nil && o.Settings.IsSet() {
		return true
	}

	return false
}

// SetSettings gets a reference to the given NullableModelMap and assigns it to the Settings field.
func (o *Source5) SetSettings(v ModelMap) {
	o.Settings.Set(&v)
}

// SetSettingsNil sets the value for Settings to be an explicit nil
func (o *Source5) SetSettingsNil() {
	o.Settings.Set(nil)
}

// UnsetSettings ensures that no value is present for Settings, not even an explicit nil
func (o *Source5) UnsetSettings() {
	o.Settings.Unset()
}

// GetLabels returns the Labels field value
func (o *Source5) GetLabels() []LabelV1 {
	if o == nil {
		var ret []LabelV1
		return ret
	}

	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value
// and a boolean to check if the value has been set.
func (o *Source5) GetLabelsOk() ([]LabelV1, bool) {
	if o == nil {
		return nil, false
	}
	return o.Labels, true
}

// SetLabels sets field value
func (o *Source5) SetLabels(v []LabelV1) {
	o.Labels = v
}

func (o Source5) MarshalJSON() ([]byte, error) {
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

type NullableSource5 struct {
	value *Source5
	isSet bool
}

func (v NullableSource5) Get() *Source5 {
	return v.value
}

func (v *NullableSource5) Set(val *Source5) {
	v.value = val
	v.isSet = true
}

func (v NullableSource5) IsSet() bool {
	return v.isSet
}

func (v *NullableSource5) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSource5(val *Source5) *NullableSource5 {
	return &NullableSource5{value: val, isSet: true}
}

func (v NullableSource5) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSource5) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
