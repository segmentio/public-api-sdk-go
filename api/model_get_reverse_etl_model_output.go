/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 58.0.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the GetReverseEtlModelOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetReverseEtlModelOutput{}

// GetReverseEtlModelOutput Defines the result of getting a Model.
type GetReverseEtlModelOutput struct {
	ReverseEtlModel ReverseEtlModel `json:"reverseEtlModel"`
}

// NewGetReverseEtlModelOutput instantiates a new GetReverseEtlModelOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetReverseEtlModelOutput(reverseEtlModel ReverseEtlModel) *GetReverseEtlModelOutput {
	this := GetReverseEtlModelOutput{}
	this.ReverseEtlModel = reverseEtlModel
	return &this
}

// NewGetReverseEtlModelOutputWithDefaults instantiates a new GetReverseEtlModelOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetReverseEtlModelOutputWithDefaults() *GetReverseEtlModelOutput {
	this := GetReverseEtlModelOutput{}
	return &this
}

// GetReverseEtlModel returns the ReverseEtlModel field value
func (o *GetReverseEtlModelOutput) GetReverseEtlModel() ReverseEtlModel {
	if o == nil {
		var ret ReverseEtlModel
		return ret
	}

	return o.ReverseEtlModel
}

// GetReverseEtlModelOk returns a tuple with the ReverseEtlModel field value
// and a boolean to check if the value has been set.
func (o *GetReverseEtlModelOutput) GetReverseEtlModelOk() (*ReverseEtlModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReverseEtlModel, true
}

// SetReverseEtlModel sets field value
func (o *GetReverseEtlModelOutput) SetReverseEtlModel(v ReverseEtlModel) {
	o.ReverseEtlModel = v
}

func (o GetReverseEtlModelOutput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetReverseEtlModelOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["reverseEtlModel"] = o.ReverseEtlModel
	return toSerialize, nil
}

type NullableGetReverseEtlModelOutput struct {
	value *GetReverseEtlModelOutput
	isSet bool
}

func (v NullableGetReverseEtlModelOutput) Get() *GetReverseEtlModelOutput {
	return v.value
}

func (v *NullableGetReverseEtlModelOutput) Set(val *GetReverseEtlModelOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableGetReverseEtlModelOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableGetReverseEtlModelOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetReverseEtlModelOutput(
	val *GetReverseEtlModelOutput,
) *NullableGetReverseEtlModelOutput {
	return &NullableGetReverseEtlModelOutput{value: val, isSet: true}
}

func (v NullableGetReverseEtlModelOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetReverseEtlModelOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
