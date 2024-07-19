/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 50.4.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the GetSourcesCatalog200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSourcesCatalog200Response{}

// GetSourcesCatalog200Response struct for GetSourcesCatalog200Response
type GetSourcesCatalog200Response struct {
	Data *GetSourcesCatalogV1Output `json:"data,omitempty"`
}

// NewGetSourcesCatalog200Response instantiates a new GetSourcesCatalog200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSourcesCatalog200Response() *GetSourcesCatalog200Response {
	this := GetSourcesCatalog200Response{}
	return &this
}

// NewGetSourcesCatalog200ResponseWithDefaults instantiates a new GetSourcesCatalog200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSourcesCatalog200ResponseWithDefaults() *GetSourcesCatalog200Response {
	this := GetSourcesCatalog200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetSourcesCatalog200Response) GetData() GetSourcesCatalogV1Output {
	if o == nil || IsNil(o.Data) {
		var ret GetSourcesCatalogV1Output
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSourcesCatalog200Response) GetDataOk() (*GetSourcesCatalogV1Output, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetSourcesCatalog200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given GetSourcesCatalogV1Output and assigns it to the Data field.
func (o *GetSourcesCatalog200Response) SetData(v GetSourcesCatalogV1Output) {
	o.Data = &v
}

func (o GetSourcesCatalog200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSourcesCatalog200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGetSourcesCatalog200Response struct {
	value *GetSourcesCatalog200Response
	isSet bool
}

func (v NullableGetSourcesCatalog200Response) Get() *GetSourcesCatalog200Response {
	return v.value
}

func (v *NullableGetSourcesCatalog200Response) Set(val *GetSourcesCatalog200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSourcesCatalog200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSourcesCatalog200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSourcesCatalog200Response(
	val *GetSourcesCatalog200Response,
) *NullableGetSourcesCatalog200Response {
	return &NullableGetSourcesCatalog200Response{value: val, isSet: true}
}

func (v NullableGetSourcesCatalog200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSourcesCatalog200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
