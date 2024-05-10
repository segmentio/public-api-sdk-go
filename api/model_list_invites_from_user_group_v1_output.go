/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 50.0.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the ListInvitesFromUserGroupV1Output type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListInvitesFromUserGroupV1Output{}

// ListInvitesFromUserGroupV1Output Returns the emails of invitees to a user group with the given group id.
type ListInvitesFromUserGroupV1Output struct {
	// The emails of the invitees to the user group.
	Emails     []string         `json:"emails"`
	Pagination PaginationOutput `json:"pagination"`
}

// NewListInvitesFromUserGroupV1Output instantiates a new ListInvitesFromUserGroupV1Output object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListInvitesFromUserGroupV1Output(
	emails []string,
	pagination PaginationOutput,
) *ListInvitesFromUserGroupV1Output {
	this := ListInvitesFromUserGroupV1Output{}
	this.Emails = emails
	this.Pagination = pagination
	return &this
}

// NewListInvitesFromUserGroupV1OutputWithDefaults instantiates a new ListInvitesFromUserGroupV1Output object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListInvitesFromUserGroupV1OutputWithDefaults() *ListInvitesFromUserGroupV1Output {
	this := ListInvitesFromUserGroupV1Output{}
	return &this
}

// GetEmails returns the Emails field value
func (o *ListInvitesFromUserGroupV1Output) GetEmails() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Emails
}

// GetEmailsOk returns a tuple with the Emails field value
// and a boolean to check if the value has been set.
func (o *ListInvitesFromUserGroupV1Output) GetEmailsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Emails, true
}

// SetEmails sets field value
func (o *ListInvitesFromUserGroupV1Output) SetEmails(v []string) {
	o.Emails = v
}

// GetPagination returns the Pagination field value
func (o *ListInvitesFromUserGroupV1Output) GetPagination() PaginationOutput {
	if o == nil {
		var ret PaginationOutput
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *ListInvitesFromUserGroupV1Output) GetPaginationOk() (*PaginationOutput, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *ListInvitesFromUserGroupV1Output) SetPagination(v PaginationOutput) {
	o.Pagination = v
}

func (o ListInvitesFromUserGroupV1Output) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListInvitesFromUserGroupV1Output) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["emails"] = o.Emails
	toSerialize["pagination"] = o.Pagination
	return toSerialize, nil
}

type NullableListInvitesFromUserGroupV1Output struct {
	value *ListInvitesFromUserGroupV1Output
	isSet bool
}

func (v NullableListInvitesFromUserGroupV1Output) Get() *ListInvitesFromUserGroupV1Output {
	return v.value
}

func (v *NullableListInvitesFromUserGroupV1Output) Set(val *ListInvitesFromUserGroupV1Output) {
	v.value = val
	v.isSet = true
}

func (v NullableListInvitesFromUserGroupV1Output) IsSet() bool {
	return v.isSet
}

func (v *NullableListInvitesFromUserGroupV1Output) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListInvitesFromUserGroupV1Output(
	val *ListInvitesFromUserGroupV1Output,
) *NullableListInvitesFromUserGroupV1Output {
	return &NullableListInvitesFromUserGroupV1Output{value: val, isSet: true}
}

func (v NullableListInvitesFromUserGroupV1Output) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListInvitesFromUserGroupV1Output) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
