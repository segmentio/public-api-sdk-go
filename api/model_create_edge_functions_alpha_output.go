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

// checks if the CreateEdgeFunctionsAlphaOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateEdgeFunctionsAlphaOutput{}

// CreateEdgeFunctionsAlphaOutput Output for CreateEdgeFunctions.
type CreateEdgeFunctionsAlphaOutput struct {
	EdgeFunctions EdgeFunctionsAlpha `json:"edgeFunctions"`
}

// NewCreateEdgeFunctionsAlphaOutput instantiates a new CreateEdgeFunctionsAlphaOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateEdgeFunctionsAlphaOutput(
	edgeFunctions EdgeFunctionsAlpha,
) *CreateEdgeFunctionsAlphaOutput {
	this := CreateEdgeFunctionsAlphaOutput{}
	this.EdgeFunctions = edgeFunctions
	return &this
}

// NewCreateEdgeFunctionsAlphaOutputWithDefaults instantiates a new CreateEdgeFunctionsAlphaOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateEdgeFunctionsAlphaOutputWithDefaults() *CreateEdgeFunctionsAlphaOutput {
	this := CreateEdgeFunctionsAlphaOutput{}
	return &this
}

// GetEdgeFunctions returns the EdgeFunctions field value
func (o *CreateEdgeFunctionsAlphaOutput) GetEdgeFunctions() EdgeFunctionsAlpha {
	if o == nil {
		var ret EdgeFunctionsAlpha
		return ret
	}

	return o.EdgeFunctions
}

// GetEdgeFunctionsOk returns a tuple with the EdgeFunctions field value
// and a boolean to check if the value has been set.
func (o *CreateEdgeFunctionsAlphaOutput) GetEdgeFunctionsOk() (*EdgeFunctionsAlpha, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EdgeFunctions, true
}

// SetEdgeFunctions sets field value
func (o *CreateEdgeFunctionsAlphaOutput) SetEdgeFunctions(v EdgeFunctionsAlpha) {
	o.EdgeFunctions = v
}

func (o CreateEdgeFunctionsAlphaOutput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateEdgeFunctionsAlphaOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["edgeFunctions"] = o.EdgeFunctions
	return toSerialize, nil
}

type NullableCreateEdgeFunctionsAlphaOutput struct {
	value *CreateEdgeFunctionsAlphaOutput
	isSet bool
}

func (v NullableCreateEdgeFunctionsAlphaOutput) Get() *CreateEdgeFunctionsAlphaOutput {
	return v.value
}

func (v *NullableCreateEdgeFunctionsAlphaOutput) Set(val *CreateEdgeFunctionsAlphaOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateEdgeFunctionsAlphaOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateEdgeFunctionsAlphaOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateEdgeFunctionsAlphaOutput(
	val *CreateEdgeFunctionsAlphaOutput,
) *NullableCreateEdgeFunctionsAlphaOutput {
	return &NullableCreateEdgeFunctionsAlphaOutput{value: val, isSet: true}
}

func (v NullableCreateEdgeFunctionsAlphaOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateEdgeFunctionsAlphaOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
