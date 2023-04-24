/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 34.3.5
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// UserV1 A user belonging to a Segment Workspace.
type UserV1 struct {
	// The unique identifier of this user.  Config API note: analogous to `name`.
	Id string `json:"id"`
	// The human-readable name of this user.
	Name string `json:"name"`
	// The email address associated with this user.
	Email string `json:"email"`
	// The permissions associated with this user.
	Permissions []PermissionV1 `json:"permissions,omitempty"`
}

// NewUserV1 instantiates a new UserV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserV1(id string, name string, email string) *UserV1 {
	this := UserV1{}
	this.Id = id
	this.Name = name
	this.Email = email
	return &this
}

// NewUserV1WithDefaults instantiates a new UserV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserV1WithDefaults() *UserV1 {
	this := UserV1{}
	return &this
}

// GetId returns the Id field value
func (o *UserV1) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UserV1) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UserV1) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *UserV1) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UserV1) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UserV1) SetName(v string) {
	o.Name = v
}

// GetEmail returns the Email field value
func (o *UserV1) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *UserV1) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *UserV1) SetEmail(v string) {
	o.Email = v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *UserV1) GetPermissions() []PermissionV1 {
	if o == nil || o.Permissions == nil {
		var ret []PermissionV1
		return ret
	}
	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserV1) GetPermissionsOk() ([]PermissionV1, bool) {
	if o == nil || o.Permissions == nil {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *UserV1) HasPermissions() bool {
	if o != nil && o.Permissions != nil {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given []PermissionV1 and assigns it to the Permissions field.
func (o *UserV1) SetPermissions(v []PermissionV1) {
	o.Permissions = v
}

func (o UserV1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["email"] = o.Email
	}
	if o.Permissions != nil {
		toSerialize["permissions"] = o.Permissions
	}
	return json.Marshal(toSerialize)
}

type NullableUserV1 struct {
	value *UserV1
	isSet bool
}

func (v NullableUserV1) Get() *UserV1 {
	return v.value
}

func (v *NullableUserV1) Set(val *UserV1) {
	v.value = val
	v.isSet = true
}

func (v NullableUserV1) IsSet() bool {
	return v.isSet
}

func (v *NullableUserV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserV1(val *UserV1) *NullableUserV1 {
	return &NullableUserV1{value: val, isSet: true}
}

func (v NullableUserV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
