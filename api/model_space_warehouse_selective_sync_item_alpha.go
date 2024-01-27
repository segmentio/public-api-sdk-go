/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 39.0.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the SpaceWarehouseSelectiveSyncItemAlpha type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SpaceWarehouseSelectiveSyncItemAlpha{}

// SpaceWarehouseSelectiveSyncItemAlpha Represents an entry in the Space Warehouse Selective Sync schema for a Warehouse and Space pair.
type SpaceWarehouseSelectiveSyncItemAlpha struct {
	// The collection within the Source.
	Collection string `json:"collection"`
	// The id of the Warehouse this sync belongs to.
	WarehouseId string `json:"warehouseId"`
	// The Enabled flag ok telling whether the Collection is enabled or not.
	Enabled bool `json:"enabled"`
	// A map that contains the properties within the collection to which the Warehouse should sync.
	Properties map[string]interface{} `json:"properties"`
}

// NewSpaceWarehouseSelectiveSyncItemAlpha instantiates a new SpaceWarehouseSelectiveSyncItemAlpha object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpaceWarehouseSelectiveSyncItemAlpha(
	collection string,
	warehouseId string,
	enabled bool,
	properties map[string]interface{},
) *SpaceWarehouseSelectiveSyncItemAlpha {
	this := SpaceWarehouseSelectiveSyncItemAlpha{}
	this.Collection = collection
	this.WarehouseId = warehouseId
	this.Enabled = enabled
	this.Properties = properties
	return &this
}

// NewSpaceWarehouseSelectiveSyncItemAlphaWithDefaults instantiates a new SpaceWarehouseSelectiveSyncItemAlpha object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpaceWarehouseSelectiveSyncItemAlphaWithDefaults() *SpaceWarehouseSelectiveSyncItemAlpha {
	this := SpaceWarehouseSelectiveSyncItemAlpha{}
	return &this
}

// GetCollection returns the Collection field value
func (o *SpaceWarehouseSelectiveSyncItemAlpha) GetCollection() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Collection
}

// GetCollectionOk returns a tuple with the Collection field value
// and a boolean to check if the value has been set.
func (o *SpaceWarehouseSelectiveSyncItemAlpha) GetCollectionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Collection, true
}

// SetCollection sets field value
func (o *SpaceWarehouseSelectiveSyncItemAlpha) SetCollection(v string) {
	o.Collection = v
}

// GetWarehouseId returns the WarehouseId field value
func (o *SpaceWarehouseSelectiveSyncItemAlpha) GetWarehouseId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WarehouseId
}

// GetWarehouseIdOk returns a tuple with the WarehouseId field value
// and a boolean to check if the value has been set.
func (o *SpaceWarehouseSelectiveSyncItemAlpha) GetWarehouseIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WarehouseId, true
}

// SetWarehouseId sets field value
func (o *SpaceWarehouseSelectiveSyncItemAlpha) SetWarehouseId(v string) {
	o.WarehouseId = v
}

// GetEnabled returns the Enabled field value
func (o *SpaceWarehouseSelectiveSyncItemAlpha) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *SpaceWarehouseSelectiveSyncItemAlpha) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *SpaceWarehouseSelectiveSyncItemAlpha) SetEnabled(v bool) {
	o.Enabled = v
}

// GetProperties returns the Properties field value
func (o *SpaceWarehouseSelectiveSyncItemAlpha) GetProperties() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
func (o *SpaceWarehouseSelectiveSyncItemAlpha) GetPropertiesOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Properties, true
}

// SetProperties sets field value
func (o *SpaceWarehouseSelectiveSyncItemAlpha) SetProperties(v map[string]interface{}) {
	o.Properties = v
}

func (o SpaceWarehouseSelectiveSyncItemAlpha) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SpaceWarehouseSelectiveSyncItemAlpha) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["collection"] = o.Collection
	toSerialize["warehouseId"] = o.WarehouseId
	toSerialize["enabled"] = o.Enabled
	toSerialize["properties"] = o.Properties
	return toSerialize, nil
}

type NullableSpaceWarehouseSelectiveSyncItemAlpha struct {
	value *SpaceWarehouseSelectiveSyncItemAlpha
	isSet bool
}

func (v NullableSpaceWarehouseSelectiveSyncItemAlpha) Get() *SpaceWarehouseSelectiveSyncItemAlpha {
	return v.value
}

func (v *NullableSpaceWarehouseSelectiveSyncItemAlpha) Set(
	val *SpaceWarehouseSelectiveSyncItemAlpha,
) {
	v.value = val
	v.isSet = true
}

func (v NullableSpaceWarehouseSelectiveSyncItemAlpha) IsSet() bool {
	return v.isSet
}

func (v *NullableSpaceWarehouseSelectiveSyncItemAlpha) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpaceWarehouseSelectiveSyncItemAlpha(
	val *SpaceWarehouseSelectiveSyncItemAlpha,
) *NullableSpaceWarehouseSelectiveSyncItemAlpha {
	return &NullableSpaceWarehouseSelectiveSyncItemAlpha{value: val, isSet: true}
}

func (v NullableSpaceWarehouseSelectiveSyncItemAlpha) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpaceWarehouseSelectiveSyncItemAlpha) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
