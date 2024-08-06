/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 52.0.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the CreateInvites201Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateInvites201Response{}

// CreateInvites201Response struct for CreateInvites201Response
type CreateInvites201Response struct {
	Data *CreateInvitesV1Output `json:"data,omitempty"`
}

// NewCreateInvites201Response instantiates a new CreateInvites201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateInvites201Response() *CreateInvites201Response {
	this := CreateInvites201Response{}
	return &this
}

// NewCreateInvites201ResponseWithDefaults instantiates a new CreateInvites201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateInvites201ResponseWithDefaults() *CreateInvites201Response {
	this := CreateInvites201Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CreateInvites201Response) GetData() CreateInvitesV1Output {
	if o == nil || IsNil(o.Data) {
		var ret CreateInvitesV1Output
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInvites201Response) GetDataOk() (*CreateInvitesV1Output, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CreateInvites201Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given CreateInvitesV1Output and assigns it to the Data field.
func (o *CreateInvites201Response) SetData(v CreateInvitesV1Output) {
	o.Data = &v
}

func (o CreateInvites201Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateInvites201Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableCreateInvites201Response struct {
	value *CreateInvites201Response
	isSet bool
}

func (v NullableCreateInvites201Response) Get() *CreateInvites201Response {
	return v.value
}

func (v *NullableCreateInvites201Response) Set(val *CreateInvites201Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateInvites201Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateInvites201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateInvites201Response(
	val *CreateInvites201Response,
) *NullableCreateInvites201Response {
	return &NullableCreateInvites201Response{value: val, isSet: true}
}

func (v NullableCreateInvites201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateInvites201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
