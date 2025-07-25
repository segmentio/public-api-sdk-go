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

// checks if the CreateSourceRegulationV1Output type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateSourceRegulationV1Output{}

// CreateSourceRegulationV1Output The output of a create Source regulation call.
type CreateSourceRegulationV1Output struct {
	// The id of the created regulation.
	RegulateId string `json:"regulateId"`
}

// NewCreateSourceRegulationV1Output instantiates a new CreateSourceRegulationV1Output object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSourceRegulationV1Output(regulateId string) *CreateSourceRegulationV1Output {
	this := CreateSourceRegulationV1Output{}
	this.RegulateId = regulateId
	return &this
}

// NewCreateSourceRegulationV1OutputWithDefaults instantiates a new CreateSourceRegulationV1Output object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSourceRegulationV1OutputWithDefaults() *CreateSourceRegulationV1Output {
	this := CreateSourceRegulationV1Output{}
	return &this
}

// GetRegulateId returns the RegulateId field value
func (o *CreateSourceRegulationV1Output) GetRegulateId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RegulateId
}

// GetRegulateIdOk returns a tuple with the RegulateId field value
// and a boolean to check if the value has been set.
func (o *CreateSourceRegulationV1Output) GetRegulateIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RegulateId, true
}

// SetRegulateId sets field value
func (o *CreateSourceRegulationV1Output) SetRegulateId(v string) {
	o.RegulateId = v
}

func (o CreateSourceRegulationV1Output) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateSourceRegulationV1Output) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["regulateId"] = o.RegulateId
	return toSerialize, nil
}

type NullableCreateSourceRegulationV1Output struct {
	value *CreateSourceRegulationV1Output
	isSet bool
}

func (v NullableCreateSourceRegulationV1Output) Get() *CreateSourceRegulationV1Output {
	return v.value
}

func (v *NullableCreateSourceRegulationV1Output) Set(val *CreateSourceRegulationV1Output) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSourceRegulationV1Output) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSourceRegulationV1Output) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSourceRegulationV1Output(
	val *CreateSourceRegulationV1Output,
) *NullableCreateSourceRegulationV1Output {
	return &NullableCreateSourceRegulationV1Output{value: val, isSet: true}
}

func (v NullableCreateSourceRegulationV1Output) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateSourceRegulationV1Output) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
