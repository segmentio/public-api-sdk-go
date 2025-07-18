/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 58.13.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the ReplaceUsersInUserGroupV1Input type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReplaceUsersInUserGroupV1Input{}

// ReplaceUsersInUserGroupV1Input Replace a user group's list of users and invites with a new one.
type ReplaceUsersInUserGroupV1Input struct {
	// The email addresses of the users and invites to replace.
	Emails []string `json:"emails"`
}

// NewReplaceUsersInUserGroupV1Input instantiates a new ReplaceUsersInUserGroupV1Input object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReplaceUsersInUserGroupV1Input(emails []string) *ReplaceUsersInUserGroupV1Input {
	this := ReplaceUsersInUserGroupV1Input{}
	this.Emails = emails
	return &this
}

// NewReplaceUsersInUserGroupV1InputWithDefaults instantiates a new ReplaceUsersInUserGroupV1Input object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReplaceUsersInUserGroupV1InputWithDefaults() *ReplaceUsersInUserGroupV1Input {
	this := ReplaceUsersInUserGroupV1Input{}
	return &this
}

// GetEmails returns the Emails field value
func (o *ReplaceUsersInUserGroupV1Input) GetEmails() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Emails
}

// GetEmailsOk returns a tuple with the Emails field value
// and a boolean to check if the value has been set.
func (o *ReplaceUsersInUserGroupV1Input) GetEmailsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Emails, true
}

// SetEmails sets field value
func (o *ReplaceUsersInUserGroupV1Input) SetEmails(v []string) {
	o.Emails = v
}

func (o ReplaceUsersInUserGroupV1Input) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReplaceUsersInUserGroupV1Input) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["emails"] = o.Emails
	return toSerialize, nil
}

type NullableReplaceUsersInUserGroupV1Input struct {
	value *ReplaceUsersInUserGroupV1Input
	isSet bool
}

func (v NullableReplaceUsersInUserGroupV1Input) Get() *ReplaceUsersInUserGroupV1Input {
	return v.value
}

func (v *NullableReplaceUsersInUserGroupV1Input) Set(val *ReplaceUsersInUserGroupV1Input) {
	v.value = val
	v.isSet = true
}

func (v NullableReplaceUsersInUserGroupV1Input) IsSet() bool {
	return v.isSet
}

func (v *NullableReplaceUsersInUserGroupV1Input) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReplaceUsersInUserGroupV1Input(
	val *ReplaceUsersInUserGroupV1Input,
) *NullableReplaceUsersInUserGroupV1Input {
	return &NullableReplaceUsersInUserGroupV1Input{value: val, isSet: true}
}

func (v NullableReplaceUsersInUserGroupV1Input) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReplaceUsersInUserGroupV1Input) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
