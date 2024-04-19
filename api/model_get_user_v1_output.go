/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 49.0.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the GetUserV1Output type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetUserV1Output{}

// GetUserV1Output Returns the user.
type GetUserV1Output struct {
	User UserV1 `json:"user"`
}

// NewGetUserV1Output instantiates a new GetUserV1Output object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetUserV1Output(user UserV1) *GetUserV1Output {
	this := GetUserV1Output{}
	this.User = user
	return &this
}

// NewGetUserV1OutputWithDefaults instantiates a new GetUserV1Output object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetUserV1OutputWithDefaults() *GetUserV1Output {
	this := GetUserV1Output{}
	return &this
}

// GetUser returns the User field value
func (o *GetUserV1Output) GetUser() UserV1 {
	if o == nil {
		var ret UserV1
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *GetUserV1Output) GetUserOk() (*UserV1, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *GetUserV1Output) SetUser(v UserV1) {
	o.User = v
}

func (o GetUserV1Output) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetUserV1Output) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["user"] = o.User
	return toSerialize, nil
}

type NullableGetUserV1Output struct {
	value *GetUserV1Output
	isSet bool
}

func (v NullableGetUserV1Output) Get() *GetUserV1Output {
	return v.value
}

func (v *NullableGetUserV1Output) Set(val *GetUserV1Output) {
	v.value = val
	v.isSet = true
}

func (v NullableGetUserV1Output) IsSet() bool {
	return v.isSet
}

func (v *NullableGetUserV1Output) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetUserV1Output(val *GetUserV1Output) *NullableGetUserV1Output {
	return &NullableGetUserV1Output{value: val, isSet: true}
}

func (v NullableGetUserV1Output) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetUserV1Output) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
