/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 50.1.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the AudienceCreateOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AudienceCreateOptions{}

// AudienceCreateOptions struct for AudienceCreateOptions
type AudienceCreateOptions struct {
	IncludeHistoricalData *bool `json:"includeHistoricalData,omitempty"`
	IncludeAnonymousUsers *bool `json:"includeAnonymousUsers,omitempty"`
}

// NewAudienceCreateOptions instantiates a new AudienceCreateOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAudienceCreateOptions() *AudienceCreateOptions {
	this := AudienceCreateOptions{}
	return &this
}

// NewAudienceCreateOptionsWithDefaults instantiates a new AudienceCreateOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAudienceCreateOptionsWithDefaults() *AudienceCreateOptions {
	this := AudienceCreateOptions{}
	return &this
}

// GetIncludeHistoricalData returns the IncludeHistoricalData field value if set, zero value otherwise.
func (o *AudienceCreateOptions) GetIncludeHistoricalData() bool {
	if o == nil || IsNil(o.IncludeHistoricalData) {
		var ret bool
		return ret
	}
	return *o.IncludeHistoricalData
}

// GetIncludeHistoricalDataOk returns a tuple with the IncludeHistoricalData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AudienceCreateOptions) GetIncludeHistoricalDataOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeHistoricalData) {
		return nil, false
	}
	return o.IncludeHistoricalData, true
}

// HasIncludeHistoricalData returns a boolean if a field has been set.
func (o *AudienceCreateOptions) HasIncludeHistoricalData() bool {
	if o != nil && !IsNil(o.IncludeHistoricalData) {
		return true
	}

	return false
}

// SetIncludeHistoricalData gets a reference to the given bool and assigns it to the IncludeHistoricalData field.
func (o *AudienceCreateOptions) SetIncludeHistoricalData(v bool) {
	o.IncludeHistoricalData = &v
}

// GetIncludeAnonymousUsers returns the IncludeAnonymousUsers field value if set, zero value otherwise.
func (o *AudienceCreateOptions) GetIncludeAnonymousUsers() bool {
	if o == nil || IsNil(o.IncludeAnonymousUsers) {
		var ret bool
		return ret
	}
	return *o.IncludeAnonymousUsers
}

// GetIncludeAnonymousUsersOk returns a tuple with the IncludeAnonymousUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AudienceCreateOptions) GetIncludeAnonymousUsersOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeAnonymousUsers) {
		return nil, false
	}
	return o.IncludeAnonymousUsers, true
}

// HasIncludeAnonymousUsers returns a boolean if a field has been set.
func (o *AudienceCreateOptions) HasIncludeAnonymousUsers() bool {
	if o != nil && !IsNil(o.IncludeAnonymousUsers) {
		return true
	}

	return false
}

// SetIncludeAnonymousUsers gets a reference to the given bool and assigns it to the IncludeAnonymousUsers field.
func (o *AudienceCreateOptions) SetIncludeAnonymousUsers(v bool) {
	o.IncludeAnonymousUsers = &v
}

func (o AudienceCreateOptions) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AudienceCreateOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IncludeHistoricalData) {
		toSerialize["includeHistoricalData"] = o.IncludeHistoricalData
	}
	if !IsNil(o.IncludeAnonymousUsers) {
		toSerialize["includeAnonymousUsers"] = o.IncludeAnonymousUsers
	}
	return toSerialize, nil
}

type NullableAudienceCreateOptions struct {
	value *AudienceCreateOptions
	isSet bool
}

func (v NullableAudienceCreateOptions) Get() *AudienceCreateOptions {
	return v.value
}

func (v *NullableAudienceCreateOptions) Set(val *AudienceCreateOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableAudienceCreateOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableAudienceCreateOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAudienceCreateOptions(val *AudienceCreateOptions) *NullableAudienceCreateOptions {
	return &NullableAudienceCreateOptions{value: val, isSet: true}
}

func (v NullableAudienceCreateOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAudienceCreateOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
