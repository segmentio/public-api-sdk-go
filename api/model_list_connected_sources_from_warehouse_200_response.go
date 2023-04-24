/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 34.3.5
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// ListConnectedSourcesFromWarehouse200Response struct for ListConnectedSourcesFromWarehouse200Response
type ListConnectedSourcesFromWarehouse200Response struct {
	Data *ListConnectedSourcesFromWarehouseV1Output `json:"data,omitempty"`
}

// NewListConnectedSourcesFromWarehouse200Response instantiates a new ListConnectedSourcesFromWarehouse200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListConnectedSourcesFromWarehouse200Response() *ListConnectedSourcesFromWarehouse200Response {
	this := ListConnectedSourcesFromWarehouse200Response{}
	return &this
}

// NewListConnectedSourcesFromWarehouse200ResponseWithDefaults instantiates a new ListConnectedSourcesFromWarehouse200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListConnectedSourcesFromWarehouse200ResponseWithDefaults() *ListConnectedSourcesFromWarehouse200Response {
	this := ListConnectedSourcesFromWarehouse200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ListConnectedSourcesFromWarehouse200Response) GetData() ListConnectedSourcesFromWarehouseV1Output {
	if o == nil || o.Data == nil {
		var ret ListConnectedSourcesFromWarehouseV1Output
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListConnectedSourcesFromWarehouse200Response) GetDataOk() (*ListConnectedSourcesFromWarehouseV1Output, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ListConnectedSourcesFromWarehouse200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given ListConnectedSourcesFromWarehouseV1Output and assigns it to the Data field.
func (o *ListConnectedSourcesFromWarehouse200Response) SetData(
	v ListConnectedSourcesFromWarehouseV1Output,
) {
	o.Data = &v
}

func (o ListConnectedSourcesFromWarehouse200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableListConnectedSourcesFromWarehouse200Response struct {
	value *ListConnectedSourcesFromWarehouse200Response
	isSet bool
}

func (v NullableListConnectedSourcesFromWarehouse200Response) Get() *ListConnectedSourcesFromWarehouse200Response {
	return v.value
}

func (v *NullableListConnectedSourcesFromWarehouse200Response) Set(
	val *ListConnectedSourcesFromWarehouse200Response,
) {
	v.value = val
	v.isSet = true
}

func (v NullableListConnectedSourcesFromWarehouse200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableListConnectedSourcesFromWarehouse200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListConnectedSourcesFromWarehouse200Response(
	val *ListConnectedSourcesFromWarehouse200Response,
) *NullableListConnectedSourcesFromWarehouse200Response {
	return &NullableListConnectedSourcesFromWarehouse200Response{value: val, isSet: true}
}

func (v NullableListConnectedSourcesFromWarehouse200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListConnectedSourcesFromWarehouse200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
