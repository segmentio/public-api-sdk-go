/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 58.1.1
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the ListAudienceConsumersFromSpaceAndAudienceOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListAudienceConsumersFromSpaceAndAudienceOutput{}

// ListAudienceConsumersFromSpaceAndAudienceOutput List Audience consumers output.
type ListAudienceConsumersFromSpaceAndAudienceOutput struct {
	// The list of audience consumers.
	Audiences  []AudienceSummary `json:"audiences"`
	Pagination PaginationOutput  `json:"pagination"`
}

// NewListAudienceConsumersFromSpaceAndAudienceOutput instantiates a new ListAudienceConsumersFromSpaceAndAudienceOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListAudienceConsumersFromSpaceAndAudienceOutput(
	audiences []AudienceSummary,
	pagination PaginationOutput,
) *ListAudienceConsumersFromSpaceAndAudienceOutput {
	this := ListAudienceConsumersFromSpaceAndAudienceOutput{}
	this.Audiences = audiences
	this.Pagination = pagination
	return &this
}

// NewListAudienceConsumersFromSpaceAndAudienceOutputWithDefaults instantiates a new ListAudienceConsumersFromSpaceAndAudienceOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListAudienceConsumersFromSpaceAndAudienceOutputWithDefaults() *ListAudienceConsumersFromSpaceAndAudienceOutput {
	this := ListAudienceConsumersFromSpaceAndAudienceOutput{}
	return &this
}

// GetAudiences returns the Audiences field value
func (o *ListAudienceConsumersFromSpaceAndAudienceOutput) GetAudiences() []AudienceSummary {
	if o == nil {
		var ret []AudienceSummary
		return ret
	}

	return o.Audiences
}

// GetAudiencesOk returns a tuple with the Audiences field value
// and a boolean to check if the value has been set.
func (o *ListAudienceConsumersFromSpaceAndAudienceOutput) GetAudiencesOk() ([]AudienceSummary, bool) {
	if o == nil {
		return nil, false
	}
	return o.Audiences, true
}

// SetAudiences sets field value
func (o *ListAudienceConsumersFromSpaceAndAudienceOutput) SetAudiences(v []AudienceSummary) {
	o.Audiences = v
}

// GetPagination returns the Pagination field value
func (o *ListAudienceConsumersFromSpaceAndAudienceOutput) GetPagination() PaginationOutput {
	if o == nil {
		var ret PaginationOutput
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *ListAudienceConsumersFromSpaceAndAudienceOutput) GetPaginationOk() (*PaginationOutput, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *ListAudienceConsumersFromSpaceAndAudienceOutput) SetPagination(v PaginationOutput) {
	o.Pagination = v
}

func (o ListAudienceConsumersFromSpaceAndAudienceOutput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListAudienceConsumersFromSpaceAndAudienceOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["audiences"] = o.Audiences
	toSerialize["pagination"] = o.Pagination
	return toSerialize, nil
}

type NullableListAudienceConsumersFromSpaceAndAudienceOutput struct {
	value *ListAudienceConsumersFromSpaceAndAudienceOutput
	isSet bool
}

func (v NullableListAudienceConsumersFromSpaceAndAudienceOutput) Get() *ListAudienceConsumersFromSpaceAndAudienceOutput {
	return v.value
}

func (v *NullableListAudienceConsumersFromSpaceAndAudienceOutput) Set(
	val *ListAudienceConsumersFromSpaceAndAudienceOutput,
) {
	v.value = val
	v.isSet = true
}

func (v NullableListAudienceConsumersFromSpaceAndAudienceOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableListAudienceConsumersFromSpaceAndAudienceOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListAudienceConsumersFromSpaceAndAudienceOutput(
	val *ListAudienceConsumersFromSpaceAndAudienceOutput,
) *NullableListAudienceConsumersFromSpaceAndAudienceOutput {
	return &NullableListAudienceConsumersFromSpaceAndAudienceOutput{value: val, isSet: true}
}

func (v NullableListAudienceConsumersFromSpaceAndAudienceOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListAudienceConsumersFromSpaceAndAudienceOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
