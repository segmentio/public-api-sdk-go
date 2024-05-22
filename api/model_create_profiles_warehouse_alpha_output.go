/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 50.2.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the CreateProfilesWarehouseAlphaOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateProfilesWarehouseAlphaOutput{}

// CreateProfilesWarehouseAlphaOutput Returns the newly created Warehouse.
type CreateProfilesWarehouseAlphaOutput struct {
	ProfilesWarehouse ProfilesWarehouseAlpha `json:"profilesWarehouse"`
}

// NewCreateProfilesWarehouseAlphaOutput instantiates a new CreateProfilesWarehouseAlphaOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateProfilesWarehouseAlphaOutput(
	profilesWarehouse ProfilesWarehouseAlpha,
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
func (o *CreateProfilesWarehouseAlphaOutput) GetProfilesWarehouse() ProfilesWarehouseAlpha {
	if o == nil {
		var ret ProfilesWarehouseAlpha
		return ret
	}

	return o.ProfilesWarehouse
}

// GetProfilesWarehouseOk returns a tuple with the ProfilesWarehouse field value
// and a boolean to check if the value has been set.
func (o *CreateProfilesWarehouseAlphaOutput) GetProfilesWarehouseOk() (*ProfilesWarehouseAlpha, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProfilesWarehouse, true
}

// SetProfilesWarehouse sets field value
func (o *CreateProfilesWarehouseAlphaOutput) SetProfilesWarehouse(v ProfilesWarehouseAlpha) {
	o.ProfilesWarehouse = v
}

func (o CreateProfilesWarehouseAlphaOutput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateProfilesWarehouseAlphaOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["profilesWarehouse"] = o.ProfilesWarehouse
	return toSerialize, nil
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
