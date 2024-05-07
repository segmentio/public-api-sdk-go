/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 49.2.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the TraitCreateOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TraitCreateOptions{}

// TraitCreateOptions struct for TraitCreateOptions
type TraitCreateOptions struct {
	IncludeHistoricalData *bool `json:"includeHistoricalData,omitempty"`
	IncludeAnonymousUsers *bool `json:"includeAnonymousUsers,omitempty"`
}

// NewTraitCreateOptions instantiates a new TraitCreateOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTraitCreateOptions() *TraitCreateOptions {
	this := TraitCreateOptions{}
	return &this
}

// NewTraitCreateOptionsWithDefaults instantiates a new TraitCreateOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTraitCreateOptionsWithDefaults() *TraitCreateOptions {
	this := TraitCreateOptions{}
	return &this
}

// GetIncludeHistoricalData returns the IncludeHistoricalData field value if set, zero value otherwise.
func (o *TraitCreateOptions) GetIncludeHistoricalData() bool {
	if o == nil || IsNil(o.IncludeHistoricalData) {
		var ret bool
		return ret
	}
	return *o.IncludeHistoricalData
}

// GetIncludeHistoricalDataOk returns a tuple with the IncludeHistoricalData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TraitCreateOptions) GetIncludeHistoricalDataOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeHistoricalData) {
		return nil, false
	}
	return o.IncludeHistoricalData, true
}

// HasIncludeHistoricalData returns a boolean if a field has been set.
func (o *TraitCreateOptions) HasIncludeHistoricalData() bool {
	if o != nil && !IsNil(o.IncludeHistoricalData) {
		return true
	}

	return false
}

// SetIncludeHistoricalData gets a reference to the given bool and assigns it to the IncludeHistoricalData field.
func (o *TraitCreateOptions) SetIncludeHistoricalData(v bool) {
	o.IncludeHistoricalData = &v
}

// GetIncludeAnonymousUsers returns the IncludeAnonymousUsers field value if set, zero value otherwise.
func (o *TraitCreateOptions) GetIncludeAnonymousUsers() bool {
	if o == nil || IsNil(o.IncludeAnonymousUsers) {
		var ret bool
		return ret
	}
	return *o.IncludeAnonymousUsers
}

// GetIncludeAnonymousUsersOk returns a tuple with the IncludeAnonymousUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TraitCreateOptions) GetIncludeAnonymousUsersOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeAnonymousUsers) {
		return nil, false
	}
	return o.IncludeAnonymousUsers, true
}

// HasIncludeAnonymousUsers returns a boolean if a field has been set.
func (o *TraitCreateOptions) HasIncludeAnonymousUsers() bool {
	if o != nil && !IsNil(o.IncludeAnonymousUsers) {
		return true
	}

	return false
}

// SetIncludeAnonymousUsers gets a reference to the given bool and assigns it to the IncludeAnonymousUsers field.
func (o *TraitCreateOptions) SetIncludeAnonymousUsers(v bool) {
	o.IncludeAnonymousUsers = &v
}

func (o TraitCreateOptions) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TraitCreateOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IncludeHistoricalData) {
		toSerialize["includeHistoricalData"] = o.IncludeHistoricalData
	}
	if !IsNil(o.IncludeAnonymousUsers) {
		toSerialize["includeAnonymousUsers"] = o.IncludeAnonymousUsers
	}
	return toSerialize, nil
}

type NullableTraitCreateOptions struct {
	value *TraitCreateOptions
	isSet bool
}

func (v NullableTraitCreateOptions) Get() *TraitCreateOptions {
	return v.value
}

func (v *NullableTraitCreateOptions) Set(val *TraitCreateOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableTraitCreateOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableTraitCreateOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTraitCreateOptions(val *TraitCreateOptions) *NullableTraitCreateOptions {
	return &NullableTraitCreateOptions{value: val, isSet: true}
}

func (v NullableTraitCreateOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTraitCreateOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
