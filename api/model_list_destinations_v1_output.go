/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 58.13.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the ListDestinationsV1Output type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListDestinationsV1Output{}

// ListDestinationsV1Output Returns all Destinations connected to the current Workspace.
type ListDestinationsV1Output struct {
	// The list that contains the Destinations connected to the Workspace.
	Destinations []DestinationV1  `json:"destinations"`
	Pagination   PaginationOutput `json:"pagination"`
}

// NewListDestinationsV1Output instantiates a new ListDestinationsV1Output object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListDestinationsV1Output(
	destinations []DestinationV1,
	pagination PaginationOutput,
) *ListDestinationsV1Output {
	this := ListDestinationsV1Output{}
	this.Destinations = destinations
	this.Pagination = pagination
	return &this
}

// NewListDestinationsV1OutputWithDefaults instantiates a new ListDestinationsV1Output object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListDestinationsV1OutputWithDefaults() *ListDestinationsV1Output {
	this := ListDestinationsV1Output{}
	return &this
}

// GetDestinations returns the Destinations field value
func (o *ListDestinationsV1Output) GetDestinations() []DestinationV1 {
	if o == nil {
		var ret []DestinationV1
		return ret
	}

	return o.Destinations
}

// GetDestinationsOk returns a tuple with the Destinations field value
// and a boolean to check if the value has been set.
func (o *ListDestinationsV1Output) GetDestinationsOk() ([]DestinationV1, bool) {
	if o == nil {
		return nil, false
	}
	return o.Destinations, true
}

// SetDestinations sets field value
func (o *ListDestinationsV1Output) SetDestinations(v []DestinationV1) {
	o.Destinations = v
}

// GetPagination returns the Pagination field value
func (o *ListDestinationsV1Output) GetPagination() PaginationOutput {
	if o == nil {
		var ret PaginationOutput
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *ListDestinationsV1Output) GetPaginationOk() (*PaginationOutput, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *ListDestinationsV1Output) SetPagination(v PaginationOutput) {
	o.Pagination = v
}

func (o ListDestinationsV1Output) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListDestinationsV1Output) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["destinations"] = o.Destinations
	toSerialize["pagination"] = o.Pagination
	return toSerialize, nil
}

type NullableListDestinationsV1Output struct {
	value *ListDestinationsV1Output
	isSet bool
}

func (v NullableListDestinationsV1Output) Get() *ListDestinationsV1Output {
	return v.value
}

func (v *NullableListDestinationsV1Output) Set(val *ListDestinationsV1Output) {
	v.value = val
	v.isSet = true
}

func (v NullableListDestinationsV1Output) IsSet() bool {
	return v.isSet
}

func (v *NullableListDestinationsV1Output) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListDestinationsV1Output(
	val *ListDestinationsV1Output,
) *NullableListDestinationsV1Output {
	return &NullableListDestinationsV1Output{value: val, isSet: true}
}

func (v NullableListDestinationsV1Output) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListDestinationsV1Output) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
