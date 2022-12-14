/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 33.0.4
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// CreateWarehouseV1Input Create a new Warehouse based on a set of parameters.
type CreateWarehouseV1Input struct {
	// The Warehouse metadata to use.
	MetadataId string `json:"metadataId"`
	// An optional human-readable name for this Warehouse.
	Name *string `json:"name,omitempty"`
	// Enable to allow this Warehouse to receive data. Defaults to true.
	Enabled *bool `json:"enabled,omitempty"`
	// A key-value object that contains instance-specific settings for a Warehouse.  Different kinds of Warehouses require different settings. The required and optional settings for a Warehouse are described in the `options` object of the associated Warehouse metadata.  You can find the full list of Warehouse metadata and related settings information in the `/catalog/warehouses` endpoint.
	Settings NullableModelMap `json:"settings"`
}

// NewCreateWarehouseV1Input instantiates a new CreateWarehouseV1Input object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateWarehouseV1Input(
	metadataId string,
	settings NullableModelMap,
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
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateWarehouseV1Input) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateWarehouseV1Input) HasName() bool {
	if o != nil && o.Name != nil {
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
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateWarehouseV1Input) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *CreateWarehouseV1Input) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *CreateWarehouseV1Input) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetSettings returns the Settings field value
// If the value is explicit nil, the zero value for ModelMap will be returned
func (o *CreateWarehouseV1Input) GetSettings() ModelMap {
	if o == nil || o.Settings.Get() == nil {
		var ret ModelMap
		return ret
	}

	return *o.Settings.Get()
}

// GetSettingsOk returns a tuple with the Settings field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateWarehouseV1Input) GetSettingsOk() (*ModelMap, bool) {
	if o == nil {
		return nil, false
	}
	return o.Settings.Get(), o.Settings.IsSet()
}

// SetSettings sets field value
func (o *CreateWarehouseV1Input) SetSettings(v ModelMap) {
	o.Settings.Set(&v)
}

func (o CreateWarehouseV1Input) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["metadataId"] = o.MetadataId
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if true {
		toSerialize["settings"] = o.Settings.Get()
	}
	return json.Marshal(toSerialize)
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
