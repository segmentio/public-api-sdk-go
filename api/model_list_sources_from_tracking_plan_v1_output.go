/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 46.0.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the ListSourcesFromTrackingPlanV1Output type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListSourcesFromTrackingPlanV1Output{}

// ListSourcesFromTrackingPlanV1Output Lists all Sources associated with a Tracking Plan.
type ListSourcesFromTrackingPlanV1Output struct {
	// A paginated list of Sources associated with the Tracking Plan.
	Sources    []SourceV1       `json:"sources"`
	Pagination PaginationOutput `json:"pagination"`
}

// NewListSourcesFromTrackingPlanV1Output instantiates a new ListSourcesFromTrackingPlanV1Output object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListSourcesFromTrackingPlanV1Output(
	sources []SourceV1,
	pagination PaginationOutput,
) *ListSourcesFromTrackingPlanV1Output {
	this := ListSourcesFromTrackingPlanV1Output{}
	this.Sources = sources
	this.Pagination = pagination
	return &this
}

// NewListSourcesFromTrackingPlanV1OutputWithDefaults instantiates a new ListSourcesFromTrackingPlanV1Output object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListSourcesFromTrackingPlanV1OutputWithDefaults() *ListSourcesFromTrackingPlanV1Output {
	this := ListSourcesFromTrackingPlanV1Output{}
	return &this
}

// GetSources returns the Sources field value
func (o *ListSourcesFromTrackingPlanV1Output) GetSources() []SourceV1 {
	if o == nil {
		var ret []SourceV1
		return ret
	}

	return o.Sources
}

// GetSourcesOk returns a tuple with the Sources field value
// and a boolean to check if the value has been set.
func (o *ListSourcesFromTrackingPlanV1Output) GetSourcesOk() ([]SourceV1, bool) {
	if o == nil {
		return nil, false
	}
	return o.Sources, true
}

// SetSources sets field value
func (o *ListSourcesFromTrackingPlanV1Output) SetSources(v []SourceV1) {
	o.Sources = v
}

// GetPagination returns the Pagination field value
func (o *ListSourcesFromTrackingPlanV1Output) GetPagination() PaginationOutput {
	if o == nil {
		var ret PaginationOutput
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *ListSourcesFromTrackingPlanV1Output) GetPaginationOk() (*PaginationOutput, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *ListSourcesFromTrackingPlanV1Output) SetPagination(v PaginationOutput) {
	o.Pagination = v
}

func (o ListSourcesFromTrackingPlanV1Output) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListSourcesFromTrackingPlanV1Output) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sources"] = o.Sources
	toSerialize["pagination"] = o.Pagination
	return toSerialize, nil
}

type NullableListSourcesFromTrackingPlanV1Output struct {
	value *ListSourcesFromTrackingPlanV1Output
	isSet bool
}

func (v NullableListSourcesFromTrackingPlanV1Output) Get() *ListSourcesFromTrackingPlanV1Output {
	return v.value
}

func (v *NullableListSourcesFromTrackingPlanV1Output) Set(
	val *ListSourcesFromTrackingPlanV1Output,
) {
	v.value = val
	v.isSet = true
}

func (v NullableListSourcesFromTrackingPlanV1Output) IsSet() bool {
	return v.isSet
}

func (v *NullableListSourcesFromTrackingPlanV1Output) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListSourcesFromTrackingPlanV1Output(
	val *ListSourcesFromTrackingPlanV1Output,
) *NullableListSourcesFromTrackingPlanV1Output {
	return &NullableListSourcesFromTrackingPlanV1Output{value: val, isSet: true}
}

func (v NullableListSourcesFromTrackingPlanV1Output) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListSourcesFromTrackingPlanV1Output) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
