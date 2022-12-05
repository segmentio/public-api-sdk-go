/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 33.0.3
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// ListSyncsFromWarehouseAndSource200Response struct for ListSyncsFromWarehouseAndSource200Response
type ListSyncsFromWarehouseAndSource200Response struct {
	Data *ListSyncsFromWarehouseAndSourceV1Output `json:"data,omitempty"`
}

// NewListSyncsFromWarehouseAndSource200Response instantiates a new ListSyncsFromWarehouseAndSource200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListSyncsFromWarehouseAndSource200Response() *ListSyncsFromWarehouseAndSource200Response {
	this := ListSyncsFromWarehouseAndSource200Response{}
	return &this
}

// NewListSyncsFromWarehouseAndSource200ResponseWithDefaults instantiates a new ListSyncsFromWarehouseAndSource200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListSyncsFromWarehouseAndSource200ResponseWithDefaults() *ListSyncsFromWarehouseAndSource200Response {
	this := ListSyncsFromWarehouseAndSource200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ListSyncsFromWarehouseAndSource200Response) GetData() ListSyncsFromWarehouseAndSourceV1Output {
	if o == nil || o.Data == nil {
		var ret ListSyncsFromWarehouseAndSourceV1Output
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSyncsFromWarehouseAndSource200Response) GetDataOk() (*ListSyncsFromWarehouseAndSourceV1Output, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ListSyncsFromWarehouseAndSource200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given ListSyncsFromWarehouseAndSourceV1Output and assigns it to the Data field.
func (o *ListSyncsFromWarehouseAndSource200Response) SetData(
	v ListSyncsFromWarehouseAndSourceV1Output,
) {
	o.Data = &v
}

func (o ListSyncsFromWarehouseAndSource200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableListSyncsFromWarehouseAndSource200Response struct {
	value *ListSyncsFromWarehouseAndSource200Response
	isSet bool
}

func (v NullableListSyncsFromWarehouseAndSource200Response) Get() *ListSyncsFromWarehouseAndSource200Response {
	return v.value
}

func (v *NullableListSyncsFromWarehouseAndSource200Response) Set(
	val *ListSyncsFromWarehouseAndSource200Response,
) {
	v.value = val
	v.isSet = true
}

func (v NullableListSyncsFromWarehouseAndSource200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableListSyncsFromWarehouseAndSource200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListSyncsFromWarehouseAndSource200Response(
	val *ListSyncsFromWarehouseAndSource200Response,
) *NullableListSyncsFromWarehouseAndSource200Response {
	return &NullableListSyncsFromWarehouseAndSource200Response{value: val, isSet: true}
}

func (v NullableListSyncsFromWarehouseAndSource200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListSyncsFromWarehouseAndSource200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
