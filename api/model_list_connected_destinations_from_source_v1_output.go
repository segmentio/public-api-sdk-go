/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 36.3.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// ListConnectedDestinationsFromSourceV1Output Returns a list of Destinations connected to a Source.
type ListConnectedDestinationsFromSourceV1Output struct {
	// A list that contains the Destinations connected to the Source.
	Destinations []DestinationV1 `json:"destinations"`
	Pagination   Pagination      `json:"pagination"`
}

// NewListConnectedDestinationsFromSourceV1Output instantiates a new ListConnectedDestinationsFromSourceV1Output object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListConnectedDestinationsFromSourceV1Output(
	destinations []DestinationV1,
	pagination Pagination,
) *ListConnectedDestinationsFromSourceV1Output {
	this := ListConnectedDestinationsFromSourceV1Output{}
	this.Destinations = destinations
	this.Pagination = pagination
	return &this
}

// NewListConnectedDestinationsFromSourceV1OutputWithDefaults instantiates a new ListConnectedDestinationsFromSourceV1Output object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListConnectedDestinationsFromSourceV1OutputWithDefaults() *ListConnectedDestinationsFromSourceV1Output {
	this := ListConnectedDestinationsFromSourceV1Output{}
	return &this
}

// GetDestinations returns the Destinations field value
func (o *ListConnectedDestinationsFromSourceV1Output) GetDestinations() []DestinationV1 {
	if o == nil {
		var ret []DestinationV1
		return ret
	}

	return o.Destinations
}

// GetDestinationsOk returns a tuple with the Destinations field value
// and a boolean to check if the value has been set.
func (o *ListConnectedDestinationsFromSourceV1Output) GetDestinationsOk() ([]DestinationV1, bool) {
	if o == nil {
		return nil, false
	}
	return o.Destinations, true
}

// SetDestinations sets field value
func (o *ListConnectedDestinationsFromSourceV1Output) SetDestinations(v []DestinationV1) {
	o.Destinations = v
}

// GetPagination returns the Pagination field value
func (o *ListConnectedDestinationsFromSourceV1Output) GetPagination() Pagination {
	if o == nil {
		var ret Pagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *ListConnectedDestinationsFromSourceV1Output) GetPaginationOk() (*Pagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *ListConnectedDestinationsFromSourceV1Output) SetPagination(v Pagination) {
	o.Pagination = v
}

func (o ListConnectedDestinationsFromSourceV1Output) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["destinations"] = o.Destinations
	}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	return json.Marshal(toSerialize)
}

type NullableListConnectedDestinationsFromSourceV1Output struct {
	value *ListConnectedDestinationsFromSourceV1Output
	isSet bool
}

func (v NullableListConnectedDestinationsFromSourceV1Output) Get() *ListConnectedDestinationsFromSourceV1Output {
	return v.value
}

func (v *NullableListConnectedDestinationsFromSourceV1Output) Set(
	val *ListConnectedDestinationsFromSourceV1Output,
) {
	v.value = val
	v.isSet = true
}

func (v NullableListConnectedDestinationsFromSourceV1Output) IsSet() bool {
	return v.isSet
}

func (v *NullableListConnectedDestinationsFromSourceV1Output) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListConnectedDestinationsFromSourceV1Output(
	val *ListConnectedDestinationsFromSourceV1Output,
) *NullableListConnectedDestinationsFromSourceV1Output {
	return &NullableListConnectedDestinationsFromSourceV1Output{value: val, isSet: true}
}

func (v NullableListConnectedDestinationsFromSourceV1Output) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListConnectedDestinationsFromSourceV1Output) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
