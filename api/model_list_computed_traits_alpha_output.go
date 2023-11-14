/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 38.0.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the ListComputedTraitsAlphaOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListComputedTraitsAlphaOutput{}

// ListComputedTraitsAlphaOutput List computed traits endpoint output.
type ListComputedTraitsAlphaOutput struct {
	// A list of computed trait summary results.
	ComputedTraits []ComputedTraitSummary `json:"computedTraits"`
	Pagination     PaginationOutput       `json:"pagination"`
}

// NewListComputedTraitsAlphaOutput instantiates a new ListComputedTraitsAlphaOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListComputedTraitsAlphaOutput(
	computedTraits []ComputedTraitSummary,
	pagination PaginationOutput,
) *ListComputedTraitsAlphaOutput {
	this := ListComputedTraitsAlphaOutput{}
	this.ComputedTraits = computedTraits
	this.Pagination = pagination
	return &this
}

// NewListComputedTraitsAlphaOutputWithDefaults instantiates a new ListComputedTraitsAlphaOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListComputedTraitsAlphaOutputWithDefaults() *ListComputedTraitsAlphaOutput {
	this := ListComputedTraitsAlphaOutput{}
	return &this
}

// GetComputedTraits returns the ComputedTraits field value
func (o *ListComputedTraitsAlphaOutput) GetComputedTraits() []ComputedTraitSummary {
	if o == nil {
		var ret []ComputedTraitSummary
		return ret
	}

	return o.ComputedTraits
}

// GetComputedTraitsOk returns a tuple with the ComputedTraits field value
// and a boolean to check if the value has been set.
func (o *ListComputedTraitsAlphaOutput) GetComputedTraitsOk() ([]ComputedTraitSummary, bool) {
	if o == nil {
		return nil, false
	}
	return o.ComputedTraits, true
}

// SetComputedTraits sets field value
func (o *ListComputedTraitsAlphaOutput) SetComputedTraits(v []ComputedTraitSummary) {
	o.ComputedTraits = v
}

// GetPagination returns the Pagination field value
func (o *ListComputedTraitsAlphaOutput) GetPagination() PaginationOutput {
	if o == nil {
		var ret PaginationOutput
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *ListComputedTraitsAlphaOutput) GetPaginationOk() (*PaginationOutput, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *ListComputedTraitsAlphaOutput) SetPagination(v PaginationOutput) {
	o.Pagination = v
}

func (o ListComputedTraitsAlphaOutput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListComputedTraitsAlphaOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["computedTraits"] = o.ComputedTraits
	toSerialize["pagination"] = o.Pagination
	return toSerialize, nil
}

type NullableListComputedTraitsAlphaOutput struct {
	value *ListComputedTraitsAlphaOutput
	isSet bool
}

func (v NullableListComputedTraitsAlphaOutput) Get() *ListComputedTraitsAlphaOutput {
	return v.value
}

func (v *NullableListComputedTraitsAlphaOutput) Set(val *ListComputedTraitsAlphaOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableListComputedTraitsAlphaOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableListComputedTraitsAlphaOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListComputedTraitsAlphaOutput(
	val *ListComputedTraitsAlphaOutput,
) *NullableListComputedTraitsAlphaOutput {
	return &NullableListComputedTraitsAlphaOutput{value: val, isSet: true}
}

func (v NullableListComputedTraitsAlphaOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListComputedTraitsAlphaOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
