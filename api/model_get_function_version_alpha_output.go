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

// checks if the GetFunctionVersionAlphaOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetFunctionVersionAlphaOutput{}

// GetFunctionVersionAlphaOutput Get Function version output.
type GetFunctionVersionAlphaOutput struct {
	Version Version `json:"version"`
}

// NewGetFunctionVersionAlphaOutput instantiates a new GetFunctionVersionAlphaOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFunctionVersionAlphaOutput(version Version) *GetFunctionVersionAlphaOutput {
	this := GetFunctionVersionAlphaOutput{}
	this.Version = version
	return &this
}

// NewGetFunctionVersionAlphaOutputWithDefaults instantiates a new GetFunctionVersionAlphaOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFunctionVersionAlphaOutputWithDefaults() *GetFunctionVersionAlphaOutput {
	this := GetFunctionVersionAlphaOutput{}
	return &this
}

// GetVersion returns the Version field value
func (o *GetFunctionVersionAlphaOutput) GetVersion() Version {
	if o == nil {
		var ret Version
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *GetFunctionVersionAlphaOutput) GetVersionOk() (*Version, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *GetFunctionVersionAlphaOutput) SetVersion(v Version) {
	o.Version = v
}

func (o GetFunctionVersionAlphaOutput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetFunctionVersionAlphaOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["version"] = o.Version
	return toSerialize, nil
}

type NullableGetFunctionVersionAlphaOutput struct {
	value *GetFunctionVersionAlphaOutput
	isSet bool
}

func (v NullableGetFunctionVersionAlphaOutput) Get() *GetFunctionVersionAlphaOutput {
	return v.value
}

func (v *NullableGetFunctionVersionAlphaOutput) Set(val *GetFunctionVersionAlphaOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFunctionVersionAlphaOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFunctionVersionAlphaOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFunctionVersionAlphaOutput(
	val *GetFunctionVersionAlphaOutput,
) *NullableGetFunctionVersionAlphaOutput {
	return &NullableGetFunctionVersionAlphaOutput{value: val, isSet: true}
}

func (v NullableGetFunctionVersionAlphaOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFunctionVersionAlphaOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
