/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 57.0.1
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the GetFilterById200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetFilterById200Response{}

// GetFilterById200Response struct for GetFilterById200Response
type GetFilterById200Response struct {
	Data *GetFilterByIdOutput `json:"data,omitempty"`
}

// NewGetFilterById200Response instantiates a new GetFilterById200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFilterById200Response() *GetFilterById200Response {
	this := GetFilterById200Response{}
	return &this
}

// NewGetFilterById200ResponseWithDefaults instantiates a new GetFilterById200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFilterById200ResponseWithDefaults() *GetFilterById200Response {
	this := GetFilterById200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetFilterById200Response) GetData() GetFilterByIdOutput {
	if o == nil || IsNil(o.Data) {
		var ret GetFilterByIdOutput
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFilterById200Response) GetDataOk() (*GetFilterByIdOutput, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetFilterById200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given GetFilterByIdOutput and assigns it to the Data field.
func (o *GetFilterById200Response) SetData(v GetFilterByIdOutput) {
	o.Data = &v
}

func (o GetFilterById200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetFilterById200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGetFilterById200Response struct {
	value *GetFilterById200Response
	isSet bool
}

func (v NullableGetFilterById200Response) Get() *GetFilterById200Response {
	return v.value
}

func (v *NullableGetFilterById200Response) Set(val *GetFilterById200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFilterById200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFilterById200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFilterById200Response(
	val *GetFilterById200Response,
) *NullableGetFilterById200Response {
	return &NullableGetFilterById200Response{value: val, isSet: true}
}

func (v NullableGetFilterById200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFilterById200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
