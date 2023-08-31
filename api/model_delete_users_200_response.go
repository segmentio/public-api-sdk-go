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

// DeleteUsers200Response struct for DeleteUsers200Response
type DeleteUsers200Response struct {
	Data *DeleteUsersV1Output `json:"data,omitempty"`
}

// NewDeleteUsers200Response instantiates a new DeleteUsers200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteUsers200Response() *DeleteUsers200Response {
	this := DeleteUsers200Response{}
	return &this
}

// NewDeleteUsers200ResponseWithDefaults instantiates a new DeleteUsers200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteUsers200ResponseWithDefaults() *DeleteUsers200Response {
	this := DeleteUsers200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *DeleteUsers200Response) GetData() DeleteUsersV1Output {
	if o == nil || o.Data == nil {
		var ret DeleteUsersV1Output
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteUsers200Response) GetDataOk() (*DeleteUsersV1Output, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *DeleteUsers200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given DeleteUsersV1Output and assigns it to the Data field.
func (o *DeleteUsers200Response) SetData(v DeleteUsersV1Output) {
	o.Data = &v
}

func (o DeleteUsers200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableDeleteUsers200Response struct {
	value *DeleteUsers200Response
	isSet bool
}

func (v NullableDeleteUsers200Response) Get() *DeleteUsers200Response {
	return v.value
}

func (v *NullableDeleteUsers200Response) Set(val *DeleteUsers200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteUsers200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteUsers200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteUsers200Response(
	val *DeleteUsers200Response,
) *NullableDeleteUsers200Response {
	return &NullableDeleteUsers200Response{value: val, isSet: true}
}

func (v NullableDeleteUsers200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteUsers200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
