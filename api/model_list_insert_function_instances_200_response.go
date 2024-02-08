/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 42.0.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the ListInsertFunctionInstances200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListInsertFunctionInstances200Response{}

// ListInsertFunctionInstances200Response struct for ListInsertFunctionInstances200Response
type ListInsertFunctionInstances200Response struct {
	Data *ListInsertFunctionInstancesAlphaOutput `json:"data,omitempty"`
}

// NewListInsertFunctionInstances200Response instantiates a new ListInsertFunctionInstances200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListInsertFunctionInstances200Response() *ListInsertFunctionInstances200Response {
	this := ListInsertFunctionInstances200Response{}
	return &this
}

// NewListInsertFunctionInstances200ResponseWithDefaults instantiates a new ListInsertFunctionInstances200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListInsertFunctionInstances200ResponseWithDefaults() *ListInsertFunctionInstances200Response {
	this := ListInsertFunctionInstances200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ListInsertFunctionInstances200Response) GetData() ListInsertFunctionInstancesAlphaOutput {
	if o == nil || IsNil(o.Data) {
		var ret ListInsertFunctionInstancesAlphaOutput
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListInsertFunctionInstances200Response) GetDataOk() (*ListInsertFunctionInstancesAlphaOutput, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ListInsertFunctionInstances200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given ListInsertFunctionInstancesAlphaOutput and assigns it to the Data field.
func (o *ListInsertFunctionInstances200Response) SetData(v ListInsertFunctionInstancesAlphaOutput) {
	o.Data = &v
}

func (o ListInsertFunctionInstances200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListInsertFunctionInstances200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableListInsertFunctionInstances200Response struct {
	value *ListInsertFunctionInstances200Response
	isSet bool
}

func (v NullableListInsertFunctionInstances200Response) Get() *ListInsertFunctionInstances200Response {
	return v.value
}

func (v *NullableListInsertFunctionInstances200Response) Set(
	val *ListInsertFunctionInstances200Response,
) {
	v.value = val
	v.isSet = true
}

func (v NullableListInsertFunctionInstances200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableListInsertFunctionInstances200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListInsertFunctionInstances200Response(
	val *ListInsertFunctionInstances200Response,
) *NullableListInsertFunctionInstances200Response {
	return &NullableListInsertFunctionInstances200Response{value: val, isSet: true}
}

func (v NullableListInsertFunctionInstances200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListInsertFunctionInstances200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
