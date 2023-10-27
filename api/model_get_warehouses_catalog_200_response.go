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

// GetWarehousesCatalog200Response struct for GetWarehousesCatalog200Response
type GetWarehousesCatalog200Response struct {
	Data *GetWarehousesCatalogV1Output `json:"data,omitempty"`
}

// NewGetWarehousesCatalog200Response instantiates a new GetWarehousesCatalog200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetWarehousesCatalog200Response() *GetWarehousesCatalog200Response {
	this := GetWarehousesCatalog200Response{}
	return &this
}

// NewGetWarehousesCatalog200ResponseWithDefaults instantiates a new GetWarehousesCatalog200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetWarehousesCatalog200ResponseWithDefaults() *GetWarehousesCatalog200Response {
	this := GetWarehousesCatalog200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetWarehousesCatalog200Response) GetData() GetWarehousesCatalogV1Output {
	if o == nil || o.Data == nil {
		var ret GetWarehousesCatalogV1Output
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetWarehousesCatalog200Response) GetDataOk() (*GetWarehousesCatalogV1Output, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetWarehousesCatalog200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GetWarehousesCatalogV1Output and assigns it to the Data field.
func (o *GetWarehousesCatalog200Response) SetData(v GetWarehousesCatalogV1Output) {
	o.Data = &v
}

func (o GetWarehousesCatalog200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGetWarehousesCatalog200Response struct {
	value *GetWarehousesCatalog200Response
	isSet bool
}

func (v NullableGetWarehousesCatalog200Response) Get() *GetWarehousesCatalog200Response {
	return v.value
}

func (v *NullableGetWarehousesCatalog200Response) Set(val *GetWarehousesCatalog200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetWarehousesCatalog200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetWarehousesCatalog200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetWarehousesCatalog200Response(
	val *GetWarehousesCatalog200Response,
) *NullableGetWarehousesCatalog200Response {
	return &NullableGetWarehousesCatalog200Response{value: val, isSet: true}
}

func (v NullableGetWarehousesCatalog200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetWarehousesCatalog200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
