/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API. 

API version: 33.0.2
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// ReplaceMessagingSubscriptionsInSpacesAlphaOutput Output for the endpoint.
type ReplaceMessagingSubscriptionsInSpacesAlphaOutput struct {
	// Array of successful subscription status.
	Successes []MessageSubscriptionResponse `json:"successes"`
	// Array of failure subscription status.
	Failures []MessageSubscriptionResponse `json:"failures"`
	Pagination *Pagination `json:"pagination,omitempty"`
}

// NewReplaceMessagingSubscriptionsInSpacesAlphaOutput instantiates a new ReplaceMessagingSubscriptionsInSpacesAlphaOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReplaceMessagingSubscriptionsInSpacesAlphaOutput(successes []MessageSubscriptionResponse, failures []MessageSubscriptionResponse) *ReplaceMessagingSubscriptionsInSpacesAlphaOutput {
	this := ReplaceMessagingSubscriptionsInSpacesAlphaOutput{}
	this.Successes = successes
	this.Failures = failures
	return &this
}

// NewReplaceMessagingSubscriptionsInSpacesAlphaOutputWithDefaults instantiates a new ReplaceMessagingSubscriptionsInSpacesAlphaOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReplaceMessagingSubscriptionsInSpacesAlphaOutputWithDefaults() *ReplaceMessagingSubscriptionsInSpacesAlphaOutput {
	this := ReplaceMessagingSubscriptionsInSpacesAlphaOutput{}
	return &this
}

// GetSuccesses returns the Successes field value
func (o *ReplaceMessagingSubscriptionsInSpacesAlphaOutput) GetSuccesses() []MessageSubscriptionResponse {
	if o == nil {
		var ret []MessageSubscriptionResponse
		return ret
	}

	return o.Successes
}

// GetSuccessesOk returns a tuple with the Successes field value
// and a boolean to check if the value has been set.
func (o *ReplaceMessagingSubscriptionsInSpacesAlphaOutput) GetSuccessesOk() ([]MessageSubscriptionResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Successes, true
}

// SetSuccesses sets field value
func (o *ReplaceMessagingSubscriptionsInSpacesAlphaOutput) SetSuccesses(v []MessageSubscriptionResponse) {
	o.Successes = v
}

// GetFailures returns the Failures field value
func (o *ReplaceMessagingSubscriptionsInSpacesAlphaOutput) GetFailures() []MessageSubscriptionResponse {
	if o == nil {
		var ret []MessageSubscriptionResponse
		return ret
	}

	return o.Failures
}

// GetFailuresOk returns a tuple with the Failures field value
// and a boolean to check if the value has been set.
func (o *ReplaceMessagingSubscriptionsInSpacesAlphaOutput) GetFailuresOk() ([]MessageSubscriptionResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Failures, true
}

// SetFailures sets field value
func (o *ReplaceMessagingSubscriptionsInSpacesAlphaOutput) SetFailures(v []MessageSubscriptionResponse) {
	o.Failures = v
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *ReplaceMessagingSubscriptionsInSpacesAlphaOutput) GetPagination() Pagination {
	if o == nil || o.Pagination == nil {
		var ret Pagination
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplaceMessagingSubscriptionsInSpacesAlphaOutput) GetPaginationOk() (*Pagination, bool) {
	if o == nil || o.Pagination == nil {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *ReplaceMessagingSubscriptionsInSpacesAlphaOutput) HasPagination() bool {
	if o != nil && o.Pagination != nil {
		return true
	}

	return false
}

// SetPagination gets a reference to the given Pagination and assigns it to the Pagination field.
func (o *ReplaceMessagingSubscriptionsInSpacesAlphaOutput) SetPagination(v Pagination) {
	o.Pagination = &v
}

func (o ReplaceMessagingSubscriptionsInSpacesAlphaOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["successes"] = o.Successes
	}
	if true {
		toSerialize["failures"] = o.Failures
	}
	if o.Pagination != nil {
		toSerialize["pagination"] = o.Pagination
	}
	return json.Marshal(toSerialize)
}

type NullableReplaceMessagingSubscriptionsInSpacesAlphaOutput struct {
	value *ReplaceMessagingSubscriptionsInSpacesAlphaOutput
	isSet bool
}

func (v NullableReplaceMessagingSubscriptionsInSpacesAlphaOutput) Get() *ReplaceMessagingSubscriptionsInSpacesAlphaOutput {
	return v.value
}

func (v *NullableReplaceMessagingSubscriptionsInSpacesAlphaOutput) Set(val *ReplaceMessagingSubscriptionsInSpacesAlphaOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableReplaceMessagingSubscriptionsInSpacesAlphaOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableReplaceMessagingSubscriptionsInSpacesAlphaOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReplaceMessagingSubscriptionsInSpacesAlphaOutput(val *ReplaceMessagingSubscriptionsInSpacesAlphaOutput) *NullableReplaceMessagingSubscriptionsInSpacesAlphaOutput {
	return &NullableReplaceMessagingSubscriptionsInSpacesAlphaOutput{value: val, isSet: true}
}

func (v NullableReplaceMessagingSubscriptionsInSpacesAlphaOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReplaceMessagingSubscriptionsInSpacesAlphaOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


