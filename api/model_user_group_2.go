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

// UserGroup2 The user group returned from the query.
type UserGroup2 struct {
	// The number of members in the user group.
	MemberCount float32 `json:"memberCount"`
	// The permissions associated with the user group.
	Permissions []PermissionV1 `json:"permissions,omitempty"`
	// The id of the user group.
	Id string `json:"id"`
	// The name of the user group.
	Name string `json:"name"`
}

// NewUserGroup2 instantiates a new UserGroup2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserGroup2(memberCount float32, id string, name string) *UserGroup2 {
	this := UserGroup2{}
	this.MemberCount = memberCount
	this.Id = id
	this.Name = name
	return &this
}

// NewUserGroup2WithDefaults instantiates a new UserGroup2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserGroup2WithDefaults() *UserGroup2 {
	this := UserGroup2{}
	return &this
}

// GetMemberCount returns the MemberCount field value
func (o *UserGroup2) GetMemberCount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.MemberCount
}

// GetMemberCountOk returns a tuple with the MemberCount field value
// and a boolean to check if the value has been set.
func (o *UserGroup2) GetMemberCountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MemberCount, true
}

// SetMemberCount sets field value
func (o *UserGroup2) SetMemberCount(v float32) {
	o.MemberCount = v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *UserGroup2) GetPermissions() []PermissionV1 {
	if o == nil || o.Permissions == nil {
		var ret []PermissionV1
		return ret
	}
	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserGroup2) GetPermissionsOk() ([]PermissionV1, bool) {
	if o == nil || o.Permissions == nil {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *UserGroup2) HasPermissions() bool {
	if o != nil && o.Permissions != nil {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given []PermissionV1 and assigns it to the Permissions field.
func (o *UserGroup2) SetPermissions(v []PermissionV1) {
	o.Permissions = v
}

// GetId returns the Id field value
func (o *UserGroup2) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UserGroup2) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UserGroup2) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *UserGroup2) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UserGroup2) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UserGroup2) SetName(v string) {
	o.Name = v
}

func (o UserGroup2) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["memberCount"] = o.MemberCount
	}
	if o.Permissions != nil {
		toSerialize["permissions"] = o.Permissions
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableUserGroup2 struct {
	value *UserGroup2
	isSet bool
}

func (v NullableUserGroup2) Get() *UserGroup2 {
	return v.value
}

func (v *NullableUserGroup2) Set(val *UserGroup2) {
	v.value = val
	v.isSet = true
}

func (v NullableUserGroup2) IsSet() bool {
	return v.isSet
}

func (v *NullableUserGroup2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserGroup2(val *UserGroup2) *NullableUserGroup2 {
	return &NullableUserGroup2{value: val, isSet: true}
}

func (v NullableUserGroup2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserGroup2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
