/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 58.0.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the WarehouseV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WarehouseV1{}

// WarehouseV1 Defines a data Warehouse used as a Destination for Segment data.
type WarehouseV1 struct {
	// The id of the Warehouse.
	Id       string              `json:"id"`
	Metadata WarehouseMetadataV1 `json:"metadata"`
	// The id of the Workspace that owns this Warehouse.
	WorkspaceId string `json:"workspaceId"`
	// When set to true, this Warehouse receives data.
	Enabled bool `json:"enabled"`
	// A key-value object that contains instance-specific Warehouse settings.
	Settings map[string]interface{} `json:"settings"`
}

// NewWarehouseV1 instantiates a new WarehouseV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWarehouseV1(
	id string,
	metadata WarehouseMetadataV1,
	workspaceId string,
	enabled bool,
	settings map[string]interface{},
) *WarehouseV1 {
	this := WarehouseV1{}
	this.Id = id
	this.Metadata = metadata
	this.WorkspaceId = workspaceId
	this.Enabled = enabled
	this.Settings = settings
	return &this
}

// NewWarehouseV1WithDefaults instantiates a new WarehouseV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWarehouseV1WithDefaults() *WarehouseV1 {
	this := WarehouseV1{}
	return &this
}

// GetId returns the Id field value
func (o *WarehouseV1) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *WarehouseV1) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *WarehouseV1) SetId(v string) {
	o.Id = v
}

// GetMetadata returns the Metadata field value
func (o *WarehouseV1) GetMetadata() WarehouseMetadataV1 {
	if o == nil {
		var ret WarehouseMetadataV1
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *WarehouseV1) GetMetadataOk() (*WarehouseMetadataV1, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *WarehouseV1) SetMetadata(v WarehouseMetadataV1) {
	o.Metadata = v
}

// GetWorkspaceId returns the WorkspaceId field value
func (o *WarehouseV1) GetWorkspaceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WorkspaceId
}

// GetWorkspaceIdOk returns a tuple with the WorkspaceId field value
// and a boolean to check if the value has been set.
func (o *WarehouseV1) GetWorkspaceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WorkspaceId, true
}

// SetWorkspaceId sets field value
func (o *WarehouseV1) SetWorkspaceId(v string) {
	o.WorkspaceId = v
}

// GetEnabled returns the Enabled field value
func (o *WarehouseV1) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *WarehouseV1) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *WarehouseV1) SetEnabled(v bool) {
	o.Enabled = v
}

// GetSettings returns the Settings field value
func (o *WarehouseV1) GetSettings() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value
// and a boolean to check if the value has been set.
func (o *WarehouseV1) GetSettingsOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Settings, true
}

// SetSettings sets field value
func (o *WarehouseV1) SetSettings(v map[string]interface{}) {
	o.Settings = v
}

func (o WarehouseV1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WarehouseV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["metadata"] = o.Metadata
	toSerialize["workspaceId"] = o.WorkspaceId
	toSerialize["enabled"] = o.Enabled
	toSerialize["settings"] = o.Settings
	return toSerialize, nil
}

type NullableWarehouseV1 struct {
	value *WarehouseV1
	isSet bool
}

func (v NullableWarehouseV1) Get() *WarehouseV1 {
	return v.value
}

func (v *NullableWarehouseV1) Set(val *WarehouseV1) {
	v.value = val
	v.isSet = true
}

func (v NullableWarehouseV1) IsSet() bool {
	return v.isSet
}

func (v *NullableWarehouseV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWarehouseV1(val *WarehouseV1) *NullableWarehouseV1 {
	return &NullableWarehouseV1{value: val, isSet: true}
}

func (v NullableWarehouseV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWarehouseV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
