/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 57.0.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the PaginationOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaginationOutput{}

// PaginationOutput Pagination metadata for a list response.  Responses return this object alongside a list of resources, which provides the necessary metadata for manipulating a paginated collection. In operations that return lists, it's always present, though some of its fields might not be.
type PaginationOutput struct {
	// The current cursor within a collection.  Consumers of the API must treat this value as opaque.
	Current string `json:"current"`
	// A pointer to the next page.  This does not return when you retrieve the last page.  Consumers of the API must treat this value as opaque.
	Next NullableString `json:"next,omitempty"`
	// A pointer to the previous page.  This does not return when you retrieve the first page.  Consumers of the API must treat this value as opaque.
	Previous NullableString `json:"previous,omitempty"`
	// The total number of entries available in the collection.  If calculating it impacts performance, the response may omit this field.
	TotalEntries *float32 `json:"totalEntries,omitempty"`
}

// NewPaginationOutput instantiates a new PaginationOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginationOutput(current string) *PaginationOutput {
	this := PaginationOutput{}
	this.Current = current
	return &this
}

// NewPaginationOutputWithDefaults instantiates a new PaginationOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginationOutputWithDefaults() *PaginationOutput {
	this := PaginationOutput{}
	return &this
}

// GetCurrent returns the Current field value
func (o *PaginationOutput) GetCurrent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Current
}

// GetCurrentOk returns a tuple with the Current field value
// and a boolean to check if the value has been set.
func (o *PaginationOutput) GetCurrentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Current, true
}

// SetCurrent sets field value
func (o *PaginationOutput) SetCurrent(v string) {
	o.Current = v
}

// GetNext returns the Next field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaginationOutput) GetNext() string {
	if o == nil || IsNil(o.Next.Get()) {
		var ret string
		return ret
	}
	return *o.Next.Get()
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaginationOutput) GetNextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Next.Get(), o.Next.IsSet()
}

// HasNext returns a boolean if a field has been set.
func (o *PaginationOutput) HasNext() bool {
	if o != nil && o.Next.IsSet() {
		return true
	}

	return false
}

// SetNext gets a reference to the given NullableString and assigns it to the Next field.
func (o *PaginationOutput) SetNext(v string) {
	o.Next.Set(&v)
}

// SetNextNil sets the value for Next to be an explicit nil
func (o *PaginationOutput) SetNextNil() {
	o.Next.Set(nil)
}

// UnsetNext ensures that no value is present for Next, not even an explicit nil
func (o *PaginationOutput) UnsetNext() {
	o.Next.Unset()
}

// GetPrevious returns the Previous field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaginationOutput) GetPrevious() string {
	if o == nil || IsNil(o.Previous.Get()) {
		var ret string
		return ret
	}
	return *o.Previous.Get()
}

// GetPreviousOk returns a tuple with the Previous field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaginationOutput) GetPreviousOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Previous.Get(), o.Previous.IsSet()
}

// HasPrevious returns a boolean if a field has been set.
func (o *PaginationOutput) HasPrevious() bool {
	if o != nil && o.Previous.IsSet() {
		return true
	}

	return false
}

// SetPrevious gets a reference to the given NullableString and assigns it to the Previous field.
func (o *PaginationOutput) SetPrevious(v string) {
	o.Previous.Set(&v)
}

// SetPreviousNil sets the value for Previous to be an explicit nil
func (o *PaginationOutput) SetPreviousNil() {
	o.Previous.Set(nil)
}

// UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
func (o *PaginationOutput) UnsetPrevious() {
	o.Previous.Unset()
}

// GetTotalEntries returns the TotalEntries field value if set, zero value otherwise.
func (o *PaginationOutput) GetTotalEntries() float32 {
	if o == nil || IsNil(o.TotalEntries) {
		var ret float32
		return ret
	}
	return *o.TotalEntries
}

// GetTotalEntriesOk returns a tuple with the TotalEntries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginationOutput) GetTotalEntriesOk() (*float32, bool) {
	if o == nil || IsNil(o.TotalEntries) {
		return nil, false
	}
	return o.TotalEntries, true
}

// HasTotalEntries returns a boolean if a field has been set.
func (o *PaginationOutput) HasTotalEntries() bool {
	if o != nil && !IsNil(o.TotalEntries) {
		return true
	}

	return false
}

// SetTotalEntries gets a reference to the given float32 and assigns it to the TotalEntries field.
func (o *PaginationOutput) SetTotalEntries(v float32) {
	o.TotalEntries = &v
}

func (o PaginationOutput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaginationOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["current"] = o.Current
	if o.Next.IsSet() {
		toSerialize["next"] = o.Next.Get()
	}
	if o.Previous.IsSet() {
		toSerialize["previous"] = o.Previous.Get()
	}
	if !IsNil(o.TotalEntries) {
		toSerialize["totalEntries"] = o.TotalEntries
	}
	return toSerialize, nil
}

type NullablePaginationOutput struct {
	value *PaginationOutput
	isSet bool
}

func (v NullablePaginationOutput) Get() *PaginationOutput {
	return v.value
}

func (v *NullablePaginationOutput) Set(val *PaginationOutput) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginationOutput) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginationOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginationOutput(val *PaginationOutput) *NullablePaginationOutput {
	return &NullablePaginationOutput{value: val, isSet: true}
}

func (v NullablePaginationOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginationOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
