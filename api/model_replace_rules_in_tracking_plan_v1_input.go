/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 34.3.5
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// ReplaceRulesInTrackingPlanV1Input Replaces Tracking Plan rules.
type ReplaceRulesInTrackingPlanV1Input struct {
	// Rules to replace.
	Rules []RuleV1 `json:"rules"`
}

// NewReplaceRulesInTrackingPlanV1Input instantiates a new ReplaceRulesInTrackingPlanV1Input object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReplaceRulesInTrackingPlanV1Input(rules []RuleV1) *ReplaceRulesInTrackingPlanV1Input {
	this := ReplaceRulesInTrackingPlanV1Input{}
	this.Rules = rules
	return &this
}

// NewReplaceRulesInTrackingPlanV1InputWithDefaults instantiates a new ReplaceRulesInTrackingPlanV1Input object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReplaceRulesInTrackingPlanV1InputWithDefaults() *ReplaceRulesInTrackingPlanV1Input {
	this := ReplaceRulesInTrackingPlanV1Input{}
	return &this
}

// GetRules returns the Rules field value
func (o *ReplaceRulesInTrackingPlanV1Input) GetRules() []RuleV1 {
	if o == nil {
		var ret []RuleV1
		return ret
	}

	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value
// and a boolean to check if the value has been set.
func (o *ReplaceRulesInTrackingPlanV1Input) GetRulesOk() ([]RuleV1, bool) {
	if o == nil {
		return nil, false
	}
	return o.Rules, true
}

// SetRules sets field value
func (o *ReplaceRulesInTrackingPlanV1Input) SetRules(v []RuleV1) {
	o.Rules = v
}

func (o ReplaceRulesInTrackingPlanV1Input) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["rules"] = o.Rules
	}
	return json.Marshal(toSerialize)
}

type NullableReplaceRulesInTrackingPlanV1Input struct {
	value *ReplaceRulesInTrackingPlanV1Input
	isSet bool
}

func (v NullableReplaceRulesInTrackingPlanV1Input) Get() *ReplaceRulesInTrackingPlanV1Input {
	return v.value
}

func (v *NullableReplaceRulesInTrackingPlanV1Input) Set(val *ReplaceRulesInTrackingPlanV1Input) {
	v.value = val
	v.isSet = true
}

func (v NullableReplaceRulesInTrackingPlanV1Input) IsSet() bool {
	return v.isSet
}

func (v *NullableReplaceRulesInTrackingPlanV1Input) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReplaceRulesInTrackingPlanV1Input(
	val *ReplaceRulesInTrackingPlanV1Input,
) *NullableReplaceRulesInTrackingPlanV1Input {
	return &NullableReplaceRulesInTrackingPlanV1Input{value: val, isSet: true}
}

func (v NullableReplaceRulesInTrackingPlanV1Input) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReplaceRulesInTrackingPlanV1Input) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
