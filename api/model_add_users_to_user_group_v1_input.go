/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 33.0.4
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// AddUsersToUserGroupV1Input Adds a list of users and invites to a user group.
type AddUsersToUserGroupV1Input struct {
	// The email addresses of the users and invites to add.
	Emails []string `json:"emails,omitempty"`
}

// NewAddUsersToUserGroupV1Input instantiates a new AddUsersToUserGroupV1Input object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddUsersToUserGroupV1Input() *AddUsersToUserGroupV1Input {
	this := AddUsersToUserGroupV1Input{}
	return &this
}

// NewAddUsersToUserGroupV1InputWithDefaults instantiates a new AddUsersToUserGroupV1Input object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddUsersToUserGroupV1InputWithDefaults() *AddUsersToUserGroupV1Input {
	this := AddUsersToUserGroupV1Input{}
	return &this
}

// GetEmails returns the Emails field value if set, zero value otherwise.
func (o *AddUsersToUserGroupV1Input) GetEmails() []string {
	if o == nil || o.Emails == nil {
		var ret []string
		return ret
	}
	return o.Emails
}

// GetEmailsOk returns a tuple with the Emails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddUsersToUserGroupV1Input) GetEmailsOk() ([]string, bool) {
	if o == nil || o.Emails == nil {
		return nil, false
	}
	return o.Emails, true
}

// HasEmails returns a boolean if a field has been set.
func (o *AddUsersToUserGroupV1Input) HasEmails() bool {
	if o != nil && o.Emails != nil {
		return true
	}

	return false
}

// SetEmails gets a reference to the given []string and assigns it to the Emails field.
func (o *AddUsersToUserGroupV1Input) SetEmails(v []string) {
	o.Emails = v
}

func (o AddUsersToUserGroupV1Input) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Emails != nil {
		toSerialize["emails"] = o.Emails
	}
	return json.Marshal(toSerialize)
}

type NullableAddUsersToUserGroupV1Input struct {
	value *AddUsersToUserGroupV1Input
	isSet bool
}

func (v NullableAddUsersToUserGroupV1Input) Get() *AddUsersToUserGroupV1Input {
	return v.value
}

func (v *NullableAddUsersToUserGroupV1Input) Set(val *AddUsersToUserGroupV1Input) {
	v.value = val
	v.isSet = true
}

func (v NullableAddUsersToUserGroupV1Input) IsSet() bool {
	return v.isSet
}

func (v *NullableAddUsersToUserGroupV1Input) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddUsersToUserGroupV1Input(
	val *AddUsersToUserGroupV1Input,
) *NullableAddUsersToUserGroupV1Input {
	return &NullableAddUsersToUserGroupV1Input{value: val, isSet: true}
}

func (v NullableAddUsersToUserGroupV1Input) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddUsersToUserGroupV1Input) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
