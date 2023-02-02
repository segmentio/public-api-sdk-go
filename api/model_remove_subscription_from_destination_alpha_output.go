/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 34.0.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// RemoveSubscriptionFromDestinationAlphaOutput Returns a Destination deletion flag.
type RemoveSubscriptionFromDestinationAlphaOutput struct {
	// The status of the operation.
	Status string `json:"status"`
}

// NewRemoveSubscriptionFromDestinationAlphaOutput instantiates a new RemoveSubscriptionFromDestinationAlphaOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemoveSubscriptionFromDestinationAlphaOutput(
	status string,
) *RemoveSubscriptionFromDestinationAlphaOutput {
	this := RemoveSubscriptionFromDestinationAlphaOutput{}
	this.Status = status
	return &this
}

// NewRemoveSubscriptionFromDestinationAlphaOutputWithDefaults instantiates a new RemoveSubscriptionFromDestinationAlphaOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemoveSubscriptionFromDestinationAlphaOutputWithDefaults() *RemoveSubscriptionFromDestinationAlphaOutput {
	this := RemoveSubscriptionFromDestinationAlphaOutput{}
	return &this
}

// GetStatus returns the Status field value
func (o *RemoveSubscriptionFromDestinationAlphaOutput) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *RemoveSubscriptionFromDestinationAlphaOutput) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *RemoveSubscriptionFromDestinationAlphaOutput) SetStatus(v string) {
	o.Status = v
}

func (o RemoveSubscriptionFromDestinationAlphaOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableRemoveSubscriptionFromDestinationAlphaOutput struct {
	value *RemoveSubscriptionFromDestinationAlphaOutput
	isSet bool
}

func (v NullableRemoveSubscriptionFromDestinationAlphaOutput) Get() *RemoveSubscriptionFromDestinationAlphaOutput {
	return v.value
}

func (v *NullableRemoveSubscriptionFromDestinationAlphaOutput) Set(
	val *RemoveSubscriptionFromDestinationAlphaOutput,
) {
	v.value = val
	v.isSet = true
}

func (v NullableRemoveSubscriptionFromDestinationAlphaOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableRemoveSubscriptionFromDestinationAlphaOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemoveSubscriptionFromDestinationAlphaOutput(
	val *RemoveSubscriptionFromDestinationAlphaOutput,
) *NullableRemoveSubscriptionFromDestinationAlphaOutput {
	return &NullableRemoveSubscriptionFromDestinationAlphaOutput{value: val, isSet: true}
}

func (v NullableRemoveSubscriptionFromDestinationAlphaOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemoveSubscriptionFromDestinationAlphaOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
