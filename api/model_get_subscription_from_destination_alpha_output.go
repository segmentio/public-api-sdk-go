/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API. 

API version: 33.0.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// GetSubscriptionFromDestinationAlphaOutput Returns a subscription for a given subscription id.
type GetSubscriptionFromDestinationAlphaOutput struct {
	Subscription Subscription `json:"subscription"`
}

// NewGetSubscriptionFromDestinationAlphaOutput instantiates a new GetSubscriptionFromDestinationAlphaOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSubscriptionFromDestinationAlphaOutput(subscription Subscription) *GetSubscriptionFromDestinationAlphaOutput {
	this := GetSubscriptionFromDestinationAlphaOutput{}
	this.Subscription = subscription
	return &this
}

// NewGetSubscriptionFromDestinationAlphaOutputWithDefaults instantiates a new GetSubscriptionFromDestinationAlphaOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSubscriptionFromDestinationAlphaOutputWithDefaults() *GetSubscriptionFromDestinationAlphaOutput {
	this := GetSubscriptionFromDestinationAlphaOutput{}
	return &this
}

// GetSubscription returns the Subscription field value
func (o *GetSubscriptionFromDestinationAlphaOutput) GetSubscription() Subscription {
	if o == nil {
		var ret Subscription
		return ret
	}

	return o.Subscription
}

// GetSubscriptionOk returns a tuple with the Subscription field value
// and a boolean to check if the value has been set.
func (o *GetSubscriptionFromDestinationAlphaOutput) GetSubscriptionOk() (*Subscription, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subscription, true
}

// SetSubscription sets field value
func (o *GetSubscriptionFromDestinationAlphaOutput) SetSubscription(v Subscription) {
	o.Subscription = v
}

func (o GetSubscriptionFromDestinationAlphaOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["subscription"] = o.Subscription
	}
	return json.Marshal(toSerialize)
}

type NullableGetSubscriptionFromDestinationAlphaOutput struct {
	value *GetSubscriptionFromDestinationAlphaOutput
	isSet bool
}

func (v NullableGetSubscriptionFromDestinationAlphaOutput) Get() *GetSubscriptionFromDestinationAlphaOutput {
	return v.value
}

func (v *NullableGetSubscriptionFromDestinationAlphaOutput) Set(val *GetSubscriptionFromDestinationAlphaOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSubscriptionFromDestinationAlphaOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSubscriptionFromDestinationAlphaOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSubscriptionFromDestinationAlphaOutput(val *GetSubscriptionFromDestinationAlphaOutput) *NullableGetSubscriptionFromDestinationAlphaOutput {
	return &NullableGetSubscriptionFromDestinationAlphaOutput{value: val, isSet: true}
}

func (v NullableGetSubscriptionFromDestinationAlphaOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSubscriptionFromDestinationAlphaOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


