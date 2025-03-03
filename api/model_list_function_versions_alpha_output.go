/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 57.4.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the ListFunctionVersionsAlphaOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListFunctionVersionsAlphaOutput{}

// ListFunctionVersionsAlphaOutput Lists Versions of a Function.
type ListFunctionVersionsAlphaOutput struct {
	// An array of Functions.
	Versions   []Version        `json:"versions"`
	Pagination PaginationOutput `json:"pagination"`
}

// NewListFunctionVersionsAlphaOutput instantiates a new ListFunctionVersionsAlphaOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListFunctionVersionsAlphaOutput(
	versions []Version,
	pagination PaginationOutput,
) *ListFunctionVersionsAlphaOutput {
	this := ListFunctionVersionsAlphaOutput{}
	this.Versions = versions
	this.Pagination = pagination
	return &this
}

// NewListFunctionVersionsAlphaOutputWithDefaults instantiates a new ListFunctionVersionsAlphaOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListFunctionVersionsAlphaOutputWithDefaults() *ListFunctionVersionsAlphaOutput {
	this := ListFunctionVersionsAlphaOutput{}
	return &this
}

// GetVersions returns the Versions field value
func (o *ListFunctionVersionsAlphaOutput) GetVersions() []Version {
	if o == nil {
		var ret []Version
		return ret
	}

	return o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value
// and a boolean to check if the value has been set.
func (o *ListFunctionVersionsAlphaOutput) GetVersionsOk() ([]Version, bool) {
	if o == nil {
		return nil, false
	}
	return o.Versions, true
}

// SetVersions sets field value
func (o *ListFunctionVersionsAlphaOutput) SetVersions(v []Version) {
	o.Versions = v
}

// GetPagination returns the Pagination field value
func (o *ListFunctionVersionsAlphaOutput) GetPagination() PaginationOutput {
	if o == nil {
		var ret PaginationOutput
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *ListFunctionVersionsAlphaOutput) GetPaginationOk() (*PaginationOutput, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *ListFunctionVersionsAlphaOutput) SetPagination(v PaginationOutput) {
	o.Pagination = v
}

func (o ListFunctionVersionsAlphaOutput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListFunctionVersionsAlphaOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["versions"] = o.Versions
	toSerialize["pagination"] = o.Pagination
	return toSerialize, nil
}

type NullableListFunctionVersionsAlphaOutput struct {
	value *ListFunctionVersionsAlphaOutput
	isSet bool
}

func (v NullableListFunctionVersionsAlphaOutput) Get() *ListFunctionVersionsAlphaOutput {
	return v.value
}

func (v *NullableListFunctionVersionsAlphaOutput) Set(val *ListFunctionVersionsAlphaOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableListFunctionVersionsAlphaOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableListFunctionVersionsAlphaOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListFunctionVersionsAlphaOutput(
	val *ListFunctionVersionsAlphaOutput,
) *NullableListFunctionVersionsAlphaOutput {
	return &NullableListFunctionVersionsAlphaOutput{value: val, isSet: true}
}

func (v NullableListFunctionVersionsAlphaOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListFunctionVersionsAlphaOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
