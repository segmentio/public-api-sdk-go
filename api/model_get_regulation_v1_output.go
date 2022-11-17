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

// GetRegulationV1Output The regulate request returned.
type GetRegulationV1Output struct {
	Regulation Regulation `json:"regulation"`
}

// NewGetRegulationV1Output instantiates a new GetRegulationV1Output object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetRegulationV1Output(regulation Regulation) *GetRegulationV1Output {
	this := GetRegulationV1Output{}
	this.Regulation = regulation
	return &this
}

// NewGetRegulationV1OutputWithDefaults instantiates a new GetRegulationV1Output object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetRegulationV1OutputWithDefaults() *GetRegulationV1Output {
	this := GetRegulationV1Output{}
	return &this
}

// GetRegulation returns the Regulation field value
func (o *GetRegulationV1Output) GetRegulation() Regulation {
	if o == nil {
		var ret Regulation
		return ret
	}

	return o.Regulation
}

// GetRegulationOk returns a tuple with the Regulation field value
// and a boolean to check if the value has been set.
func (o *GetRegulationV1Output) GetRegulationOk() (*Regulation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Regulation, true
}

// SetRegulation sets field value
func (o *GetRegulationV1Output) SetRegulation(v Regulation) {
	o.Regulation = v
}

func (o GetRegulationV1Output) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["regulation"] = o.Regulation
	}
	return json.Marshal(toSerialize)
}

type NullableGetRegulationV1Output struct {
	value *GetRegulationV1Output
	isSet bool
}

func (v NullableGetRegulationV1Output) Get() *GetRegulationV1Output {
	return v.value
}

func (v *NullableGetRegulationV1Output) Set(val *GetRegulationV1Output) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRegulationV1Output) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRegulationV1Output) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRegulationV1Output(val *GetRegulationV1Output) *NullableGetRegulationV1Output {
	return &NullableGetRegulationV1Output{value: val, isSet: true}
}

func (v NullableGetRegulationV1Output) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetRegulationV1Output) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


