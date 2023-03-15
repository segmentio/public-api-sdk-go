/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 34.3.3
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// ListProfilesWarehouseInSpaceAlphaOutput Returns all Profiles Warehouse based on spaceID and pagination.
type ListProfilesWarehouseInSpaceAlphaOutput struct {
	// A list of Warehouses that belong to the Workspace.
	ProfilesWarehouses []ProfilesWarehouseAlpha `json:"profilesWarehouses"`
	Pagination         *Pagination              `json:"pagination,omitempty"`
}

// NewListProfilesWarehouseInSpaceAlphaOutput instantiates a new ListProfilesWarehouseInSpaceAlphaOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListProfilesWarehouseInSpaceAlphaOutput(
	profilesWarehouses []ProfilesWarehouseAlpha,
) *ListProfilesWarehouseInSpaceAlphaOutput {
	this := ListProfilesWarehouseInSpaceAlphaOutput{}
	this.ProfilesWarehouses = profilesWarehouses
	return &this
}

// NewListProfilesWarehouseInSpaceAlphaOutputWithDefaults instantiates a new ListProfilesWarehouseInSpaceAlphaOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListProfilesWarehouseInSpaceAlphaOutputWithDefaults() *ListProfilesWarehouseInSpaceAlphaOutput {
	this := ListProfilesWarehouseInSpaceAlphaOutput{}
	return &this
}

// GetProfilesWarehouses returns the ProfilesWarehouses field value
func (o *ListProfilesWarehouseInSpaceAlphaOutput) GetProfilesWarehouses() []ProfilesWarehouseAlpha {
	if o == nil {
		var ret []ProfilesWarehouseAlpha
		return ret
	}

	return o.ProfilesWarehouses
}

// GetProfilesWarehousesOk returns a tuple with the ProfilesWarehouses field value
// and a boolean to check if the value has been set.
func (o *ListProfilesWarehouseInSpaceAlphaOutput) GetProfilesWarehousesOk() ([]ProfilesWarehouseAlpha, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProfilesWarehouses, true
}

// SetProfilesWarehouses sets field value
func (o *ListProfilesWarehouseInSpaceAlphaOutput) SetProfilesWarehouses(
	v []ProfilesWarehouseAlpha,
) {
	o.ProfilesWarehouses = v
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *ListProfilesWarehouseInSpaceAlphaOutput) GetPagination() Pagination {
	if o == nil || o.Pagination == nil {
		var ret Pagination
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListProfilesWarehouseInSpaceAlphaOutput) GetPaginationOk() (*Pagination, bool) {
	if o == nil || o.Pagination == nil {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *ListProfilesWarehouseInSpaceAlphaOutput) HasPagination() bool {
	if o != nil && o.Pagination != nil {
		return true
	}

	return false
}

// SetPagination gets a reference to the given Pagination and assigns it to the Pagination field.
func (o *ListProfilesWarehouseInSpaceAlphaOutput) SetPagination(v Pagination) {
	o.Pagination = &v
}

func (o ListProfilesWarehouseInSpaceAlphaOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["profilesWarehouses"] = o.ProfilesWarehouses
	}
	if o.Pagination != nil {
		toSerialize["pagination"] = o.Pagination
	}
	return json.Marshal(toSerialize)
}

type NullableListProfilesWarehouseInSpaceAlphaOutput struct {
	value *ListProfilesWarehouseInSpaceAlphaOutput
	isSet bool
}

func (v NullableListProfilesWarehouseInSpaceAlphaOutput) Get() *ListProfilesWarehouseInSpaceAlphaOutput {
	return v.value
}

func (v *NullableListProfilesWarehouseInSpaceAlphaOutput) Set(
	val *ListProfilesWarehouseInSpaceAlphaOutput,
) {
	v.value = val
	v.isSet = true
}

func (v NullableListProfilesWarehouseInSpaceAlphaOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableListProfilesWarehouseInSpaceAlphaOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListProfilesWarehouseInSpaceAlphaOutput(
	val *ListProfilesWarehouseInSpaceAlphaOutput,
) *NullableListProfilesWarehouseInSpaceAlphaOutput {
	return &NullableListProfilesWarehouseInSpaceAlphaOutput{value: val, isSet: true}
}

func (v NullableListProfilesWarehouseInSpaceAlphaOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListProfilesWarehouseInSpaceAlphaOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}