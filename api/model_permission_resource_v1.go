/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 43.0.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the PermissionResourceV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PermissionResourceV1{}

// PermissionResourceV1 The most basic representation of a resource belonging to a set of permissions.
type PermissionResourceV1 struct {
	// The id of this resource.
	Id string `json:"id"`
	// The type for this resource.
	Type string `json:"type"`
	// The labels that further refine access to this resource. Labels are exclusive to Workspace-level permissions.
	Labels []AllowedLabelBeta `json:"labels,omitempty"`
}

// NewPermissionResourceV1 instantiates a new PermissionResourceV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPermissionResourceV1(id string, type_ string) *PermissionResourceV1 {
	this := PermissionResourceV1{}
	this.Id = id
	this.Type = type_
	return &this
}

// NewPermissionResourceV1WithDefaults instantiates a new PermissionResourceV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPermissionResourceV1WithDefaults() *PermissionResourceV1 {
	this := PermissionResourceV1{}
	return &this
}

// GetId returns the Id field value
func (o *PermissionResourceV1) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PermissionResourceV1) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PermissionResourceV1) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value
func (o *PermissionResourceV1) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PermissionResourceV1) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PermissionResourceV1) SetType(v string) {
	o.Type = v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *PermissionResourceV1) GetLabels() []AllowedLabelBeta {
	if o == nil || IsNil(o.Labels) {
		var ret []AllowedLabelBeta
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionResourceV1) GetLabelsOk() ([]AllowedLabelBeta, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *PermissionResourceV1) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []AllowedLabelBeta and assigns it to the Labels field.
func (o *PermissionResourceV1) SetLabels(v []AllowedLabelBeta) {
	o.Labels = v
}

func (o PermissionResourceV1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PermissionResourceV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	return toSerialize, nil
}

type NullablePermissionResourceV1 struct {
	value *PermissionResourceV1
	isSet bool
}

func (v NullablePermissionResourceV1) Get() *PermissionResourceV1 {
	return v.value
}

func (v *NullablePermissionResourceV1) Set(val *PermissionResourceV1) {
	v.value = val
	v.isSet = true
}

func (v NullablePermissionResourceV1) IsSet() bool {
	return v.isSet
}

func (v *NullablePermissionResourceV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePermissionResourceV1(val *PermissionResourceV1) *NullablePermissionResourceV1 {
	return &NullablePermissionResourceV1{value: val, isSet: true}
}

func (v NullablePermissionResourceV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePermissionResourceV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
