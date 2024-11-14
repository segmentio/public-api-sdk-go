/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 57.0.1
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the ListUsersFromUserGroupV1Output type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListUsersFromUserGroupV1Output{}

// ListUsersFromUserGroupV1Output Returns the members of a user group with the given group id.
type ListUsersFromUserGroupV1Output struct {
	// The users of the user group.
	Users      []MinimalUserV1  `json:"users"`
	Pagination PaginationOutput `json:"pagination"`
}

// NewListUsersFromUserGroupV1Output instantiates a new ListUsersFromUserGroupV1Output object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListUsersFromUserGroupV1Output(
	users []MinimalUserV1,
	pagination PaginationOutput,
) *ListUsersFromUserGroupV1Output {
	this := ListUsersFromUserGroupV1Output{}
	this.Users = users
	this.Pagination = pagination
	return &this
}

// NewListUsersFromUserGroupV1OutputWithDefaults instantiates a new ListUsersFromUserGroupV1Output object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListUsersFromUserGroupV1OutputWithDefaults() *ListUsersFromUserGroupV1Output {
	this := ListUsersFromUserGroupV1Output{}
	return &this
}

// GetUsers returns the Users field value
func (o *ListUsersFromUserGroupV1Output) GetUsers() []MinimalUserV1 {
	if o == nil {
		var ret []MinimalUserV1
		return ret
	}

	return o.Users
}

// GetUsersOk returns a tuple with the Users field value
// and a boolean to check if the value has been set.
func (o *ListUsersFromUserGroupV1Output) GetUsersOk() ([]MinimalUserV1, bool) {
	if o == nil {
		return nil, false
	}
	return o.Users, true
}

// SetUsers sets field value
func (o *ListUsersFromUserGroupV1Output) SetUsers(v []MinimalUserV1) {
	o.Users = v
}

// GetPagination returns the Pagination field value
func (o *ListUsersFromUserGroupV1Output) GetPagination() PaginationOutput {
	if o == nil {
		var ret PaginationOutput
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *ListUsersFromUserGroupV1Output) GetPaginationOk() (*PaginationOutput, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *ListUsersFromUserGroupV1Output) SetPagination(v PaginationOutput) {
	o.Pagination = v
}

func (o ListUsersFromUserGroupV1Output) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListUsersFromUserGroupV1Output) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["users"] = o.Users
	toSerialize["pagination"] = o.Pagination
	return toSerialize, nil
}

type NullableListUsersFromUserGroupV1Output struct {
	value *ListUsersFromUserGroupV1Output
	isSet bool
}

func (v NullableListUsersFromUserGroupV1Output) Get() *ListUsersFromUserGroupV1Output {
	return v.value
}

func (v *NullableListUsersFromUserGroupV1Output) Set(val *ListUsersFromUserGroupV1Output) {
	v.value = val
	v.isSet = true
}

func (v NullableListUsersFromUserGroupV1Output) IsSet() bool {
	return v.isSet
}

func (v *NullableListUsersFromUserGroupV1Output) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListUsersFromUserGroupV1Output(
	val *ListUsersFromUserGroupV1Output,
) *NullableListUsersFromUserGroupV1Output {
	return &NullableListUsersFromUserGroupV1Output{value: val, isSet: true}
}

func (v NullableListUsersFromUserGroupV1Output) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListUsersFromUserGroupV1Output) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
