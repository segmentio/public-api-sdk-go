/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 36.3.1
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// CreateValidationInWarehouseV1Input Verifies a set of Warehouse credentials by attempting to connect to it.
type CreateValidationInWarehouseV1Input struct {
	// The id of the Warehouse metadata type.
	MetadataId string `json:"metadataId"`
	// The settings to check.
	Settings NullableModelMap `json:"settings"`
}

// NewCreateValidationInWarehouseV1Input instantiates a new CreateValidationInWarehouseV1Input object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateValidationInWarehouseV1Input(
	metadataId string,
	settings NullableModelMap,
) *CreateValidationInWarehouseV1Input {
	this := CreateValidationInWarehouseV1Input{}
	this.MetadataId = metadataId
	this.Settings = settings
	return &this
}

// NewCreateValidationInWarehouseV1InputWithDefaults instantiates a new CreateValidationInWarehouseV1Input object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateValidationInWarehouseV1InputWithDefaults() *CreateValidationInWarehouseV1Input {
	this := CreateValidationInWarehouseV1Input{}
	return &this
}

// GetMetadataId returns the MetadataId field value
func (o *CreateValidationInWarehouseV1Input) GetMetadataId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MetadataId
}

// GetMetadataIdOk returns a tuple with the MetadataId field value
// and a boolean to check if the value has been set.
func (o *CreateValidationInWarehouseV1Input) GetMetadataIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetadataId, true
}

// SetMetadataId sets field value
func (o *CreateValidationInWarehouseV1Input) SetMetadataId(v string) {
	o.MetadataId = v
}

// GetSettings returns the Settings field value
// If the value is explicit nil, the zero value for ModelMap will be returned
func (o *CreateValidationInWarehouseV1Input) GetSettings() ModelMap {
	if o == nil || o.Settings.Get() == nil {
		var ret ModelMap
		return ret
	}

	return *o.Settings.Get()
}

// GetSettingsOk returns a tuple with the Settings field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateValidationInWarehouseV1Input) GetSettingsOk() (*ModelMap, bool) {
	if o == nil {
		return nil, false
	}
	return o.Settings.Get(), o.Settings.IsSet()
}

// SetSettings sets field value
func (o *CreateValidationInWarehouseV1Input) SetSettings(v ModelMap) {
	o.Settings.Set(&v)
}

func (o CreateValidationInWarehouseV1Input) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["metadataId"] = o.MetadataId
	}
	if true {
		toSerialize["settings"] = o.Settings.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableCreateValidationInWarehouseV1Input struct {
	value *CreateValidationInWarehouseV1Input
	isSet bool
}

func (v NullableCreateValidationInWarehouseV1Input) Get() *CreateValidationInWarehouseV1Input {
	return v.value
}

func (v *NullableCreateValidationInWarehouseV1Input) Set(val *CreateValidationInWarehouseV1Input) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateValidationInWarehouseV1Input) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateValidationInWarehouseV1Input) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateValidationInWarehouseV1Input(
	val *CreateValidationInWarehouseV1Input,
) *NullableCreateValidationInWarehouseV1Input {
	return &NullableCreateValidationInWarehouseV1Input{value: val, isSet: true}
}

func (v NullableCreateValidationInWarehouseV1Input) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateValidationInWarehouseV1Input) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
