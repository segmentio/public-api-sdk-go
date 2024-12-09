/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 57.2.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the CreateCloudSourceRegulationV1Output type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateCloudSourceRegulationV1Output{}

// CreateCloudSourceRegulationV1Output The output of a create Cloud Source regulation call.
type CreateCloudSourceRegulationV1Output struct {
	// The id of the created regulation.
	RegulateId string `json:"regulateId"`
}

// NewCreateCloudSourceRegulationV1Output instantiates a new CreateCloudSourceRegulationV1Output object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateCloudSourceRegulationV1Output(
	regulateId string,
) *CreateCloudSourceRegulationV1Output {
	this := CreateCloudSourceRegulationV1Output{}
	this.RegulateId = regulateId
	return &this
}

// NewCreateCloudSourceRegulationV1OutputWithDefaults instantiates a new CreateCloudSourceRegulationV1Output object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateCloudSourceRegulationV1OutputWithDefaults() *CreateCloudSourceRegulationV1Output {
	this := CreateCloudSourceRegulationV1Output{}
	return &this
}

// GetRegulateId returns the RegulateId field value
func (o *CreateCloudSourceRegulationV1Output) GetRegulateId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RegulateId
}

// GetRegulateIdOk returns a tuple with the RegulateId field value
// and a boolean to check if the value has been set.
func (o *CreateCloudSourceRegulationV1Output) GetRegulateIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RegulateId, true
}

// SetRegulateId sets field value
func (o *CreateCloudSourceRegulationV1Output) SetRegulateId(v string) {
	o.RegulateId = v
}

func (o CreateCloudSourceRegulationV1Output) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateCloudSourceRegulationV1Output) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["regulateId"] = o.RegulateId
	return toSerialize, nil
}

type NullableCreateCloudSourceRegulationV1Output struct {
	value *CreateCloudSourceRegulationV1Output
	isSet bool
}

func (v NullableCreateCloudSourceRegulationV1Output) Get() *CreateCloudSourceRegulationV1Output {
	return v.value
}

func (v *NullableCreateCloudSourceRegulationV1Output) Set(
	val *CreateCloudSourceRegulationV1Output,
) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateCloudSourceRegulationV1Output) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateCloudSourceRegulationV1Output) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateCloudSourceRegulationV1Output(
	val *CreateCloudSourceRegulationV1Output,
) *NullableCreateCloudSourceRegulationV1Output {
	return &NullableCreateCloudSourceRegulationV1Output{value: val, isSet: true}
}

func (v NullableCreateCloudSourceRegulationV1Output) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateCloudSourceRegulationV1Output) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
