/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 58.2.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the ListFiltersForSpaceOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListFiltersForSpaceOutput{}

// ListFiltersForSpaceOutput Output for ListFiltersByIntegrationId.
type ListFiltersForSpaceOutput struct {
	// Filter output.
	Filters    []Filter                     `json:"filters,omitempty"`
	Pagination *ListFiltersPaginationOutput `json:"pagination,omitempty"`
}

// NewListFiltersForSpaceOutput instantiates a new ListFiltersForSpaceOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListFiltersForSpaceOutput() *ListFiltersForSpaceOutput {
	this := ListFiltersForSpaceOutput{}
	return &this
}

// NewListFiltersForSpaceOutputWithDefaults instantiates a new ListFiltersForSpaceOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListFiltersForSpaceOutputWithDefaults() *ListFiltersForSpaceOutput {
	this := ListFiltersForSpaceOutput{}
	return &this
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *ListFiltersForSpaceOutput) GetFilters() []Filter {
	if o == nil || IsNil(o.Filters) {
		var ret []Filter
		return ret
	}
	return o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListFiltersForSpaceOutput) GetFiltersOk() ([]Filter, bool) {
	if o == nil || IsNil(o.Filters) {
		return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *ListFiltersForSpaceOutput) HasFilters() bool {
	if o != nil && !IsNil(o.Filters) {
		return true
	}

	return false
}

// SetFilters gets a reference to the given []Filter and assigns it to the Filters field.
func (o *ListFiltersForSpaceOutput) SetFilters(v []Filter) {
	o.Filters = v
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *ListFiltersForSpaceOutput) GetPagination() ListFiltersPaginationOutput {
	if o == nil || IsNil(o.Pagination) {
		var ret ListFiltersPaginationOutput
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListFiltersForSpaceOutput) GetPaginationOk() (*ListFiltersPaginationOutput, bool) {
	if o == nil || IsNil(o.Pagination) {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *ListFiltersForSpaceOutput) HasPagination() bool {
	if o != nil && !IsNil(o.Pagination) {
		return true
	}

	return false
}

// SetPagination gets a reference to the given ListFiltersPaginationOutput and assigns it to the Pagination field.
func (o *ListFiltersForSpaceOutput) SetPagination(v ListFiltersPaginationOutput) {
	o.Pagination = &v
}

func (o ListFiltersForSpaceOutput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListFiltersForSpaceOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Filters) {
		toSerialize["filters"] = o.Filters
	}
	if !IsNil(o.Pagination) {
		toSerialize["pagination"] = o.Pagination
	}
	return toSerialize, nil
}

type NullableListFiltersForSpaceOutput struct {
	value *ListFiltersForSpaceOutput
	isSet bool
}

func (v NullableListFiltersForSpaceOutput) Get() *ListFiltersForSpaceOutput {
	return v.value
}

func (v *NullableListFiltersForSpaceOutput) Set(val *ListFiltersForSpaceOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableListFiltersForSpaceOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableListFiltersForSpaceOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListFiltersForSpaceOutput(
	val *ListFiltersForSpaceOutput,
) *NullableListFiltersForSpaceOutput {
	return &NullableListFiltersForSpaceOutput{value: val, isSet: true}
}

func (v NullableListFiltersForSpaceOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListFiltersForSpaceOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
