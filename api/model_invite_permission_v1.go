/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 54.1.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the InvitePermissionV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InvitePermissionV1{}

// InvitePermissionV1 Defines a permission to apply to the user in an invite.
type InvitePermissionV1 struct {
	// The id of the role.
	RoleId string `json:"roleId"`
	// The resources to grant the invited users access to.
	Resources []ResourceV1 `json:"resources,omitempty"`
	// The labels that determine which resources to grant users access to.
	Labels []AllowedLabelBeta `json:"labels,omitempty"`
}

// NewInvitePermissionV1 instantiates a new InvitePermissionV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvitePermissionV1(roleId string) *InvitePermissionV1 {
	this := InvitePermissionV1{}
	this.RoleId = roleId
	return &this
}

// NewInvitePermissionV1WithDefaults instantiates a new InvitePermissionV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvitePermissionV1WithDefaults() *InvitePermissionV1 {
	this := InvitePermissionV1{}
	return &this
}

// GetRoleId returns the RoleId field value
func (o *InvitePermissionV1) GetRoleId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RoleId
}

// GetRoleIdOk returns a tuple with the RoleId field value
// and a boolean to check if the value has been set.
func (o *InvitePermissionV1) GetRoleIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RoleId, true
}

// SetRoleId sets field value
func (o *InvitePermissionV1) SetRoleId(v string) {
	o.RoleId = v
}

// GetResources returns the Resources field value if set, zero value otherwise.
func (o *InvitePermissionV1) GetResources() []ResourceV1 {
	if o == nil || IsNil(o.Resources) {
		var ret []ResourceV1
		return ret
	}
	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvitePermissionV1) GetResourcesOk() ([]ResourceV1, bool) {
	if o == nil || IsNil(o.Resources) {
		return nil, false
	}
	return o.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *InvitePermissionV1) HasResources() bool {
	if o != nil && !IsNil(o.Resources) {
		return true
	}

	return false
}

// SetResources gets a reference to the given []ResourceV1 and assigns it to the Resources field.
func (o *InvitePermissionV1) SetResources(v []ResourceV1) {
	o.Resources = v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *InvitePermissionV1) GetLabels() []AllowedLabelBeta {
	if o == nil || IsNil(o.Labels) {
		var ret []AllowedLabelBeta
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvitePermissionV1) GetLabelsOk() ([]AllowedLabelBeta, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *InvitePermissionV1) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []AllowedLabelBeta and assigns it to the Labels field.
func (o *InvitePermissionV1) SetLabels(v []AllowedLabelBeta) {
	o.Labels = v
}

func (o InvitePermissionV1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InvitePermissionV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["roleId"] = o.RoleId
	if !IsNil(o.Resources) {
		toSerialize["resources"] = o.Resources
	}
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	return toSerialize, nil
}

type NullableInvitePermissionV1 struct {
	value *InvitePermissionV1
	isSet bool
}

func (v NullableInvitePermissionV1) Get() *InvitePermissionV1 {
	return v.value
}

func (v *NullableInvitePermissionV1) Set(val *InvitePermissionV1) {
	v.value = val
	v.isSet = true
}

func (v NullableInvitePermissionV1) IsSet() bool {
	return v.isSet
}

func (v *NullableInvitePermissionV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvitePermissionV1(val *InvitePermissionV1) *NullableInvitePermissionV1 {
	return &NullableInvitePermissionV1{value: val, isSet: true}
}

func (v NullableInvitePermissionV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvitePermissionV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
