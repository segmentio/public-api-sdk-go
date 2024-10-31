/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 56.0.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the GetDailyWorkspaceAPICallsUsageV1Output type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetDailyWorkspaceAPICallsUsageV1Output{}

// GetDailyWorkspaceAPICallsUsageV1Output Returns a list of daily aggregations of Workspace MTU counts.
type GetDailyWorkspaceAPICallsUsageV1Output struct {
	// The list of daily Workspace API calls count aggregates.
	DailyWorkspaceAPICallsUsage []APICallSnapshotV1 `json:"dailyWorkspaceAPICallsUsage"`
	Pagination                  PaginationOutput    `json:"pagination"`
}

// NewGetDailyWorkspaceAPICallsUsageV1Output instantiates a new GetDailyWorkspaceAPICallsUsageV1Output object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDailyWorkspaceAPICallsUsageV1Output(
	dailyWorkspaceAPICallsUsage []APICallSnapshotV1,
	pagination PaginationOutput,
) *GetDailyWorkspaceAPICallsUsageV1Output {
	this := GetDailyWorkspaceAPICallsUsageV1Output{}
	this.DailyWorkspaceAPICallsUsage = dailyWorkspaceAPICallsUsage
	this.Pagination = pagination
	return &this
}

// NewGetDailyWorkspaceAPICallsUsageV1OutputWithDefaults instantiates a new GetDailyWorkspaceAPICallsUsageV1Output object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDailyWorkspaceAPICallsUsageV1OutputWithDefaults() *GetDailyWorkspaceAPICallsUsageV1Output {
	this := GetDailyWorkspaceAPICallsUsageV1Output{}
	return &this
}

// GetDailyWorkspaceAPICallsUsage returns the DailyWorkspaceAPICallsUsage field value
func (o *GetDailyWorkspaceAPICallsUsageV1Output) GetDailyWorkspaceAPICallsUsage() []APICallSnapshotV1 {
	if o == nil {
		var ret []APICallSnapshotV1
		return ret
	}

	return o.DailyWorkspaceAPICallsUsage
}

// GetDailyWorkspaceAPICallsUsageOk returns a tuple with the DailyWorkspaceAPICallsUsage field value
// and a boolean to check if the value has been set.
func (o *GetDailyWorkspaceAPICallsUsageV1Output) GetDailyWorkspaceAPICallsUsageOk() ([]APICallSnapshotV1, bool) {
	if o == nil {
		return nil, false
	}
	return o.DailyWorkspaceAPICallsUsage, true
}

// SetDailyWorkspaceAPICallsUsage sets field value
func (o *GetDailyWorkspaceAPICallsUsageV1Output) SetDailyWorkspaceAPICallsUsage(
	v []APICallSnapshotV1,
) {
	o.DailyWorkspaceAPICallsUsage = v
}

// GetPagination returns the Pagination field value
func (o *GetDailyWorkspaceAPICallsUsageV1Output) GetPagination() PaginationOutput {
	if o == nil {
		var ret PaginationOutput
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *GetDailyWorkspaceAPICallsUsageV1Output) GetPaginationOk() (*PaginationOutput, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *GetDailyWorkspaceAPICallsUsageV1Output) SetPagination(v PaginationOutput) {
	o.Pagination = v
}

func (o GetDailyWorkspaceAPICallsUsageV1Output) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetDailyWorkspaceAPICallsUsageV1Output) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dailyWorkspaceAPICallsUsage"] = o.DailyWorkspaceAPICallsUsage
	toSerialize["pagination"] = o.Pagination
	return toSerialize, nil
}

type NullableGetDailyWorkspaceAPICallsUsageV1Output struct {
	value *GetDailyWorkspaceAPICallsUsageV1Output
	isSet bool
}

func (v NullableGetDailyWorkspaceAPICallsUsageV1Output) Get() *GetDailyWorkspaceAPICallsUsageV1Output {
	return v.value
}

func (v *NullableGetDailyWorkspaceAPICallsUsageV1Output) Set(
	val *GetDailyWorkspaceAPICallsUsageV1Output,
) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDailyWorkspaceAPICallsUsageV1Output) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDailyWorkspaceAPICallsUsageV1Output) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDailyWorkspaceAPICallsUsageV1Output(
	val *GetDailyWorkspaceAPICallsUsageV1Output,
) *NullableGetDailyWorkspaceAPICallsUsageV1Output {
	return &NullableGetDailyWorkspaceAPICallsUsageV1Output{value: val, isSet: true}
}

func (v NullableGetDailyWorkspaceAPICallsUsageV1Output) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDailyWorkspaceAPICallsUsageV1Output) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
