/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API. 

API version: 32.0.4
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// UpdateSubscriptionForDestinationAlphaInput The basic input parameters for updating a Destination subscription.
type UpdateSubscriptionForDestinationAlphaInput struct {
	Input Input `json:"input"`
}

// NewUpdateSubscriptionForDestinationAlphaInput instantiates a new UpdateSubscriptionForDestinationAlphaInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateSubscriptionForDestinationAlphaInput(input Input) *UpdateSubscriptionForDestinationAlphaInput {
	this := UpdateSubscriptionForDestinationAlphaInput{}
	this.Input = input
	return &this
}

// NewUpdateSubscriptionForDestinationAlphaInputWithDefaults instantiates a new UpdateSubscriptionForDestinationAlphaInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateSubscriptionForDestinationAlphaInputWithDefaults() *UpdateSubscriptionForDestinationAlphaInput {
	this := UpdateSubscriptionForDestinationAlphaInput{}
	return &this
}

// GetInput returns the Input field value
func (o *UpdateSubscriptionForDestinationAlphaInput) GetInput() Input {
	if o == nil {
		var ret Input
		return ret
	}

	return o.Input
}

// GetInputOk returns a tuple with the Input field value
// and a boolean to check if the value has been set.
func (o *UpdateSubscriptionForDestinationAlphaInput) GetInputOk() (*Input, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Input, true
}

// SetInput sets field value
func (o *UpdateSubscriptionForDestinationAlphaInput) SetInput(v Input) {
	o.Input = v
}

func (o UpdateSubscriptionForDestinationAlphaInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["input"] = o.Input
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateSubscriptionForDestinationAlphaInput struct {
	value *UpdateSubscriptionForDestinationAlphaInput
	isSet bool
}

func (v NullableUpdateSubscriptionForDestinationAlphaInput) Get() *UpdateSubscriptionForDestinationAlphaInput {
	return v.value
}

func (v *NullableUpdateSubscriptionForDestinationAlphaInput) Set(val *UpdateSubscriptionForDestinationAlphaInput) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateSubscriptionForDestinationAlphaInput) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateSubscriptionForDestinationAlphaInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateSubscriptionForDestinationAlphaInput(val *UpdateSubscriptionForDestinationAlphaInput) *NullableUpdateSubscriptionForDestinationAlphaInput {
	return &NullableUpdateSubscriptionForDestinationAlphaInput{value: val, isSet: true}
}

func (v NullableUpdateSubscriptionForDestinationAlphaInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateSubscriptionForDestinationAlphaInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


