/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 53.3.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the GetWarehousesCatalogV1Output type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetWarehousesCatalogV1Output{}

// GetWarehousesCatalogV1Output Returns a list of all Warehouse catalog items contained within a given page.
type GetWarehousesCatalogV1Output struct {
	// All Warehouse catalog items contained within the requested page.
	WarehousesCatalog []WarehouseMetadataV1 `json:"warehousesCatalog"`
	Pagination        PaginationOutput      `json:"pagination"`
}

// NewGetWarehousesCatalogV1Output instantiates a new GetWarehousesCatalogV1Output object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetWarehousesCatalogV1Output(
	warehousesCatalog []WarehouseMetadataV1,
	pagination PaginationOutput,
) *GetWarehousesCatalogV1Output {
	this := GetWarehousesCatalogV1Output{}
	this.WarehousesCatalog = warehousesCatalog
	this.Pagination = pagination
	return &this
}

// NewGetWarehousesCatalogV1OutputWithDefaults instantiates a new GetWarehousesCatalogV1Output object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetWarehousesCatalogV1OutputWithDefaults() *GetWarehousesCatalogV1Output {
	this := GetWarehousesCatalogV1Output{}
	return &this
}

// GetWarehousesCatalog returns the WarehousesCatalog field value
func (o *GetWarehousesCatalogV1Output) GetWarehousesCatalog() []WarehouseMetadataV1 {
	if o == nil {
		var ret []WarehouseMetadataV1
		return ret
	}

	return o.WarehousesCatalog
}

// GetWarehousesCatalogOk returns a tuple with the WarehousesCatalog field value
// and a boolean to check if the value has been set.
func (o *GetWarehousesCatalogV1Output) GetWarehousesCatalogOk() ([]WarehouseMetadataV1, bool) {
	if o == nil {
		return nil, false
	}
	return o.WarehousesCatalog, true
}

// SetWarehousesCatalog sets field value
func (o *GetWarehousesCatalogV1Output) SetWarehousesCatalog(v []WarehouseMetadataV1) {
	o.WarehousesCatalog = v
}

// GetPagination returns the Pagination field value
func (o *GetWarehousesCatalogV1Output) GetPagination() PaginationOutput {
	if o == nil {
		var ret PaginationOutput
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *GetWarehousesCatalogV1Output) GetPaginationOk() (*PaginationOutput, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *GetWarehousesCatalogV1Output) SetPagination(v PaginationOutput) {
	o.Pagination = v
}

func (o GetWarehousesCatalogV1Output) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetWarehousesCatalogV1Output) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["warehousesCatalog"] = o.WarehousesCatalog
	toSerialize["pagination"] = o.Pagination
	return toSerialize, nil
}

type NullableGetWarehousesCatalogV1Output struct {
	value *GetWarehousesCatalogV1Output
	isSet bool
}

func (v NullableGetWarehousesCatalogV1Output) Get() *GetWarehousesCatalogV1Output {
	return v.value
}

func (v *NullableGetWarehousesCatalogV1Output) Set(val *GetWarehousesCatalogV1Output) {
	v.value = val
	v.isSet = true
}

func (v NullableGetWarehousesCatalogV1Output) IsSet() bool {
	return v.isSet
}

func (v *NullableGetWarehousesCatalogV1Output) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetWarehousesCatalogV1Output(
	val *GetWarehousesCatalogV1Output,
) *NullableGetWarehousesCatalogV1Output {
	return &NullableGetWarehousesCatalogV1Output{value: val, isSet: true}
}

func (v NullableGetWarehousesCatalogV1Output) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetWarehousesCatalogV1Output) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
