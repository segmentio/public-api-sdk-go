/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 58.13.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the GetActivationFromAudience200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetActivationFromAudience200Response{}

// GetActivationFromAudience200Response struct for GetActivationFromAudience200Response
type GetActivationFromAudience200Response struct {
	Data *GetActivationFromAudienceOutput `json:"data,omitempty"`
}

// NewGetActivationFromAudience200Response instantiates a new GetActivationFromAudience200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetActivationFromAudience200Response() *GetActivationFromAudience200Response {
	this := GetActivationFromAudience200Response{}
	return &this
}

// NewGetActivationFromAudience200ResponseWithDefaults instantiates a new GetActivationFromAudience200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetActivationFromAudience200ResponseWithDefaults() *GetActivationFromAudience200Response {
	this := GetActivationFromAudience200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetActivationFromAudience200Response) GetData() GetActivationFromAudienceOutput {
	if o == nil || IsNil(o.Data) {
		var ret GetActivationFromAudienceOutput
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetActivationFromAudience200Response) GetDataOk() (*GetActivationFromAudienceOutput, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetActivationFromAudience200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given GetActivationFromAudienceOutput and assigns it to the Data field.
func (o *GetActivationFromAudience200Response) SetData(v GetActivationFromAudienceOutput) {
	o.Data = &v
}

func (o GetActivationFromAudience200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetActivationFromAudience200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGetActivationFromAudience200Response struct {
	value *GetActivationFromAudience200Response
	isSet bool
}

func (v NullableGetActivationFromAudience200Response) Get() *GetActivationFromAudience200Response {
	return v.value
}

func (v *NullableGetActivationFromAudience200Response) Set(
	val *GetActivationFromAudience200Response,
) {
	v.value = val
	v.isSet = true
}

func (v NullableGetActivationFromAudience200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetActivationFromAudience200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetActivationFromAudience200Response(
	val *GetActivationFromAudience200Response,
) *NullableGetActivationFromAudience200Response {
	return &NullableGetActivationFromAudience200Response{value: val, isSet: true}
}

func (v NullableGetActivationFromAudience200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetActivationFromAudience200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
