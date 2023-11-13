/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 37.2.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// RemoveAudienceFromSpaceAlphaOutput Delete audience endpoint output.
type RemoveAudienceFromSpaceAlphaOutput struct {
	// The status of the operation.
	Status string `json:"status"`
}

// NewRemoveAudienceFromSpaceAlphaOutput instantiates a new RemoveAudienceFromSpaceAlphaOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemoveAudienceFromSpaceAlphaOutput(status string) *RemoveAudienceFromSpaceAlphaOutput {
	this := RemoveAudienceFromSpaceAlphaOutput{}
	this.Status = status
	return &this
}

// NewRemoveAudienceFromSpaceAlphaOutputWithDefaults instantiates a new RemoveAudienceFromSpaceAlphaOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemoveAudienceFromSpaceAlphaOutputWithDefaults() *RemoveAudienceFromSpaceAlphaOutput {
	this := RemoveAudienceFromSpaceAlphaOutput{}
	return &this
}

// GetStatus returns the Status field value
func (o *RemoveAudienceFromSpaceAlphaOutput) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *RemoveAudienceFromSpaceAlphaOutput) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *RemoveAudienceFromSpaceAlphaOutput) SetStatus(v string) {
	o.Status = v
}

func (o RemoveAudienceFromSpaceAlphaOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableRemoveAudienceFromSpaceAlphaOutput struct {
	value *RemoveAudienceFromSpaceAlphaOutput
	isSet bool
}

func (v NullableRemoveAudienceFromSpaceAlphaOutput) Get() *RemoveAudienceFromSpaceAlphaOutput {
	return v.value
}

func (v *NullableRemoveAudienceFromSpaceAlphaOutput) Set(val *RemoveAudienceFromSpaceAlphaOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableRemoveAudienceFromSpaceAlphaOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableRemoveAudienceFromSpaceAlphaOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemoveAudienceFromSpaceAlphaOutput(
	val *RemoveAudienceFromSpaceAlphaOutput,
) *NullableRemoveAudienceFromSpaceAlphaOutput {
	return &NullableRemoveAudienceFromSpaceAlphaOutput{value: val, isSet: true}
}

func (v NullableRemoveAudienceFromSpaceAlphaOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemoveAudienceFromSpaceAlphaOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
