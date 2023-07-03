/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 35.0.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// ReplaceUsersInUserGroup200Response struct for ReplaceUsersInUserGroup200Response
type ReplaceUsersInUserGroup200Response struct {
	Data *ReplaceUsersInUserGroupV1Output `json:"data,omitempty"`
}

// NewReplaceUsersInUserGroup200Response instantiates a new ReplaceUsersInUserGroup200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReplaceUsersInUserGroup200Response() *ReplaceUsersInUserGroup200Response {
	this := ReplaceUsersInUserGroup200Response{}
	return &this
}

// NewReplaceUsersInUserGroup200ResponseWithDefaults instantiates a new ReplaceUsersInUserGroup200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReplaceUsersInUserGroup200ResponseWithDefaults() *ReplaceUsersInUserGroup200Response {
	this := ReplaceUsersInUserGroup200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ReplaceUsersInUserGroup200Response) GetData() ReplaceUsersInUserGroupV1Output {
	if o == nil || o.Data == nil {
		var ret ReplaceUsersInUserGroupV1Output
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplaceUsersInUserGroup200Response) GetDataOk() (*ReplaceUsersInUserGroupV1Output, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ReplaceUsersInUserGroup200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given ReplaceUsersInUserGroupV1Output and assigns it to the Data field.
func (o *ReplaceUsersInUserGroup200Response) SetData(v ReplaceUsersInUserGroupV1Output) {
	o.Data = &v
}

func (o ReplaceUsersInUserGroup200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableReplaceUsersInUserGroup200Response struct {
	value *ReplaceUsersInUserGroup200Response
	isSet bool
}

func (v NullableReplaceUsersInUserGroup200Response) Get() *ReplaceUsersInUserGroup200Response {
	return v.value
}

func (v *NullableReplaceUsersInUserGroup200Response) Set(val *ReplaceUsersInUserGroup200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableReplaceUsersInUserGroup200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableReplaceUsersInUserGroup200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReplaceUsersInUserGroup200Response(
	val *ReplaceUsersInUserGroup200Response,
) *NullableReplaceUsersInUserGroup200Response {
	return &NullableReplaceUsersInUserGroup200Response{value: val, isSet: true}
}

func (v NullableReplaceUsersInUserGroup200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReplaceUsersInUserGroup200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
