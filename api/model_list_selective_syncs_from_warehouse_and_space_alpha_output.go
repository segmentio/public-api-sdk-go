/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 49.2.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the ListSelectiveSyncsFromWarehouseAndSpaceAlphaOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListSelectiveSyncsFromWarehouseAndSpaceAlphaOutput{}

// ListSelectiveSyncsFromWarehouseAndSpaceAlphaOutput Results containing the Selective Sync configuration for a Space Warehouse Connection.
type ListSelectiveSyncsFromWarehouseAndSpaceAlphaOutput struct {
	// Represents a list of collections and properties synced to the Warehouse.
	Items []SpaceWarehouseSelectiveSyncItemAlpha `json:"items"`
	// Optional. Represents the enabled state of all event tables.
	EnableEventTables *bool            `json:"enableEventTables,omitempty"`
	Pagination        PaginationOutput `json:"pagination"`
}

// NewListSelectiveSyncsFromWarehouseAndSpaceAlphaOutput instantiates a new ListSelectiveSyncsFromWarehouseAndSpaceAlphaOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListSelectiveSyncsFromWarehouseAndSpaceAlphaOutput(
	items []SpaceWarehouseSelectiveSyncItemAlpha,
	pagination PaginationOutput,
) *ListSelectiveSyncsFromWarehouseAndSpaceAlphaOutput {
	this := ListSelectiveSyncsFromWarehouseAndSpaceAlphaOutput{}
	this.Items = items
	this.Pagination = pagination
	return &this
}

// NewListSelectiveSyncsFromWarehouseAndSpaceAlphaOutputWithDefaults instantiates a new ListSelectiveSyncsFromWarehouseAndSpaceAlphaOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListSelectiveSyncsFromWarehouseAndSpaceAlphaOutputWithDefaults() *ListSelectiveSyncsFromWarehouseAndSpaceAlphaOutput {
	this := ListSelectiveSyncsFromWarehouseAndSpaceAlphaOutput{}
	return &this
}

// GetItems returns the Items field value
func (o *ListSelectiveSyncsFromWarehouseAndSpaceAlphaOutput) GetItems() []SpaceWarehouseSelectiveSyncItemAlpha {
	if o == nil {
		var ret []SpaceWarehouseSelectiveSyncItemAlpha
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *ListSelectiveSyncsFromWarehouseAndSpaceAlphaOutput) GetItemsOk() ([]SpaceWarehouseSelectiveSyncItemAlpha, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *ListSelectiveSyncsFromWarehouseAndSpaceAlphaOutput) SetItems(
	v []SpaceWarehouseSelectiveSyncItemAlpha,
) {
	o.Items = v
}

// GetEnableEventTables returns the EnableEventTables field value if set, zero value otherwise.
func (o *ListSelectiveSyncsFromWarehouseAndSpaceAlphaOutput) GetEnableEventTables() bool {
	if o == nil || IsNil(o.EnableEventTables) {
		var ret bool
		return ret
	}
	return *o.EnableEventTables
}

// GetEnableEventTablesOk returns a tuple with the EnableEventTables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSelectiveSyncsFromWarehouseAndSpaceAlphaOutput) GetEnableEventTablesOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableEventTables) {
		return nil, false
	}
	return o.EnableEventTables, true
}

// HasEnableEventTables returns a boolean if a field has been set.
func (o *ListSelectiveSyncsFromWarehouseAndSpaceAlphaOutput) HasEnableEventTables() bool {
	if o != nil && !IsNil(o.EnableEventTables) {
		return true
	}

	return false
}

// SetEnableEventTables gets a reference to the given bool and assigns it to the EnableEventTables field.
func (o *ListSelectiveSyncsFromWarehouseAndSpaceAlphaOutput) SetEnableEventTables(v bool) {
	o.EnableEventTables = &v
}

// GetPagination returns the Pagination field value
func (o *ListSelectiveSyncsFromWarehouseAndSpaceAlphaOutput) GetPagination() PaginationOutput {
	if o == nil {
		var ret PaginationOutput
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *ListSelectiveSyncsFromWarehouseAndSpaceAlphaOutput) GetPaginationOk() (*PaginationOutput, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *ListSelectiveSyncsFromWarehouseAndSpaceAlphaOutput) SetPagination(v PaginationOutput) {
	o.Pagination = v
}

func (o ListSelectiveSyncsFromWarehouseAndSpaceAlphaOutput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListSelectiveSyncsFromWarehouseAndSpaceAlphaOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["items"] = o.Items
	if !IsNil(o.EnableEventTables) {
		toSerialize["enableEventTables"] = o.EnableEventTables
	}
	toSerialize["pagination"] = o.Pagination
	return toSerialize, nil
}

type NullableListSelectiveSyncsFromWarehouseAndSpaceAlphaOutput struct {
	value *ListSelectiveSyncsFromWarehouseAndSpaceAlphaOutput
	isSet bool
}

func (v NullableListSelectiveSyncsFromWarehouseAndSpaceAlphaOutput) Get() *ListSelectiveSyncsFromWarehouseAndSpaceAlphaOutput {
	return v.value
}

func (v *NullableListSelectiveSyncsFromWarehouseAndSpaceAlphaOutput) Set(
	val *ListSelectiveSyncsFromWarehouseAndSpaceAlphaOutput,
) {
	v.value = val
	v.isSet = true
}

func (v NullableListSelectiveSyncsFromWarehouseAndSpaceAlphaOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableListSelectiveSyncsFromWarehouseAndSpaceAlphaOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListSelectiveSyncsFromWarehouseAndSpaceAlphaOutput(
	val *ListSelectiveSyncsFromWarehouseAndSpaceAlphaOutput,
) *NullableListSelectiveSyncsFromWarehouseAndSpaceAlphaOutput {
	return &NullableListSelectiveSyncsFromWarehouseAndSpaceAlphaOutput{value: val, isSet: true}
}

func (v NullableListSelectiveSyncsFromWarehouseAndSpaceAlphaOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListSelectiveSyncsFromWarehouseAndSpaceAlphaOutput) UnmarshalJSON(
	src []byte,
) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
