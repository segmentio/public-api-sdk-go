/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 57.1.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the UpdateComputedTraitForSpace200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateComputedTraitForSpace200Response{}

// UpdateComputedTraitForSpace200Response struct for UpdateComputedTraitForSpace200Response
type UpdateComputedTraitForSpace200Response struct {
	Data *UpdateComputedTraitForSpaceAlphaOutput `json:"data,omitempty"`
}

// NewUpdateComputedTraitForSpace200Response instantiates a new UpdateComputedTraitForSpace200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateComputedTraitForSpace200Response() *UpdateComputedTraitForSpace200Response {
	this := UpdateComputedTraitForSpace200Response{}
	return &this
}

// NewUpdateComputedTraitForSpace200ResponseWithDefaults instantiates a new UpdateComputedTraitForSpace200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateComputedTraitForSpace200ResponseWithDefaults() *UpdateComputedTraitForSpace200Response {
	this := UpdateComputedTraitForSpace200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *UpdateComputedTraitForSpace200Response) GetData() UpdateComputedTraitForSpaceAlphaOutput {
	if o == nil || IsNil(o.Data) {
		var ret UpdateComputedTraitForSpaceAlphaOutput
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateComputedTraitForSpace200Response) GetDataOk() (*UpdateComputedTraitForSpaceAlphaOutput, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *UpdateComputedTraitForSpace200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given UpdateComputedTraitForSpaceAlphaOutput and assigns it to the Data field.
func (o *UpdateComputedTraitForSpace200Response) SetData(v UpdateComputedTraitForSpaceAlphaOutput) {
	o.Data = &v
}

func (o UpdateComputedTraitForSpace200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateComputedTraitForSpace200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableUpdateComputedTraitForSpace200Response struct {
	value *UpdateComputedTraitForSpace200Response
	isSet bool
}

func (v NullableUpdateComputedTraitForSpace200Response) Get() *UpdateComputedTraitForSpace200Response {
	return v.value
}

func (v *NullableUpdateComputedTraitForSpace200Response) Set(
	val *UpdateComputedTraitForSpace200Response,
) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateComputedTraitForSpace200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateComputedTraitForSpace200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateComputedTraitForSpace200Response(
	val *UpdateComputedTraitForSpace200Response,
) *NullableUpdateComputedTraitForSpace200Response {
	return &NullableUpdateComputedTraitForSpace200Response{value: val, isSet: true}
}

func (v NullableUpdateComputedTraitForSpace200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateComputedTraitForSpace200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
