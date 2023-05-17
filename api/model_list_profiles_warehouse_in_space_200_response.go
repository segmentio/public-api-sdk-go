/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 34.3.6
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// ListProfilesWarehouseInSpace200Response struct for ListProfilesWarehouseInSpace200Response
type ListProfilesWarehouseInSpace200Response struct {
	Data *ListProfilesWarehouseInSpaceAlphaOutput `json:"data,omitempty"`
}

// NewListProfilesWarehouseInSpace200Response instantiates a new ListProfilesWarehouseInSpace200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListProfilesWarehouseInSpace200Response() *ListProfilesWarehouseInSpace200Response {
	this := ListProfilesWarehouseInSpace200Response{}
	return &this
}

// NewListProfilesWarehouseInSpace200ResponseWithDefaults instantiates a new ListProfilesWarehouseInSpace200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListProfilesWarehouseInSpace200ResponseWithDefaults() *ListProfilesWarehouseInSpace200Response {
	this := ListProfilesWarehouseInSpace200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ListProfilesWarehouseInSpace200Response) GetData() ListProfilesWarehouseInSpaceAlphaOutput {
	if o == nil || o.Data == nil {
		var ret ListProfilesWarehouseInSpaceAlphaOutput
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListProfilesWarehouseInSpace200Response) GetDataOk() (*ListProfilesWarehouseInSpaceAlphaOutput, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ListProfilesWarehouseInSpace200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given ListProfilesWarehouseInSpaceAlphaOutput and assigns it to the Data field.
func (o *ListProfilesWarehouseInSpace200Response) SetData(
	v ListProfilesWarehouseInSpaceAlphaOutput,
) {
	o.Data = &v
}

func (o ListProfilesWarehouseInSpace200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableListProfilesWarehouseInSpace200Response struct {
	value *ListProfilesWarehouseInSpace200Response
	isSet bool
}

func (v NullableListProfilesWarehouseInSpace200Response) Get() *ListProfilesWarehouseInSpace200Response {
	return v.value
}

func (v *NullableListProfilesWarehouseInSpace200Response) Set(
	val *ListProfilesWarehouseInSpace200Response,
) {
	v.value = val
	v.isSet = true
}

func (v NullableListProfilesWarehouseInSpace200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableListProfilesWarehouseInSpace200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListProfilesWarehouseInSpace200Response(
	val *ListProfilesWarehouseInSpace200Response,
) *NullableListProfilesWarehouseInSpace200Response {
	return &NullableListProfilesWarehouseInSpace200Response{value: val, isSet: true}
}

func (v NullableListProfilesWarehouseInSpace200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListProfilesWarehouseInSpace200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
