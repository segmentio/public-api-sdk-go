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

// checks if the ListSelectiveSyncsFromWarehouseAndSourceV1Output type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListSelectiveSyncsFromWarehouseAndSourceV1Output{}

// ListSelectiveSyncsFromWarehouseAndSourceV1Output Results containing the Selective Sync configuration for a Warehouse.
type ListSelectiveSyncsFromWarehouseAndSourceV1Output struct {
	// Represents a list of Source, collection, and properties synced to the Warehouse.
	Items      []WarehouseSelectiveSyncItemV1 `json:"items"`
	Pagination PaginationOutput               `json:"pagination"`
}

// NewListSelectiveSyncsFromWarehouseAndSourceV1Output instantiates a new ListSelectiveSyncsFromWarehouseAndSourceV1Output object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListSelectiveSyncsFromWarehouseAndSourceV1Output(
	items []WarehouseSelectiveSyncItemV1,
	pagination PaginationOutput,
) *ListSelectiveSyncsFromWarehouseAndSourceV1Output {
	this := ListSelectiveSyncsFromWarehouseAndSourceV1Output{}
	this.Items = items
	this.Pagination = pagination
	return &this
}

// NewListSelectiveSyncsFromWarehouseAndSourceV1OutputWithDefaults instantiates a new ListSelectiveSyncsFromWarehouseAndSourceV1Output object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListSelectiveSyncsFromWarehouseAndSourceV1OutputWithDefaults() *ListSelectiveSyncsFromWarehouseAndSourceV1Output {
	this := ListSelectiveSyncsFromWarehouseAndSourceV1Output{}
	return &this
}

// GetItems returns the Items field value
func (o *ListSelectiveSyncsFromWarehouseAndSourceV1Output) GetItems() []WarehouseSelectiveSyncItemV1 {
	if o == nil {
		var ret []WarehouseSelectiveSyncItemV1
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *ListSelectiveSyncsFromWarehouseAndSourceV1Output) GetItemsOk() ([]WarehouseSelectiveSyncItemV1, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *ListSelectiveSyncsFromWarehouseAndSourceV1Output) SetItems(
	v []WarehouseSelectiveSyncItemV1,
) {
	o.Items = v
}

// GetPagination returns the Pagination field value
func (o *ListSelectiveSyncsFromWarehouseAndSourceV1Output) GetPagination() PaginationOutput {
	if o == nil {
		var ret PaginationOutput
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *ListSelectiveSyncsFromWarehouseAndSourceV1Output) GetPaginationOk() (*PaginationOutput, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *ListSelectiveSyncsFromWarehouseAndSourceV1Output) SetPagination(v PaginationOutput) {
	o.Pagination = v
}

func (o ListSelectiveSyncsFromWarehouseAndSourceV1Output) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListSelectiveSyncsFromWarehouseAndSourceV1Output) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["items"] = o.Items
	toSerialize["pagination"] = o.Pagination
	return toSerialize, nil
}

type NullableListSelectiveSyncsFromWarehouseAndSourceV1Output struct {
	value *ListSelectiveSyncsFromWarehouseAndSourceV1Output
	isSet bool
}

func (v NullableListSelectiveSyncsFromWarehouseAndSourceV1Output) Get() *ListSelectiveSyncsFromWarehouseAndSourceV1Output {
	return v.value
}

func (v *NullableListSelectiveSyncsFromWarehouseAndSourceV1Output) Set(
	val *ListSelectiveSyncsFromWarehouseAndSourceV1Output,
) {
	v.value = val
	v.isSet = true
}

func (v NullableListSelectiveSyncsFromWarehouseAndSourceV1Output) IsSet() bool {
	return v.isSet
}

func (v *NullableListSelectiveSyncsFromWarehouseAndSourceV1Output) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListSelectiveSyncsFromWarehouseAndSourceV1Output(
	val *ListSelectiveSyncsFromWarehouseAndSourceV1Output,
) *NullableListSelectiveSyncsFromWarehouseAndSourceV1Output {
	return &NullableListSelectiveSyncsFromWarehouseAndSourceV1Output{value: val, isSet: true}
}

func (v NullableListSelectiveSyncsFromWarehouseAndSourceV1Output) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListSelectiveSyncsFromWarehouseAndSourceV1Output) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
