/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 58.13.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the ListAudiencesPaginationInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListAudiencesPaginationInput{}

// ListAudiencesPaginationInput Fork of Autobahn's PaginationInput. Count is limited to 200 in this fork.  Pagination parameters.  Every resource that returns a list of items in its `Output` object may contain a `PaginationInput` in its `Input` object. Required, though some of its fields are optional.
type ListAudiencesPaginationInput struct {
	// The page to request.  Acceptable values to use here are in PaginationOutput objects, in the `current`, `next`, and `previous` keys.  Consumers of the API must treat this value as opaque.
	Cursor *string `json:"cursor,omitempty"`
	// The number of items to retrieve in a page, between 1 and 200.
	Count float32 `json:"count"`
}

// NewListAudiencesPaginationInput instantiates a new ListAudiencesPaginationInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListAudiencesPaginationInput(count float32) *ListAudiencesPaginationInput {
	this := ListAudiencesPaginationInput{}
	this.Count = count
	return &this
}

// NewListAudiencesPaginationInputWithDefaults instantiates a new ListAudiencesPaginationInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListAudiencesPaginationInputWithDefaults() *ListAudiencesPaginationInput {
	this := ListAudiencesPaginationInput{}
	return &this
}

// GetCursor returns the Cursor field value if set, zero value otherwise.
func (o *ListAudiencesPaginationInput) GetCursor() string {
	if o == nil || IsNil(o.Cursor) {
		var ret string
		return ret
	}
	return *o.Cursor
}

// GetCursorOk returns a tuple with the Cursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAudiencesPaginationInput) GetCursorOk() (*string, bool) {
	if o == nil || IsNil(o.Cursor) {
		return nil, false
	}
	return o.Cursor, true
}

// HasCursor returns a boolean if a field has been set.
func (o *ListAudiencesPaginationInput) HasCursor() bool {
	if o != nil && !IsNil(o.Cursor) {
		return true
	}

	return false
}

// SetCursor gets a reference to the given string and assigns it to the Cursor field.
func (o *ListAudiencesPaginationInput) SetCursor(v string) {
	o.Cursor = &v
}

// GetCount returns the Count field value
func (o *ListAudiencesPaginationInput) GetCount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *ListAudiencesPaginationInput) GetCountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *ListAudiencesPaginationInput) SetCount(v float32) {
	o.Count = v
}

func (o ListAudiencesPaginationInput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListAudiencesPaginationInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Cursor) {
		toSerialize["cursor"] = o.Cursor
	}
	toSerialize["count"] = o.Count
	return toSerialize, nil
}

type NullableListAudiencesPaginationInput struct {
	value *ListAudiencesPaginationInput
	isSet bool
}

func (v NullableListAudiencesPaginationInput) Get() *ListAudiencesPaginationInput {
	return v.value
}

func (v *NullableListAudiencesPaginationInput) Set(val *ListAudiencesPaginationInput) {
	v.value = val
	v.isSet = true
}

func (v NullableListAudiencesPaginationInput) IsSet() bool {
	return v.isSet
}

func (v *NullableListAudiencesPaginationInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListAudiencesPaginationInput(
	val *ListAudiencesPaginationInput,
) *NullableListAudiencesPaginationInput {
	return &NullableListAudiencesPaginationInput{value: val, isSet: true}
}

func (v NullableListAudiencesPaginationInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListAudiencesPaginationInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
