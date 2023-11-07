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

// checks if the GetSpaceAlphaOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSpaceAlphaOutput{}

// GetSpaceAlphaOutput Response for the getSpaceById endpoint.
type GetSpaceAlphaOutput struct {
	Space NullableSpace `json:"space"`
}

// NewGetSpaceAlphaOutput instantiates a new GetSpaceAlphaOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSpaceAlphaOutput(space NullableSpace) *GetSpaceAlphaOutput {
	this := GetSpaceAlphaOutput{}
	this.Space = space
	return &this
}

// NewGetSpaceAlphaOutputWithDefaults instantiates a new GetSpaceAlphaOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSpaceAlphaOutputWithDefaults() *GetSpaceAlphaOutput {
	this := GetSpaceAlphaOutput{}
	return &this
}

// GetSpace returns the Space field value
// If the value is explicit nil, the zero value for Space will be returned
func (o *GetSpaceAlphaOutput) GetSpace() Space {
	if o == nil || o.Space.Get() == nil {
		var ret Space
		return ret
	}

	return *o.Space.Get()
}

// GetSpaceOk returns a tuple with the Space field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetSpaceAlphaOutput) GetSpaceOk() (*Space, bool) {
	if o == nil {
		return nil, false
	}
	return o.Space.Get(), o.Space.IsSet()
}

// SetSpace sets field value
func (o *GetSpaceAlphaOutput) SetSpace(v Space) {
	o.Space.Set(&v)
}

func (o GetSpaceAlphaOutput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSpaceAlphaOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["space"] = o.Space.Get()
	return toSerialize, nil
}

type NullableGetSpaceAlphaOutput struct {
	value *GetSpaceAlphaOutput
	isSet bool
}

func (v NullableGetSpaceAlphaOutput) Get() *GetSpaceAlphaOutput {
	return v.value
}

func (v *NullableGetSpaceAlphaOutput) Set(val *GetSpaceAlphaOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSpaceAlphaOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSpaceAlphaOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSpaceAlphaOutput(val *GetSpaceAlphaOutput) *NullableGetSpaceAlphaOutput {
	return &NullableGetSpaceAlphaOutput{value: val, isSet: true}
}

func (v NullableGetSpaceAlphaOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSpaceAlphaOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
