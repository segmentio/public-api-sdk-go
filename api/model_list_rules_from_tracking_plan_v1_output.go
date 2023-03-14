/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 34.3.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// ListRulesFromTrackingPlanV1Output Lists a Tracking Plan's rules.
type ListRulesFromTrackingPlanV1Output struct {
	// Rules associated with the given Tracking Plan.
	Rules      []RuleV1   `json:"rules"`
	Pagination Pagination `json:"pagination"`
}

// NewListRulesFromTrackingPlanV1Output instantiates a new ListRulesFromTrackingPlanV1Output object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListRulesFromTrackingPlanV1Output(
	rules []RuleV1,
	pagination Pagination,
) *ListRulesFromTrackingPlanV1Output {
	this := ListRulesFromTrackingPlanV1Output{}
	this.Rules = rules
	this.Pagination = pagination
	return &this
}

// NewListRulesFromTrackingPlanV1OutputWithDefaults instantiates a new ListRulesFromTrackingPlanV1Output object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListRulesFromTrackingPlanV1OutputWithDefaults() *ListRulesFromTrackingPlanV1Output {
	this := ListRulesFromTrackingPlanV1Output{}
	return &this
}

// GetRules returns the Rules field value
func (o *ListRulesFromTrackingPlanV1Output) GetRules() []RuleV1 {
	if o == nil {
		var ret []RuleV1
		return ret
	}

	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value
// and a boolean to check if the value has been set.
func (o *ListRulesFromTrackingPlanV1Output) GetRulesOk() ([]RuleV1, bool) {
	if o == nil {
		return nil, false
	}
	return o.Rules, true
}

// SetRules sets field value
func (o *ListRulesFromTrackingPlanV1Output) SetRules(v []RuleV1) {
	o.Rules = v
}

// GetPagination returns the Pagination field value
func (o *ListRulesFromTrackingPlanV1Output) GetPagination() Pagination {
	if o == nil {
		var ret Pagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *ListRulesFromTrackingPlanV1Output) GetPaginationOk() (*Pagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *ListRulesFromTrackingPlanV1Output) SetPagination(v Pagination) {
	o.Pagination = v
}

func (o ListRulesFromTrackingPlanV1Output) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["rules"] = o.Rules
	}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	return json.Marshal(toSerialize)
}

type NullableListRulesFromTrackingPlanV1Output struct {
	value *ListRulesFromTrackingPlanV1Output
	isSet bool
}

func (v NullableListRulesFromTrackingPlanV1Output) Get() *ListRulesFromTrackingPlanV1Output {
	return v.value
}

func (v *NullableListRulesFromTrackingPlanV1Output) Set(val *ListRulesFromTrackingPlanV1Output) {
	v.value = val
	v.isSet = true
}

func (v NullableListRulesFromTrackingPlanV1Output) IsSet() bool {
	return v.isSet
}

func (v *NullableListRulesFromTrackingPlanV1Output) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListRulesFromTrackingPlanV1Output(
	val *ListRulesFromTrackingPlanV1Output,
) *NullableListRulesFromTrackingPlanV1Output {
	return &NullableListRulesFromTrackingPlanV1Output{value: val, isSet: true}
}

func (v NullableListRulesFromTrackingPlanV1Output) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListRulesFromTrackingPlanV1Output) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
