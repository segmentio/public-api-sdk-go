/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 36.3.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// ListSyncsFromWarehouseV1Output Returns an overview page that contains the last syncs for a Warehouse.
type ListSyncsFromWarehouseV1Output struct {
	// A list that contains the latest syncs for the specified Warehouse.
	Reports    []SyncV1   `json:"reports"`
	Pagination Pagination `json:"pagination"`
}

// NewListSyncsFromWarehouseV1Output instantiates a new ListSyncsFromWarehouseV1Output object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListSyncsFromWarehouseV1Output(
	reports []SyncV1,
	pagination Pagination,
) *ListSyncsFromWarehouseV1Output {
	this := ListSyncsFromWarehouseV1Output{}
	this.Reports = reports
	this.Pagination = pagination
	return &this
}

// NewListSyncsFromWarehouseV1OutputWithDefaults instantiates a new ListSyncsFromWarehouseV1Output object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListSyncsFromWarehouseV1OutputWithDefaults() *ListSyncsFromWarehouseV1Output {
	this := ListSyncsFromWarehouseV1Output{}
	return &this
}

// GetReports returns the Reports field value
func (o *ListSyncsFromWarehouseV1Output) GetReports() []SyncV1 {
	if o == nil {
		var ret []SyncV1
		return ret
	}

	return o.Reports
}

// GetReportsOk returns a tuple with the Reports field value
// and a boolean to check if the value has been set.
func (o *ListSyncsFromWarehouseV1Output) GetReportsOk() ([]SyncV1, bool) {
	if o == nil {
		return nil, false
	}
	return o.Reports, true
}

// SetReports sets field value
func (o *ListSyncsFromWarehouseV1Output) SetReports(v []SyncV1) {
	o.Reports = v
}

// GetPagination returns the Pagination field value
func (o *ListSyncsFromWarehouseV1Output) GetPagination() Pagination {
	if o == nil {
		var ret Pagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *ListSyncsFromWarehouseV1Output) GetPaginationOk() (*Pagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *ListSyncsFromWarehouseV1Output) SetPagination(v Pagination) {
	o.Pagination = v
}

func (o ListSyncsFromWarehouseV1Output) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["reports"] = o.Reports
	}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	return json.Marshal(toSerialize)
}

type NullableListSyncsFromWarehouseV1Output struct {
	value *ListSyncsFromWarehouseV1Output
	isSet bool
}

func (v NullableListSyncsFromWarehouseV1Output) Get() *ListSyncsFromWarehouseV1Output {
	return v.value
}

func (v *NullableListSyncsFromWarehouseV1Output) Set(val *ListSyncsFromWarehouseV1Output) {
	v.value = val
	v.isSet = true
}

func (v NullableListSyncsFromWarehouseV1Output) IsSet() bool {
	return v.isSet
}

func (v *NullableListSyncsFromWarehouseV1Output) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListSyncsFromWarehouseV1Output(
	val *ListSyncsFromWarehouseV1Output,
) *NullableListSyncsFromWarehouseV1Output {
	return &NullableListSyncsFromWarehouseV1Output{value: val, isSet: true}
}

func (v NullableListSyncsFromWarehouseV1Output) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListSyncsFromWarehouseV1Output) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
