/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 38.3.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the GetWorkspaceV1Output type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetWorkspaceV1Output{}

// GetWorkspaceV1Output Represents the output of loading the Workspace.
type GetWorkspaceV1Output struct {
	Workspace WorkspaceV1 `json:"workspace"`
}

// NewGetWorkspaceV1Output instantiates a new GetWorkspaceV1Output object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetWorkspaceV1Output(workspace WorkspaceV1) *GetWorkspaceV1Output {
	this := GetWorkspaceV1Output{}
	this.Workspace = workspace
	return &this
}

// NewGetWorkspaceV1OutputWithDefaults instantiates a new GetWorkspaceV1Output object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetWorkspaceV1OutputWithDefaults() *GetWorkspaceV1Output {
	this := GetWorkspaceV1Output{}
	return &this
}

// GetWorkspace returns the Workspace field value
func (o *GetWorkspaceV1Output) GetWorkspace() WorkspaceV1 {
	if o == nil {
		var ret WorkspaceV1
		return ret
	}

	return o.Workspace
}

// GetWorkspaceOk returns a tuple with the Workspace field value
// and a boolean to check if the value has been set.
func (o *GetWorkspaceV1Output) GetWorkspaceOk() (*WorkspaceV1, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Workspace, true
}

// SetWorkspace sets field value
func (o *GetWorkspaceV1Output) SetWorkspace(v WorkspaceV1) {
	o.Workspace = v
}

func (o GetWorkspaceV1Output) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetWorkspaceV1Output) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["workspace"] = o.Workspace
	return toSerialize, nil
}

type NullableGetWorkspaceV1Output struct {
	value *GetWorkspaceV1Output
	isSet bool
}

func (v NullableGetWorkspaceV1Output) Get() *GetWorkspaceV1Output {
	return v.value
}

func (v *NullableGetWorkspaceV1Output) Set(val *GetWorkspaceV1Output) {
	v.value = val
	v.isSet = true
}

func (v NullableGetWorkspaceV1Output) IsSet() bool {
	return v.isSet
}

func (v *NullableGetWorkspaceV1Output) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetWorkspaceV1Output(val *GetWorkspaceV1Output) *NullableGetWorkspaceV1Output {
	return &NullableGetWorkspaceV1Output{value: val, isSet: true}
}

func (v NullableGetWorkspaceV1Output) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetWorkspaceV1Output) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
