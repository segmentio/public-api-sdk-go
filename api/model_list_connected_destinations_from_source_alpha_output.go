/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 53.2.1
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the ListConnectedDestinationsFromSourceAlphaOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListConnectedDestinationsFromSourceAlphaOutput{}

// ListConnectedDestinationsFromSourceAlphaOutput Returns a list of Destinations connected to a Source.
type ListConnectedDestinationsFromSourceAlphaOutput struct {
	// A list that contains the Destinations connected to the Source.
	Destinations []DestinationV1  `json:"destinations"`
	Pagination   PaginationOutput `json:"pagination"`
}

// NewListConnectedDestinationsFromSourceAlphaOutput instantiates a new ListConnectedDestinationsFromSourceAlphaOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListConnectedDestinationsFromSourceAlphaOutput(
	destinations []DestinationV1,
	pagination PaginationOutput,
) *ListConnectedDestinationsFromSourceAlphaOutput {
	this := ListConnectedDestinationsFromSourceAlphaOutput{}
	this.Destinations = destinations
	this.Pagination = pagination
	return &this
}

// NewListConnectedDestinationsFromSourceAlphaOutputWithDefaults instantiates a new ListConnectedDestinationsFromSourceAlphaOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListConnectedDestinationsFromSourceAlphaOutputWithDefaults() *ListConnectedDestinationsFromSourceAlphaOutput {
	this := ListConnectedDestinationsFromSourceAlphaOutput{}
	return &this
}

// GetDestinations returns the Destinations field value
func (o *ListConnectedDestinationsFromSourceAlphaOutput) GetDestinations() []DestinationV1 {
	if o == nil {
		var ret []DestinationV1
		return ret
	}

	return o.Destinations
}

// GetDestinationsOk returns a tuple with the Destinations field value
// and a boolean to check if the value has been set.
func (o *ListConnectedDestinationsFromSourceAlphaOutput) GetDestinationsOk() ([]DestinationV1, bool) {
	if o == nil {
		return nil, false
	}
	return o.Destinations, true
}

// SetDestinations sets field value
func (o *ListConnectedDestinationsFromSourceAlphaOutput) SetDestinations(v []DestinationV1) {
	o.Destinations = v
}

// GetPagination returns the Pagination field value
func (o *ListConnectedDestinationsFromSourceAlphaOutput) GetPagination() PaginationOutput {
	if o == nil {
		var ret PaginationOutput
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *ListConnectedDestinationsFromSourceAlphaOutput) GetPaginationOk() (*PaginationOutput, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *ListConnectedDestinationsFromSourceAlphaOutput) SetPagination(v PaginationOutput) {
	o.Pagination = v
}

func (o ListConnectedDestinationsFromSourceAlphaOutput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListConnectedDestinationsFromSourceAlphaOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["destinations"] = o.Destinations
	toSerialize["pagination"] = o.Pagination
	return toSerialize, nil
}

type NullableListConnectedDestinationsFromSourceAlphaOutput struct {
	value *ListConnectedDestinationsFromSourceAlphaOutput
	isSet bool
}

func (v NullableListConnectedDestinationsFromSourceAlphaOutput) Get() *ListConnectedDestinationsFromSourceAlphaOutput {
	return v.value
}

func (v *NullableListConnectedDestinationsFromSourceAlphaOutput) Set(
	val *ListConnectedDestinationsFromSourceAlphaOutput,
) {
	v.value = val
	v.isSet = true
}

func (v NullableListConnectedDestinationsFromSourceAlphaOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableListConnectedDestinationsFromSourceAlphaOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListConnectedDestinationsFromSourceAlphaOutput(
	val *ListConnectedDestinationsFromSourceAlphaOutput,
) *NullableListConnectedDestinationsFromSourceAlphaOutput {
	return &NullableListConnectedDestinationsFromSourceAlphaOutput{value: val, isSet: true}
}

func (v NullableListConnectedDestinationsFromSourceAlphaOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListConnectedDestinationsFromSourceAlphaOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
