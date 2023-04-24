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

// PermissionV1 A registered set of permissions for a subject, extending a role to a resource.
type PermissionV1 struct {
	// The name of the role associated with this permission.
	RoleName string `json:"roleName"`
	// The id of the role associated with this permission.
	RoleId string `json:"roleId"`
	// The resources associated with this permission.
	Resources []PermissionResourceV1 `json:"resources"`
	// The labels to attach to this permission.
	Labels []AllowedLabelBeta `json:"labels,omitempty"`
}

// NewPermissionV1 instantiates a new PermissionV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPermissionV1(
	roleName string,
	roleId string,
	resources []PermissionResourceV1,
) *PermissionV1 {
	this := PermissionV1{}
	this.RoleName = roleName
	this.RoleId = roleId
	this.Resources = resources
	return &this
}

// NewPermissionV1WithDefaults instantiates a new PermissionV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPermissionV1WithDefaults() *PermissionV1 {
	this := PermissionV1{}
	return &this
}

// GetRoleName returns the RoleName field value
func (o *PermissionV1) GetRoleName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RoleName
}

// GetRoleNameOk returns a tuple with the RoleName field value
// and a boolean to check if the value has been set.
func (o *PermissionV1) GetRoleNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RoleName, true
}

// SetRoleName sets field value
func (o *PermissionV1) SetRoleName(v string) {
	o.RoleName = v
}

// GetRoleId returns the RoleId field value
func (o *PermissionV1) GetRoleId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RoleId
}

// GetRoleIdOk returns a tuple with the RoleId field value
// and a boolean to check if the value has been set.
func (o *PermissionV1) GetRoleIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RoleId, true
}

// SetRoleId sets field value
func (o *PermissionV1) SetRoleId(v string) {
	o.RoleId = v
}

// GetResources returns the Resources field value
func (o *PermissionV1) GetResources() []PermissionResourceV1 {
	if o == nil {
		var ret []PermissionResourceV1
		return ret
	}

	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value
// and a boolean to check if the value has been set.
func (o *PermissionV1) GetResourcesOk() ([]PermissionResourceV1, bool) {
	if o == nil {
		return nil, false
	}
	return o.Resources, true
}

// SetResources sets field value
func (o *PermissionV1) SetResources(v []PermissionResourceV1) {
	o.Resources = v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *PermissionV1) GetLabels() []AllowedLabelBeta {
	if o == nil || o.Labels == nil {
		var ret []AllowedLabelBeta
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionV1) GetLabelsOk() ([]AllowedLabelBeta, bool) {
	if o == nil || o.Labels == nil {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *PermissionV1) HasLabels() bool {
	if o != nil && o.Labels != nil {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []AllowedLabelBeta and assigns it to the Labels field.
func (o *PermissionV1) SetLabels(v []AllowedLabelBeta) {
	o.Labels = v
}

func (o PermissionV1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["roleName"] = o.RoleName
	}
	if true {
		toSerialize["roleId"] = o.RoleId
	}
	if true {
		toSerialize["resources"] = o.Resources
	}
	if o.Labels != nil {
		toSerialize["labels"] = o.Labels
	}
	return json.Marshal(toSerialize)
}

type NullablePermissionV1 struct {
	value *PermissionV1
	isSet bool
}

func (v NullablePermissionV1) Get() *PermissionV1 {
	return v.value
}

func (v *NullablePermissionV1) Set(val *PermissionV1) {
	v.value = val
	v.isSet = true
}

func (v NullablePermissionV1) IsSet() bool {
	return v.isSet
}

func (v *NullablePermissionV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePermissionV1(val *PermissionV1) *NullablePermissionV1 {
	return &NullablePermissionV1{value: val, isSet: true}
}

func (v NullablePermissionV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePermissionV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
