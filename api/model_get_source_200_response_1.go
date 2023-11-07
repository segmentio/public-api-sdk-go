/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 37.2.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// GetSource200Response1 struct for GetSource200Response1
type GetSource200Response1 struct {
	Data *GetSourceAlphaOutput `json:"data,omitempty"`
}

// NewGetSource200Response1 instantiates a new GetSource200Response1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSource200Response1() *GetSource200Response1 {
	this := GetSource200Response1{}
	return &this
}

// NewGetSource200Response1WithDefaults instantiates a new GetSource200Response1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSource200Response1WithDefaults() *GetSource200Response1 {
	this := GetSource200Response1{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetSource200Response1) GetData() GetSourceAlphaOutput {
	if o == nil || o.Data == nil {
		var ret GetSourceAlphaOutput
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSource200Response1) GetDataOk() (*GetSourceAlphaOutput, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetSource200Response1) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GetSourceAlphaOutput and assigns it to the Data field.
func (o *GetSource200Response1) SetData(v GetSourceAlphaOutput) {
	o.Data = &v
}

func (o GetSource200Response1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGetSource200Response1 struct {
	value *GetSource200Response1
	isSet bool
}

func (v NullableGetSource200Response1) Get() *GetSource200Response1 {
	return v.value
}

func (v *NullableGetSource200Response1) Set(val *GetSource200Response1) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSource200Response1) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSource200Response1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSource200Response1(val *GetSource200Response1) *NullableGetSource200Response1 {
	return &NullableGetSource200Response1{value: val, isSet: true}
}

func (v NullableGetSource200Response1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSource200Response1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
