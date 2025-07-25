/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 58.13.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the ReplaceRulesInTrackingPlanV1Output type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReplaceRulesInTrackingPlanV1Output{}

// ReplaceRulesInTrackingPlanV1Output Replaces Tracking Plan rules.
type ReplaceRulesInTrackingPlanV1Output struct {
	// The operation status.
	Status string `json:"status"`
}

// NewReplaceRulesInTrackingPlanV1Output instantiates a new ReplaceRulesInTrackingPlanV1Output object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReplaceRulesInTrackingPlanV1Output(status string) *ReplaceRulesInTrackingPlanV1Output {
	this := ReplaceRulesInTrackingPlanV1Output{}
	this.Status = status
	return &this
}

// NewReplaceRulesInTrackingPlanV1OutputWithDefaults instantiates a new ReplaceRulesInTrackingPlanV1Output object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReplaceRulesInTrackingPlanV1OutputWithDefaults() *ReplaceRulesInTrackingPlanV1Output {
	this := ReplaceRulesInTrackingPlanV1Output{}
	return &this
}

// GetStatus returns the Status field value
func (o *ReplaceRulesInTrackingPlanV1Output) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ReplaceRulesInTrackingPlanV1Output) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ReplaceRulesInTrackingPlanV1Output) SetStatus(v string) {
	o.Status = v
}

func (o ReplaceRulesInTrackingPlanV1Output) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReplaceRulesInTrackingPlanV1Output) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

type NullableReplaceRulesInTrackingPlanV1Output struct {
	value *ReplaceRulesInTrackingPlanV1Output
	isSet bool
}

func (v NullableReplaceRulesInTrackingPlanV1Output) Get() *ReplaceRulesInTrackingPlanV1Output {
	return v.value
}

func (v *NullableReplaceRulesInTrackingPlanV1Output) Set(val *ReplaceRulesInTrackingPlanV1Output) {
	v.value = val
	v.isSet = true
}

func (v NullableReplaceRulesInTrackingPlanV1Output) IsSet() bool {
	return v.isSet
}

func (v *NullableReplaceRulesInTrackingPlanV1Output) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReplaceRulesInTrackingPlanV1Output(
	val *ReplaceRulesInTrackingPlanV1Output,
) *NullableReplaceRulesInTrackingPlanV1Output {
	return &NullableReplaceRulesInTrackingPlanV1Output{value: val, isSet: true}
}

func (v NullableReplaceRulesInTrackingPlanV1Output) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReplaceRulesInTrackingPlanV1Output) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
