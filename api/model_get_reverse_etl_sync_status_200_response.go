/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 57.4.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the GetReverseETLSyncStatus200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetReverseETLSyncStatus200Response{}

// GetReverseETLSyncStatus200Response struct for GetReverseETLSyncStatus200Response
type GetReverseETLSyncStatus200Response struct {
	Data *GetReverseETLSyncStatusOutput `json:"data,omitempty"`
}

// NewGetReverseETLSyncStatus200Response instantiates a new GetReverseETLSyncStatus200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetReverseETLSyncStatus200Response() *GetReverseETLSyncStatus200Response {
	this := GetReverseETLSyncStatus200Response{}
	return &this
}

// NewGetReverseETLSyncStatus200ResponseWithDefaults instantiates a new GetReverseETLSyncStatus200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetReverseETLSyncStatus200ResponseWithDefaults() *GetReverseETLSyncStatus200Response {
	this := GetReverseETLSyncStatus200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetReverseETLSyncStatus200Response) GetData() GetReverseETLSyncStatusOutput {
	if o == nil || IsNil(o.Data) {
		var ret GetReverseETLSyncStatusOutput
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetReverseETLSyncStatus200Response) GetDataOk() (*GetReverseETLSyncStatusOutput, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetReverseETLSyncStatus200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given GetReverseETLSyncStatusOutput and assigns it to the Data field.
func (o *GetReverseETLSyncStatus200Response) SetData(v GetReverseETLSyncStatusOutput) {
	o.Data = &v
}

func (o GetReverseETLSyncStatus200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetReverseETLSyncStatus200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGetReverseETLSyncStatus200Response struct {
	value *GetReverseETLSyncStatus200Response
	isSet bool
}

func (v NullableGetReverseETLSyncStatus200Response) Get() *GetReverseETLSyncStatus200Response {
	return v.value
}

func (v *NullableGetReverseETLSyncStatus200Response) Set(val *GetReverseETLSyncStatus200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetReverseETLSyncStatus200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetReverseETLSyncStatus200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetReverseETLSyncStatus200Response(
	val *GetReverseETLSyncStatus200Response,
) *NullableGetReverseETLSyncStatus200Response {
	return &NullableGetReverseETLSyncStatus200Response{value: val, isSet: true}
}

func (v NullableGetReverseETLSyncStatus200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetReverseETLSyncStatus200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
