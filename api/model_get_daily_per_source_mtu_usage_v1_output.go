/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 58.13.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the GetDailyPerSourceMTUUsageV1Output type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetDailyPerSourceMTUUsageV1Output{}

// GetDailyPerSourceMTUUsageV1Output Returns a list of daily aggregations of Source level MTU counts.
type GetDailyPerSourceMTUUsageV1Output struct {
	// The list of daily per Source MTU count aggregates.
	DailyPerSourceMTUUsage []UsersPerSourceSnapshotV1 `json:"dailyPerSourceMTUUsage"`
	Pagination             PaginationOutput           `json:"pagination"`
}

// NewGetDailyPerSourceMTUUsageV1Output instantiates a new GetDailyPerSourceMTUUsageV1Output object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDailyPerSourceMTUUsageV1Output(
	dailyPerSourceMTUUsage []UsersPerSourceSnapshotV1,
	pagination PaginationOutput,
) *GetDailyPerSourceMTUUsageV1Output {
	this := GetDailyPerSourceMTUUsageV1Output{}
	this.DailyPerSourceMTUUsage = dailyPerSourceMTUUsage
	this.Pagination = pagination
	return &this
}

// NewGetDailyPerSourceMTUUsageV1OutputWithDefaults instantiates a new GetDailyPerSourceMTUUsageV1Output object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDailyPerSourceMTUUsageV1OutputWithDefaults() *GetDailyPerSourceMTUUsageV1Output {
	this := GetDailyPerSourceMTUUsageV1Output{}
	return &this
}

// GetDailyPerSourceMTUUsage returns the DailyPerSourceMTUUsage field value
func (o *GetDailyPerSourceMTUUsageV1Output) GetDailyPerSourceMTUUsage() []UsersPerSourceSnapshotV1 {
	if o == nil {
		var ret []UsersPerSourceSnapshotV1
		return ret
	}

	return o.DailyPerSourceMTUUsage
}

// GetDailyPerSourceMTUUsageOk returns a tuple with the DailyPerSourceMTUUsage field value
// and a boolean to check if the value has been set.
func (o *GetDailyPerSourceMTUUsageV1Output) GetDailyPerSourceMTUUsageOk() ([]UsersPerSourceSnapshotV1, bool) {
	if o == nil {
		return nil, false
	}
	return o.DailyPerSourceMTUUsage, true
}

// SetDailyPerSourceMTUUsage sets field value
func (o *GetDailyPerSourceMTUUsageV1Output) SetDailyPerSourceMTUUsage(
	v []UsersPerSourceSnapshotV1,
) {
	o.DailyPerSourceMTUUsage = v
}

// GetPagination returns the Pagination field value
func (o *GetDailyPerSourceMTUUsageV1Output) GetPagination() PaginationOutput {
	if o == nil {
		var ret PaginationOutput
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *GetDailyPerSourceMTUUsageV1Output) GetPaginationOk() (*PaginationOutput, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *GetDailyPerSourceMTUUsageV1Output) SetPagination(v PaginationOutput) {
	o.Pagination = v
}

func (o GetDailyPerSourceMTUUsageV1Output) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetDailyPerSourceMTUUsageV1Output) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dailyPerSourceMTUUsage"] = o.DailyPerSourceMTUUsage
	toSerialize["pagination"] = o.Pagination
	return toSerialize, nil
}

type NullableGetDailyPerSourceMTUUsageV1Output struct {
	value *GetDailyPerSourceMTUUsageV1Output
	isSet bool
}

func (v NullableGetDailyPerSourceMTUUsageV1Output) Get() *GetDailyPerSourceMTUUsageV1Output {
	return v.value
}

func (v *NullableGetDailyPerSourceMTUUsageV1Output) Set(val *GetDailyPerSourceMTUUsageV1Output) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDailyPerSourceMTUUsageV1Output) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDailyPerSourceMTUUsageV1Output) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDailyPerSourceMTUUsageV1Output(
	val *GetDailyPerSourceMTUUsageV1Output,
) *NullableGetDailyPerSourceMTUUsageV1Output {
	return &NullableGetDailyPerSourceMTUUsageV1Output{value: val, isSet: true}
}

func (v NullableGetDailyPerSourceMTUUsageV1Output) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDailyPerSourceMTUUsageV1Output) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
