/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 38.4.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the GetWarehouseMetadata200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetWarehouseMetadata200Response{}

// GetWarehouseMetadata200Response struct for GetWarehouseMetadata200Response
type GetWarehouseMetadata200Response struct {
	Data *GetWarehouseMetadataV1Output `json:"data,omitempty"`
}

// NewGetWarehouseMetadata200Response instantiates a new GetWarehouseMetadata200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetWarehouseMetadata200Response() *GetWarehouseMetadata200Response {
	this := GetWarehouseMetadata200Response{}
	return &this
}

// NewGetWarehouseMetadata200ResponseWithDefaults instantiates a new GetWarehouseMetadata200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetWarehouseMetadata200ResponseWithDefaults() *GetWarehouseMetadata200Response {
	this := GetWarehouseMetadata200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetWarehouseMetadata200Response) GetData() GetWarehouseMetadataV1Output {
	if o == nil || IsNil(o.Data) {
		var ret GetWarehouseMetadataV1Output
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetWarehouseMetadata200Response) GetDataOk() (*GetWarehouseMetadataV1Output, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetWarehouseMetadata200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given GetWarehouseMetadataV1Output and assigns it to the Data field.
func (o *GetWarehouseMetadata200Response) SetData(v GetWarehouseMetadataV1Output) {
	o.Data = &v
}

func (o GetWarehouseMetadata200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetWarehouseMetadata200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGetWarehouseMetadata200Response struct {
	value *GetWarehouseMetadata200Response
	isSet bool
}

func (v NullableGetWarehouseMetadata200Response) Get() *GetWarehouseMetadata200Response {
	return v.value
}

func (v *NullableGetWarehouseMetadata200Response) Set(val *GetWarehouseMetadata200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetWarehouseMetadata200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetWarehouseMetadata200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetWarehouseMetadata200Response(
	val *GetWarehouseMetadata200Response,
) *NullableGetWarehouseMetadata200Response {
	return &NullableGetWarehouseMetadata200Response{value: val, isSet: true}
}

func (v NullableGetWarehouseMetadata200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetWarehouseMetadata200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
