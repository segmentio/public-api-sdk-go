/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 34.3.4
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// ReplacePermissionsForUserGroup200Response struct for ReplacePermissionsForUserGroup200Response
type ReplacePermissionsForUserGroup200Response struct {
	Data *ReplacePermissionsForUserGroupV1Output `json:"data,omitempty"`
}

// NewReplacePermissionsForUserGroup200Response instantiates a new ReplacePermissionsForUserGroup200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReplacePermissionsForUserGroup200Response() *ReplacePermissionsForUserGroup200Response {
	this := ReplacePermissionsForUserGroup200Response{}
	return &this
}

// NewReplacePermissionsForUserGroup200ResponseWithDefaults instantiates a new ReplacePermissionsForUserGroup200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReplacePermissionsForUserGroup200ResponseWithDefaults() *ReplacePermissionsForUserGroup200Response {
	this := ReplacePermissionsForUserGroup200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ReplacePermissionsForUserGroup200Response) GetData() ReplacePermissionsForUserGroupV1Output {
	if o == nil || o.Data == nil {
		var ret ReplacePermissionsForUserGroupV1Output
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplacePermissionsForUserGroup200Response) GetDataOk() (*ReplacePermissionsForUserGroupV1Output, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ReplacePermissionsForUserGroup200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given ReplacePermissionsForUserGroupV1Output and assigns it to the Data field.
func (o *ReplacePermissionsForUserGroup200Response) SetData(
	v ReplacePermissionsForUserGroupV1Output,
) {
	o.Data = &v
}

func (o ReplacePermissionsForUserGroup200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableReplacePermissionsForUserGroup200Response struct {
	value *ReplacePermissionsForUserGroup200Response
	isSet bool
}

func (v NullableReplacePermissionsForUserGroup200Response) Get() *ReplacePermissionsForUserGroup200Response {
	return v.value
}

func (v *NullableReplacePermissionsForUserGroup200Response) Set(
	val *ReplacePermissionsForUserGroup200Response,
) {
	v.value = val
	v.isSet = true
}

func (v NullableReplacePermissionsForUserGroup200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableReplacePermissionsForUserGroup200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReplacePermissionsForUserGroup200Response(
	val *ReplacePermissionsForUserGroup200Response,
) *NullableReplacePermissionsForUserGroup200Response {
	return &NullableReplacePermissionsForUserGroup200Response{value: val, isSet: true}
}

func (v NullableReplacePermissionsForUserGroup200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReplacePermissionsForUserGroup200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
