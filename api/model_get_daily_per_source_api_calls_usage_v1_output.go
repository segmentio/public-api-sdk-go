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

// checks if the GetDailyPerSourceAPICallsUsageV1Output type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetDailyPerSourceAPICallsUsageV1Output{}

// GetDailyPerSourceAPICallsUsageV1Output Returns a list of daily aggregations of Source level API calls counts.
type GetDailyPerSourceAPICallsUsageV1Output struct {
	// The list of daily per Source API calls count aggregates.
	DailyPerSourceAPICallsUsage []SourceAPICallSnapshotV1 `json:"dailyPerSourceAPICallsUsage"`
	Pagination                  PaginationOutput          `json:"pagination"`
}

// NewGetDailyPerSourceAPICallsUsageV1Output instantiates a new GetDailyPerSourceAPICallsUsageV1Output object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDailyPerSourceAPICallsUsageV1Output(
	dailyPerSourceAPICallsUsage []SourceAPICallSnapshotV1,
	pagination PaginationOutput,
) *GetDailyPerSourceAPICallsUsageV1Output {
	this := GetDailyPerSourceAPICallsUsageV1Output{}
	this.DailyPerSourceAPICallsUsage = dailyPerSourceAPICallsUsage
	this.Pagination = pagination
	return &this
}

// NewGetDailyPerSourceAPICallsUsageV1OutputWithDefaults instantiates a new GetDailyPerSourceAPICallsUsageV1Output object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDailyPerSourceAPICallsUsageV1OutputWithDefaults() *GetDailyPerSourceAPICallsUsageV1Output {
	this := GetDailyPerSourceAPICallsUsageV1Output{}
	return &this
}

// GetDailyPerSourceAPICallsUsage returns the DailyPerSourceAPICallsUsage field value
func (o *GetDailyPerSourceAPICallsUsageV1Output) GetDailyPerSourceAPICallsUsage() []SourceAPICallSnapshotV1 {
	if o == nil {
		var ret []SourceAPICallSnapshotV1
		return ret
	}

	return o.DailyPerSourceAPICallsUsage
}

// GetDailyPerSourceAPICallsUsageOk returns a tuple with the DailyPerSourceAPICallsUsage field value
// and a boolean to check if the value has been set.
func (o *GetDailyPerSourceAPICallsUsageV1Output) GetDailyPerSourceAPICallsUsageOk() ([]SourceAPICallSnapshotV1, bool) {
	if o == nil {
		return nil, false
	}
	return o.DailyPerSourceAPICallsUsage, true
}

// SetDailyPerSourceAPICallsUsage sets field value
func (o *GetDailyPerSourceAPICallsUsageV1Output) SetDailyPerSourceAPICallsUsage(
	v []SourceAPICallSnapshotV1,
) {
	o.DailyPerSourceAPICallsUsage = v
}

// GetPagination returns the Pagination field value
func (o *GetDailyPerSourceAPICallsUsageV1Output) GetPagination() PaginationOutput {
	if o == nil {
		var ret PaginationOutput
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *GetDailyPerSourceAPICallsUsageV1Output) GetPaginationOk() (*PaginationOutput, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *GetDailyPerSourceAPICallsUsageV1Output) SetPagination(v PaginationOutput) {
	o.Pagination = v
}

func (o GetDailyPerSourceAPICallsUsageV1Output) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetDailyPerSourceAPICallsUsageV1Output) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dailyPerSourceAPICallsUsage"] = o.DailyPerSourceAPICallsUsage
	toSerialize["pagination"] = o.Pagination
	return toSerialize, nil
}

type NullableGetDailyPerSourceAPICallsUsageV1Output struct {
	value *GetDailyPerSourceAPICallsUsageV1Output
	isSet bool
}

func (v NullableGetDailyPerSourceAPICallsUsageV1Output) Get() *GetDailyPerSourceAPICallsUsageV1Output {
	return v.value
}

func (v *NullableGetDailyPerSourceAPICallsUsageV1Output) Set(
	val *GetDailyPerSourceAPICallsUsageV1Output,
) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDailyPerSourceAPICallsUsageV1Output) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDailyPerSourceAPICallsUsageV1Output) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDailyPerSourceAPICallsUsageV1Output(
	val *GetDailyPerSourceAPICallsUsageV1Output,
) *NullableGetDailyPerSourceAPICallsUsageV1Output {
	return &NullableGetDailyPerSourceAPICallsUsageV1Output{value: val, isSet: true}
}

func (v NullableGetDailyPerSourceAPICallsUsageV1Output) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDailyPerSourceAPICallsUsageV1Output) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
