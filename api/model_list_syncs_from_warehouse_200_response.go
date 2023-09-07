/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 36.1.1
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// ListSyncsFromWarehouse200Response struct for ListSyncsFromWarehouse200Response
type ListSyncsFromWarehouse200Response struct {
	Data *ListSyncsFromWarehouseV1Output `json:"data,omitempty"`
}

// NewListSyncsFromWarehouse200Response instantiates a new ListSyncsFromWarehouse200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListSyncsFromWarehouse200Response() *ListSyncsFromWarehouse200Response {
	this := ListSyncsFromWarehouse200Response{}
	return &this
}

// NewListSyncsFromWarehouse200ResponseWithDefaults instantiates a new ListSyncsFromWarehouse200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListSyncsFromWarehouse200ResponseWithDefaults() *ListSyncsFromWarehouse200Response {
	this := ListSyncsFromWarehouse200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ListSyncsFromWarehouse200Response) GetData() ListSyncsFromWarehouseV1Output {
	if o == nil || o.Data == nil {
		var ret ListSyncsFromWarehouseV1Output
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSyncsFromWarehouse200Response) GetDataOk() (*ListSyncsFromWarehouseV1Output, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ListSyncsFromWarehouse200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given ListSyncsFromWarehouseV1Output and assigns it to the Data field.
func (o *ListSyncsFromWarehouse200Response) SetData(v ListSyncsFromWarehouseV1Output) {
	o.Data = &v
}

func (o ListSyncsFromWarehouse200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableListSyncsFromWarehouse200Response struct {
	value *ListSyncsFromWarehouse200Response
	isSet bool
}

func (v NullableListSyncsFromWarehouse200Response) Get() *ListSyncsFromWarehouse200Response {
	return v.value
}

func (v *NullableListSyncsFromWarehouse200Response) Set(val *ListSyncsFromWarehouse200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableListSyncsFromWarehouse200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableListSyncsFromWarehouse200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListSyncsFromWarehouse200Response(
	val *ListSyncsFromWarehouse200Response,
) *NullableListSyncsFromWarehouse200Response {
	return &NullableListSyncsFromWarehouse200Response{value: val, isSet: true}
}

func (v NullableListSyncsFromWarehouse200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListSyncsFromWarehouse200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
