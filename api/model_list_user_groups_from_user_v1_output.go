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

// ListUserGroupsFromUserV1Output Returns all user groups the user belongs to.
type ListUserGroupsFromUserV1Output struct {
	// The user groups that the user belongs to.
	Groups     []MinimalUserGroupV1 `json:"groups"`
	Pagination Pagination           `json:"pagination"`
}

// NewListUserGroupsFromUserV1Output instantiates a new ListUserGroupsFromUserV1Output object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListUserGroupsFromUserV1Output(
	groups []MinimalUserGroupV1,
	pagination Pagination,
) *ListUserGroupsFromUserV1Output {
	this := ListUserGroupsFromUserV1Output{}
	this.Groups = groups
	this.Pagination = pagination
	return &this
}

// NewListUserGroupsFromUserV1OutputWithDefaults instantiates a new ListUserGroupsFromUserV1Output object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListUserGroupsFromUserV1OutputWithDefaults() *ListUserGroupsFromUserV1Output {
	this := ListUserGroupsFromUserV1Output{}
	return &this
}

// GetGroups returns the Groups field value
func (o *ListUserGroupsFromUserV1Output) GetGroups() []MinimalUserGroupV1 {
	if o == nil {
		var ret []MinimalUserGroupV1
		return ret
	}

	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value
// and a boolean to check if the value has been set.
func (o *ListUserGroupsFromUserV1Output) GetGroupsOk() ([]MinimalUserGroupV1, bool) {
	if o == nil {
		return nil, false
	}
	return o.Groups, true
}

// SetGroups sets field value
func (o *ListUserGroupsFromUserV1Output) SetGroups(v []MinimalUserGroupV1) {
	o.Groups = v
}

// GetPagination returns the Pagination field value
func (o *ListUserGroupsFromUserV1Output) GetPagination() Pagination {
	if o == nil {
		var ret Pagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *ListUserGroupsFromUserV1Output) GetPaginationOk() (*Pagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *ListUserGroupsFromUserV1Output) SetPagination(v Pagination) {
	o.Pagination = v
}

func (o ListUserGroupsFromUserV1Output) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["groups"] = o.Groups
	}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	return json.Marshal(toSerialize)
}

type NullableListUserGroupsFromUserV1Output struct {
	value *ListUserGroupsFromUserV1Output
	isSet bool
}

func (v NullableListUserGroupsFromUserV1Output) Get() *ListUserGroupsFromUserV1Output {
	return v.value
}

func (v *NullableListUserGroupsFromUserV1Output) Set(val *ListUserGroupsFromUserV1Output) {
	v.value = val
	v.isSet = true
}

func (v NullableListUserGroupsFromUserV1Output) IsSet() bool {
	return v.isSet
}

func (v *NullableListUserGroupsFromUserV1Output) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListUserGroupsFromUserV1Output(
	val *ListUserGroupsFromUserV1Output,
) *NullableListUserGroupsFromUserV1Output {
	return &NullableListUserGroupsFromUserV1Output{value: val, isSet: true}
}

func (v NullableListUserGroupsFromUserV1Output) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListUserGroupsFromUserV1Output) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
