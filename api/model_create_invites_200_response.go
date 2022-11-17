/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API. 

API version: 32.0.4
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// CreateInvites200Response struct for CreateInvites200Response
type CreateInvites200Response struct {
	Data *CreateInvitesV1Output `json:"data,omitempty"`
}

// NewCreateInvites200Response instantiates a new CreateInvites200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateInvites200Response() *CreateInvites200Response {
	this := CreateInvites200Response{}
	return &this
}

// NewCreateInvites200ResponseWithDefaults instantiates a new CreateInvites200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateInvites200ResponseWithDefaults() *CreateInvites200Response {
	this := CreateInvites200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CreateInvites200Response) GetData() CreateInvitesV1Output {
	if o == nil || o.Data == nil {
		var ret CreateInvitesV1Output
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInvites200Response) GetDataOk() (*CreateInvitesV1Output, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CreateInvites200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given CreateInvitesV1Output and assigns it to the Data field.
func (o *CreateInvites200Response) SetData(v CreateInvitesV1Output) {
	o.Data = &v
}

func (o CreateInvites200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableCreateInvites200Response struct {
	value *CreateInvites200Response
	isSet bool
}

func (v NullableCreateInvites200Response) Get() *CreateInvites200Response {
	return v.value
}

func (v *NullableCreateInvites200Response) Set(val *CreateInvites200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateInvites200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateInvites200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateInvites200Response(val *CreateInvites200Response) *NullableCreateInvites200Response {
	return &NullableCreateInvites200Response{value: val, isSet: true}
}

func (v NullableCreateInvites200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateInvites200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


