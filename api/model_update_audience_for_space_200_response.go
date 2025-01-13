/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 57.3.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the UpdateAudienceForSpace200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateAudienceForSpace200Response{}

// UpdateAudienceForSpace200Response struct for UpdateAudienceForSpace200Response
type UpdateAudienceForSpace200Response struct {
	Data *UpdateAudienceForSpaceAlphaOutput `json:"data,omitempty"`
}

// NewUpdateAudienceForSpace200Response instantiates a new UpdateAudienceForSpace200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateAudienceForSpace200Response() *UpdateAudienceForSpace200Response {
	this := UpdateAudienceForSpace200Response{}
	return &this
}

// NewUpdateAudienceForSpace200ResponseWithDefaults instantiates a new UpdateAudienceForSpace200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateAudienceForSpace200ResponseWithDefaults() *UpdateAudienceForSpace200Response {
	this := UpdateAudienceForSpace200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *UpdateAudienceForSpace200Response) GetData() UpdateAudienceForSpaceAlphaOutput {
	if o == nil || IsNil(o.Data) {
		var ret UpdateAudienceForSpaceAlphaOutput
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAudienceForSpace200Response) GetDataOk() (*UpdateAudienceForSpaceAlphaOutput, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *UpdateAudienceForSpace200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given UpdateAudienceForSpaceAlphaOutput and assigns it to the Data field.
func (o *UpdateAudienceForSpace200Response) SetData(v UpdateAudienceForSpaceAlphaOutput) {
	o.Data = &v
}

func (o UpdateAudienceForSpace200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateAudienceForSpace200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableUpdateAudienceForSpace200Response struct {
	value *UpdateAudienceForSpace200Response
	isSet bool
}

func (v NullableUpdateAudienceForSpace200Response) Get() *UpdateAudienceForSpace200Response {
	return v.value
}

func (v *NullableUpdateAudienceForSpace200Response) Set(val *UpdateAudienceForSpace200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateAudienceForSpace200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateAudienceForSpace200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateAudienceForSpace200Response(
	val *UpdateAudienceForSpace200Response,
) *NullableUpdateAudienceForSpace200Response {
	return &NullableUpdateAudienceForSpace200Response{value: val, isSet: true}
}

func (v NullableUpdateAudienceForSpace200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateAudienceForSpace200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
