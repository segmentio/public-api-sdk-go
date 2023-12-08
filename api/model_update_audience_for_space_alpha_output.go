/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 38.4.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the UpdateAudienceForSpaceAlphaOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateAudienceForSpaceAlphaOutput{}

// UpdateAudienceForSpaceAlphaOutput Audience output for get.
type UpdateAudienceForSpaceAlphaOutput struct {
	Audience AudienceSummary `json:"audience"`
}

// NewUpdateAudienceForSpaceAlphaOutput instantiates a new UpdateAudienceForSpaceAlphaOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateAudienceForSpaceAlphaOutput(
	audience AudienceSummary,
) *UpdateAudienceForSpaceAlphaOutput {
	this := UpdateAudienceForSpaceAlphaOutput{}
	this.Audience = audience
	return &this
}

// NewUpdateAudienceForSpaceAlphaOutputWithDefaults instantiates a new UpdateAudienceForSpaceAlphaOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateAudienceForSpaceAlphaOutputWithDefaults() *UpdateAudienceForSpaceAlphaOutput {
	this := UpdateAudienceForSpaceAlphaOutput{}
	return &this
}

// GetAudience returns the Audience field value
func (o *UpdateAudienceForSpaceAlphaOutput) GetAudience() AudienceSummary {
	if o == nil {
		var ret AudienceSummary
		return ret
	}

	return o.Audience
}

// GetAudienceOk returns a tuple with the Audience field value
// and a boolean to check if the value has been set.
func (o *UpdateAudienceForSpaceAlphaOutput) GetAudienceOk() (*AudienceSummary, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Audience, true
}

// SetAudience sets field value
func (o *UpdateAudienceForSpaceAlphaOutput) SetAudience(v AudienceSummary) {
	o.Audience = v
}

func (o UpdateAudienceForSpaceAlphaOutput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateAudienceForSpaceAlphaOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["audience"] = o.Audience
	return toSerialize, nil
}

type NullableUpdateAudienceForSpaceAlphaOutput struct {
	value *UpdateAudienceForSpaceAlphaOutput
	isSet bool
}

func (v NullableUpdateAudienceForSpaceAlphaOutput) Get() *UpdateAudienceForSpaceAlphaOutput {
	return v.value
}

func (v *NullableUpdateAudienceForSpaceAlphaOutput) Set(val *UpdateAudienceForSpaceAlphaOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateAudienceForSpaceAlphaOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateAudienceForSpaceAlphaOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateAudienceForSpaceAlphaOutput(
	val *UpdateAudienceForSpaceAlphaOutput,
) *NullableUpdateAudienceForSpaceAlphaOutput {
	return &NullableUpdateAudienceForSpaceAlphaOutput{value: val, isSet: true}
}

func (v NullableUpdateAudienceForSpaceAlphaOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateAudienceForSpaceAlphaOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
