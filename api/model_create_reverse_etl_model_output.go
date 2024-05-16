/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 50.1.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the CreateReverseEtlModelOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateReverseEtlModelOutput{}

// CreateReverseEtlModelOutput Defines the results of creating a Model.
type CreateReverseEtlModelOutput struct {
	ReverseEtlModel ReverseEtlModel `json:"reverseEtlModel"`
}

// NewCreateReverseEtlModelOutput instantiates a new CreateReverseEtlModelOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateReverseEtlModelOutput(reverseEtlModel ReverseEtlModel) *CreateReverseEtlModelOutput {
	this := CreateReverseEtlModelOutput{}
	this.ReverseEtlModel = reverseEtlModel
	return &this
}

// NewCreateReverseEtlModelOutputWithDefaults instantiates a new CreateReverseEtlModelOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateReverseEtlModelOutputWithDefaults() *CreateReverseEtlModelOutput {
	this := CreateReverseEtlModelOutput{}
	return &this
}

// GetReverseEtlModel returns the ReverseEtlModel field value
func (o *CreateReverseEtlModelOutput) GetReverseEtlModel() ReverseEtlModel {
	if o == nil {
		var ret ReverseEtlModel
		return ret
	}

	return o.ReverseEtlModel
}

// GetReverseEtlModelOk returns a tuple with the ReverseEtlModel field value
// and a boolean to check if the value has been set.
func (o *CreateReverseEtlModelOutput) GetReverseEtlModelOk() (*ReverseEtlModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReverseEtlModel, true
}

// SetReverseEtlModel sets field value
func (o *CreateReverseEtlModelOutput) SetReverseEtlModel(v ReverseEtlModel) {
	o.ReverseEtlModel = v
}

func (o CreateReverseEtlModelOutput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateReverseEtlModelOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["reverseEtlModel"] = o.ReverseEtlModel
	return toSerialize, nil
}

type NullableCreateReverseEtlModelOutput struct {
	value *CreateReverseEtlModelOutput
	isSet bool
}

func (v NullableCreateReverseEtlModelOutput) Get() *CreateReverseEtlModelOutput {
	return v.value
}

func (v *NullableCreateReverseEtlModelOutput) Set(val *CreateReverseEtlModelOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateReverseEtlModelOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateReverseEtlModelOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateReverseEtlModelOutput(
	val *CreateReverseEtlModelOutput,
) *NullableCreateReverseEtlModelOutput {
	return &NullableCreateReverseEtlModelOutput{value: val, isSet: true}
}

func (v NullableCreateReverseEtlModelOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateReverseEtlModelOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
