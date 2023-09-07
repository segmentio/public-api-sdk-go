/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 36.1.1
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// CreateDestinationSubscriptionAlphaOutput Returns a newly created Destination subscription.
type CreateDestinationSubscriptionAlphaOutput struct {
	DestinationSubscription DestinationSubscription `json:"destinationSubscription"`
}

// NewCreateDestinationSubscriptionAlphaOutput instantiates a new CreateDestinationSubscriptionAlphaOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateDestinationSubscriptionAlphaOutput(
	destinationSubscription DestinationSubscription,
) *CreateDestinationSubscriptionAlphaOutput {
	this := CreateDestinationSubscriptionAlphaOutput{}
	this.DestinationSubscription = destinationSubscription
	return &this
}

// NewCreateDestinationSubscriptionAlphaOutputWithDefaults instantiates a new CreateDestinationSubscriptionAlphaOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateDestinationSubscriptionAlphaOutputWithDefaults() *CreateDestinationSubscriptionAlphaOutput {
	this := CreateDestinationSubscriptionAlphaOutput{}
	return &this
}

// GetDestinationSubscription returns the DestinationSubscription field value
func (o *CreateDestinationSubscriptionAlphaOutput) GetDestinationSubscription() DestinationSubscription {
	if o == nil {
		var ret DestinationSubscription
		return ret
	}

	return o.DestinationSubscription
}

// GetDestinationSubscriptionOk returns a tuple with the DestinationSubscription field value
// and a boolean to check if the value has been set.
func (o *CreateDestinationSubscriptionAlphaOutput) GetDestinationSubscriptionOk() (*DestinationSubscription, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DestinationSubscription, true
}

// SetDestinationSubscription sets field value
func (o *CreateDestinationSubscriptionAlphaOutput) SetDestinationSubscription(
	v DestinationSubscription,
) {
	o.DestinationSubscription = v
}

func (o CreateDestinationSubscriptionAlphaOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["destinationSubscription"] = o.DestinationSubscription
	}
	return json.Marshal(toSerialize)
}

type NullableCreateDestinationSubscriptionAlphaOutput struct {
	value *CreateDestinationSubscriptionAlphaOutput
	isSet bool
}

func (v NullableCreateDestinationSubscriptionAlphaOutput) Get() *CreateDestinationSubscriptionAlphaOutput {
	return v.value
}

func (v *NullableCreateDestinationSubscriptionAlphaOutput) Set(
	val *CreateDestinationSubscriptionAlphaOutput,
) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateDestinationSubscriptionAlphaOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateDestinationSubscriptionAlphaOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateDestinationSubscriptionAlphaOutput(
	val *CreateDestinationSubscriptionAlphaOutput,
) *NullableCreateDestinationSubscriptionAlphaOutput {
	return &NullableCreateDestinationSubscriptionAlphaOutput{value: val, isSet: true}
}

func (v NullableCreateDestinationSubscriptionAlphaOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateDestinationSubscriptionAlphaOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
