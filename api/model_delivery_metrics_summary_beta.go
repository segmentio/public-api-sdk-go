/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 37.2.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// DeliveryMetricsSummaryBeta Defines the summary of delivery metrics for a Destination.
type DeliveryMetricsSummaryBeta struct {
	// The Source id.  Config API note: analogous to `parent`.
	SourceId string `json:"sourceId"`
	// The Destination metadata id.
	DestinationMetadataId string `json:"destinationMetadataId"`
	// The summary of event delivery metrics for the requested Destination.
	Metrics []MetricBeta `json:"metrics"`
}

// NewDeliveryMetricsSummaryBeta instantiates a new DeliveryMetricsSummaryBeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeliveryMetricsSummaryBeta(
	sourceId string,
	destinationMetadataId string,
	metrics []MetricBeta,
) *DeliveryMetricsSummaryBeta {
	this := DeliveryMetricsSummaryBeta{}
	this.SourceId = sourceId
	this.DestinationMetadataId = destinationMetadataId
	this.Metrics = metrics
	return &this
}

// NewDeliveryMetricsSummaryBetaWithDefaults instantiates a new DeliveryMetricsSummaryBeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeliveryMetricsSummaryBetaWithDefaults() *DeliveryMetricsSummaryBeta {
	this := DeliveryMetricsSummaryBeta{}
	return &this
}

// GetSourceId returns the SourceId field value
func (o *DeliveryMetricsSummaryBeta) GetSourceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceId
}

// GetSourceIdOk returns a tuple with the SourceId field value
// and a boolean to check if the value has been set.
func (o *DeliveryMetricsSummaryBeta) GetSourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceId, true
}

// SetSourceId sets field value
func (o *DeliveryMetricsSummaryBeta) SetSourceId(v string) {
	o.SourceId = v
}

// GetDestinationMetadataId returns the DestinationMetadataId field value
func (o *DeliveryMetricsSummaryBeta) GetDestinationMetadataId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DestinationMetadataId
}

// GetDestinationMetadataIdOk returns a tuple with the DestinationMetadataId field value
// and a boolean to check if the value has been set.
func (o *DeliveryMetricsSummaryBeta) GetDestinationMetadataIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DestinationMetadataId, true
}

// SetDestinationMetadataId sets field value
func (o *DeliveryMetricsSummaryBeta) SetDestinationMetadataId(v string) {
	o.DestinationMetadataId = v
}

// GetMetrics returns the Metrics field value
func (o *DeliveryMetricsSummaryBeta) GetMetrics() []MetricBeta {
	if o == nil {
		var ret []MetricBeta
		return ret
	}

	return o.Metrics
}

// GetMetricsOk returns a tuple with the Metrics field value
// and a boolean to check if the value has been set.
func (o *DeliveryMetricsSummaryBeta) GetMetricsOk() ([]MetricBeta, bool) {
	if o == nil {
		return nil, false
	}
	return o.Metrics, true
}

// SetMetrics sets field value
func (o *DeliveryMetricsSummaryBeta) SetMetrics(v []MetricBeta) {
	o.Metrics = v
}

func (o DeliveryMetricsSummaryBeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["sourceId"] = o.SourceId
	}
	if true {
		toSerialize["destinationMetadataId"] = o.DestinationMetadataId
	}
	if true {
		toSerialize["metrics"] = o.Metrics
	}
	return json.Marshal(toSerialize)
}

type NullableDeliveryMetricsSummaryBeta struct {
	value *DeliveryMetricsSummaryBeta
	isSet bool
}

func (v NullableDeliveryMetricsSummaryBeta) Get() *DeliveryMetricsSummaryBeta {
	return v.value
}

func (v *NullableDeliveryMetricsSummaryBeta) Set(val *DeliveryMetricsSummaryBeta) {
	v.value = val
	v.isSet = true
}

func (v NullableDeliveryMetricsSummaryBeta) IsSet() bool {
	return v.isSet
}

func (v *NullableDeliveryMetricsSummaryBeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeliveryMetricsSummaryBeta(
	val *DeliveryMetricsSummaryBeta,
) *NullableDeliveryMetricsSummaryBeta {
	return &NullableDeliveryMetricsSummaryBeta{value: val, isSet: true}
}

func (v NullableDeliveryMetricsSummaryBeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeliveryMetricsSummaryBeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
