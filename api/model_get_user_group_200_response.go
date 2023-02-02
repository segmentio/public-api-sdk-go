/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 34.0.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// GetUserGroup200Response struct for GetUserGroup200Response
type GetUserGroup200Response struct {
	Data *GetUserGroupV1Output `json:"data,omitempty"`
}

// NewGetUserGroup200Response instantiates a new GetUserGroup200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetUserGroup200Response() *GetUserGroup200Response {
	this := GetUserGroup200Response{}
	return &this
}

// NewGetUserGroup200ResponseWithDefaults instantiates a new GetUserGroup200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetUserGroup200ResponseWithDefaults() *GetUserGroup200Response {
	this := GetUserGroup200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetUserGroup200Response) GetData() GetUserGroupV1Output {
	if o == nil || o.Data == nil {
		var ret GetUserGroupV1Output
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUserGroup200Response) GetDataOk() (*GetUserGroupV1Output, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetUserGroup200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GetUserGroupV1Output and assigns it to the Data field.
func (o *GetUserGroup200Response) SetData(v GetUserGroupV1Output) {
	o.Data = &v
}

func (o GetUserGroup200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGetUserGroup200Response struct {
	value *GetUserGroup200Response
	isSet bool
}

func (v NullableGetUserGroup200Response) Get() *GetUserGroup200Response {
	return v.value
}

func (v *NullableGetUserGroup200Response) Set(val *GetUserGroup200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetUserGroup200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetUserGroup200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetUserGroup200Response(
	val *GetUserGroup200Response,
) *NullableGetUserGroup200Response {
	return &NullableGetUserGroup200Response{value: val, isSet: true}
}

func (v NullableGetUserGroup200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetUserGroup200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
