/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API. 

API version: 33.0.1
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// ListSuppressionsV1Output The output of a list suppressed call for a Workspace.
type ListSuppressionsV1Output struct {
	// An array that lists the suppressions from the Workspace.  Config API note: equal to `attributes`.
	Suppressed []SuppressedInner `json:"suppressed"`
	Pagination Pagination `json:"pagination"`
}

// NewListSuppressionsV1Output instantiates a new ListSuppressionsV1Output object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListSuppressionsV1Output(suppressed []SuppressedInner, pagination Pagination) *ListSuppressionsV1Output {
	this := ListSuppressionsV1Output{}
	this.Suppressed = suppressed
	this.Pagination = pagination
	return &this
}

// NewListSuppressionsV1OutputWithDefaults instantiates a new ListSuppressionsV1Output object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListSuppressionsV1OutputWithDefaults() *ListSuppressionsV1Output {
	this := ListSuppressionsV1Output{}
	return &this
}

// GetSuppressed returns the Suppressed field value
func (o *ListSuppressionsV1Output) GetSuppressed() []SuppressedInner {
	if o == nil {
		var ret []SuppressedInner
		return ret
	}

	return o.Suppressed
}

// GetSuppressedOk returns a tuple with the Suppressed field value
// and a boolean to check if the value has been set.
func (o *ListSuppressionsV1Output) GetSuppressedOk() ([]SuppressedInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Suppressed, true
}

// SetSuppressed sets field value
func (o *ListSuppressionsV1Output) SetSuppressed(v []SuppressedInner) {
	o.Suppressed = v
}

// GetPagination returns the Pagination field value
func (o *ListSuppressionsV1Output) GetPagination() Pagination {
	if o == nil {
		var ret Pagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *ListSuppressionsV1Output) GetPaginationOk() (*Pagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *ListSuppressionsV1Output) SetPagination(v Pagination) {
	o.Pagination = v
}

func (o ListSuppressionsV1Output) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["suppressed"] = o.Suppressed
	}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	return json.Marshal(toSerialize)
}

type NullableListSuppressionsV1Output struct {
	value *ListSuppressionsV1Output
	isSet bool
}

func (v NullableListSuppressionsV1Output) Get() *ListSuppressionsV1Output {
	return v.value
}

func (v *NullableListSuppressionsV1Output) Set(val *ListSuppressionsV1Output) {
	v.value = val
	v.isSet = true
}

func (v NullableListSuppressionsV1Output) IsSet() bool {
	return v.isSet
}

func (v *NullableListSuppressionsV1Output) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListSuppressionsV1Output(val *ListSuppressionsV1Output) *NullableListSuppressionsV1Output {
	return &NullableListSuppressionsV1Output{value: val, isSet: true}
}

func (v NullableListSuppressionsV1Output) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListSuppressionsV1Output) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


