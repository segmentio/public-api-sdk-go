/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API. 

API version: 32.0.7
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// ListWorkspaceRegulationsV1Output Output of all Workspace-scoped regulations in a given Workspace.
type ListWorkspaceRegulationsV1Output struct {
	// List of Workspace-scoped regulations with statuses.
	Regulations []RegulationListEntryV1 `json:"regulations"`
	Pagination Pagination `json:"pagination"`
}

// NewListWorkspaceRegulationsV1Output instantiates a new ListWorkspaceRegulationsV1Output object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListWorkspaceRegulationsV1Output(regulations []RegulationListEntryV1, pagination Pagination) *ListWorkspaceRegulationsV1Output {
	this := ListWorkspaceRegulationsV1Output{}
	this.Regulations = regulations
	this.Pagination = pagination
	return &this
}

// NewListWorkspaceRegulationsV1OutputWithDefaults instantiates a new ListWorkspaceRegulationsV1Output object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListWorkspaceRegulationsV1OutputWithDefaults() *ListWorkspaceRegulationsV1Output {
	this := ListWorkspaceRegulationsV1Output{}
	return &this
}

// GetRegulations returns the Regulations field value
func (o *ListWorkspaceRegulationsV1Output) GetRegulations() []RegulationListEntryV1 {
	if o == nil {
		var ret []RegulationListEntryV1
		return ret
	}

	return o.Regulations
}

// GetRegulationsOk returns a tuple with the Regulations field value
// and a boolean to check if the value has been set.
func (o *ListWorkspaceRegulationsV1Output) GetRegulationsOk() ([]RegulationListEntryV1, bool) {
	if o == nil {
		return nil, false
	}
	return o.Regulations, true
}

// SetRegulations sets field value
func (o *ListWorkspaceRegulationsV1Output) SetRegulations(v []RegulationListEntryV1) {
	o.Regulations = v
}

// GetPagination returns the Pagination field value
func (o *ListWorkspaceRegulationsV1Output) GetPagination() Pagination {
	if o == nil {
		var ret Pagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *ListWorkspaceRegulationsV1Output) GetPaginationOk() (*Pagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *ListWorkspaceRegulationsV1Output) SetPagination(v Pagination) {
	o.Pagination = v
}

func (o ListWorkspaceRegulationsV1Output) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["regulations"] = o.Regulations
	}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	return json.Marshal(toSerialize)
}

type NullableListWorkspaceRegulationsV1Output struct {
	value *ListWorkspaceRegulationsV1Output
	isSet bool
}

func (v NullableListWorkspaceRegulationsV1Output) Get() *ListWorkspaceRegulationsV1Output {
	return v.value
}

func (v *NullableListWorkspaceRegulationsV1Output) Set(val *ListWorkspaceRegulationsV1Output) {
	v.value = val
	v.isSet = true
}

func (v NullableListWorkspaceRegulationsV1Output) IsSet() bool {
	return v.isSet
}

func (v *NullableListWorkspaceRegulationsV1Output) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListWorkspaceRegulationsV1Output(val *ListWorkspaceRegulationsV1Output) *NullableListWorkspaceRegulationsV1Output {
	return &NullableListWorkspaceRegulationsV1Output{value: val, isSet: true}
}

func (v NullableListWorkspaceRegulationsV1Output) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListWorkspaceRegulationsV1Output) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


