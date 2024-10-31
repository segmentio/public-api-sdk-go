/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 55.2.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the GetSourcesCatalogV1Output type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSourcesCatalogV1Output{}

// GetSourcesCatalogV1Output Returns a list of all Source catalog items contained within a given page.
type GetSourcesCatalogV1Output struct {
	// All Source catalog items contained within the requested page.
	SourcesCatalog []SourceMetadataV1 `json:"sourcesCatalog"`
	Pagination     PaginationOutput   `json:"pagination"`
}

// NewGetSourcesCatalogV1Output instantiates a new GetSourcesCatalogV1Output object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSourcesCatalogV1Output(
	sourcesCatalog []SourceMetadataV1,
	pagination PaginationOutput,
) *GetSourcesCatalogV1Output {
	this := GetSourcesCatalogV1Output{}
	this.SourcesCatalog = sourcesCatalog
	this.Pagination = pagination
	return &this
}

// NewGetSourcesCatalogV1OutputWithDefaults instantiates a new GetSourcesCatalogV1Output object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSourcesCatalogV1OutputWithDefaults() *GetSourcesCatalogV1Output {
	this := GetSourcesCatalogV1Output{}
	return &this
}

// GetSourcesCatalog returns the SourcesCatalog field value
func (o *GetSourcesCatalogV1Output) GetSourcesCatalog() []SourceMetadataV1 {
	if o == nil {
		var ret []SourceMetadataV1
		return ret
	}

	return o.SourcesCatalog
}

// GetSourcesCatalogOk returns a tuple with the SourcesCatalog field value
// and a boolean to check if the value has been set.
func (o *GetSourcesCatalogV1Output) GetSourcesCatalogOk() ([]SourceMetadataV1, bool) {
	if o == nil {
		return nil, false
	}
	return o.SourcesCatalog, true
}

// SetSourcesCatalog sets field value
func (o *GetSourcesCatalogV1Output) SetSourcesCatalog(v []SourceMetadataV1) {
	o.SourcesCatalog = v
}

// GetPagination returns the Pagination field value
func (o *GetSourcesCatalogV1Output) GetPagination() PaginationOutput {
	if o == nil {
		var ret PaginationOutput
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *GetSourcesCatalogV1Output) GetPaginationOk() (*PaginationOutput, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *GetSourcesCatalogV1Output) SetPagination(v PaginationOutput) {
	o.Pagination = v
}

func (o GetSourcesCatalogV1Output) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSourcesCatalogV1Output) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sourcesCatalog"] = o.SourcesCatalog
	toSerialize["pagination"] = o.Pagination
	return toSerialize, nil
}

type NullableGetSourcesCatalogV1Output struct {
	value *GetSourcesCatalogV1Output
	isSet bool
}

func (v NullableGetSourcesCatalogV1Output) Get() *GetSourcesCatalogV1Output {
	return v.value
}

func (v *NullableGetSourcesCatalogV1Output) Set(val *GetSourcesCatalogV1Output) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSourcesCatalogV1Output) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSourcesCatalogV1Output) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSourcesCatalogV1Output(
	val *GetSourcesCatalogV1Output,
) *NullableGetSourcesCatalogV1Output {
	return &NullableGetSourcesCatalogV1Output{value: val, isSet: true}
}

func (v NullableGetSourcesCatalogV1Output) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSourcesCatalogV1Output) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
