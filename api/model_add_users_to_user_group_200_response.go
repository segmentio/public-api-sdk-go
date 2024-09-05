/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 53.3.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the AddUsersToUserGroup200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddUsersToUserGroup200Response{}

// AddUsersToUserGroup200Response struct for AddUsersToUserGroup200Response
type AddUsersToUserGroup200Response struct {
	Data *AddUsersToUserGroupV1Output `json:"data,omitempty"`
}

// NewAddUsersToUserGroup200Response instantiates a new AddUsersToUserGroup200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddUsersToUserGroup200Response() *AddUsersToUserGroup200Response {
	this := AddUsersToUserGroup200Response{}
	return &this
}

// NewAddUsersToUserGroup200ResponseWithDefaults instantiates a new AddUsersToUserGroup200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddUsersToUserGroup200ResponseWithDefaults() *AddUsersToUserGroup200Response {
	this := AddUsersToUserGroup200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AddUsersToUserGroup200Response) GetData() AddUsersToUserGroupV1Output {
	if o == nil || IsNil(o.Data) {
		var ret AddUsersToUserGroupV1Output
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddUsersToUserGroup200Response) GetDataOk() (*AddUsersToUserGroupV1Output, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AddUsersToUserGroup200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given AddUsersToUserGroupV1Output and assigns it to the Data field.
func (o *AddUsersToUserGroup200Response) SetData(v AddUsersToUserGroupV1Output) {
	o.Data = &v
}

func (o AddUsersToUserGroup200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddUsersToUserGroup200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableAddUsersToUserGroup200Response struct {
	value *AddUsersToUserGroup200Response
	isSet bool
}

func (v NullableAddUsersToUserGroup200Response) Get() *AddUsersToUserGroup200Response {
	return v.value
}

func (v *NullableAddUsersToUserGroup200Response) Set(val *AddUsersToUserGroup200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableAddUsersToUserGroup200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableAddUsersToUserGroup200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddUsersToUserGroup200Response(
	val *AddUsersToUserGroup200Response,
) *NullableAddUsersToUserGroup200Response {
	return &NullableAddUsersToUserGroup200Response{value: val, isSet: true}
}

func (v NullableAddUsersToUserGroup200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddUsersToUserGroup200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
