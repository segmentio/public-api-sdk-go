/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 34.4.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// GetSpace200Response struct for GetSpace200Response
type GetSpace200Response struct {
	Data *GetSpaceAlphaOutput `json:"data,omitempty"`
}

// NewGetSpace200Response instantiates a new GetSpace200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSpace200Response() *GetSpace200Response {
	this := GetSpace200Response{}
	return &this
}

// NewGetSpace200ResponseWithDefaults instantiates a new GetSpace200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSpace200ResponseWithDefaults() *GetSpace200Response {
	this := GetSpace200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetSpace200Response) GetData() GetSpaceAlphaOutput {
	if o == nil || o.Data == nil {
		var ret GetSpaceAlphaOutput
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSpace200Response) GetDataOk() (*GetSpaceAlphaOutput, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetSpace200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GetSpaceAlphaOutput and assigns it to the Data field.
func (o *GetSpace200Response) SetData(v GetSpaceAlphaOutput) {
	o.Data = &v
}

func (o GetSpace200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGetSpace200Response struct {
	value *GetSpace200Response
	isSet bool
}

func (v NullableGetSpace200Response) Get() *GetSpace200Response {
	return v.value
}

func (v *NullableGetSpace200Response) Set(val *GetSpace200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSpace200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSpace200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSpace200Response(val *GetSpace200Response) *NullableGetSpace200Response {
	return &NullableGetSpace200Response{value: val, isSet: true}
}

func (v NullableGetSpace200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSpace200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
