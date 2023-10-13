/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 36.2.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// GetAudience200Response struct for GetAudience200Response
type GetAudience200Response struct {
	Data *GetAudienceAlphaOutput `json:"data,omitempty"`
}

// NewGetAudience200Response instantiates a new GetAudience200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAudience200Response() *GetAudience200Response {
	this := GetAudience200Response{}
	return &this
}

// NewGetAudience200ResponseWithDefaults instantiates a new GetAudience200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAudience200ResponseWithDefaults() *GetAudience200Response {
	this := GetAudience200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetAudience200Response) GetData() GetAudienceAlphaOutput {
	if o == nil || o.Data == nil {
		var ret GetAudienceAlphaOutput
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAudience200Response) GetDataOk() (*GetAudienceAlphaOutput, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetAudience200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GetAudienceAlphaOutput and assigns it to the Data field.
func (o *GetAudience200Response) SetData(v GetAudienceAlphaOutput) {
	o.Data = &v
}

func (o GetAudience200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGetAudience200Response struct {
	value *GetAudience200Response
	isSet bool
}

func (v NullableGetAudience200Response) Get() *GetAudience200Response {
	return v.value
}

func (v *NullableGetAudience200Response) Set(val *GetAudience200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAudience200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAudience200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAudience200Response(
	val *GetAudience200Response,
) *NullableGetAudience200Response {
	return &NullableGetAudience200Response{value: val, isSet: true}
}

func (v NullableGetAudience200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAudience200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}