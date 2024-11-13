/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 57.0.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the MetricBeta type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetricBeta{}

// MetricBeta The event delivery metric.
type MetricBeta struct {
	// The name of the metric.
	MetricName string `json:"metricName"`
	// Number of occurrences of the metric.
	Total float32 `json:"total"`
	// Breakdown of metrics within a metric.
	Breakdown []BreakdownBeta `json:"breakdown,omitempty"`
}

// NewMetricBeta instantiates a new MetricBeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricBeta(metricName string, total float32) *MetricBeta {
	this := MetricBeta{}
	this.MetricName = metricName
	this.Total = total
	return &this
}

// NewMetricBetaWithDefaults instantiates a new MetricBeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricBetaWithDefaults() *MetricBeta {
	this := MetricBeta{}
	return &this
}

// GetMetricName returns the MetricName field value
func (o *MetricBeta) GetMetricName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MetricName
}

// GetMetricNameOk returns a tuple with the MetricName field value
// and a boolean to check if the value has been set.
func (o *MetricBeta) GetMetricNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetricName, true
}

// SetMetricName sets field value
func (o *MetricBeta) SetMetricName(v string) {
	o.MetricName = v
}

// GetTotal returns the Total field value
func (o *MetricBeta) GetTotal() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *MetricBeta) GetTotalOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *MetricBeta) SetTotal(v float32) {
	o.Total = v
}

// GetBreakdown returns the Breakdown field value if set, zero value otherwise.
func (o *MetricBeta) GetBreakdown() []BreakdownBeta {
	if o == nil || IsNil(o.Breakdown) {
		var ret []BreakdownBeta
		return ret
	}
	return o.Breakdown
}

// GetBreakdownOk returns a tuple with the Breakdown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricBeta) GetBreakdownOk() ([]BreakdownBeta, bool) {
	if o == nil || IsNil(o.Breakdown) {
		return nil, false
	}
	return o.Breakdown, true
}

// HasBreakdown returns a boolean if a field has been set.
func (o *MetricBeta) HasBreakdown() bool {
	if o != nil && !IsNil(o.Breakdown) {
		return true
	}

	return false
}

// SetBreakdown gets a reference to the given []BreakdownBeta and assigns it to the Breakdown field.
func (o *MetricBeta) SetBreakdown(v []BreakdownBeta) {
	o.Breakdown = v
}

func (o MetricBeta) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetricBeta) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["metricName"] = o.MetricName
	toSerialize["total"] = o.Total
	if !IsNil(o.Breakdown) {
		toSerialize["breakdown"] = o.Breakdown
	}
	return toSerialize, nil
}

type NullableMetricBeta struct {
	value *MetricBeta
	isSet bool
}

func (v NullableMetricBeta) Get() *MetricBeta {
	return v.value
}

func (v *NullableMetricBeta) Set(val *MetricBeta) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricBeta) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricBeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricBeta(val *MetricBeta) *NullableMetricBeta {
	return &NullableMetricBeta{value: val, isSet: true}
}

func (v NullableMetricBeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricBeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
