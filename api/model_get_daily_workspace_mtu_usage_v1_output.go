/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 35.0.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// GetDailyWorkspaceMTUUsageV1Output Returns a list of daily aggregations of Workspace MTU counts.
type GetDailyWorkspaceMTUUsageV1Output struct {
	// The list of daily Workspace MTU count aggregates.
	DailyWorkspaceMTUUsage []MtuSnapshotV1 `json:"dailyWorkspaceMTUUsage"`
	Pagination             Pagination      `json:"pagination"`
}

// NewGetDailyWorkspaceMTUUsageV1Output instantiates a new GetDailyWorkspaceMTUUsageV1Output object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDailyWorkspaceMTUUsageV1Output(
	dailyWorkspaceMTUUsage []MtuSnapshotV1,
	pagination Pagination,
) *GetDailyWorkspaceMTUUsageV1Output {
	this := GetDailyWorkspaceMTUUsageV1Output{}
	this.DailyWorkspaceMTUUsage = dailyWorkspaceMTUUsage
	this.Pagination = pagination
	return &this
}

// NewGetDailyWorkspaceMTUUsageV1OutputWithDefaults instantiates a new GetDailyWorkspaceMTUUsageV1Output object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDailyWorkspaceMTUUsageV1OutputWithDefaults() *GetDailyWorkspaceMTUUsageV1Output {
	this := GetDailyWorkspaceMTUUsageV1Output{}
	return &this
}

// GetDailyWorkspaceMTUUsage returns the DailyWorkspaceMTUUsage field value
func (o *GetDailyWorkspaceMTUUsageV1Output) GetDailyWorkspaceMTUUsage() []MtuSnapshotV1 {
	if o == nil {
		var ret []MtuSnapshotV1
		return ret
	}

	return o.DailyWorkspaceMTUUsage
}

// GetDailyWorkspaceMTUUsageOk returns a tuple with the DailyWorkspaceMTUUsage field value
// and a boolean to check if the value has been set.
func (o *GetDailyWorkspaceMTUUsageV1Output) GetDailyWorkspaceMTUUsageOk() ([]MtuSnapshotV1, bool) {
	if o == nil {
		return nil, false
	}
	return o.DailyWorkspaceMTUUsage, true
}

// SetDailyWorkspaceMTUUsage sets field value
func (o *GetDailyWorkspaceMTUUsageV1Output) SetDailyWorkspaceMTUUsage(v []MtuSnapshotV1) {
	o.DailyWorkspaceMTUUsage = v
}

// GetPagination returns the Pagination field value
func (o *GetDailyWorkspaceMTUUsageV1Output) GetPagination() Pagination {
	if o == nil {
		var ret Pagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *GetDailyWorkspaceMTUUsageV1Output) GetPaginationOk() (*Pagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *GetDailyWorkspaceMTUUsageV1Output) SetPagination(v Pagination) {
	o.Pagination = v
}

func (o GetDailyWorkspaceMTUUsageV1Output) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["dailyWorkspaceMTUUsage"] = o.DailyWorkspaceMTUUsage
	}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	return json.Marshal(toSerialize)
}

type NullableGetDailyWorkspaceMTUUsageV1Output struct {
	value *GetDailyWorkspaceMTUUsageV1Output
	isSet bool
}

func (v NullableGetDailyWorkspaceMTUUsageV1Output) Get() *GetDailyWorkspaceMTUUsageV1Output {
	return v.value
}

func (v *NullableGetDailyWorkspaceMTUUsageV1Output) Set(val *GetDailyWorkspaceMTUUsageV1Output) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDailyWorkspaceMTUUsageV1Output) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDailyWorkspaceMTUUsageV1Output) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDailyWorkspaceMTUUsageV1Output(
	val *GetDailyWorkspaceMTUUsageV1Output,
) *NullableGetDailyWorkspaceMTUUsageV1Output {
	return &NullableGetDailyWorkspaceMTUUsageV1Output{value: val, isSet: true}
}

func (v NullableGetDailyWorkspaceMTUUsageV1Output) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDailyWorkspaceMTUUsageV1Output) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
