/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 34.3.5
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// GetWarehouseMetadataV1Output Returns a Warehouse catalog item looked up by id.
type GetWarehouseMetadataV1Output struct {
	WarehouseMetadata WarehouseMetadata `json:"warehouseMetadata"`
}

// NewGetWarehouseMetadataV1Output instantiates a new GetWarehouseMetadataV1Output object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetWarehouseMetadataV1Output(
	warehouseMetadata WarehouseMetadata,
) *GetWarehouseMetadataV1Output {
	this := GetWarehouseMetadataV1Output{}
	this.WarehouseMetadata = warehouseMetadata
	return &this
}

// NewGetWarehouseMetadataV1OutputWithDefaults instantiates a new GetWarehouseMetadataV1Output object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetWarehouseMetadataV1OutputWithDefaults() *GetWarehouseMetadataV1Output {
	this := GetWarehouseMetadataV1Output{}
	return &this
}

// GetWarehouseMetadata returns the WarehouseMetadata field value
func (o *GetWarehouseMetadataV1Output) GetWarehouseMetadata() WarehouseMetadata {
	if o == nil {
		var ret WarehouseMetadata
		return ret
	}

	return o.WarehouseMetadata
}

// GetWarehouseMetadataOk returns a tuple with the WarehouseMetadata field value
// and a boolean to check if the value has been set.
func (o *GetWarehouseMetadataV1Output) GetWarehouseMetadataOk() (*WarehouseMetadata, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WarehouseMetadata, true
}

// SetWarehouseMetadata sets field value
func (o *GetWarehouseMetadataV1Output) SetWarehouseMetadata(v WarehouseMetadata) {
	o.WarehouseMetadata = v
}

func (o GetWarehouseMetadataV1Output) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["warehouseMetadata"] = o.WarehouseMetadata
	}
	return json.Marshal(toSerialize)
}

type NullableGetWarehouseMetadataV1Output struct {
	value *GetWarehouseMetadataV1Output
	isSet bool
}

func (v NullableGetWarehouseMetadataV1Output) Get() *GetWarehouseMetadataV1Output {
	return v.value
}

func (v *NullableGetWarehouseMetadataV1Output) Set(val *GetWarehouseMetadataV1Output) {
	v.value = val
	v.isSet = true
}

func (v NullableGetWarehouseMetadataV1Output) IsSet() bool {
	return v.isSet
}

func (v *NullableGetWarehouseMetadataV1Output) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetWarehouseMetadataV1Output(
	val *GetWarehouseMetadataV1Output,
) *NullableGetWarehouseMetadataV1Output {
	return &NullableGetWarehouseMetadataV1Output{value: val, isSet: true}
}

func (v NullableGetWarehouseMetadataV1Output) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetWarehouseMetadataV1Output) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
