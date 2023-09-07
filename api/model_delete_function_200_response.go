/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 36.2.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// DeleteFunction200Response struct for DeleteFunction200Response
type DeleteFunction200Response struct {
	Data *DeleteFunctionV1Output `json:"data,omitempty"`
}

// NewDeleteFunction200Response instantiates a new DeleteFunction200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteFunction200Response() *DeleteFunction200Response {
	this := DeleteFunction200Response{}
	return &this
}

// NewDeleteFunction200ResponseWithDefaults instantiates a new DeleteFunction200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteFunction200ResponseWithDefaults() *DeleteFunction200Response {
	this := DeleteFunction200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *DeleteFunction200Response) GetData() DeleteFunctionV1Output {
	if o == nil || o.Data == nil {
		var ret DeleteFunctionV1Output
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteFunction200Response) GetDataOk() (*DeleteFunctionV1Output, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *DeleteFunction200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given DeleteFunctionV1Output and assigns it to the Data field.
func (o *DeleteFunction200Response) SetData(v DeleteFunctionV1Output) {
	o.Data = &v
}

func (o DeleteFunction200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableDeleteFunction200Response struct {
	value *DeleteFunction200Response
	isSet bool
}

func (v NullableDeleteFunction200Response) Get() *DeleteFunction200Response {
	return v.value
}

func (v *NullableDeleteFunction200Response) Set(val *DeleteFunction200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteFunction200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteFunction200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteFunction200Response(
	val *DeleteFunction200Response,
) *NullableDeleteFunction200Response {
	return &NullableDeleteFunction200Response{value: val, isSet: true}
}

func (v NullableDeleteFunction200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteFunction200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
