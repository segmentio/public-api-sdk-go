/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 34.1.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// ListInvites200Response struct for ListInvites200Response
type ListInvites200Response struct {
	Data *ListInvitesV1Output `json:"data,omitempty"`
}

// NewListInvites200Response instantiates a new ListInvites200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListInvites200Response() *ListInvites200Response {
	this := ListInvites200Response{}
	return &this
}

// NewListInvites200ResponseWithDefaults instantiates a new ListInvites200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListInvites200ResponseWithDefaults() *ListInvites200Response {
	this := ListInvites200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ListInvites200Response) GetData() ListInvitesV1Output {
	if o == nil || o.Data == nil {
		var ret ListInvitesV1Output
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListInvites200Response) GetDataOk() (*ListInvitesV1Output, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ListInvites200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given ListInvitesV1Output and assigns it to the Data field.
func (o *ListInvites200Response) SetData(v ListInvitesV1Output) {
	o.Data = &v
}

func (o ListInvites200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableListInvites200Response struct {
	value *ListInvites200Response
	isSet bool
}

func (v NullableListInvites200Response) Get() *ListInvites200Response {
	return v.value
}

func (v *NullableListInvites200Response) Set(val *ListInvites200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableListInvites200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableListInvites200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListInvites200Response(
	val *ListInvites200Response,
) *NullableListInvites200Response {
	return &NullableListInvites200Response{value: val, isSet: true}
}

func (v NullableListInvites200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListInvites200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
