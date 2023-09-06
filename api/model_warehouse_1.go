/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 36.2.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// Warehouse1 The newly created Warehouse.
type Warehouse1 struct {
	// The id of the Warehouse.
	Id       string    `json:"id"`
	Metadata Metadata1 `json:"metadata"`
	// The id of the Workspace that owns this Warehouse.
	WorkspaceId string `json:"workspaceId"`
	// When set to true, this Warehouse receives data.
	Enabled bool `json:"enabled"`
	// The settings associated with this Warehouse.  Common settings are connection-related configuration used to connect to it, for example host, username, and port.
	Settings NullableModelMap `json:"settings"`
}

// NewWarehouse1 instantiates a new Warehouse1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWarehouse1(
	id string,
	metadata Metadata1,
	workspaceId string,
	enabled bool,
	settings NullableModelMap,
) *Warehouse1 {
	this := Warehouse1{}
	this.Id = id
	this.Metadata = metadata
	this.WorkspaceId = workspaceId
	this.Enabled = enabled
	this.Settings = settings
	return &this
}

// NewWarehouse1WithDefaults instantiates a new Warehouse1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWarehouse1WithDefaults() *Warehouse1 {
	this := Warehouse1{}
	return &this
}

// GetId returns the Id field value
func (o *Warehouse1) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Warehouse1) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Warehouse1) SetId(v string) {
	o.Id = v
}

// GetMetadata returns the Metadata field value
func (o *Warehouse1) GetMetadata() Metadata1 {
	if o == nil {
		var ret Metadata1
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *Warehouse1) GetMetadataOk() (*Metadata1, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *Warehouse1) SetMetadata(v Metadata1) {
	o.Metadata = v
}

// GetWorkspaceId returns the WorkspaceId field value
func (o *Warehouse1) GetWorkspaceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WorkspaceId
}

// GetWorkspaceIdOk returns a tuple with the WorkspaceId field value
// and a boolean to check if the value has been set.
func (o *Warehouse1) GetWorkspaceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WorkspaceId, true
}

// SetWorkspaceId sets field value
func (o *Warehouse1) SetWorkspaceId(v string) {
	o.WorkspaceId = v
}

// GetEnabled returns the Enabled field value
func (o *Warehouse1) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *Warehouse1) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *Warehouse1) SetEnabled(v bool) {
	o.Enabled = v
}

// GetSettings returns the Settings field value
// If the value is explicit nil, the zero value for ModelMap will be returned
func (o *Warehouse1) GetSettings() ModelMap {
	if o == nil || o.Settings.Get() == nil {
		var ret ModelMap
		return ret
	}

	return *o.Settings.Get()
}

// GetSettingsOk returns a tuple with the Settings field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Warehouse1) GetSettingsOk() (*ModelMap, bool) {
	if o == nil {
		return nil, false
	}
	return o.Settings.Get(), o.Settings.IsSet()
}

// SetSettings sets field value
func (o *Warehouse1) SetSettings(v ModelMap) {
	o.Settings.Set(&v)
}

func (o Warehouse1) MarshalJSON() ([]byte, error) {
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

type NullableWarehouse1 struct {
	value *Warehouse1
	isSet bool
}

func (v NullableWarehouse1) Get() *Warehouse1 {
	return v.value
}

func (v *NullableWarehouse1) Set(val *Warehouse1) {
	v.value = val
	v.isSet = true
}

func (v NullableWarehouse1) IsSet() bool {
	return v.isSet
}

func (v *NullableWarehouse1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWarehouse1(val *Warehouse1) *NullableWarehouse1 {
	return &NullableWarehouse1{value: val, isSet: true}
}

func (v NullableWarehouse1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWarehouse1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
