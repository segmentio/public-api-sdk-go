/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 54.0.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the SpaceWarehouseSchemaOverride type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SpaceWarehouseSchemaOverride{}

// SpaceWarehouseSchemaOverride Overrides the enabled or disabled state of the specified collection and / or properties within the schema.
type SpaceWarehouseSchemaOverride struct {
	// The collection within the Source.
	Collection string `json:"collection"`
	// Represents the overridden enabled state for the listed collection and / or properties.
	Enabled bool `json:"enabled"`
	// A map that contains the properties within the collection to which the Warehouse should sync.
	Property *string `json:"property,omitempty"`
}

// NewSpaceWarehouseSchemaOverride instantiates a new SpaceWarehouseSchemaOverride object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpaceWarehouseSchemaOverride(
	collection string,
	enabled bool,
) *SpaceWarehouseSchemaOverride {
	this := SpaceWarehouseSchemaOverride{}
	this.Collection = collection
	this.Enabled = enabled
	return &this
}

// NewSpaceWarehouseSchemaOverrideWithDefaults instantiates a new SpaceWarehouseSchemaOverride object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpaceWarehouseSchemaOverrideWithDefaults() *SpaceWarehouseSchemaOverride {
	this := SpaceWarehouseSchemaOverride{}
	return &this
}

// GetCollection returns the Collection field value
func (o *SpaceWarehouseSchemaOverride) GetCollection() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Collection
}

// GetCollectionOk returns a tuple with the Collection field value
// and a boolean to check if the value has been set.
func (o *SpaceWarehouseSchemaOverride) GetCollectionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Collection, true
}

// SetCollection sets field value
func (o *SpaceWarehouseSchemaOverride) SetCollection(v string) {
	o.Collection = v
}

// GetEnabled returns the Enabled field value
func (o *SpaceWarehouseSchemaOverride) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *SpaceWarehouseSchemaOverride) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *SpaceWarehouseSchemaOverride) SetEnabled(v bool) {
	o.Enabled = v
}

// GetProperty returns the Property field value if set, zero value otherwise.
func (o *SpaceWarehouseSchemaOverride) GetProperty() string {
	if o == nil || IsNil(o.Property) {
		var ret string
		return ret
	}
	return *o.Property
}

// GetPropertyOk returns a tuple with the Property field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpaceWarehouseSchemaOverride) GetPropertyOk() (*string, bool) {
	if o == nil || IsNil(o.Property) {
		return nil, false
	}
	return o.Property, true
}

// HasProperty returns a boolean if a field has been set.
func (o *SpaceWarehouseSchemaOverride) HasProperty() bool {
	if o != nil && !IsNil(o.Property) {
		return true
	}

	return false
}

// SetProperty gets a reference to the given string and assigns it to the Property field.
func (o *SpaceWarehouseSchemaOverride) SetProperty(v string) {
	o.Property = &v
}

func (o SpaceWarehouseSchemaOverride) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SpaceWarehouseSchemaOverride) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["collection"] = o.Collection
	toSerialize["enabled"] = o.Enabled
	if !IsNil(o.Property) {
		toSerialize["property"] = o.Property
	}
	return toSerialize, nil
}

type NullableSpaceWarehouseSchemaOverride struct {
	value *SpaceWarehouseSchemaOverride
	isSet bool
}

func (v NullableSpaceWarehouseSchemaOverride) Get() *SpaceWarehouseSchemaOverride {
	return v.value
}

func (v *NullableSpaceWarehouseSchemaOverride) Set(val *SpaceWarehouseSchemaOverride) {
	v.value = val
	v.isSet = true
}

func (v NullableSpaceWarehouseSchemaOverride) IsSet() bool {
	return v.isSet
}

func (v *NullableSpaceWarehouseSchemaOverride) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpaceWarehouseSchemaOverride(
	val *SpaceWarehouseSchemaOverride,
) *NullableSpaceWarehouseSchemaOverride {
	return &NullableSpaceWarehouseSchemaOverride{value: val, isSet: true}
}

func (v NullableSpaceWarehouseSchemaOverride) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpaceWarehouseSchemaOverride) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
