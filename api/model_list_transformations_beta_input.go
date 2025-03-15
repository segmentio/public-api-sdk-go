/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 58.0.1
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the ListTransformationsBetaInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListTransformationsBetaInput{}

// ListTransformationsBetaInput Lists the Transformations associated with the current Workspace.
type ListTransformationsBetaInput struct {
	Pagination *PaginationInput `json:"pagination,omitempty"`
}

// NewListTransformationsBetaInput instantiates a new ListTransformationsBetaInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListTransformationsBetaInput() *ListTransformationsBetaInput {
	this := ListTransformationsBetaInput{}
	return &this
}

// NewListTransformationsBetaInputWithDefaults instantiates a new ListTransformationsBetaInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListTransformationsBetaInputWithDefaults() *ListTransformationsBetaInput {
	this := ListTransformationsBetaInput{}
	return &this
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *ListTransformationsBetaInput) GetPagination() PaginationInput {
	if o == nil || IsNil(o.Pagination) {
		var ret PaginationInput
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListTransformationsBetaInput) GetPaginationOk() (*PaginationInput, bool) {
	if o == nil || IsNil(o.Pagination) {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *ListTransformationsBetaInput) HasPagination() bool {
	if o != nil && !IsNil(o.Pagination) {
		return true
	}

	return false
}

// SetPagination gets a reference to the given PaginationInput and assigns it to the Pagination field.
func (o *ListTransformationsBetaInput) SetPagination(v PaginationInput) {
	o.Pagination = &v
}

func (o ListTransformationsBetaInput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListTransformationsBetaInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Pagination) {
		toSerialize["pagination"] = o.Pagination
	}
	return toSerialize, nil
}

type NullableListTransformationsBetaInput struct {
	value *ListTransformationsBetaInput
	isSet bool
}

func (v NullableListTransformationsBetaInput) Get() *ListTransformationsBetaInput {
	return v.value
}

func (v *NullableListTransformationsBetaInput) Set(val *ListTransformationsBetaInput) {
	v.value = val
	v.isSet = true
}

func (v NullableListTransformationsBetaInput) IsSet() bool {
	return v.isSet
}

func (v *NullableListTransformationsBetaInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListTransformationsBetaInput(
	val *ListTransformationsBetaInput,
) *NullableListTransformationsBetaInput {
	return &NullableListTransformationsBetaInput{value: val, isSet: true}
}

func (v NullableListTransformationsBetaInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListTransformationsBetaInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
