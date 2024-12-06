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

// checks if the RemoveComputedTraitFromSpace200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RemoveComputedTraitFromSpace200Response{}

// RemoveComputedTraitFromSpace200Response struct for RemoveComputedTraitFromSpace200Response
type RemoveComputedTraitFromSpace200Response struct {
	Data *RemoveComputedTraitFromSpaceAlphaOutput `json:"data,omitempty"`
}

// NewRemoveComputedTraitFromSpace200Response instantiates a new RemoveComputedTraitFromSpace200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemoveComputedTraitFromSpace200Response() *RemoveComputedTraitFromSpace200Response {
	this := RemoveComputedTraitFromSpace200Response{}
	return &this
}

// NewRemoveComputedTraitFromSpace200ResponseWithDefaults instantiates a new RemoveComputedTraitFromSpace200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemoveComputedTraitFromSpace200ResponseWithDefaults() *RemoveComputedTraitFromSpace200Response {
	this := RemoveComputedTraitFromSpace200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *RemoveComputedTraitFromSpace200Response) GetData() RemoveComputedTraitFromSpaceAlphaOutput {
	if o == nil || IsNil(o.Data) {
		var ret RemoveComputedTraitFromSpaceAlphaOutput
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoveComputedTraitFromSpace200Response) GetDataOk() (*RemoveComputedTraitFromSpaceAlphaOutput, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *RemoveComputedTraitFromSpace200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given RemoveComputedTraitFromSpaceAlphaOutput and assigns it to the Data field.
func (o *RemoveComputedTraitFromSpace200Response) SetData(
	v RemoveComputedTraitFromSpaceAlphaOutput,
) {
	o.Data = &v
}

func (o RemoveComputedTraitFromSpace200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RemoveComputedTraitFromSpace200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableRemoveComputedTraitFromSpace200Response struct {
	value *RemoveComputedTraitFromSpace200Response
	isSet bool
}

func (v NullableRemoveComputedTraitFromSpace200Response) Get() *RemoveComputedTraitFromSpace200Response {
	return v.value
}

func (v *NullableRemoveComputedTraitFromSpace200Response) Set(
	val *RemoveComputedTraitFromSpace200Response,
) {
	v.value = val
	v.isSet = true
}

func (v NullableRemoveComputedTraitFromSpace200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableRemoveComputedTraitFromSpace200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemoveComputedTraitFromSpace200Response(
	val *RemoveComputedTraitFromSpace200Response,
) *NullableRemoveComputedTraitFromSpace200Response {
	return &NullableRemoveComputedTraitFromSpace200Response{value: val, isSet: true}
}

func (v NullableRemoveComputedTraitFromSpace200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemoveComputedTraitFromSpace200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
