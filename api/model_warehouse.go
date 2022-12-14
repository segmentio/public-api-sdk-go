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

// Warehouse The returned Warehouse object.
type Warehouse struct {
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

// NewWarehouse instantiates a new Warehouse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWarehouse(
	id string,
	metadata Metadata2,
	workspaceId string,
	enabled bool,
	settings NullableModelMap,
) *Warehouse {
	this := Warehouse{}
	this.Id = id
	this.Metadata = metadata
	this.WorkspaceId = workspaceId
	this.Enabled = enabled
	this.Settings = settings
	return &this
}

// NewWarehouseWithDefaults instantiates a new Warehouse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWarehouseWithDefaults() *Warehouse {
	this := Warehouse{}
	return &this
}

// GetId returns the Id field value
func (o *Warehouse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Warehouse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Warehouse) SetId(v string) {
	o.Id = v
}

// GetMetadata returns the Metadata field value
func (o *Warehouse) GetMetadata() Metadata2 {
	if o == nil {
		var ret Metadata2
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *Warehouse) GetMetadataOk() (*Metadata2, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *Warehouse) SetMetadata(v Metadata2) {
	o.Metadata = v
}

// GetWorkspaceId returns the WorkspaceId field value
func (o *Warehouse) GetWorkspaceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WorkspaceId
}

// GetWorkspaceIdOk returns a tuple with the WorkspaceId field value
// and a boolean to check if the value has been set.
func (o *Warehouse) GetWorkspaceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WorkspaceId, true
}

// SetWorkspaceId sets field value
func (o *Warehouse) SetWorkspaceId(v string) {
	o.WorkspaceId = v
}

// GetEnabled returns the Enabled field value
func (o *Warehouse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *Warehouse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *Warehouse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetSettings returns the Settings field value
// If the value is explicit nil, the zero value for ModelMap will be returned
func (o *Warehouse) GetSettings() ModelMap {
	if o == nil || o.Settings.Get() == nil {
		var ret ModelMap
		return ret
	}

	return *o.Settings.Get()
}

// GetSettingsOk returns a tuple with the Settings field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Warehouse) GetSettingsOk() (*ModelMap, bool) {
	if o == nil {
		return nil, false
	}
	return o.Settings.Get(), o.Settings.IsSet()
}

// SetSettings sets field value
func (o *Warehouse) SetSettings(v ModelMap) {
	o.Settings.Set(&v)
}

func (o Warehouse) MarshalJSON() ([]byte, error) {
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

type NullableWarehouse struct {
	value *Warehouse
	isSet bool
}

func (v NullableWarehouse) Get() *Warehouse {
	return v.value
}

func (v *NullableWarehouse) Set(val *Warehouse) {
	v.value = val
	v.isSet = true
}

func (v NullableWarehouse) IsSet() bool {
	return v.isSet
}

func (v *NullableWarehouse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWarehouse(val *Warehouse) *NullableWarehouse {
	return &NullableWarehouse{value: val, isSet: true}
}

func (v NullableWarehouse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWarehouse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
