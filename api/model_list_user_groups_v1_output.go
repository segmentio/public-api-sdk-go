/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 54.1.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the ListUserGroupsV1Output type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListUserGroupsV1Output{}

// ListUserGroupsV1Output Returns a list of user groups with the given Workspace id.
type ListUserGroupsV1Output struct {
	// The user group returned from the query.
	UserGroups []UserGroupV1    `json:"userGroups"`
	Pagination PaginationOutput `json:"pagination"`
}

// NewListUserGroupsV1Output instantiates a new ListUserGroupsV1Output object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListUserGroupsV1Output(
	userGroups []UserGroupV1,
	pagination PaginationOutput,
) *ListUserGroupsV1Output {
	this := ListUserGroupsV1Output{}
	this.UserGroups = userGroups
	this.Pagination = pagination
	return &this
}

// NewListUserGroupsV1OutputWithDefaults instantiates a new ListUserGroupsV1Output object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListUserGroupsV1OutputWithDefaults() *ListUserGroupsV1Output {
	this := ListUserGroupsV1Output{}
	return &this
}

// GetUserGroups returns the UserGroups field value
func (o *ListUserGroupsV1Output) GetUserGroups() []UserGroupV1 {
	if o == nil {
		var ret []UserGroupV1
		return ret
	}

	return o.UserGroups
}

// GetUserGroupsOk returns a tuple with the UserGroups field value
// and a boolean to check if the value has been set.
func (o *ListUserGroupsV1Output) GetUserGroupsOk() ([]UserGroupV1, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserGroups, true
}

// SetUserGroups sets field value
func (o *ListUserGroupsV1Output) SetUserGroups(v []UserGroupV1) {
	o.UserGroups = v
}

// GetPagination returns the Pagination field value
func (o *ListUserGroupsV1Output) GetPagination() PaginationOutput {
	if o == nil {
		var ret PaginationOutput
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *ListUserGroupsV1Output) GetPaginationOk() (*PaginationOutput, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *ListUserGroupsV1Output) SetPagination(v PaginationOutput) {
	o.Pagination = v
}

func (o ListUserGroupsV1Output) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListUserGroupsV1Output) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["userGroups"] = o.UserGroups
	toSerialize["pagination"] = o.Pagination
	return toSerialize, nil
}

type NullableListUserGroupsV1Output struct {
	value *ListUserGroupsV1Output
	isSet bool
}

func (v NullableListUserGroupsV1Output) Get() *ListUserGroupsV1Output {
	return v.value
}

func (v *NullableListUserGroupsV1Output) Set(val *ListUserGroupsV1Output) {
	v.value = val
	v.isSet = true
}

func (v NullableListUserGroupsV1Output) IsSet() bool {
	return v.isSet
}

func (v *NullableListUserGroupsV1Output) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListUserGroupsV1Output(
	val *ListUserGroupsV1Output,
) *NullableListUserGroupsV1Output {
	return &NullableListUserGroupsV1Output{value: val, isSet: true}
}

func (v NullableListUserGroupsV1Output) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListUserGroupsV1Output) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
