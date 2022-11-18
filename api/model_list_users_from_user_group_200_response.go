/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API. 

API version: 33.0.1
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// ListUsersFromUserGroup200Response struct for ListUsersFromUserGroup200Response
type ListUsersFromUserGroup200Response struct {
	Data *ListUsersFromUserGroupV1Output `json:"data,omitempty"`
}

// NewListUsersFromUserGroup200Response instantiates a new ListUsersFromUserGroup200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListUsersFromUserGroup200Response() *ListUsersFromUserGroup200Response {
	this := ListUsersFromUserGroup200Response{}
	return &this
}

// NewListUsersFromUserGroup200ResponseWithDefaults instantiates a new ListUsersFromUserGroup200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListUsersFromUserGroup200ResponseWithDefaults() *ListUsersFromUserGroup200Response {
	this := ListUsersFromUserGroup200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ListUsersFromUserGroup200Response) GetData() ListUsersFromUserGroupV1Output {
	if o == nil || o.Data == nil {
		var ret ListUsersFromUserGroupV1Output
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListUsersFromUserGroup200Response) GetDataOk() (*ListUsersFromUserGroupV1Output, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ListUsersFromUserGroup200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given ListUsersFromUserGroupV1Output and assigns it to the Data field.
func (o *ListUsersFromUserGroup200Response) SetData(v ListUsersFromUserGroupV1Output) {
	o.Data = &v
}

func (o ListUsersFromUserGroup200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableListUsersFromUserGroup200Response struct {
	value *ListUsersFromUserGroup200Response
	isSet bool
}

func (v NullableListUsersFromUserGroup200Response) Get() *ListUsersFromUserGroup200Response {
	return v.value
}

func (v *NullableListUsersFromUserGroup200Response) Set(val *ListUsersFromUserGroup200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableListUsersFromUserGroup200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableListUsersFromUserGroup200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListUsersFromUserGroup200Response(val *ListUsersFromUserGroup200Response) *NullableListUsersFromUserGroup200Response {
	return &NullableListUsersFromUserGroup200Response{value: val, isSet: true}
}

func (v NullableListUsersFromUserGroup200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListUsersFromUserGroup200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


