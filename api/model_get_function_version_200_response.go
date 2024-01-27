/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 39.0.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the GetFunctionVersion200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetFunctionVersion200Response{}

// GetFunctionVersion200Response struct for GetFunctionVersion200Response
type GetFunctionVersion200Response struct {
	Data *GetFunctionVersionAlphaOutput `json:"data,omitempty"`
}

// NewGetFunctionVersion200Response instantiates a new GetFunctionVersion200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFunctionVersion200Response() *GetFunctionVersion200Response {
	this := GetFunctionVersion200Response{}
	return &this
}

// NewGetFunctionVersion200ResponseWithDefaults instantiates a new GetFunctionVersion200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFunctionVersion200ResponseWithDefaults() *GetFunctionVersion200Response {
	this := GetFunctionVersion200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetFunctionVersion200Response) GetData() GetFunctionVersionAlphaOutput {
	if o == nil || IsNil(o.Data) {
		var ret GetFunctionVersionAlphaOutput
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFunctionVersion200Response) GetDataOk() (*GetFunctionVersionAlphaOutput, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetFunctionVersion200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given GetFunctionVersionAlphaOutput and assigns it to the Data field.
func (o *GetFunctionVersion200Response) SetData(v GetFunctionVersionAlphaOutput) {
	o.Data = &v
}

func (o GetFunctionVersion200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetFunctionVersion200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGetFunctionVersion200Response struct {
	value *GetFunctionVersion200Response
	isSet bool
}

func (v NullableGetFunctionVersion200Response) Get() *GetFunctionVersion200Response {
	return v.value
}

func (v *NullableGetFunctionVersion200Response) Set(val *GetFunctionVersion200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFunctionVersion200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFunctionVersion200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFunctionVersion200Response(
	val *GetFunctionVersion200Response,
) *NullableGetFunctionVersion200Response {
	return &NullableGetFunctionVersion200Response{value: val, isSet: true}
}

func (v NullableGetFunctionVersion200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFunctionVersion200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
