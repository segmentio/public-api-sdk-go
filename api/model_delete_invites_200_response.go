/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 58.13.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the DeleteInvites200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteInvites200Response{}

// DeleteInvites200Response struct for DeleteInvites200Response
type DeleteInvites200Response struct {
	Data *DeleteInvitesV1Output `json:"data,omitempty"`
}

// NewDeleteInvites200Response instantiates a new DeleteInvites200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteInvites200Response() *DeleteInvites200Response {
	this := DeleteInvites200Response{}
	return &this
}

// NewDeleteInvites200ResponseWithDefaults instantiates a new DeleteInvites200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteInvites200ResponseWithDefaults() *DeleteInvites200Response {
	this := DeleteInvites200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *DeleteInvites200Response) GetData() DeleteInvitesV1Output {
	if o == nil || IsNil(o.Data) {
		var ret DeleteInvitesV1Output
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteInvites200Response) GetDataOk() (*DeleteInvitesV1Output, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *DeleteInvites200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given DeleteInvitesV1Output and assigns it to the Data field.
func (o *DeleteInvites200Response) SetData(v DeleteInvitesV1Output) {
	o.Data = &v
}

func (o DeleteInvites200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteInvites200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableDeleteInvites200Response struct {
	value *DeleteInvites200Response
	isSet bool
}

func (v NullableDeleteInvites200Response) Get() *DeleteInvites200Response {
	return v.value
}

func (v *NullableDeleteInvites200Response) Set(val *DeleteInvites200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteInvites200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteInvites200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteInvites200Response(
	val *DeleteInvites200Response,
) *NullableDeleteInvites200Response {
	return &NullableDeleteInvites200Response{value: val, isSet: true}
}

func (v NullableDeleteInvites200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteInvites200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
