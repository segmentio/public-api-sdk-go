/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 33.0.3
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// WarehouseV1 Defines a data Warehouse used as a Destination for Segment data.
type WarehouseV1 struct {
	// The id of the Warehouse.
	Id       string    `json:"id"`
	Metadata Metadata2 `json:"metadata"`
	// The id of the Workspace that owns this Warehouse.
	WorkspaceId string `json:"workspaceId"`
	// When set to true, this Warehouse receives data.
	Enabled bool `json:"enabled"`
	// The settings associated with this Warehouse.  Common settings are connection-related configuration used to connect to it, for example host, username, and port.
	Settings NullableModelMap `json:"settings"`
}

// NewWarehouseV1 instantiates a new WarehouseV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWarehouseV1(
	id string,
	metadata Metadata2,
	workspaceId string,
	enabled bool,
	settings NullableModelMap,
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
func (o *WarehouseV1) GetMetadata() Metadata2 {
	if o == nil {
		var ret Metadata2
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *WarehouseV1) GetMetadataOk() (*Metadata2, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *WarehouseV1) SetMetadata(v Metadata2) {
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
// If the value is explicit nil, the zero value for ModelMap will be returned
func (o *WarehouseV1) GetSettings() ModelMap {
	if o == nil || o.Settings.Get() == nil {
		var ret ModelMap
		return ret
	}

	return *o.Settings.Get()
}

// GetSettingsOk returns a tuple with the Settings field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WarehouseV1) GetSettingsOk() (*ModelMap, bool) {
	if o == nil {
		return nil, false
	}
	return o.Settings.Get(), o.Settings.IsSet()
}

// SetSettings sets field value
func (o *WarehouseV1) SetSettings(v ModelMap) {
	o.Settings.Set(&v)
}

func (o WarehouseV1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
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
		toSerialize["settings"] = o.Settings.Get()
	}
	return json.Marshal(toSerialize)
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
