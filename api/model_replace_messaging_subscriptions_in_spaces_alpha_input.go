/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 54.2.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the ReplaceMessagingSubscriptionsInSpacesAlphaInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReplaceMessagingSubscriptionsInSpacesAlphaInput{}

// ReplaceMessagingSubscriptionsInSpacesAlphaInput Input for the endpoint.
type ReplaceMessagingSubscriptionsInSpacesAlphaInput struct {
	// An array of subscriptions.
	Subscriptions []MessagesSubscriptionRequest `json:"subscriptions"`
}

// NewReplaceMessagingSubscriptionsInSpacesAlphaInput instantiates a new ReplaceMessagingSubscriptionsInSpacesAlphaInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReplaceMessagingSubscriptionsInSpacesAlphaInput(
	subscriptions []MessagesSubscriptionRequest,
) *ReplaceMessagingSubscriptionsInSpacesAlphaInput {
	this := ReplaceMessagingSubscriptionsInSpacesAlphaInput{}
	this.Subscriptions = subscriptions
	return &this
}

// NewReplaceMessagingSubscriptionsInSpacesAlphaInputWithDefaults instantiates a new ReplaceMessagingSubscriptionsInSpacesAlphaInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReplaceMessagingSubscriptionsInSpacesAlphaInputWithDefaults() *ReplaceMessagingSubscriptionsInSpacesAlphaInput {
	this := ReplaceMessagingSubscriptionsInSpacesAlphaInput{}
	return &this
}

// GetSubscriptions returns the Subscriptions field value
func (o *ReplaceMessagingSubscriptionsInSpacesAlphaInput) GetSubscriptions() []MessagesSubscriptionRequest {
	if o == nil {
		var ret []MessagesSubscriptionRequest
		return ret
	}

	return o.Subscriptions
}

// GetSubscriptionsOk returns a tuple with the Subscriptions field value
// and a boolean to check if the value has been set.
func (o *ReplaceMessagingSubscriptionsInSpacesAlphaInput) GetSubscriptionsOk() ([]MessagesSubscriptionRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Subscriptions, true
}

// SetSubscriptions sets field value
func (o *ReplaceMessagingSubscriptionsInSpacesAlphaInput) SetSubscriptions(
	v []MessagesSubscriptionRequest,
) {
	o.Subscriptions = v
}

func (o ReplaceMessagingSubscriptionsInSpacesAlphaInput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReplaceMessagingSubscriptionsInSpacesAlphaInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["subscriptions"] = o.Subscriptions
	return toSerialize, nil
}

type NullableReplaceMessagingSubscriptionsInSpacesAlphaInput struct {
	value *ReplaceMessagingSubscriptionsInSpacesAlphaInput
	isSet bool
}

func (v NullableReplaceMessagingSubscriptionsInSpacesAlphaInput) Get() *ReplaceMessagingSubscriptionsInSpacesAlphaInput {
	return v.value
}

func (v *NullableReplaceMessagingSubscriptionsInSpacesAlphaInput) Set(
	val *ReplaceMessagingSubscriptionsInSpacesAlphaInput,
) {
	v.value = val
	v.isSet = true
}

func (v NullableReplaceMessagingSubscriptionsInSpacesAlphaInput) IsSet() bool {
	return v.isSet
}

func (v *NullableReplaceMessagingSubscriptionsInSpacesAlphaInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReplaceMessagingSubscriptionsInSpacesAlphaInput(
	val *ReplaceMessagingSubscriptionsInSpacesAlphaInput,
) *NullableReplaceMessagingSubscriptionsInSpacesAlphaInput {
	return &NullableReplaceMessagingSubscriptionsInSpacesAlphaInput{value: val, isSet: true}
}

func (v NullableReplaceMessagingSubscriptionsInSpacesAlphaInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReplaceMessagingSubscriptionsInSpacesAlphaInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
