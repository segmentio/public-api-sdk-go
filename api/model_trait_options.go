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

// checks if the TraitOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TraitOptions{}

// TraitOptions struct for TraitOptions
type TraitOptions struct {
	// Determines whether data prior to the computed trait being created is included when determining the computed trait value. Note that including historical data may be needed in order to properly handle the definition specified. In these cases, Segment will automatically handle including historical data and the response will return the includeHistoricalData parameter as true.
	IncludeHistoricalData *bool `json:"includeHistoricalData,omitempty"`
	// Determines whether anonymous users should be included when determining the computed trait value.
	IncludeAnonymousUsers *bool `json:"includeAnonymousUsers,omitempty"`
}

// NewTraitOptions instantiates a new TraitOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTraitOptions() *TraitOptions {
	this := TraitOptions{}
	return &this
}

// NewTraitOptionsWithDefaults instantiates a new TraitOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTraitOptionsWithDefaults() *TraitOptions {
	this := TraitOptions{}
	return &this
}

// GetIncludeHistoricalData returns the IncludeHistoricalData field value if set, zero value otherwise.
func (o *TraitOptions) GetIncludeHistoricalData() bool {
	if o == nil || IsNil(o.IncludeHistoricalData) {
		var ret bool
		return ret
	}
	return *o.IncludeHistoricalData
}

// GetIncludeHistoricalDataOk returns a tuple with the IncludeHistoricalData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TraitOptions) GetIncludeHistoricalDataOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeHistoricalData) {
		return nil, false
	}
	return o.IncludeHistoricalData, true
}

// HasIncludeHistoricalData returns a boolean if a field has been set.
func (o *TraitOptions) HasIncludeHistoricalData() bool {
	if o != nil && !IsNil(o.IncludeHistoricalData) {
		return true
	}

	return false
}

// SetIncludeHistoricalData gets a reference to the given bool and assigns it to the IncludeHistoricalData field.
func (o *TraitOptions) SetIncludeHistoricalData(v bool) {
	o.IncludeHistoricalData = &v
}

// GetIncludeAnonymousUsers returns the IncludeAnonymousUsers field value if set, zero value otherwise.
func (o *TraitOptions) GetIncludeAnonymousUsers() bool {
	if o == nil || IsNil(o.IncludeAnonymousUsers) {
		var ret bool
		return ret
	}
	return *o.IncludeAnonymousUsers
}

// GetIncludeAnonymousUsersOk returns a tuple with the IncludeAnonymousUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TraitOptions) GetIncludeAnonymousUsersOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeAnonymousUsers) {
		return nil, false
	}
	return o.IncludeAnonymousUsers, true
}

// HasIncludeAnonymousUsers returns a boolean if a field has been set.
func (o *TraitOptions) HasIncludeAnonymousUsers() bool {
	if o != nil && !IsNil(o.IncludeAnonymousUsers) {
		return true
	}

	return false
}

// SetIncludeAnonymousUsers gets a reference to the given bool and assigns it to the IncludeAnonymousUsers field.
func (o *TraitOptions) SetIncludeAnonymousUsers(v bool) {
	o.IncludeAnonymousUsers = &v
}

func (o TraitOptions) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TraitOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IncludeHistoricalData) {
		toSerialize["includeHistoricalData"] = o.IncludeHistoricalData
	}
	if !IsNil(o.IncludeAnonymousUsers) {
		toSerialize["includeAnonymousUsers"] = o.IncludeAnonymousUsers
	}
	return toSerialize, nil
}

type NullableTraitOptions struct {
	value *TraitOptions
	isSet bool
}

func (v NullableTraitOptions) Get() *TraitOptions {
	return v.value
}

func (v *NullableTraitOptions) Set(val *TraitOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableTraitOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableTraitOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTraitOptions(val *TraitOptions) *NullableTraitOptions {
	return &NullableTraitOptions{value: val, isSet: true}
}

func (v NullableTraitOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTraitOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
