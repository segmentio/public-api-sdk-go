/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 37.1.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// UpdateInsertFunctionInstanceAlphaOutput Returns the updated insert Function instance.
type UpdateInsertFunctionInstanceAlphaOutput struct {
	InsertFunctionInstance InsertFunctionInstance1 `json:"insertFunctionInstance"`
}

// NewUpdateInsertFunctionInstanceAlphaOutput instantiates a new UpdateInsertFunctionInstanceAlphaOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateInsertFunctionInstanceAlphaOutput(
	insertFunctionInstance InsertFunctionInstance1,
) *UpdateInsertFunctionInstanceAlphaOutput {
	this := UpdateInsertFunctionInstanceAlphaOutput{}
	this.InsertFunctionInstance = insertFunctionInstance
	return &this
}

// NewUpdateInsertFunctionInstanceAlphaOutputWithDefaults instantiates a new UpdateInsertFunctionInstanceAlphaOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateInsertFunctionInstanceAlphaOutputWithDefaults() *UpdateInsertFunctionInstanceAlphaOutput {
	this := UpdateInsertFunctionInstanceAlphaOutput{}
	return &this
}

// GetInsertFunctionInstance returns the InsertFunctionInstance field value
func (o *UpdateInsertFunctionInstanceAlphaOutput) GetInsertFunctionInstance() InsertFunctionInstance1 {
	if o == nil {
		var ret InsertFunctionInstance1
		return ret
	}

	return o.InsertFunctionInstance
}

// GetInsertFunctionInstanceOk returns a tuple with the InsertFunctionInstance field value
// and a boolean to check if the value has been set.
func (o *UpdateInsertFunctionInstanceAlphaOutput) GetInsertFunctionInstanceOk() (*InsertFunctionInstance1, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InsertFunctionInstance, true
}

// SetInsertFunctionInstance sets field value
func (o *UpdateInsertFunctionInstanceAlphaOutput) SetInsertFunctionInstance(
	v InsertFunctionInstance1,
) {
	o.InsertFunctionInstance = v
}

func (o UpdateInsertFunctionInstanceAlphaOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["insertFunctionInstance"] = o.InsertFunctionInstance
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateInsertFunctionInstanceAlphaOutput struct {
	value *UpdateInsertFunctionInstanceAlphaOutput
	isSet bool
}

func (v NullableUpdateInsertFunctionInstanceAlphaOutput) Get() *UpdateInsertFunctionInstanceAlphaOutput {
	return v.value
}

func (v *NullableUpdateInsertFunctionInstanceAlphaOutput) Set(
	val *UpdateInsertFunctionInstanceAlphaOutput,
) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateInsertFunctionInstanceAlphaOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateInsertFunctionInstanceAlphaOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateInsertFunctionInstanceAlphaOutput(
	val *UpdateInsertFunctionInstanceAlphaOutput,
) *NullableUpdateInsertFunctionInstanceAlphaOutput {
	return &NullableUpdateInsertFunctionInstanceAlphaOutput{value: val, isSet: true}
}

func (v NullableUpdateInsertFunctionInstanceAlphaOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateInsertFunctionInstanceAlphaOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
