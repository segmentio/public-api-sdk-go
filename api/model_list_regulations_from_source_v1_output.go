/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API. 

API version: 32.0.8
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// ListRegulationsFromSourceV1Output Output of all Source-scoped regulations.
type ListRegulationsFromSourceV1Output struct {
	// List of Workspace-scoped regulations with statuses.
	Regulations []RegulationListEntryV1 `json:"regulations"`
}

// NewListRegulationsFromSourceV1Output instantiates a new ListRegulationsFromSourceV1Output object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListRegulationsFromSourceV1Output(regulations []RegulationListEntryV1) *ListRegulationsFromSourceV1Output {
	this := ListRegulationsFromSourceV1Output{}
	this.Regulations = regulations
	return &this
}

// NewListRegulationsFromSourceV1OutputWithDefaults instantiates a new ListRegulationsFromSourceV1Output object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListRegulationsFromSourceV1OutputWithDefaults() *ListRegulationsFromSourceV1Output {
	this := ListRegulationsFromSourceV1Output{}
	return &this
}

// GetRegulations returns the Regulations field value
func (o *ListRegulationsFromSourceV1Output) GetRegulations() []RegulationListEntryV1 {
	if o == nil {
		var ret []RegulationListEntryV1
		return ret
	}

	return o.Regulations
}

// GetRegulationsOk returns a tuple with the Regulations field value
// and a boolean to check if the value has been set.
func (o *ListRegulationsFromSourceV1Output) GetRegulationsOk() ([]RegulationListEntryV1, bool) {
	if o == nil {
		return nil, false
	}
	return o.Regulations, true
}

// SetRegulations sets field value
func (o *ListRegulationsFromSourceV1Output) SetRegulations(v []RegulationListEntryV1) {
	o.Regulations = v
}

func (o ListRegulationsFromSourceV1Output) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["regulations"] = o.Regulations
	}
	return json.Marshal(toSerialize)
}

type NullableListRegulationsFromSourceV1Output struct {
	value *ListRegulationsFromSourceV1Output
	isSet bool
}

func (v NullableListRegulationsFromSourceV1Output) Get() *ListRegulationsFromSourceV1Output {
	return v.value
}

func (v *NullableListRegulationsFromSourceV1Output) Set(val *ListRegulationsFromSourceV1Output) {
	v.value = val
	v.isSet = true
}

func (v NullableListRegulationsFromSourceV1Output) IsSet() bool {
	return v.isSet
}

func (v *NullableListRegulationsFromSourceV1Output) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListRegulationsFromSourceV1Output(val *ListRegulationsFromSourceV1Output) *NullableListRegulationsFromSourceV1Output {
	return &NullableListRegulationsFromSourceV1Output{value: val, isSet: true}
}

func (v NullableListRegulationsFromSourceV1Output) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListRegulationsFromSourceV1Output) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


