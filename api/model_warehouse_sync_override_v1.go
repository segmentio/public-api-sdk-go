/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 34.3.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// WarehouseSyncOverrideV1 Represents the override for a Source/collection/property? path to apply to a Warehouse.
type WarehouseSyncOverrideV1 struct {
	// The id of the Source this schema item applies to.
	SourceId string `json:"sourceId"`
	// The collection the schema item override applies to.
	Collection *string `json:"collection,omitempty"`
	// Enable to apply the override.
	Enabled bool `json:"enabled"`
	// A property within the collection to which the override applies.
	Property *string `json:"property,omitempty"`
}

// NewWarehouseSyncOverrideV1 instantiates a new WarehouseSyncOverrideV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWarehouseSyncOverrideV1(sourceId string, enabled bool) *WarehouseSyncOverrideV1 {
	this := WarehouseSyncOverrideV1{}
	this.SourceId = sourceId
	this.Enabled = enabled
	return &this
}

// NewWarehouseSyncOverrideV1WithDefaults instantiates a new WarehouseSyncOverrideV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWarehouseSyncOverrideV1WithDefaults() *WarehouseSyncOverrideV1 {
	this := WarehouseSyncOverrideV1{}
	return &this
}

// GetSourceId returns the SourceId field value
func (o *WarehouseSyncOverrideV1) GetSourceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceId
}

// GetSourceIdOk returns a tuple with the SourceId field value
// and a boolean to check if the value has been set.
func (o *WarehouseSyncOverrideV1) GetSourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceId, true
}

// SetSourceId sets field value
func (o *WarehouseSyncOverrideV1) SetSourceId(v string) {
	o.SourceId = v
}

// GetCollection returns the Collection field value if set, zero value otherwise.
func (o *WarehouseSyncOverrideV1) GetCollection() string {
	if o == nil || o.Collection == nil {
		var ret string
		return ret
	}
	return *o.Collection
}

// GetCollectionOk returns a tuple with the Collection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WarehouseSyncOverrideV1) GetCollectionOk() (*string, bool) {
	if o == nil || o.Collection == nil {
		return nil, false
	}
	return o.Collection, true
}

// HasCollection returns a boolean if a field has been set.
func (o *WarehouseSyncOverrideV1) HasCollection() bool {
	if o != nil && o.Collection != nil {
		return true
	}

	return false
}

// SetCollection gets a reference to the given string and assigns it to the Collection field.
func (o *WarehouseSyncOverrideV1) SetCollection(v string) {
	o.Collection = &v
}

// GetEnabled returns the Enabled field value
func (o *WarehouseSyncOverrideV1) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *WarehouseSyncOverrideV1) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *WarehouseSyncOverrideV1) SetEnabled(v bool) {
	o.Enabled = v
}

// GetProperty returns the Property field value if set, zero value otherwise.
func (o *WarehouseSyncOverrideV1) GetProperty() string {
	if o == nil || o.Property == nil {
		var ret string
		return ret
	}
	return *o.Property
}

// GetPropertyOk returns a tuple with the Property field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WarehouseSyncOverrideV1) GetPropertyOk() (*string, bool) {
	if o == nil || o.Property == nil {
		return nil, false
	}
	return o.Property, true
}

// HasProperty returns a boolean if a field has been set.
func (o *WarehouseSyncOverrideV1) HasProperty() bool {
	if o != nil && o.Property != nil {
		return true
	}

	return false
}

// SetProperty gets a reference to the given string and assigns it to the Property field.
func (o *WarehouseSyncOverrideV1) SetProperty(v string) {
	o.Property = &v
}

func (o WarehouseSyncOverrideV1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["sourceId"] = o.SourceId
	}
	if o.Collection != nil {
		toSerialize["collection"] = o.Collection
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Property != nil {
		toSerialize["property"] = o.Property
	}
	return json.Marshal(toSerialize)
}

type NullableWarehouseSyncOverrideV1 struct {
	value *WarehouseSyncOverrideV1
	isSet bool
}

func (v NullableWarehouseSyncOverrideV1) Get() *WarehouseSyncOverrideV1 {
	return v.value
}

func (v *NullableWarehouseSyncOverrideV1) Set(val *WarehouseSyncOverrideV1) {
	v.value = val
	v.isSet = true
}

func (v NullableWarehouseSyncOverrideV1) IsSet() bool {
	return v.isSet
}

func (v *NullableWarehouseSyncOverrideV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWarehouseSyncOverrideV1(
	val *WarehouseSyncOverrideV1,
) *NullableWarehouseSyncOverrideV1 {
	return &NullableWarehouseSyncOverrideV1{value: val, isSet: true}
}

func (v NullableWarehouseSyncOverrideV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWarehouseSyncOverrideV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
