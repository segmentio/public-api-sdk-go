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

// UpdateRulesInTrackingPlanV1Input Updates Tracking Plan rules. Non-existent rules are added.
type UpdateRulesInTrackingPlanV1Input struct {
	// Rules to update or insert.
	Rules []UpsertRuleV1 `json:"rules"`
}

// NewUpdateRulesInTrackingPlanV1Input instantiates a new UpdateRulesInTrackingPlanV1Input object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateRulesInTrackingPlanV1Input(rules []UpsertRuleV1) *UpdateRulesInTrackingPlanV1Input {
	this := UpdateRulesInTrackingPlanV1Input{}
	this.Rules = rules
	return &this
}

// NewUpdateRulesInTrackingPlanV1InputWithDefaults instantiates a new UpdateRulesInTrackingPlanV1Input object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateRulesInTrackingPlanV1InputWithDefaults() *UpdateRulesInTrackingPlanV1Input {
	this := UpdateRulesInTrackingPlanV1Input{}
	return &this
}

// GetRules returns the Rules field value
func (o *UpdateRulesInTrackingPlanV1Input) GetRules() []UpsertRuleV1 {
	if o == nil {
		var ret []UpsertRuleV1
		return ret
	}

	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value
// and a boolean to check if the value has been set.
func (o *UpdateRulesInTrackingPlanV1Input) GetRulesOk() ([]UpsertRuleV1, bool) {
	if o == nil {
		return nil, false
	}
	return o.Rules, true
}

// SetRules sets field value
func (o *UpdateRulesInTrackingPlanV1Input) SetRules(v []UpsertRuleV1) {
	o.Rules = v
}

func (o UpdateRulesInTrackingPlanV1Input) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["rules"] = o.Rules
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateRulesInTrackingPlanV1Input struct {
	value *UpdateRulesInTrackingPlanV1Input
	isSet bool
}

func (v NullableUpdateRulesInTrackingPlanV1Input) Get() *UpdateRulesInTrackingPlanV1Input {
	return v.value
}

func (v *NullableUpdateRulesInTrackingPlanV1Input) Set(val *UpdateRulesInTrackingPlanV1Input) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateRulesInTrackingPlanV1Input) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateRulesInTrackingPlanV1Input) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateRulesInTrackingPlanV1Input(
	val *UpdateRulesInTrackingPlanV1Input,
) *NullableUpdateRulesInTrackingPlanV1Input {
	return &NullableUpdateRulesInTrackingPlanV1Input{value: val, isSet: true}
}

func (v NullableUpdateRulesInTrackingPlanV1Input) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateRulesInTrackingPlanV1Input) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
