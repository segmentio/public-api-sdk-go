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

// CreateLabel200Response struct for CreateLabel200Response
type CreateLabel200Response struct {
	Data *CreateLabelV1Output `json:"data,omitempty"`
}

// NewCreateLabel200Response instantiates a new CreateLabel200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateLabel200Response() *CreateLabel200Response {
	this := CreateLabel200Response{}
	return &this
}

// NewCreateLabel200ResponseWithDefaults instantiates a new CreateLabel200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateLabel200ResponseWithDefaults() *CreateLabel200Response {
	this := CreateLabel200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CreateLabel200Response) GetData() CreateLabelV1Output {
	if o == nil || o.Data == nil {
		var ret CreateLabelV1Output
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLabel200Response) GetDataOk() (*CreateLabelV1Output, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CreateLabel200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given CreateLabelV1Output and assigns it to the Data field.
func (o *CreateLabel200Response) SetData(v CreateLabelV1Output) {
	o.Data = &v
}

func (o CreateLabel200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableCreateLabel200Response struct {
	value *CreateLabel200Response
	isSet bool
}

func (v NullableCreateLabel200Response) Get() *CreateLabel200Response {
	return v.value
}

func (v *NullableCreateLabel200Response) Set(val *CreateLabel200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateLabel200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateLabel200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateLabel200Response(
	val *CreateLabel200Response,
) *NullableCreateLabel200Response {
	return &NullableCreateLabel200Response{value: val, isSet: true}
}

func (v NullableCreateLabel200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateLabel200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
