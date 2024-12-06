/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 57.1.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the ProfilesWarehouseAlpha type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProfilesWarehouseAlpha{}

// ProfilesWarehouseAlpha Defines a Profiles data Warehouse used as a Destination for Segment data.
type ProfilesWarehouseAlpha struct {
	// The id of the Warehouse.
	Id string `json:"id"`
	// The Space id.
	SpaceId  string              `json:"spaceId"`
	Metadata WarehouseMetadataV1 `json:"metadata"`
	// The id of the Workspace that owns this Warehouse.
	WorkspaceId string `json:"workspaceId"`
	// When set to true, this Warehouse receives data.
	Enabled bool `json:"enabled"`
	// A key-value object that contains instance-specific Warehouse settings.
	Settings map[string]interface{} `json:"settings"`
	// The custom schema name that Segment uses on the Warehouse side.
	SchemaName *string `json:"schemaName,omitempty"`
}

// NewProfilesWarehouseAlpha instantiates a new ProfilesWarehouseAlpha object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProfilesWarehouseAlpha(
	id string,
	spaceId string,
	metadata WarehouseMetadataV1,
	workspaceId string,
	enabled bool,
	settings map[string]interface{},
) *ProfilesWarehouseAlpha {
	this := ProfilesWarehouseAlpha{}
	this.Id = id
	this.SpaceId = spaceId
	this.Metadata = metadata
	this.WorkspaceId = workspaceId
	this.Enabled = enabled
	this.Settings = settings
	return &this
}

// NewProfilesWarehouseAlphaWithDefaults instantiates a new ProfilesWarehouseAlpha object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProfilesWarehouseAlphaWithDefaults() *ProfilesWarehouseAlpha {
	this := ProfilesWarehouseAlpha{}
	return &this
}

// GetId returns the Id field value
func (o *ProfilesWarehouseAlpha) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ProfilesWarehouseAlpha) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ProfilesWarehouseAlpha) SetId(v string) {
	o.Id = v
}

// GetSpaceId returns the SpaceId field value
func (o *ProfilesWarehouseAlpha) GetSpaceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SpaceId
}

// GetSpaceIdOk returns a tuple with the SpaceId field value
// and a boolean to check if the value has been set.
func (o *ProfilesWarehouseAlpha) GetSpaceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SpaceId, true
}

// SetSpaceId sets field value
func (o *ProfilesWarehouseAlpha) SetSpaceId(v string) {
	o.SpaceId = v
}

// GetMetadata returns the Metadata field value
func (o *ProfilesWarehouseAlpha) GetMetadata() WarehouseMetadataV1 {
	if o == nil {
		var ret WarehouseMetadataV1
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *ProfilesWarehouseAlpha) GetMetadataOk() (*WarehouseMetadataV1, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *ProfilesWarehouseAlpha) SetMetadata(v WarehouseMetadataV1) {
	o.Metadata = v
}

// GetWorkspaceId returns the WorkspaceId field value
func (o *ProfilesWarehouseAlpha) GetWorkspaceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WorkspaceId
}

// GetWorkspaceIdOk returns a tuple with the WorkspaceId field value
// and a boolean to check if the value has been set.
func (o *ProfilesWarehouseAlpha) GetWorkspaceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WorkspaceId, true
}

// SetWorkspaceId sets field value
func (o *ProfilesWarehouseAlpha) SetWorkspaceId(v string) {
	o.WorkspaceId = v
}

// GetEnabled returns the Enabled field value
func (o *ProfilesWarehouseAlpha) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *ProfilesWarehouseAlpha) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *ProfilesWarehouseAlpha) SetEnabled(v bool) {
	o.Enabled = v
}

// GetSettings returns the Settings field value
func (o *ProfilesWarehouseAlpha) GetSettings() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value
// and a boolean to check if the value has been set.
func (o *ProfilesWarehouseAlpha) GetSettingsOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Settings, true
}

// SetSettings sets field value
func (o *ProfilesWarehouseAlpha) SetSettings(v map[string]interface{}) {
	o.Settings = v
}

// GetSchemaName returns the SchemaName field value if set, zero value otherwise.
func (o *ProfilesWarehouseAlpha) GetSchemaName() string {
	if o == nil || IsNil(o.SchemaName) {
		var ret string
		return ret
	}
	return *o.SchemaName
}

// GetSchemaNameOk returns a tuple with the SchemaName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfilesWarehouseAlpha) GetSchemaNameOk() (*string, bool) {
	if o == nil || IsNil(o.SchemaName) {
		return nil, false
	}
	return o.SchemaName, true
}

// HasSchemaName returns a boolean if a field has been set.
func (o *ProfilesWarehouseAlpha) HasSchemaName() bool {
	if o != nil && !IsNil(o.SchemaName) {
		return true
	}

	return false
}

// SetSchemaName gets a reference to the given string and assigns it to the SchemaName field.
func (o *ProfilesWarehouseAlpha) SetSchemaName(v string) {
	o.SchemaName = &v
}

func (o ProfilesWarehouseAlpha) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProfilesWarehouseAlpha) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["spaceId"] = o.SpaceId
	toSerialize["metadata"] = o.Metadata
	toSerialize["workspaceId"] = o.WorkspaceId
	toSerialize["enabled"] = o.Enabled
	toSerialize["settings"] = o.Settings
	if !IsNil(o.SchemaName) {
		toSerialize["schemaName"] = o.SchemaName
	}
	return toSerialize, nil
}

type NullableProfilesWarehouseAlpha struct {
	value *ProfilesWarehouseAlpha
	isSet bool
}

func (v NullableProfilesWarehouseAlpha) Get() *ProfilesWarehouseAlpha {
	return v.value
}

func (v *NullableProfilesWarehouseAlpha) Set(val *ProfilesWarehouseAlpha) {
	v.value = val
	v.isSet = true
}

func (v NullableProfilesWarehouseAlpha) IsSet() bool {
	return v.isSet
}

func (v *NullableProfilesWarehouseAlpha) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProfilesWarehouseAlpha(
	val *ProfilesWarehouseAlpha,
) *NullableProfilesWarehouseAlpha {
	return &NullableProfilesWarehouseAlpha{value: val, isSet: true}
}

func (v NullableProfilesWarehouseAlpha) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProfilesWarehouseAlpha) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
