/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 36.2.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// GetDestinationsCatalogV1Output Returns a list of all Destination catalog items contained within a given page.
type GetDestinationsCatalogV1Output struct {
	// All Destination catalog items contained within the requested page.
	DestinationsCatalog []DestinationMetadataV1 `json:"destinationsCatalog"`
	Pagination          Pagination              `json:"pagination"`
}

// NewGetDestinationsCatalogV1Output instantiates a new GetDestinationsCatalogV1Output object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDestinationsCatalogV1Output(
	destinationsCatalog []DestinationMetadataV1,
	pagination Pagination,
) *GetDestinationsCatalogV1Output {
	this := GetDestinationsCatalogV1Output{}
	this.DestinationsCatalog = destinationsCatalog
	this.Pagination = pagination
	return &this
}

// NewGetDestinationsCatalogV1OutputWithDefaults instantiates a new GetDestinationsCatalogV1Output object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDestinationsCatalogV1OutputWithDefaults() *GetDestinationsCatalogV1Output {
	this := GetDestinationsCatalogV1Output{}
	return &this
}

// GetDestinationsCatalog returns the DestinationsCatalog field value
func (o *GetDestinationsCatalogV1Output) GetDestinationsCatalog() []DestinationMetadataV1 {
	if o == nil {
		var ret []DestinationMetadataV1
		return ret
	}

	return o.DestinationsCatalog
}

// GetDestinationsCatalogOk returns a tuple with the DestinationsCatalog field value
// and a boolean to check if the value has been set.
func (o *GetDestinationsCatalogV1Output) GetDestinationsCatalogOk() ([]DestinationMetadataV1, bool) {
	if o == nil {
		return nil, false
	}
	return o.DestinationsCatalog, true
}

// SetDestinationsCatalog sets field value
func (o *GetDestinationsCatalogV1Output) SetDestinationsCatalog(v []DestinationMetadataV1) {
	o.DestinationsCatalog = v
}

// GetPagination returns the Pagination field value
func (o *GetDestinationsCatalogV1Output) GetPagination() Pagination {
	if o == nil {
		var ret Pagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *GetDestinationsCatalogV1Output) GetPaginationOk() (*Pagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *GetDestinationsCatalogV1Output) SetPagination(v Pagination) {
	o.Pagination = v
}

func (o GetDestinationsCatalogV1Output) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["destinationsCatalog"] = o.DestinationsCatalog
	}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	return json.Marshal(toSerialize)
}

type NullableGetDestinationsCatalogV1Output struct {
	value *GetDestinationsCatalogV1Output
	isSet bool
}

func (v NullableGetDestinationsCatalogV1Output) Get() *GetDestinationsCatalogV1Output {
	return v.value
}

func (v *NullableGetDestinationsCatalogV1Output) Set(val *GetDestinationsCatalogV1Output) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDestinationsCatalogV1Output) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDestinationsCatalogV1Output) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDestinationsCatalogV1Output(
	val *GetDestinationsCatalogV1Output,
) *NullableGetDestinationsCatalogV1Output {
	return &NullableGetDestinationsCatalogV1Output{value: val, isSet: true}
}

func (v NullableGetDestinationsCatalogV1Output) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDestinationsCatalogV1Output) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
