/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 53.1.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the GetSourceAlphaOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSourceAlphaOutput{}

// GetSourceAlphaOutput Returns a Source.
type GetSourceAlphaOutput struct {
	Source SourceAlpha `json:"source"`
	// The id of the Tracking Plan connected to the Source.
	TrackingPlanId NullableString `json:"trackingPlanId"`
}

// NewGetSourceAlphaOutput instantiates a new GetSourceAlphaOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSourceAlphaOutput(
	source SourceAlpha,
	trackingPlanId NullableString,
) *GetSourceAlphaOutput {
	this := GetSourceAlphaOutput{}
	this.Source = source
	this.TrackingPlanId = trackingPlanId
	return &this
}

// NewGetSourceAlphaOutputWithDefaults instantiates a new GetSourceAlphaOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSourceAlphaOutputWithDefaults() *GetSourceAlphaOutput {
	this := GetSourceAlphaOutput{}
	return &this
}

// GetSource returns the Source field value
func (o *GetSourceAlphaOutput) GetSource() SourceAlpha {
	if o == nil {
		var ret SourceAlpha
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *GetSourceAlphaOutput) GetSourceOk() (*SourceAlpha, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *GetSourceAlphaOutput) SetSource(v SourceAlpha) {
	o.Source = v
}

// GetTrackingPlanId returns the TrackingPlanId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *GetSourceAlphaOutput) GetTrackingPlanId() string {
	if o == nil || o.TrackingPlanId.Get() == nil {
		var ret string
		return ret
	}

	return *o.TrackingPlanId.Get()
}

// GetTrackingPlanIdOk returns a tuple with the TrackingPlanId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetSourceAlphaOutput) GetTrackingPlanIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TrackingPlanId.Get(), o.TrackingPlanId.IsSet()
}

// SetTrackingPlanId sets field value
func (o *GetSourceAlphaOutput) SetTrackingPlanId(v string) {
	o.TrackingPlanId.Set(&v)
}

func (o GetSourceAlphaOutput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSourceAlphaOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["source"] = o.Source
	toSerialize["trackingPlanId"] = o.TrackingPlanId.Get()
	return toSerialize, nil
}

type NullableGetSourceAlphaOutput struct {
	value *GetSourceAlphaOutput
	isSet bool
}

func (v NullableGetSourceAlphaOutput) Get() *GetSourceAlphaOutput {
	return v.value
}

func (v *NullableGetSourceAlphaOutput) Set(val *GetSourceAlphaOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSourceAlphaOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSourceAlphaOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSourceAlphaOutput(val *GetSourceAlphaOutput) *NullableGetSourceAlphaOutput {
	return &NullableGetSourceAlphaOutput{value: val, isSet: true}
}

func (v NullableGetSourceAlphaOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSourceAlphaOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
