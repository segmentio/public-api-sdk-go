/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 50.0.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the ListFiltersFromDestinationV1Output type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListFiltersFromDestinationV1Output{}

// ListFiltersFromDestinationV1Output Output for ListDestinationFiltersV1.
type ListFiltersFromDestinationV1Output struct {
	// A list of the filters that belong to the specified Destination instance.
	Filters    []DestinationFilterV1 `json:"filters"`
	Pagination PaginationOutput      `json:"pagination"`
}

// NewListFiltersFromDestinationV1Output instantiates a new ListFiltersFromDestinationV1Output object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListFiltersFromDestinationV1Output(
	filters []DestinationFilterV1,
	pagination PaginationOutput,
) *ListFiltersFromDestinationV1Output {
	this := ListFiltersFromDestinationV1Output{}
	this.Filters = filters
	this.Pagination = pagination
	return &this
}

// NewListFiltersFromDestinationV1OutputWithDefaults instantiates a new ListFiltersFromDestinationV1Output object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListFiltersFromDestinationV1OutputWithDefaults() *ListFiltersFromDestinationV1Output {
	this := ListFiltersFromDestinationV1Output{}
	return &this
}

// GetFilters returns the Filters field value
func (o *ListFiltersFromDestinationV1Output) GetFilters() []DestinationFilterV1 {
	if o == nil {
		var ret []DestinationFilterV1
		return ret
	}

	return o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value
// and a boolean to check if the value has been set.
func (o *ListFiltersFromDestinationV1Output) GetFiltersOk() ([]DestinationFilterV1, bool) {
	if o == nil {
		return nil, false
	}
	return o.Filters, true
}

// SetFilters sets field value
func (o *ListFiltersFromDestinationV1Output) SetFilters(v []DestinationFilterV1) {
	o.Filters = v
}

// GetPagination returns the Pagination field value
func (o *ListFiltersFromDestinationV1Output) GetPagination() PaginationOutput {
	if o == nil {
		var ret PaginationOutput
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *ListFiltersFromDestinationV1Output) GetPaginationOk() (*PaginationOutput, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *ListFiltersFromDestinationV1Output) SetPagination(v PaginationOutput) {
	o.Pagination = v
}

func (o ListFiltersFromDestinationV1Output) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListFiltersFromDestinationV1Output) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["filters"] = o.Filters
	toSerialize["pagination"] = o.Pagination
	return toSerialize, nil
}

type NullableListFiltersFromDestinationV1Output struct {
	value *ListFiltersFromDestinationV1Output
	isSet bool
}

func (v NullableListFiltersFromDestinationV1Output) Get() *ListFiltersFromDestinationV1Output {
	return v.value
}

func (v *NullableListFiltersFromDestinationV1Output) Set(val *ListFiltersFromDestinationV1Output) {
	v.value = val
	v.isSet = true
}

func (v NullableListFiltersFromDestinationV1Output) IsSet() bool {
	return v.isSet
}

func (v *NullableListFiltersFromDestinationV1Output) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListFiltersFromDestinationV1Output(
	val *ListFiltersFromDestinationV1Output,
) *NullableListFiltersFromDestinationV1Output {
	return &NullableListFiltersFromDestinationV1Output{value: val, isSet: true}
}

func (v NullableListFiltersFromDestinationV1Output) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListFiltersFromDestinationV1Output) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
