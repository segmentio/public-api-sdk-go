/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 38.0.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the RestoreFunctionVersionAlphaOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RestoreFunctionVersionAlphaOutput{}

// RestoreFunctionVersionAlphaOutput Restore version output.
type RestoreFunctionVersionAlphaOutput struct {
	Version Version `json:"version"`
}

// NewRestoreFunctionVersionAlphaOutput instantiates a new RestoreFunctionVersionAlphaOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRestoreFunctionVersionAlphaOutput(version Version) *RestoreFunctionVersionAlphaOutput {
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
func (o *RestoreFunctionVersionAlphaOutput) GetVersion() Version {
	if o == nil {
		var ret Version
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *RestoreFunctionVersionAlphaOutput) GetVersionOk() (*Version, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *RestoreFunctionVersionAlphaOutput) SetVersion(v Version) {
	o.Version = v
}

func (o RestoreFunctionVersionAlphaOutput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RestoreFunctionVersionAlphaOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["version"] = o.Version
	return toSerialize, nil
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
