/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API. 

API version: 32.0.2
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// UpdateSubscriptionForDestinationAlphaOutput Returns the updated Destination subscription.
type UpdateSubscriptionForDestinationAlphaOutput struct {
	Subscription Subscription `json:"subscription"`
}

// NewUpdateSubscriptionForDestinationAlphaOutput instantiates a new UpdateSubscriptionForDestinationAlphaOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateSubscriptionForDestinationAlphaOutput(subscription Subscription) *UpdateSubscriptionForDestinationAlphaOutput {
	this := UpdateSubscriptionForDestinationAlphaOutput{}
	this.Subscription = subscription
	return &this
}

// NewUpdateSubscriptionForDestinationAlphaOutputWithDefaults instantiates a new UpdateSubscriptionForDestinationAlphaOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateSubscriptionForDestinationAlphaOutputWithDefaults() *UpdateSubscriptionForDestinationAlphaOutput {
	this := UpdateSubscriptionForDestinationAlphaOutput{}
	return &this
}

// GetSubscription returns the Subscription field value
func (o *UpdateSubscriptionForDestinationAlphaOutput) GetSubscription() Subscription {
	if o == nil {
		var ret Subscription
		return ret
	}

	return o.Subscription
}

// GetSubscriptionOk returns a tuple with the Subscription field value
// and a boolean to check if the value has been set.
func (o *UpdateSubscriptionForDestinationAlphaOutput) GetSubscriptionOk() (*Subscription, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subscription, true
}

// SetSubscription sets field value
func (o *UpdateSubscriptionForDestinationAlphaOutput) SetSubscription(v Subscription) {
	o.Subscription = v
}

func (o UpdateSubscriptionForDestinationAlphaOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["subscription"] = o.Subscription
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateSubscriptionForDestinationAlphaOutput struct {
	value *UpdateSubscriptionForDestinationAlphaOutput
	isSet bool
}

func (v NullableUpdateSubscriptionForDestinationAlphaOutput) Get() *UpdateSubscriptionForDestinationAlphaOutput {
	return v.value
}

func (v *NullableUpdateSubscriptionForDestinationAlphaOutput) Set(val *UpdateSubscriptionForDestinationAlphaOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateSubscriptionForDestinationAlphaOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateSubscriptionForDestinationAlphaOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateSubscriptionForDestinationAlphaOutput(val *UpdateSubscriptionForDestinationAlphaOutput) *NullableUpdateSubscriptionForDestinationAlphaOutput {
	return &NullableUpdateSubscriptionForDestinationAlphaOutput{value: val, isSet: true}
}

func (v NullableUpdateSubscriptionForDestinationAlphaOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateSubscriptionForDestinationAlphaOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


