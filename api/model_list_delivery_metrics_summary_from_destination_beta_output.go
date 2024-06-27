/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 50.2.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the ListDeliveryMetricsSummaryFromDestinationBetaOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListDeliveryMetricsSummaryFromDestinationBetaOutput{}

// ListDeliveryMetricsSummaryFromDestinationBetaOutput Output to retrieve event delivery metrics summary for a Destination.
type ListDeliveryMetricsSummaryFromDestinationBetaOutput struct {
	DeliveryMetricsSummary DeliveryMetricsSummaryBeta `json:"deliveryMetricsSummary"`
}

// NewListDeliveryMetricsSummaryFromDestinationBetaOutput instantiates a new ListDeliveryMetricsSummaryFromDestinationBetaOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListDeliveryMetricsSummaryFromDestinationBetaOutput(
	deliveryMetricsSummary DeliveryMetricsSummaryBeta,
) *ListDeliveryMetricsSummaryFromDestinationBetaOutput {
	this := ListDeliveryMetricsSummaryFromDestinationBetaOutput{}
	this.DeliveryMetricsSummary = deliveryMetricsSummary
	return &this
}

// NewListDeliveryMetricsSummaryFromDestinationBetaOutputWithDefaults instantiates a new ListDeliveryMetricsSummaryFromDestinationBetaOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListDeliveryMetricsSummaryFromDestinationBetaOutputWithDefaults() *ListDeliveryMetricsSummaryFromDestinationBetaOutput {
	this := ListDeliveryMetricsSummaryFromDestinationBetaOutput{}
	return &this
}

// GetDeliveryMetricsSummary returns the DeliveryMetricsSummary field value
func (o *ListDeliveryMetricsSummaryFromDestinationBetaOutput) GetDeliveryMetricsSummary() DeliveryMetricsSummaryBeta {
	if o == nil {
		var ret DeliveryMetricsSummaryBeta
		return ret
	}

	return o.DeliveryMetricsSummary
}

// GetDeliveryMetricsSummaryOk returns a tuple with the DeliveryMetricsSummary field value
// and a boolean to check if the value has been set.
func (o *ListDeliveryMetricsSummaryFromDestinationBetaOutput) GetDeliveryMetricsSummaryOk() (*DeliveryMetricsSummaryBeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeliveryMetricsSummary, true
}

// SetDeliveryMetricsSummary sets field value
func (o *ListDeliveryMetricsSummaryFromDestinationBetaOutput) SetDeliveryMetricsSummary(
	v DeliveryMetricsSummaryBeta,
) {
	o.DeliveryMetricsSummary = v
}

func (o ListDeliveryMetricsSummaryFromDestinationBetaOutput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListDeliveryMetricsSummaryFromDestinationBetaOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["deliveryMetricsSummary"] = o.DeliveryMetricsSummary
	return toSerialize, nil
}

type NullableListDeliveryMetricsSummaryFromDestinationBetaOutput struct {
	value *ListDeliveryMetricsSummaryFromDestinationBetaOutput
	isSet bool
}

func (v NullableListDeliveryMetricsSummaryFromDestinationBetaOutput) Get() *ListDeliveryMetricsSummaryFromDestinationBetaOutput {
	return v.value
}

func (v *NullableListDeliveryMetricsSummaryFromDestinationBetaOutput) Set(
	val *ListDeliveryMetricsSummaryFromDestinationBetaOutput,
) {
	v.value = val
	v.isSet = true
}

func (v NullableListDeliveryMetricsSummaryFromDestinationBetaOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableListDeliveryMetricsSummaryFromDestinationBetaOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListDeliveryMetricsSummaryFromDestinationBetaOutput(
	val *ListDeliveryMetricsSummaryFromDestinationBetaOutput,
) *NullableListDeliveryMetricsSummaryFromDestinationBetaOutput {
	return &NullableListDeliveryMetricsSummaryFromDestinationBetaOutput{value: val, isSet: true}
}

func (v NullableListDeliveryMetricsSummaryFromDestinationBetaOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListDeliveryMetricsSummaryFromDestinationBetaOutput) UnmarshalJSON(
	src []byte,
) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
