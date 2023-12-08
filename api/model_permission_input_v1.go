/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 38.4.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the PermissionInputV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PermissionInputV1{}

// PermissionInputV1 The input for a permission, associated with a resource and/or labels.
type PermissionInputV1 struct {
	// The role id of this permission.
	RoleId string `json:"roleId"`
	// The resources to link the selected role to.
	Resources []PermissionResourceV1 `json:"resources"`
}

// NewPermissionInputV1 instantiates a new PermissionInputV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPermissionInputV1(roleId string, resources []PermissionResourceV1) *PermissionInputV1 {
	this := PermissionInputV1{}
	this.RoleId = roleId
	this.Resources = resources
	return &this
}

// NewPermissionInputV1WithDefaults instantiates a new PermissionInputV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPermissionInputV1WithDefaults() *PermissionInputV1 {
	this := PermissionInputV1{}
	return &this
}

// GetRoleId returns the RoleId field value
func (o *PermissionInputV1) GetRoleId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RoleId
}

// GetRoleIdOk returns a tuple with the RoleId field value
// and a boolean to check if the value has been set.
func (o *PermissionInputV1) GetRoleIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RoleId, true
}

// SetRoleId sets field value
func (o *PermissionInputV1) SetRoleId(v string) {
	o.RoleId = v
}

// GetResources returns the Resources field value
func (o *PermissionInputV1) GetResources() []PermissionResourceV1 {
	if o == nil {
		var ret []PermissionResourceV1
		return ret
	}

	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value
// and a boolean to check if the value has been set.
func (o *PermissionInputV1) GetResourcesOk() ([]PermissionResourceV1, bool) {
	if o == nil {
		return nil, false
	}
	return o.Resources, true
}

// SetResources sets field value
func (o *PermissionInputV1) SetResources(v []PermissionResourceV1) {
	o.Resources = v
}

func (o PermissionInputV1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PermissionInputV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["roleId"] = o.RoleId
	toSerialize["resources"] = o.Resources
	return toSerialize, nil
}

type NullablePermissionInputV1 struct {
	value *PermissionInputV1
	isSet bool
}

func (v NullablePermissionInputV1) Get() *PermissionInputV1 {
	return v.value
}

func (v *NullablePermissionInputV1) Set(val *PermissionInputV1) {
	v.value = val
	v.isSet = true
}

func (v NullablePermissionInputV1) IsSet() bool {
	return v.isSet
}

func (v *NullablePermissionInputV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePermissionInputV1(val *PermissionInputV1) *NullablePermissionInputV1 {
	return &NullablePermissionInputV1{value: val, isSet: true}
}

func (v NullablePermissionInputV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePermissionInputV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
