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

// ListReverseEtlModels200Response struct for ListReverseEtlModels200Response
type ListReverseEtlModels200Response struct {
	Data *ListReverseEtlModelsOutput `json:"data,omitempty"`
}

// NewListReverseEtlModels200Response instantiates a new ListReverseEtlModels200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListReverseEtlModels200Response() *ListReverseEtlModels200Response {
	this := ListReverseEtlModels200Response{}
	return &this
}

// NewListReverseEtlModels200ResponseWithDefaults instantiates a new ListReverseEtlModels200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListReverseEtlModels200ResponseWithDefaults() *ListReverseEtlModels200Response {
	this := ListReverseEtlModels200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ListReverseEtlModels200Response) GetData() ListReverseEtlModelsOutput {
	if o == nil || o.Data == nil {
		var ret ListReverseEtlModelsOutput
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListReverseEtlModels200Response) GetDataOk() (*ListReverseEtlModelsOutput, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ListReverseEtlModels200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given ListReverseEtlModelsOutput and assigns it to the Data field.
func (o *ListReverseEtlModels200Response) SetData(v ListReverseEtlModelsOutput) {
	o.Data = &v
}

func (o ListReverseEtlModels200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableListReverseEtlModels200Response struct {
	value *ListReverseEtlModels200Response
	isSet bool
}

func (v NullableListReverseEtlModels200Response) Get() *ListReverseEtlModels200Response {
	return v.value
}

func (v *NullableListReverseEtlModels200Response) Set(val *ListReverseEtlModels200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableListReverseEtlModels200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableListReverseEtlModels200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListReverseEtlModels200Response(
	val *ListReverseEtlModels200Response,
) *NullableListReverseEtlModels200Response {
	return &NullableListReverseEtlModels200Response{value: val, isSet: true}
}

func (v NullableListReverseEtlModels200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListReverseEtlModels200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
