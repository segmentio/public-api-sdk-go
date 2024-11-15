/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 57.0.1
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the CreateProfilesWarehouseAlphaInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateProfilesWarehouseAlphaInput{}

// CreateProfilesWarehouseAlphaInput Create a new Profiles Warehouse based on a set of parameters.
type CreateProfilesWarehouseAlphaInput struct {
	// The Warehouse metadata to use.
	MetadataId string `json:"metadataId"`
	// An optional human-readable name for this Warehouse.
	Name *string `json:"name,omitempty"`
	// Enable to allow this Warehouse to receive data. Defaults to true.
	Enabled *bool `json:"enabled,omitempty"`
	// A key-value object that contains instance-specific Warehouse settings.
	Settings map[string]interface{} `json:"settings"`
	// The custom schema name that Segment uses on the Warehouse side. The space slug value is default otherwise. The schema name cannot be an existing schema name in the Warehouse. To use an existing schema name, please create a profiles Warehouse through the Segment app UI.
	SchemaName *string `json:"schemaName,omitempty"`
}

// NewCreateProfilesWarehouseAlphaInput instantiates a new CreateProfilesWarehouseAlphaInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateProfilesWarehouseAlphaInput(
	metadataId string,
	settings map[string]interface{},
) *CreateProfilesWarehouseAlphaInput {
	this := CreateProfilesWarehouseAlphaInput{}
	this.MetadataId = metadataId
	this.Settings = settings
	return &this
}

// NewCreateProfilesWarehouseAlphaInputWithDefaults instantiates a new CreateProfilesWarehouseAlphaInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateProfilesWarehouseAlphaInputWithDefaults() *CreateProfilesWarehouseAlphaInput {
	this := CreateProfilesWarehouseAlphaInput{}
	return &this
}

// GetMetadataId returns the MetadataId field value
func (o *CreateProfilesWarehouseAlphaInput) GetMetadataId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MetadataId
}

// GetMetadataIdOk returns a tuple with the MetadataId field value
// and a boolean to check if the value has been set.
func (o *CreateProfilesWarehouseAlphaInput) GetMetadataIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetadataId, true
}

// SetMetadataId sets field value
func (o *CreateProfilesWarehouseAlphaInput) SetMetadataId(v string) {
	o.MetadataId = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateProfilesWarehouseAlphaInput) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateProfilesWarehouseAlphaInput) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateProfilesWarehouseAlphaInput) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreateProfilesWarehouseAlphaInput) SetName(v string) {
	o.Name = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *CreateProfilesWarehouseAlphaInput) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateProfilesWarehouseAlphaInput) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *CreateProfilesWarehouseAlphaInput) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *CreateProfilesWarehouseAlphaInput) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetSettings returns the Settings field value
func (o *CreateProfilesWarehouseAlphaInput) GetSettings() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value
// and a boolean to check if the value has been set.
func (o *CreateProfilesWarehouseAlphaInput) GetSettingsOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Settings, true
}

// SetSettings sets field value
func (o *CreateProfilesWarehouseAlphaInput) SetSettings(v map[string]interface{}) {
	o.Settings = v
}

// GetSchemaName returns the SchemaName field value if set, zero value otherwise.
func (o *CreateProfilesWarehouseAlphaInput) GetSchemaName() string {
	if o == nil || IsNil(o.SchemaName) {
		var ret string
		return ret
	}
	return *o.SchemaName
}

// GetSchemaNameOk returns a tuple with the SchemaName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateProfilesWarehouseAlphaInput) GetSchemaNameOk() (*string, bool) {
	if o == nil || IsNil(o.SchemaName) {
		return nil, false
	}
	return o.SchemaName, true
}

// HasSchemaName returns a boolean if a field has been set.
func (o *CreateProfilesWarehouseAlphaInput) HasSchemaName() bool {
	if o != nil && !IsNil(o.SchemaName) {
		return true
	}

	return false
}

// SetSchemaName gets a reference to the given string and assigns it to the SchemaName field.
func (o *CreateProfilesWarehouseAlphaInput) SetSchemaName(v string) {
	o.SchemaName = &v
}

func (o CreateProfilesWarehouseAlphaInput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateProfilesWarehouseAlphaInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["metadataId"] = o.MetadataId
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

type NullableCreateProfilesWarehouseAlphaInput struct {
	value *CreateProfilesWarehouseAlphaInput
	isSet bool
}

func (v NullableCreateProfilesWarehouseAlphaInput) Get() *CreateProfilesWarehouseAlphaInput {
	return v.value
}

func (v *NullableCreateProfilesWarehouseAlphaInput) Set(val *CreateProfilesWarehouseAlphaInput) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateProfilesWarehouseAlphaInput) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateProfilesWarehouseAlphaInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateProfilesWarehouseAlphaInput(
	val *CreateProfilesWarehouseAlphaInput,
) *NullableCreateProfilesWarehouseAlphaInput {
	return &NullableCreateProfilesWarehouseAlphaInput{value: val, isSet: true}
}

func (v NullableCreateProfilesWarehouseAlphaInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateProfilesWarehouseAlphaInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
