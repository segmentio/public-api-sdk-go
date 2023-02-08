/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 34.1.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PaginationInput Pagination parameters.  Every resource that returns a list of items in its `Output` object may contain a `PaginationInput` in its `Input` object. Required, though some of its fields are optional.
type PaginationInput struct {
	// The page to request.  Acceptable values to use here are in PaginationOutput objects, in the `current`, `next`, and `previous` keys.  Consumers of the API must treat this value as opaque.
	Cursor *string `json:"cursor,omitempty"`
	// The number of items to retrieve in a page, between 1 and 200.
	Count float32 `json:"count"`
}

// NewPaginationInput instantiates a new PaginationInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginationInput(count float32) *PaginationInput {
	this := PaginationInput{}
	this.Count = count
	return &this
}

// NewPaginationInputWithDefaults instantiates a new PaginationInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginationInputWithDefaults() *PaginationInput {
	this := PaginationInput{}
	return &this
}

// GetCursor returns the Cursor field value if set, zero value otherwise.
func (o *PaginationInput) GetCursor() string {
	if o == nil || o.Cursor == nil {
		var ret string
		return ret
	}
	return *o.Cursor
}

// GetCursorOk returns a tuple with the Cursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginationInput) GetCursorOk() (*string, bool) {
	if o == nil || o.Cursor == nil {
		return nil, false
	}
	return o.Cursor, true
}

// HasCursor returns a boolean if a field has been set.
func (o *PaginationInput) HasCursor() bool {
	if o != nil && o.Cursor != nil {
		return true
	}

	return false
}

// SetCursor gets a reference to the given string and assigns it to the Cursor field.
func (o *PaginationInput) SetCursor(v string) {
	o.Cursor = &v
}

// GetCount returns the Count field value
func (o *PaginationInput) GetCount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *PaginationInput) GetCountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *PaginationInput) SetCount(v float32) {
	o.Count = v
}

func (o PaginationInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Cursor != nil {
		toSerialize["cursor"] = o.Cursor
	}
	if true {
		toSerialize["count"] = o.Count
	}
	return json.Marshal(toSerialize)
}

type NullablePaginationInput struct {
	value *PaginationInput
	isSet bool
}

func (v NullablePaginationInput) Get() *PaginationInput {
	return v.value
}

func (v *NullablePaginationInput) Set(val *PaginationInput) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginationInput) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginationInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginationInput(val *PaginationInput) *NullablePaginationInput {
	return &NullablePaginationInput{value: val, isSet: true}
}

func (v NullablePaginationInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginationInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
