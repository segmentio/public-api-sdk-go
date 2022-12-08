/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 33.0.4
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// UpdateRulesInTrackingPlanV1Output Updates Tracking Plan rules. Non-existent rules are added.
type UpdateRulesInTrackingPlanV1Output struct {
	// The operation status.
	Status string `json:"status"`
}

// NewUpdateRulesInTrackingPlanV1Output instantiates a new UpdateRulesInTrackingPlanV1Output object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateRulesInTrackingPlanV1Output(status string) *UpdateRulesInTrackingPlanV1Output {
	this := UpdateRulesInTrackingPlanV1Output{}
	this.Status = status
	return &this
}

// NewUpdateRulesInTrackingPlanV1OutputWithDefaults instantiates a new UpdateRulesInTrackingPlanV1Output object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateRulesInTrackingPlanV1OutputWithDefaults() *UpdateRulesInTrackingPlanV1Output {
	this := UpdateRulesInTrackingPlanV1Output{}
	return &this
}

// GetStatus returns the Status field value
func (o *UpdateRulesInTrackingPlanV1Output) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *UpdateRulesInTrackingPlanV1Output) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *UpdateRulesInTrackingPlanV1Output) SetStatus(v string) {
	o.Status = v
}

func (o UpdateRulesInTrackingPlanV1Output) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateRulesInTrackingPlanV1Output struct {
	value *UpdateRulesInTrackingPlanV1Output
	isSet bool
}

func (v NullableUpdateRulesInTrackingPlanV1Output) Get() *UpdateRulesInTrackingPlanV1Output {
	return v.value
}

func (v *NullableUpdateRulesInTrackingPlanV1Output) Set(val *UpdateRulesInTrackingPlanV1Output) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateRulesInTrackingPlanV1Output) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateRulesInTrackingPlanV1Output) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateRulesInTrackingPlanV1Output(
	val *UpdateRulesInTrackingPlanV1Output,
) *NullableUpdateRulesInTrackingPlanV1Output {
	return &NullableUpdateRulesInTrackingPlanV1Output{value: val, isSet: true}
}

func (v NullableUpdateRulesInTrackingPlanV1Output) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateRulesInTrackingPlanV1Output) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
