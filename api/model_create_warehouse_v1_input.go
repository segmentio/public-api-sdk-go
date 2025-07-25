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

// checks if the CreateWarehouseV1Input type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateWarehouseV1Input{}

// CreateWarehouseV1Input Create a new Warehouse based on a set of parameters.
type CreateWarehouseV1Input struct {
	// The Warehouse metadata to use.
	MetadataId string `json:"metadataId"`
	// An optional human-readable name for this Warehouse.
	Name *string `json:"name,omitempty"`
	// Enable to allow this Warehouse to receive data. Defaults to true.
	Enabled *bool `json:"enabled,omitempty"`
	// A key-value object that contains instance-specific Warehouse settings.
	Settings map[string]interface{} `json:"settings"`
	// Whether to disconnect all Sources from this Warehouse.
	DisconnectAllSources *bool `json:"disconnectAllSources,omitempty"`
}

// NewCreateWarehouseV1Input instantiates a new CreateWarehouseV1Input object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateWarehouseV1Input(
	metadataId string,
	settings map[string]interface{},
) *CreateWarehouseV1Input {
	this := CreateWarehouseV1Input{}
	this.MetadataId = metadataId
	this.Settings = settings
	return &this
}

// NewCreateWarehouseV1InputWithDefaults instantiates a new CreateWarehouseV1Input object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateWarehouseV1InputWithDefaults() *CreateWarehouseV1Input {
	this := CreateWarehouseV1Input{}
	return &this
}

// GetMetadataId returns the MetadataId field value
func (o *CreateWarehouseV1Input) GetMetadataId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MetadataId
}

// GetMetadataIdOk returns a tuple with the MetadataId field value
// and a boolean to check if the value has been set.
func (o *CreateWarehouseV1Input) GetMetadataIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetadataId, true
}

// SetMetadataId sets field value
func (o *CreateWarehouseV1Input) SetMetadataId(v string) {
	o.MetadataId = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateWarehouseV1Input) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateWarehouseV1Input) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateWarehouseV1Input) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreateWarehouseV1Input) SetName(v string) {
	o.Name = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *CreateWarehouseV1Input) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateWarehouseV1Input) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *CreateWarehouseV1Input) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *CreateWarehouseV1Input) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetSettings returns the Settings field value
func (o *CreateWarehouseV1Input) GetSettings() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value
// and a boolean to check if the value has been set.
func (o *CreateWarehouseV1Input) GetSettingsOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Settings, true
}

// SetSettings sets field value
func (o *CreateWarehouseV1Input) SetSettings(v map[string]interface{}) {
	o.Settings = v
}

// GetDisconnectAllSources returns the DisconnectAllSources field value if set, zero value otherwise.
func (o *CreateWarehouseV1Input) GetDisconnectAllSources() bool {
	if o == nil || IsNil(o.DisconnectAllSources) {
		var ret bool
		return ret
	}
	return *o.DisconnectAllSources
}

// GetDisconnectAllSourcesOk returns a tuple with the DisconnectAllSources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateWarehouseV1Input) GetDisconnectAllSourcesOk() (*bool, bool) {
	if o == nil || IsNil(o.DisconnectAllSources) {
		return nil, false
	}
	return o.DisconnectAllSources, true
}

// HasDisconnectAllSources returns a boolean if a field has been set.
func (o *CreateWarehouseV1Input) HasDisconnectAllSources() bool {
	if o != nil && !IsNil(o.DisconnectAllSources) {
		return true
	}

	return false
}

// SetDisconnectAllSources gets a reference to the given bool and assigns it to the DisconnectAllSources field.
func (o *CreateWarehouseV1Input) SetDisconnectAllSources(v bool) {
	o.DisconnectAllSources = &v
}

func (o CreateWarehouseV1Input) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateWarehouseV1Input) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["metadataId"] = o.MetadataId
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	toSerialize["settings"] = o.Settings
	if !IsNil(o.DisconnectAllSources) {
		toSerialize["disconnectAllSources"] = o.DisconnectAllSources
	}
	return toSerialize, nil
}

type NullableCreateWarehouseV1Input struct {
	value *CreateWarehouseV1Input
	isSet bool
}

func (v NullableCreateWarehouseV1Input) Get() *CreateWarehouseV1Input {
	return v.value
}

func (v *NullableCreateWarehouseV1Input) Set(val *CreateWarehouseV1Input) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateWarehouseV1Input) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateWarehouseV1Input) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateWarehouseV1Input(
	val *CreateWarehouseV1Input,
) *NullableCreateWarehouseV1Input {
	return &NullableCreateWarehouseV1Input{value: val, isSet: true}
}

func (v NullableCreateWarehouseV1Input) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateWarehouseV1Input) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
