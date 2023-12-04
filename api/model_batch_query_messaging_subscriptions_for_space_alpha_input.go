/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 38.2.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the BatchQueryMessagingSubscriptionsForSpaceAlphaInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BatchQueryMessagingSubscriptionsForSpaceAlphaInput{}

// BatchQueryMessagingSubscriptionsForSpaceAlphaInput Batch get request.
type BatchQueryMessagingSubscriptionsForSpaceAlphaInput struct {
	// A list of subscriptions to retrieve subscription status.
	Subscriptions []GetSubscriptionRequest `json:"subscriptions"`
}

// NewBatchQueryMessagingSubscriptionsForSpaceAlphaInput instantiates a new BatchQueryMessagingSubscriptionsForSpaceAlphaInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBatchQueryMessagingSubscriptionsForSpaceAlphaInput(
	subscriptions []GetSubscriptionRequest,
) *BatchQueryMessagingSubscriptionsForSpaceAlphaInput {
	this := BatchQueryMessagingSubscriptionsForSpaceAlphaInput{}
	this.Subscriptions = subscriptions
	return &this
}

// NewBatchQueryMessagingSubscriptionsForSpaceAlphaInputWithDefaults instantiates a new BatchQueryMessagingSubscriptionsForSpaceAlphaInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBatchQueryMessagingSubscriptionsForSpaceAlphaInputWithDefaults() *BatchQueryMessagingSubscriptionsForSpaceAlphaInput {
	this := BatchQueryMessagingSubscriptionsForSpaceAlphaInput{}
	return &this
}

// GetSubscriptions returns the Subscriptions field value
func (o *BatchQueryMessagingSubscriptionsForSpaceAlphaInput) GetSubscriptions() []GetSubscriptionRequest {
	if o == nil {
		var ret []GetSubscriptionRequest
		return ret
	}

	return o.Subscriptions
}

// GetSubscriptionsOk returns a tuple with the Subscriptions field value
// and a boolean to check if the value has been set.
func (o *BatchQueryMessagingSubscriptionsForSpaceAlphaInput) GetSubscriptionsOk() ([]GetSubscriptionRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Subscriptions, true
}

// SetSubscriptions sets field value
func (o *BatchQueryMessagingSubscriptionsForSpaceAlphaInput) SetSubscriptions(
	v []GetSubscriptionRequest,
) {
	o.Subscriptions = v
}

func (o BatchQueryMessagingSubscriptionsForSpaceAlphaInput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BatchQueryMessagingSubscriptionsForSpaceAlphaInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["subscriptions"] = o.Subscriptions
	return toSerialize, nil
}

type NullableBatchQueryMessagingSubscriptionsForSpaceAlphaInput struct {
	value *BatchQueryMessagingSubscriptionsForSpaceAlphaInput
	isSet bool
}

func (v NullableBatchQueryMessagingSubscriptionsForSpaceAlphaInput) Get() *BatchQueryMessagingSubscriptionsForSpaceAlphaInput {
	return v.value
}

func (v *NullableBatchQueryMessagingSubscriptionsForSpaceAlphaInput) Set(
	val *BatchQueryMessagingSubscriptionsForSpaceAlphaInput,
) {
	v.value = val
	v.isSet = true
}

func (v NullableBatchQueryMessagingSubscriptionsForSpaceAlphaInput) IsSet() bool {
	return v.isSet
}

func (v *NullableBatchQueryMessagingSubscriptionsForSpaceAlphaInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBatchQueryMessagingSubscriptionsForSpaceAlphaInput(
	val *BatchQueryMessagingSubscriptionsForSpaceAlphaInput,
) *NullableBatchQueryMessagingSubscriptionsForSpaceAlphaInput {
	return &NullableBatchQueryMessagingSubscriptionsForSpaceAlphaInput{value: val, isSet: true}
}

func (v NullableBatchQueryMessagingSubscriptionsForSpaceAlphaInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBatchQueryMessagingSubscriptionsForSpaceAlphaInput) UnmarshalJSON(
	src []byte,
) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
