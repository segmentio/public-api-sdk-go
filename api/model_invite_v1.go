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

// InviteV1 Defines an invitation to join a Workspace.
type InviteV1 struct {
	// The invited user's email to attach the permissions to.
	Email string `json:"email"`
	// The permissions to attach to the invited user.
	Permissions []InvitePermissionV1 `json:"permissions,omitempty"`
}

// NewInviteV1 instantiates a new InviteV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInviteV1(email string) *InviteV1 {
	this := InviteV1{}
	this.Email = email
	return &this
}

// NewInviteV1WithDefaults instantiates a new InviteV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInviteV1WithDefaults() *InviteV1 {
	this := InviteV1{}
	return &this
}

// GetEmail returns the Email field value
func (o *InviteV1) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *InviteV1) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *InviteV1) SetEmail(v string) {
	o.Email = v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *InviteV1) GetPermissions() []InvitePermissionV1 {
	if o == nil || o.Permissions == nil {
		var ret []InvitePermissionV1
		return ret
	}
	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InviteV1) GetPermissionsOk() ([]InvitePermissionV1, bool) {
	if o == nil || o.Permissions == nil {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *InviteV1) HasPermissions() bool {
	if o != nil && o.Permissions != nil {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given []InvitePermissionV1 and assigns it to the Permissions field.
func (o *InviteV1) SetPermissions(v []InvitePermissionV1) {
	o.Permissions = v
}

func (o InviteV1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["email"] = o.Email
	}
	if o.Permissions != nil {
		toSerialize["permissions"] = o.Permissions
	}
	return json.Marshal(toSerialize)
}

type NullableInviteV1 struct {
	value *InviteV1
	isSet bool
}

func (v NullableInviteV1) Get() *InviteV1 {
	return v.value
}

func (v *NullableInviteV1) Set(val *InviteV1) {
	v.value = val
	v.isSet = true
}

func (v NullableInviteV1) IsSet() bool {
	return v.isSet
}

func (v *NullableInviteV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInviteV1(val *InviteV1) *NullableInviteV1 {
	return &NullableInviteV1{value: val, isSet: true}
}

func (v NullableInviteV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInviteV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
