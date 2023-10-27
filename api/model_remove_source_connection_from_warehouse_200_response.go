/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 36.3.1
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// RemoveSourceConnectionFromWarehouse200Response struct for RemoveSourceConnectionFromWarehouse200Response
type RemoveSourceConnectionFromWarehouse200Response struct {
	Data *RemoveSourceConnectionFromWarehouseV1Output `json:"data,omitempty"`
}

// NewRemoveSourceConnectionFromWarehouse200Response instantiates a new RemoveSourceConnectionFromWarehouse200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemoveSourceConnectionFromWarehouse200Response() *RemoveSourceConnectionFromWarehouse200Response {
	this := RemoveSourceConnectionFromWarehouse200Response{}
	return &this
}

// NewRemoveSourceConnectionFromWarehouse200ResponseWithDefaults instantiates a new RemoveSourceConnectionFromWarehouse200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemoveSourceConnectionFromWarehouse200ResponseWithDefaults() *RemoveSourceConnectionFromWarehouse200Response {
	this := RemoveSourceConnectionFromWarehouse200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *RemoveSourceConnectionFromWarehouse200Response) GetData() RemoveSourceConnectionFromWarehouseV1Output {
	if o == nil || o.Data == nil {
		var ret RemoveSourceConnectionFromWarehouseV1Output
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoveSourceConnectionFromWarehouse200Response) GetDataOk() (*RemoveSourceConnectionFromWarehouseV1Output, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *RemoveSourceConnectionFromWarehouse200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given RemoveSourceConnectionFromWarehouseV1Output and assigns it to the Data field.
func (o *RemoveSourceConnectionFromWarehouse200Response) SetData(
	v RemoveSourceConnectionFromWarehouseV1Output,
) {
	o.Data = &v
}

func (o RemoveSourceConnectionFromWarehouse200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableRemoveSourceConnectionFromWarehouse200Response struct {
	value *RemoveSourceConnectionFromWarehouse200Response
	isSet bool
}

func (v NullableRemoveSourceConnectionFromWarehouse200Response) Get() *RemoveSourceConnectionFromWarehouse200Response {
	return v.value
}

func (v *NullableRemoveSourceConnectionFromWarehouse200Response) Set(
	val *RemoveSourceConnectionFromWarehouse200Response,
) {
	v.value = val
	v.isSet = true
}

func (v NullableRemoveSourceConnectionFromWarehouse200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableRemoveSourceConnectionFromWarehouse200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemoveSourceConnectionFromWarehouse200Response(
	val *RemoveSourceConnectionFromWarehouse200Response,
) *NullableRemoveSourceConnectionFromWarehouse200Response {
	return &NullableRemoveSourceConnectionFromWarehouse200Response{value: val, isSet: true}
}

func (v NullableRemoveSourceConnectionFromWarehouse200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemoveSourceConnectionFromWarehouse200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
