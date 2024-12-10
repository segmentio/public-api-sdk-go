/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 57.2.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the RemoveProfilesWarehouseFromSpace200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RemoveProfilesWarehouseFromSpace200Response{}

// RemoveProfilesWarehouseFromSpace200Response struct for RemoveProfilesWarehouseFromSpace200Response
type RemoveProfilesWarehouseFromSpace200Response struct {
	Data *RemoveProfilesWarehouseFromSpaceAlphaOutput `json:"data,omitempty"`
}

// NewRemoveProfilesWarehouseFromSpace200Response instantiates a new RemoveProfilesWarehouseFromSpace200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemoveProfilesWarehouseFromSpace200Response() *RemoveProfilesWarehouseFromSpace200Response {
	this := RemoveProfilesWarehouseFromSpace200Response{}
	return &this
}

// NewRemoveProfilesWarehouseFromSpace200ResponseWithDefaults instantiates a new RemoveProfilesWarehouseFromSpace200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemoveProfilesWarehouseFromSpace200ResponseWithDefaults() *RemoveProfilesWarehouseFromSpace200Response {
	this := RemoveProfilesWarehouseFromSpace200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *RemoveProfilesWarehouseFromSpace200Response) GetData() RemoveProfilesWarehouseFromSpaceAlphaOutput {
	if o == nil || IsNil(o.Data) {
		var ret RemoveProfilesWarehouseFromSpaceAlphaOutput
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoveProfilesWarehouseFromSpace200Response) GetDataOk() (*RemoveProfilesWarehouseFromSpaceAlphaOutput, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *RemoveProfilesWarehouseFromSpace200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given RemoveProfilesWarehouseFromSpaceAlphaOutput and assigns it to the Data field.
func (o *RemoveProfilesWarehouseFromSpace200Response) SetData(
	v RemoveProfilesWarehouseFromSpaceAlphaOutput,
) {
	o.Data = &v
}

func (o RemoveProfilesWarehouseFromSpace200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RemoveProfilesWarehouseFromSpace200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableRemoveProfilesWarehouseFromSpace200Response struct {
	value *RemoveProfilesWarehouseFromSpace200Response
	isSet bool
}

func (v NullableRemoveProfilesWarehouseFromSpace200Response) Get() *RemoveProfilesWarehouseFromSpace200Response {
	return v.value
}

func (v *NullableRemoveProfilesWarehouseFromSpace200Response) Set(
	val *RemoveProfilesWarehouseFromSpace200Response,
) {
	v.value = val
	v.isSet = true
}

func (v NullableRemoveProfilesWarehouseFromSpace200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableRemoveProfilesWarehouseFromSpace200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemoveProfilesWarehouseFromSpace200Response(
	val *RemoveProfilesWarehouseFromSpace200Response,
) *NullableRemoveProfilesWarehouseFromSpace200Response {
	return &NullableRemoveProfilesWarehouseFromSpace200Response{value: val, isSet: true}
}

func (v NullableRemoveProfilesWarehouseFromSpace200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemoveProfilesWarehouseFromSpace200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
