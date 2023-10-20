/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 36.2.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// BreakdownBeta The breakdown of a metric.
type BreakdownBeta struct {
	// The name of the metric.
	MetricName string `json:"metricName"`
	// Number of occurrences of the metric.
	Value float32 `json:"value"`
}

// NewBreakdownBeta instantiates a new BreakdownBeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBreakdownBeta(metricName string, value float32) *BreakdownBeta {
	this := BreakdownBeta{}
	this.MetricName = metricName
	this.Value = value
	return &this
}

// NewBreakdownBetaWithDefaults instantiates a new BreakdownBeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBreakdownBetaWithDefaults() *BreakdownBeta {
	this := BreakdownBeta{}
	return &this
}

// GetMetricName returns the MetricName field value
func (o *BreakdownBeta) GetMetricName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MetricName
}

// GetMetricNameOk returns a tuple with the MetricName field value
// and a boolean to check if the value has been set.
func (o *BreakdownBeta) GetMetricNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetricName, true
}

// SetMetricName sets field value
func (o *BreakdownBeta) SetMetricName(v string) {
	o.MetricName = v
}

// GetValue returns the Value field value
func (o *BreakdownBeta) GetValue() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *BreakdownBeta) GetValueOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *BreakdownBeta) SetValue(v float32) {
	o.Value = v
}

func (o BreakdownBeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["metricName"] = o.MetricName
	}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableBreakdownBeta struct {
	value *BreakdownBeta
	isSet bool
}

func (v NullableBreakdownBeta) Get() *BreakdownBeta {
	return v.value
}

func (v *NullableBreakdownBeta) Set(val *BreakdownBeta) {
	v.value = val
	v.isSet = true
}

func (v NullableBreakdownBeta) IsSet() bool {
	return v.isSet
}

func (v *NullableBreakdownBeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBreakdownBeta(val *BreakdownBeta) *NullableBreakdownBeta {
	return &NullableBreakdownBeta{value: val, isSet: true}
}

func (v NullableBreakdownBeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBreakdownBeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
