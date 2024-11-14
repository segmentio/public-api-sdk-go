/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 57.0.1
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the CreateWorkspaceRegulationV1Output type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateWorkspaceRegulationV1Output{}

// CreateWorkspaceRegulationV1Output The output of a create Workspace regulation call.
type CreateWorkspaceRegulationV1Output struct {
	// The id of the created regulation.
	RegulateId string `json:"regulateId"`
}

// NewCreateWorkspaceRegulationV1Output instantiates a new CreateWorkspaceRegulationV1Output object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateWorkspaceRegulationV1Output(regulateId string) *CreateWorkspaceRegulationV1Output {
	this := CreateWorkspaceRegulationV1Output{}
	this.RegulateId = regulateId
	return &this
}

// NewCreateWorkspaceRegulationV1OutputWithDefaults instantiates a new CreateWorkspaceRegulationV1Output object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateWorkspaceRegulationV1OutputWithDefaults() *CreateWorkspaceRegulationV1Output {
	this := CreateWorkspaceRegulationV1Output{}
	return &this
}

// GetRegulateId returns the RegulateId field value
func (o *CreateWorkspaceRegulationV1Output) GetRegulateId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RegulateId
}

// GetRegulateIdOk returns a tuple with the RegulateId field value
// and a boolean to check if the value has been set.
func (o *CreateWorkspaceRegulationV1Output) GetRegulateIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RegulateId, true
}

// SetRegulateId sets field value
func (o *CreateWorkspaceRegulationV1Output) SetRegulateId(v string) {
	o.RegulateId = v
}

func (o CreateWorkspaceRegulationV1Output) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateWorkspaceRegulationV1Output) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["regulateId"] = o.RegulateId
	return toSerialize, nil
}

type NullableCreateWorkspaceRegulationV1Output struct {
	value *CreateWorkspaceRegulationV1Output
	isSet bool
}

func (v NullableCreateWorkspaceRegulationV1Output) Get() *CreateWorkspaceRegulationV1Output {
	return v.value
}

func (v *NullableCreateWorkspaceRegulationV1Output) Set(val *CreateWorkspaceRegulationV1Output) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateWorkspaceRegulationV1Output) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateWorkspaceRegulationV1Output) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateWorkspaceRegulationV1Output(
	val *CreateWorkspaceRegulationV1Output,
) *NullableCreateWorkspaceRegulationV1Output {
	return &NullableCreateWorkspaceRegulationV1Output{value: val, isSet: true}
}

func (v NullableCreateWorkspaceRegulationV1Output) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateWorkspaceRegulationV1Output) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
