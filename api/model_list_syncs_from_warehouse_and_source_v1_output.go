/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 38.0.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the ListSyncsFromWarehouseAndSourceV1Output type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListSyncsFromWarehouseAndSourceV1Output{}

// ListSyncsFromWarehouseAndSourceV1Output Returns a list that contains the most recent syncs for a Warehouse-source pair, filtered and constrained by a given set of pagination parameters.
type ListSyncsFromWarehouseAndSourceV1Output struct {
	// A list that contains the latest syncs for the specified Warehouse-source pair.
	Reports    []SyncV1         `json:"reports"`
	Pagination PaginationOutput `json:"pagination"`
}

// NewListSyncsFromWarehouseAndSourceV1Output instantiates a new ListSyncsFromWarehouseAndSourceV1Output object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListSyncsFromWarehouseAndSourceV1Output(
	reports []SyncV1,
	pagination PaginationOutput,
) *ListSyncsFromWarehouseAndSourceV1Output {
	this := ListSyncsFromWarehouseAndSourceV1Output{}
	this.Reports = reports
	this.Pagination = pagination
	return &this
}

// NewListSyncsFromWarehouseAndSourceV1OutputWithDefaults instantiates a new ListSyncsFromWarehouseAndSourceV1Output object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListSyncsFromWarehouseAndSourceV1OutputWithDefaults() *ListSyncsFromWarehouseAndSourceV1Output {
	this := ListSyncsFromWarehouseAndSourceV1Output{}
	return &this
}

// GetReports returns the Reports field value
func (o *ListSyncsFromWarehouseAndSourceV1Output) GetReports() []SyncV1 {
	if o == nil {
		var ret []SyncV1
		return ret
	}

	return o.Reports
}

// GetReportsOk returns a tuple with the Reports field value
// and a boolean to check if the value has been set.
func (o *ListSyncsFromWarehouseAndSourceV1Output) GetReportsOk() ([]SyncV1, bool) {
	if o == nil {
		return nil, false
	}
	return o.Reports, true
}

// SetReports sets field value
func (o *ListSyncsFromWarehouseAndSourceV1Output) SetReports(v []SyncV1) {
	o.Reports = v
}

// GetPagination returns the Pagination field value
func (o *ListSyncsFromWarehouseAndSourceV1Output) GetPagination() PaginationOutput {
	if o == nil {
		var ret PaginationOutput
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *ListSyncsFromWarehouseAndSourceV1Output) GetPaginationOk() (*PaginationOutput, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *ListSyncsFromWarehouseAndSourceV1Output) SetPagination(v PaginationOutput) {
	o.Pagination = v
}

func (o ListSyncsFromWarehouseAndSourceV1Output) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListSyncsFromWarehouseAndSourceV1Output) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["reports"] = o.Reports
	toSerialize["pagination"] = o.Pagination
	return toSerialize, nil
}

type NullableListSyncsFromWarehouseAndSourceV1Output struct {
	value *ListSyncsFromWarehouseAndSourceV1Output
	isSet bool
}

func (v NullableListSyncsFromWarehouseAndSourceV1Output) Get() *ListSyncsFromWarehouseAndSourceV1Output {
	return v.value
}

func (v *NullableListSyncsFromWarehouseAndSourceV1Output) Set(
	val *ListSyncsFromWarehouseAndSourceV1Output,
) {
	v.value = val
	v.isSet = true
}

func (v NullableListSyncsFromWarehouseAndSourceV1Output) IsSet() bool {
	return v.isSet
}

func (v *NullableListSyncsFromWarehouseAndSourceV1Output) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListSyncsFromWarehouseAndSourceV1Output(
	val *ListSyncsFromWarehouseAndSourceV1Output,
) *NullableListSyncsFromWarehouseAndSourceV1Output {
	return &NullableListSyncsFromWarehouseAndSourceV1Output{value: val, isSet: true}
}

func (v NullableListSyncsFromWarehouseAndSourceV1Output) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListSyncsFromWarehouseAndSourceV1Output) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
