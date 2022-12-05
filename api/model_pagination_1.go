/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 33.0.3
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// Pagination1 Pagination options.
type Pagination1 struct {
	// The page to request.  Acceptable values to use here are in PaginationOutput objects, in the `current`, `next`, and `previous` keys.  Consumers of the API must treat this value as opaque.
	Cursor *string `json:"cursor,omitempty"`
	// The number of items to retrieve in a page, between 1 and 200.
	Count float32 `json:"count"`
}

// NewPagination1 instantiates a new Pagination1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPagination1(count float32) *Pagination1 {
	this := Pagination1{}
	this.Count = count
	return &this
}

// NewPagination1WithDefaults instantiates a new Pagination1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPagination1WithDefaults() *Pagination1 {
	this := Pagination1{}
	return &this
}

// GetCursor returns the Cursor field value if set, zero value otherwise.
func (o *Pagination1) GetCursor() string {
	if o == nil || o.Cursor == nil {
		var ret string
		return ret
	}
	return *o.Cursor
}

// GetCursorOk returns a tuple with the Cursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pagination1) GetCursorOk() (*string, bool) {
	if o == nil || o.Cursor == nil {
		return nil, false
	}
	return o.Cursor, true
}

// HasCursor returns a boolean if a field has been set.
func (o *Pagination1) HasCursor() bool {
	if o != nil && o.Cursor != nil {
		return true
	}

	return false
}

// SetCursor gets a reference to the given string and assigns it to the Cursor field.
func (o *Pagination1) SetCursor(v string) {
	o.Cursor = &v
}

// GetCount returns the Count field value
func (o *Pagination1) GetCount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *Pagination1) GetCountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *Pagination1) SetCount(v float32) {
	o.Count = v
}

func (o Pagination1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Cursor != nil {
		toSerialize["cursor"] = o.Cursor
	}
	if true {
		toSerialize["count"] = o.Count
	}
	return json.Marshal(toSerialize)
}

type NullablePagination1 struct {
	value *Pagination1
	isSet bool
}

func (v NullablePagination1) Get() *Pagination1 {
	return v.value
}

func (v *NullablePagination1) Set(val *Pagination1) {
	v.value = val
	v.isSet = true
}

func (v NullablePagination1) IsSet() bool {
	return v.isSet
}

func (v *NullablePagination1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePagination1(val *Pagination1) *NullablePagination1 {
	return &NullablePagination1{value: val, isSet: true}
}

func (v NullablePagination1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePagination1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
