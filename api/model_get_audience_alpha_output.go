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

// GetAudienceAlphaOutput Audience output for update.
type GetAudienceAlphaOutput struct {
	Audience Audience `json:"audience"`
}

// NewGetAudienceAlphaOutput instantiates a new GetAudienceAlphaOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAudienceAlphaOutput(audience Audience) *GetAudienceAlphaOutput {
	this := GetAudienceAlphaOutput{}
	this.Audience = audience
	return &this
}

// NewGetAudienceAlphaOutputWithDefaults instantiates a new GetAudienceAlphaOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAudienceAlphaOutputWithDefaults() *GetAudienceAlphaOutput {
	this := GetAudienceAlphaOutput{}
	return &this
}

// GetAudience returns the Audience field value
func (o *GetAudienceAlphaOutput) GetAudience() Audience {
	if o == nil {
		var ret Audience
		return ret
	}

	return o.Audience
}

// GetAudienceOk returns a tuple with the Audience field value
// and a boolean to check if the value has been set.
func (o *GetAudienceAlphaOutput) GetAudienceOk() (*Audience, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Audience, true
}

// SetAudience sets field value
func (o *GetAudienceAlphaOutput) SetAudience(v Audience) {
	o.Audience = v
}

func (o GetAudienceAlphaOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["audience"] = o.Audience
	}
	return json.Marshal(toSerialize)
}

type NullableGetAudienceAlphaOutput struct {
	value *GetAudienceAlphaOutput
	isSet bool
}

func (v NullableGetAudienceAlphaOutput) Get() *GetAudienceAlphaOutput {
	return v.value
}

func (v *NullableGetAudienceAlphaOutput) Set(val *GetAudienceAlphaOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAudienceAlphaOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAudienceAlphaOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAudienceAlphaOutput(
	val *GetAudienceAlphaOutput,
) *NullableGetAudienceAlphaOutput {
	return &NullableGetAudienceAlphaOutput{value: val, isSet: true}
}

func (v NullableGetAudienceAlphaOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAudienceAlphaOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
