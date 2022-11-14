/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API. 

API version: 32.0.4
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// GetTransformation200Response struct for GetTransformation200Response
type GetTransformation200Response struct {
	Data *GetTransformationBetaOutput `json:"data,omitempty"`
}

// NewGetTransformation200Response instantiates a new GetTransformation200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTransformation200Response() *GetTransformation200Response {
	this := GetTransformation200Response{}
	return &this
}

// NewGetTransformation200ResponseWithDefaults instantiates a new GetTransformation200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTransformation200ResponseWithDefaults() *GetTransformation200Response {
	this := GetTransformation200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetTransformation200Response) GetData() GetTransformationBetaOutput {
	if o == nil || o.Data == nil {
		var ret GetTransformationBetaOutput
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTransformation200Response) GetDataOk() (*GetTransformationBetaOutput, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetTransformation200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GetTransformationBetaOutput and assigns it to the Data field.
func (o *GetTransformation200Response) SetData(v GetTransformationBetaOutput) {
	o.Data = &v
}

func (o GetTransformation200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGetTransformation200Response struct {
	value *GetTransformation200Response
	isSet bool
}

func (v NullableGetTransformation200Response) Get() *GetTransformation200Response {
	return v.value
}

func (v *NullableGetTransformation200Response) Set(val *GetTransformation200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTransformation200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTransformation200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTransformation200Response(val *GetTransformation200Response) *NullableGetTransformation200Response {
	return &NullableGetTransformation200Response{value: val, isSet: true}
}

func (v NullableGetTransformation200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTransformation200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


