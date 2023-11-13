/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 37.2.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// ListRolesV1Output Returns the list of roles.
type ListRolesV1Output struct {
	// The list of roles.
	Roles      []RoleV1   `json:"roles"`
	Pagination Pagination `json:"pagination"`
}

// NewListRolesV1Output instantiates a new ListRolesV1Output object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListRolesV1Output(roles []RoleV1, pagination Pagination) *ListRolesV1Output {
	this := ListRolesV1Output{}
	this.Roles = roles
	this.Pagination = pagination
	return &this
}

// NewListRolesV1OutputWithDefaults instantiates a new ListRolesV1Output object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListRolesV1OutputWithDefaults() *ListRolesV1Output {
	this := ListRolesV1Output{}
	return &this
}

// GetRoles returns the Roles field value
func (o *ListRolesV1Output) GetRoles() []RoleV1 {
	if o == nil {
		var ret []RoleV1
		return ret
	}

	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value
// and a boolean to check if the value has been set.
func (o *ListRolesV1Output) GetRolesOk() ([]RoleV1, bool) {
	if o == nil {
		return nil, false
	}
	return o.Roles, true
}

// SetRoles sets field value
func (o *ListRolesV1Output) SetRoles(v []RoleV1) {
	o.Roles = v
}

// GetPagination returns the Pagination field value
func (o *ListRolesV1Output) GetPagination() Pagination {
	if o == nil {
		var ret Pagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *ListRolesV1Output) GetPaginationOk() (*Pagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *ListRolesV1Output) SetPagination(v Pagination) {
	o.Pagination = v
}

func (o ListRolesV1Output) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["roles"] = o.Roles
	}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	return json.Marshal(toSerialize)
}

type NullableListRolesV1Output struct {
	value *ListRolesV1Output
	isSet bool
}

func (v NullableListRolesV1Output) Get() *ListRolesV1Output {
	return v.value
}

func (v *NullableListRolesV1Output) Set(val *ListRolesV1Output) {
	v.value = val
	v.isSet = true
}

func (v NullableListRolesV1Output) IsSet() bool {
	return v.isSet
}

func (v *NullableListRolesV1Output) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListRolesV1Output(val *ListRolesV1Output) *NullableListRolesV1Output {
	return &NullableListRolesV1Output{value: val, isSet: true}
}

func (v NullableListRolesV1Output) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListRolesV1Output) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
