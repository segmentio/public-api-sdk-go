/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 34.4.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// RemoveSourceFromTrackingPlanV1Output Disconnects a Source from a Tracking Plan.
type RemoveSourceFromTrackingPlanV1Output struct {
	// The operation status.
	Status string `json:"status"`
}

// NewRemoveSourceFromTrackingPlanV1Output instantiates a new RemoveSourceFromTrackingPlanV1Output object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemoveSourceFromTrackingPlanV1Output(status string) *RemoveSourceFromTrackingPlanV1Output {
	this := RemoveSourceFromTrackingPlanV1Output{}
	this.Status = status
	return &this
}

// NewRemoveSourceFromTrackingPlanV1OutputWithDefaults instantiates a new RemoveSourceFromTrackingPlanV1Output object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemoveSourceFromTrackingPlanV1OutputWithDefaults() *RemoveSourceFromTrackingPlanV1Output {
	this := RemoveSourceFromTrackingPlanV1Output{}
	return &this
}

// GetStatus returns the Status field value
func (o *RemoveSourceFromTrackingPlanV1Output) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *RemoveSourceFromTrackingPlanV1Output) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *RemoveSourceFromTrackingPlanV1Output) SetStatus(v string) {
	o.Status = v
}

func (o RemoveSourceFromTrackingPlanV1Output) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableRemoveSourceFromTrackingPlanV1Output struct {
	value *RemoveSourceFromTrackingPlanV1Output
	isSet bool
}

func (v NullableRemoveSourceFromTrackingPlanV1Output) Get() *RemoveSourceFromTrackingPlanV1Output {
	return v.value
}

func (v *NullableRemoveSourceFromTrackingPlanV1Output) Set(
	val *RemoveSourceFromTrackingPlanV1Output,
) {
	v.value = val
	v.isSet = true
}

func (v NullableRemoveSourceFromTrackingPlanV1Output) IsSet() bool {
	return v.isSet
}

func (v *NullableRemoveSourceFromTrackingPlanV1Output) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemoveSourceFromTrackingPlanV1Output(
	val *RemoveSourceFromTrackingPlanV1Output,
) *NullableRemoveSourceFromTrackingPlanV1Output {
	return &NullableRemoveSourceFromTrackingPlanV1Output{value: val, isSet: true}
}

func (v NullableRemoveSourceFromTrackingPlanV1Output) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemoveSourceFromTrackingPlanV1Output) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
