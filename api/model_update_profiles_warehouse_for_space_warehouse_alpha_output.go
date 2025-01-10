/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 57.3.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the UpdateProfilesWarehouseForSpaceWarehouseAlphaOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateProfilesWarehouseForSpaceWarehouseAlphaOutput{}

// UpdateProfilesWarehouseForSpaceWarehouseAlphaOutput Returns the updated Warehouse.
type UpdateProfilesWarehouseForSpaceWarehouseAlphaOutput struct {
	ProfilesWarehouse ProfilesWarehouseAlpha `json:"profilesWarehouse"`
}

// NewUpdateProfilesWarehouseForSpaceWarehouseAlphaOutput instantiates a new UpdateProfilesWarehouseForSpaceWarehouseAlphaOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateProfilesWarehouseForSpaceWarehouseAlphaOutput(
	profilesWarehouse ProfilesWarehouseAlpha,
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
func (o *UpdateProfilesWarehouseForSpaceWarehouseAlphaOutput) GetProfilesWarehouse() ProfilesWarehouseAlpha {
	if o == nil {
		var ret ProfilesWarehouseAlpha
		return ret
	}

	return o.ProfilesWarehouse
}

// GetProfilesWarehouseOk returns a tuple with the ProfilesWarehouse field value
// and a boolean to check if the value has been set.
func (o *UpdateProfilesWarehouseForSpaceWarehouseAlphaOutput) GetProfilesWarehouseOk() (*ProfilesWarehouseAlpha, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProfilesWarehouse, true
}

// SetProfilesWarehouse sets field value
func (o *UpdateProfilesWarehouseForSpaceWarehouseAlphaOutput) SetProfilesWarehouse(
	v ProfilesWarehouseAlpha,
) {
	o.ProfilesWarehouse = v
}

func (o UpdateProfilesWarehouseForSpaceWarehouseAlphaOutput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateProfilesWarehouseForSpaceWarehouseAlphaOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["profilesWarehouse"] = o.ProfilesWarehouse
	return toSerialize, nil
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
