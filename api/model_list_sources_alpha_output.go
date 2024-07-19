/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 50.4.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the ListSourcesAlphaOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListSourcesAlphaOutput{}

// ListSourcesAlphaOutput Returns a list of Sources that belong to the current Workspace.
type ListSourcesAlphaOutput struct {
	// A list of Sources that belong to the Workspace.
	Sources    []SourceAlpha    `json:"sources"`
	Pagination PaginationOutput `json:"pagination"`
}

// NewListSourcesAlphaOutput instantiates a new ListSourcesAlphaOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListSourcesAlphaOutput(
	sources []SourceAlpha,
	pagination PaginationOutput,
) *ListSourcesAlphaOutput {
	this := ListSourcesAlphaOutput{}
	this.Sources = sources
	this.Pagination = pagination
	return &this
}

// NewListSourcesAlphaOutputWithDefaults instantiates a new ListSourcesAlphaOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListSourcesAlphaOutputWithDefaults() *ListSourcesAlphaOutput {
	this := ListSourcesAlphaOutput{}
	return &this
}

// GetSources returns the Sources field value
func (o *ListSourcesAlphaOutput) GetSources() []SourceAlpha {
	if o == nil {
		var ret []SourceAlpha
		return ret
	}

	return o.Sources
}

// GetSourcesOk returns a tuple with the Sources field value
// and a boolean to check if the value has been set.
func (o *ListSourcesAlphaOutput) GetSourcesOk() ([]SourceAlpha, bool) {
	if o == nil {
		return nil, false
	}
	return o.Sources, true
}

// SetSources sets field value
func (o *ListSourcesAlphaOutput) SetSources(v []SourceAlpha) {
	o.Sources = v
}

// GetPagination returns the Pagination field value
func (o *ListSourcesAlphaOutput) GetPagination() PaginationOutput {
	if o == nil {
		var ret PaginationOutput
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *ListSourcesAlphaOutput) GetPaginationOk() (*PaginationOutput, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *ListSourcesAlphaOutput) SetPagination(v PaginationOutput) {
	o.Pagination = v
}

func (o ListSourcesAlphaOutput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListSourcesAlphaOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sources"] = o.Sources
	toSerialize["pagination"] = o.Pagination
	return toSerialize, nil
}

type NullableListSourcesAlphaOutput struct {
	value *ListSourcesAlphaOutput
	isSet bool
}

func (v NullableListSourcesAlphaOutput) Get() *ListSourcesAlphaOutput {
	return v.value
}

func (v *NullableListSourcesAlphaOutput) Set(val *ListSourcesAlphaOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableListSourcesAlphaOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableListSourcesAlphaOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListSourcesAlphaOutput(
	val *ListSourcesAlphaOutput,
) *NullableListSourcesAlphaOutput {
	return &NullableListSourcesAlphaOutput{value: val, isSet: true}
}

func (v NullableListSourcesAlphaOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListSourcesAlphaOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
