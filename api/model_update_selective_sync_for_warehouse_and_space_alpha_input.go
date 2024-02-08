/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 42.0.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the UpdateSelectiveSyncForWarehouseAndSpaceAlphaInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateSelectiveSyncForWarehouseAndSpaceAlphaInput{}

// UpdateSelectiveSyncForWarehouseAndSpaceAlphaInput Updates the schema for a Space Warehouse connection.
type UpdateSelectiveSyncForWarehouseAndSpaceAlphaInput struct {
	// A list of sync Schema overrides to apply to this Space Warehouse. Note: Selective Sync is not supported if the enableEventTables flag is false.
	SyncOverrides []SpaceWarehouseSchemaOverride `json:"syncOverrides,omitempty"`
	// A flag to enable or disable all event Tables. This field is optional.
	EnableEventTables *bool `json:"enableEventTables,omitempty"`
}

// NewUpdateSelectiveSyncForWarehouseAndSpaceAlphaInput instantiates a new UpdateSelectiveSyncForWarehouseAndSpaceAlphaInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateSelectiveSyncForWarehouseAndSpaceAlphaInput() *UpdateSelectiveSyncForWarehouseAndSpaceAlphaInput {
	this := UpdateSelectiveSyncForWarehouseAndSpaceAlphaInput{}
	return &this
}

// NewUpdateSelectiveSyncForWarehouseAndSpaceAlphaInputWithDefaults instantiates a new UpdateSelectiveSyncForWarehouseAndSpaceAlphaInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateSelectiveSyncForWarehouseAndSpaceAlphaInputWithDefaults() *UpdateSelectiveSyncForWarehouseAndSpaceAlphaInput {
	this := UpdateSelectiveSyncForWarehouseAndSpaceAlphaInput{}
	return &this
}

// GetSyncOverrides returns the SyncOverrides field value if set, zero value otherwise.
func (o *UpdateSelectiveSyncForWarehouseAndSpaceAlphaInput) GetSyncOverrides() []SpaceWarehouseSchemaOverride {
	if o == nil || IsNil(o.SyncOverrides) {
		var ret []SpaceWarehouseSchemaOverride
		return ret
	}
	return o.SyncOverrides
}

// GetSyncOverridesOk returns a tuple with the SyncOverrides field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSelectiveSyncForWarehouseAndSpaceAlphaInput) GetSyncOverridesOk() ([]SpaceWarehouseSchemaOverride, bool) {
	if o == nil || IsNil(o.SyncOverrides) {
		return nil, false
	}
	return o.SyncOverrides, true
}

// HasSyncOverrides returns a boolean if a field has been set.
func (o *UpdateSelectiveSyncForWarehouseAndSpaceAlphaInput) HasSyncOverrides() bool {
	if o != nil && !IsNil(o.SyncOverrides) {
		return true
	}

	return false
}

// SetSyncOverrides gets a reference to the given []SpaceWarehouseSchemaOverride and assigns it to the SyncOverrides field.
func (o *UpdateSelectiveSyncForWarehouseAndSpaceAlphaInput) SetSyncOverrides(
	v []SpaceWarehouseSchemaOverride,
) {
	o.SyncOverrides = v
}

// GetEnableEventTables returns the EnableEventTables field value if set, zero value otherwise.
func (o *UpdateSelectiveSyncForWarehouseAndSpaceAlphaInput) GetEnableEventTables() bool {
	if o == nil || IsNil(o.EnableEventTables) {
		var ret bool
		return ret
	}
	return *o.EnableEventTables
}

// GetEnableEventTablesOk returns a tuple with the EnableEventTables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSelectiveSyncForWarehouseAndSpaceAlphaInput) GetEnableEventTablesOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableEventTables) {
		return nil, false
	}
	return o.EnableEventTables, true
}

// HasEnableEventTables returns a boolean if a field has been set.
func (o *UpdateSelectiveSyncForWarehouseAndSpaceAlphaInput) HasEnableEventTables() bool {
	if o != nil && !IsNil(o.EnableEventTables) {
		return true
	}

	return false
}

// SetEnableEventTables gets a reference to the given bool and assigns it to the EnableEventTables field.
func (o *UpdateSelectiveSyncForWarehouseAndSpaceAlphaInput) SetEnableEventTables(v bool) {
	o.EnableEventTables = &v
}

func (o UpdateSelectiveSyncForWarehouseAndSpaceAlphaInput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateSelectiveSyncForWarehouseAndSpaceAlphaInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SyncOverrides) {
		toSerialize["syncOverrides"] = o.SyncOverrides
	}
	if !IsNil(o.EnableEventTables) {
		toSerialize["enableEventTables"] = o.EnableEventTables
	}
	return toSerialize, nil
}

type NullableUpdateSelectiveSyncForWarehouseAndSpaceAlphaInput struct {
	value *UpdateSelectiveSyncForWarehouseAndSpaceAlphaInput
	isSet bool
}

func (v NullableUpdateSelectiveSyncForWarehouseAndSpaceAlphaInput) Get() *UpdateSelectiveSyncForWarehouseAndSpaceAlphaInput {
	return v.value
}

func (v *NullableUpdateSelectiveSyncForWarehouseAndSpaceAlphaInput) Set(
	val *UpdateSelectiveSyncForWarehouseAndSpaceAlphaInput,
) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateSelectiveSyncForWarehouseAndSpaceAlphaInput) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateSelectiveSyncForWarehouseAndSpaceAlphaInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateSelectiveSyncForWarehouseAndSpaceAlphaInput(
	val *UpdateSelectiveSyncForWarehouseAndSpaceAlphaInput,
) *NullableUpdateSelectiveSyncForWarehouseAndSpaceAlphaInput {
	return &NullableUpdateSelectiveSyncForWarehouseAndSpaceAlphaInput{value: val, isSet: true}
}

func (v NullableUpdateSelectiveSyncForWarehouseAndSpaceAlphaInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateSelectiveSyncForWarehouseAndSpaceAlphaInput) UnmarshalJSON(
	src []byte,
) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
