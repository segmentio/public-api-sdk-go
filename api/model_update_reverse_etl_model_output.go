/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 37.0.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// UpdateReverseEtlModelOutput Defines the results of updating a Model.
type UpdateReverseEtlModelOutput struct {
	ReverseEtlModel ReverseEtlModel `json:"reverseEtlModel"`
}

// NewUpdateReverseEtlModelOutput instantiates a new UpdateReverseEtlModelOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateReverseEtlModelOutput(reverseEtlModel ReverseEtlModel) *UpdateReverseEtlModelOutput {
	this := UpdateReverseEtlModelOutput{}
	this.ReverseEtlModel = reverseEtlModel
	return &this
}

// NewUpdateReverseEtlModelOutputWithDefaults instantiates a new UpdateReverseEtlModelOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateReverseEtlModelOutputWithDefaults() *UpdateReverseEtlModelOutput {
	this := UpdateReverseEtlModelOutput{}
	return &this
}

// GetReverseEtlModel returns the ReverseEtlModel field value
func (o *UpdateReverseEtlModelOutput) GetReverseEtlModel() ReverseEtlModel {
	if o == nil {
		var ret ReverseEtlModel
		return ret
	}

	return o.ReverseEtlModel
}

// GetReverseEtlModelOk returns a tuple with the ReverseEtlModel field value
// and a boolean to check if the value has been set.
func (o *UpdateReverseEtlModelOutput) GetReverseEtlModelOk() (*ReverseEtlModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReverseEtlModel, true
}

// SetReverseEtlModel sets field value
func (o *UpdateReverseEtlModelOutput) SetReverseEtlModel(v ReverseEtlModel) {
	o.ReverseEtlModel = v
}

func (o UpdateReverseEtlModelOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["reverseEtlModel"] = o.ReverseEtlModel
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateReverseEtlModelOutput struct {
	value *UpdateReverseEtlModelOutput
	isSet bool
}

func (v NullableUpdateReverseEtlModelOutput) Get() *UpdateReverseEtlModelOutput {
	return v.value
}

func (v *NullableUpdateReverseEtlModelOutput) Set(val *UpdateReverseEtlModelOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateReverseEtlModelOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateReverseEtlModelOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateReverseEtlModelOutput(
	val *UpdateReverseEtlModelOutput,
) *NullableUpdateReverseEtlModelOutput {
	return &NullableUpdateReverseEtlModelOutput{value: val, isSet: true}
}

func (v NullableUpdateReverseEtlModelOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateReverseEtlModelOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
