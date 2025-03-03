/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 57.4.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the AccessPermissionV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccessPermissionV1{}

// AccessPermissionV1 A permission governing access to a resource.
type AccessPermissionV1 struct {
	// The id of the role that applies to this permission.
	RoleId string `json:"roleId"`
	// The name of the role that applies to this permission.
	RoleName string `json:"roleName"`
	// The resources included with this permission.
	Resources []PermissionResourceV1 `json:"resources"`
}

// NewAccessPermissionV1 instantiates a new AccessPermissionV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessPermissionV1(
	roleId string,
	roleName string,
	resources []PermissionResourceV1,
) *AccessPermissionV1 {
	this := AccessPermissionV1{}
	this.RoleId = roleId
	this.RoleName = roleName
	this.Resources = resources
	return &this
}

// NewAccessPermissionV1WithDefaults instantiates a new AccessPermissionV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessPermissionV1WithDefaults() *AccessPermissionV1 {
	this := AccessPermissionV1{}
	return &this
}

// GetRoleId returns the RoleId field value
func (o *AccessPermissionV1) GetRoleId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RoleId
}

// GetRoleIdOk returns a tuple with the RoleId field value
// and a boolean to check if the value has been set.
func (o *AccessPermissionV1) GetRoleIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RoleId, true
}

// SetRoleId sets field value
func (o *AccessPermissionV1) SetRoleId(v string) {
	o.RoleId = v
}

// GetRoleName returns the RoleName field value
func (o *AccessPermissionV1) GetRoleName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RoleName
}

// GetRoleNameOk returns a tuple with the RoleName field value
// and a boolean to check if the value has been set.
func (o *AccessPermissionV1) GetRoleNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RoleName, true
}

// SetRoleName sets field value
func (o *AccessPermissionV1) SetRoleName(v string) {
	o.RoleName = v
}

// GetResources returns the Resources field value
func (o *AccessPermissionV1) GetResources() []PermissionResourceV1 {
	if o == nil {
		var ret []PermissionResourceV1
		return ret
	}

	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value
// and a boolean to check if the value has been set.
func (o *AccessPermissionV1) GetResourcesOk() ([]PermissionResourceV1, bool) {
	if o == nil {
		return nil, false
	}
	return o.Resources, true
}

// SetResources sets field value
func (o *AccessPermissionV1) SetResources(v []PermissionResourceV1) {
	o.Resources = v
}

func (o AccessPermissionV1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccessPermissionV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["roleId"] = o.RoleId
	toSerialize["roleName"] = o.RoleName
	toSerialize["resources"] = o.Resources
	return toSerialize, nil
}

type NullableAccessPermissionV1 struct {
	value *AccessPermissionV1
	isSet bool
}

func (v NullableAccessPermissionV1) Get() *AccessPermissionV1 {
	return v.value
}

func (v *NullableAccessPermissionV1) Set(val *AccessPermissionV1) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessPermissionV1) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessPermissionV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessPermissionV1(val *AccessPermissionV1) *NullableAccessPermissionV1 {
	return &NullableAccessPermissionV1{value: val, isSet: true}
}

func (v NullableAccessPermissionV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessPermissionV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
