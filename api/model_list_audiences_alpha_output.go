/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 57.1.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the ListAudiencesAlphaOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListAudiencesAlphaOutput{}

// ListAudiencesAlphaOutput List audiences endpoint output.
type ListAudiencesAlphaOutput struct {
	// A list of audience summary results.
	Audiences  []AudienceSummary `json:"audiences"`
	Pagination PaginationOutput  `json:"pagination"`
}

// NewListAudiencesAlphaOutput instantiates a new ListAudiencesAlphaOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListAudiencesAlphaOutput(
	audiences []AudienceSummary,
	pagination PaginationOutput,
) *ListAudiencesAlphaOutput {
	this := ListAudiencesAlphaOutput{}
	this.Audiences = audiences
	this.Pagination = pagination
	return &this
}

// NewListAudiencesAlphaOutputWithDefaults instantiates a new ListAudiencesAlphaOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListAudiencesAlphaOutputWithDefaults() *ListAudiencesAlphaOutput {
	this := ListAudiencesAlphaOutput{}
	return &this
}

// GetAudiences returns the Audiences field value
func (o *ListAudiencesAlphaOutput) GetAudiences() []AudienceSummary {
	if o == nil {
		var ret []AudienceSummary
		return ret
	}

	return o.Audiences
}

// GetAudiencesOk returns a tuple with the Audiences field value
// and a boolean to check if the value has been set.
func (o *ListAudiencesAlphaOutput) GetAudiencesOk() ([]AudienceSummary, bool) {
	if o == nil {
		return nil, false
	}
	return o.Audiences, true
}

// SetAudiences sets field value
func (o *ListAudiencesAlphaOutput) SetAudiences(v []AudienceSummary) {
	o.Audiences = v
}

// GetPagination returns the Pagination field value
func (o *ListAudiencesAlphaOutput) GetPagination() PaginationOutput {
	if o == nil {
		var ret PaginationOutput
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *ListAudiencesAlphaOutput) GetPaginationOk() (*PaginationOutput, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *ListAudiencesAlphaOutput) SetPagination(v PaginationOutput) {
	o.Pagination = v
}

func (o ListAudiencesAlphaOutput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListAudiencesAlphaOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["audiences"] = o.Audiences
	toSerialize["pagination"] = o.Pagination
	return toSerialize, nil
}

type NullableListAudiencesAlphaOutput struct {
	value *ListAudiencesAlphaOutput
	isSet bool
}

func (v NullableListAudiencesAlphaOutput) Get() *ListAudiencesAlphaOutput {
	return v.value
}

func (v *NullableListAudiencesAlphaOutput) Set(val *ListAudiencesAlphaOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableListAudiencesAlphaOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableListAudiencesAlphaOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListAudiencesAlphaOutput(
	val *ListAudiencesAlphaOutput,
) *NullableListAudiencesAlphaOutput {
	return &NullableListAudiencesAlphaOutput{value: val, isSet: true}
}

func (v NullableListAudiencesAlphaOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListAudiencesAlphaOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
