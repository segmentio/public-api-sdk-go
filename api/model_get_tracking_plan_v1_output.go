/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 34.3.1
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// GetTrackingPlanV1Output Gets a single Tracking Plan.
type GetTrackingPlanV1Output struct {
	TrackingPlan TrackingPlan `json:"trackingPlan"`
}

// NewGetTrackingPlanV1Output instantiates a new GetTrackingPlanV1Output object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTrackingPlanV1Output(trackingPlan TrackingPlan) *GetTrackingPlanV1Output {
	this := GetTrackingPlanV1Output{}
	this.TrackingPlan = trackingPlan
	return &this
}

// NewGetTrackingPlanV1OutputWithDefaults instantiates a new GetTrackingPlanV1Output object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTrackingPlanV1OutputWithDefaults() *GetTrackingPlanV1Output {
	this := GetTrackingPlanV1Output{}
	return &this
}

// GetTrackingPlan returns the TrackingPlan field value
func (o *GetTrackingPlanV1Output) GetTrackingPlan() TrackingPlan {
	if o == nil {
		var ret TrackingPlan
		return ret
	}

	return o.TrackingPlan
}

// GetTrackingPlanOk returns a tuple with the TrackingPlan field value
// and a boolean to check if the value has been set.
func (o *GetTrackingPlanV1Output) GetTrackingPlanOk() (*TrackingPlan, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TrackingPlan, true
}

// SetTrackingPlan sets field value
func (o *GetTrackingPlanV1Output) SetTrackingPlan(v TrackingPlan) {
	o.TrackingPlan = v
}

func (o GetTrackingPlanV1Output) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["trackingPlan"] = o.TrackingPlan
	}
	return json.Marshal(toSerialize)
}

type NullableGetTrackingPlanV1Output struct {
	value *GetTrackingPlanV1Output
	isSet bool
}

func (v NullableGetTrackingPlanV1Output) Get() *GetTrackingPlanV1Output {
	return v.value
}

func (v *NullableGetTrackingPlanV1Output) Set(val *GetTrackingPlanV1Output) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTrackingPlanV1Output) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTrackingPlanV1Output) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTrackingPlanV1Output(
	val *GetTrackingPlanV1Output,
) *NullableGetTrackingPlanV1Output {
	return &NullableGetTrackingPlanV1Output{value: val, isSet: true}
}

func (v NullableGetTrackingPlanV1Output) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTrackingPlanV1Output) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
