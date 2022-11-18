/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API. 

API version: 33.0.1
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// CreateInvitesV1Output Returns the emails of the invited users.
type CreateInvitesV1Output struct {
	// The list of emails invited to the Workspace.
	Emails []string `json:"emails"`
}

// NewCreateInvitesV1Output instantiates a new CreateInvitesV1Output object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateInvitesV1Output(emails []string) *CreateInvitesV1Output {
	this := CreateInvitesV1Output{}
	this.Emails = emails
	return &this
}

// NewCreateInvitesV1OutputWithDefaults instantiates a new CreateInvitesV1Output object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateInvitesV1OutputWithDefaults() *CreateInvitesV1Output {
	this := CreateInvitesV1Output{}
	return &this
}

// GetEmails returns the Emails field value
func (o *CreateInvitesV1Output) GetEmails() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Emails
}

// GetEmailsOk returns a tuple with the Emails field value
// and a boolean to check if the value has been set.
func (o *CreateInvitesV1Output) GetEmailsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Emails, true
}

// SetEmails sets field value
func (o *CreateInvitesV1Output) SetEmails(v []string) {
	o.Emails = v
}

func (o CreateInvitesV1Output) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["emails"] = o.Emails
	}
	return json.Marshal(toSerialize)
}

type NullableCreateInvitesV1Output struct {
	value *CreateInvitesV1Output
	isSet bool
}

func (v NullableCreateInvitesV1Output) Get() *CreateInvitesV1Output {
	return v.value
}

func (v *NullableCreateInvitesV1Output) Set(val *CreateInvitesV1Output) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateInvitesV1Output) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateInvitesV1Output) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateInvitesV1Output(val *CreateInvitesV1Output) *NullableCreateInvitesV1Output {
	return &NullableCreateInvitesV1Output{value: val, isSet: true}
}

func (v NullableCreateInvitesV1Output) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateInvitesV1Output) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


