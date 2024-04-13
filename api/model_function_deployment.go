/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 48.0.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the FunctionDeployment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FunctionDeployment{}

// FunctionDeployment The status of the operation.
type FunctionDeployment struct {
	Status string `json:"status"`
}

// NewFunctionDeployment instantiates a new FunctionDeployment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFunctionDeployment(status string) *FunctionDeployment {
	this := FunctionDeployment{}
	this.Status = status
	return &this
}

// NewFunctionDeploymentWithDefaults instantiates a new FunctionDeployment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFunctionDeploymentWithDefaults() *FunctionDeployment {
	this := FunctionDeployment{}
	return &this
}

// GetStatus returns the Status field value
func (o *FunctionDeployment) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *FunctionDeployment) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *FunctionDeployment) SetStatus(v string) {
	o.Status = v
}

func (o FunctionDeployment) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FunctionDeployment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

type NullableFunctionDeployment struct {
	value *FunctionDeployment
	isSet bool
}

func (v NullableFunctionDeployment) Get() *FunctionDeployment {
	return v.value
}

func (v *NullableFunctionDeployment) Set(val *FunctionDeployment) {
	v.value = val
	v.isSet = true
}

func (v NullableFunctionDeployment) IsSet() bool {
	return v.isSet
}

func (v *NullableFunctionDeployment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFunctionDeployment(val *FunctionDeployment) *NullableFunctionDeployment {
	return &NullableFunctionDeployment{value: val, isSet: true}
}

func (v NullableFunctionDeployment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFunctionDeployment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
