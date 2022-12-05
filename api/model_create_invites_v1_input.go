/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 33.0.3
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// CreateInvitesV1Input Invites a user to a Workspace with specified permissions.
type CreateInvitesV1Input struct {
	// The list of invites.
	Invites []InviteV1 `json:"invites"`
}

// NewCreateInvitesV1Input instantiates a new CreateInvitesV1Input object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateInvitesV1Input(invites []InviteV1) *CreateInvitesV1Input {
	this := CreateInvitesV1Input{}
	this.Invites = invites
	return &this
}

// NewCreateInvitesV1InputWithDefaults instantiates a new CreateInvitesV1Input object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateInvitesV1InputWithDefaults() *CreateInvitesV1Input {
	this := CreateInvitesV1Input{}
	return &this
}

// GetInvites returns the Invites field value
func (o *CreateInvitesV1Input) GetInvites() []InviteV1 {
	if o == nil {
		var ret []InviteV1
		return ret
	}

	return o.Invites
}

// GetInvitesOk returns a tuple with the Invites field value
// and a boolean to check if the value has been set.
func (o *CreateInvitesV1Input) GetInvitesOk() ([]InviteV1, bool) {
	if o == nil {
		return nil, false
	}
	return o.Invites, true
}

// SetInvites sets field value
func (o *CreateInvitesV1Input) SetInvites(v []InviteV1) {
	o.Invites = v
}

func (o CreateInvitesV1Input) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["invites"] = o.Invites
	}
	return json.Marshal(toSerialize)
}

type NullableCreateInvitesV1Input struct {
	value *CreateInvitesV1Input
	isSet bool
}

func (v NullableCreateInvitesV1Input) Get() *CreateInvitesV1Input {
	return v.value
}

func (v *NullableCreateInvitesV1Input) Set(val *CreateInvitesV1Input) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateInvitesV1Input) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateInvitesV1Input) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateInvitesV1Input(val *CreateInvitesV1Input) *NullableCreateInvitesV1Input {
	return &NullableCreateInvitesV1Input{value: val, isSet: true}
}

func (v NullableCreateInvitesV1Input) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateInvitesV1Input) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
