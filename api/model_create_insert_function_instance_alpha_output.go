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

// checks if the CreateInsertFunctionInstanceAlphaOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateInsertFunctionInstanceAlphaOutput{}

// CreateInsertFunctionInstanceAlphaOutput Creates an insert Function instance.
type CreateInsertFunctionInstanceAlphaOutput struct {
	InsertFunctionInstance InsertFunctionInstanceAlpha `json:"insertFunctionInstance"`
}

// NewCreateInsertFunctionInstanceAlphaOutput instantiates a new CreateInsertFunctionInstanceAlphaOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateInsertFunctionInstanceAlphaOutput(
	insertFunctionInstance InsertFunctionInstanceAlpha,
) *CreateInsertFunctionInstanceAlphaOutput {
	this := CreateInsertFunctionInstanceAlphaOutput{}
	this.InsertFunctionInstance = insertFunctionInstance
	return &this
}

// NewCreateInsertFunctionInstanceAlphaOutputWithDefaults instantiates a new CreateInsertFunctionInstanceAlphaOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateInsertFunctionInstanceAlphaOutputWithDefaults() *CreateInsertFunctionInstanceAlphaOutput {
	this := CreateInsertFunctionInstanceAlphaOutput{}
	return &this
}

// GetInsertFunctionInstance returns the InsertFunctionInstance field value
func (o *CreateInsertFunctionInstanceAlphaOutput) GetInsertFunctionInstance() InsertFunctionInstanceAlpha {
	if o == nil {
		var ret InsertFunctionInstanceAlpha
		return ret
	}

	return o.InsertFunctionInstance
}

// GetInsertFunctionInstanceOk returns a tuple with the InsertFunctionInstance field value
// and a boolean to check if the value has been set.
func (o *CreateInsertFunctionInstanceAlphaOutput) GetInsertFunctionInstanceOk() (*InsertFunctionInstanceAlpha, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InsertFunctionInstance, true
}

// SetInsertFunctionInstance sets field value
func (o *CreateInsertFunctionInstanceAlphaOutput) SetInsertFunctionInstance(
	v InsertFunctionInstanceAlpha,
) {
	o.InsertFunctionInstance = v
}

func (o CreateInsertFunctionInstanceAlphaOutput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateInsertFunctionInstanceAlphaOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["insertFunctionInstance"] = o.InsertFunctionInstance
	return toSerialize, nil
}

type NullableCreateInsertFunctionInstanceAlphaOutput struct {
	value *CreateInsertFunctionInstanceAlphaOutput
	isSet bool
}

func (v NullableCreateInsertFunctionInstanceAlphaOutput) Get() *CreateInsertFunctionInstanceAlphaOutput {
	return v.value
}

func (v *NullableCreateInsertFunctionInstanceAlphaOutput) Set(
	val *CreateInsertFunctionInstanceAlphaOutput,
) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateInsertFunctionInstanceAlphaOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateInsertFunctionInstanceAlphaOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateInsertFunctionInstanceAlphaOutput(
	val *CreateInsertFunctionInstanceAlphaOutput,
) *NullableCreateInsertFunctionInstanceAlphaOutput {
	return &NullableCreateInsertFunctionInstanceAlphaOutput{value: val, isSet: true}
}

func (v NullableCreateInsertFunctionInstanceAlphaOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateInsertFunctionInstanceAlphaOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
