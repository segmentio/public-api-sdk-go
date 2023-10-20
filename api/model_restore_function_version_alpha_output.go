/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 36.3.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// RestoreFunctionVersionAlphaOutput Restore version output.
type RestoreFunctionVersionAlphaOutput struct {
	Version Version1 `json:"version"`
}

// NewRestoreFunctionVersionAlphaOutput instantiates a new RestoreFunctionVersionAlphaOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRestoreFunctionVersionAlphaOutput(version Version1) *RestoreFunctionVersionAlphaOutput {
	this := RestoreFunctionVersionAlphaOutput{}
	this.Version = version
	return &this
}

// NewRestoreFunctionVersionAlphaOutputWithDefaults instantiates a new RestoreFunctionVersionAlphaOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRestoreFunctionVersionAlphaOutputWithDefaults() *RestoreFunctionVersionAlphaOutput {
	this := RestoreFunctionVersionAlphaOutput{}
	return &this
}

// GetVersion returns the Version field value
func (o *RestoreFunctionVersionAlphaOutput) GetVersion() Version1 {
	if o == nil {
		var ret Version1
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *RestoreFunctionVersionAlphaOutput) GetVersionOk() (*Version1, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *RestoreFunctionVersionAlphaOutput) SetVersion(v Version1) {
	o.Version = v
}

func (o RestoreFunctionVersionAlphaOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableRestoreFunctionVersionAlphaOutput struct {
	value *RestoreFunctionVersionAlphaOutput
	isSet bool
}

func (v NullableRestoreFunctionVersionAlphaOutput) Get() *RestoreFunctionVersionAlphaOutput {
	return v.value
}

func (v *NullableRestoreFunctionVersionAlphaOutput) Set(val *RestoreFunctionVersionAlphaOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableRestoreFunctionVersionAlphaOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableRestoreFunctionVersionAlphaOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRestoreFunctionVersionAlphaOutput(
	val *RestoreFunctionVersionAlphaOutput,
) *NullableRestoreFunctionVersionAlphaOutput {
	return &NullableRestoreFunctionVersionAlphaOutput{value: val, isSet: true}
}

func (v NullableRestoreFunctionVersionAlphaOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRestoreFunctionVersionAlphaOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
