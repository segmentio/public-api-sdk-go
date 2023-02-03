/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 34.0.1
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// GetLatestFromEdgeFunctions200Response struct for GetLatestFromEdgeFunctions200Response
type GetLatestFromEdgeFunctions200Response struct {
	Data *GetLatestFromEdgeFunctionsAlphaOutput `json:"data,omitempty"`
}

// NewGetLatestFromEdgeFunctions200Response instantiates a new GetLatestFromEdgeFunctions200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetLatestFromEdgeFunctions200Response() *GetLatestFromEdgeFunctions200Response {
	this := GetLatestFromEdgeFunctions200Response{}
	return &this
}

// NewGetLatestFromEdgeFunctions200ResponseWithDefaults instantiates a new GetLatestFromEdgeFunctions200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetLatestFromEdgeFunctions200ResponseWithDefaults() *GetLatestFromEdgeFunctions200Response {
	this := GetLatestFromEdgeFunctions200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetLatestFromEdgeFunctions200Response) GetData() GetLatestFromEdgeFunctionsAlphaOutput {
	if o == nil || o.Data == nil {
		var ret GetLatestFromEdgeFunctionsAlphaOutput
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLatestFromEdgeFunctions200Response) GetDataOk() (*GetLatestFromEdgeFunctionsAlphaOutput, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetLatestFromEdgeFunctions200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GetLatestFromEdgeFunctionsAlphaOutput and assigns it to the Data field.
func (o *GetLatestFromEdgeFunctions200Response) SetData(v GetLatestFromEdgeFunctionsAlphaOutput) {
	o.Data = &v
}

func (o GetLatestFromEdgeFunctions200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGetLatestFromEdgeFunctions200Response struct {
	value *GetLatestFromEdgeFunctions200Response
	isSet bool
}

func (v NullableGetLatestFromEdgeFunctions200Response) Get() *GetLatestFromEdgeFunctions200Response {
	return v.value
}

func (v *NullableGetLatestFromEdgeFunctions200Response) Set(
	val *GetLatestFromEdgeFunctions200Response,
) {
	v.value = val
	v.isSet = true
}

func (v NullableGetLatestFromEdgeFunctions200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetLatestFromEdgeFunctions200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetLatestFromEdgeFunctions200Response(
	val *GetLatestFromEdgeFunctions200Response,
) *NullableGetLatestFromEdgeFunctions200Response {
	return &NullableGetLatestFromEdgeFunctions200Response{value: val, isSet: true}
}

func (v NullableGetLatestFromEdgeFunctions200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetLatestFromEdgeFunctions200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
