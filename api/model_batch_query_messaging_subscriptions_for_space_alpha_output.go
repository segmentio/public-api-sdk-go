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

// checks if the BatchQueryMessagingSubscriptionsForSpaceAlphaOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BatchQueryMessagingSubscriptionsForSpaceAlphaOutput{}

// BatchQueryMessagingSubscriptionsForSpaceAlphaOutput Batch get response.
type BatchQueryMessagingSubscriptionsForSpaceAlphaOutput struct {
	// Array of successful subscription status.
	Successes []GetMessagingSubscriptionSuccessResponse `json:"successes"`
	// Validation errors due to invalid types or email/phone numbers.
	Failures []GetMessagingSubscriptionFailureResponse `json:"failures"`
	// General errors when making the request such as invalid payload or wrong http method errors.
	Errors     []MessageSubscriptionResponseError `json:"errors"`
	Pagination *PaginationOutput                  `json:"pagination,omitempty"`
}

// NewBatchQueryMessagingSubscriptionsForSpaceAlphaOutput instantiates a new BatchQueryMessagingSubscriptionsForSpaceAlphaOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBatchQueryMessagingSubscriptionsForSpaceAlphaOutput(
	successes []GetMessagingSubscriptionSuccessResponse,
	failures []GetMessagingSubscriptionFailureResponse,
	errors []MessageSubscriptionResponseError,
) *BatchQueryMessagingSubscriptionsForSpaceAlphaOutput {
	this := BatchQueryMessagingSubscriptionsForSpaceAlphaOutput{}
	this.Successes = successes
	this.Failures = failures
	this.Errors = errors
	return &this
}

// NewBatchQueryMessagingSubscriptionsForSpaceAlphaOutputWithDefaults instantiates a new BatchQueryMessagingSubscriptionsForSpaceAlphaOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBatchQueryMessagingSubscriptionsForSpaceAlphaOutputWithDefaults() *BatchQueryMessagingSubscriptionsForSpaceAlphaOutput {
	this := BatchQueryMessagingSubscriptionsForSpaceAlphaOutput{}
	return &this
}

// GetSuccesses returns the Successes field value
func (o *BatchQueryMessagingSubscriptionsForSpaceAlphaOutput) GetSuccesses() []GetMessagingSubscriptionSuccessResponse {
	if o == nil {
		var ret []GetMessagingSubscriptionSuccessResponse
		return ret
	}

	return o.Successes
}

// GetSuccessesOk returns a tuple with the Successes field value
// and a boolean to check if the value has been set.
func (o *BatchQueryMessagingSubscriptionsForSpaceAlphaOutput) GetSuccessesOk() ([]GetMessagingSubscriptionSuccessResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Successes, true
}

// SetSuccesses sets field value
func (o *BatchQueryMessagingSubscriptionsForSpaceAlphaOutput) SetSuccesses(
	v []GetMessagingSubscriptionSuccessResponse,
) {
	o.Successes = v
}

// GetFailures returns the Failures field value
func (o *BatchQueryMessagingSubscriptionsForSpaceAlphaOutput) GetFailures() []GetMessagingSubscriptionFailureResponse {
	if o == nil {
		var ret []GetMessagingSubscriptionFailureResponse
		return ret
	}

	return o.Failures
}

// GetFailuresOk returns a tuple with the Failures field value
// and a boolean to check if the value has been set.
func (o *BatchQueryMessagingSubscriptionsForSpaceAlphaOutput) GetFailuresOk() ([]GetMessagingSubscriptionFailureResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Failures, true
}

// SetFailures sets field value
func (o *BatchQueryMessagingSubscriptionsForSpaceAlphaOutput) SetFailures(
	v []GetMessagingSubscriptionFailureResponse,
) {
	o.Failures = v
}

// GetErrors returns the Errors field value
func (o *BatchQueryMessagingSubscriptionsForSpaceAlphaOutput) GetErrors() []MessageSubscriptionResponseError {
	if o == nil {
		var ret []MessageSubscriptionResponseError
		return ret
	}

	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value
// and a boolean to check if the value has been set.
func (o *BatchQueryMessagingSubscriptionsForSpaceAlphaOutput) GetErrorsOk() ([]MessageSubscriptionResponseError, bool) {
	if o == nil {
		return nil, false
	}
	return o.Errors, true
}

// SetErrors sets field value
func (o *BatchQueryMessagingSubscriptionsForSpaceAlphaOutput) SetErrors(
	v []MessageSubscriptionResponseError,
) {
	o.Errors = v
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *BatchQueryMessagingSubscriptionsForSpaceAlphaOutput) GetPagination() PaginationOutput {
	if o == nil || IsNil(o.Pagination) {
		var ret PaginationOutput
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchQueryMessagingSubscriptionsForSpaceAlphaOutput) GetPaginationOk() (*PaginationOutput, bool) {
	if o == nil || IsNil(o.Pagination) {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *BatchQueryMessagingSubscriptionsForSpaceAlphaOutput) HasPagination() bool {
	if o != nil && !IsNil(o.Pagination) {
		return true
	}

	return false
}

// SetPagination gets a reference to the given PaginationOutput and assigns it to the Pagination field.
func (o *BatchQueryMessagingSubscriptionsForSpaceAlphaOutput) SetPagination(v PaginationOutput) {
	o.Pagination = &v
}

func (o BatchQueryMessagingSubscriptionsForSpaceAlphaOutput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BatchQueryMessagingSubscriptionsForSpaceAlphaOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["successes"] = o.Successes
	toSerialize["failures"] = o.Failures
	toSerialize["errors"] = o.Errors
	if !IsNil(o.Pagination) {
		toSerialize["pagination"] = o.Pagination
	}
	return toSerialize, nil
}

type NullableBatchQueryMessagingSubscriptionsForSpaceAlphaOutput struct {
	value *BatchQueryMessagingSubscriptionsForSpaceAlphaOutput
	isSet bool
}

func (v NullableBatchQueryMessagingSubscriptionsForSpaceAlphaOutput) Get() *BatchQueryMessagingSubscriptionsForSpaceAlphaOutput {
	return v.value
}

func (v *NullableBatchQueryMessagingSubscriptionsForSpaceAlphaOutput) Set(
	val *BatchQueryMessagingSubscriptionsForSpaceAlphaOutput,
) {
	v.value = val
	v.isSet = true
}

func (v NullableBatchQueryMessagingSubscriptionsForSpaceAlphaOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableBatchQueryMessagingSubscriptionsForSpaceAlphaOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBatchQueryMessagingSubscriptionsForSpaceAlphaOutput(
	val *BatchQueryMessagingSubscriptionsForSpaceAlphaOutput,
) *NullableBatchQueryMessagingSubscriptionsForSpaceAlphaOutput {
	return &NullableBatchQueryMessagingSubscriptionsForSpaceAlphaOutput{value: val, isSet: true}
}

func (v NullableBatchQueryMessagingSubscriptionsForSpaceAlphaOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBatchQueryMessagingSubscriptionsForSpaceAlphaOutput) UnmarshalJSON(
	src []byte,
) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
