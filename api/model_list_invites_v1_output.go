/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 34.0.1
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// ListInvitesV1Output Returns the list of invites.
type ListInvitesV1Output struct {
	// The list of invites.
	Invites    []string   `json:"invites"`
	Pagination Pagination `json:"pagination"`
}

// NewListInvitesV1Output instantiates a new ListInvitesV1Output object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListInvitesV1Output(invites []string, pagination Pagination) *ListInvitesV1Output {
	this := ListInvitesV1Output{}
	this.Invites = invites
	this.Pagination = pagination
	return &this
}

// NewListInvitesV1OutputWithDefaults instantiates a new ListInvitesV1Output object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListInvitesV1OutputWithDefaults() *ListInvitesV1Output {
	this := ListInvitesV1Output{}
	return &this
}

// GetInvites returns the Invites field value
func (o *ListInvitesV1Output) GetInvites() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Invites
}

// GetInvitesOk returns a tuple with the Invites field value
// and a boolean to check if the value has been set.
func (o *ListInvitesV1Output) GetInvitesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Invites, true
}

// SetInvites sets field value
func (o *ListInvitesV1Output) SetInvites(v []string) {
	o.Invites = v
}

// GetPagination returns the Pagination field value
func (o *ListInvitesV1Output) GetPagination() Pagination {
	if o == nil {
		var ret Pagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *ListInvitesV1Output) GetPaginationOk() (*Pagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *ListInvitesV1Output) SetPagination(v Pagination) {
	o.Pagination = v
}

func (o ListInvitesV1Output) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["invites"] = o.Invites
	}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	return json.Marshal(toSerialize)
}

type NullableListInvitesV1Output struct {
	value *ListInvitesV1Output
	isSet bool
}

func (v NullableListInvitesV1Output) Get() *ListInvitesV1Output {
	return v.value
}

func (v *NullableListInvitesV1Output) Set(val *ListInvitesV1Output) {
	v.value = val
	v.isSet = true
}

func (v NullableListInvitesV1Output) IsSet() bool {
	return v.isSet
}

func (v *NullableListInvitesV1Output) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListInvitesV1Output(val *ListInvitesV1Output) *NullableListInvitesV1Output {
	return &NullableListInvitesV1Output{value: val, isSet: true}
}

func (v NullableListInvitesV1Output) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListInvitesV1Output) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
