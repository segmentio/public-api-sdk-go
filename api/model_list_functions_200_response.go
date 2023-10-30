/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 36.3.1
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// ListFunctions200Response struct for ListFunctions200Response
type ListFunctions200Response struct {
	Data *ListFunctionsV1Output `json:"data,omitempty"`
}

// NewListFunctions200Response instantiates a new ListFunctions200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListFunctions200Response() *ListFunctions200Response {
	this := ListFunctions200Response{}
	return &this
}

// NewListFunctions200ResponseWithDefaults instantiates a new ListFunctions200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListFunctions200ResponseWithDefaults() *ListFunctions200Response {
	this := ListFunctions200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ListFunctions200Response) GetData() ListFunctionsV1Output {
	if o == nil || o.Data == nil {
		var ret ListFunctionsV1Output
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListFunctions200Response) GetDataOk() (*ListFunctionsV1Output, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ListFunctions200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given ListFunctionsV1Output and assigns it to the Data field.
func (o *ListFunctions200Response) SetData(v ListFunctionsV1Output) {
	o.Data = &v
}

func (o ListFunctions200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableListFunctions200Response struct {
	value *ListFunctions200Response
	isSet bool
}

func (v NullableListFunctions200Response) Get() *ListFunctions200Response {
	return v.value
}

func (v *NullableListFunctions200Response) Set(val *ListFunctions200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableListFunctions200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableListFunctions200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListFunctions200Response(
	val *ListFunctions200Response,
) *NullableListFunctions200Response {
	return &NullableListFunctions200Response{value: val, isSet: true}
}

func (v NullableListFunctions200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListFunctions200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
