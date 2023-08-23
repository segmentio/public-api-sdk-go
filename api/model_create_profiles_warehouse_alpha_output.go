/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 36.1.1
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// CreateProfilesWarehouseAlphaOutput Returns the newly created Warehouse.
type CreateProfilesWarehouseAlphaOutput struct {
	ProfilesWarehouse ProfilesWarehouse `json:"profilesWarehouse"`
}

// NewCreateProfilesWarehouseAlphaOutput instantiates a new CreateProfilesWarehouseAlphaOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateProfilesWarehouseAlphaOutput(
	profilesWarehouse ProfilesWarehouse,
) *CreateProfilesWarehouseAlphaOutput {
	this := CreateProfilesWarehouseAlphaOutput{}
	this.ProfilesWarehouse = profilesWarehouse
	return &this
}

// NewCreateProfilesWarehouseAlphaOutputWithDefaults instantiates a new CreateProfilesWarehouseAlphaOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateProfilesWarehouseAlphaOutputWithDefaults() *CreateProfilesWarehouseAlphaOutput {
	this := CreateProfilesWarehouseAlphaOutput{}
	return &this
}

// GetProfilesWarehouse returns the ProfilesWarehouse field value
func (o *CreateProfilesWarehouseAlphaOutput) GetProfilesWarehouse() ProfilesWarehouse {
	if o == nil {
		var ret ProfilesWarehouse
		return ret
	}

	return o.ProfilesWarehouse
}

// GetProfilesWarehouseOk returns a tuple with the ProfilesWarehouse field value
// and a boolean to check if the value has been set.
func (o *CreateProfilesWarehouseAlphaOutput) GetProfilesWarehouseOk() (*ProfilesWarehouse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProfilesWarehouse, true
}

// SetProfilesWarehouse sets field value
func (o *CreateProfilesWarehouseAlphaOutput) SetProfilesWarehouse(v ProfilesWarehouse) {
	o.ProfilesWarehouse = v
}

func (o CreateProfilesWarehouseAlphaOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["profilesWarehouse"] = o.ProfilesWarehouse
	}
	return json.Marshal(toSerialize)
}

type NullableCreateProfilesWarehouseAlphaOutput struct {
	value *CreateProfilesWarehouseAlphaOutput
	isSet bool
}

func (v NullableCreateProfilesWarehouseAlphaOutput) Get() *CreateProfilesWarehouseAlphaOutput {
	return v.value
}

func (v *NullableCreateProfilesWarehouseAlphaOutput) Set(val *CreateProfilesWarehouseAlphaOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateProfilesWarehouseAlphaOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateProfilesWarehouseAlphaOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateProfilesWarehouseAlphaOutput(
	val *CreateProfilesWarehouseAlphaOutput,
) *NullableCreateProfilesWarehouseAlphaOutput {
	return &NullableCreateProfilesWarehouseAlphaOutput{value: val, isSet: true}
}

func (v NullableCreateProfilesWarehouseAlphaOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateProfilesWarehouseAlphaOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
