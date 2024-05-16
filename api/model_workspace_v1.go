/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 50.1.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the WorkspaceV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkspaceV1{}

// WorkspaceV1 An organized group of Sources and Destinations managed by a team.
type WorkspaceV1 struct {
	// The unique identifier.
	Id string `json:"id"`
	// The URL-friendly slug.
	Slug string `json:"slug"`
	// The human-readable name.
	Name string `json:"name"`
}

// NewWorkspaceV1 instantiates a new WorkspaceV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkspaceV1(id string, slug string, name string) *WorkspaceV1 {
	this := WorkspaceV1{}
	this.Id = id
	this.Slug = slug
	this.Name = name
	return &this
}

// NewWorkspaceV1WithDefaults instantiates a new WorkspaceV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkspaceV1WithDefaults() *WorkspaceV1 {
	this := WorkspaceV1{}
	return &this
}

// GetId returns the Id field value
func (o *WorkspaceV1) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *WorkspaceV1) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *WorkspaceV1) SetId(v string) {
	o.Id = v
}

// GetSlug returns the Slug field value
func (o *WorkspaceV1) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *WorkspaceV1) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *WorkspaceV1) SetSlug(v string) {
	o.Slug = v
}

// GetName returns the Name field value
func (o *WorkspaceV1) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *WorkspaceV1) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *WorkspaceV1) SetName(v string) {
	o.Name = v
}

func (o WorkspaceV1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkspaceV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["slug"] = o.Slug
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullableWorkspaceV1 struct {
	value *WorkspaceV1
	isSet bool
}

func (v NullableWorkspaceV1) Get() *WorkspaceV1 {
	return v.value
}

func (v *NullableWorkspaceV1) Set(val *WorkspaceV1) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkspaceV1) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkspaceV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkspaceV1(val *WorkspaceV1) *NullableWorkspaceV1 {
	return &NullableWorkspaceV1{value: val, isSet: true}
}

func (v NullableWorkspaceV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkspaceV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
