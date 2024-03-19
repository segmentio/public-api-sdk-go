/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 46.0.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the GetDeliveryOverviewMetricsAlphaOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetDeliveryOverviewMetricsAlphaOutput{}

// GetDeliveryOverviewMetricsAlphaOutput Output of the Delivery Overview public api endpoints.
type GetDeliveryOverviewMetricsAlphaOutput struct {
	// The total number of events for the returned dataset.
	Total float32 `json:"total"`
	// Represents the list of series broken down by the dimensions and timeframe requested.
	Dataset    []DeliveryOverviewMetricsDataset `json:"dataset"`
	Pagination PaginationOutput                 `json:"pagination"`
}

// NewGetDeliveryOverviewMetricsAlphaOutput instantiates a new GetDeliveryOverviewMetricsAlphaOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDeliveryOverviewMetricsAlphaOutput(
	total float32,
	dataset []DeliveryOverviewMetricsDataset,
	pagination PaginationOutput,
) *GetDeliveryOverviewMetricsAlphaOutput {
	this := GetDeliveryOverviewMetricsAlphaOutput{}
	this.Total = total
	this.Dataset = dataset
	this.Pagination = pagination
	return &this
}

// NewGetDeliveryOverviewMetricsAlphaOutputWithDefaults instantiates a new GetDeliveryOverviewMetricsAlphaOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDeliveryOverviewMetricsAlphaOutputWithDefaults() *GetDeliveryOverviewMetricsAlphaOutput {
	this := GetDeliveryOverviewMetricsAlphaOutput{}
	return &this
}

// GetTotal returns the Total field value
func (o *GetDeliveryOverviewMetricsAlphaOutput) GetTotal() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *GetDeliveryOverviewMetricsAlphaOutput) GetTotalOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *GetDeliveryOverviewMetricsAlphaOutput) SetTotal(v float32) {
	o.Total = v
}

// GetDataset returns the Dataset field value
func (o *GetDeliveryOverviewMetricsAlphaOutput) GetDataset() []DeliveryOverviewMetricsDataset {
	if o == nil {
		var ret []DeliveryOverviewMetricsDataset
		return ret
	}

	return o.Dataset
}

// GetDatasetOk returns a tuple with the Dataset field value
// and a boolean to check if the value has been set.
func (o *GetDeliveryOverviewMetricsAlphaOutput) GetDatasetOk() ([]DeliveryOverviewMetricsDataset, bool) {
	if o == nil {
		return nil, false
	}
	return o.Dataset, true
}

// SetDataset sets field value
func (o *GetDeliveryOverviewMetricsAlphaOutput) SetDataset(v []DeliveryOverviewMetricsDataset) {
	o.Dataset = v
}

// GetPagination returns the Pagination field value
func (o *GetDeliveryOverviewMetricsAlphaOutput) GetPagination() PaginationOutput {
	if o == nil {
		var ret PaginationOutput
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *GetDeliveryOverviewMetricsAlphaOutput) GetPaginationOk() (*PaginationOutput, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *GetDeliveryOverviewMetricsAlphaOutput) SetPagination(v PaginationOutput) {
	o.Pagination = v
}

func (o GetDeliveryOverviewMetricsAlphaOutput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetDeliveryOverviewMetricsAlphaOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["total"] = o.Total
	toSerialize["dataset"] = o.Dataset
	toSerialize["pagination"] = o.Pagination
	return toSerialize, nil
}

type NullableGetDeliveryOverviewMetricsAlphaOutput struct {
	value *GetDeliveryOverviewMetricsAlphaOutput
	isSet bool
}

func (v NullableGetDeliveryOverviewMetricsAlphaOutput) Get() *GetDeliveryOverviewMetricsAlphaOutput {
	return v.value
}

func (v *NullableGetDeliveryOverviewMetricsAlphaOutput) Set(
	val *GetDeliveryOverviewMetricsAlphaOutput,
) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDeliveryOverviewMetricsAlphaOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDeliveryOverviewMetricsAlphaOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDeliveryOverviewMetricsAlphaOutput(
	val *GetDeliveryOverviewMetricsAlphaOutput,
) *NullableGetDeliveryOverviewMetricsAlphaOutput {
	return &NullableGetDeliveryOverviewMetricsAlphaOutput{value: val, isSet: true}
}

func (v NullableGetDeliveryOverviewMetricsAlphaOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDeliveryOverviewMetricsAlphaOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
