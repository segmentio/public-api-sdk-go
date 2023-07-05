/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 36.0.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// UpdateProfilesWarehouseForSpaceWarehouseAlphaOutput Returns the updated Warehouse.
type UpdateProfilesWarehouseForSpaceWarehouseAlphaOutput struct {
	ProfilesWarehouse ProfilesWarehouse1 `json:"profilesWarehouse"`
}

// NewUpdateProfilesWarehouseForSpaceWarehouseAlphaOutput instantiates a new UpdateProfilesWarehouseForSpaceWarehouseAlphaOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateProfilesWarehouseForSpaceWarehouseAlphaOutput(
	profilesWarehouse ProfilesWarehouse1,
) *UpdateProfilesWarehouseForSpaceWarehouseAlphaOutput {
	this := UpdateProfilesWarehouseForSpaceWarehouseAlphaOutput{}
	this.ProfilesWarehouse = profilesWarehouse
	return &this
}

// NewUpdateProfilesWarehouseForSpaceWarehouseAlphaOutputWithDefaults instantiates a new UpdateProfilesWarehouseForSpaceWarehouseAlphaOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateProfilesWarehouseForSpaceWarehouseAlphaOutputWithDefaults() *UpdateProfilesWarehouseForSpaceWarehouseAlphaOutput {
	this := UpdateProfilesWarehouseForSpaceWarehouseAlphaOutput{}
	return &this
}

// GetProfilesWarehouse returns the ProfilesWarehouse field value
func (o *UpdateProfilesWarehouseForSpaceWarehouseAlphaOutput) GetProfilesWarehouse() ProfilesWarehouse1 {
	if o == nil {
		var ret ProfilesWarehouse1
		return ret
	}

	return o.ProfilesWarehouse
}

// GetProfilesWarehouseOk returns a tuple with the ProfilesWarehouse field value
// and a boolean to check if the value has been set.
func (o *UpdateProfilesWarehouseForSpaceWarehouseAlphaOutput) GetProfilesWarehouseOk() (*ProfilesWarehouse1, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProfilesWarehouse, true
}

// SetProfilesWarehouse sets field value
func (o *UpdateProfilesWarehouseForSpaceWarehouseAlphaOutput) SetProfilesWarehouse(
	v ProfilesWarehouse1,
) {
	o.ProfilesWarehouse = v
}

func (o UpdateProfilesWarehouseForSpaceWarehouseAlphaOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["profilesWarehouse"] = o.ProfilesWarehouse
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateProfilesWarehouseForSpaceWarehouseAlphaOutput struct {
	value *UpdateProfilesWarehouseForSpaceWarehouseAlphaOutput
	isSet bool
}

func (v NullableUpdateProfilesWarehouseForSpaceWarehouseAlphaOutput) Get() *UpdateProfilesWarehouseForSpaceWarehouseAlphaOutput {
	return v.value
}

func (v *NullableUpdateProfilesWarehouseForSpaceWarehouseAlphaOutput) Set(
	val *UpdateProfilesWarehouseForSpaceWarehouseAlphaOutput,
) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateProfilesWarehouseForSpaceWarehouseAlphaOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateProfilesWarehouseForSpaceWarehouseAlphaOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateProfilesWarehouseForSpaceWarehouseAlphaOutput(
	val *UpdateProfilesWarehouseForSpaceWarehouseAlphaOutput,
) *NullableUpdateProfilesWarehouseForSpaceWarehouseAlphaOutput {
	return &NullableUpdateProfilesWarehouseForSpaceWarehouseAlphaOutput{value: val, isSet: true}
}

func (v NullableUpdateProfilesWarehouseForSpaceWarehouseAlphaOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateProfilesWarehouseForSpaceWarehouseAlphaOutput) UnmarshalJSON(
	src []byte,
) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
