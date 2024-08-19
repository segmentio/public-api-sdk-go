/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 53.1.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the ListSpacesAlphaOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListSpacesAlphaOutput{}

// ListSpacesAlphaOutput Response for the list spaces endpoint.
type ListSpacesAlphaOutput struct {
	// A list of spaces.
	Spaces     []Space          `json:"spaces"`
	Pagination PaginationOutput `json:"pagination"`
}

// NewListSpacesAlphaOutput instantiates a new ListSpacesAlphaOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListSpacesAlphaOutput(spaces []Space, pagination PaginationOutput) *ListSpacesAlphaOutput {
	this := ListSpacesAlphaOutput{}
	this.Spaces = spaces
	this.Pagination = pagination
	return &this
}

// NewListSpacesAlphaOutputWithDefaults instantiates a new ListSpacesAlphaOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListSpacesAlphaOutputWithDefaults() *ListSpacesAlphaOutput {
	this := ListSpacesAlphaOutput{}
	return &this
}

// GetSpaces returns the Spaces field value
func (o *ListSpacesAlphaOutput) GetSpaces() []Space {
	if o == nil {
		var ret []Space
		return ret
	}

	return o.Spaces
}

// GetSpacesOk returns a tuple with the Spaces field value
// and a boolean to check if the value has been set.
func (o *ListSpacesAlphaOutput) GetSpacesOk() ([]Space, bool) {
	if o == nil {
		return nil, false
	}
	return o.Spaces, true
}

// SetSpaces sets field value
func (o *ListSpacesAlphaOutput) SetSpaces(v []Space) {
	o.Spaces = v
}

// GetPagination returns the Pagination field value
func (o *ListSpacesAlphaOutput) GetPagination() PaginationOutput {
	if o == nil {
		var ret PaginationOutput
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *ListSpacesAlphaOutput) GetPaginationOk() (*PaginationOutput, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *ListSpacesAlphaOutput) SetPagination(v PaginationOutput) {
	o.Pagination = v
}

func (o ListSpacesAlphaOutput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListSpacesAlphaOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["spaces"] = o.Spaces
	toSerialize["pagination"] = o.Pagination
	return toSerialize, nil
}

type NullableListSpacesAlphaOutput struct {
	value *ListSpacesAlphaOutput
	isSet bool
}

func (v NullableListSpacesAlphaOutput) Get() *ListSpacesAlphaOutput {
	return v.value
}

func (v *NullableListSpacesAlphaOutput) Set(val *ListSpacesAlphaOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableListSpacesAlphaOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableListSpacesAlphaOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListSpacesAlphaOutput(val *ListSpacesAlphaOutput) *NullableListSpacesAlphaOutput {
	return &NullableListSpacesAlphaOutput{value: val, isSet: true}
}

func (v NullableListSpacesAlphaOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListSpacesAlphaOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
