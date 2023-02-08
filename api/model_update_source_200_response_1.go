/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 34.1.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// UpdateSource200Response1 struct for UpdateSource200Response1
type UpdateSource200Response1 struct {
	Data *UpdateSourceAlphaOutput `json:"data,omitempty"`
}

// NewUpdateSource200Response1 instantiates a new UpdateSource200Response1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateSource200Response1() *UpdateSource200Response1 {
	this := UpdateSource200Response1{}
	return &this
}

// NewUpdateSource200Response1WithDefaults instantiates a new UpdateSource200Response1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateSource200Response1WithDefaults() *UpdateSource200Response1 {
	this := UpdateSource200Response1{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *UpdateSource200Response1) GetData() UpdateSourceAlphaOutput {
	if o == nil || o.Data == nil {
		var ret UpdateSourceAlphaOutput
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSource200Response1) GetDataOk() (*UpdateSourceAlphaOutput, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *UpdateSource200Response1) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given UpdateSourceAlphaOutput and assigns it to the Data field.
func (o *UpdateSource200Response1) SetData(v UpdateSourceAlphaOutput) {
	o.Data = &v
}

func (o UpdateSource200Response1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateSource200Response1 struct {
	value *UpdateSource200Response1
	isSet bool
}

func (v NullableUpdateSource200Response1) Get() *UpdateSource200Response1 {
	return v.value
}

func (v *NullableUpdateSource200Response1) Set(val *UpdateSource200Response1) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateSource200Response1) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateSource200Response1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateSource200Response1(
	val *UpdateSource200Response1,
) *NullableUpdateSource200Response1 {
	return &NullableUpdateSource200Response1{value: val, isSet: true}
}

func (v NullableUpdateSource200Response1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateSource200Response1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
