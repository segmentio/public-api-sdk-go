/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 40.0.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the WarehouseSelectiveSyncItemV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WarehouseSelectiveSyncItemV1{}

// WarehouseSelectiveSyncItemV1 Represents an entry in the Warehouse Selective Sync schema for a Warehouse and Source pair.
type WarehouseSelectiveSyncItemV1 struct {
	// The Source id attached to this Warehouse.
	SourceId string `json:"sourceId"`
	// The collection within the Source.
	Collection string `json:"collection"`
	// The id of the Warehouse this sync belongs to.
	WarehouseId string `json:"warehouseId"`
	// Whether this Selective Sync item is enabled.
	Enabled bool `json:"enabled"`
	// A map that contains the properties within the collection to which the Warehouse should sync.
	Properties map[string]interface{} `json:"properties"`
}

// NewWarehouseSelectiveSyncItemV1 instantiates a new WarehouseSelectiveSyncItemV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWarehouseSelectiveSyncItemV1(
	sourceId string,
	collection string,
	warehouseId string,
	enabled bool,
	properties map[string]interface{},
) *WarehouseSelectiveSyncItemV1 {
	this := WarehouseSelectiveSyncItemV1{}
	this.SourceId = sourceId
	this.Collection = collection
	this.WarehouseId = warehouseId
	this.Enabled = enabled
	this.Properties = properties
	return &this
}

// NewWarehouseSelectiveSyncItemV1WithDefaults instantiates a new WarehouseSelectiveSyncItemV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWarehouseSelectiveSyncItemV1WithDefaults() *WarehouseSelectiveSyncItemV1 {
	this := WarehouseSelectiveSyncItemV1{}
	return &this
}

// GetSourceId returns the SourceId field value
func (o *WarehouseSelectiveSyncItemV1) GetSourceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceId
}

// GetSourceIdOk returns a tuple with the SourceId field value
// and a boolean to check if the value has been set.
func (o *WarehouseSelectiveSyncItemV1) GetSourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceId, true
}

// SetSourceId sets field value
func (o *WarehouseSelectiveSyncItemV1) SetSourceId(v string) {
	o.SourceId = v
}

// GetCollection returns the Collection field value
func (o *WarehouseSelectiveSyncItemV1) GetCollection() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Collection
}

// GetCollectionOk returns a tuple with the Collection field value
// and a boolean to check if the value has been set.
func (o *WarehouseSelectiveSyncItemV1) GetCollectionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Collection, true
}

// SetCollection sets field value
func (o *WarehouseSelectiveSyncItemV1) SetCollection(v string) {
	o.Collection = v
}

// GetWarehouseId returns the WarehouseId field value
func (o *WarehouseSelectiveSyncItemV1) GetWarehouseId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WarehouseId
}

// GetWarehouseIdOk returns a tuple with the WarehouseId field value
// and a boolean to check if the value has been set.
func (o *WarehouseSelectiveSyncItemV1) GetWarehouseIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WarehouseId, true
}

// SetWarehouseId sets field value
func (o *WarehouseSelectiveSyncItemV1) SetWarehouseId(v string) {
	o.WarehouseId = v
}

// GetEnabled returns the Enabled field value
func (o *WarehouseSelectiveSyncItemV1) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *WarehouseSelectiveSyncItemV1) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *WarehouseSelectiveSyncItemV1) SetEnabled(v bool) {
	o.Enabled = v
}

// GetProperties returns the Properties field value
func (o *WarehouseSelectiveSyncItemV1) GetProperties() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
func (o *WarehouseSelectiveSyncItemV1) GetPropertiesOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Properties, true
}

// SetProperties sets field value
func (o *WarehouseSelectiveSyncItemV1) SetProperties(v map[string]interface{}) {
	o.Properties = v
}

func (o WarehouseSelectiveSyncItemV1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WarehouseSelectiveSyncItemV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sourceId"] = o.SourceId
	toSerialize["collection"] = o.Collection
	toSerialize["warehouseId"] = o.WarehouseId
	toSerialize["enabled"] = o.Enabled
	toSerialize["properties"] = o.Properties
	return toSerialize, nil
}

type NullableWarehouseSelectiveSyncItemV1 struct {
	value *WarehouseSelectiveSyncItemV1
	isSet bool
}

func (v NullableWarehouseSelectiveSyncItemV1) Get() *WarehouseSelectiveSyncItemV1 {
	return v.value
}

func (v *NullableWarehouseSelectiveSyncItemV1) Set(val *WarehouseSelectiveSyncItemV1) {
	v.value = val
	v.isSet = true
}

func (v NullableWarehouseSelectiveSyncItemV1) IsSet() bool {
	return v.isSet
}

func (v *NullableWarehouseSelectiveSyncItemV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWarehouseSelectiveSyncItemV1(
	val *WarehouseSelectiveSyncItemV1,
) *NullableWarehouseSelectiveSyncItemV1 {
	return &NullableWarehouseSelectiveSyncItemV1{value: val, isSet: true}
}

func (v NullableWarehouseSelectiveSyncItemV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWarehouseSelectiveSyncItemV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
