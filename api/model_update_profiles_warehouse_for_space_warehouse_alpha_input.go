/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 46.0.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the UpdateProfilesWarehouseForSpaceWarehouseAlphaInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateProfilesWarehouseForSpaceWarehouseAlphaInput{}

// UpdateProfilesWarehouseForSpaceWarehouseAlphaInput Updates a Profiles Warehouse based on a set of parameters.
type UpdateProfilesWarehouseForSpaceWarehouseAlphaInput struct {
	// An optional human-readable name for this Warehouse.
	Name *string `json:"name,omitempty"`
	// Enable to allow this Warehouse to receive data. Defaults to true.
	Enabled *bool `json:"enabled,omitempty"`
	// A key-value object that contains instance-specific Warehouse settings.
	Settings map[string]interface{} `json:"settings"`
	// The custom schema name that Segment uses on the Warehouse side. The space slug value is default otherwise.
	SchemaName *string `json:"schemaName,omitempty"`
}

// NewUpdateProfilesWarehouseForSpaceWarehouseAlphaInput instantiates a new UpdateProfilesWarehouseForSpaceWarehouseAlphaInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateProfilesWarehouseForSpaceWarehouseAlphaInput(
	settings map[string]interface{},
) *UpdateProfilesWarehouseForSpaceWarehouseAlphaInput {
	this := UpdateProfilesWarehouseForSpaceWarehouseAlphaInput{}
	this.Settings = settings
	return &this
}

// NewUpdateProfilesWarehouseForSpaceWarehouseAlphaInputWithDefaults instantiates a new UpdateProfilesWarehouseForSpaceWarehouseAlphaInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateProfilesWarehouseForSpaceWarehouseAlphaInputWithDefaults() *UpdateProfilesWarehouseForSpaceWarehouseAlphaInput {
	this := UpdateProfilesWarehouseForSpaceWarehouseAlphaInput{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateProfilesWarehouseForSpaceWarehouseAlphaInput) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProfilesWarehouseForSpaceWarehouseAlphaInput) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateProfilesWarehouseForSpaceWarehouseAlphaInput) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateProfilesWarehouseForSpaceWarehouseAlphaInput) SetName(v string) {
	o.Name = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *UpdateProfilesWarehouseForSpaceWarehouseAlphaInput) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProfilesWarehouseForSpaceWarehouseAlphaInput) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *UpdateProfilesWarehouseForSpaceWarehouseAlphaInput) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *UpdateProfilesWarehouseForSpaceWarehouseAlphaInput) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetSettings returns the Settings field value
func (o *UpdateProfilesWarehouseForSpaceWarehouseAlphaInput) GetSettings() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value
// and a boolean to check if the value has been set.
func (o *UpdateProfilesWarehouseForSpaceWarehouseAlphaInput) GetSettingsOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Settings, true
}

// SetSettings sets field value
func (o *UpdateProfilesWarehouseForSpaceWarehouseAlphaInput) SetSettings(v map[string]interface{}) {
	o.Settings = v
}

// GetSchemaName returns the SchemaName field value if set, zero value otherwise.
func (o *UpdateProfilesWarehouseForSpaceWarehouseAlphaInput) GetSchemaName() string {
	if o == nil || IsNil(o.SchemaName) {
		var ret string
		return ret
	}
	return *o.SchemaName
}

// GetSchemaNameOk returns a tuple with the SchemaName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProfilesWarehouseForSpaceWarehouseAlphaInput) GetSchemaNameOk() (*string, bool) {
	if o == nil || IsNil(o.SchemaName) {
		return nil, false
	}
	return o.SchemaName, true
}

// HasSchemaName returns a boolean if a field has been set.
func (o *UpdateProfilesWarehouseForSpaceWarehouseAlphaInput) HasSchemaName() bool {
	if o != nil && !IsNil(o.SchemaName) {
		return true
	}

	return false
}

// SetSchemaName gets a reference to the given string and assigns it to the SchemaName field.
func (o *UpdateProfilesWarehouseForSpaceWarehouseAlphaInput) SetSchemaName(v string) {
	o.SchemaName = &v
}

func (o UpdateProfilesWarehouseForSpaceWarehouseAlphaInput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateProfilesWarehouseForSpaceWarehouseAlphaInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	toSerialize["settings"] = o.Settings
	if !IsNil(o.SchemaName) {
		toSerialize["schemaName"] = o.SchemaName
	}
	return toSerialize, nil
}

type NullableUpdateProfilesWarehouseForSpaceWarehouseAlphaInput struct {
	value *UpdateProfilesWarehouseForSpaceWarehouseAlphaInput
	isSet bool
}

func (v NullableUpdateProfilesWarehouseForSpaceWarehouseAlphaInput) Get() *UpdateProfilesWarehouseForSpaceWarehouseAlphaInput {
	return v.value
}

func (v *NullableUpdateProfilesWarehouseForSpaceWarehouseAlphaInput) Set(
	val *UpdateProfilesWarehouseForSpaceWarehouseAlphaInput,
) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateProfilesWarehouseForSpaceWarehouseAlphaInput) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateProfilesWarehouseForSpaceWarehouseAlphaInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateProfilesWarehouseForSpaceWarehouseAlphaInput(
	val *UpdateProfilesWarehouseForSpaceWarehouseAlphaInput,
) *NullableUpdateProfilesWarehouseForSpaceWarehouseAlphaInput {
	return &NullableUpdateProfilesWarehouseForSpaceWarehouseAlphaInput{value: val, isSet: true}
}

func (v NullableUpdateProfilesWarehouseForSpaceWarehouseAlphaInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateProfilesWarehouseForSpaceWarehouseAlphaInput) UnmarshalJSON(
	src []byte,
) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
