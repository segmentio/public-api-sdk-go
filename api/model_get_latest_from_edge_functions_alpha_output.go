/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 37.1.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// GetLatestFromEdgeFunctionsAlphaOutput Output for GetLatestFromEdgeFunctions.
type GetLatestFromEdgeFunctionsAlphaOutput struct {
	EdgeFunctions EdgeFunctions1 `json:"edgeFunctions"`
}

// NewGetLatestFromEdgeFunctionsAlphaOutput instantiates a new GetLatestFromEdgeFunctionsAlphaOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetLatestFromEdgeFunctionsAlphaOutput(
	edgeFunctions EdgeFunctions1,
) *GetLatestFromEdgeFunctionsAlphaOutput {
	this := GetLatestFromEdgeFunctionsAlphaOutput{}
	this.EdgeFunctions = edgeFunctions
	return &this
}

// NewGetLatestFromEdgeFunctionsAlphaOutputWithDefaults instantiates a new GetLatestFromEdgeFunctionsAlphaOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetLatestFromEdgeFunctionsAlphaOutputWithDefaults() *GetLatestFromEdgeFunctionsAlphaOutput {
	this := GetLatestFromEdgeFunctionsAlphaOutput{}
	return &this
}

// GetEdgeFunctions returns the EdgeFunctions field value
func (o *GetLatestFromEdgeFunctionsAlphaOutput) GetEdgeFunctions() EdgeFunctions1 {
	if o == nil {
		var ret EdgeFunctions1
		return ret
	}

	return o.EdgeFunctions
}

// GetEdgeFunctionsOk returns a tuple with the EdgeFunctions field value
// and a boolean to check if the value has been set.
func (o *GetLatestFromEdgeFunctionsAlphaOutput) GetEdgeFunctionsOk() (*EdgeFunctions1, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EdgeFunctions, true
}

// SetEdgeFunctions sets field value
func (o *GetLatestFromEdgeFunctionsAlphaOutput) SetEdgeFunctions(v EdgeFunctions1) {
	o.EdgeFunctions = v
}

func (o GetLatestFromEdgeFunctionsAlphaOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["edgeFunctions"] = o.EdgeFunctions
	}
	return json.Marshal(toSerialize)
}

type NullableGetLatestFromEdgeFunctionsAlphaOutput struct {
	value *GetLatestFromEdgeFunctionsAlphaOutput
	isSet bool
}

func (v NullableGetLatestFromEdgeFunctionsAlphaOutput) Get() *GetLatestFromEdgeFunctionsAlphaOutput {
	return v.value
}

func (v *NullableGetLatestFromEdgeFunctionsAlphaOutput) Set(
	val *GetLatestFromEdgeFunctionsAlphaOutput,
) {
	v.value = val
	v.isSet = true
}

func (v NullableGetLatestFromEdgeFunctionsAlphaOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableGetLatestFromEdgeFunctionsAlphaOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetLatestFromEdgeFunctionsAlphaOutput(
	val *GetLatestFromEdgeFunctionsAlphaOutput,
) *NullableGetLatestFromEdgeFunctionsAlphaOutput {
	return &NullableGetLatestFromEdgeFunctionsAlphaOutput{value: val, isSet: true}
}

func (v NullableGetLatestFromEdgeFunctionsAlphaOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetLatestFromEdgeFunctionsAlphaOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
