/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 57.3.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the ListFiltersPaginationInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListFiltersPaginationInput{}

// ListFiltersPaginationInput Pagination for list filters.
type ListFiltersPaginationInput struct {
	// The offset.
	Cursor *float32 `json:"cursor,omitempty"`
	// The number of items to retrieve in a page, between 1 and 10.
	Count float32 `json:"count"`
}

// NewListFiltersPaginationInput instantiates a new ListFiltersPaginationInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListFiltersPaginationInput(count float32) *ListFiltersPaginationInput {
	this := ListFiltersPaginationInput{}
	this.Count = count
	return &this
}

// NewListFiltersPaginationInputWithDefaults instantiates a new ListFiltersPaginationInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListFiltersPaginationInputWithDefaults() *ListFiltersPaginationInput {
	this := ListFiltersPaginationInput{}
	return &this
}

// GetCursor returns the Cursor field value if set, zero value otherwise.
func (o *ListFiltersPaginationInput) GetCursor() float32 {
	if o == nil || IsNil(o.Cursor) {
		var ret float32
		return ret
	}
	return *o.Cursor
}

// GetCursorOk returns a tuple with the Cursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListFiltersPaginationInput) GetCursorOk() (*float32, bool) {
	if o == nil || IsNil(o.Cursor) {
		return nil, false
	}
	return o.Cursor, true
}

// HasCursor returns a boolean if a field has been set.
func (o *ListFiltersPaginationInput) HasCursor() bool {
	if o != nil && !IsNil(o.Cursor) {
		return true
	}

	return false
}

// SetCursor gets a reference to the given float32 and assigns it to the Cursor field.
func (o *ListFiltersPaginationInput) SetCursor(v float32) {
	o.Cursor = &v
}

// GetCount returns the Count field value
func (o *ListFiltersPaginationInput) GetCount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *ListFiltersPaginationInput) GetCountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *ListFiltersPaginationInput) SetCount(v float32) {
	o.Count = v
}

func (o ListFiltersPaginationInput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListFiltersPaginationInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Cursor) {
		toSerialize["cursor"] = o.Cursor
	}
	toSerialize["count"] = o.Count
	return toSerialize, nil
}

type NullableListFiltersPaginationInput struct {
	value *ListFiltersPaginationInput
	isSet bool
}

func (v NullableListFiltersPaginationInput) Get() *ListFiltersPaginationInput {
	return v.value
}

func (v *NullableListFiltersPaginationInput) Set(val *ListFiltersPaginationInput) {
	v.value = val
	v.isSet = true
}

func (v NullableListFiltersPaginationInput) IsSet() bool {
	return v.isSet
}

func (v *NullableListFiltersPaginationInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListFiltersPaginationInput(
	val *ListFiltersPaginationInput,
) *NullableListFiltersPaginationInput {
	return &NullableListFiltersPaginationInput{value: val, isSet: true}
}

func (v NullableListFiltersPaginationInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListFiltersPaginationInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
